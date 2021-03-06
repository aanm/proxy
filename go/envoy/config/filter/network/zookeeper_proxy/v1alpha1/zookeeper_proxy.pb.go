// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/zookeeper_proxy/v1alpha1/zookeeper_proxy.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/lyft/protoc-gen-validate/validate"
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

// [#protodoc-title: ZooKeeper proxy]
// ZooKeeper Proxy :ref:`configuration overview <config_network_filters_zookeeper_proxy>`.
type ZooKeeperProxy struct {
	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_zookeeper_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// [#not-implemented-hide:] The optional path to use for writing ZooKeeper access logs.
	// If the access log field is empty, access logs will not be written.
	AccessLog string `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// Messages — requests, responses and events — that are bigger than this value will
	// be ignored. If it is not set, the default value is 1Mb.
	//
	// The value here should match the jute.maxbuffer property in your cluster configuration:
	//
	// https://zookeeper.apache.org/doc/r3.4.10/zookeeperAdmin.html#Unsafe+Options
	//
	// if that is set. If it isn't, ZooKeeper's default is also 1Mb.
	MaxPacketBytes       *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=max_packet_bytes,json=maxPacketBytes,proto3" json:"max_packet_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ZooKeeperProxy) Reset()         { *m = ZooKeeperProxy{} }
func (m *ZooKeeperProxy) String() string { return proto.CompactTextString(m) }
func (*ZooKeeperProxy) ProtoMessage()    {}
func (*ZooKeeperProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_05247350458709ad, []int{0}
}

func (m *ZooKeeperProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZooKeeperProxy.Unmarshal(m, b)
}
func (m *ZooKeeperProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZooKeeperProxy.Marshal(b, m, deterministic)
}
func (m *ZooKeeperProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZooKeeperProxy.Merge(m, src)
}
func (m *ZooKeeperProxy) XXX_Size() int {
	return xxx_messageInfo_ZooKeeperProxy.Size(m)
}
func (m *ZooKeeperProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ZooKeeperProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ZooKeeperProxy proto.InternalMessageInfo

func (m *ZooKeeperProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ZooKeeperProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func (m *ZooKeeperProxy) GetMaxPacketBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxPacketBytes
	}
	return nil
}

func init() {
	proto.RegisterType((*ZooKeeperProxy)(nil), "envoy.config.filter.network.zookeeper_proxy.v1alpha1.ZooKeeperProxy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/zookeeper_proxy/v1alpha1/zookeeper_proxy.proto", fileDescriptor_05247350458709ad)
}

var fileDescriptor_05247350458709ad = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xbf, 0x4e, 0xfb, 0x30,
	0x10, 0xc7, 0xe5, 0xf6, 0xa7, 0x9f, 0xa8, 0x2b, 0x55, 0x28, 0x0c, 0x54, 0x15, 0xa0, 0x8a, 0xa9,
	0x62, 0xb0, 0xd5, 0x96, 0x27, 0xc8, 0x80, 0xc4, 0x9f, 0x21, 0x8a, 0x04, 0x43, 0x96, 0xc8, 0x09,
	0x97, 0x10, 0xc5, 0xcd, 0x59, 0x8e, 0x9b, 0x26, 0xbc, 0x0e, 0x6f, 0xc1, 0xc4, 0xeb, 0xf0, 0x16,
	0x28, 0x36, 0x59, 0x3a, 0xb2, 0xd9, 0xfe, 0xdc, 0x7d, 0xfc, 0xbd, 0xa3, 0x0f, 0x50, 0x35, 0xd8,
	0xf1, 0x14, 0xab, 0xac, 0xc8, 0x79, 0x56, 0x48, 0x03, 0x9a, 0x57, 0x60, 0x0e, 0xa8, 0x4b, 0xfe,
	0x8e, 0x58, 0x02, 0x28, 0xd0, 0xb1, 0xd2, 0xd8, 0x76, 0xbc, 0x59, 0x0b, 0xa9, 0xde, 0xc4, 0xfa,
	0x18, 0x30, 0xa5, 0xd1, 0xa0, 0x77, 0x6b, 0x5d, 0xcc, 0xb9, 0x98, 0x73, 0xb1, 0x5f, 0x17, 0x3b,
	0x6e, 0x19, 0x5c, 0x8b, 0xf3, 0x46, 0xc8, 0xe2, 0x55, 0x18, 0xe0, 0xc3, 0xc1, 0xe9, 0x16, 0x57,
	0x39, 0x62, 0x2e, 0x81, 0xdb, 0x5b, 0xb2, 0xcf, 0xf8, 0x41, 0x0b, 0xa5, 0x40, 0xd7, 0x8e, 0x5f,
	0x7f, 0x10, 0x3a, 0x8b, 0x10, 0x1f, 0xad, 0x35, 0xe8, 0xa5, 0xde, 0x0d, 0x9d, 0xd6, 0x46, 0x98,
	0x58, 0x69, 0xc8, 0x8a, 0x76, 0x4e, 0x96, 0x64, 0x35, 0xf1, 0x27, 0x9f, 0xdf, 0x5f, 0xe3, 0x7f,
	0x7a, 0xb4, 0x24, 0x21, 0xed, 0x69, 0x60, 0xa1, 0x77, 0x49, 0xa9, 0x48, 0x53, 0xa8, 0xeb, 0x58,
	0x62, 0x3e, 0x1f, 0xf5, 0xa5, 0xe1, 0xc4, 0xbd, 0x3c, 0x61, 0xee, 0xdd, 0xd1, 0xd3, 0x9d, 0x68,
	0x63, 0x25, 0xd2, 0x12, 0x4c, 0x9c, 0x74, 0x06, 0xea, 0xf9, 0x78, 0x49, 0x56, 0xd3, 0xcd, 0x05,
	0x73, 0xc1, 0xd8, 0x10, 0x8c, 0x3d, 0xdf, 0x57, 0x66, 0xbb, 0x79, 0x11, 0x72, 0x0f, 0xe1, 0x6c,
	0x27, 0xda, 0xc0, 0x36, 0xf9, 0x7d, 0x8f, 0x0f, 0xd4, 0x2f, 0x90, 0xd9, 0xcd, 0xb8, 0xc1, 0xff,
	0xb2, 0x24, 0xff, 0x2c, 0x1a, 0x88, 0x1d, 0x34, 0xe8, 0x7f, 0x0e, 0x48, 0x74, 0x32, 0x14, 0x24,
	0xff, 0x6d, 0x98, 0xed, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xf0, 0x9f, 0x6a, 0xd0, 0x01,
	0x00, 0x00,
}
