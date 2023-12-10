// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/nebulaclouds/nebula/datacatalog/pkg/repositories/models"
)

// PartitionRepo is an autogenerated mock type for the PartitionRepo type
type PartitionRepo struct {
	mock.Mock
}

type PartitionRepo_Create struct {
	*mock.Call
}

func (_m PartitionRepo_Create) Return(_a0 error) *PartitionRepo_Create {
	return &PartitionRepo_Create{Call: _m.Call.Return(_a0)}
}

func (_m *PartitionRepo) OnCreate(ctx context.Context, in models.Partition) *PartitionRepo_Create {
	c_call := _m.On("Create", ctx, in)
	return &PartitionRepo_Create{Call: c_call}
}

func (_m *PartitionRepo) OnCreateMatch(matchers ...interface{}) *PartitionRepo_Create {
	c_call := _m.On("Create", matchers...)
	return &PartitionRepo_Create{Call: c_call}
}

// Create provides a mock function with given fields: ctx, in
func (_m *PartitionRepo) Create(ctx context.Context, in models.Partition) error {
	ret := _m.Called(ctx, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Partition) error); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
