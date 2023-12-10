package util

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"

	nebulaAdminErrors "github.com/nebulaclouds/nebula/nebulaadmin/pkg/errors"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/manager/impl/testutils"
	managerMocks "github.com/nebulaclouds/nebula/nebulaadmin/pkg/manager/mocks"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/interfaces"
	repositoryMocks "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/mocks"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
	runtimeMocks "github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/mocks"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
)

func TestGenerateNodeNameFromTask(t *testing.T) {
	assert.EqualValues(t, "foo-12345", generateNodeNameFromTask("foo@`!=+- 1,2$3(4)5"))
	assert.EqualValues(t, "nametoadefinedtasksomewhereincodeveryverynestedcrazy",
		generateNodeNameFromTask("app.long.path.name.to.a.defined.task.somewhere.in.code.very.very.nested.crazy"))
}

func TestGenerateWorkflowNameFromTask(t *testing.T) {
	assert.EqualValues(t, ".nebulagen.SingleTask", generateWorkflowNameFromTask("SingleTask"))
}

func TestGenerateBindings(t *testing.T) {
	nodeID := "nodeID"
	outputs := core.VariableMap{
		Variables: map[string]*core.Variable{
			"output1": {},
			"output2": {},
		},
	}
	generatedBindings := generateBindings(outputs, nodeID)
	assert.Contains(t, generatedBindings,
		&core.Binding{
			Var: "output1",
			Binding: &core.BindingData{
				Value: &core.BindingData_Promise{
					Promise: &core.OutputReference{
						NodeId: nodeID,
						Var:    "output1",
					},
				},
			},
		},
	)
	assert.Contains(t, generatedBindings,
		&core.Binding{
			Var: "output2",
			Binding: &core.BindingData{
				Value: &core.BindingData_Promise{
					Promise: &core.OutputReference{
						NodeId: nodeID,
						Var:    "output2",
					},
				},
			},
		})
}

func TestCreateOrGetWorkflowModel(t *testing.T) {
	repository := repositoryMocks.NewMockRepository()
	var getCalledCount = 0
	var newlyCreatedWorkflow models.Workflow
	workflowcreateFunc := func(input models.Workflow, descriptionEntity *models.DescriptionEntity) error {
		newlyCreatedWorkflow = input
		return nil
	}
	repository.WorkflowRepo().(*repositoryMocks.MockWorkflowRepo).SetCreateCallback(workflowcreateFunc)

	workflowGetFunc := func(input interfaces.Identifier) (models.Workflow, error) {
		if getCalledCount == 0 {
			getCalledCount++
			return models.Workflow{}, nebulaAdminErrors.NewNebulaAdminErrorf(codes.NotFound, "not found")
		}
		getCalledCount++
		return newlyCreatedWorkflow, nil
	}
	repository.WorkflowRepo().(*repositoryMocks.MockWorkflowRepo).SetGetCallback(workflowGetFunc)

	mockNamedEntityManager := managerMocks.NamedEntityManager{}
	mockNamedEntityManager.UpdateNamedEntityFunc = func(ctx context.Context, request admin.NamedEntityUpdateRequest) (*admin.NamedEntityUpdateResponse, error) {
		assert.Equal(t, request.ResourceType, core.ResourceType_WORKFLOW)
		assert.True(t, proto.Equal(request.Id, &admin.NamedEntityIdentifier{
			Project: "nebulakit",
			Domain:  "production",
			Name:    ".nebulagen.app.workflows.MyWorkflow.my_task",
		}), fmt.Sprintf("%+v", request.Id))
		assert.True(t, proto.Equal(request.Metadata, &admin.NamedEntityMetadata{
			State: admin.NamedEntityState_SYSTEM_GENERATED,
		}))
		return &admin.NamedEntityUpdateResponse{}, nil
	}

	mockWorkflowManager := managerMocks.MockWorkflowManager{}
	mockWorkflowManager.SetCreateCallback(func(ctx context.Context, request admin.WorkflowCreateRequest) (*admin.WorkflowCreateResponse, error) {
		assert.True(t, proto.Equal(request.Id, &core.Identifier{
			ResourceType: core.ResourceType_WORKFLOW,
			Project:      "nebulakit",
			Domain:       "production",
			Name:         ".nebulagen.app.workflows.MyWorkflow.my_task",
			Version:      "12345",
		}), fmt.Sprintf("%+v", request.Id))
		return &admin.WorkflowCreateResponse{}, nil
	})
	taskIdentifier := &core.Identifier{
		ResourceType: core.ResourceType_TASK,
		Project:      "nebulakit",
		Domain:       "production",
		Name:         "app.workflows.MyWorkflow.my_task",
		Version:      "12345",
	}
	var task = &admin.Task{
		Id: taskIdentifier,
		Closure: &admin.TaskClosure{
			CompiledTask: &core.CompiledTask{
				Template: &core.TaskTemplate{
					Id: taskIdentifier,
					Interface: &core.TypedInterface{
						Inputs: &core.VariableMap{
							Variables: map[string]*core.Variable{
								"an_int": {
									Type: &core.LiteralType{
										Type: &core.LiteralType_Simple{
											Simple: core.SimpleType_INTEGER,
										},
									},
								},
							},
						},
						Outputs: &core.VariableMap{
							Variables: map[string]*core.Variable{
								"an_output": {
									Type: &core.LiteralType{
										Type: &core.LiteralType_Simple{
											Simple: core.SimpleType_DURATION,
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
	workflowModel, err := CreateOrGetWorkflowModel(context.Background(), admin.ExecutionCreateRequest{
		Project: "nebulakit",
		Domain:  "production",
		Name:    "SingleTaskExecution",
		Spec: &admin.ExecutionSpec{
			LaunchPlan: taskIdentifier,
		},
	}, repository, &mockWorkflowManager, &mockNamedEntityManager, taskIdentifier, task)
	assert.NoError(t, err)
	assert.EqualValues(t, workflowModel, &newlyCreatedWorkflow)
}

func TestCreateOrGetLaunchPlan(t *testing.T) {
	repository := repositoryMocks.NewMockRepository()
	var getCalledCount = 0
	var newlyCreatedLP models.LaunchPlan
	launchPlanCreateFunc := func(input models.LaunchPlan) error {
		newlyCreatedLP = input
		return nil
	}
	repository.LaunchPlanRepo().(*repositoryMocks.MockLaunchPlanRepo).SetCreateCallback(launchPlanCreateFunc)

	launchPlanGetFunc := func(input interfaces.Identifier) (models.LaunchPlan, error) {
		if getCalledCount == 0 {
			getCalledCount++
			return models.LaunchPlan{}, nebulaAdminErrors.NewNebulaAdminErrorf(codes.NotFound, "not found")
		}
		getCalledCount++
		return newlyCreatedLP, nil
	}
	repository.LaunchPlanRepo().(*repositoryMocks.MockLaunchPlanRepo).SetGetCallback(launchPlanGetFunc)

	workflowInterface := &core.TypedInterface{
		Inputs: &core.VariableMap{
			Variables: map[string]*core.Variable{
				"an_int": {
					Type: &core.LiteralType{
						Type: &core.LiteralType_Simple{
							Simple: core.SimpleType_INTEGER,
						},
					},
				},
			},
		},
		Outputs: &core.VariableMap{
			Variables: map[string]*core.Variable{
				"an_output": {
					Type: &core.LiteralType{
						Type: &core.LiteralType_Simple{
							Simple: core.SimpleType_DURATION,
						},
					},
				},
			},
		},
	}
	workflowID := uint(12)
	taskIdentifier := &core.Identifier{
		ResourceType: core.ResourceType_TASK,
		Project:      "nebulakit",
		Domain:       "production",
		Name:         "app.workflows.MyWorkflow.my_task",
		Version:      "12345",
	}
	config := runtimeMocks.NewMockConfigurationProvider(
		testutils.GetApplicationConfigWithDefaultDomains(), nil, nil, nil, nil, nil)
	spec := admin.ExecutionSpec{
		LaunchPlan: taskIdentifier,
		AuthRole: &admin.AuthRole{
			AssumableIamRole: "assumable_role",
		},
	}
	launchPlan, err := CreateOrGetLaunchPlan(
		context.Background(), repository, config, taskIdentifier, workflowInterface, workflowID, &spec)
	assert.NoError(t, err)
	assert.True(t, proto.Equal(&core.Identifier{
		ResourceType: core.ResourceType_LAUNCH_PLAN,
		Project:      "nebulakit",
		Domain:       "production",
		Name:         ".nebulagen.app.workflows.MyWorkflow.my_task",
		Version:      "12345",
	}, launchPlan.Id))
	assert.True(t, proto.Equal(launchPlan.Closure.ExpectedOutputs, workflowInterface.Outputs))
	assert.True(t, proto.Equal(launchPlan.Spec.AuthRole, spec.AuthRole))
}
