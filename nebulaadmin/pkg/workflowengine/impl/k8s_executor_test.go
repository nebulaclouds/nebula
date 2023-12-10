package impl

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	k8_api_err "k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/executioncluster"
	execClusterIfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/executioncluster/interfaces"
	clusterMock "github.com/nebulaclouds/nebula/nebulaadmin/pkg/executioncluster/mocks"
	runtimeInterfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/interfaces"
	runtimeMocks "github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/mocks"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/workflowengine/interfaces"
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/workflowengine/mocks"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"
	"github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	"github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	nebulaclient "github.com/nebulaclouds/nebula/nebulapropeller/pkg/client/clientset/versioned"
	v1alpha12 "github.com/nebulaclouds/nebula/nebulapropeller/pkg/client/clientset/versioned/typed/nebulaworkflow/v1alpha1"
)

var fakeNebulaWF = FakeNebulaWorkflowV1alpha1{}

type createCallback func(*v1alpha1.NebulaWorkflow, v1.CreateOptions) (*v1alpha1.NebulaWorkflow, error)
type deleteCallback func(name string, options *v1.DeleteOptions) error
type FakeNebulaWorkflow struct {
	v1alpha12.NebulaWorkflowInterface
	createCallback createCallback
	deleteCallback deleteCallback
}

func (b *FakeNebulaWorkflow) Create(ctx context.Context, wf *v1alpha1.NebulaWorkflow, opts v1.CreateOptions) (*v1alpha1.NebulaWorkflow, error) {
	if b.createCallback != nil {
		return b.createCallback(wf, opts)
	}
	return nil, nil
}

func (b *FakeNebulaWorkflow) Delete(ctx context.Context, name string, options v1.DeleteOptions) error {
	if b.deleteCallback != nil {
		return b.deleteCallback(name, &options)
	}
	return nil
}

type nebulaWorkflowsCallback func(string) v1alpha12.NebulaWorkflowInterface

type FakeNebulaWorkflowV1alpha1 struct {
	v1alpha12.NebulaworkflowV1alpha1Interface
	nebulaWorkflowsCallback nebulaWorkflowsCallback
}

func (b *FakeNebulaWorkflowV1alpha1) NebulaWorkflows(namespace string) v1alpha12.NebulaWorkflowInterface {
	if b.nebulaWorkflowsCallback != nil {
		return b.nebulaWorkflowsCallback(namespace)
	}
	return &FakeNebulaWorkflow{}
}

type FakeK8NebulaClient struct {
	nebulaclient.Interface
	ID string
}

func (b *FakeK8NebulaClient) NebulaworkflowV1alpha1() v1alpha12.NebulaworkflowV1alpha1Interface {
	return &fakeNebulaWF
}

const namespace = "p-d"

const clusterID = "C1"

var execID = &core.WorkflowExecutionIdentifier{
	Project: "proj",
	Domain:  "domain",
	Name:    "name",
}

var nebulaWf = &v1alpha1.NebulaWorkflow{
	ExecutionID: v1alpha1.ExecutionID{
		WorkflowExecutionIdentifier: execID,
	},
}

var testInputs = &core.LiteralMap{
	Literals: map[string]*core.Literal{
		"foo": {
			Value: &core.Literal_Scalar{
				Scalar: &core.Scalar{
					Value: &core.Scalar_Primitive{
						Primitive: &core.Primitive{
							Value: &core.Primitive_Integer{
								Integer: 4,
							},
						},
					},
				},
			},
		},
	},
}

func getFakeExecutionCluster() execClusterIfaces.ClusterInterface {
	fakeCluster := clusterMock.MockCluster{}
	fakeCluster.SetGetTargetCallback(func(ctx context.Context, spec *executioncluster.ExecutionTargetSpec) (target *executioncluster.ExecutionTarget, e error) {
		return &executioncluster.ExecutionTarget{
			ID:          clusterID,
			NebulaClient: &FakeK8NebulaClient{},
		}, nil
	})
	return &fakeCluster
}

func TestGetID(t *testing.T) {
	executor := K8sWorkflowExecutor{}
	assert.Equal(t, defaultIdentifier, executor.ID())
}

func TestExecute(t *testing.T) {
	fakeNebulaWorkflow := FakeNebulaWorkflow{}
	fakeNebulaWorkflow.createCallback = func(nebulaWorkflow *v1alpha1.NebulaWorkflow, opts v1.CreateOptions) (*v1alpha1.NebulaWorkflow, error) {
		assert.Equal(t, nebulaWf, nebulaWorkflow)
		assert.Empty(t, opts)
		return nil, nil
	}
	fakeNebulaWF.nebulaWorkflowsCallback = func(ns string) v1alpha12.NebulaWorkflowInterface {
		assert.Equal(t, namespace, ns)
		return &fakeNebulaWorkflow
	}

	mockApplicationConfig := runtimeMocks.MockApplicationProvider{}
	mockApplicationConfig.SetTopLevelConfig(runtimeInterfaces.ApplicationConfig{
		UseOffloadedWorkflowClosure: false,
	})
	mockRuntime := runtimeMocks.NewMockConfigurationProvider(&mockApplicationConfig, nil, nil, nil, nil, nil)

	mockBuilder := mocks.NebulaWorkflowBuilder{}
	workflowClosure := core.CompiledWorkflowClosure{
		Primary: &core.CompiledWorkflow{
			Template: &core.WorkflowTemplate{
				Id: &core.Identifier{
					Project: "p",
					Domain:  "d",
					Name:    "n",
					Version: "version",
				},
			},
		},
	}
	mockBuilder.OnBuildMatch(mock.MatchedBy(func(wfClosure *core.CompiledWorkflowClosure) bool {
		return proto.Equal(wfClosure, &workflowClosure)
	}), mock.MatchedBy(func(inputs *core.LiteralMap) bool {
		return proto.Equal(inputs, testInputs)
	}), mock.MatchedBy(func(executionID *core.WorkflowExecutionIdentifier) bool {
		return proto.Equal(executionID, execID)
	}), namespace).Return(nebulaWf, nil)
	executor := K8sWorkflowExecutor{
		config:           mockRuntime,
		workflowBuilder:  &mockBuilder,
		executionCluster: getFakeExecutionCluster(),
	}

	resp, err := executor.Execute(context.TODO(), interfaces.ExecutionData{
		Namespace:               namespace,
		ExecutionID:             execID,
		ReferenceWorkflowName:   "ref_workflow_name",
		ReferenceLaunchPlanName: "ref_lp_name",
		WorkflowClosure:         &workflowClosure,
		ExecutionParameters: interfaces.ExecutionParameters{
			Inputs: testInputs,
			ExecutionConfig: &admin.WorkflowExecutionConfig{
				SecurityContext: &core.SecurityContext{
					RunAs: &core.Identity{
						IamRole:           testRoleSc,
						K8SServiceAccount: testK8sServiceAccountSc,
					},
				},
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, resp.Cluster, clusterID)
}

func TestExecute_AlreadyExists(t *testing.T) {
	fakeNebulaWorkflow := FakeNebulaWorkflow{}
	fakeNebulaWorkflow.createCallback = func(nebulaWorkflow *v1alpha1.NebulaWorkflow, opts v1.CreateOptions) (*v1alpha1.NebulaWorkflow, error) {
		return nil, k8_api_err.NewAlreadyExists(schema.GroupResource{}, "")
	}
	fakeNebulaWF.nebulaWorkflowsCallback = func(ns string) v1alpha12.NebulaWorkflowInterface {
		assert.Equal(t, namespace, ns)
		return &fakeNebulaWorkflow
	}

	mockApplicationConfig := runtimeMocks.MockApplicationProvider{}
	mockApplicationConfig.SetTopLevelConfig(runtimeInterfaces.ApplicationConfig{
		UseOffloadedWorkflowClosure: false,
	})
	mockRuntime := runtimeMocks.NewMockConfigurationProvider(&mockApplicationConfig, nil, nil, nil, nil, nil)

	mockBuilder := mocks.NebulaWorkflowBuilder{}
	mockBuilder.OnBuildMatch(mock.Anything, mock.Anything, mock.Anything, namespace).Return(nebulaWf, nil)
	executor := K8sWorkflowExecutor{
		config:           mockRuntime,
		workflowBuilder:  &mockBuilder,
		executionCluster: getFakeExecutionCluster(),
	}

	resp, err := executor.Execute(context.TODO(), interfaces.ExecutionData{
		Namespace:               namespace,
		ExecutionID:             execID,
		ReferenceWorkflowName:   "ref_workflow_name",
		ReferenceLaunchPlanName: "ref_lp_name",
		ExecutionParameters: interfaces.ExecutionParameters{
			ExecutionConfig: &admin.WorkflowExecutionConfig{
				SecurityContext: &core.SecurityContext{
					RunAs: &core.Identity{
						IamRole:           testRoleSc,
						K8SServiceAccount: testK8sServiceAccountSc,
					},
				},
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, resp.Cluster, clusterID)
}

func TestExecute_MiscError(t *testing.T) {
	fakeNebulaWorkflow := FakeNebulaWorkflow{}
	fakeNebulaWorkflow.createCallback = func(nebulaWorkflow *v1alpha1.NebulaWorkflow, opts v1.CreateOptions) (*v1alpha1.NebulaWorkflow, error) {
		return nil, errors.New("call failed")
	}
	fakeNebulaWF.nebulaWorkflowsCallback = func(ns string) v1alpha12.NebulaWorkflowInterface {
		assert.Equal(t, namespace, ns)
		return &fakeNebulaWorkflow
	}

	mockApplicationConfig := runtimeMocks.MockApplicationProvider{}
	mockApplicationConfig.SetTopLevelConfig(runtimeInterfaces.ApplicationConfig{
		UseOffloadedWorkflowClosure: false,
	})
	mockRuntime := runtimeMocks.NewMockConfigurationProvider(&mockApplicationConfig, nil, nil, nil, nil, nil)

	mockBuilder := mocks.NebulaWorkflowBuilder{}
	mockBuilder.OnBuildMatch(mock.Anything, mock.Anything, mock.Anything, namespace).Return(nebulaWf, nil)
	executor := K8sWorkflowExecutor{
		config:           mockRuntime,
		workflowBuilder:  &mockBuilder,
		executionCluster: getFakeExecutionCluster(),
	}

	_, err := executor.Execute(context.TODO(), interfaces.ExecutionData{
		Namespace:               namespace,
		ExecutionID:             execID,
		ReferenceWorkflowName:   "ref_workflow_name",
		ReferenceLaunchPlanName: "ref_lp_name",
		ExecutionParameters: interfaces.ExecutionParameters{
			ExecutionConfig: &admin.WorkflowExecutionConfig{
				SecurityContext: &core.SecurityContext{
					RunAs: &core.Identity{
						IamRole:           testRoleSc,
						K8SServiceAccount: testK8sServiceAccountSc,
					},
				},
			},
		},
	})
	assert.EqualError(t, err, "failed to create workflow in propeller call failed")
}

func TestAbort(t *testing.T) {
	fakeNebulaWorkflow := FakeNebulaWorkflow{}
	fakeNebulaWorkflow.deleteCallback = func(name string, options *v1.DeleteOptions) error {
		assert.Equal(t, execID.Name, name)
		assert.Equal(t, options.PropagationPolicy, &deletePropagationBackground)
		return nil
	}
	fakeNebulaWF.nebulaWorkflowsCallback = func(ns string) v1alpha12.NebulaWorkflowInterface {
		assert.Equal(t, namespace, ns)
		return &fakeNebulaWorkflow
	}
	executor := K8sWorkflowExecutor{
		executionCluster: getFakeExecutionCluster(),
	}
	err := executor.Abort(context.TODO(), interfaces.AbortData{
		Namespace:   namespace,
		ExecutionID: execID,
		Cluster:     clusterID,
	})
	assert.NoError(t, err)
}

func TestAbort_Notfound(t *testing.T) {
	fakeNebulaWorkflow := FakeNebulaWorkflow{}
	fakeNebulaWorkflow.deleteCallback = func(name string, options *v1.DeleteOptions) error {
		return k8_api_err.NewNotFound(schema.GroupResource{
			Group:    "foo",
			Resource: "bar",
		}, execID.Name)
	}
	fakeNebulaWF.nebulaWorkflowsCallback = func(ns string) v1alpha12.NebulaWorkflowInterface {
		assert.Equal(t, namespace, ns)
		return &fakeNebulaWorkflow
	}
	executor := K8sWorkflowExecutor{
		executionCluster: getFakeExecutionCluster(),
	}
	err := executor.Abort(context.TODO(), interfaces.AbortData{
		Namespace:   namespace,
		ExecutionID: execID,
		Cluster:     clusterID,
	})
	assert.NoError(t, err)
}

func TestAbort_MiscError(t *testing.T) {
	fakeNebulaWorkflow := FakeNebulaWorkflow{}
	fakeNebulaWorkflow.deleteCallback = func(name string, options *v1.DeleteOptions) error {
		return errors.New("call failed")
	}
	fakeNebulaWF.nebulaWorkflowsCallback = func(ns string) v1alpha12.NebulaWorkflowInterface {
		assert.Equal(t, namespace, ns)
		return &fakeNebulaWorkflow
	}
	executor := K8sWorkflowExecutor{
		executionCluster: getFakeExecutionCluster(),
	}
	err := executor.Abort(context.TODO(), interfaces.AbortData{
		Namespace:   namespace,
		ExecutionID: execID,
		Cluster:     clusterID,
	})
	assert.EqualError(t, err, "failed to terminate execution: project:\"proj\" domain:\"domain\" name:\"name\"  with err call failed")
}

func TestExecute_OffloadWorkflowClosure(t *testing.T) {
	offloadedNebulaWf := &v1alpha1.NebulaWorkflow{
		ExecutionID: v1alpha1.ExecutionID{
			WorkflowExecutionIdentifier: execID,
		},
		WorkflowSpec: &v1alpha1.WorkflowSpec{},
		Tasks:        make(map[v1alpha1.TaskID]*v1alpha1.TaskSpec),
		SubWorkflows: make(map[v1alpha1.WorkflowID]*v1alpha1.WorkflowSpec),
	}

	mockApplicationConfig := runtimeMocks.MockApplicationProvider{}
	mockApplicationConfig.SetTopLevelConfig(runtimeInterfaces.ApplicationConfig{
		UseOffloadedWorkflowClosure: true,
	})
	mockRuntime := runtimeMocks.NewMockConfigurationProvider(&mockApplicationConfig, nil, nil, nil, nil, nil)

	mockBuilder := mocks.NebulaWorkflowBuilder{}
	workflowClosure := core.CompiledWorkflowClosure{
		Primary: &core.CompiledWorkflow{
			Template: &core.WorkflowTemplate{
				Id: &core.Identifier{
					Project: "p",
					Domain:  "d",
					Name:    "n",
					Version: "version",
				},
			},
		},
		SubWorkflows: []*core.CompiledWorkflow{},
		Tasks:        []*core.CompiledTask{},
	}
	mockBuilder.OnBuildMatch(mock.MatchedBy(func(wfClosure *core.CompiledWorkflowClosure) bool {
		return proto.Equal(wfClosure, &workflowClosure)
	}), mock.MatchedBy(func(inputs *core.LiteralMap) bool {
		return proto.Equal(inputs, testInputs)
	}), mock.MatchedBy(func(executionID *core.WorkflowExecutionIdentifier) bool {
		return proto.Equal(executionID, execID)
	}), namespace).Return(offloadedNebulaWf, nil)
	executor := K8sWorkflowExecutor{
		config:           mockRuntime,
		workflowBuilder:  &mockBuilder,
		executionCluster: getFakeExecutionCluster(),
	}

	assert.NotNil(t, offloadedNebulaWf.WorkflowSpec)
	assert.NotNil(t, offloadedNebulaWf.Tasks)
	assert.NotNil(t, offloadedNebulaWf.SubWorkflows)

	resp, err := executor.Execute(context.TODO(), interfaces.ExecutionData{
		Namespace:               namespace,
		ExecutionID:             execID,
		ReferenceWorkflowName:   "ref_workflow_name",
		ReferenceLaunchPlanName: "ref_lp_name",
		WorkflowClosure:         &workflowClosure,
		ExecutionParameters: interfaces.ExecutionParameters{
			Inputs: testInputs,
			ExecutionConfig: &admin.WorkflowExecutionConfig{
				SecurityContext: &core.SecurityContext{
					RunAs: &core.Identity{
						IamRole:           testRoleSc,
						K8SServiceAccount: testK8sServiceAccountSc,
					},
				},
			},
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, resp.Cluster, clusterID)

	assert.Nil(t, offloadedNebulaWf.WorkflowSpec)
	assert.Nil(t, offloadedNebulaWf.Tasks)
	assert.Nil(t, offloadedNebulaWf.SubWorkflows)
}
