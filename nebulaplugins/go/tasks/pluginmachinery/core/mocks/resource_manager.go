// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/nebulaclouds/nebula/nebulaplugins/go/tasks/pluginmachinery/core"
	mock "github.com/stretchr/testify/mock"
)

// ResourceManager is an autogenerated mock type for the ResourceManager type
type ResourceManager struct {
	mock.Mock
}

type ResourceManager_AllocateResource struct {
	*mock.Call
}

func (_m ResourceManager_AllocateResource) Return(_a0 core.AllocationStatus, _a1 error) *ResourceManager_AllocateResource {
	return &ResourceManager_AllocateResource{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ResourceManager) OnAllocateResource(ctx context.Context, namespace core.ResourceNamespace, allocationToken string, constraintsSpec core.ResourceConstraintsSpec) *ResourceManager_AllocateResource {
	c_call := _m.On("AllocateResource", ctx, namespace, allocationToken, constraintsSpec)
	return &ResourceManager_AllocateResource{Call: c_call}
}

func (_m *ResourceManager) OnAllocateResourceMatch(matchers ...interface{}) *ResourceManager_AllocateResource {
	c_call := _m.On("AllocateResource", matchers...)
	return &ResourceManager_AllocateResource{Call: c_call}
}

// AllocateResource provides a mock function with given fields: ctx, namespace, allocationToken, constraintsSpec
func (_m *ResourceManager) AllocateResource(ctx context.Context, namespace core.ResourceNamespace, allocationToken string, constraintsSpec core.ResourceConstraintsSpec) (core.AllocationStatus, error) {
	ret := _m.Called(ctx, namespace, allocationToken, constraintsSpec)

	var r0 core.AllocationStatus
	if rf, ok := ret.Get(0).(func(context.Context, core.ResourceNamespace, string, core.ResourceConstraintsSpec) core.AllocationStatus); ok {
		r0 = rf(ctx, namespace, allocationToken, constraintsSpec)
	} else {
		r0 = ret.Get(0).(core.AllocationStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, core.ResourceNamespace, string, core.ResourceConstraintsSpec) error); ok {
		r1 = rf(ctx, namespace, allocationToken, constraintsSpec)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ResourceManager_GetID struct {
	*mock.Call
}

func (_m ResourceManager_GetID) Return(_a0 string) *ResourceManager_GetID {
	return &ResourceManager_GetID{Call: _m.Call.Return(_a0)}
}

func (_m *ResourceManager) OnGetID() *ResourceManager_GetID {
	c_call := _m.On("GetID")
	return &ResourceManager_GetID{Call: c_call}
}

func (_m *ResourceManager) OnGetIDMatch(matchers ...interface{}) *ResourceManager_GetID {
	c_call := _m.On("GetID", matchers...)
	return &ResourceManager_GetID{Call: c_call}
}

// GetID provides a mock function with given fields:
func (_m *ResourceManager) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ResourceManager_ReleaseResource struct {
	*mock.Call
}

func (_m ResourceManager_ReleaseResource) Return(_a0 error) *ResourceManager_ReleaseResource {
	return &ResourceManager_ReleaseResource{Call: _m.Call.Return(_a0)}
}

func (_m *ResourceManager) OnReleaseResource(ctx context.Context, namespace core.ResourceNamespace, allocationToken string) *ResourceManager_ReleaseResource {
	c_call := _m.On("ReleaseResource", ctx, namespace, allocationToken)
	return &ResourceManager_ReleaseResource{Call: c_call}
}

func (_m *ResourceManager) OnReleaseResourceMatch(matchers ...interface{}) *ResourceManager_ReleaseResource {
	c_call := _m.On("ReleaseResource", matchers...)
	return &ResourceManager_ReleaseResource{Call: c_call}
}

// ReleaseResource provides a mock function with given fields: ctx, namespace, allocationToken
func (_m *ResourceManager) ReleaseResource(ctx context.Context, namespace core.ResourceNamespace, allocationToken string) error {
	ret := _m.Called(ctx, namespace, allocationToken)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.ResourceNamespace, string) error); ok {
		r0 = rf(ctx, namespace, allocationToken)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
