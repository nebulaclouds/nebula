package bigquery

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/api/bigquery/v2"
	"google.golang.org/api/googleapi"
	"k8s.io/apimachinery/pkg/util/rand"

	nebulaIdlCore "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/core"
	pluginsCore "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/core"
	coreMocks "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/core/mocks"
	"github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/io"
	ioMocks "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/io/mocks"
	"github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/webapi/mocks"
	"github.com/nebulaclouds/nebula/nebulastdlib/contextutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils/labeled"
	"github.com/nebulaclouds/nebula/nebulastdlib/storage"
)

func init() {
	labeled.SetMetricKeys(contextutils.NamespaceKey)
}

func TestFormatJobReference(t *testing.T) {
	t.Run("format job reference", func(t *testing.T) {
		jobReference := bigquery.JobReference{
			JobId:     "my-job-id",
			Location:  "EU",
			ProjectId: "nebula-test",
		}

		str := formatJobReference(jobReference)

		assert.Equal(t, "nebula-test:EU.my-job-id", str)
	})
}

func TestConstructOutputLocation(t *testing.T) {
	job := &bigquery.Job{
		Configuration: &bigquery.JobConfiguration{
			Query: &bigquery.JobConfigurationQuery{
				DestinationTable: &bigquery.TableReference{
					ProjectId: "project",
					DatasetId: "dataset",
					TableId:   "table",
				},
			},
		},
	}
	ol := constructOutputLocation(context.Background(), job)
	assert.Equal(t, ol, "bq://project:dataset.table")

	job.Configuration.Query.DestinationTable = nil
	ol = constructOutputLocation(context.Background(), job)
	assert.Equal(t, ol, "")
}

func TestCreateTaskInfo(t *testing.T) {
	t.Run("create task info", func(t *testing.T) {
		resourceMeta := ResourceMetaWrapper{
			JobReference: bigquery.JobReference{
				JobId:     "my-job-id",
				Location:  "EU",
				ProjectId: "nebula-test",
			},
		}

		taskInfo := createTaskInfo(&resourceMeta)

		assert.Equal(t, 1, len(taskInfo.Logs))
		assert.Equal(t, nebulaIdlCore.TaskLog{
			Uri:  "https://console.cloud.google.com/bigquery?project=nebula-test&j=bq:EU:my-job-id&page=queryresults",
			Name: "BigQuery Console",
		}, *taskInfo.Logs[0])
	})
}

func TestOutputWriter(t *testing.T) {
	ctx := context.Background()
	statusContext := &mocks.StatusContext{}

	template := nebulaIdlCore.TaskTemplate{}
	tr := &coreMocks.TaskReader{}
	tr.OnRead(ctx).Return(&template, nil)
	statusContext.OnTaskReader().Return(tr)

	outputLocation := "bq://project:nebula.table"
	err := writeOutput(ctx, statusContext, outputLocation)
	assert.NoError(t, err)

	ds, err := storage.NewDataStore(&storage.Config{Type: storage.TypeMemory}, promutils.NewTestScope())
	assert.NoError(t, err)

	outputWriter := &ioMocks.OutputWriter{}
	outputWriter.OnPutMatch(mock.Anything, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		or := args.Get(1).(io.OutputReader)
		literals, ee, err := or.Read(ctx)
		assert.NoError(t, err)

		sd := literals.GetLiterals()["results"].GetScalar().GetStructuredDataset()
		assert.Equal(t, sd.Uri, outputLocation)
		assert.Equal(t, sd.Metadata.GetStructuredDatasetType().Columns[0].Name, "col1")
		assert.Equal(t, sd.Metadata.GetStructuredDatasetType().Columns[0].LiteralType.GetSimple(), nebulaIdlCore.SimpleType_INTEGER)

		if ee != nil {
			assert.NoError(t, ds.WriteProtobuf(ctx, outputWriter.GetErrorPath(), storage.Options{}, ee))
		}

		if literals != nil {
			assert.NoError(t, ds.WriteProtobuf(ctx, outputWriter.GetOutputPath(), storage.Options{}, literals))
		}
	})

	execID := rand.String(3)
	basePrefix := storage.DataReference("fake://bucket/prefix/" + execID)
	outputWriter.OnGetOutputPath().Return(basePrefix + "/outputs.pb")
	statusContext.OnOutputWriter().Return(outputWriter)

	template = nebulaIdlCore.TaskTemplate{
		Interface: &nebulaIdlCore.TypedInterface{
			Outputs: &nebulaIdlCore.VariableMap{
				Variables: map[string]*nebulaIdlCore.Variable{
					"results": {
						Type: &nebulaIdlCore.LiteralType{
							Type: &nebulaIdlCore.LiteralType_StructuredDatasetType{
								StructuredDatasetType: &nebulaIdlCore.StructuredDatasetType{
									Columns: []*nebulaIdlCore.StructuredDatasetType_DatasetColumn{
										{
											Name: "col1",
											LiteralType: &nebulaIdlCore.LiteralType{
												Type: &nebulaIdlCore.LiteralType_Simple{
													Simple: nebulaIdlCore.SimpleType_INTEGER,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	tr.OnRead(ctx).Return(&template, nil)
	statusContext.OnTaskReader().Return(tr)
	err = writeOutput(ctx, statusContext, outputLocation)
	assert.NoError(t, err)
}

func TestHandleCreateError(t *testing.T) {
	occurredAt := time.Now()
	taskInfo := core.TaskInfo{OccurredAt: &occurredAt}

	t.Run("handle 401", func(t *testing.T) {
		createError := googleapi.Error{
			Code:    401,
			Message: "user xxx is not authorized",
		}

		phase := handleCreateError(&createError, &taskInfo)

		assert.Equal(t, nebulaIdlCore.ExecutionError{
			Code:    "http401",
			Message: "user xxx is not authorized",
			Kind:    nebulaIdlCore.ExecutionError_USER,
		}, *phase.Err())
		assert.Equal(t, taskInfo, *phase.Info())
	})

	t.Run("handle 500", func(t *testing.T) {
		createError := googleapi.Error{
			Code:    500,
			Message: "oops",
		}

		phase := handleCreateError(&createError, &taskInfo)

		assert.Equal(t, nebulaIdlCore.ExecutionError{
			Code:    "http500",
			Message: "oops",
			Kind:    nebulaIdlCore.ExecutionError_SYSTEM,
		}, *phase.Err())
		assert.Equal(t, taskInfo, *phase.Info())
	})
}

func TestHandleErrorResult(t *testing.T) {
	occurredAt := time.Now()
	taskInfo := core.TaskInfo{OccurredAt: &occurredAt}

	type args struct {
		reason    string
		phase     core.Phase
		errorKind nebulaIdlCore.ExecutionError_ErrorKind
	}

	tests := []args{
		{
			reason:    "accessDenied",
			phase:     pluginsCore.PhasePermanentFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},
		{
			reason:    "backendError",
			phase:     pluginsCore.PhaseRetryableFailure,
			errorKind: nebulaIdlCore.ExecutionError_SYSTEM,
		},
		{
			reason:    "billingNotEnabled",
			phase:     pluginsCore.PhasePermanentFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},
		{
			reason:    "blocked",
			phase:     pluginsCore.PhasePermanentFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},
		{
			reason:    "duplicate",
			phase:     pluginsCore.PhasePermanentFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},
		{
			reason:    "internalError",
			phase:     pluginsCore.PhaseRetryableFailure,
			errorKind: nebulaIdlCore.ExecutionError_SYSTEM,
		},
		{
			reason:    "invalid",
			phase:     pluginsCore.PhasePermanentFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},
		{
			reason:    "invalidQuery",
			phase:     pluginsCore.PhasePermanentFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},
		{
			reason:    "invalidUser",
			phase:     pluginsCore.PhaseRetryableFailure,
			errorKind: nebulaIdlCore.ExecutionError_SYSTEM,
		},
		{
			reason:    "notFound",
			phase:     pluginsCore.PhasePermanentFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},
		{
			reason:    "notImplemented",
			phase:     pluginsCore.PhasePermanentFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},
		{
			reason:    "quotaExceeded",
			phase:     pluginsCore.PhaseRetryableFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},
		{
			reason:    "rateLimitExceeded",
			phase:     pluginsCore.PhaseRetryableFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},
		{
			reason:    "resourceInUse",
			phase:     pluginsCore.PhaseRetryableFailure,
			errorKind: nebulaIdlCore.ExecutionError_SYSTEM,
		},

		{
			reason:    "resourcesExceeded",
			phase:     pluginsCore.PhasePermanentFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},

		{
			reason:    "responseTooLarge",
			phase:     pluginsCore.PhasePermanentFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},
		{
			reason:    "stopped",
			phase:     pluginsCore.PhasePermanentFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},
		{
			reason:    "tableUnavailable",
			phase:     pluginsCore.PhaseRetryableFailure,
			errorKind: nebulaIdlCore.ExecutionError_SYSTEM,
		},
		{
			reason:    "timeout",
			phase:     pluginsCore.PhasePermanentFailure,
			errorKind: nebulaIdlCore.ExecutionError_USER,
		},
	}

	for _, test := range tests {
		t.Run(test.reason, func(t *testing.T) {
			phaseInfo := handleErrorResult(test.reason, "message", &taskInfo)

			assert.Equal(t, test.phase, phaseInfo.Phase())
			assert.Equal(t, test.reason, phaseInfo.Err().Code)
			assert.Equal(t, test.errorKind, phaseInfo.Err().Kind)
			assert.Equal(t, "message", phaseInfo.Err().Message)
		})
	}
}
