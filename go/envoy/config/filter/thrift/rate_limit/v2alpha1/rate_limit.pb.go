// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/thrift/rate_limit/v2alpha1/rate_limit.proto

package v2alpha1

import (
	fmt "fmt"
	v2 "github.com/cilium/proxy/go/envoy/config/ratelimit/v2"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// [#comment:next free field: 5]
type RateLimit struct {
	// The rate limit domain to use in the rate limit service request.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Specifies the rate limit configuration stage. Each configured rate limit filter performs a
	// rate limit check using descriptors configured in the
	// :ref:`envoy_api_msg_config.filter.network.thrift_proxy.v2alpha1.RouteAction` for the request.
	// Only those entries with a matching stage number are used for a given filter. If not set, the
	// default stage number is 0.
	//
	// .. note::
	//
	//  The filter supports a range of 0 - 10 inclusively for stage numbers.
	Stage uint32 `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *duration.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny bool `protobuf:"varint,4,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	// Configuration for an external rate limit service provider. If not
	// specified, any calls to the rate limit service will immediately return
	// success.
	RateLimitService     *v2.RateLimitServiceConfig `protobuf:"bytes,5,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_961acdee13c1bd42, []int{0}
}

func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetStage() uint32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *RateLimit) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v2.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.thrift.rate_limit.v2alpha1.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/thrift/rate_limit/v2alpha1/rate_limit.proto", fileDescriptor_961acdee13c1bd42)
}

var fileDescriptor_961acdee13c1bd42 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xbd, 0x4e, 0xe3, 0x40,
	0x18, 0xd4, 0x3a, 0x3f, 0x97, 0xec, 0xe9, 0xee, 0x72, 0xd6, 0x49, 0x67, 0x52, 0x10, 0x03, 0x8d,
	0x95, 0x62, 0x57, 0x09, 0x15, 0x12, 0x12, 0x92, 0x49, 0x09, 0x52, 0x64, 0xba, 0x34, 0xd6, 0x26,
	0x5e, 0x3b, 0x2b, 0x6c, 0x7f, 0xd1, 0x66, 0x6d, 0xe1, 0x57, 0xe0, 0x09, 0xe8, 0x79, 0x0b, 0x2a,
	0xde, 0x84, 0x9a, 0xb7, 0x40, 0xf6, 0xda, 0x21, 0xa1, 0xa3, 0xfb, 0xfc, 0xcd, 0x78, 0x66, 0xbe,
	0x59, 0x7c, 0xc5, 0xd3, 0x1c, 0x0a, 0xba, 0x82, 0x34, 0x14, 0x11, 0x0d, 0x45, 0xac, 0xb8, 0xa4,
	0x6a, 0x2d, 0x45, 0xa8, 0xa8, 0x64, 0x8a, 0xfb, 0xb1, 0x48, 0x84, 0xa2, 0xf9, 0x94, 0xc5, 0x9b,
	0x35, 0x9b, 0xec, 0xed, 0xc8, 0x46, 0x82, 0x02, 0x93, 0x54, 0x02, 0x44, 0x0b, 0x10, 0x2d, 0x40,
	0xb4, 0x00, 0xd9, 0x23, 0x37, 0x02, 0xc3, 0xb3, 0x03, 0xc3, 0x92, 0xd1, 0x38, 0x50, 0x19, 0x6f,
	0xb5, 0xe8, 0xf0, 0x38, 0x02, 0x88, 0x62, 0x4e, 0xab, 0xaf, 0x65, 0x16, 0xd2, 0x20, 0x93, 0x4c,
	0x09, 0x48, 0x6b, 0xfc, 0x7f, 0xce, 0x62, 0x11, 0x30, 0xc5, 0x69, 0x33, 0xd4, 0xc0, 0xbf, 0x08,
	0x22, 0xa8, 0x46, 0x5a, 0x4e, 0x7a, 0x7b, 0xfa, 0x6c, 0xe0, 0xbe, 0xc7, 0x14, 0xbf, 0x29, 0x9d,
	0xcc, 0x13, 0xdc, 0x0d, 0x20, 0x61, 0x22, 0xb5, 0x90, 0x8d, 0x9c, 0xbe, 0xdb, 0x7f, 0x79, 0x7f,
	0x6d, 0xb5, 0xa5, 0x61, 0x23, 0xaf, 0x06, 0xcc, 0x11, 0xee, 0x6c, 0x15, 0x8b, 0xb8, 0x65, 0xd8,
	0xc8, 0xf9, 0x55, 0x33, 0xc6, 0x86, 0x85, 0x3d, 0xbd, 0x37, 0x2f, 0xf0, 0x0f, 0x25, 0x12, 0x0e,
	0x99, 0xb2, 0x5a, 0x36, 0x72, 0x7e, 0x4e, 0x8f, 0x88, 0x8e, 0x4c, 0x9a, 0xc8, 0x64, 0x56, 0x47,
	0x76, 0xdb, 0x4f, 0x6f, 0x23, 0xe4, 0x35, 0x7c, 0x73, 0x8c, 0xff, 0x86, 0x4c, 0xc4, 0x99, 0xe4,
	0x7e, 0x02, 0x01, 0xf7, 0x03, 0x9e, 0x16, 0x56, 0xdb, 0x46, 0x4e, 0xcf, 0xfb, 0x53, 0x03, 0xb7,
	0x10, 0xf0, 0x19, 0x4f, 0x0b, 0xf3, 0x1e, 0x9b, 0x9f, 0x1d, 0xfa, 0x5b, 0x2e, 0x73, 0xb1, 0xe2,
	0x56, 0xa7, 0x72, 0x9c, 0x1c, 0x36, 0xbf, 0x6b, 0x92, 0xe4, 0x53, 0xb2, 0x3b, 0xf6, 0x4e, 0xff,
	0x72, 0x5d, 0x71, 0x5c, 0x5c, 0xde, 0xd1, 0x79, 0x44, 0xc6, 0x00, 0x79, 0x03, 0xf9, 0x85, 0xe3,
	0x2e, 0xf0, 0xa5, 0x00, 0x2d, 0xba, 0x91, 0xf0, 0x50, 0x7c, 0xf3, 0x65, 0xdd, 0xdf, 0x3b, 0xd7,
	0x79, 0xd9, 0xc1, 0x1c, 0x2d, 0x7a, 0x0d, 0xb6, 0xec, 0x56, 0xb5, 0x9c, 0x7f, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xb4, 0x29, 0xac, 0xd4, 0x6f, 0x02, 0x00, 0x00,
}
