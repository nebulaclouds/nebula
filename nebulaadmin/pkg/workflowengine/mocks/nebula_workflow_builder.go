// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"

	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
)

// NebulaWorkflowBuilder is an autogenerated mock type for the NebulaWorkflowBuilder type
type NebulaWorkflowBuilder struct {
	mock.Mock
}

type NebulaWorkflowBuilder_Build struct {
	*mock.Call
}

func (_m NebulaWorkflowBuilder_Build) Return(_a0 *v1alpha1.NebulaWorkflow, _a1 error) *NebulaWorkflowBuilder_Build {
	return &NebulaWorkflowBuilder_Build{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *NebulaWorkflowBuilder) OnBuild(wfClosure *core.CompiledWorkflowClosure, inputs *core.LiteralMap, executionID *core.WorkflowExecutionIdentifier, namespace string) *NebulaWorkflowBuilder_Build {
	c_call := _m.On("Build", wfClosure, inputs, executionID, namespace)
	return &NebulaWorkflowBuilder_Build{Call: c_call}
}

func (_m *NebulaWorkflowBuilder) OnBuildMatch(matchers ...interface{}) *NebulaWorkflowBuilder_Build {
	c_call := _m.On("Build", matchers...)
	return &NebulaWorkflowBuilder_Build{Call: c_call}
}

// Build provides a mock function with given fields: wfClosure, inputs, executionID, namespace
func (_m *NebulaWorkflowBuilder) Build(wfClosure *core.CompiledWorkflowClosure, inputs *core.LiteralMap, executionID *core.WorkflowExecutionIdentifier, namespace string) (*v1alpha1.NebulaWorkflow, error) {
	ret := _m.Called(wfClosure, inputs, executionID, namespace)

	var r0 *v1alpha1.NebulaWorkflow
	if rf, ok := ret.Get(0).(func(*core.CompiledWorkflowClosure, *core.LiteralMap, *core.WorkflowExecutionIdentifier, string) *v1alpha1.NebulaWorkflow); ok {
		r0 = rf(wfClosure, inputs, executionID, namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.NebulaWorkflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*core.CompiledWorkflowClosure, *core.LiteralMap, *core.WorkflowExecutionIdentifier, string) error); ok {
		r1 = rf(wfClosure, inputs, executionID, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
