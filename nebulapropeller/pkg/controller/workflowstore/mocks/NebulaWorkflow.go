// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	v1alpha1 "github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
	workflowstore "github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/workflowstore"
	mock "github.com/stretchr/testify/mock"
)

// NebulaWorkflow is an autogenerated mock type for the NebulaWorkflow type
type NebulaWorkflow struct {
	mock.Mock
}

type NebulaWorkflow_Get struct {
	*mock.Call
}

func (_m NebulaWorkflow_Get) Return(_a0 *v1alpha1.NebulaWorkflow, _a1 error) *NebulaWorkflow_Get {
	return &NebulaWorkflow_Get{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *NebulaWorkflow) OnGet(ctx context.Context, namespace string, name string) *NebulaWorkflow_Get {
	c_call := _m.On("Get", ctx, namespace, name)
	return &NebulaWorkflow_Get{Call: c_call}
}

func (_m *NebulaWorkflow) OnGetMatch(matchers ...interface{}) *NebulaWorkflow_Get {
	c_call := _m.On("Get", matchers...)
	return &NebulaWorkflow_Get{Call: c_call}
}

// Get provides a mock function with given fields: ctx, namespace, name
func (_m *NebulaWorkflow) Get(ctx context.Context, namespace string, name string) (*v1alpha1.NebulaWorkflow, error) {
	ret := _m.Called(ctx, namespace, name)

	var r0 *v1alpha1.NebulaWorkflow
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *v1alpha1.NebulaWorkflow); ok {
		r0 = rf(ctx, namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.NebulaWorkflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NebulaWorkflow_Update struct {
	*mock.Call
}

func (_m NebulaWorkflow_Update) Return(newWF *v1alpha1.NebulaWorkflow, err error) *NebulaWorkflow_Update {
	return &NebulaWorkflow_Update{Call: _m.Call.Return(newWF, err)}
}

func (_m *NebulaWorkflow) OnUpdate(ctx context.Context, workflow *v1alpha1.NebulaWorkflow, priorityClass workflowstore.PriorityClass) *NebulaWorkflow_Update {
	c_call := _m.On("Update", ctx, workflow, priorityClass)
	return &NebulaWorkflow_Update{Call: c_call}
}

func (_m *NebulaWorkflow) OnUpdateMatch(matchers ...interface{}) *NebulaWorkflow_Update {
	c_call := _m.On("Update", matchers...)
	return &NebulaWorkflow_Update{Call: c_call}
}

// Update provides a mock function with given fields: ctx, workflow, priorityClass
func (_m *NebulaWorkflow) Update(ctx context.Context, workflow *v1alpha1.NebulaWorkflow, priorityClass workflowstore.PriorityClass) (*v1alpha1.NebulaWorkflow, error) {
	ret := _m.Called(ctx, workflow, priorityClass)

	var r0 *v1alpha1.NebulaWorkflow
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.NebulaWorkflow, workflowstore.PriorityClass) *v1alpha1.NebulaWorkflow); ok {
		r0 = rf(ctx, workflow, priorityClass)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.NebulaWorkflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.NebulaWorkflow, workflowstore.PriorityClass) error); ok {
		r1 = rf(ctx, workflow, priorityClass)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NebulaWorkflow_UpdateStatus struct {
	*mock.Call
}

func (_m NebulaWorkflow_UpdateStatus) Return(newWF *v1alpha1.NebulaWorkflow, err error) *NebulaWorkflow_UpdateStatus {
	return &NebulaWorkflow_UpdateStatus{Call: _m.Call.Return(newWF, err)}
}

func (_m *NebulaWorkflow) OnUpdateStatus(ctx context.Context, workflow *v1alpha1.NebulaWorkflow, priorityClass workflowstore.PriorityClass) *NebulaWorkflow_UpdateStatus {
	c_call := _m.On("UpdateStatus", ctx, workflow, priorityClass)
	return &NebulaWorkflow_UpdateStatus{Call: c_call}
}

func (_m *NebulaWorkflow) OnUpdateStatusMatch(matchers ...interface{}) *NebulaWorkflow_UpdateStatus {
	c_call := _m.On("UpdateStatus", matchers...)
	return &NebulaWorkflow_UpdateStatus{Call: c_call}
}

// UpdateStatus provides a mock function with given fields: ctx, workflow, priorityClass
func (_m *NebulaWorkflow) UpdateStatus(ctx context.Context, workflow *v1alpha1.NebulaWorkflow, priorityClass workflowstore.PriorityClass) (*v1alpha1.NebulaWorkflow, error) {
	ret := _m.Called(ctx, workflow, priorityClass)

	var r0 *v1alpha1.NebulaWorkflow
	if rf, ok := ret.Get(0).(func(context.Context, *v1alpha1.NebulaWorkflow, workflowstore.PriorityClass) *v1alpha1.NebulaWorkflow); ok {
		r0 = rf(ctx, workflow, priorityClass)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.NebulaWorkflow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1alpha1.NebulaWorkflow, workflowstore.PriorityClass) error); ok {
		r1 = rf(ctx, workflow, priorityClass)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}