// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nebulaidl/plugins/dask.proto

package plugins

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

// Custom Proto for Dask Plugin.
type DaskJob struct {
	// Spec for the scheduler pod.
	Scheduler *DaskScheduler `protobuf:"bytes,1,opt,name=scheduler,proto3" json:"scheduler,omitempty"`
	// Spec of the default worker group.
	Workers              *DaskWorkerGroup `protobuf:"bytes,2,opt,name=workers,proto3" json:"workers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DaskJob) Reset()         { *m = DaskJob{} }
func (m *DaskJob) String() string { return proto.CompactTextString(m) }
func (*DaskJob) ProtoMessage()    {}
func (*DaskJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_d719e18eb4f4b89f, []int{0}
}

func (m *DaskJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DaskJob.Unmarshal(m, b)
}
func (m *DaskJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DaskJob.Marshal(b, m, deterministic)
}
func (m *DaskJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DaskJob.Merge(m, src)
}
func (m *DaskJob) XXX_Size() int {
	return xxx_messageInfo_DaskJob.Size(m)
}
func (m *DaskJob) XXX_DiscardUnknown() {
	xxx_messageInfo_DaskJob.DiscardUnknown(m)
}

var xxx_messageInfo_DaskJob proto.InternalMessageInfo

func (m *DaskJob) GetScheduler() *DaskScheduler {
	if m != nil {
		return m.Scheduler
	}
	return nil
}

func (m *DaskJob) GetWorkers() *DaskWorkerGroup {
	if m != nil {
		return m.Workers
	}
	return nil
}

// Specification for the scheduler pod.
type DaskScheduler struct {
	// Optional image to use. If unset, will use the default image.
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// Resources assigned to the scheduler pod.
	Resources            *core.Resources `protobuf:"bytes,2,opt,name=resources,proto3" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DaskScheduler) Reset()         { *m = DaskScheduler{} }
func (m *DaskScheduler) String() string { return proto.CompactTextString(m) }
func (*DaskScheduler) ProtoMessage()    {}
func (*DaskScheduler) Descriptor() ([]byte, []int) {
	return fileDescriptor_d719e18eb4f4b89f, []int{1}
}

func (m *DaskScheduler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DaskScheduler.Unmarshal(m, b)
}
func (m *DaskScheduler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DaskScheduler.Marshal(b, m, deterministic)
}
func (m *DaskScheduler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DaskScheduler.Merge(m, src)
}
func (m *DaskScheduler) XXX_Size() int {
	return xxx_messageInfo_DaskScheduler.Size(m)
}
func (m *DaskScheduler) XXX_DiscardUnknown() {
	xxx_messageInfo_DaskScheduler.DiscardUnknown(m)
}

var xxx_messageInfo_DaskScheduler proto.InternalMessageInfo

func (m *DaskScheduler) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *DaskScheduler) GetResources() *core.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

type DaskWorkerGroup struct {
	// Number of workers in the group.
	NumberOfWorkers uint32 `protobuf:"varint,1,opt,name=number_of_workers,json=numberOfWorkers,proto3" json:"number_of_workers,omitempty"`
	// Optional image to use for the pods of the worker group. If unset, will use the default image.
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	// Resources assigned to the all pods of the worker group.
	// As per https://kubernetes.dask.org/en/latest/kubecluster.html?highlight=limit#best-practices
	// it is advised to only set limits. If requests are not explicitly set, the plugin will make
	// sure to set requests==limits.
	// The plugin sets ` --memory-limit` as well as `--nthreads` for the workers according to the limit.
	Resources            *core.Resources `protobuf:"bytes,3,opt,name=resources,proto3" json:"resources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DaskWorkerGroup) Reset()         { *m = DaskWorkerGroup{} }
func (m *DaskWorkerGroup) String() string { return proto.CompactTextString(m) }
func (*DaskWorkerGroup) ProtoMessage()    {}
func (*DaskWorkerGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_d719e18eb4f4b89f, []int{2}
}

func (m *DaskWorkerGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DaskWorkerGroup.Unmarshal(m, b)
}
func (m *DaskWorkerGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DaskWorkerGroup.Marshal(b, m, deterministic)
}
func (m *DaskWorkerGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DaskWorkerGroup.Merge(m, src)
}
func (m *DaskWorkerGroup) XXX_Size() int {
	return xxx_messageInfo_DaskWorkerGroup.Size(m)
}
func (m *DaskWorkerGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_DaskWorkerGroup.DiscardUnknown(m)
}

var xxx_messageInfo_DaskWorkerGroup proto.InternalMessageInfo

func (m *DaskWorkerGroup) GetNumberOfWorkers() uint32 {
	if m != nil {
		return m.NumberOfWorkers
	}
	return 0
}

func (m *DaskWorkerGroup) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *DaskWorkerGroup) GetResources() *core.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto.RegisterType((*DaskJob)(nil), "nebulaidl.plugins.DaskJob")
	proto.RegisterType((*DaskScheduler)(nil), "nebulaidl.plugins.DaskScheduler")
	proto.RegisterType((*DaskWorkerGroup)(nil), "nebulaidl.plugins.DaskWorkerGroup")
}

func init() { proto.RegisterFile("nebulaidl/plugins/dask.proto", fileDescriptor_d719e18eb4f4b89f) }

var fileDescriptor_d719e18eb4f4b89f = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xc4, 0x30,
	0x14, 0xc4, 0xe9, 0x8a, 0x2e, 0x8d, 0x2c, 0xab, 0xc1, 0x43, 0xd5, 0x83, 0xda, 0x93, 0x08, 0x26,
	0xa0, 0xe0, 0x45, 0x16, 0x41, 0x04, 0xc1, 0x8b, 0x10, 0x0f, 0x82, 0x20, 0x4b, 0xff, 0xa4, 0xd9,
	0xd2, 0x3f, 0xaf, 0xbc, 0x34, 0x88, 0x1f, 0xc0, 0x93, 0x5f, 0x5a, 0xda, 0x98, 0x2d, 0x16, 0x05,
	0x6f, 0x2d, 0xf3, 0x9b, 0x37, 0x13, 0x86, 0x1c, 0x66, 0xe5, 0x7b, 0x2b, 0xf3, 0xb4, 0xe4, 0x4d,
	0x69, 0x54, 0x5e, 0x6b, 0x9e, 0x46, 0xba, 0x60, 0x0d, 0x42, 0x0b, 0x74, 0xc7, 0x89, 0xec, 0x5b,
	0x3c, 0xd8, 0x5f, 0xe3, 0x09, 0xa0, 0xe4, 0x6d, 0xa4, 0x0b, 0x6d, 0xe1, 0xf0, 0xc3, 0x23, 0xd3,
	0xbb, 0x48, 0x17, 0x0f, 0x10, 0xd3, 0x05, 0xf1, 0x75, 0xb2, 0x92, 0xa9, 0x29, 0x25, 0x06, 0xde,
	0xb1, 0x77, 0xba, 0x7d, 0x71, 0xc4, 0xc6, 0xc7, 0x58, 0x47, 0x3f, 0x39, 0x4c, 0x0c, 0x0e, 0x7a,
	0x4d, 0xa6, 0x6f, 0x80, 0x85, 0x44, 0x1d, 0x4c, 0x7a, 0xf3, 0xc9, 0xef, 0xe6, 0xe7, 0x1e, 0xba,
	0x47, 0x30, 0x8d, 0x70, 0x8e, 0xf0, 0x95, 0xcc, 0x7e, 0x1c, 0xa6, 0x7b, 0x64, 0x33, 0xaf, 0x22,
	0x25, 0xfb, 0x22, 0xbe, 0xb0, 0x3f, 0xf4, 0x8a, 0xf8, 0x28, 0x35, 0x18, 0x4c, 0xa4, 0x4b, 0x09,
	0x86, 0x94, 0xee, 0x75, 0x4c, 0x38, 0x5d, 0x0c, 0x68, 0xf8, 0xe9, 0x91, 0xf9, 0x28, 0x9b, 0x9e,
	0x91, 0xdd, 0xda, 0x54, 0xb1, 0xc4, 0x25, 0x64, 0x4b, 0xd7, 0xbc, 0x4b, 0x9b, 0x89, 0xb9, 0x15,
	0x1e, 0x33, 0xcb, 0xeb, 0xa1, 0xcd, 0xe4, 0xcf, 0x36, 0x1b, 0xff, 0x6e, 0x73, 0x7b, 0xf3, 0xb2,
	0x50, 0x79, 0xbb, 0x32, 0x31, 0x4b, 0xa0, 0xe2, 0xbd, 0x01, 0x50, 0xd9, 0x0f, 0xbe, 0xde, 0x4a,
	0xc9, 0x9a, 0x37, 0xf1, 0xb9, 0x02, 0x3e, 0x5e, 0x3b, 0xde, 0xea, 0xc7, 0xbb, 0xfc, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x6f, 0x43, 0xcc, 0x98, 0x08, 0x02, 0x00, 0x00,
}
