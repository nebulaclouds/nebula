// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	storage "github.com/nebulaclouds/nebula/nebulastdlib/storage"
)

// TaskTemplatePath is an autogenerated mock type for the TaskTemplatePath type
type TaskTemplatePath struct {
	mock.Mock
}

type TaskTemplatePath_Path struct {
	*mock.Call
}

func (_m TaskTemplatePath_Path) Return(_a0 storage.DataReference, _a1 error) *TaskTemplatePath_Path {
	return &TaskTemplatePath_Path{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *TaskTemplatePath) OnPath(ctx context.Context) *TaskTemplatePath_Path {
	c_call := _m.On("Path", ctx)
	return &TaskTemplatePath_Path{Call: c_call}
}

func (_m *TaskTemplatePath) OnPathMatch(matchers ...interface{}) *TaskTemplatePath_Path {
	c_call := _m.On("Path", matchers...)
	return &TaskTemplatePath_Path{Call: c_call}
}

// Path provides a mock function with given fields: ctx
func (_m *TaskTemplatePath) Path(ctx context.Context) (storage.DataReference, error) {
	ret := _m.Called(ctx)

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func(context.Context) storage.DataReference); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
