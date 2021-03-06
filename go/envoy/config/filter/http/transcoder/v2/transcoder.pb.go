// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/transcoder/v2/transcoder.proto

package v2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GrpcJsonTranscoder struct {
	// Types that are valid to be assigned to DescriptorSet:
	//	*GrpcJsonTranscoder_ProtoDescriptor
	//	*GrpcJsonTranscoder_ProtoDescriptorBin
	DescriptorSet isGrpcJsonTranscoder_DescriptorSet `protobuf_oneof:"descriptor_set"`
	// A list of strings that
	// supplies the fully qualified service names (i.e. "package_name.service_name") that
	// the transcoder will translate. If the service name doesn't exist in ``proto_descriptor``,
	// Envoy will fail at startup. The ``proto_descriptor`` may contain more services than
	// the service names specified here, but they won't be translated.
	Services []string `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
	// Control options for response JSON. These options are passed directly to
	// `JsonPrintOptions <https://developers.google.com/protocol-buffers/docs/reference/cpp/
	// google.protobuf.util.json_util#JsonPrintOptions>`_.
	PrintOptions *GrpcJsonTranscoder_PrintOptions `protobuf:"bytes,3,opt,name=print_options,json=printOptions,proto3" json:"print_options,omitempty"`
	// Whether to keep the incoming request route after the outgoing headers have been transformed to
	// the match the upstream gRPC service. Note: This means that routes for gRPC services that are
	// not transcoded cannot be used in combination with *match_incoming_request_route*.
	MatchIncomingRequestRoute bool `protobuf:"varint,5,opt,name=match_incoming_request_route,json=matchIncomingRequestRoute,proto3" json:"match_incoming_request_route,omitempty"`
	// A list of query parameters to be ignored for transcoding method mapping.
	// By default, the transcoder filter will not transcode a request if there are any
	// unknown/invalid query parameters.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {
	//         option (google.api.http) = {
	//           get: "/shelves/{shelf}"
	//         };
	//       }
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The request ``/shelves/100?foo=bar`` will not be mapped to ``GetShelf``` because variable
	// binding for ``foo`` is not defined. Adding ``foo`` to ``ignored_query_parameters`` will allow
	// the same request to be mapped to ``GetShelf``.
	IgnoredQueryParameters []string `protobuf:"bytes,6,rep,name=ignored_query_parameters,json=ignoredQueryParameters,proto3" json:"ignored_query_parameters,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *GrpcJsonTranscoder) Reset()         { *m = GrpcJsonTranscoder{} }
func (m *GrpcJsonTranscoder) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder) ProtoMessage()    {}
func (*GrpcJsonTranscoder) Descriptor() ([]byte, []int) {
	return fileDescriptor_540f2f6de8b0585c, []int{0}
}

func (m *GrpcJsonTranscoder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder.Marshal(b, m, deterministic)
}
func (m *GrpcJsonTranscoder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder.Merge(m, src)
}
func (m *GrpcJsonTranscoder) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder.Size(m)
}
func (m *GrpcJsonTranscoder) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder proto.InternalMessageInfo

type isGrpcJsonTranscoder_DescriptorSet interface {
	isGrpcJsonTranscoder_DescriptorSet()
}

type GrpcJsonTranscoder_ProtoDescriptor struct {
	ProtoDescriptor string `protobuf:"bytes,1,opt,name=proto_descriptor,json=protoDescriptor,proto3,oneof"`
}

type GrpcJsonTranscoder_ProtoDescriptorBin struct {
	ProtoDescriptorBin []byte `protobuf:"bytes,4,opt,name=proto_descriptor_bin,json=protoDescriptorBin,proto3,oneof"`
}

func (*GrpcJsonTranscoder_ProtoDescriptor) isGrpcJsonTranscoder_DescriptorSet() {}

func (*GrpcJsonTranscoder_ProtoDescriptorBin) isGrpcJsonTranscoder_DescriptorSet() {}

func (m *GrpcJsonTranscoder) GetDescriptorSet() isGrpcJsonTranscoder_DescriptorSet {
	if m != nil {
		return m.DescriptorSet
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetProtoDescriptor() string {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptor); ok {
		return x.ProtoDescriptor
	}
	return ""
}

func (m *GrpcJsonTranscoder) GetProtoDescriptorBin() []byte {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptorBin); ok {
		return x.ProtoDescriptorBin
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetPrintOptions() *GrpcJsonTranscoder_PrintOptions {
	if m != nil {
		return m.PrintOptions
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetMatchIncomingRequestRoute() bool {
	if m != nil {
		return m.MatchIncomingRequestRoute
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoredQueryParameters() []string {
	if m != nil {
		return m.IgnoredQueryParameters
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GrpcJsonTranscoder) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GrpcJsonTranscoder_ProtoDescriptor)(nil),
		(*GrpcJsonTranscoder_ProtoDescriptorBin)(nil),
	}
}

type GrpcJsonTranscoder_PrintOptions struct {
	// Whether to add spaces, line breaks and indentation to make the JSON
	// output easy to read. Defaults to false.
	AddWhitespace bool `protobuf:"varint,1,opt,name=add_whitespace,json=addWhitespace,proto3" json:"add_whitespace,omitempty"`
	// Whether to always print primitive fields. By default primitive
	// fields with default values will be omitted in JSON output. For
	// example, an int32 field set to 0 will be omitted. Setting this flag to
	// true will override the default behavior and print primitive fields
	// regardless of their values. Defaults to false.
	AlwaysPrintPrimitiveFields bool `protobuf:"varint,2,opt,name=always_print_primitive_fields,json=alwaysPrintPrimitiveFields,proto3" json:"always_print_primitive_fields,omitempty"`
	// Whether to always print enums as ints. By default they are rendered
	// as strings. Defaults to false.
	AlwaysPrintEnumsAsInts bool `protobuf:"varint,3,opt,name=always_print_enums_as_ints,json=alwaysPrintEnumsAsInts,proto3" json:"always_print_enums_as_ints,omitempty"`
	// Whether to preserve proto field names. By default protobuf will
	// generate JSON field names using the ``json_name`` option, or lower camel case,
	// in that order. Setting this flag will preserve the original field names. Defaults to false.
	PreserveProtoFieldNames bool     `protobuf:"varint,4,opt,name=preserve_proto_field_names,json=preserveProtoFieldNames,proto3" json:"preserve_proto_field_names,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *GrpcJsonTranscoder_PrintOptions) Reset()         { *m = GrpcJsonTranscoder_PrintOptions{} }
func (m *GrpcJsonTranscoder_PrintOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder_PrintOptions) ProtoMessage()    {}
func (*GrpcJsonTranscoder_PrintOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_540f2f6de8b0585c, []int{0, 0}
}

func (m *GrpcJsonTranscoder_PrintOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Marshal(b, m, deterministic)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Merge(m, src)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Size(m)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder_PrintOptions proto.InternalMessageInfo

func (m *GrpcJsonTranscoder_PrintOptions) GetAddWhitespace() bool {
	if m != nil {
		return m.AddWhitespace
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintPrimitiveFields() bool {
	if m != nil {
		return m.AlwaysPrintPrimitiveFields
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintEnumsAsInts() bool {
	if m != nil {
		return m.AlwaysPrintEnumsAsInts
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetPreserveProtoFieldNames() bool {
	if m != nil {
		return m.PreserveProtoFieldNames
	}
	return false
}

func init() {
	proto.RegisterType((*GrpcJsonTranscoder)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder")
	proto.RegisterType((*GrpcJsonTranscoder_PrintOptions)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder.PrintOptions")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/transcoder/v2/transcoder.proto", fileDescriptor_540f2f6de8b0585c)
}

var fileDescriptor_540f2f6de8b0585c = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x6a, 0x14, 0x31,
	0x14, 0x86, 0x9b, 0xae, 0x5b, 0xa6, 0x71, 0xdb, 0x4a, 0x90, 0xee, 0xb8, 0x28, 0x2c, 0x82, 0x65,
	0x41, 0x98, 0x81, 0x55, 0x50, 0xf4, 0x42, 0xba, 0xa8, 0x6d, 0x05, 0x75, 0x0d, 0x82, 0xe0, 0x4d,
	0x48, 0x67, 0xce, 0xee, 0x06, 0x76, 0x92, 0x34, 0xc9, 0x4e, 0xdd, 0xd7, 0xd0, 0x97, 0x11, 0xaf,
	0xfa, 0x3a, 0xbd, 0xf4, 0x0d, 0x24, 0x49, 0x77, 0x1d, 0xeb, 0x4d, 0xef, 0x26, 0xf3, 0x7f, 0xff,
	0x39, 0xc9, 0xf9, 0x0f, 0x7e, 0x06, 0xb2, 0x56, 0xcb, 0xbc, 0x50, 0x72, 0x22, 0xa6, 0xf9, 0x44,
	0xcc, 0x1d, 0x98, 0x7c, 0xe6, 0x9c, 0xce, 0x9d, 0xe1, 0xd2, 0x16, 0xaa, 0x04, 0x93, 0xd7, 0xc3,
	0xc6, 0x29, 0xd3, 0x46, 0x39, 0x45, 0x0e, 0x82, 0x31, 0x8b, 0xc6, 0x2c, 0x1a, 0x33, 0x6f, 0xcc,
	0x1a, 0x68, 0x3d, 0xec, 0x75, 0x6b, 0x3e, 0x17, 0x25, 0x77, 0x90, 0xaf, 0x3e, 0x62, 0x81, 0x87,
	0x3f, 0xda, 0x98, 0x1c, 0x19, 0x5d, 0xbc, 0xb3, 0x4a, 0x7e, 0x5e, 0x5b, 0xc8, 0x63, 0x7c, 0x27,
	0xe8, 0xac, 0x04, 0x5b, 0x18, 0xa1, 0x9d, 0x32, 0x29, 0xea, 0xa3, 0xc1, 0xf6, 0xf1, 0x06, 0xdd,
	0x0b, 0xca, 0xeb, 0xb5, 0x40, 0x86, 0xf8, 0xee, 0x75, 0x98, 0x9d, 0x0a, 0x99, 0xde, 0xea, 0xa3,
	0x41, 0xe7, 0x78, 0x83, 0x92, 0x6b, 0x86, 0x91, 0x90, 0xe4, 0x00, 0x27, 0x16, 0x4c, 0x2d, 0x0a,
	0xb0, 0xe9, 0x66, 0xbf, 0x35, 0xd8, 0x1e, 0xe1, 0x5f, 0x97, 0x17, 0xad, 0xf6, 0x77, 0xb4, 0x99,
	0x20, 0xba, 0xd6, 0xc8, 0x1c, 0xef, 0x68, 0x23, 0xa4, 0x63, 0x4a, 0x3b, 0xa1, 0xa4, 0x4d, 0x5b,
	0x7d, 0x34, 0xb8, 0x3d, 0x3c, 0xca, 0x6e, 0xf6, 0xf0, 0xec, 0xff, 0xb7, 0x65, 0x63, 0x5f, 0xef,
	0x63, 0x2c, 0x47, 0x3b, 0xba, 0x71, 0x22, 0xaf, 0xf0, 0xfd, 0x8a, 0xbb, 0x62, 0xc6, 0x84, 0x2c,
	0x54, 0x25, 0xe4, 0x94, 0x19, 0x38, 0x5b, 0x80, 0x75, 0xcc, 0xa8, 0x85, 0x83, 0xb4, 0xdd, 0x47,
	0x83, 0x84, 0xde, 0x0b, 0xcc, 0xc9, 0x15, 0x42, 0x23, 0x41, 0x3d, 0x40, 0x9e, 0xe3, 0x54, 0x4c,
	0xa5, 0x32, 0x50, 0xb2, 0xb3, 0x05, 0x98, 0x25, 0xd3, 0xdc, 0xf0, 0x0a, 0x1c, 0x18, 0x9b, 0x6e,
	0xf9, 0x67, 0xd2, 0xfd, 0x2b, 0xfd, 0x93, 0x97, 0xc7, 0x6b, 0xb5, 0xf7, 0x1b, 0xe1, 0x4e, 0xf3,
	0x66, 0xe4, 0x11, 0xde, 0xe5, 0x65, 0xc9, 0xce, 0x67, 0xc2, 0x81, 0xd5, 0xbc, 0x80, 0x10, 0x40,
	0x42, 0x77, 0x78, 0x59, 0x7e, 0x59, 0xff, 0x24, 0x87, 0xf8, 0x01, 0x9f, 0x9f, 0xf3, 0xa5, 0x65,
	0x71, 0x4e, 0xda, 0x88, 0x4a, 0x38, 0x51, 0x03, 0x9b, 0x08, 0x98, 0x97, 0x7e, 0xba, 0xde, 0xd5,
	0x8b, 0x50, 0xe8, 0x30, 0x5e, 0x21, 0x6f, 0x03, 0x41, 0x5e, 0xe0, 0xde, 0x3f, 0x25, 0x40, 0x2e,
	0x2a, 0xcb, 0xb8, 0x65, 0x42, 0xba, 0x38, 0xf0, 0x84, 0xee, 0x37, 0xfc, 0x6f, 0xbc, 0x7e, 0x68,
	0x4f, 0xa4, 0xb3, 0xe4, 0x25, 0xee, 0x69, 0x03, 0x3e, 0x2e, 0x60, 0x71, 0x09, 0x42, 0x5b, 0x26,
	0x79, 0x05, 0x36, 0x6c, 0x40, 0x42, 0xbb, 0x2b, 0x62, 0xec, 0x81, 0xd0, 0xf4, 0x83, 0x97, 0x47,
	0x5d, 0xbc, 0xdb, 0x58, 0x19, 0x0b, 0x8e, 0xb4, 0x7f, 0x5e, 0x5e, 0xb4, 0xd0, 0xe8, 0x3d, 0x7e,
	0x2a, 0x54, 0x8c, 0x58, 0x1b, 0xf5, 0x6d, 0x79, 0xc3, 0xb4, 0x47, 0x7b, 0x7f, 0x63, 0x0e, 0xbd,
	0xc6, 0xe8, 0xeb, 0x66, 0x3d, 0x3c, 0xdd, 0x0a, 0x37, 0x7b, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x50, 0xe7, 0xde, 0xfd, 0x67, 0x03, 0x00, 0x00,
}
