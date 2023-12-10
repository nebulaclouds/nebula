// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/admin/task_execution.proto

package admin

import (
	fmt "fmt"
	core "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	event "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/event"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A message used to fetch a single task execution entity.
// See :ref:`ref_nebulaidl.admin.TaskExecution` for more details
type TaskExecutionGetRequest struct {
	// Unique identifier for the task execution.
	// +required
	Id                   *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *TaskExecutionGetRequest) Reset()         { *m = TaskExecutionGetRequest{} }
func (m *TaskExecutionGetRequest) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionGetRequest) ProtoMessage()    {}
func (*TaskExecutionGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cde4c3aa101642e, []int{0}
}

func (m *TaskExecutionGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionGetRequest.Unmarshal(m, b)
}
func (m *TaskExecutionGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionGetRequest.Marshal(b, m, deterministic)
}
func (m *TaskExecutionGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionGetRequest.Merge(m, src)
}
func (m *TaskExecutionGetRequest) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionGetRequest.Size(m)
}
func (m *TaskExecutionGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionGetRequest proto.InternalMessageInfo

func (m *TaskExecutionGetRequest) GetId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

// Represents a request structure to retrieve a list of task execution entities yielded by a specific node execution.
// See :ref:`ref_nebulaidl.admin.TaskExecution` for more details
type TaskExecutionListRequest struct {
	// Indicates the node execution to filter by.
	// +required
	NodeExecutionId *core.NodeExecutionIdentifier `protobuf:"bytes,1,opt,name=node_execution_id,json=nodeExecutionId,proto3" json:"node_execution_id,omitempty"`
	// Indicates the number of resources to be returned.
	// +required
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query.
	// +optional
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	// Indicates a list of filters passed as string.
	// More info on constructing filters : <Link>
	// +optional
	Filters string `protobuf:"bytes,4,opt,name=filters,proto3" json:"filters,omitempty"`
	// Sort ordering for returned list.
	// +optional
	SortBy               *Sort    `protobuf:"bytes,5,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecutionListRequest) Reset()         { *m = TaskExecutionListRequest{} }
func (m *TaskExecutionListRequest) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionListRequest) ProtoMessage()    {}
func (*TaskExecutionListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cde4c3aa101642e, []int{1}
}

func (m *TaskExecutionListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionListRequest.Unmarshal(m, b)
}
func (m *TaskExecutionListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionListRequest.Marshal(b, m, deterministic)
}
func (m *TaskExecutionListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionListRequest.Merge(m, src)
}
func (m *TaskExecutionListRequest) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionListRequest.Size(m)
}
func (m *TaskExecutionListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionListRequest proto.InternalMessageInfo

func (m *TaskExecutionListRequest) GetNodeExecutionId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.NodeExecutionId
	}
	return nil
}

func (m *TaskExecutionListRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TaskExecutionListRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *TaskExecutionListRequest) GetFilters() string {
	if m != nil {
		return m.Filters
	}
	return ""
}

func (m *TaskExecutionListRequest) GetSortBy() *Sort {
	if m != nil {
		return m.SortBy
	}
	return nil
}

// Encapsulates all details for a single task execution entity.
// A task execution represents an instantiated task, including all inputs and additional
// metadata as well as computed results included state, outputs, and duration-based attributes.
type TaskExecution struct {
	// Unique identifier for the task execution.
	Id *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Path to remote data store where input blob is stored.
	InputUri string `protobuf:"bytes,2,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// Task execution details and results.
	Closure *TaskExecutionClosure `protobuf:"bytes,3,opt,name=closure,proto3" json:"closure,omitempty"`
	// Whether this task spawned nodes.
	IsParent             bool     `protobuf:"varint,4,opt,name=is_parent,json=isParent,proto3" json:"is_parent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecution) Reset()         { *m = TaskExecution{} }
func (m *TaskExecution) String() string { return proto.CompactTextString(m) }
func (*TaskExecution) ProtoMessage()    {}
func (*TaskExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cde4c3aa101642e, []int{2}
}

func (m *TaskExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecution.Unmarshal(m, b)
}
func (m *TaskExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecution.Marshal(b, m, deterministic)
}
func (m *TaskExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecution.Merge(m, src)
}
func (m *TaskExecution) XXX_Size() int {
	return xxx_messageInfo_TaskExecution.Size(m)
}
func (m *TaskExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecution.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecution proto.InternalMessageInfo

func (m *TaskExecution) GetId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TaskExecution) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

func (m *TaskExecution) GetClosure() *TaskExecutionClosure {
	if m != nil {
		return m.Closure
	}
	return nil
}

func (m *TaskExecution) GetIsParent() bool {
	if m != nil {
		return m.IsParent
	}
	return false
}

// Response structure for a query to list of task execution entities.
// See :ref:`ref_nebulaidl.admin.TaskExecution` for more details
type TaskExecutionList struct {
	TaskExecutions []*TaskExecution `protobuf:"bytes,1,rep,name=task_executions,json=taskExecutions,proto3" json:"task_executions,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query. If there are no more results, this value will be empty.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecutionList) Reset()         { *m = TaskExecutionList{} }
func (m *TaskExecutionList) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionList) ProtoMessage()    {}
func (*TaskExecutionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cde4c3aa101642e, []int{3}
}

func (m *TaskExecutionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionList.Unmarshal(m, b)
}
func (m *TaskExecutionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionList.Marshal(b, m, deterministic)
}
func (m *TaskExecutionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionList.Merge(m, src)
}
func (m *TaskExecutionList) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionList.Size(m)
}
func (m *TaskExecutionList) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionList.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionList proto.InternalMessageInfo

func (m *TaskExecutionList) GetTaskExecutions() []*TaskExecution {
	if m != nil {
		return m.TaskExecutions
	}
	return nil
}

func (m *TaskExecutionList) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Container for task execution details and results.
type TaskExecutionClosure struct {
	// Types that are valid to be assigned to OutputResult:
	//	*TaskExecutionClosure_OutputUri
	//	*TaskExecutionClosure_Error
	//	*TaskExecutionClosure_OutputData
	OutputResult isTaskExecutionClosure_OutputResult `protobuf_oneof:"output_result"`
	// The last recorded phase for this task execution.
	Phase core.TaskExecution_Phase `protobuf:"varint,3,opt,name=phase,proto3,enum=nebulaidl.core.TaskExecution_Phase" json:"phase,omitempty"`
	// Detailed log information output by the task execution.
	Logs []*core.TaskLog `protobuf:"bytes,4,rep,name=logs,proto3" json:"logs,omitempty"`
	// Time at which the task execution began running.
	StartedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// The amount of time the task execution spent running.
	Duration *duration.Duration `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
	// Time at which the task execution was created.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Time at which the task execution was last updated.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Custom data specific to the task plugin.
	CustomInfo *_struct.Struct `protobuf:"bytes,9,opt,name=custom_info,json=customInfo,proto3" json:"custom_info,omitempty"`
	// If there is an explanation for the most recent phase transition, the reason will capture it.
	Reason string `protobuf:"bytes,10,opt,name=reason,proto3" json:"reason,omitempty"`
	// A predefined yet extensible Task type identifier.
	TaskType string `protobuf:"bytes,11,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// Metadata around how a task was executed.
	Metadata *event.TaskExecutionMetadata `protobuf:"bytes,16,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The event version is used to indicate versioned changes in how data is maintained using this
	// proto message. For example, event_verison > 0 means that maps tasks logs use the
	// TaskExecutionMetadata ExternalResourceInfo fields for each subtask rather than the TaskLog
	// in this message.
	EventVersion int32 `protobuf:"varint,17,opt,name=event_version,json=eventVersion,proto3" json:"event_version,omitempty"`
	// A time-series of the phase transition or update explanations. This, when compared to storing a singular reason
	// as previously done, is much more valuable in visualizing and understanding historical evaluations.
	Reasons              []*Reason `protobuf:"bytes,18,rep,name=reasons,proto3" json:"reasons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TaskExecutionClosure) Reset()         { *m = TaskExecutionClosure{} }
func (m *TaskExecutionClosure) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionClosure) ProtoMessage()    {}
func (*TaskExecutionClosure) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cde4c3aa101642e, []int{4}
}

func (m *TaskExecutionClosure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionClosure.Unmarshal(m, b)
}
func (m *TaskExecutionClosure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionClosure.Marshal(b, m, deterministic)
}
func (m *TaskExecutionClosure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionClosure.Merge(m, src)
}
func (m *TaskExecutionClosure) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionClosure.Size(m)
}
func (m *TaskExecutionClosure) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionClosure.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionClosure proto.InternalMessageInfo

type isTaskExecutionClosure_OutputResult interface {
	isTaskExecutionClosure_OutputResult()
}

type TaskExecutionClosure_OutputUri struct {
	OutputUri string `protobuf:"bytes,1,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type TaskExecutionClosure_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type TaskExecutionClosure_OutputData struct {
	OutputData *core.LiteralMap `protobuf:"bytes,12,opt,name=output_data,json=outputData,proto3,oneof"`
}

func (*TaskExecutionClosure_OutputUri) isTaskExecutionClosure_OutputResult() {}

func (*TaskExecutionClosure_Error) isTaskExecutionClosure_OutputResult() {}

func (*TaskExecutionClosure_OutputData) isTaskExecutionClosure_OutputResult() {}

func (m *TaskExecutionClosure) GetOutputResult() isTaskExecutionClosure_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

// Deprecated: Do not use.
func (m *TaskExecutionClosure) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*TaskExecutionClosure_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *TaskExecutionClosure) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*TaskExecutionClosure_Error); ok {
		return x.Error
	}
	return nil
}

// Deprecated: Do not use.
func (m *TaskExecutionClosure) GetOutputData() *core.LiteralMap {
	if x, ok := m.GetOutputResult().(*TaskExecutionClosure_OutputData); ok {
		return x.OutputData
	}
	return nil
}

func (m *TaskExecutionClosure) GetPhase() core.TaskExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.TaskExecution_UNDEFINED
}

func (m *TaskExecutionClosure) GetLogs() []*core.TaskLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *TaskExecutionClosure) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *TaskExecutionClosure) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *TaskExecutionClosure) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *TaskExecutionClosure) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *TaskExecutionClosure) GetCustomInfo() *_struct.Struct {
	if m != nil {
		return m.CustomInfo
	}
	return nil
}

func (m *TaskExecutionClosure) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *TaskExecutionClosure) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *TaskExecutionClosure) GetMetadata() *event.TaskExecutionMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TaskExecutionClosure) GetEventVersion() int32 {
	if m != nil {
		return m.EventVersion
	}
	return 0
}

func (m *TaskExecutionClosure) GetReasons() []*Reason {
	if m != nil {
		return m.Reasons
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TaskExecutionClosure) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TaskExecutionClosure_OutputUri)(nil),
		(*TaskExecutionClosure_Error)(nil),
		(*TaskExecutionClosure_OutputData)(nil),
	}
}

// Reason is a single message annotated with a timestamp to indicate the instant the reason occurred.
type Reason struct {
	// occurred_at is the timestamp indicating the instant that this reason happened.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,1,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// message is the explanation for the most recent phase transition or status update.
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reason) Reset()         { *m = Reason{} }
func (m *Reason) String() string { return proto.CompactTextString(m) }
func (*Reason) ProtoMessage()    {}
func (*Reason) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cde4c3aa101642e, []int{5}
}

func (m *Reason) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reason.Unmarshal(m, b)
}
func (m *Reason) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reason.Marshal(b, m, deterministic)
}
func (m *Reason) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reason.Merge(m, src)
}
func (m *Reason) XXX_Size() int {
	return xxx_messageInfo_Reason.Size(m)
}
func (m *Reason) XXX_DiscardUnknown() {
	xxx_messageInfo_Reason.DiscardUnknown(m)
}

var xxx_messageInfo_Reason proto.InternalMessageInfo

func (m *Reason) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *Reason) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Request structure to fetch inputs and output for a task execution.
// By default this data is not returned inline in :ref:`ref_nebulaidl.admin.TaskExecutionGetRequest`
type TaskExecutionGetDataRequest struct {
	// The identifier of the task execution for which to fetch inputs and outputs.
	// +required
	Id                   *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *TaskExecutionGetDataRequest) Reset()         { *m = TaskExecutionGetDataRequest{} }
func (m *TaskExecutionGetDataRequest) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionGetDataRequest) ProtoMessage()    {}
func (*TaskExecutionGetDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cde4c3aa101642e, []int{6}
}

func (m *TaskExecutionGetDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionGetDataRequest.Unmarshal(m, b)
}
func (m *TaskExecutionGetDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionGetDataRequest.Marshal(b, m, deterministic)
}
func (m *TaskExecutionGetDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionGetDataRequest.Merge(m, src)
}
func (m *TaskExecutionGetDataRequest) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionGetDataRequest.Size(m)
}
func (m *TaskExecutionGetDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionGetDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionGetDataRequest proto.InternalMessageInfo

func (m *TaskExecutionGetDataRequest) GetId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

// Response structure for TaskExecutionGetDataRequest which contains inputs and outputs for a task execution.
type TaskExecutionGetDataResponse struct {
	// Signed url to fetch a core.LiteralMap of task execution inputs.
	// Deprecated: Please use full_inputs instead.
	Inputs *UrlBlob `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"` // Deprecated: Do not use.
	// Signed url to fetch a core.LiteralMap of task execution outputs.
	// Deprecated: Please use full_outputs instead.
	Outputs *UrlBlob `protobuf:"bytes,2,opt,name=outputs,proto3" json:"outputs,omitempty"` // Deprecated: Do not use.
	// Full_inputs will only be populated if they are under a configured size threshold.
	FullInputs *core.LiteralMap `protobuf:"bytes,3,opt,name=full_inputs,json=fullInputs,proto3" json:"full_inputs,omitempty"`
	// Full_outputs will only be populated if they are under a configured size threshold.
	FullOutputs *core.LiteralMap `protobuf:"bytes,4,opt,name=full_outputs,json=fullOutputs,proto3" json:"full_outputs,omitempty"`
	// nebula tiny url to fetch a core.LiteralMap of task execution's IO
	// Deck will be empty for task
	NebulaUrls            *NebulaURLs `protobuf:"bytes,5,opt,name=nebula_urls,json=nebulaUrls,proto3" json:"nebula_urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TaskExecutionGetDataResponse) Reset()         { *m = TaskExecutionGetDataResponse{} }
func (m *TaskExecutionGetDataResponse) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionGetDataResponse) ProtoMessage()    {}
func (*TaskExecutionGetDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cde4c3aa101642e, []int{7}
}

func (m *TaskExecutionGetDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionGetDataResponse.Unmarshal(m, b)
}
func (m *TaskExecutionGetDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionGetDataResponse.Marshal(b, m, deterministic)
}
func (m *TaskExecutionGetDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionGetDataResponse.Merge(m, src)
}
func (m *TaskExecutionGetDataResponse) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionGetDataResponse.Size(m)
}
func (m *TaskExecutionGetDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionGetDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionGetDataResponse proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *TaskExecutionGetDataResponse) GetInputs() *UrlBlob {
	if m != nil {
		return m.Inputs
	}
	return nil
}

// Deprecated: Do not use.
func (m *TaskExecutionGetDataResponse) GetOutputs() *UrlBlob {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *TaskExecutionGetDataResponse) GetFullInputs() *core.LiteralMap {
	if m != nil {
		return m.FullInputs
	}
	return nil
}

func (m *TaskExecutionGetDataResponse) GetFullOutputs() *core.LiteralMap {
	if m != nil {
		return m.FullOutputs
	}
	return nil
}

func (m *TaskExecutionGetDataResponse) GetNebulaUrls() *NebulaURLs {
	if m != nil {
		return m.NebulaUrls
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskExecutionGetRequest)(nil), "nebulaidl.admin.TaskExecutionGetRequest")
	proto.RegisterType((*TaskExecutionListRequest)(nil), "nebulaidl.admin.TaskExecutionListRequest")
	proto.RegisterType((*TaskExecution)(nil), "nebulaidl.admin.TaskExecution")
	proto.RegisterType((*TaskExecutionList)(nil), "nebulaidl.admin.TaskExecutionList")
	proto.RegisterType((*TaskExecutionClosure)(nil), "nebulaidl.admin.TaskExecutionClosure")
	proto.RegisterType((*Reason)(nil), "nebulaidl.admin.Reason")
	proto.RegisterType((*TaskExecutionGetDataRequest)(nil), "nebulaidl.admin.TaskExecutionGetDataRequest")
	proto.RegisterType((*TaskExecutionGetDataResponse)(nil), "nebulaidl.admin.TaskExecutionGetDataResponse")
}

func init() {
	proto.RegisterFile("nebulaidl/admin/task_execution.proto", fileDescriptor_8cde4c3aa101642e)
}

var fileDescriptor_8cde4c3aa101642e = []byte{
	// 927 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xeb, 0x6e, 0x1b, 0x45,
	0x14, 0xee, 0x3a, 0x89, 0x2f, 0xc7, 0xb9, 0x90, 0x51, 0xd4, 0x6c, 0x93, 0x14, 0x2c, 0x07, 0x90,
	0x85, 0xd4, 0x35, 0x4a, 0x15, 0x14, 0x28, 0x20, 0x6c, 0xda, 0xd2, 0x48, 0x29, 0x94, 0x69, 0xcc,
	0x0f, 0xfe, 0xac, 0xc6, 0xbb, 0x63, 0x77, 0x94, 0xdd, 0x9d, 0xed, 0xcc, 0x6c, 0x85, 0xdf, 0x85,
	0x67, 0x41, 0xe2, 0x59, 0x78, 0x11, 0x34, 0x97, 0x75, 0xb2, 0xdb, 0x10, 0x4b, 0xf4, 0x8f, 0xe5,
	0x73, 0xce, 0xf7, 0x9d, 0xfb, 0x19, 0x2d, 0x1c, 0xcf, 0x92, 0x85, 0xa2, 0x2c, 0x4e, 0x86, 0x24,
	0x4e, 0x59, 0x36, 0x54, 0x44, 0x5e, 0x85, 0xf4, 0x0f, 0x1a, 0x15, 0x8a, 0xf1, 0x2c, 0xc8, 0x05,
	0x57, 0x1c, 0x6d, 0x97, 0xa0, 0xc0, 0x80, 0x0e, 0x0e, 0x6b, 0xa4, 0x88, 0xa7, 0x69, 0x09, 0x3e,
	0x78, 0xb8, 0x34, 0x46, 0x5c, 0xd0, 0x61, 0xcd, 0xd7, 0xc1, 0xc7, 0x55, 0x33, 0x8b, 0x69, 0xa6,
	0xd8, 0x8c, 0x51, 0xe1, 0xec, 0x47, 0x55, 0x7b, 0xc2, 0x14, 0x15, 0x24, 0x91, 0xce, 0x7a, 0xb0,
	0xb4, 0xd2, 0x77, 0x34, 0x53, 0xf6, 0xd7, 0xd9, 0x3e, 0x99, 0x73, 0x3e, 0x4f, 0xe8, 0xd0, 0x48,
	0xd3, 0x62, 0x36, 0x54, 0x2c, 0xa5, 0x52, 0x91, 0x34, 0x2f, 0x43, 0xd7, 0x01, 0x71, 0x21, 0xc8,
	0x8d, 0xd4, 0x8e, 0xea, 0x76, 0xa9, 0x44, 0x11, 0x39, 0xf7, 0xfd, 0x5f, 0x61, 0xff, 0x92, 0xc8,
	0xab, 0x67, 0x65, 0x3d, 0x3f, 0x51, 0x85, 0xe9, 0xdb, 0x82, 0x4a, 0x85, 0xbe, 0x82, 0x06, 0x8b,
	0x7d, 0xaf, 0xe7, 0x0d, 0xba, 0x27, 0x9f, 0x07, 0xcb, 0x66, 0xe9, 0x02, 0x82, 0x0a, 0xe7, 0x7c,
	0x59, 0x2d, 0x6e, 0xb0, 0xb8, 0xff, 0x8f, 0x07, 0x7e, 0xc5, 0x7e, 0xc1, 0xe4, 0xd2, 0x29, 0x86,
	0xdd, 0x8c, 0xc7, 0xf4, 0x7a, 0x18, 0xe1, 0x7f, 0xc6, 0xf8, 0x99, 0xc7, 0xf4, 0xb6, 0x18, 0x3b,
	0x59, 0xd5, 0x80, 0xf6, 0x60, 0x23, 0x61, 0x29, 0x53, 0x7e, 0xa3, 0xe7, 0x0d, 0xb6, 0xb0, 0x15,
	0xb4, 0x56, 0xf1, 0x2b, 0x9a, 0xf9, 0x6b, 0x3d, 0x6f, 0xd0, 0xc1, 0x56, 0x40, 0x3e, 0xb4, 0x66,
	0x2c, 0x51, 0x54, 0x48, 0x7f, 0xdd, 0xe8, 0x4b, 0x11, 0x3d, 0x82, 0x96, 0xe4, 0x42, 0x85, 0xd3,
	0x85, 0xbf, 0x61, 0xf2, 0xd9, 0x0b, 0xaa, 0x0b, 0x12, 0xbc, 0xe6, 0x42, 0xe1, 0xa6, 0x06, 0x8d,
	0x17, 0xfd, 0xbf, 0x3d, 0xd8, 0xaa, 0x54, 0xf9, 0x7f, 0xfb, 0x85, 0x0e, 0xa1, 0xc3, 0xb2, 0xbc,
	0x50, 0x61, 0x21, 0x98, 0x29, 0xa1, 0x83, 0xdb, 0x46, 0x31, 0x11, 0x0c, 0x7d, 0x0f, 0xad, 0x28,
	0xe1, 0xb2, 0x10, 0xd4, 0xd4, 0xd1, 0x3d, 0xf9, 0xb4, 0x9e, 0x55, 0xc5, 0xf5, 0x8f, 0x16, 0x8b,
	0x4b, 0x92, 0x71, 0x2e, 0xc3, 0x9c, 0x08, 0x9a, 0x29, 0x53, 0x71, 0x1b, 0xb7, 0x99, 0x7c, 0x65,
	0xe4, 0xfe, 0x5b, 0xd8, 0x7d, 0x6f, 0x50, 0xe8, 0x39, 0xec, 0x54, 0xcf, 0x45, 0xfa, 0x5e, 0x6f,
	0x6d, 0xd0, 0x3d, 0x79, 0x78, 0x67, 0x64, 0xbc, 0xad, 0x6e, 0x8a, 0xf2, 0xba, 0xff, 0x8d, 0x1b,
	0xfd, 0xef, 0xff, 0xd9, 0x84, 0xbd, 0xdb, 0x32, 0x46, 0xc7, 0x00, 0xbc, 0x50, 0x65, 0x1b, 0x74,
	0x17, 0x3b, 0xe3, 0x86, 0xef, 0xbd, 0xb8, 0x87, 0x3b, 0x56, 0xaf, 0xbb, 0x71, 0x0a, 0x1b, 0x54,
	0x08, 0x2e, 0x8c, 0xcf, 0x4a, 0x46, 0xa6, 0xcb, 0x4b, 0xa7, 0xcf, 0x34, 0xe8, 0xc5, 0x3d, 0x6c,
	0xd1, 0xe8, 0x07, 0xe8, 0x3a, 0xdf, 0x31, 0x51, 0xc4, 0xdf, 0x34, 0xe4, 0x07, 0x35, 0xf2, 0x85,
	0xbd, 0xc9, 0x97, 0x24, 0x77, 0x71, 0x5d, 0x3e, 0x4f, 0x89, 0x22, 0xe8, 0x0c, 0x36, 0xf2, 0x37,
	0x44, 0xda, 0x21, 0x6c, 0x9f, 0xf4, 0xef, 0x1a, 0x6f, 0xf0, 0x4a, 0x23, 0xb1, 0x25, 0xa0, 0x2f,
	0x60, 0x3d, 0xe1, 0x73, 0xbd, 0x6d, 0xba, 0x87, 0xf7, 0x6f, 0x21, 0x5e, 0xf0, 0x39, 0x36, 0x18,
	0xf4, 0x35, 0x80, 0x54, 0x44, 0x28, 0x1a, 0x87, 0x44, 0xb9, 0x2d, 0x3c, 0x08, 0xec, 0xfd, 0x06,
	0xe5, 0xfd, 0x06, 0x97, 0xe5, 0x03, 0x80, 0x3b, 0x0e, 0x3d, 0x52, 0xe8, 0x14, 0xda, 0xe5, 0xdd,
	0xfb, 0x4d, 0x57, 0x5f, 0x9d, 0xf8, 0xd4, 0x01, 0xf0, 0x12, 0xaa, 0x23, 0x46, 0x82, 0x12, 0x17,
	0xb1, 0xb5, 0x3a, 0xa2, 0x43, 0x8f, 0x94, 0xa6, 0x16, 0x79, 0x5c, 0x52, 0xdb, 0xab, 0xa9, 0x0e,
	0x3d, 0x52, 0xe8, 0x0c, 0xba, 0x51, 0x21, 0x15, 0x4f, 0x43, 0x96, 0xcd, 0xb8, 0xdf, 0x31, 0xdc,
	0xfd, 0xf7, 0xb8, 0xaf, 0xcd, 0x43, 0x85, 0xc1, 0x62, 0xcf, 0xb3, 0x19, 0x47, 0xf7, 0xa1, 0x29,
	0x28, 0x91, 0x3c, 0xf3, 0xc1, 0x6c, 0x95, 0x93, 0xf4, 0x9a, 0x9b, 0xa5, 0x55, 0x8b, 0x9c, 0xfa,
	0x5d, 0x7b, 0x43, 0x5a, 0x71, 0xb9, 0xc8, 0x29, 0x1a, 0x41, 0x3b, 0xa5, 0x8a, 0x98, 0xd9, 0x7f,
	0x64, 0x62, 0x7d, 0x76, 0x3d, 0x06, 0xfb, 0xd6, 0x56, 0x06, 0xf8, 0xd2, 0x81, 0xf1, 0x92, 0x86,
	0x8e, 0x61, 0xcb, 0x00, 0xc3, 0x77, 0x54, 0x48, 0xdd, 0xe3, 0xdd, 0x9e, 0x37, 0xd8, 0xc0, 0x9b,
	0x46, 0xf9, 0x9b, 0xd5, 0xa1, 0x2f, 0xa1, 0x65, 0xd3, 0x91, 0x3e, 0xaa, 0x4f, 0xdb, 0x5e, 0x0c,
	0x36, 0x66, 0x5c, 0xc2, 0xc6, 0x3b, 0xb0, 0xe5, 0x16, 0x53, 0x50, 0x59, 0x24, 0xaa, 0x1f, 0x42,
	0xd3, 0x62, 0xd0, 0x13, 0xe8, 0xf2, 0x28, 0x2a, 0x84, 0xb0, 0xfd, 0xf5, 0x56, 0xf6, 0x17, 0x4a,
	0xf8, 0x48, 0xe9, 0x57, 0x2e, 0xa5, 0x52, 0x92, 0x39, 0x75, 0xd7, 0x57, 0x8a, 0xfd, 0x09, 0x1c,
	0xd6, 0xdf, 0x7b, 0xbd, 0xe0, 0x1f, 0xfa, 0xe6, 0xff, 0xd5, 0x80, 0xa3, 0xdb, 0xfd, 0xca, 0x9c,
	0x67, 0x92, 0xa2, 0xc7, 0xd0, 0x34, 0x6f, 0x9a, 0x74, 0xce, 0xf7, 0xeb, 0xad, 0x99, 0x88, 0x64,
	0x9c, 0xf0, 0xa9, 0xbe, 0x3d, 0xec, 0xa0, 0xe8, 0x14, 0x5a, 0xb6, 0x3d, 0xd2, 0x1d, 0xfc, 0x9d,
	0xac, 0x12, 0x8b, 0xbe, 0x81, 0xee, 0xac, 0x48, 0x92, 0xd0, 0x05, 0x5c, 0x5b, 0x71, 0xee, 0x18,
	0x34, 0xfa, 0xdc, 0x86, 0xfc, 0x16, 0x36, 0x0d, 0xb7, 0x8c, 0xbb, 0xbe, 0x8a, 0x6c, 0x42, 0xfd,
	0xe2, 0x22, 0x9f, 0x01, 0x18, 0x60, 0x58, 0x88, 0x44, 0xba, 0x03, 0x7e, 0x50, 0xcf, 0xf9, 0xb9,
	0x16, 0x27, 0xf8, 0x42, 0xe2, 0x8e, 0xb1, 0x4c, 0x44, 0x22, 0xc7, 0xdf, 0xfd, 0xfe, 0x64, 0xce,
	0xd4, 0x9b, 0x62, 0x1a, 0x44, 0x3c, 0x1d, 0x1a, 0x3d, 0x17, 0x73, 0xfb, 0x67, 0xb8, 0xfc, 0x3c,
	0x98, 0xd3, 0x6c, 0x98, 0x4f, 0x1f, 0xcd, 0xf9, 0xb0, 0xfa, 0xad, 0x32, 0x6d, 0x9a, 0x85, 0x78,
	0xfc, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xcb, 0x61, 0x27, 0xf9, 0x08, 0x00, 0x00,
}