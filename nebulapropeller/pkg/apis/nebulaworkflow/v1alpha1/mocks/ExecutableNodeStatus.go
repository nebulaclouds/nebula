// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	mock "github.com/stretchr/testify/mock"

	storage "github.com/nebulaclouds/nebula/nebulastdlib/storage"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
)

// ExecutableNodeStatus is an autogenerated mock type for the ExecutableNodeStatus type
type ExecutableNodeStatus struct {
	mock.Mock
}

// ClearArrayNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) ClearArrayNodeStatus() {
	_m.Called()
}

// ClearDynamicNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) ClearDynamicNodeStatus() {
	_m.Called()
}

// ClearGateNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) ClearGateNodeStatus() {
	_m.Called()
}

// ClearLastAttemptStartedAt provides a mock function with given fields:
func (_m *ExecutableNodeStatus) ClearLastAttemptStartedAt() {
	_m.Called()
}

// ClearSubNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) ClearSubNodeStatus() {
	_m.Called()
}

// ClearTaskStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) ClearTaskStatus() {
	_m.Called()
}

// ClearWorkflowStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) ClearWorkflowStatus() {
	_m.Called()
}

type ExecutableNodeStatus_GetArrayNodeStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetArrayNodeStatus) Return(_a0 v1alpha1.MutableArrayNodeStatus) *ExecutableNodeStatus_GetArrayNodeStatus {
	return &ExecutableNodeStatus_GetArrayNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetArrayNodeStatus() *ExecutableNodeStatus_GetArrayNodeStatus {
	c_call := _m.On("GetArrayNodeStatus")
	return &ExecutableNodeStatus_GetArrayNodeStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetArrayNodeStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetArrayNodeStatus {
	c_call := _m.On("GetArrayNodeStatus", matchers...)
	return &ExecutableNodeStatus_GetArrayNodeStatus{Call: c_call}
}

// GetArrayNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetArrayNodeStatus() v1alpha1.MutableArrayNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableArrayNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableArrayNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableArrayNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetAttempts struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetAttempts) Return(_a0 uint32) *ExecutableNodeStatus_GetAttempts {
	return &ExecutableNodeStatus_GetAttempts{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetAttempts() *ExecutableNodeStatus_GetAttempts {
	c_call := _m.On("GetAttempts")
	return &ExecutableNodeStatus_GetAttempts{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetAttemptsMatch(matchers ...interface{}) *ExecutableNodeStatus_GetAttempts {
	c_call := _m.On("GetAttempts", matchers...)
	return &ExecutableNodeStatus_GetAttempts{Call: c_call}
}

// GetAttempts provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetAttempts() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type ExecutableNodeStatus_GetBranchStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetBranchStatus) Return(_a0 v1alpha1.MutableBranchNodeStatus) *ExecutableNodeStatus_GetBranchStatus {
	return &ExecutableNodeStatus_GetBranchStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetBranchStatus() *ExecutableNodeStatus_GetBranchStatus {
	c_call := _m.On("GetBranchStatus")
	return &ExecutableNodeStatus_GetBranchStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetBranchStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetBranchStatus {
	c_call := _m.On("GetBranchStatus", matchers...)
	return &ExecutableNodeStatus_GetBranchStatus{Call: c_call}
}

// GetBranchStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetBranchStatus() v1alpha1.MutableBranchNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableBranchNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableBranchNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableBranchNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetDataDir struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetDataDir) Return(_a0 storage.DataReference) *ExecutableNodeStatus_GetDataDir {
	return &ExecutableNodeStatus_GetDataDir{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetDataDir() *ExecutableNodeStatus_GetDataDir {
	c_call := _m.On("GetDataDir")
	return &ExecutableNodeStatus_GetDataDir{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetDataDirMatch(matchers ...interface{}) *ExecutableNodeStatus_GetDataDir {
	c_call := _m.On("GetDataDir", matchers...)
	return &ExecutableNodeStatus_GetDataDir{Call: c_call}
}

// GetDataDir provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetDataDir() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

type ExecutableNodeStatus_GetDynamicNodeStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetDynamicNodeStatus) Return(_a0 v1alpha1.MutableDynamicNodeStatus) *ExecutableNodeStatus_GetDynamicNodeStatus {
	return &ExecutableNodeStatus_GetDynamicNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetDynamicNodeStatus() *ExecutableNodeStatus_GetDynamicNodeStatus {
	c_call := _m.On("GetDynamicNodeStatus")
	return &ExecutableNodeStatus_GetDynamicNodeStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetDynamicNodeStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetDynamicNodeStatus {
	c_call := _m.On("GetDynamicNodeStatus", matchers...)
	return &ExecutableNodeStatus_GetDynamicNodeStatus{Call: c_call}
}

// GetDynamicNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetDynamicNodeStatus() v1alpha1.MutableDynamicNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableDynamicNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableDynamicNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableDynamicNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetExecutionError struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetExecutionError) Return(_a0 *core.ExecutionError) *ExecutableNodeStatus_GetExecutionError {
	return &ExecutableNodeStatus_GetExecutionError{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetExecutionError() *ExecutableNodeStatus_GetExecutionError {
	c_call := _m.On("GetExecutionError")
	return &ExecutableNodeStatus_GetExecutionError{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetExecutionErrorMatch(matchers ...interface{}) *ExecutableNodeStatus_GetExecutionError {
	c_call := _m.On("GetExecutionError", matchers...)
	return &ExecutableNodeStatus_GetExecutionError{Call: c_call}
}

// GetExecutionError provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetExecutionError() *core.ExecutionError {
	ret := _m.Called()

	var r0 *core.ExecutionError
	if rf, ok := ret.Get(0).(func() *core.ExecutionError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ExecutionError)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetGateNodeStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetGateNodeStatus) Return(_a0 v1alpha1.MutableGateNodeStatus) *ExecutableNodeStatus_GetGateNodeStatus {
	return &ExecutableNodeStatus_GetGateNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetGateNodeStatus() *ExecutableNodeStatus_GetGateNodeStatus {
	c_call := _m.On("GetGateNodeStatus")
	return &ExecutableNodeStatus_GetGateNodeStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetGateNodeStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetGateNodeStatus {
	c_call := _m.On("GetGateNodeStatus", matchers...)
	return &ExecutableNodeStatus_GetGateNodeStatus{Call: c_call}
}

// GetGateNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetGateNodeStatus() v1alpha1.MutableGateNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableGateNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableGateNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableGateNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetLastAttemptStartedAt struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetLastAttemptStartedAt) Return(_a0 *v1.Time) *ExecutableNodeStatus_GetLastAttemptStartedAt {
	return &ExecutableNodeStatus_GetLastAttemptStartedAt{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetLastAttemptStartedAt() *ExecutableNodeStatus_GetLastAttemptStartedAt {
	c_call := _m.On("GetLastAttemptStartedAt")
	return &ExecutableNodeStatus_GetLastAttemptStartedAt{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetLastAttemptStartedAtMatch(matchers ...interface{}) *ExecutableNodeStatus_GetLastAttemptStartedAt {
	c_call := _m.On("GetLastAttemptStartedAt", matchers...)
	return &ExecutableNodeStatus_GetLastAttemptStartedAt{Call: c_call}
}

// GetLastAttemptStartedAt provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetLastAttemptStartedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetLastUpdatedAt struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetLastUpdatedAt) Return(_a0 *v1.Time) *ExecutableNodeStatus_GetLastUpdatedAt {
	return &ExecutableNodeStatus_GetLastUpdatedAt{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetLastUpdatedAt() *ExecutableNodeStatus_GetLastUpdatedAt {
	c_call := _m.On("GetLastUpdatedAt")
	return &ExecutableNodeStatus_GetLastUpdatedAt{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetLastUpdatedAtMatch(matchers ...interface{}) *ExecutableNodeStatus_GetLastUpdatedAt {
	c_call := _m.On("GetLastUpdatedAt", matchers...)
	return &ExecutableNodeStatus_GetLastUpdatedAt{Call: c_call}
}

// GetLastUpdatedAt provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetLastUpdatedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetMessage struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetMessage) Return(_a0 string) *ExecutableNodeStatus_GetMessage {
	return &ExecutableNodeStatus_GetMessage{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetMessage() *ExecutableNodeStatus_GetMessage {
	c_call := _m.On("GetMessage")
	return &ExecutableNodeStatus_GetMessage{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetMessageMatch(matchers ...interface{}) *ExecutableNodeStatus_GetMessage {
	c_call := _m.On("GetMessage", matchers...)
	return &ExecutableNodeStatus_GetMessage{Call: c_call}
}

// GetMessage provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetMessage() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ExecutableNodeStatus_GetNodeExecutionStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetNodeExecutionStatus) Return(_a0 v1alpha1.ExecutableNodeStatus) *ExecutableNodeStatus_GetNodeExecutionStatus {
	return &ExecutableNodeStatus_GetNodeExecutionStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetNodeExecutionStatus(ctx context.Context, id string) *ExecutableNodeStatus_GetNodeExecutionStatus {
	c_call := _m.On("GetNodeExecutionStatus", ctx, id)
	return &ExecutableNodeStatus_GetNodeExecutionStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetNodeExecutionStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetNodeExecutionStatus {
	c_call := _m.On("GetNodeExecutionStatus", matchers...)
	return &ExecutableNodeStatus_GetNodeExecutionStatus{Call: c_call}
}

// GetNodeExecutionStatus provides a mock function with given fields: ctx, id
func (_m *ExecutableNodeStatus) GetNodeExecutionStatus(ctx context.Context, id string) v1alpha1.ExecutableNodeStatus {
	ret := _m.Called(ctx, id)

	var r0 v1alpha1.ExecutableNodeStatus
	if rf, ok := ret.Get(0).(func(context.Context, string) v1alpha1.ExecutableNodeStatus); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetOrCreateArrayNodeStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetOrCreateArrayNodeStatus) Return(_a0 v1alpha1.MutableArrayNodeStatus) *ExecutableNodeStatus_GetOrCreateArrayNodeStatus {
	return &ExecutableNodeStatus_GetOrCreateArrayNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetOrCreateArrayNodeStatus() *ExecutableNodeStatus_GetOrCreateArrayNodeStatus {
	c_call := _m.On("GetOrCreateArrayNodeStatus")
	return &ExecutableNodeStatus_GetOrCreateArrayNodeStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetOrCreateArrayNodeStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetOrCreateArrayNodeStatus {
	c_call := _m.On("GetOrCreateArrayNodeStatus", matchers...)
	return &ExecutableNodeStatus_GetOrCreateArrayNodeStatus{Call: c_call}
}

// GetOrCreateArrayNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetOrCreateArrayNodeStatus() v1alpha1.MutableArrayNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableArrayNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableArrayNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableArrayNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetOrCreateBranchStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetOrCreateBranchStatus) Return(_a0 v1alpha1.MutableBranchNodeStatus) *ExecutableNodeStatus_GetOrCreateBranchStatus {
	return &ExecutableNodeStatus_GetOrCreateBranchStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetOrCreateBranchStatus() *ExecutableNodeStatus_GetOrCreateBranchStatus {
	c_call := _m.On("GetOrCreateBranchStatus")
	return &ExecutableNodeStatus_GetOrCreateBranchStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetOrCreateBranchStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetOrCreateBranchStatus {
	c_call := _m.On("GetOrCreateBranchStatus", matchers...)
	return &ExecutableNodeStatus_GetOrCreateBranchStatus{Call: c_call}
}

// GetOrCreateBranchStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetOrCreateBranchStatus() v1alpha1.MutableBranchNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableBranchNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableBranchNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableBranchNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetOrCreateDynamicNodeStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetOrCreateDynamicNodeStatus) Return(_a0 v1alpha1.MutableDynamicNodeStatus) *ExecutableNodeStatus_GetOrCreateDynamicNodeStatus {
	return &ExecutableNodeStatus_GetOrCreateDynamicNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetOrCreateDynamicNodeStatus() *ExecutableNodeStatus_GetOrCreateDynamicNodeStatus {
	c_call := _m.On("GetOrCreateDynamicNodeStatus")
	return &ExecutableNodeStatus_GetOrCreateDynamicNodeStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetOrCreateDynamicNodeStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetOrCreateDynamicNodeStatus {
	c_call := _m.On("GetOrCreateDynamicNodeStatus", matchers...)
	return &ExecutableNodeStatus_GetOrCreateDynamicNodeStatus{Call: c_call}
}

// GetOrCreateDynamicNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetOrCreateDynamicNodeStatus() v1alpha1.MutableDynamicNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableDynamicNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableDynamicNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableDynamicNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetOrCreateGateNodeStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetOrCreateGateNodeStatus) Return(_a0 v1alpha1.MutableGateNodeStatus) *ExecutableNodeStatus_GetOrCreateGateNodeStatus {
	return &ExecutableNodeStatus_GetOrCreateGateNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetOrCreateGateNodeStatus() *ExecutableNodeStatus_GetOrCreateGateNodeStatus {
	c_call := _m.On("GetOrCreateGateNodeStatus")
	return &ExecutableNodeStatus_GetOrCreateGateNodeStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetOrCreateGateNodeStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetOrCreateGateNodeStatus {
	c_call := _m.On("GetOrCreateGateNodeStatus", matchers...)
	return &ExecutableNodeStatus_GetOrCreateGateNodeStatus{Call: c_call}
}

// GetOrCreateGateNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetOrCreateGateNodeStatus() v1alpha1.MutableGateNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableGateNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableGateNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableGateNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetOrCreateTaskStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetOrCreateTaskStatus) Return(_a0 v1alpha1.MutableTaskNodeStatus) *ExecutableNodeStatus_GetOrCreateTaskStatus {
	return &ExecutableNodeStatus_GetOrCreateTaskStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetOrCreateTaskStatus() *ExecutableNodeStatus_GetOrCreateTaskStatus {
	c_call := _m.On("GetOrCreateTaskStatus")
	return &ExecutableNodeStatus_GetOrCreateTaskStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetOrCreateTaskStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetOrCreateTaskStatus {
	c_call := _m.On("GetOrCreateTaskStatus", matchers...)
	return &ExecutableNodeStatus_GetOrCreateTaskStatus{Call: c_call}
}

// GetOrCreateTaskStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetOrCreateTaskStatus() v1alpha1.MutableTaskNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableTaskNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableTaskNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableTaskNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetOrCreateWorkflowStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetOrCreateWorkflowStatus) Return(_a0 v1alpha1.MutableWorkflowNodeStatus) *ExecutableNodeStatus_GetOrCreateWorkflowStatus {
	return &ExecutableNodeStatus_GetOrCreateWorkflowStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetOrCreateWorkflowStatus() *ExecutableNodeStatus_GetOrCreateWorkflowStatus {
	c_call := _m.On("GetOrCreateWorkflowStatus")
	return &ExecutableNodeStatus_GetOrCreateWorkflowStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetOrCreateWorkflowStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetOrCreateWorkflowStatus {
	c_call := _m.On("GetOrCreateWorkflowStatus", matchers...)
	return &ExecutableNodeStatus_GetOrCreateWorkflowStatus{Call: c_call}
}

// GetOrCreateWorkflowStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetOrCreateWorkflowStatus() v1alpha1.MutableWorkflowNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableWorkflowNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableWorkflowNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableWorkflowNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetOutputDir struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetOutputDir) Return(_a0 storage.DataReference) *ExecutableNodeStatus_GetOutputDir {
	return &ExecutableNodeStatus_GetOutputDir{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetOutputDir() *ExecutableNodeStatus_GetOutputDir {
	c_call := _m.On("GetOutputDir")
	return &ExecutableNodeStatus_GetOutputDir{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetOutputDirMatch(matchers ...interface{}) *ExecutableNodeStatus_GetOutputDir {
	c_call := _m.On("GetOutputDir", matchers...)
	return &ExecutableNodeStatus_GetOutputDir{Call: c_call}
}

// GetOutputDir provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetOutputDir() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

type ExecutableNodeStatus_GetParentNodeID struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetParentNodeID) Return(_a0 *string) *ExecutableNodeStatus_GetParentNodeID {
	return &ExecutableNodeStatus_GetParentNodeID{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetParentNodeID() *ExecutableNodeStatus_GetParentNodeID {
	c_call := _m.On("GetParentNodeID")
	return &ExecutableNodeStatus_GetParentNodeID{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetParentNodeIDMatch(matchers ...interface{}) *ExecutableNodeStatus_GetParentNodeID {
	c_call := _m.On("GetParentNodeID", matchers...)
	return &ExecutableNodeStatus_GetParentNodeID{Call: c_call}
}

// GetParentNodeID provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetParentNodeID() *string {
	ret := _m.Called()

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetParentTaskID struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetParentTaskID) Return(_a0 *core.TaskExecutionIdentifier) *ExecutableNodeStatus_GetParentTaskID {
	return &ExecutableNodeStatus_GetParentTaskID{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetParentTaskID() *ExecutableNodeStatus_GetParentTaskID {
	c_call := _m.On("GetParentTaskID")
	return &ExecutableNodeStatus_GetParentTaskID{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetParentTaskIDMatch(matchers ...interface{}) *ExecutableNodeStatus_GetParentTaskID {
	c_call := _m.On("GetParentTaskID", matchers...)
	return &ExecutableNodeStatus_GetParentTaskID{Call: c_call}
}

// GetParentTaskID provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetParentTaskID() *core.TaskExecutionIdentifier {
	ret := _m.Called()

	var r0 *core.TaskExecutionIdentifier
	if rf, ok := ret.Get(0).(func() *core.TaskExecutionIdentifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.TaskExecutionIdentifier)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetPhase struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetPhase) Return(_a0 v1alpha1.NodePhase) *ExecutableNodeStatus_GetPhase {
	return &ExecutableNodeStatus_GetPhase{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetPhase() *ExecutableNodeStatus_GetPhase {
	c_call := _m.On("GetPhase")
	return &ExecutableNodeStatus_GetPhase{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetPhaseMatch(matchers ...interface{}) *ExecutableNodeStatus_GetPhase {
	c_call := _m.On("GetPhase", matchers...)
	return &ExecutableNodeStatus_GetPhase{Call: c_call}
}

// GetPhase provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetPhase() v1alpha1.NodePhase {
	ret := _m.Called()

	var r0 v1alpha1.NodePhase
	if rf, ok := ret.Get(0).(func() v1alpha1.NodePhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.NodePhase)
	}

	return r0
}

type ExecutableNodeStatus_GetQueuedAt struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetQueuedAt) Return(_a0 *v1.Time) *ExecutableNodeStatus_GetQueuedAt {
	return &ExecutableNodeStatus_GetQueuedAt{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetQueuedAt() *ExecutableNodeStatus_GetQueuedAt {
	c_call := _m.On("GetQueuedAt")
	return &ExecutableNodeStatus_GetQueuedAt{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetQueuedAtMatch(matchers ...interface{}) *ExecutableNodeStatus_GetQueuedAt {
	c_call := _m.On("GetQueuedAt", matchers...)
	return &ExecutableNodeStatus_GetQueuedAt{Call: c_call}
}

// GetQueuedAt provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetQueuedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetStartedAt struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetStartedAt) Return(_a0 *v1.Time) *ExecutableNodeStatus_GetStartedAt {
	return &ExecutableNodeStatus_GetStartedAt{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetStartedAt() *ExecutableNodeStatus_GetStartedAt {
	c_call := _m.On("GetStartedAt")
	return &ExecutableNodeStatus_GetStartedAt{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetStartedAtMatch(matchers ...interface{}) *ExecutableNodeStatus_GetStartedAt {
	c_call := _m.On("GetStartedAt", matchers...)
	return &ExecutableNodeStatus_GetStartedAt{Call: c_call}
}

// GetStartedAt provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetStartedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetStoppedAt struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetStoppedAt) Return(_a0 *v1.Time) *ExecutableNodeStatus_GetStoppedAt {
	return &ExecutableNodeStatus_GetStoppedAt{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetStoppedAt() *ExecutableNodeStatus_GetStoppedAt {
	c_call := _m.On("GetStoppedAt")
	return &ExecutableNodeStatus_GetStoppedAt{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetStoppedAtMatch(matchers ...interface{}) *ExecutableNodeStatus_GetStoppedAt {
	c_call := _m.On("GetStoppedAt", matchers...)
	return &ExecutableNodeStatus_GetStoppedAt{Call: c_call}
}

// GetStoppedAt provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetStoppedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetSystemFailures struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetSystemFailures) Return(_a0 uint32) *ExecutableNodeStatus_GetSystemFailures {
	return &ExecutableNodeStatus_GetSystemFailures{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetSystemFailures() *ExecutableNodeStatus_GetSystemFailures {
	c_call := _m.On("GetSystemFailures")
	return &ExecutableNodeStatus_GetSystemFailures{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetSystemFailuresMatch(matchers ...interface{}) *ExecutableNodeStatus_GetSystemFailures {
	c_call := _m.On("GetSystemFailures", matchers...)
	return &ExecutableNodeStatus_GetSystemFailures{Call: c_call}
}

// GetSystemFailures provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetSystemFailures() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type ExecutableNodeStatus_GetTaskNodeStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetTaskNodeStatus) Return(_a0 v1alpha1.ExecutableTaskNodeStatus) *ExecutableNodeStatus_GetTaskNodeStatus {
	return &ExecutableNodeStatus_GetTaskNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetTaskNodeStatus() *ExecutableNodeStatus_GetTaskNodeStatus {
	c_call := _m.On("GetTaskNodeStatus")
	return &ExecutableNodeStatus_GetTaskNodeStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetTaskNodeStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetTaskNodeStatus {
	c_call := _m.On("GetTaskNodeStatus", matchers...)
	return &ExecutableNodeStatus_GetTaskNodeStatus{Call: c_call}
}

// GetTaskNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetTaskNodeStatus() v1alpha1.ExecutableTaskNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableTaskNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableTaskNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableTaskNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetTaskStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetTaskStatus) Return(_a0 v1alpha1.MutableTaskNodeStatus) *ExecutableNodeStatus_GetTaskStatus {
	return &ExecutableNodeStatus_GetTaskStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetTaskStatus() *ExecutableNodeStatus_GetTaskStatus {
	c_call := _m.On("GetTaskStatus")
	return &ExecutableNodeStatus_GetTaskStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetTaskStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetTaskStatus {
	c_call := _m.On("GetTaskStatus", matchers...)
	return &ExecutableNodeStatus_GetTaskStatus{Call: c_call}
}

// GetTaskStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetTaskStatus() v1alpha1.MutableTaskNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableTaskNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableTaskNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableTaskNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetWorkflowNodeStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetWorkflowNodeStatus) Return(_a0 v1alpha1.ExecutableWorkflowNodeStatus) *ExecutableNodeStatus_GetWorkflowNodeStatus {
	return &ExecutableNodeStatus_GetWorkflowNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetWorkflowNodeStatus() *ExecutableNodeStatus_GetWorkflowNodeStatus {
	c_call := _m.On("GetWorkflowNodeStatus")
	return &ExecutableNodeStatus_GetWorkflowNodeStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetWorkflowNodeStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetWorkflowNodeStatus {
	c_call := _m.On("GetWorkflowNodeStatus", matchers...)
	return &ExecutableNodeStatus_GetWorkflowNodeStatus{Call: c_call}
}

// GetWorkflowNodeStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetWorkflowNodeStatus() v1alpha1.ExecutableWorkflowNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableWorkflowNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableWorkflowNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableWorkflowNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_GetWorkflowStatus struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_GetWorkflowStatus) Return(_a0 v1alpha1.MutableWorkflowNodeStatus) *ExecutableNodeStatus_GetWorkflowStatus {
	return &ExecutableNodeStatus_GetWorkflowStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnGetWorkflowStatus() *ExecutableNodeStatus_GetWorkflowStatus {
	c_call := _m.On("GetWorkflowStatus")
	return &ExecutableNodeStatus_GetWorkflowStatus{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnGetWorkflowStatusMatch(matchers ...interface{}) *ExecutableNodeStatus_GetWorkflowStatus {
	c_call := _m.On("GetWorkflowStatus", matchers...)
	return &ExecutableNodeStatus_GetWorkflowStatus{Call: c_call}
}

// GetWorkflowStatus provides a mock function with given fields:
func (_m *ExecutableNodeStatus) GetWorkflowStatus() v1alpha1.MutableWorkflowNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableWorkflowNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableWorkflowNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableWorkflowNodeStatus)
		}
	}

	return r0
}

type ExecutableNodeStatus_IncrementAttempts struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_IncrementAttempts) Return(_a0 uint32) *ExecutableNodeStatus_IncrementAttempts {
	return &ExecutableNodeStatus_IncrementAttempts{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnIncrementAttempts() *ExecutableNodeStatus_IncrementAttempts {
	c_call := _m.On("IncrementAttempts")
	return &ExecutableNodeStatus_IncrementAttempts{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnIncrementAttemptsMatch(matchers ...interface{}) *ExecutableNodeStatus_IncrementAttempts {
	c_call := _m.On("IncrementAttempts", matchers...)
	return &ExecutableNodeStatus_IncrementAttempts{Call: c_call}
}

// IncrementAttempts provides a mock function with given fields:
func (_m *ExecutableNodeStatus) IncrementAttempts() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type ExecutableNodeStatus_IncrementSystemFailures struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_IncrementSystemFailures) Return(_a0 uint32) *ExecutableNodeStatus_IncrementSystemFailures {
	return &ExecutableNodeStatus_IncrementSystemFailures{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnIncrementSystemFailures() *ExecutableNodeStatus_IncrementSystemFailures {
	c_call := _m.On("IncrementSystemFailures")
	return &ExecutableNodeStatus_IncrementSystemFailures{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnIncrementSystemFailuresMatch(matchers ...interface{}) *ExecutableNodeStatus_IncrementSystemFailures {
	c_call := _m.On("IncrementSystemFailures", matchers...)
	return &ExecutableNodeStatus_IncrementSystemFailures{Call: c_call}
}

// IncrementSystemFailures provides a mock function with given fields:
func (_m *ExecutableNodeStatus) IncrementSystemFailures() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type ExecutableNodeStatus_IsCached struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_IsCached) Return(_a0 bool) *ExecutableNodeStatus_IsCached {
	return &ExecutableNodeStatus_IsCached{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnIsCached() *ExecutableNodeStatus_IsCached {
	c_call := _m.On("IsCached")
	return &ExecutableNodeStatus_IsCached{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnIsCachedMatch(matchers ...interface{}) *ExecutableNodeStatus_IsCached {
	c_call := _m.On("IsCached", matchers...)
	return &ExecutableNodeStatus_IsCached{Call: c_call}
}

// IsCached provides a mock function with given fields:
func (_m *ExecutableNodeStatus) IsCached() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type ExecutableNodeStatus_IsDirty struct {
	*mock.Call
}

func (_m ExecutableNodeStatus_IsDirty) Return(_a0 bool) *ExecutableNodeStatus_IsDirty {
	return &ExecutableNodeStatus_IsDirty{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNodeStatus) OnIsDirty() *ExecutableNodeStatus_IsDirty {
	c_call := _m.On("IsDirty")
	return &ExecutableNodeStatus_IsDirty{Call: c_call}
}

func (_m *ExecutableNodeStatus) OnIsDirtyMatch(matchers ...interface{}) *ExecutableNodeStatus_IsDirty {
	c_call := _m.On("IsDirty", matchers...)
	return &ExecutableNodeStatus_IsDirty{Call: c_call}
}

// IsDirty provides a mock function with given fields:
func (_m *ExecutableNodeStatus) IsDirty() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ResetDirty provides a mock function with given fields:
func (_m *ExecutableNodeStatus) ResetDirty() {
	_m.Called()
}

// SetCached provides a mock function with given fields:
func (_m *ExecutableNodeStatus) SetCached() {
	_m.Called()
}

// SetDataDir provides a mock function with given fields: _a0
func (_m *ExecutableNodeStatus) SetDataDir(_a0 storage.DataReference) {
	_m.Called(_a0)
}

// SetOutputDir provides a mock function with given fields: d
func (_m *ExecutableNodeStatus) SetOutputDir(d storage.DataReference) {
	_m.Called(d)
}

// SetParentNodeID provides a mock function with given fields: n
func (_m *ExecutableNodeStatus) SetParentNodeID(n *string) {
	_m.Called(n)
}

// SetParentTaskID provides a mock function with given fields: t
func (_m *ExecutableNodeStatus) SetParentTaskID(t *core.TaskExecutionIdentifier) {
	_m.Called(t)
}

// UpdatePhase provides a mock function with given fields: phase, occurredAt, reason, err
func (_m *ExecutableNodeStatus) UpdatePhase(phase v1alpha1.NodePhase, occurredAt v1.Time, reason string, err *core.ExecutionError) {
	_m.Called(phase, occurredAt, reason, err)
}

// VisitNodeStatuses provides a mock function with given fields: visitor
func (_m *ExecutableNodeStatus) VisitNodeStatuses(visitor func(string, v1alpha1.ExecutableNodeStatus)) {
	_m.Called(visitor)
}
