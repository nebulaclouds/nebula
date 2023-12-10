// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	service "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/service"
)

// DataProxyServiceClient is an autogenerated mock type for the DataProxyServiceClient type
type DataProxyServiceClient struct {
	mock.Mock
}

type DataProxyServiceClient_CreateDownloadLink struct {
	*mock.Call
}

func (_m DataProxyServiceClient_CreateDownloadLink) Return(_a0 *service.CreateDownloadLinkResponse, _a1 error) *DataProxyServiceClient_CreateDownloadLink {
	return &DataProxyServiceClient_CreateDownloadLink{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataProxyServiceClient) OnCreateDownloadLink(ctx context.Context, in *service.CreateDownloadLinkRequest, opts ...grpc.CallOption) *DataProxyServiceClient_CreateDownloadLink {
	c_call := _m.On("CreateDownloadLink", ctx, in, opts)
	return &DataProxyServiceClient_CreateDownloadLink{Call: c_call}
}

func (_m *DataProxyServiceClient) OnCreateDownloadLinkMatch(matchers ...interface{}) *DataProxyServiceClient_CreateDownloadLink {
	c_call := _m.On("CreateDownloadLink", matchers...)
	return &DataProxyServiceClient_CreateDownloadLink{Call: c_call}
}

// CreateDownloadLink provides a mock function with given fields: ctx, in, opts
func (_m *DataProxyServiceClient) CreateDownloadLink(ctx context.Context, in *service.CreateDownloadLinkRequest, opts ...grpc.CallOption) (*service.CreateDownloadLinkResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *service.CreateDownloadLinkResponse
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateDownloadLinkRequest, ...grpc.CallOption) *service.CreateDownloadLinkResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CreateDownloadLinkResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.CreateDownloadLinkRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataProxyServiceClient_CreateDownloadLocation struct {
	*mock.Call
}

func (_m DataProxyServiceClient_CreateDownloadLocation) Return(_a0 *service.CreateDownloadLocationResponse, _a1 error) *DataProxyServiceClient_CreateDownloadLocation {
	return &DataProxyServiceClient_CreateDownloadLocation{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataProxyServiceClient) OnCreateDownloadLocation(ctx context.Context, in *service.CreateDownloadLocationRequest, opts ...grpc.CallOption) *DataProxyServiceClient_CreateDownloadLocation {
	c_call := _m.On("CreateDownloadLocation", ctx, in, opts)
	return &DataProxyServiceClient_CreateDownloadLocation{Call: c_call}
}

func (_m *DataProxyServiceClient) OnCreateDownloadLocationMatch(matchers ...interface{}) *DataProxyServiceClient_CreateDownloadLocation {
	c_call := _m.On("CreateDownloadLocation", matchers...)
	return &DataProxyServiceClient_CreateDownloadLocation{Call: c_call}
}

// CreateDownloadLocation provides a mock function with given fields: ctx, in, opts
func (_m *DataProxyServiceClient) CreateDownloadLocation(ctx context.Context, in *service.CreateDownloadLocationRequest, opts ...grpc.CallOption) (*service.CreateDownloadLocationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *service.CreateDownloadLocationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateDownloadLocationRequest, ...grpc.CallOption) *service.CreateDownloadLocationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CreateDownloadLocationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.CreateDownloadLocationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataProxyServiceClient_CreateUploadLocation struct {
	*mock.Call
}

func (_m DataProxyServiceClient_CreateUploadLocation) Return(_a0 *service.CreateUploadLocationResponse, _a1 error) *DataProxyServiceClient_CreateUploadLocation {
	return &DataProxyServiceClient_CreateUploadLocation{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataProxyServiceClient) OnCreateUploadLocation(ctx context.Context, in *service.CreateUploadLocationRequest, opts ...grpc.CallOption) *DataProxyServiceClient_CreateUploadLocation {
	c_call := _m.On("CreateUploadLocation", ctx, in, opts)
	return &DataProxyServiceClient_CreateUploadLocation{Call: c_call}
}

func (_m *DataProxyServiceClient) OnCreateUploadLocationMatch(matchers ...interface{}) *DataProxyServiceClient_CreateUploadLocation {
	c_call := _m.On("CreateUploadLocation", matchers...)
	return &DataProxyServiceClient_CreateUploadLocation{Call: c_call}
}

// CreateUploadLocation provides a mock function with given fields: ctx, in, opts
func (_m *DataProxyServiceClient) CreateUploadLocation(ctx context.Context, in *service.CreateUploadLocationRequest, opts ...grpc.CallOption) (*service.CreateUploadLocationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *service.CreateUploadLocationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *service.CreateUploadLocationRequest, ...grpc.CallOption) *service.CreateUploadLocationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CreateUploadLocationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.CreateUploadLocationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataProxyServiceClient_GetData struct {
	*mock.Call
}

func (_m DataProxyServiceClient_GetData) Return(_a0 *service.GetDataResponse, _a1 error) *DataProxyServiceClient_GetData {
	return &DataProxyServiceClient_GetData{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataProxyServiceClient) OnGetData(ctx context.Context, in *service.GetDataRequest, opts ...grpc.CallOption) *DataProxyServiceClient_GetData {
	c_call := _m.On("GetData", ctx, in, opts)
	return &DataProxyServiceClient_GetData{Call: c_call}
}

func (_m *DataProxyServiceClient) OnGetDataMatch(matchers ...interface{}) *DataProxyServiceClient_GetData {
	c_call := _m.On("GetData", matchers...)
	return &DataProxyServiceClient_GetData{Call: c_call}
}

// GetData provides a mock function with given fields: ctx, in, opts
func (_m *DataProxyServiceClient) GetData(ctx context.Context, in *service.GetDataRequest, opts ...grpc.CallOption) (*service.GetDataResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *service.GetDataResponse
	if rf, ok := ret.Get(0).(func(context.Context, *service.GetDataRequest, ...grpc.CallOption) *service.GetDataResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.GetDataResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.GetDataRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
