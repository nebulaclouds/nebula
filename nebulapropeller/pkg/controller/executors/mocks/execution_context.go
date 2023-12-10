// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	executors "github.com/nebulaclouds/nebula/nebulapropeller/pkg/controller/executors"

	mock "github.com/stretchr/testify/mock"

	types "k8s.io/apimachinery/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/nebulaclouds/nebula/nebulapropeller/pkg/apis/nebulaworkflow/v1alpha1"
)

// ExecutionContext is an autogenerated mock type for the ExecutionContext type
type ExecutionContext struct {
	mock.Mock
}

type ExecutionContext_CurrentParallelism struct {
	*mock.Call
}

func (_m ExecutionContext_CurrentParallelism) Return(_a0 uint32) *ExecutionContext_CurrentParallelism {
	return &ExecutionContext_CurrentParallelism{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnCurrentParallelism() *ExecutionContext_CurrentParallelism {
	c_call := _m.On("CurrentParallelism")
	return &ExecutionContext_CurrentParallelism{Call: c_call}
}

func (_m *ExecutionContext) OnCurrentParallelismMatch(matchers ...interface{}) *ExecutionContext_CurrentParallelism {
	c_call := _m.On("CurrentParallelism", matchers...)
	return &ExecutionContext_CurrentParallelism{Call: c_call}
}

// CurrentParallelism provides a mock function with given fields:
func (_m *ExecutionContext) CurrentParallelism() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type ExecutionContext_FindSubWorkflow struct {
	*mock.Call
}

func (_m ExecutionContext_FindSubWorkflow) Return(_a0 v1alpha1.ExecutableSubWorkflow) *ExecutionContext_FindSubWorkflow {
	return &ExecutionContext_FindSubWorkflow{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnFindSubWorkflow(subID string) *ExecutionContext_FindSubWorkflow {
	c_call := _m.On("FindSubWorkflow", subID)
	return &ExecutionContext_FindSubWorkflow{Call: c_call}
}

func (_m *ExecutionContext) OnFindSubWorkflowMatch(matchers ...interface{}) *ExecutionContext_FindSubWorkflow {
	c_call := _m.On("FindSubWorkflow", matchers...)
	return &ExecutionContext_FindSubWorkflow{Call: c_call}
}

// FindSubWorkflow provides a mock function with given fields: subID
func (_m *ExecutionContext) FindSubWorkflow(subID string) v1alpha1.ExecutableSubWorkflow {
	ret := _m.Called(subID)

	var r0 v1alpha1.ExecutableSubWorkflow
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableSubWorkflow); ok {
		r0 = rf(subID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableSubWorkflow)
		}
	}

	return r0
}

type ExecutionContext_GetAnnotations struct {
	*mock.Call
}

func (_m ExecutionContext_GetAnnotations) Return(_a0 map[string]string) *ExecutionContext_GetAnnotations {
	return &ExecutionContext_GetAnnotations{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetAnnotations() *ExecutionContext_GetAnnotations {
	c_call := _m.On("GetAnnotations")
	return &ExecutionContext_GetAnnotations{Call: c_call}
}

func (_m *ExecutionContext) OnGetAnnotationsMatch(matchers ...interface{}) *ExecutionContext_GetAnnotations {
	c_call := _m.On("GetAnnotations", matchers...)
	return &ExecutionContext_GetAnnotations{Call: c_call}
}

// GetAnnotations provides a mock function with given fields:
func (_m *ExecutionContext) GetAnnotations() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

type ExecutionContext_GetCreationTimestamp struct {
	*mock.Call
}

func (_m ExecutionContext_GetCreationTimestamp) Return(_a0 v1.Time) *ExecutionContext_GetCreationTimestamp {
	return &ExecutionContext_GetCreationTimestamp{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetCreationTimestamp() *ExecutionContext_GetCreationTimestamp {
	c_call := _m.On("GetCreationTimestamp")
	return &ExecutionContext_GetCreationTimestamp{Call: c_call}
}

func (_m *ExecutionContext) OnGetCreationTimestampMatch(matchers ...interface{}) *ExecutionContext_GetCreationTimestamp {
	c_call := _m.On("GetCreationTimestamp", matchers...)
	return &ExecutionContext_GetCreationTimestamp{Call: c_call}
}

// GetCreationTimestamp provides a mock function with given fields:
func (_m *ExecutionContext) GetCreationTimestamp() v1.Time {
	ret := _m.Called()

	var r0 v1.Time
	if rf, ok := ret.Get(0).(func() v1.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.Time)
	}

	return r0
}

type ExecutionContext_GetDefinitionVersion struct {
	*mock.Call
}

func (_m ExecutionContext_GetDefinitionVersion) Return(_a0 v1alpha1.WorkflowDefinitionVersion) *ExecutionContext_GetDefinitionVersion {
	return &ExecutionContext_GetDefinitionVersion{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetDefinitionVersion() *ExecutionContext_GetDefinitionVersion {
	c_call := _m.On("GetDefinitionVersion")
	return &ExecutionContext_GetDefinitionVersion{Call: c_call}
}

func (_m *ExecutionContext) OnGetDefinitionVersionMatch(matchers ...interface{}) *ExecutionContext_GetDefinitionVersion {
	c_call := _m.On("GetDefinitionVersion", matchers...)
	return &ExecutionContext_GetDefinitionVersion{Call: c_call}
}

// GetDefinitionVersion provides a mock function with given fields:
func (_m *ExecutionContext) GetDefinitionVersion() v1alpha1.WorkflowDefinitionVersion {
	ret := _m.Called()

	var r0 v1alpha1.WorkflowDefinitionVersion
	if rf, ok := ret.Get(0).(func() v1alpha1.WorkflowDefinitionVersion); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.WorkflowDefinitionVersion)
	}

	return r0
}

type ExecutionContext_GetEventVersion struct {
	*mock.Call
}

func (_m ExecutionContext_GetEventVersion) Return(_a0 v1alpha1.EventVersion) *ExecutionContext_GetEventVersion {
	return &ExecutionContext_GetEventVersion{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetEventVersion() *ExecutionContext_GetEventVersion {
	c_call := _m.On("GetEventVersion")
	return &ExecutionContext_GetEventVersion{Call: c_call}
}

func (_m *ExecutionContext) OnGetEventVersionMatch(matchers ...interface{}) *ExecutionContext_GetEventVersion {
	c_call := _m.On("GetEventVersion", matchers...)
	return &ExecutionContext_GetEventVersion{Call: c_call}
}

// GetEventVersion provides a mock function with given fields:
func (_m *ExecutionContext) GetEventVersion() v1alpha1.EventVersion {
	ret := _m.Called()

	var r0 v1alpha1.EventVersion
	if rf, ok := ret.Get(0).(func() v1alpha1.EventVersion); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.EventVersion)
	}

	return r0
}

type ExecutionContext_GetExecutionConfig struct {
	*mock.Call
}

func (_m ExecutionContext_GetExecutionConfig) Return(_a0 v1alpha1.ExecutionConfig) *ExecutionContext_GetExecutionConfig {
	return &ExecutionContext_GetExecutionConfig{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetExecutionConfig() *ExecutionContext_GetExecutionConfig {
	c_call := _m.On("GetExecutionConfig")
	return &ExecutionContext_GetExecutionConfig{Call: c_call}
}

func (_m *ExecutionContext) OnGetExecutionConfigMatch(matchers ...interface{}) *ExecutionContext_GetExecutionConfig {
	c_call := _m.On("GetExecutionConfig", matchers...)
	return &ExecutionContext_GetExecutionConfig{Call: c_call}
}

// GetExecutionConfig provides a mock function with given fields:
func (_m *ExecutionContext) GetExecutionConfig() v1alpha1.ExecutionConfig {
	ret := _m.Called()

	var r0 v1alpha1.ExecutionConfig
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutionConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.ExecutionConfig)
	}

	return r0
}

type ExecutionContext_GetExecutionID struct {
	*mock.Call
}

func (_m ExecutionContext_GetExecutionID) Return(_a0 v1alpha1.WorkflowExecutionIdentifier) *ExecutionContext_GetExecutionID {
	return &ExecutionContext_GetExecutionID{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetExecutionID() *ExecutionContext_GetExecutionID {
	c_call := _m.On("GetExecutionID")
	return &ExecutionContext_GetExecutionID{Call: c_call}
}

func (_m *ExecutionContext) OnGetExecutionIDMatch(matchers ...interface{}) *ExecutionContext_GetExecutionID {
	c_call := _m.On("GetExecutionID", matchers...)
	return &ExecutionContext_GetExecutionID{Call: c_call}
}

// GetExecutionID provides a mock function with given fields:
func (_m *ExecutionContext) GetExecutionID() v1alpha1.WorkflowExecutionIdentifier {
	ret := _m.Called()

	var r0 v1alpha1.WorkflowExecutionIdentifier
	if rf, ok := ret.Get(0).(func() v1alpha1.WorkflowExecutionIdentifier); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.WorkflowExecutionIdentifier)
	}

	return r0
}

type ExecutionContext_GetID struct {
	*mock.Call
}

func (_m ExecutionContext_GetID) Return(_a0 string) *ExecutionContext_GetID {
	return &ExecutionContext_GetID{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetID() *ExecutionContext_GetID {
	c_call := _m.On("GetID")
	return &ExecutionContext_GetID{Call: c_call}
}

func (_m *ExecutionContext) OnGetIDMatch(matchers ...interface{}) *ExecutionContext_GetID {
	c_call := _m.On("GetID", matchers...)
	return &ExecutionContext_GetID{Call: c_call}
}

// GetID provides a mock function with given fields:
func (_m *ExecutionContext) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ExecutionContext_GetK8sWorkflowID struct {
	*mock.Call
}

func (_m ExecutionContext_GetK8sWorkflowID) Return(_a0 types.NamespacedName) *ExecutionContext_GetK8sWorkflowID {
	return &ExecutionContext_GetK8sWorkflowID{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetK8sWorkflowID() *ExecutionContext_GetK8sWorkflowID {
	c_call := _m.On("GetK8sWorkflowID")
	return &ExecutionContext_GetK8sWorkflowID{Call: c_call}
}

func (_m *ExecutionContext) OnGetK8sWorkflowIDMatch(matchers ...interface{}) *ExecutionContext_GetK8sWorkflowID {
	c_call := _m.On("GetK8sWorkflowID", matchers...)
	return &ExecutionContext_GetK8sWorkflowID{Call: c_call}
}

// GetK8sWorkflowID provides a mock function with given fields:
func (_m *ExecutionContext) GetK8sWorkflowID() types.NamespacedName {
	ret := _m.Called()

	var r0 types.NamespacedName
	if rf, ok := ret.Get(0).(func() types.NamespacedName); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.NamespacedName)
	}

	return r0
}

type ExecutionContext_GetLabels struct {
	*mock.Call
}

func (_m ExecutionContext_GetLabels) Return(_a0 map[string]string) *ExecutionContext_GetLabels {
	return &ExecutionContext_GetLabels{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetLabels() *ExecutionContext_GetLabels {
	c_call := _m.On("GetLabels")
	return &ExecutionContext_GetLabels{Call: c_call}
}

func (_m *ExecutionContext) OnGetLabelsMatch(matchers ...interface{}) *ExecutionContext_GetLabels {
	c_call := _m.On("GetLabels", matchers...)
	return &ExecutionContext_GetLabels{Call: c_call}
}

// GetLabels provides a mock function with given fields:
func (_m *ExecutionContext) GetLabels() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

type ExecutionContext_GetName struct {
	*mock.Call
}

func (_m ExecutionContext_GetName) Return(_a0 string) *ExecutionContext_GetName {
	return &ExecutionContext_GetName{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetName() *ExecutionContext_GetName {
	c_call := _m.On("GetName")
	return &ExecutionContext_GetName{Call: c_call}
}

func (_m *ExecutionContext) OnGetNameMatch(matchers ...interface{}) *ExecutionContext_GetName {
	c_call := _m.On("GetName", matchers...)
	return &ExecutionContext_GetName{Call: c_call}
}

// GetName provides a mock function with given fields:
func (_m *ExecutionContext) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ExecutionContext_GetNamespace struct {
	*mock.Call
}

func (_m ExecutionContext_GetNamespace) Return(_a0 string) *ExecutionContext_GetNamespace {
	return &ExecutionContext_GetNamespace{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetNamespace() *ExecutionContext_GetNamespace {
	c_call := _m.On("GetNamespace")
	return &ExecutionContext_GetNamespace{Call: c_call}
}

func (_m *ExecutionContext) OnGetNamespaceMatch(matchers ...interface{}) *ExecutionContext_GetNamespace {
	c_call := _m.On("GetNamespace", matchers...)
	return &ExecutionContext_GetNamespace{Call: c_call}
}

// GetNamespace provides a mock function with given fields:
func (_m *ExecutionContext) GetNamespace() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ExecutionContext_GetOnFailurePolicy struct {
	*mock.Call
}

func (_m ExecutionContext_GetOnFailurePolicy) Return(_a0 v1alpha1.WorkflowOnFailurePolicy) *ExecutionContext_GetOnFailurePolicy {
	return &ExecutionContext_GetOnFailurePolicy{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetOnFailurePolicy() *ExecutionContext_GetOnFailurePolicy {
	c_call := _m.On("GetOnFailurePolicy")
	return &ExecutionContext_GetOnFailurePolicy{Call: c_call}
}

func (_m *ExecutionContext) OnGetOnFailurePolicyMatch(matchers ...interface{}) *ExecutionContext_GetOnFailurePolicy {
	c_call := _m.On("GetOnFailurePolicy", matchers...)
	return &ExecutionContext_GetOnFailurePolicy{Call: c_call}
}

// GetOnFailurePolicy provides a mock function with given fields:
func (_m *ExecutionContext) GetOnFailurePolicy() v1alpha1.WorkflowOnFailurePolicy {
	ret := _m.Called()

	var r0 v1alpha1.WorkflowOnFailurePolicy
	if rf, ok := ret.Get(0).(func() v1alpha1.WorkflowOnFailurePolicy); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.WorkflowOnFailurePolicy)
	}

	return r0
}

type ExecutionContext_GetOwnerReference struct {
	*mock.Call
}

func (_m ExecutionContext_GetOwnerReference) Return(_a0 v1.OwnerReference) *ExecutionContext_GetOwnerReference {
	return &ExecutionContext_GetOwnerReference{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetOwnerReference() *ExecutionContext_GetOwnerReference {
	c_call := _m.On("GetOwnerReference")
	return &ExecutionContext_GetOwnerReference{Call: c_call}
}

func (_m *ExecutionContext) OnGetOwnerReferenceMatch(matchers ...interface{}) *ExecutionContext_GetOwnerReference {
	c_call := _m.On("GetOwnerReference", matchers...)
	return &ExecutionContext_GetOwnerReference{Call: c_call}
}

// GetOwnerReference provides a mock function with given fields:
func (_m *ExecutionContext) GetOwnerReference() v1.OwnerReference {
	ret := _m.Called()

	var r0 v1.OwnerReference
	if rf, ok := ret.Get(0).(func() v1.OwnerReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1.OwnerReference)
	}

	return r0
}

type ExecutionContext_GetParentInfo struct {
	*mock.Call
}

func (_m ExecutionContext_GetParentInfo) Return(_a0 executors.ImmutableParentInfo) *ExecutionContext_GetParentInfo {
	return &ExecutionContext_GetParentInfo{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetParentInfo() *ExecutionContext_GetParentInfo {
	c_call := _m.On("GetParentInfo")
	return &ExecutionContext_GetParentInfo{Call: c_call}
}

func (_m *ExecutionContext) OnGetParentInfoMatch(matchers ...interface{}) *ExecutionContext_GetParentInfo {
	c_call := _m.On("GetParentInfo", matchers...)
	return &ExecutionContext_GetParentInfo{Call: c_call}
}

// GetParentInfo provides a mock function with given fields:
func (_m *ExecutionContext) GetParentInfo() executors.ImmutableParentInfo {
	ret := _m.Called()

	var r0 executors.ImmutableParentInfo
	if rf, ok := ret.Get(0).(func() executors.ImmutableParentInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(executors.ImmutableParentInfo)
		}
	}

	return r0
}

type ExecutionContext_GetRawOutputDataConfig struct {
	*mock.Call
}

func (_m ExecutionContext_GetRawOutputDataConfig) Return(_a0 v1alpha1.RawOutputDataConfig) *ExecutionContext_GetRawOutputDataConfig {
	return &ExecutionContext_GetRawOutputDataConfig{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetRawOutputDataConfig() *ExecutionContext_GetRawOutputDataConfig {
	c_call := _m.On("GetRawOutputDataConfig")
	return &ExecutionContext_GetRawOutputDataConfig{Call: c_call}
}

func (_m *ExecutionContext) OnGetRawOutputDataConfigMatch(matchers ...interface{}) *ExecutionContext_GetRawOutputDataConfig {
	c_call := _m.On("GetRawOutputDataConfig", matchers...)
	return &ExecutionContext_GetRawOutputDataConfig{Call: c_call}
}

// GetRawOutputDataConfig provides a mock function with given fields:
func (_m *ExecutionContext) GetRawOutputDataConfig() v1alpha1.RawOutputDataConfig {
	ret := _m.Called()

	var r0 v1alpha1.RawOutputDataConfig
	if rf, ok := ret.Get(0).(func() v1alpha1.RawOutputDataConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.RawOutputDataConfig)
	}

	return r0
}

type ExecutionContext_GetSecurityContext struct {
	*mock.Call
}

func (_m ExecutionContext_GetSecurityContext) Return(_a0 core.SecurityContext) *ExecutionContext_GetSecurityContext {
	return &ExecutionContext_GetSecurityContext{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetSecurityContext() *ExecutionContext_GetSecurityContext {
	c_call := _m.On("GetSecurityContext")
	return &ExecutionContext_GetSecurityContext{Call: c_call}
}

func (_m *ExecutionContext) OnGetSecurityContextMatch(matchers ...interface{}) *ExecutionContext_GetSecurityContext {
	c_call := _m.On("GetSecurityContext", matchers...)
	return &ExecutionContext_GetSecurityContext{Call: c_call}
}

// GetSecurityContext provides a mock function with given fields:
func (_m *ExecutionContext) GetSecurityContext() core.SecurityContext {
	ret := _m.Called()

	var r0 core.SecurityContext
	if rf, ok := ret.Get(0).(func() core.SecurityContext); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(core.SecurityContext)
	}

	return r0
}

type ExecutionContext_GetServiceAccountName struct {
	*mock.Call
}

func (_m ExecutionContext_GetServiceAccountName) Return(_a0 string) *ExecutionContext_GetServiceAccountName {
	return &ExecutionContext_GetServiceAccountName{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnGetServiceAccountName() *ExecutionContext_GetServiceAccountName {
	c_call := _m.On("GetServiceAccountName")
	return &ExecutionContext_GetServiceAccountName{Call: c_call}
}

func (_m *ExecutionContext) OnGetServiceAccountNameMatch(matchers ...interface{}) *ExecutionContext_GetServiceAccountName {
	c_call := _m.On("GetServiceAccountName", matchers...)
	return &ExecutionContext_GetServiceAccountName{Call: c_call}
}

// GetServiceAccountName provides a mock function with given fields:
func (_m *ExecutionContext) GetServiceAccountName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ExecutionContext_GetTask struct {
	*mock.Call
}

func (_m ExecutionContext_GetTask) Return(_a0 v1alpha1.ExecutableTask, _a1 error) *ExecutionContext_GetTask {
	return &ExecutionContext_GetTask{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ExecutionContext) OnGetTask(id string) *ExecutionContext_GetTask {
	c_call := _m.On("GetTask", id)
	return &ExecutionContext_GetTask{Call: c_call}
}

func (_m *ExecutionContext) OnGetTaskMatch(matchers ...interface{}) *ExecutionContext_GetTask {
	c_call := _m.On("GetTask", matchers...)
	return &ExecutionContext_GetTask{Call: c_call}
}

// GetTask provides a mock function with given fields: id
func (_m *ExecutionContext) GetTask(id string) (v1alpha1.ExecutableTask, error) {
	ret := _m.Called(id)

	var r0 v1alpha1.ExecutableTask
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableTask); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableTask)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ExecutionContext_IncrementParallelism struct {
	*mock.Call
}

func (_m ExecutionContext_IncrementParallelism) Return(_a0 uint32) *ExecutionContext_IncrementParallelism {
	return &ExecutionContext_IncrementParallelism{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnIncrementParallelism() *ExecutionContext_IncrementParallelism {
	c_call := _m.On("IncrementParallelism")
	return &ExecutionContext_IncrementParallelism{Call: c_call}
}

func (_m *ExecutionContext) OnIncrementParallelismMatch(matchers ...interface{}) *ExecutionContext_IncrementParallelism {
	c_call := _m.On("IncrementParallelism", matchers...)
	return &ExecutionContext_IncrementParallelism{Call: c_call}
}

// IncrementParallelism provides a mock function with given fields:
func (_m *ExecutionContext) IncrementParallelism() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type ExecutionContext_IsInterruptible struct {
	*mock.Call
}

func (_m ExecutionContext_IsInterruptible) Return(_a0 bool) *ExecutionContext_IsInterruptible {
	return &ExecutionContext_IsInterruptible{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutionContext) OnIsInterruptible() *ExecutionContext_IsInterruptible {
	c_call := _m.On("IsInterruptible")
	return &ExecutionContext_IsInterruptible{Call: c_call}
}

func (_m *ExecutionContext) OnIsInterruptibleMatch(matchers ...interface{}) *ExecutionContext_IsInterruptible {
	c_call := _m.On("IsInterruptible", matchers...)
	return &ExecutionContext_IsInterruptible{Call: c_call}
}

// IsInterruptible provides a mock function with given fields:
func (_m *ExecutionContext) IsInterruptible() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
