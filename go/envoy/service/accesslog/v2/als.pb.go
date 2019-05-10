// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/accesslog/v2/als.proto

package v2

import (
	context "context"
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	v2 "github.com/cilium/proxy/go/envoy/data/accesslog/v2"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Empty response for the StreamAccessLogs API. Will never be sent. See below.
type StreamAccessLogsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamAccessLogsResponse) Reset()         { *m = StreamAccessLogsResponse{} }
func (m *StreamAccessLogsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsResponse) ProtoMessage()    {}
func (*StreamAccessLogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{0}
}

func (m *StreamAccessLogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsResponse.Unmarshal(m, b)
}
func (m *StreamAccessLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsResponse.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsResponse.Merge(m, src)
}
func (m *StreamAccessLogsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsResponse.Size(m)
}
func (m *StreamAccessLogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsResponse proto.InternalMessageInfo

// Stream message for the StreamAccessLogs API. Envoy will open a stream to the server and stream
// access logs without ever expecting a response.
type StreamAccessLogsMessage struct {
	// Identifier data that will only be sent in the first message on the stream. This is effectively
	// structured metadata and is a performance optimization.
	Identifier *StreamAccessLogsMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Batches of log entries of a single type. Generally speaking, a given stream should only
	// ever include one type of log entry.
	//
	// Types that are valid to be assigned to LogEntries:
	//	*StreamAccessLogsMessage_HttpLogs
	//	*StreamAccessLogsMessage_TcpLogs
	LogEntries           isStreamAccessLogsMessage_LogEntries `protobuf_oneof:"log_entries"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *StreamAccessLogsMessage) Reset()         { *m = StreamAccessLogsMessage{} }
func (m *StreamAccessLogsMessage) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsMessage) ProtoMessage()    {}
func (*StreamAccessLogsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{1}
}

func (m *StreamAccessLogsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage.Merge(m, src)
}
func (m *StreamAccessLogsMessage) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage.Size(m)
}
func (m *StreamAccessLogsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage proto.InternalMessageInfo

func (m *StreamAccessLogsMessage) GetIdentifier() *StreamAccessLogsMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

type isStreamAccessLogsMessage_LogEntries interface {
	isStreamAccessLogsMessage_LogEntries()
}

type StreamAccessLogsMessage_HttpLogs struct {
	HttpLogs *StreamAccessLogsMessage_HTTPAccessLogEntries `protobuf:"bytes,2,opt,name=http_logs,json=httpLogs,proto3,oneof"`
}

type StreamAccessLogsMessage_TcpLogs struct {
	TcpLogs *StreamAccessLogsMessage_TCPAccessLogEntries `protobuf:"bytes,3,opt,name=tcp_logs,json=tcpLogs,proto3,oneof"`
}

func (*StreamAccessLogsMessage_HttpLogs) isStreamAccessLogsMessage_LogEntries() {}

func (*StreamAccessLogsMessage_TcpLogs) isStreamAccessLogsMessage_LogEntries() {}

func (m *StreamAccessLogsMessage) GetLogEntries() isStreamAccessLogsMessage_LogEntries {
	if m != nil {
		return m.LogEntries
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetHttpLogs() *StreamAccessLogsMessage_HTTPAccessLogEntries {
	if x, ok := m.GetLogEntries().(*StreamAccessLogsMessage_HttpLogs); ok {
		return x.HttpLogs
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetTcpLogs() *StreamAccessLogsMessage_TCPAccessLogEntries {
	if x, ok := m.GetLogEntries().(*StreamAccessLogsMessage_TcpLogs); ok {
		return x.TcpLogs
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StreamAccessLogsMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StreamAccessLogsMessage_HttpLogs)(nil),
		(*StreamAccessLogsMessage_TcpLogs)(nil),
	}
}

type StreamAccessLogsMessage_Identifier struct {
	// The node sending the access log messages over the stream.
	Node *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// The friendly name of the log configured in :ref:`CommonGrpcAccessLogConfig
	// <envoy_api_msg_config.accesslog.v2.CommonGrpcAccessLogConfig>`.
	LogName              string   `protobuf:"bytes,2,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamAccessLogsMessage_Identifier) Reset()         { *m = StreamAccessLogsMessage_Identifier{} }
func (m *StreamAccessLogsMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsMessage_Identifier) ProtoMessage()    {}
func (*StreamAccessLogsMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{1, 0}
}

func (m *StreamAccessLogsMessage_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage_Identifier.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage_Identifier.Merge(m, src)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage_Identifier.Size(m)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage_Identifier proto.InternalMessageInfo

func (m *StreamAccessLogsMessage_Identifier) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *StreamAccessLogsMessage_Identifier) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

// Wrapper for batches of HTTP access log entries.
type StreamAccessLogsMessage_HTTPAccessLogEntries struct {
	LogEntry             []*v2.HTTPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) Reset() {
	*m = StreamAccessLogsMessage_HTTPAccessLogEntries{}
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) String() string {
	return proto.CompactTextString(m)
}
func (*StreamAccessLogsMessage_HTTPAccessLogEntries) ProtoMessage() {}
func (*StreamAccessLogsMessage_HTTPAccessLogEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{1, 1}
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Merge(m, src)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Size(m)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries proto.InternalMessageInfo

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) GetLogEntry() []*v2.HTTPAccessLogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

// [#not-implemented-hide:]
// Wrapper for batches of TCP access log entries.
type StreamAccessLogsMessage_TCPAccessLogEntries struct {
	LogEntry             []*v2.TCPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) Reset() {
	*m = StreamAccessLogsMessage_TCPAccessLogEntries{}
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) String() string {
	return proto.CompactTextString(m)
}
func (*StreamAccessLogsMessage_TCPAccessLogEntries) ProtoMessage() {}
func (*StreamAccessLogsMessage_TCPAccessLogEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{1, 2}
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Merge(m, src)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Size(m)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries proto.InternalMessageInfo

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) GetLogEntry() []*v2.TCPAccessLogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamAccessLogsResponse)(nil), "envoy.service.accesslog.v2.StreamAccessLogsResponse")
	proto.RegisterType((*StreamAccessLogsMessage)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage")
	proto.RegisterType((*StreamAccessLogsMessage_Identifier)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.Identifier")
	proto.RegisterType((*StreamAccessLogsMessage_HTTPAccessLogEntries)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.HTTPAccessLogEntries")
	proto.RegisterType((*StreamAccessLogsMessage_TCPAccessLogEntries)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.TCPAccessLogEntries")
}

func init() {
	proto.RegisterFile("envoy/service/accesslog/v2/als.proto", fileDescriptor_e4f3a3a69261b513)
}

var fileDescriptor_e4f3a3a69261b513 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x77, 0xd2, 0xad, 0x9b, 0xbe, 0x5e, 0xca, 0xb8, 0xd0, 0x12, 0x3c, 0x94, 0x65, 0xc1,
	0xa2, 0x90, 0x40, 0x56, 0xf0, 0xa4, 0xb0, 0x11, 0xb1, 0x82, 0x2e, 0x25, 0xed, 0xc9, 0x83, 0xcb,
	0x6c, 0xf2, 0x8c, 0xa3, 0x69, 0x26, 0xcc, 0x0c, 0xc1, 0x1e, 0xbd, 0x89, 0x47, 0x0f, 0x7e, 0x03,
	0xbf, 0x83, 0x78, 0xda, 0xaf, 0xb3, 0xdf, 0x42, 0x66, 0x92, 0x8d, 0xda, 0x6d, 0x85, 0xed, 0x2d,
	0xcc, 0xbc, 0xff, 0xef, 0x37, 0x6f, 0x5e, 0x06, 0x8e, 0xb1, 0xa8, 0xc4, 0x2a, 0x50, 0x28, 0x2b,
	0x9e, 0x60, 0xc0, 0x92, 0x04, 0x95, 0xca, 0x45, 0x16, 0x54, 0x61, 0xc0, 0x72, 0xe5, 0x97, 0x52,
	0x68, 0x41, 0x3d, 0x5b, 0xe5, 0x37, 0x55, 0x7e, 0x5b, 0xe5, 0x57, 0xa1, 0x77, 0xaf, 0x26, 0xb0,
	0x92, 0x9b, 0x4c, 0x22, 0x24, 0x06, 0x17, 0x4c, 0x61, 0x9d, 0xf4, 0xee, 0xd7, 0xbb, 0x29, 0xd3,
	0x6c, 0x0d, 0xde, 0x32, 0xea, 0xc2, 0x61, 0xc5, 0x72, 0x9e, 0x32, 0x8d, 0xc1, 0xf5, 0x47, 0xbd,
	0x71, 0xe4, 0xc1, 0x68, 0xae, 0x25, 0xb2, 0xe5, 0xa9, 0x4d, 0xbc, 0x12, 0x99, 0x8a, 0x51, 0x95,
	0xa2, 0x50, 0x78, 0xf4, 0xa3, 0x0b, 0xc3, 0xf5, 0xcd, 0xd7, 0xa8, 0x14, 0xcb, 0x90, 0xbe, 0x05,
	0xe0, 0x29, 0x16, 0x9a, 0xbf, 0xe3, 0x28, 0x47, 0x64, 0x4c, 0x26, 0xfd, 0xf0, 0xa9, 0xbf, 0xbd,
	0x11, 0x7f, 0x0b, 0xc8, 0x7f, 0xd9, 0x52, 0xe2, 0xbf, 0x88, 0x34, 0x83, 0xde, 0x7b, 0xad, 0xcb,
	0xf3, 0x5c, 0x64, 0x6a, 0xe4, 0x58, 0xfc, 0x74, 0x17, 0xfc, 0x74, 0xb1, 0x98, 0xb5, 0xab, 0xcf,
	0x0b, 0x2d, 0x39, 0xaa, 0xe9, 0x5e, 0xec, 0x1a, 0xb8, 0xa9, 0xa3, 0x29, 0xb8, 0x3a, 0x69, 0x3c,
	0x1d, 0xeb, 0x79, 0xb1, 0x8b, 0x67, 0xf1, 0x6c, 0x93, 0xe6, 0x40, 0x27, 0xd6, 0xe2, 0x7d, 0x04,
	0xf8, 0xd3, 0x28, 0x7d, 0x0c, 0xfb, 0x85, 0x48, 0xb1, 0xb9, 0xb6, 0x61, 0xe3, 0x63, 0x25, 0x37,
	0x06, 0x33, 0x63, 0xff, 0x4c, 0xa4, 0x18, 0xc1, 0xaf, 0xab, 0xcb, 0x4e, 0xf7, 0x2b, 0x71, 0x06,
	0x24, 0xb6, 0x01, 0x7a, 0x0c, 0x6e, 0x2e, 0xb2, 0xf3, 0x82, 0x2d, 0xd1, 0x5e, 0x4a, 0x2f, 0xea,
	0x99, 0x9a, 0x7d, 0xe9, 0x8c, 0x49, 0x7c, 0x90, 0x8b, 0xec, 0x8c, 0x2d, 0xd1, 0xcb, 0xe1, 0x70,
	0x53, 0xdb, 0x74, 0x01, 0x3d, 0x93, 0xc6, 0x42, 0xcb, 0xd5, 0x88, 0x8c, 0x3b, 0x93, 0x7e, 0xf8,
	0xb0, 0x71, 0x9b, 0x3f, 0xe8, 0xdf, 0x46, 0x6f, 0x10, 0x56, 0xcd, 0x79, 0xbe, 0x11, 0xc7, 0x25,
	0xb1, 0x39, 0x87, 0x5d, 0xf5, 0x3e, 0xc0, 0xdd, 0x0d, 0xcd, 0xd3, 0xf9, 0x4d, 0xd9, 0x83, 0xad,
	0xb2, 0x75, 0xc0, 0x16, 0x57, 0x74, 0x08, 0xfd, 0x6b, 0xa8, 0x71, 0x74, 0x7f, 0x5e, 0x5d, 0x76,
	0x48, 0xf8, 0x9d, 0xc0, 0xa0, 0x8d, 0xcf, 0xeb, 0xa9, 0xd1, 0xcf, 0x04, 0x06, 0xeb, 0xc3, 0xa2,
	0x27, 0x3b, 0x8c, 0xd6, 0x7b, 0x74, 0x9b, 0x50, 0xfb, 0x78, 0xf6, 0x26, 0x24, 0x7a, 0x02, 0x13,
	0x2e, 0xea, 0x74, 0x29, 0xc5, 0xa7, 0xd5, 0x7f, 0x40, 0x91, 0x7b, 0x9a, 0xab, 0x99, 0x79, 0x92,
	0x33, 0xf2, 0xc6, 0xa9, 0xc2, 0x2f, 0x84, 0x5c, 0xdc, 0xb1, 0x4f, 0xf4, 0xe4, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x19, 0xfd, 0x86, 0xb5, 0x46, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccessLogServiceClient is the client API for AccessLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccessLogServiceClient interface {
	// Envoy will connect and send StreamAccessLogsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect. In the future we may decide to add a different
	// API for "critical" access logs in which Envoy will buffer access logs for some period of time
	// until it gets an ACK so it could then retry. This API is designed for high throughput with the
	// expectation that it might be lossy.
	StreamAccessLogs(ctx context.Context, opts ...grpc.CallOption) (AccessLogService_StreamAccessLogsClient, error)
}

type accessLogServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccessLogServiceClient(cc *grpc.ClientConn) AccessLogServiceClient {
	return &accessLogServiceClient{cc}
}

func (c *accessLogServiceClient) StreamAccessLogs(ctx context.Context, opts ...grpc.CallOption) (AccessLogService_StreamAccessLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AccessLogService_serviceDesc.Streams[0], "/envoy.service.accesslog.v2.AccessLogService/StreamAccessLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &accessLogServiceStreamAccessLogsClient{stream}
	return x, nil
}

type AccessLogService_StreamAccessLogsClient interface {
	Send(*StreamAccessLogsMessage) error
	CloseAndRecv() (*StreamAccessLogsResponse, error)
	grpc.ClientStream
}

type accessLogServiceStreamAccessLogsClient struct {
	grpc.ClientStream
}

func (x *accessLogServiceStreamAccessLogsClient) Send(m *StreamAccessLogsMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *accessLogServiceStreamAccessLogsClient) CloseAndRecv() (*StreamAccessLogsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamAccessLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccessLogServiceServer is the server API for AccessLogService service.
type AccessLogServiceServer interface {
	// Envoy will connect and send StreamAccessLogsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect. In the future we may decide to add a different
	// API for "critical" access logs in which Envoy will buffer access logs for some period of time
	// until it gets an ACK so it could then retry. This API is designed for high throughput with the
	// expectation that it might be lossy.
	StreamAccessLogs(AccessLogService_StreamAccessLogsServer) error
}

// UnimplementedAccessLogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccessLogServiceServer struct {
}

func (*UnimplementedAccessLogServiceServer) StreamAccessLogs(srv AccessLogService_StreamAccessLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAccessLogs not implemented")
}

func RegisterAccessLogServiceServer(s *grpc.Server, srv AccessLogServiceServer) {
	s.RegisterService(&_AccessLogService_serviceDesc, srv)
}

func _AccessLogService_StreamAccessLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccessLogServiceServer).StreamAccessLogs(&accessLogServiceStreamAccessLogsServer{stream})
}

type AccessLogService_StreamAccessLogsServer interface {
	SendAndClose(*StreamAccessLogsResponse) error
	Recv() (*StreamAccessLogsMessage, error)
	grpc.ServerStream
}

type accessLogServiceStreamAccessLogsServer struct {
	grpc.ServerStream
}

func (x *accessLogServiceStreamAccessLogsServer) SendAndClose(m *StreamAccessLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *accessLogServiceStreamAccessLogsServer) Recv() (*StreamAccessLogsMessage, error) {
	m := new(StreamAccessLogsMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AccessLogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.accesslog.v2.AccessLogService",
	HandlerType: (*AccessLogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAccessLogs",
			Handler:       _AccessLogService_StreamAccessLogs_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/accesslog/v2/als.proto",
}
