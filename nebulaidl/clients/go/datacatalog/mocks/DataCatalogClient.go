// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	datacatalog "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/datacatalog"
	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// DataCatalogClient is an autogenerated mock type for the DataCatalogClient type
type DataCatalogClient struct {
	mock.Mock
}

type DataCatalogClient_AddTag struct {
	*mock.Call
}

func (_m DataCatalogClient_AddTag) Return(_a0 *datacatalog.AddTagResponse, _a1 error) *DataCatalogClient_AddTag {
	return &DataCatalogClient_AddTag{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataCatalogClient) OnAddTag(ctx context.Context, in *datacatalog.AddTagRequest, opts ...grpc.CallOption) *DataCatalogClient_AddTag {
	c_call := _m.On("AddTag", ctx, in, opts)
	return &DataCatalogClient_AddTag{Call: c_call}
}

func (_m *DataCatalogClient) OnAddTagMatch(matchers ...interface{}) *DataCatalogClient_AddTag {
	c_call := _m.On("AddTag", matchers...)
	return &DataCatalogClient_AddTag{Call: c_call}
}

// AddTag provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) AddTag(ctx context.Context, in *datacatalog.AddTagRequest, opts ...grpc.CallOption) (*datacatalog.AddTagResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.AddTagResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.AddTagRequest, ...grpc.CallOption) *datacatalog.AddTagResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.AddTagResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.AddTagRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataCatalogClient_CreateArtifact struct {
	*mock.Call
}

func (_m DataCatalogClient_CreateArtifact) Return(_a0 *datacatalog.CreateArtifactResponse, _a1 error) *DataCatalogClient_CreateArtifact {
	return &DataCatalogClient_CreateArtifact{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataCatalogClient) OnCreateArtifact(ctx context.Context, in *datacatalog.CreateArtifactRequest, opts ...grpc.CallOption) *DataCatalogClient_CreateArtifact {
	c_call := _m.On("CreateArtifact", ctx, in, opts)
	return &DataCatalogClient_CreateArtifact{Call: c_call}
}

func (_m *DataCatalogClient) OnCreateArtifactMatch(matchers ...interface{}) *DataCatalogClient_CreateArtifact {
	c_call := _m.On("CreateArtifact", matchers...)
	return &DataCatalogClient_CreateArtifact{Call: c_call}
}

// CreateArtifact provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) CreateArtifact(ctx context.Context, in *datacatalog.CreateArtifactRequest, opts ...grpc.CallOption) (*datacatalog.CreateArtifactResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.CreateArtifactResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.CreateArtifactRequest, ...grpc.CallOption) *datacatalog.CreateArtifactResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.CreateArtifactResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.CreateArtifactRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataCatalogClient_CreateDataset struct {
	*mock.Call
}

func (_m DataCatalogClient_CreateDataset) Return(_a0 *datacatalog.CreateDatasetResponse, _a1 error) *DataCatalogClient_CreateDataset {
	return &DataCatalogClient_CreateDataset{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataCatalogClient) OnCreateDataset(ctx context.Context, in *datacatalog.CreateDatasetRequest, opts ...grpc.CallOption) *DataCatalogClient_CreateDataset {
	c_call := _m.On("CreateDataset", ctx, in, opts)
	return &DataCatalogClient_CreateDataset{Call: c_call}
}

func (_m *DataCatalogClient) OnCreateDatasetMatch(matchers ...interface{}) *DataCatalogClient_CreateDataset {
	c_call := _m.On("CreateDataset", matchers...)
	return &DataCatalogClient_CreateDataset{Call: c_call}
}

// CreateDataset provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) CreateDataset(ctx context.Context, in *datacatalog.CreateDatasetRequest, opts ...grpc.CallOption) (*datacatalog.CreateDatasetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.CreateDatasetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.CreateDatasetRequest, ...grpc.CallOption) *datacatalog.CreateDatasetResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.CreateDatasetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.CreateDatasetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataCatalogClient_GetArtifact struct {
	*mock.Call
}

func (_m DataCatalogClient_GetArtifact) Return(_a0 *datacatalog.GetArtifactResponse, _a1 error) *DataCatalogClient_GetArtifact {
	return &DataCatalogClient_GetArtifact{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataCatalogClient) OnGetArtifact(ctx context.Context, in *datacatalog.GetArtifactRequest, opts ...grpc.CallOption) *DataCatalogClient_GetArtifact {
	c_call := _m.On("GetArtifact", ctx, in, opts)
	return &DataCatalogClient_GetArtifact{Call: c_call}
}

func (_m *DataCatalogClient) OnGetArtifactMatch(matchers ...interface{}) *DataCatalogClient_GetArtifact {
	c_call := _m.On("GetArtifact", matchers...)
	return &DataCatalogClient_GetArtifact{Call: c_call}
}

// GetArtifact provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) GetArtifact(ctx context.Context, in *datacatalog.GetArtifactRequest, opts ...grpc.CallOption) (*datacatalog.GetArtifactResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.GetArtifactResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.GetArtifactRequest, ...grpc.CallOption) *datacatalog.GetArtifactResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.GetArtifactResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.GetArtifactRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataCatalogClient_GetDataset struct {
	*mock.Call
}

func (_m DataCatalogClient_GetDataset) Return(_a0 *datacatalog.GetDatasetResponse, _a1 error) *DataCatalogClient_GetDataset {
	return &DataCatalogClient_GetDataset{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataCatalogClient) OnGetDataset(ctx context.Context, in *datacatalog.GetDatasetRequest, opts ...grpc.CallOption) *DataCatalogClient_GetDataset {
	c_call := _m.On("GetDataset", ctx, in, opts)
	return &DataCatalogClient_GetDataset{Call: c_call}
}

func (_m *DataCatalogClient) OnGetDatasetMatch(matchers ...interface{}) *DataCatalogClient_GetDataset {
	c_call := _m.On("GetDataset", matchers...)
	return &DataCatalogClient_GetDataset{Call: c_call}
}

// GetDataset provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) GetDataset(ctx context.Context, in *datacatalog.GetDatasetRequest, opts ...grpc.CallOption) (*datacatalog.GetDatasetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.GetDatasetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.GetDatasetRequest, ...grpc.CallOption) *datacatalog.GetDatasetResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.GetDatasetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.GetDatasetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataCatalogClient_GetOrExtendReservation struct {
	*mock.Call
}

func (_m DataCatalogClient_GetOrExtendReservation) Return(_a0 *datacatalog.GetOrExtendReservationResponse, _a1 error) *DataCatalogClient_GetOrExtendReservation {
	return &DataCatalogClient_GetOrExtendReservation{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataCatalogClient) OnGetOrExtendReservation(ctx context.Context, in *datacatalog.GetOrExtendReservationRequest, opts ...grpc.CallOption) *DataCatalogClient_GetOrExtendReservation {
	c_call := _m.On("GetOrExtendReservation", ctx, in, opts)
	return &DataCatalogClient_GetOrExtendReservation{Call: c_call}
}

func (_m *DataCatalogClient) OnGetOrExtendReservationMatch(matchers ...interface{}) *DataCatalogClient_GetOrExtendReservation {
	c_call := _m.On("GetOrExtendReservation", matchers...)
	return &DataCatalogClient_GetOrExtendReservation{Call: c_call}
}

// GetOrExtendReservation provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) GetOrExtendReservation(ctx context.Context, in *datacatalog.GetOrExtendReservationRequest, opts ...grpc.CallOption) (*datacatalog.GetOrExtendReservationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.GetOrExtendReservationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.GetOrExtendReservationRequest, ...grpc.CallOption) *datacatalog.GetOrExtendReservationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.GetOrExtendReservationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.GetOrExtendReservationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataCatalogClient_ListArtifacts struct {
	*mock.Call
}

func (_m DataCatalogClient_ListArtifacts) Return(_a0 *datacatalog.ListArtifactsResponse, _a1 error) *DataCatalogClient_ListArtifacts {
	return &DataCatalogClient_ListArtifacts{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataCatalogClient) OnListArtifacts(ctx context.Context, in *datacatalog.ListArtifactsRequest, opts ...grpc.CallOption) *DataCatalogClient_ListArtifacts {
	c_call := _m.On("ListArtifacts", ctx, in, opts)
	return &DataCatalogClient_ListArtifacts{Call: c_call}
}

func (_m *DataCatalogClient) OnListArtifactsMatch(matchers ...interface{}) *DataCatalogClient_ListArtifacts {
	c_call := _m.On("ListArtifacts", matchers...)
	return &DataCatalogClient_ListArtifacts{Call: c_call}
}

// ListArtifacts provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) ListArtifacts(ctx context.Context, in *datacatalog.ListArtifactsRequest, opts ...grpc.CallOption) (*datacatalog.ListArtifactsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.ListArtifactsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.ListArtifactsRequest, ...grpc.CallOption) *datacatalog.ListArtifactsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.ListArtifactsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.ListArtifactsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataCatalogClient_ListDatasets struct {
	*mock.Call
}

func (_m DataCatalogClient_ListDatasets) Return(_a0 *datacatalog.ListDatasetsResponse, _a1 error) *DataCatalogClient_ListDatasets {
	return &DataCatalogClient_ListDatasets{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataCatalogClient) OnListDatasets(ctx context.Context, in *datacatalog.ListDatasetsRequest, opts ...grpc.CallOption) *DataCatalogClient_ListDatasets {
	c_call := _m.On("ListDatasets", ctx, in, opts)
	return &DataCatalogClient_ListDatasets{Call: c_call}
}

func (_m *DataCatalogClient) OnListDatasetsMatch(matchers ...interface{}) *DataCatalogClient_ListDatasets {
	c_call := _m.On("ListDatasets", matchers...)
	return &DataCatalogClient_ListDatasets{Call: c_call}
}

// ListDatasets provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) ListDatasets(ctx context.Context, in *datacatalog.ListDatasetsRequest, opts ...grpc.CallOption) (*datacatalog.ListDatasetsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.ListDatasetsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.ListDatasetsRequest, ...grpc.CallOption) *datacatalog.ListDatasetsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.ListDatasetsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.ListDatasetsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataCatalogClient_ReleaseReservation struct {
	*mock.Call
}

func (_m DataCatalogClient_ReleaseReservation) Return(_a0 *datacatalog.ReleaseReservationResponse, _a1 error) *DataCatalogClient_ReleaseReservation {
	return &DataCatalogClient_ReleaseReservation{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataCatalogClient) OnReleaseReservation(ctx context.Context, in *datacatalog.ReleaseReservationRequest, opts ...grpc.CallOption) *DataCatalogClient_ReleaseReservation {
	c_call := _m.On("ReleaseReservation", ctx, in, opts)
	return &DataCatalogClient_ReleaseReservation{Call: c_call}
}

func (_m *DataCatalogClient) OnReleaseReservationMatch(matchers ...interface{}) *DataCatalogClient_ReleaseReservation {
	c_call := _m.On("ReleaseReservation", matchers...)
	return &DataCatalogClient_ReleaseReservation{Call: c_call}
}

// ReleaseReservation provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) ReleaseReservation(ctx context.Context, in *datacatalog.ReleaseReservationRequest, opts ...grpc.CallOption) (*datacatalog.ReleaseReservationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.ReleaseReservationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.ReleaseReservationRequest, ...grpc.CallOption) *datacatalog.ReleaseReservationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.ReleaseReservationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.ReleaseReservationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DataCatalogClient_UpdateArtifact struct {
	*mock.Call
}

func (_m DataCatalogClient_UpdateArtifact) Return(_a0 *datacatalog.UpdateArtifactResponse, _a1 error) *DataCatalogClient_UpdateArtifact {
	return &DataCatalogClient_UpdateArtifact{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *DataCatalogClient) OnUpdateArtifact(ctx context.Context, in *datacatalog.UpdateArtifactRequest, opts ...grpc.CallOption) *DataCatalogClient_UpdateArtifact {
	c_call := _m.On("UpdateArtifact", ctx, in, opts)
	return &DataCatalogClient_UpdateArtifact{Call: c_call}
}

func (_m *DataCatalogClient) OnUpdateArtifactMatch(matchers ...interface{}) *DataCatalogClient_UpdateArtifact {
	c_call := _m.On("UpdateArtifact", matchers...)
	return &DataCatalogClient_UpdateArtifact{Call: c_call}
}

// UpdateArtifact provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) UpdateArtifact(ctx context.Context, in *datacatalog.UpdateArtifactRequest, opts ...grpc.CallOption) (*datacatalog.UpdateArtifactResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.UpdateArtifactResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.UpdateArtifactRequest, ...grpc.CallOption) *datacatalog.UpdateArtifactResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.UpdateArtifactResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.UpdateArtifactRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
