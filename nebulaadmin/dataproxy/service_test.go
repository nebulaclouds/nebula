package dataproxy

import (
	"bytes"
	"context"
	"crypto/md5" // #nosec
	"net/url"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/durationpb"

	commonMocks "github.com/nebulaclouds/nebula/nebulaadmin/pkg/common/mocks"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/config"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/manager/mocks"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/service"
	stdlibConfig "github.com/nebulaclouds/nebula/nebulastdlib/config"
	"github.com/nebulaclouds/nebula/nebulastdlib/contextutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils"
	"github.com/nebulaclouds/nebula/nebulastdlib/promutils/labeled"
	"github.com/nebulaclouds/nebula/nebulastdlib/storage"
)

func TestNewService(t *testing.T) {
	dataStore, err := storage.NewDataStore(&storage.Config{Type: storage.TypeMemory}, promutils.NewTestScope())
	assert.NoError(t, err)

	nodeExecutionManager := &mocks.MockNodeExecutionManager{}
	taskExecutionManager := &mocks.MockTaskExecutionManager{}
	s, err := NewService(config.DataProxyConfig{
		Upload: config.DataProxyUploadConfig{},
	}, nodeExecutionManager, dataStore, taskExecutionManager)
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func init() {
	labeled.SetMetricKeys(contextutils.DomainKey)
}

func Test_createStorageLocation(t *testing.T) {
	dataStore, err := storage.NewDataStore(&storage.Config{Type: storage.TypeMemory}, promutils.NewTestScope())
	assert.NoError(t, err)
	loc, err := createStorageLocation(context.Background(), dataStore, config.DataProxyUploadConfig{
		StoragePrefix: "blah",
	})
	assert.NoError(t, err)
	assert.Equal(t, "/blah", loc.String())
}

func TestCreateUploadLocation(t *testing.T) {
	dataStore, err := storage.NewDataStore(&storage.Config{Type: storage.TypeMemory}, promutils.NewTestScope())
	assert.NoError(t, err)
	nodeExecutionManager := &mocks.MockNodeExecutionManager{}
	taskExecutionManager := &mocks.MockTaskExecutionManager{}
	s, err := NewService(config.DataProxyConfig{}, nodeExecutionManager, dataStore, taskExecutionManager)
	assert.NoError(t, err)
	t.Run("No project/domain", func(t *testing.T) {
		_, err = s.CreateUploadLocation(context.Background(), &service.CreateUploadLocationRequest{})
		assert.Error(t, err)
	})

	t.Run("unsupported operation by InMemory DataStore", func(t *testing.T) {
		_, err = s.CreateUploadLocation(context.Background(), &service.CreateUploadLocationRequest{
			Project: "hello",
			Domain:  "world",
		})
		assert.Error(t, err)
	})

	t.Run("Invalid expiry", func(t *testing.T) {
		_, err = s.CreateUploadLocation(context.Background(), &service.CreateUploadLocationRequest{
			Project:   "hello",
			Domain:    "world",
			ExpiresIn: durationpb.New(-time.Hour),
		})
		assert.Error(t, err)
	})
}

type UploadableMemProtobufStore struct {
	storage.ComposedProtobufStore
}

func (u *UploadableMemProtobufStore) CreateSignedURL(ctx context.Context, reference storage.DataReference, properties storage.SignedURLProperties) (storage.SignedURLResponse, error) {
	return storage.SignedURLResponse{
		URL: url.URL{
			Scheme: "http",
			Host:   "localhost",
			Path:   "/blah",
		},
	}, nil
}

func TestCreateUploadLocationMore(t *testing.T) {
	dataStore, err := storage.NewDataStore(&storage.Config{Type: storage.TypeMemory}, promutils.NewTestScope())
	// Overwrite the CreateSignedURL function to return a fixed value
	uploadable := &UploadableMemProtobufStore{
		dataStore.ComposedProtobufStore,
	}

	ds := storage.DataStore{
		ComposedProtobufStore: uploadable,
		ReferenceConstructor:  dataStore.ReferenceConstructor,
	}

	assert.NoError(t, err)
	nodeExecutionManager := &mocks.MockNodeExecutionManager{}
	taskExecutionManager := &mocks.MockTaskExecutionManager{}
	s, err := NewService(config.DataProxyConfig{}, nodeExecutionManager, &ds, taskExecutionManager)
	assert.NoError(t, err)

	exists, err := createStorageLocation(context.TODO(), s.dataStore, s.cfg.Upload,
		"nebulasnacks", "development", "known_parent_folder", "myfile.txt")
	assert.NoError(t, err)
	err = dataStore.WriteRaw(context.TODO(), exists, 5, storage.Options{}, bytes.NewReader([]byte("hello")))
	assert.NoError(t, err)

	t.Run("no need for content md5", func(t *testing.T) {
		_, err = s.CreateUploadLocation(context.Background(), &service.CreateUploadLocationRequest{
			Project:      "nebulasnacks",
			Domain:       "development",
			Filename:     "myotherfile.txt",
			ContentMd5:   nil,
			FilenameRoot: "known_parent_folder",
		})
		assert.NoError(t, err)
	})

	t.Run("already exists errors", func(t *testing.T) {
		_, err = s.CreateUploadLocation(context.Background(), &service.CreateUploadLocationRequest{
			Project:      "nebulasnacks",
			Domain:       "development",
			Filename:     "myfile.txt",
			ContentMd5:   nil,
			FilenameRoot: "known_parent_folder",
		})
		assert.ErrorContainsf(t, err, "already exists", "expected error to contain already exists")
	})

	t.Run("already exists but matching content md5", func(t *testing.T) {
		m := md5.Sum([]byte("hello")) // #nosec
		resp, err := s.CreateUploadLocation(context.Background(), &service.CreateUploadLocationRequest{
			Project:      "nebulasnacks",
			Domain:       "development",
			Filename:     "myfile.txt",
			ContentMd5:   m[:],
			FilenameRoot: "known_parent_folder",
		})
		assert.NoError(t, err)
		assert.Equal(t, "/nebulasnacks/development/known_parent_folder/myfile.txt", resp.GetNativeUrl())
	})
}

func TestCreateDownloadLink(t *testing.T) {
	dataStore := commonMocks.GetMockStorageClient()
	nodeExecutionManager := &mocks.MockNodeExecutionManager{}
	nodeExecutionManager.SetGetNodeExecutionFunc(func(ctx context.Context, request admin.NodeExecutionGetRequest) (*admin.NodeExecution, error) {
		return &admin.NodeExecution{
			Closure: &admin.NodeExecutionClosure{
				DeckUri: "s3://something/something",
			},
		}, nil
	})
	taskExecutionManager := &mocks.MockTaskExecutionManager{}

	s, err := NewService(config.DataProxyConfig{Download: config.DataProxyDownloadConfig{MaxExpiresIn: stdlibConfig.Duration{Duration: time.Hour}}}, nodeExecutionManager, dataStore, taskExecutionManager)
	assert.NoError(t, err)

	t.Run("Invalid expiry", func(t *testing.T) {
		_, err = s.CreateDownloadLink(context.Background(), &service.CreateDownloadLinkRequest{
			ExpiresIn: durationpb.New(-time.Hour),
		})
		assert.Error(t, err)
	})

	t.Run("valid config", func(t *testing.T) {
		_, err = s.CreateDownloadLink(context.Background(), &service.CreateDownloadLinkRequest{
			ArtifactType: service.ArtifactType_ARTIFACT_TYPE_DECK,
			Source: &service.CreateDownloadLinkRequest_NodeExecutionId{
				NodeExecutionId: &core.NodeExecutionIdentifier{},
			},
			ExpiresIn: durationpb.New(time.Hour),
		})
		assert.NoError(t, err)
	})

	t.Run("use default ExpiresIn", func(t *testing.T) {
		_, err = s.CreateDownloadLink(context.Background(), &service.CreateDownloadLinkRequest{
			ArtifactType: service.ArtifactType_ARTIFACT_TYPE_DECK,
			Source: &service.CreateDownloadLinkRequest_NodeExecutionId{
				NodeExecutionId: &core.NodeExecutionIdentifier{},
			},
		})
		assert.NoError(t, err)
	})
}

func TestCreateDownloadLocation(t *testing.T) {
	dataStore := commonMocks.GetMockStorageClient()
	nodeExecutionManager := &mocks.MockNodeExecutionManager{}
	taskExecutionManager := &mocks.MockTaskExecutionManager{}
	s, err := NewService(config.DataProxyConfig{Download: config.DataProxyDownloadConfig{MaxExpiresIn: stdlibConfig.Duration{Duration: time.Hour}}}, nodeExecutionManager, dataStore, taskExecutionManager)
	assert.NoError(t, err)

	t.Run("Invalid expiry", func(t *testing.T) {
		_, err = s.CreateDownloadLocation(context.Background(), &service.CreateDownloadLocationRequest{
			NativeUrl: "s3://bucket/key",
			ExpiresIn: durationpb.New(-time.Hour),
		})
		assert.Error(t, err)
	})

	t.Run("valid config", func(t *testing.T) {
		_, err = s.CreateDownloadLocation(context.Background(), &service.CreateDownloadLocationRequest{
			NativeUrl: "s3://bucket/key",
			ExpiresIn: durationpb.New(time.Hour),
		})
		assert.NoError(t, err)
	})

	t.Run("use default ExpiresIn", func(t *testing.T) {
		_, err = s.CreateDownloadLocation(context.Background(), &service.CreateDownloadLocationRequest{
			NativeUrl: "s3://bucket/key",
		})
		assert.NoError(t, err)
	})

	t.Run("invalid URL", func(t *testing.T) {
		_, err = s.CreateDownloadLocation(context.Background(), &service.CreateDownloadLocationRequest{
			NativeUrl: "bucket/key",
		})
		assert.NoError(t, err)
	})
}

func TestService_GetData(t *testing.T) {
	dataStore := commonMocks.GetMockStorageClient()
	nodeExecutionManager := &mocks.MockNodeExecutionManager{}
	taskExecutionManager := &mocks.MockTaskExecutionManager{}
	s, err := NewService(config.DataProxyConfig{}, nodeExecutionManager, dataStore, taskExecutionManager)
	assert.NoError(t, err)

	inputsLM := &core.LiteralMap{
		Literals: map[string]*core.Literal{
			"input": {
				Value: &core.Literal_Scalar{
					Scalar: &core.Scalar{
						Value: &core.Scalar_Primitive{
							Primitive: &core.Primitive{
								Value: &core.Primitive_StringValue{
									StringValue: "hello",
								},
							},
						},
					},
				},
			},
		},
	}
	outputsLM := &core.LiteralMap{
		Literals: map[string]*core.Literal{
			"output": {
				Value: &core.Literal_Scalar{
					Scalar: &core.Scalar{
						Value: &core.Scalar_Primitive{
							Primitive: &core.Primitive{
								Value: &core.Primitive_StringValue{
									StringValue: "world",
								},
							},
						},
					},
				},
			},
		},
	}

	nodeExecutionManager.SetGetNodeExecutionDataFunc(
		func(ctx context.Context, request admin.NodeExecutionGetDataRequest) (*admin.NodeExecutionGetDataResponse, error) {
			return &admin.NodeExecutionGetDataResponse{
				FullInputs:  inputsLM,
				FullOutputs: outputsLM,
			}, nil
		},
	)
	taskExecutionManager.SetListTaskExecutionsCallback(func(ctx context.Context, request admin.TaskExecutionListRequest) (*admin.TaskExecutionList, error) {
		return &admin.TaskExecutionList{
			TaskExecutions: []*admin.TaskExecution{
				{
					Id: &core.TaskExecutionIdentifier{
						TaskId: &core.Identifier{
							ResourceType: core.ResourceType_TASK,
							Project:      "proj",
							Domain:       "dev",
							Name:         "task",
							Version:      "v1",
						},
						NodeExecutionId: &core.NodeExecutionIdentifier{
							NodeId: "n0",
							ExecutionId: &core.WorkflowExecutionIdentifier{
								Project: "proj",
								Domain:  "dev",
								Name:    "wfexecid",
							},
						},
						RetryAttempt: 5,
					},
				},
			},
		}, nil
	})
	taskExecutionManager.SetGetTaskExecutionDataCallback(func(ctx context.Context, request admin.TaskExecutionGetDataRequest) (*admin.TaskExecutionGetDataResponse, error) {
		return &admin.TaskExecutionGetDataResponse{
			FullInputs:  inputsLM,
			FullOutputs: outputsLM,
		}, nil
	})

	t.Run("get a working set of urls without retry attempt", func(t *testing.T) {
		res, err := s.GetData(context.Background(), &service.GetDataRequest{
			NebulaUrl: "nebula://v1/proj/dev/wfexecid/n0-d0/i",
		})
		assert.NoError(t, err)
		assert.True(t, proto.Equal(inputsLM, res.GetLiteralMap()))
		assert.Nil(t, res.GetPreSignedUrls())
	})

	t.Run("get a working set of urls with a retry attempt", func(t *testing.T) {
		res, err := s.GetData(context.Background(), &service.GetDataRequest{
			NebulaUrl: "nebula://v1/proj/dev/wfexecid/n0-d0/5/o",
		})
		assert.NoError(t, err)
		assert.True(t, proto.Equal(outputsLM, res.GetLiteralMap()))
		assert.Nil(t, res.GetPreSignedUrls())
	})

	t.Run("Bad URL", func(t *testing.T) {
		_, err = s.GetData(context.Background(), &service.GetDataRequest{
			NebulaUrl: "nebula://v3/blah/lorem/m0-fdj",
		})
		assert.Error(t, err)
	})

	t.Run("get individual literal without retry attempt", func(t *testing.T) {
		res, err := s.GetData(context.Background(), &service.GetDataRequest{
			NebulaUrl: "nebula://v1/proj/dev/wfexecid/n0-d0/i/input",
		})
		assert.NoError(t, err)
		assert.True(t, proto.Equal(inputsLM.GetLiterals()["input"], res.GetLiteral()))
		assert.Nil(t, res.GetPreSignedUrls())
	})

	t.Run("get individual literal with a retry attempt", func(t *testing.T) {
		res, err := s.GetData(context.Background(), &service.GetDataRequest{
			NebulaUrl: "nebula://v1/proj/dev/wfexecid/n0-d0/5/o/output",
		})
		assert.NoError(t, err)
		assert.True(t, proto.Equal(outputsLM.GetLiterals()["output"], res.GetLiteral()))
		assert.Nil(t, res.GetPreSignedUrls())
	})

	t.Run("error requesting missing name without retry attempt", func(t *testing.T) {
		_, err := s.GetData(context.Background(), &service.GetDataRequest{
			NebulaUrl: "nebula://v1/proj/dev/wfexecid/n0-d0/i/o5",
		})
		assert.Error(t, err)
	})

	t.Run("error requesting missing name with a retry attempt", func(t *testing.T) {
		_, err := s.GetData(context.Background(), &service.GetDataRequest{
			NebulaUrl: "nebula://v1/proj/dev/wfexecid/n0-d0/5/o/o1",
		})
		assert.Error(t, err)
	})
}

func TestService_Error(t *testing.T) {
	dataStore := commonMocks.GetMockStorageClient()
	nodeExecutionManager := &mocks.MockNodeExecutionManager{}
	taskExecutionManager := &mocks.MockTaskExecutionManager{}
	s, err := NewService(config.DataProxyConfig{}, nodeExecutionManager, dataStore, taskExecutionManager)
	assert.NoError(t, err)

	t.Run("get a working set of urls without retry attempt", func(t *testing.T) {
		taskExecutionManager.SetListTaskExecutionsCallback(func(ctx context.Context, request admin.TaskExecutionListRequest) (*admin.TaskExecutionList, error) {
			return nil, errors.NewNebulaAdminErrorf(1, "not found")
		})
		nodeExecID := core.NodeExecutionIdentifier{
			NodeId: "n0",
			ExecutionId: &core.WorkflowExecutionIdentifier{
				Project: "proj",
				Domain:  "dev",
				Name:    "wfexecid",
			},
		}
		_, err := s.GetTaskExecutionID(context.Background(), 0, nodeExecID)
		assert.Error(t, err, "failed to list")
	})

	t.Run("get a working set of urls without retry attempt", func(t *testing.T) {
		taskExecutionManager.SetListTaskExecutionsCallback(func(ctx context.Context, request admin.TaskExecutionListRequest) (*admin.TaskExecutionList, error) {
			return &admin.TaskExecutionList{
				TaskExecutions: nil,
				Token:          "",
			}, nil
		})
		nodeExecID := core.NodeExecutionIdentifier{
			NodeId: "n0",
			ExecutionId: &core.WorkflowExecutionIdentifier{
				Project: "proj",
				Domain:  "dev",
				Name:    "wfexecid",
			},
		}
		_, err := s.GetTaskExecutionID(context.Background(), 0, nodeExecID)
		assert.Error(t, err, "no task executions")
	})
}
