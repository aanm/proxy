// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/fault/v2/fault.proto

package v2

import (
	fmt "fmt"
	route "github.com/cilium/proxy/go/envoy/api/v2/route"
	v2 "github.com/cilium/proxy/go/envoy/config/filter/fault/v2"
	_type "github.com/cilium/proxy/go/envoy/type"
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

type FaultAbort struct {
	// Types that are valid to be assigned to ErrorType:
	//	*FaultAbort_HttpStatus
	ErrorType isFaultAbort_ErrorType `protobuf_oneof:"error_type"`
	// The percentage of requests/operations/connections that will be aborted with the error code
	// provided.
	Percentage           *_type.FractionalPercent `protobuf:"bytes,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FaultAbort) Reset()         { *m = FaultAbort{} }
func (m *FaultAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort) ProtoMessage()    {}
func (*FaultAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_26070db6b6576d5c, []int{0}
}

func (m *FaultAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultAbort.Unmarshal(m, b)
}
func (m *FaultAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultAbort.Marshal(b, m, deterministic)
}
func (m *FaultAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort.Merge(m, src)
}
func (m *FaultAbort) XXX_Size() int {
	return xxx_messageInfo_FaultAbort.Size(m)
}
func (m *FaultAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort proto.InternalMessageInfo

type isFaultAbort_ErrorType interface {
	isFaultAbort_ErrorType()
}

type FaultAbort_HttpStatus struct {
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

func (*FaultAbort_HttpStatus) isFaultAbort_ErrorType() {}

func (m *FaultAbort) GetErrorType() isFaultAbort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *FaultAbort) GetHttpStatus() uint32 {
	if x, ok := m.GetErrorType().(*FaultAbort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (m *FaultAbort) GetPercentage() *_type.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultAbort) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultAbort_HttpStatus)(nil),
	}
}

type HTTPFault struct {
	// If specified, the filter will inject delays based on the values in the
	// object.
	Delay *v2.FaultDelay `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	// If specified, the filter will abort requests based on the values in
	// the object. At least *abort* or *delay* must be specified.
	Abort *FaultAbort `protobuf:"bytes,2,opt,name=abort,proto3" json:"abort,omitempty"`
	// Specifies the name of the (destination) upstream cluster that the
	// filter should match on. Fault injection will be restricted to requests
	// bound to the specific upstream cluster.
	UpstreamCluster string `protobuf:"bytes,3,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	// Specifies a set of headers that the filter should match on. The fault
	// injection filter can be applied selectively to requests that match a set of
	// headers specified in the fault filter config. The chances of actual fault
	// injection further depend on the value of the :ref:`percentage
	// <envoy_api_field_config.filter.http.fault.v2.FaultAbort.percentage>` field.
	// The filter will check the request's headers against all the specified
	// headers in the filter config. A match will happen if all the headers in the
	// config are present in the request with the same values (or based on
	// presence if the *value* field is not in the config).
	Headers []*route.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	// Faults are injected for the specified list of downstream hosts. If this
	// setting is not set, faults are injected for all downstream nodes.
	// Downstream node name is taken from :ref:`the HTTP
	// x-envoy-downstream-service-node
	// <config_http_conn_man_headers_downstream-service-node>` header and compared
	// against downstream_nodes list.
	DownstreamNodes []string `protobuf:"bytes,5,rep,name=downstream_nodes,json=downstreamNodes,proto3" json:"downstream_nodes,omitempty"`
	// The maximum number of faults that can be active at a single time via the configured fault
	// filter. Note that because this setting can be overridden at the route level, it's possible
	// for the number of active faults to be greater than this value (if injected via a different
	// route). If not specified, defaults to unlimited. This setting can be overridden via
	// `runtime <config_http_filters_fault_injection_runtime>` and any faults that are not injected
	// due to overflow will be indicated via the `faults_overflow
	// <config_http_filters_fault_injection_stats>` stat.
	//
	// .. attention::
	//   Like other :ref:`circuit breakers <arch_overview_circuit_break>` in Envoy, this is a fuzzy
	//   limit. It's possible for the number of active faults to rise slightly above the configured
	//   amount due to the implementation details.
	MaxActiveFaults *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=max_active_faults,json=maxActiveFaults,proto3" json:"max_active_faults,omitempty"`
	// The response rate limit to be applied to the response body of the stream. When configured,
	// the percentage can be overridden by the :ref:`fault.http.rate_limit.response_percent
	// <config_http_filters_fault_injection_runtime>` runtime key.
	//
	// .. attention::
	//  This is a per-stream limit versus a connection level limit. This means that concurrent streams
	//  will each get an independent limit.
	ResponseRateLimit    *v2.FaultRateLimit `protobuf:"bytes,7,opt,name=response_rate_limit,json=responseRateLimit,proto3" json:"response_rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HTTPFault) Reset()         { *m = HTTPFault{} }
func (m *HTTPFault) String() string { return proto.CompactTextString(m) }
func (*HTTPFault) ProtoMessage()    {}
func (*HTTPFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_26070db6b6576d5c, []int{1}
}

func (m *HTTPFault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPFault.Unmarshal(m, b)
}
func (m *HTTPFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPFault.Marshal(b, m, deterministic)
}
func (m *HTTPFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFault.Merge(m, src)
}
func (m *HTTPFault) XXX_Size() int {
	return xxx_messageInfo_HTTPFault.Size(m)
}
func (m *HTTPFault) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFault.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFault proto.InternalMessageInfo

func (m *HTTPFault) GetDelay() *v2.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFault) GetAbort() *FaultAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *HTTPFault) GetUpstreamCluster() string {
	if m != nil {
		return m.UpstreamCluster
	}
	return ""
}

func (m *HTTPFault) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPFault) GetDownstreamNodes() []string {
	if m != nil {
		return m.DownstreamNodes
	}
	return nil
}

func (m *HTTPFault) GetMaxActiveFaults() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxActiveFaults
	}
	return nil
}

func (m *HTTPFault) GetResponseRateLimit() *v2.FaultRateLimit {
	if m != nil {
		return m.ResponseRateLimit
	}
	return nil
}

func init() {
	proto.RegisterType((*FaultAbort)(nil), "envoy.config.filter.http.fault.v2.FaultAbort")
	proto.RegisterType((*HTTPFault)(nil), "envoy.config.filter.http.fault.v2.HTTPFault")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/fault/v2/fault.proto", fileDescriptor_26070db6b6576d5c)
}

var fileDescriptor_26070db6b6576d5c = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x71, 0x92, 0xb6, 0x1b, 0x55, 0x4d, 0xdd, 0x03, 0x56, 0x04, 0x55, 0xda, 0x93, 0x41,
	0x74, 0x8d, 0xdc, 0x23, 0x02, 0xa9, 0x09, 0xaa, 0x02, 0x02, 0x14, 0x99, 0xc2, 0x01, 0x21, 0x59,
	0x93, 0x78, 0x92, 0x58, 0x72, 0xbc, 0xab, 0xdd, 0xb5, 0x9b, 0xfc, 0x09, 0x9f, 0x82, 0x38, 0xf5,
	0x88, 0xc4, 0x0f, 0xf0, 0x0b, 0xfd, 0x0b, 0xb4, 0xbb, 0x31, 0xe5, 0x10, 0xa9, 0x5c, 0xa2, 0xcd,
	0xcc, 0x7b, 0xf3, 0xe6, 0x3d, 0x0f, 0x39, 0xc3, 0xa2, 0x62, 0xeb, 0x70, 0xca, 0x8a, 0x59, 0x36,
	0x0f, 0x67, 0x59, 0xae, 0x50, 0x84, 0x0b, 0xa5, 0x78, 0x38, 0x83, 0x32, 0x57, 0x61, 0x15, 0xd9,
	0x07, 0xe5, 0x82, 0x29, 0xe6, 0x9d, 0x18, 0x38, 0xb5, 0x70, 0x6a, 0xe1, 0x54, 0xc3, 0xa9, 0x45,
	0x55, 0x51, 0xef, 0xd8, 0x4e, 0x04, 0x9e, 0x69, 0xb2, 0x60, 0xa5, 0x42, 0xfb, 0x6b, 0x47, 0xf4,
	0x82, 0x6d, 0x8a, 0xdb, 0xc4, 0x7a, 0xbe, 0x45, 0xaa, 0x35, 0xc7, 0x90, 0xa3, 0x98, 0x62, 0x51,
	0x77, 0x8e, 0xe7, 0x8c, 0xcd, 0x73, 0x0c, 0xcd, 0xbf, 0x49, 0x39, 0x0b, 0xaf, 0x05, 0x70, 0x8e,
	0x42, 0x6e, 0xfa, 0x0f, 0x2b, 0xc8, 0xb3, 0x14, 0x14, 0x86, 0xf5, 0xc3, 0x36, 0x4e, 0xbf, 0x39,
	0x84, 0x5c, 0x6a, 0x89, 0x8b, 0x09, 0x13, 0xca, 0x7b, 0x4e, 0x3a, 0x7a, 0xf9, 0x44, 0x2a, 0x50,
	0xa5, 0xf4, 0x1b, 0x7d, 0x27, 0xd8, 0x1f, 0xec, 0xff, 0xb8, 0xbd, 0x71, 0x77, 0x9f, 0xb6, 0xbb,
	0xbf, 0x9b, 0xc1, 0x4f, 0x67, 0xf4, 0x20, 0x26, 0x1a, 0xf3, 0xd1, 0x40, 0xbc, 0x97, 0x84, 0x6c,
	0x56, 0x81, 0x39, 0xfa, 0x6e, 0xdf, 0x09, 0x3a, 0xd1, 0x63, 0x6a, 0x53, 0xd1, 0x8b, 0xd2, 0x4b,
	0x01, 0x53, 0x95, 0xb1, 0x02, 0xf2, 0xb1, 0xc5, 0xc5, 0xff, 0x10, 0x06, 0x47, 0x84, 0xa0, 0x10,
	0x4c, 0x24, 0x1a, 0xeb, 0xb5, 0xbe, 0xdf, 0xde, 0xb8, 0xce, 0xdb, 0xe6, 0xae, 0xd3, 0x6d, 0x9c,
	0xfe, 0x72, 0xc9, 0xde, 0xe8, 0xea, 0x6a, 0x6c, 0xd6, 0xf3, 0x5e, 0x91, 0x56, 0x8a, 0x39, 0xac,
	0x7d, 0xc7, 0x48, 0x04, 0x74, 0x5b, 0xf0, 0x75, 0xe6, 0xd4, 0x70, 0x5e, 0x6b, 0x7c, 0x6c, 0x69,
	0xde, 0x90, 0xb4, 0x40, 0x5b, 0x34, 0x9e, 0x3a, 0xd1, 0x19, 0xbd, 0xf7, 0xc3, 0xd1, 0xbb, 0x5c,
	0x62, 0xcb, 0xf5, 0x9e, 0x90, 0x6e, 0xc9, 0xa5, 0x12, 0x08, 0xcb, 0x64, 0x9a, 0x97, 0x52, 0xa1,
	0x30, 0x96, 0xf7, 0xe2, 0x83, 0xba, 0x3e, 0xb4, 0x65, 0xef, 0x05, 0xd9, 0x59, 0x20, 0xa4, 0x28,
	0xa4, 0xdf, 0xec, 0xbb, 0x41, 0x27, 0x3a, 0xd9, 0x28, 0x02, 0xcf, 0xf4, 0x70, 0x7b, 0x01, 0x23,
	0x03, 0x79, 0x0f, 0x6a, 0xba, 0x40, 0x11, 0xd7, 0x0c, 0xad, 0x93, 0xb2, 0xeb, 0x62, 0xa3, 0x54,
	0xb0, 0x14, 0xa5, 0xdf, 0xea, 0xbb, 0x5a, 0xe7, 0xae, 0xfe, 0x41, 0x97, 0xbd, 0x11, 0x39, 0x5c,
	0xc2, 0x2a, 0xd1, 0x19, 0x57, 0x98, 0x98, 0xdd, 0xa5, 0xdf, 0x36, 0x1e, 0x1f, 0x51, 0x7b, 0x15,
	0xb4, 0xbe, 0x0a, 0xfa, 0xe9, 0x4d, 0xa1, 0xce, 0xa3, 0xcf, 0x90, 0x97, 0x18, 0x1f, 0x2c, 0x61,
	0x75, 0x61, 0x58, 0xc6, 0xa7, 0xf4, 0xbe, 0x92, 0x23, 0x81, 0x92, 0xb3, 0x42, 0x62, 0x22, 0x40,
	0x61, 0x92, 0x67, 0xcb, 0x4c, 0xf9, 0x3b, 0x66, 0xd6, 0xb3, 0xff, 0xc8, 0x3b, 0x06, 0x85, 0xef,
	0x34, 0x27, 0x3e, 0xac, 0x07, 0xfd, 0x2d, 0x0d, 0x86, 0x24, 0xcc, 0x98, 0x1d, 0xc2, 0x05, 0x5b,
	0xad, 0xef, 0xcf, 0x7f, 0x60, 0x0f, 0x73, 0xac, 0x97, 0x1f, 0x3b, 0x5f, 0x1a, 0x55, 0x34, 0x69,
	0x1b, 0x27, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x37, 0x72, 0x46, 0x3f, 0xa5, 0x03, 0x00,
	0x00,
}
