// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	protoiface "google.golang.org/protobuf/runtime/protoiface"

	storage "github.com/nebulaclouds/nebula/nebulastdlib/storage"
)

// ComposedProtobufStore is an autogenerated mock type for the ComposedProtobufStore type
type ComposedProtobufStore struct {
	mock.Mock
}

type ComposedProtobufStore_CopyRaw struct {
	*mock.Call
}

func (_m ComposedProtobufStore_CopyRaw) Return(_a0 error) *ComposedProtobufStore_CopyRaw {
	return &ComposedProtobufStore_CopyRaw{Call: _m.Call.Return(_a0)}
}

func (_m *ComposedProtobufStore) OnCopyRaw(ctx context.Context, source storage.DataReference, destination storage.DataReference, opts storage.Options) *ComposedProtobufStore_CopyRaw {
	c := _m.On("CopyRaw", ctx, source, destination, opts)
	return &ComposedProtobufStore_CopyRaw{Call: c}
}

func (_m *ComposedProtobufStore) OnCopyRawMatch(matchers ...interface{}) *ComposedProtobufStore_CopyRaw {
	c := _m.On("CopyRaw", matchers...)
	return &ComposedProtobufStore_CopyRaw{Call: c}
}

// CopyRaw provides a mock function with given fields: ctx, source, destination, opts
func (_m *ComposedProtobufStore) CopyRaw(ctx context.Context, source storage.DataReference, destination storage.DataReference, opts storage.Options) error {
	ret := _m.Called(ctx, source, destination, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, storage.DataReference, storage.Options) error); ok {
		r0 = rf(ctx, source, destination, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type ComposedProtobufStore_CreateSignedURL struct {
	*mock.Call
}

func (_m ComposedProtobufStore_CreateSignedURL) Return(_a0 storage.SignedURLResponse, _a1 error) *ComposedProtobufStore_CreateSignedURL {
	return &ComposedProtobufStore_CreateSignedURL{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ComposedProtobufStore) OnCreateSignedURL(ctx context.Context, reference storage.DataReference, properties storage.SignedURLProperties) *ComposedProtobufStore_CreateSignedURL {
	c := _m.On("CreateSignedURL", ctx, reference, properties)
	return &ComposedProtobufStore_CreateSignedURL{Call: c}
}

func (_m *ComposedProtobufStore) OnCreateSignedURLMatch(matchers ...interface{}) *ComposedProtobufStore_CreateSignedURL {
	c := _m.On("CreateSignedURL", matchers...)
	return &ComposedProtobufStore_CreateSignedURL{Call: c}
}

// CreateSignedURL provides a mock function with given fields: ctx, reference, properties
func (_m *ComposedProtobufStore) CreateSignedURL(ctx context.Context, reference storage.DataReference, properties storage.SignedURLProperties) (storage.SignedURLResponse, error) {
	ret := _m.Called(ctx, reference, properties)

	var r0 storage.SignedURLResponse
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, storage.SignedURLProperties) storage.SignedURLResponse); ok {
		r0 = rf(ctx, reference, properties)
	} else {
		r0 = ret.Get(0).(storage.SignedURLResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, storage.DataReference, storage.SignedURLProperties) error); ok {
		r1 = rf(ctx, reference, properties)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ComposedProtobufStore_Delete struct {
	*mock.Call
}

func (_m ComposedProtobufStore_Delete) Return(_a0 error) *ComposedProtobufStore_Delete {
	return &ComposedProtobufStore_Delete{Call: _m.Call.Return(_a0)}
}

func (_m *ComposedProtobufStore) OnDelete(ctx context.Context, reference storage.DataReference) *ComposedProtobufStore_Delete {
	c := _m.On("Delete", ctx, reference)
	return &ComposedProtobufStore_Delete{Call: c}
}

func (_m *ComposedProtobufStore) OnDeleteMatch(matchers ...interface{}) *ComposedProtobufStore_Delete {
	c := _m.On("Delete", matchers...)
	return &ComposedProtobufStore_Delete{Call: c}
}

// Delete provides a mock function with given fields: ctx, reference
func (_m *ComposedProtobufStore) Delete(ctx context.Context, reference storage.DataReference) error {
	ret := _m.Called(ctx, reference)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference) error); ok {
		r0 = rf(ctx, reference)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type ComposedProtobufStore_GetBaseContainerFQN struct {
	*mock.Call
}

func (_m ComposedProtobufStore_GetBaseContainerFQN) Return(_a0 storage.DataReference) *ComposedProtobufStore_GetBaseContainerFQN {
	return &ComposedProtobufStore_GetBaseContainerFQN{Call: _m.Call.Return(_a0)}
}

func (_m *ComposedProtobufStore) OnGetBaseContainerFQN(ctx context.Context) *ComposedProtobufStore_GetBaseContainerFQN {
	c := _m.On("GetBaseContainerFQN", ctx)
	return &ComposedProtobufStore_GetBaseContainerFQN{Call: c}
}

func (_m *ComposedProtobufStore) OnGetBaseContainerFQNMatch(matchers ...interface{}) *ComposedProtobufStore_GetBaseContainerFQN {
	c := _m.On("GetBaseContainerFQN", matchers...)
	return &ComposedProtobufStore_GetBaseContainerFQN{Call: c}
}

// GetBaseContainerFQN provides a mock function with given fields: ctx
func (_m *ComposedProtobufStore) GetBaseContainerFQN(ctx context.Context) storage.DataReference {
	ret := _m.Called(ctx)

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func(context.Context) storage.DataReference); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

type ComposedProtobufStore_Head struct {
	*mock.Call
}

func (_m ComposedProtobufStore_Head) Return(_a0 storage.Metadata, _a1 error) *ComposedProtobufStore_Head {
	return &ComposedProtobufStore_Head{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ComposedProtobufStore) OnHead(ctx context.Context, reference storage.DataReference) *ComposedProtobufStore_Head {
	c := _m.On("Head", ctx, reference)
	return &ComposedProtobufStore_Head{Call: c}
}

func (_m *ComposedProtobufStore) OnHeadMatch(matchers ...interface{}) *ComposedProtobufStore_Head {
	c := _m.On("Head", matchers...)
	return &ComposedProtobufStore_Head{Call: c}
}

// Head provides a mock function with given fields: ctx, reference
func (_m *ComposedProtobufStore) Head(ctx context.Context, reference storage.DataReference) (storage.Metadata, error) {
	ret := _m.Called(ctx, reference)

	var r0 storage.Metadata
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference) storage.Metadata); ok {
		r0 = rf(ctx, reference)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.Metadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, storage.DataReference) error); ok {
		r1 = rf(ctx, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ComposedProtobufStore_ReadProtobuf struct {
	*mock.Call
}

func (_m ComposedProtobufStore_ReadProtobuf) Return(_a0 error) *ComposedProtobufStore_ReadProtobuf {
	return &ComposedProtobufStore_ReadProtobuf{Call: _m.Call.Return(_a0)}
}

func (_m *ComposedProtobufStore) OnReadProtobuf(ctx context.Context, reference storage.DataReference, msg protoiface.MessageV1) *ComposedProtobufStore_ReadProtobuf {
	c := _m.On("ReadProtobuf", ctx, reference, msg)
	return &ComposedProtobufStore_ReadProtobuf{Call: c}
}

func (_m *ComposedProtobufStore) OnReadProtobufMatch(matchers ...interface{}) *ComposedProtobufStore_ReadProtobuf {
	c := _m.On("ReadProtobuf", matchers...)
	return &ComposedProtobufStore_ReadProtobuf{Call: c}
}

// ReadProtobuf provides a mock function with given fields: ctx, reference, msg
func (_m *ComposedProtobufStore) ReadProtobuf(ctx context.Context, reference storage.DataReference, msg protoiface.MessageV1) error {
	ret := _m.Called(ctx, reference, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, protoiface.MessageV1) error); ok {
		r0 = rf(ctx, reference, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type ComposedProtobufStore_ReadRaw struct {
	*mock.Call
}

func (_m ComposedProtobufStore_ReadRaw) Return(_a0 io.ReadCloser, _a1 error) *ComposedProtobufStore_ReadRaw {
	return &ComposedProtobufStore_ReadRaw{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ComposedProtobufStore) OnReadRaw(ctx context.Context, reference storage.DataReference) *ComposedProtobufStore_ReadRaw {
	c := _m.On("ReadRaw", ctx, reference)
	return &ComposedProtobufStore_ReadRaw{Call: c}
}

func (_m *ComposedProtobufStore) OnReadRawMatch(matchers ...interface{}) *ComposedProtobufStore_ReadRaw {
	c := _m.On("ReadRaw", matchers...)
	return &ComposedProtobufStore_ReadRaw{Call: c}
}

// ReadRaw provides a mock function with given fields: ctx, reference
func (_m *ComposedProtobufStore) ReadRaw(ctx context.Context, reference storage.DataReference) (io.ReadCloser, error) {
	ret := _m.Called(ctx, reference)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference) io.ReadCloser); ok {
		r0 = rf(ctx, reference)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, storage.DataReference) error); ok {
		r1 = rf(ctx, reference)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ComposedProtobufStore_WriteProtobuf struct {
	*mock.Call
}

func (_m ComposedProtobufStore_WriteProtobuf) Return(_a0 error) *ComposedProtobufStore_WriteProtobuf {
	return &ComposedProtobufStore_WriteProtobuf{Call: _m.Call.Return(_a0)}
}

func (_m *ComposedProtobufStore) OnWriteProtobuf(ctx context.Context, reference storage.DataReference, opts storage.Options, msg protoiface.MessageV1) *ComposedProtobufStore_WriteProtobuf {
	c := _m.On("WriteProtobuf", ctx, reference, opts, msg)
	return &ComposedProtobufStore_WriteProtobuf{Call: c}
}

func (_m *ComposedProtobufStore) OnWriteProtobufMatch(matchers ...interface{}) *ComposedProtobufStore_WriteProtobuf {
	c := _m.On("WriteProtobuf", matchers...)
	return &ComposedProtobufStore_WriteProtobuf{Call: c}
}

// WriteProtobuf provides a mock function with given fields: ctx, reference, opts, msg
func (_m *ComposedProtobufStore) WriteProtobuf(ctx context.Context, reference storage.DataReference, opts storage.Options, msg protoiface.MessageV1) error {
	ret := _m.Called(ctx, reference, opts, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, storage.Options, protoiface.MessageV1) error); ok {
		r0 = rf(ctx, reference, opts, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type ComposedProtobufStore_WriteRaw struct {
	*mock.Call
}

func (_m ComposedProtobufStore_WriteRaw) Return(_a0 error) *ComposedProtobufStore_WriteRaw {
	return &ComposedProtobufStore_WriteRaw{Call: _m.Call.Return(_a0)}
}

func (_m *ComposedProtobufStore) OnWriteRaw(ctx context.Context, reference storage.DataReference, size int64, opts storage.Options, raw io.Reader) *ComposedProtobufStore_WriteRaw {
	c := _m.On("WriteRaw", ctx, reference, size, opts, raw)
	return &ComposedProtobufStore_WriteRaw{Call: c}
}

func (_m *ComposedProtobufStore) OnWriteRawMatch(matchers ...interface{}) *ComposedProtobufStore_WriteRaw {
	c := _m.On("WriteRaw", matchers...)
	return &ComposedProtobufStore_WriteRaw{Call: c}
}

// WriteRaw provides a mock function with given fields: ctx, reference, size, opts, raw
func (_m *ComposedProtobufStore) WriteRaw(ctx context.Context, reference storage.DataReference, size int64, opts storage.Options, raw io.Reader) error {
	ret := _m.Called(ctx, reference, size, opts, raw)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, int64, storage.Options, io.Reader) error); ok {
		r0 = rf(ctx, reference, size, opts, raw)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}