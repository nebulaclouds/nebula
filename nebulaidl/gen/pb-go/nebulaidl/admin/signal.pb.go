// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/admin/signal.proto

package admin

import (
	fmt "fmt"
	core "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/core"
	proto "github.com/golang/protobuf/proto"
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

// SignalGetOrCreateRequest represents a request structure to retrieve or create a signal.
// See :ref:`ref_nebulaidl.admin.Signal` for more details
type SignalGetOrCreateRequest struct {
	// A unique identifier for the requested signal.
	Id *core.SignalIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A type denoting the required value type for this signal.
	Type                 *core.LiteralType `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SignalGetOrCreateRequest) Reset()         { *m = SignalGetOrCreateRequest{} }
func (m *SignalGetOrCreateRequest) String() string { return proto.CompactTextString(m) }
func (*SignalGetOrCreateRequest) ProtoMessage()    {}
func (*SignalGetOrCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b1c2d4aa3794afa, []int{0}
}

func (m *SignalGetOrCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalGetOrCreateRequest.Unmarshal(m, b)
}
func (m *SignalGetOrCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalGetOrCreateRequest.Marshal(b, m, deterministic)
}
func (m *SignalGetOrCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalGetOrCreateRequest.Merge(m, src)
}
func (m *SignalGetOrCreateRequest) XXX_Size() int {
	return xxx_messageInfo_SignalGetOrCreateRequest.Size(m)
}
func (m *SignalGetOrCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalGetOrCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignalGetOrCreateRequest proto.InternalMessageInfo

func (m *SignalGetOrCreateRequest) GetId() *core.SignalIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SignalGetOrCreateRequest) GetType() *core.LiteralType {
	if m != nil {
		return m.Type
	}
	return nil
}

// SignalListRequest represents a request structure to retrieve a collection of signals.
// See :ref:`ref_nebulaidl.admin.Signal` for more details
type SignalListRequest struct {
	// Indicates the workflow execution to filter by.
	// +required
	WorkflowExecutionId *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=workflow_execution_id,json=workflowExecutionId,proto3" json:"workflow_execution_id,omitempty"`
	// Indicates the number of resources to be returned.
	// +required
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// In the case of multiple pages of results, the, server-provided token can be used to fetch the next page
	// in a query.
	// +optional
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	// Indicates a list of filters passed as string.
	// +optional
	Filters string `protobuf:"bytes,4,opt,name=filters,proto3" json:"filters,omitempty"`
	// Sort ordering.
	// +optional
	SortBy               *Sort    `protobuf:"bytes,5,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignalListRequest) Reset()         { *m = SignalListRequest{} }
func (m *SignalListRequest) String() string { return proto.CompactTextString(m) }
func (*SignalListRequest) ProtoMessage()    {}
func (*SignalListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b1c2d4aa3794afa, []int{1}
}

func (m *SignalListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalListRequest.Unmarshal(m, b)
}
func (m *SignalListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalListRequest.Marshal(b, m, deterministic)
}
func (m *SignalListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalListRequest.Merge(m, src)
}
func (m *SignalListRequest) XXX_Size() int {
	return xxx_messageInfo_SignalListRequest.Size(m)
}
func (m *SignalListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignalListRequest proto.InternalMessageInfo

func (m *SignalListRequest) GetWorkflowExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.WorkflowExecutionId
	}
	return nil
}

func (m *SignalListRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SignalListRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SignalListRequest) GetFilters() string {
	if m != nil {
		return m.Filters
	}
	return ""
}

func (m *SignalListRequest) GetSortBy() *Sort {
	if m != nil {
		return m.SortBy
	}
	return nil
}

// SignalList represents collection of signals along with the token of the last result.
// See :ref:`ref_nebulaidl.admin.Signal` for more details
type SignalList struct {
	// A list of signals matching the input filters.
	Signals []*Signal `protobuf:"bytes,1,rep,name=signals,proto3" json:"signals,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query. If there are no more results, this value will be empty.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignalList) Reset()         { *m = SignalList{} }
func (m *SignalList) String() string { return proto.CompactTextString(m) }
func (*SignalList) ProtoMessage()    {}
func (*SignalList) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b1c2d4aa3794afa, []int{2}
}

func (m *SignalList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalList.Unmarshal(m, b)
}
func (m *SignalList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalList.Marshal(b, m, deterministic)
}
func (m *SignalList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalList.Merge(m, src)
}
func (m *SignalList) XXX_Size() int {
	return xxx_messageInfo_SignalList.Size(m)
}
func (m *SignalList) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalList.DiscardUnknown(m)
}

var xxx_messageInfo_SignalList proto.InternalMessageInfo

func (m *SignalList) GetSignals() []*Signal {
	if m != nil {
		return m.Signals
	}
	return nil
}

func (m *SignalList) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// SignalSetRequest represents a request structure to set the value on a signal. Setting a signal
// effetively satisfies the signal condition within a Nebula workflow.
// See :ref:`ref_nebulaidl.admin.Signal` for more details
type SignalSetRequest struct {
	// A unique identifier for the requested signal.
	Id *core.SignalIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The value of this signal, must match the defining signal type.
	Value                *core.Literal `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SignalSetRequest) Reset()         { *m = SignalSetRequest{} }
func (m *SignalSetRequest) String() string { return proto.CompactTextString(m) }
func (*SignalSetRequest) ProtoMessage()    {}
func (*SignalSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b1c2d4aa3794afa, []int{3}
}

func (m *SignalSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalSetRequest.Unmarshal(m, b)
}
func (m *SignalSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalSetRequest.Marshal(b, m, deterministic)
}
func (m *SignalSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalSetRequest.Merge(m, src)
}
func (m *SignalSetRequest) XXX_Size() int {
	return xxx_messageInfo_SignalSetRequest.Size(m)
}
func (m *SignalSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignalSetRequest proto.InternalMessageInfo

func (m *SignalSetRequest) GetId() *core.SignalIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SignalSetRequest) GetValue() *core.Literal {
	if m != nil {
		return m.Value
	}
	return nil
}

// SignalSetResponse represents a response structure if signal setting succeeds.
type SignalSetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignalSetResponse) Reset()         { *m = SignalSetResponse{} }
func (m *SignalSetResponse) String() string { return proto.CompactTextString(m) }
func (*SignalSetResponse) ProtoMessage()    {}
func (*SignalSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b1c2d4aa3794afa, []int{4}
}

func (m *SignalSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalSetResponse.Unmarshal(m, b)
}
func (m *SignalSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalSetResponse.Marshal(b, m, deterministic)
}
func (m *SignalSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalSetResponse.Merge(m, src)
}
func (m *SignalSetResponse) XXX_Size() int {
	return xxx_messageInfo_SignalSetResponse.Size(m)
}
func (m *SignalSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignalSetResponse proto.InternalMessageInfo

// Signal encapsulates a unique identifier, associated metadata, and a value for a single Nebula
// signal. Signals may exist either without a set value (representing a signal request) or with a
// populated value (indicating the signal has been given).
type Signal struct {
	// A unique identifier for the requested signal.
	Id *core.SignalIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A type denoting the required value type for this signal.
	Type *core.LiteralType `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// The value of the signal. This is only available if the signal has been "set" and must match
	// the defined the type.
	Value                *core.Literal `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Signal) Reset()         { *m = Signal{} }
func (m *Signal) String() string { return proto.CompactTextString(m) }
func (*Signal) ProtoMessage()    {}
func (*Signal) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b1c2d4aa3794afa, []int{5}
}

func (m *Signal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signal.Unmarshal(m, b)
}
func (m *Signal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signal.Marshal(b, m, deterministic)
}
func (m *Signal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signal.Merge(m, src)
}
func (m *Signal) XXX_Size() int {
	return xxx_messageInfo_Signal.Size(m)
}
func (m *Signal) XXX_DiscardUnknown() {
	xxx_messageInfo_Signal.DiscardUnknown(m)
}

var xxx_messageInfo_Signal proto.InternalMessageInfo

func (m *Signal) GetId() *core.SignalIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Signal) GetType() *core.LiteralType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Signal) GetValue() *core.Literal {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*SignalGetOrCreateRequest)(nil), "nebulaidl.admin.SignalGetOrCreateRequest")
	proto.RegisterType((*SignalListRequest)(nil), "nebulaidl.admin.SignalListRequest")
	proto.RegisterType((*SignalList)(nil), "nebulaidl.admin.SignalList")
	proto.RegisterType((*SignalSetRequest)(nil), "nebulaidl.admin.SignalSetRequest")
	proto.RegisterType((*SignalSetResponse)(nil), "nebulaidl.admin.SignalSetResponse")
	proto.RegisterType((*Signal)(nil), "nebulaidl.admin.Signal")
}

func init() { proto.RegisterFile("nebulaidl/admin/signal.proto", fileDescriptor_4b1c2d4aa3794afa) }

var fileDescriptor_4b1c2d4aa3794afa = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xd1, 0x8a, 0x13, 0x31,
	0x14, 0x86, 0x99, 0x76, 0xdb, 0x62, 0x16, 0x45, 0x67, 0xd7, 0x25, 0x56, 0xd1, 0x32, 0x57, 0x45,
	0xdc, 0x44, 0xd6, 0x4b, 0xf1, 0x66, 0x45, 0x44, 0x58, 0x10, 0xd2, 0x05, 0xc1, 0x0b, 0xcb, 0xb4,
	0x73, 0x3a, 0x86, 0xcd, 0x24, 0xb3, 0xc9, 0xa9, 0x75, 0xf0, 0x41, 0x7c, 0x45, 0x1f, 0x43, 0x9a,
	0xcc, 0x4c, 0x77, 0xba, 0x5e, 0x88, 0xe0, 0xdd, 0xe4, 0xfc, 0xff, 0x39, 0xf9, 0x86, 0x3f, 0x87,
	0x3c, 0x5e, 0xa9, 0x0a, 0x41, 0x66, 0x8a, 0xa7, 0x59, 0x21, 0x35, 0x77, 0x32, 0xd7, 0xa9, 0x62,
	0xa5, 0x35, 0x68, 0xe2, 0x7b, 0x8d, 0xc8, 0xbc, 0x38, 0xde, 0x37, 0x2f, 0x4d, 0x51, 0x18, 0x1d,
	0xcc, 0xe3, 0xa7, 0xad, 0xb8, 0x34, 0x16, 0xb8, 0xcc, 0x40, 0xa3, 0x5c, 0x49, 0xb0, 0xb5, 0xfe,
	0xa4, 0xab, 0x2b, 0x89, 0x60, 0x53, 0xe5, 0x6a, 0xf5, 0x51, 0x57, 0xc5, 0xaa, 0x84, 0x5a, 0x4a,
	0x7e, 0x10, 0x3a, 0xf3, 0x54, 0xef, 0x01, 0x3f, 0xda, 0xb7, 0x16, 0x52, 0x04, 0x01, 0xd7, 0x6b,
	0x70, 0x18, 0x73, 0xd2, 0x93, 0x19, 0x8d, 0x26, 0xd1, 0xf4, 0xf0, 0xec, 0x19, 0x6b, 0x71, 0xb7,
	0x33, 0x58, 0x68, 0xfa, 0xd0, 0x72, 0x88, 0x9e, 0xcc, 0x62, 0x46, 0x0e, 0xb6, 0xb3, 0x69, 0xcf,
	0xb7, 0x8c, 0xf7, 0x5a, 0x2e, 0x02, 0xd4, 0x65, 0x55, 0x82, 0xf0, 0xbe, 0xe4, 0x57, 0x44, 0x1e,
	0x84, 0x41, 0x17, 0xd2, 0x61, 0x73, 0xed, 0x17, 0xf2, 0x70, 0x63, 0xec, 0xd5, 0x4a, 0x99, 0xcd,
	0x1c, 0xbe, 0xc3, 0x72, 0x8d, 0xd2, 0xe8, 0x79, 0x4b, 0xf2, 0x7c, 0x6f, 0xec, 0xa7, 0xda, 0xfb,
	0xae, 0xb1, 0xde, 0x80, 0x3a, 0xda, 0xdc, 0x16, 0xe3, 0x63, 0x32, 0x50, 0xb2, 0x90, 0xe8, 0x31,
	0xef, 0x8a, 0x70, 0xd8, 0x56, 0xd1, 0x5c, 0x81, 0xa6, 0xfd, 0x49, 0x34, 0xbd, 0x23, 0xc2, 0x21,
	0xa6, 0x64, 0xb4, 0x92, 0x0a, 0xc1, 0x3a, 0x7a, 0xe0, 0xeb, 0xcd, 0x31, 0x3e, 0x25, 0x23, 0x67,
	0x2c, 0xce, 0x17, 0x15, 0x1d, 0x78, 0xae, 0x63, 0xd6, 0x0d, 0x94, 0xcd, 0x8c, 0x45, 0x31, 0xdc,
	0x9a, 0xce, 0xab, 0xe4, 0x92, 0x90, 0xdd, 0x9f, 0xc6, 0x2f, 0xc9, 0x28, 0xbc, 0x05, 0x47, 0xa3,
	0x49, 0x7f, 0x7a, 0x78, 0x76, 0x72, 0xab, 0xd9, 0xcb, 0xa2, 0xb1, 0xed, 0xf0, 0x7a, 0x37, 0xf0,
	0x92, 0x6b, 0x72, 0x3f, 0x18, 0x67, 0x80, 0xff, 0x9c, 0xda, 0x0b, 0x32, 0xf8, 0x96, 0xaa, 0x75,
	0x13, 0xdb, 0xc9, 0x9f, 0x63, 0x13, 0xc1, 0x94, 0x1c, 0x35, 0x91, 0xf9, 0x2b, 0x5d, 0x69, 0xb4,
	0x83, 0xe4, 0x67, 0x44, 0x86, 0xa1, 0xfa, 0xdf, 0x1f, 0xcd, 0x0e, 0xb7, 0xff, 0x17, 0xb8, 0xe7,
	0x6f, 0x3e, 0xbf, 0xce, 0x25, 0x7e, 0x5d, 0x2f, 0xd8, 0xd2, 0x14, 0xdc, 0x5b, 0x8d, 0xcd, 0xc3,
	0x07, 0x6f, 0xd7, 0x22, 0x07, 0xcd, 0xcb, 0xc5, 0x69, 0x6e, 0x78, 0x77, 0x09, 0x17, 0x43, 0xbf,
	0x25, 0xaf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x5e, 0xa5, 0xbf, 0xca, 0x03, 0x00, 0x00,
}
