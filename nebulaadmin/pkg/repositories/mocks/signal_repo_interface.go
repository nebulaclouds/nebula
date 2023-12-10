// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	interfaces "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/interfaces"
	mock "github.com/stretchr/testify/mock"

	models "github.com/nebulaclouds/nebula/nebulaadmin/pkg/repositories/models"
)

// SignalRepoInterface is an autogenerated mock type for the SignalRepoInterface type
type SignalRepoInterface struct {
	mock.Mock
}

type SignalRepoInterface_Get struct {
	*mock.Call
}

func (_m SignalRepoInterface_Get) Return(_a0 models.Signal, _a1 error) *SignalRepoInterface_Get {
	return &SignalRepoInterface_Get{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *SignalRepoInterface) OnGet(ctx context.Context, input models.SignalKey) *SignalRepoInterface_Get {
	c_call := _m.On("Get", ctx, input)
	return &SignalRepoInterface_Get{Call: c_call}
}

func (_m *SignalRepoInterface) OnGetMatch(matchers ...interface{}) *SignalRepoInterface_Get {
	c_call := _m.On("Get", matchers...)
	return &SignalRepoInterface_Get{Call: c_call}
}

// Get provides a mock function with given fields: ctx, input
func (_m *SignalRepoInterface) Get(ctx context.Context, input models.SignalKey) (models.Signal, error) {
	ret := _m.Called(ctx, input)

	var r0 models.Signal
	if rf, ok := ret.Get(0).(func(context.Context, models.SignalKey) models.Signal); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(models.Signal)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.SignalKey) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type SignalRepoInterface_GetOrCreate struct {
	*mock.Call
}

func (_m SignalRepoInterface_GetOrCreate) Return(_a0 error) *SignalRepoInterface_GetOrCreate {
	return &SignalRepoInterface_GetOrCreate{Call: _m.Call.Return(_a0)}
}

func (_m *SignalRepoInterface) OnGetOrCreate(ctx context.Context, input *models.Signal) *SignalRepoInterface_GetOrCreate {
	c_call := _m.On("GetOrCreate", ctx, input)
	return &SignalRepoInterface_GetOrCreate{Call: c_call}
}

func (_m *SignalRepoInterface) OnGetOrCreateMatch(matchers ...interface{}) *SignalRepoInterface_GetOrCreate {
	c_call := _m.On("GetOrCreate", matchers...)
	return &SignalRepoInterface_GetOrCreate{Call: c_call}
}

// GetOrCreate provides a mock function with given fields: ctx, input
func (_m *SignalRepoInterface) GetOrCreate(ctx context.Context, input *models.Signal) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Signal) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type SignalRepoInterface_List struct {
	*mock.Call
}

func (_m SignalRepoInterface_List) Return(_a0 []models.Signal, _a1 error) *SignalRepoInterface_List {
	return &SignalRepoInterface_List{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *SignalRepoInterface) OnList(ctx context.Context, input interfaces.ListResourceInput) *SignalRepoInterface_List {
	c_call := _m.On("List", ctx, input)
	return &SignalRepoInterface_List{Call: c_call}
}

func (_m *SignalRepoInterface) OnListMatch(matchers ...interface{}) *SignalRepoInterface_List {
	c_call := _m.On("List", matchers...)
	return &SignalRepoInterface_List{Call: c_call}
}

// List provides a mock function with given fields: ctx, input
func (_m *SignalRepoInterface) List(ctx context.Context, input interfaces.ListResourceInput) ([]models.Signal, error) {
	ret := _m.Called(ctx, input)

	var r0 []models.Signal
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.ListResourceInput) []models.Signal); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Signal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interfaces.ListResourceInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type SignalRepoInterface_Update struct {
	*mock.Call
}

func (_m SignalRepoInterface_Update) Return(_a0 error) *SignalRepoInterface_Update {
	return &SignalRepoInterface_Update{Call: _m.Call.Return(_a0)}
}

func (_m *SignalRepoInterface) OnUpdate(ctx context.Context, input models.SignalKey, value []byte) *SignalRepoInterface_Update {
	c_call := _m.On("Update", ctx, input, value)
	return &SignalRepoInterface_Update{Call: c_call}
}

func (_m *SignalRepoInterface) OnUpdateMatch(matchers ...interface{}) *SignalRepoInterface_Update {
	c_call := _m.On("Update", matchers...)
	return &SignalRepoInterface_Update{Call: c_call}
}

// Update provides a mock function with given fields: ctx, input, value
func (_m *SignalRepoInterface) Update(ctx context.Context, input models.SignalKey, value []byte) error {
	ret := _m.Called(ctx, input, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.SignalKey, []byte) error); ok {
		r0 = rf(ctx, input, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
