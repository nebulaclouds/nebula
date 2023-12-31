// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
)

// Workflow is an autogenerated mock type for the Workflow type
type Workflow struct {
	mock.Mock
}

type Workflow_HandleAbortedWorkflow struct {
	*mock.Call
}

func (_m Workflow_HandleAbortedWorkflow) Return(_a0 error) *Workflow_HandleAbortedWorkflow {
	return &Workflow_HandleAbortedWorkflow{Call: _m.Call.Return(_a0)}
}

func (_m *Workflow) OnHandleAbortedWorkflow(ctx context.Context, w *v1alpha1.NebulaWorkflow, maxRetries uint32) *Workflow_HandleAbortedWorkflow {
	c_call := _m.On("HandleAbortedWorkflow", ctx, w, maxRetries)
	return &Workflow_HandleAbortedWorkflow{Call: c_call}
}

func (_m *Workflow) OnHandleAbortedWorkflowMatch(matchers ...interface{}) *Workflow_HandleAbortedWorkflow {
	c_call := _m.On("HandleAbortedWorkflow", matchers...)
	return &Workflow_HandleAbortedWorkflow{Call: c_call}
}

// HandleAbortedWorkflow provides a mock function with given fields: ctx, w, maxRetries
func (_m *Workflow) HandleAbortedWorkflow(ctx context.Context, w *v1alpha1.NebulaWorkflow, maxRetries uint32) error {
	ret := _m.Called(ctx, w, maxRetries)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.NebulaWorkflow, uint32) error); ok {
		r0 = rf(ctx, w, maxRetries)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Workflow_HandleNebulaWorkflow struct {
	*mock.Call
}

func (_m Workflow_HandleNebulaWorkflow) Return(_a0 error) *Workflow_HandleNebulaWorkflow {
	return &Workflow_HandleNebulaWorkflow{Call: _m.Call.Return(_a0)}
}

func (_m *Workflow) OnHandleNebulaWorkflow(ctx context.Context, w *v1alpha1.NebulaWorkflow) *Workflow_HandleNebulaWorkflow {
	c_call := _m.On("HandleNebulaWorkflow", ctx, w)
	return &Workflow_HandleNebulaWorkflow{Call: c_call}
}

func (_m *Workflow) OnHandleNebulaWorkflowMatch(matchers ...interface{}) *Workflow_HandleNebulaWorkflow {
	c_call := _m.On("HandleNebulaWorkflow", matchers...)
	return &Workflow_HandleNebulaWorkflow{Call: c_call}
}

// HandleNebulaWorkflow provides a mock function with given fields: ctx, w
func (_m *Workflow) HandleNebulaWorkflow(ctx context.Context, w *v1alpha1.NebulaWorkflow) error {
	ret := _m.Called(ctx, w)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.NebulaWorkflow) error); ok {
		r0 = rf(ctx, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Workflow_Initialize struct {
	*mock.Call
}

func (_m Workflow_Initialize) Return(_a0 error) *Workflow_Initialize {
	return &Workflow_Initialize{Call: _m.Call.Return(_a0)}
}

func (_m *Workflow) OnInitialize(ctx context.Context) *Workflow_Initialize {
	c_call := _m.On("Initialize", ctx)
	return &Workflow_Initialize{Call: c_call}
}

func (_m *Workflow) OnInitializeMatch(matchers ...interface{}) *Workflow_Initialize {
	c_call := _m.On("Initialize", matchers...)
	return &Workflow_Initialize{Call: c_call}
}

// Initialize provides a mock function with given fields: ctx
func (_m *Workflow) Initialize(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
