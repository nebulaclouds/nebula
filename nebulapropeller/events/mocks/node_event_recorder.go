// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	config "github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/config"

	event "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/event"

	mock "github.com/stretchr/testify/mock"
)

// NodeEventRecorder is an autogenerated mock type for the NodeEventRecorder type
type NodeEventRecorder struct {
	mock.Mock
}

type NodeEventRecorder_RecordNodeEvent struct {
	*mock.Call
}

func (_m NodeEventRecorder_RecordNodeEvent) Return(_a0 error) *NodeEventRecorder_RecordNodeEvent {
	return &NodeEventRecorder_RecordNodeEvent{Call: _m.Call.Return(_a0)}
}

func (_m *NodeEventRecorder) OnRecordNodeEvent(ctx context.Context, _a1 *event.NodeExecutionEvent, eventConfig *config.EventConfig) *NodeEventRecorder_RecordNodeEvent {
	c_call := _m.On("RecordNodeEvent", ctx, _a1, eventConfig)
	return &NodeEventRecorder_RecordNodeEvent{Call: c_call}
}

func (_m *NodeEventRecorder) OnRecordNodeEventMatch(matchers ...interface{}) *NodeEventRecorder_RecordNodeEvent {
	c_call := _m.On("RecordNodeEvent", matchers...)
	return &NodeEventRecorder_RecordNodeEvent{Call: c_call}
}

// RecordNodeEvent provides a mock function with given fields: ctx, _a1, eventConfig
func (_m *NodeEventRecorder) RecordNodeEvent(ctx context.Context, _a1 *event.NodeExecutionEvent, eventConfig *config.EventConfig) error {
	ret := _m.Called(ctx, _a1, eventConfig)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *event.NodeExecutionEvent, *config.EventConfig) error); ok {
		r0 = rf(ctx, _a1, eventConfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
