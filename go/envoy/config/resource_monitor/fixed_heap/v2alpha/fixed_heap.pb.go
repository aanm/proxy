// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/resource_monitor/fixed_heap/v2alpha/fixed_heap.proto

package v2alpha

import (
	fmt "fmt"
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

// The fixed heap resource monitor reports the Envoy process memory pressure, computed as a
// fraction of currently reserved heap memory divided by a statically configured maximum
// specified in the FixedHeapConfig.
type FixedHeapConfig struct {
	MaxHeapSizeBytes     uint64   `protobuf:"varint,1,opt,name=max_heap_size_bytes,json=maxHeapSizeBytes,proto3" json:"max_heap_size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FixedHeapConfig) Reset()         { *m = FixedHeapConfig{} }
func (m *FixedHeapConfig) String() string { return proto.CompactTextString(m) }
func (*FixedHeapConfig) ProtoMessage()    {}
func (*FixedHeapConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_141ee15b3c15e2df, []int{0}
}

func (m *FixedHeapConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FixedHeapConfig.Unmarshal(m, b)
}
func (m *FixedHeapConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FixedHeapConfig.Marshal(b, m, deterministic)
}
func (m *FixedHeapConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixedHeapConfig.Merge(m, src)
}
func (m *FixedHeapConfig) XXX_Size() int {
	return xxx_messageInfo_FixedHeapConfig.Size(m)
}
func (m *FixedHeapConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FixedHeapConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FixedHeapConfig proto.InternalMessageInfo

func (m *FixedHeapConfig) GetMaxHeapSizeBytes() uint64 {
	if m != nil {
		return m.MaxHeapSizeBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*FixedHeapConfig)(nil), "envoy.config.resource_monitor.fixed_heap.v2alpha.FixedHeapConfig")
}

func init() {
	proto.RegisterFile("envoy/config/resource_monitor/fixed_heap/v2alpha/fixed_heap.proto", fileDescriptor_141ee15b3c15e2df)
}

var fileDescriptor_141ee15b3c15e2df = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a,
	0x4e, 0x8d, 0xcf, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9, 0x2f, 0xd2, 0x4f, 0xcb, 0xac, 0x48, 0x4d, 0x89,
	0xcf, 0x48, 0x4d, 0x2c, 0xd0, 0x2f, 0x33, 0x4a, 0xcc, 0x29, 0xc8, 0x48, 0x44, 0x12, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0x00, 0x1b, 0xa1, 0x07, 0x31, 0x42, 0x0f, 0xdd, 0x08, 0x3d,
	0x24, 0xf5, 0x50, 0x23, 0x94, 0x1c, 0xb8, 0xf8, 0xdd, 0x40, 0xa2, 0x1e, 0xa9, 0x89, 0x05, 0xce,
	0x60, 0x6d, 0x42, 0xba, 0x5c, 0xc2, 0xb9, 0x89, 0x15, 0x60, 0x65, 0xf1, 0xc5, 0x99, 0x55, 0xa9,
	0xf1, 0x49, 0x95, 0x25, 0xa9, 0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x02, 0xb9, 0x89,
	0x15, 0x20, 0xb5, 0xc1, 0x99, 0x55, 0xa9, 0x4e, 0x20, 0x71, 0xa7, 0x68, 0x2e, 0xbb, 0xcc, 0x7c,
	0x3d, 0xb0, 0xc5, 0x05, 0x45, 0xf9, 0x15, 0x95, 0x7a, 0xa4, 0xba, 0xc1, 0x89, 0x0f, 0xee, 0x82,
	0x00, 0x90, 0x2f, 0x02, 0x18, 0xa3, 0xd8, 0xa1, 0x52, 0x49, 0x6c, 0x60, 0x7f, 0x19, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xa5, 0xec, 0x3e, 0x55, 0x1c, 0x01, 0x00, 0x00,
}
