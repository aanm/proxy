// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/server_info.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type ServerInfo_State int32

const (
	// Server is live and serving traffic.
	ServerInfo_LIVE ServerInfo_State = 0
	// Server is draining listeners in response to external health checks failing.
	ServerInfo_DRAINING ServerInfo_State = 1
	// Server has not yet completed cluster manager initialization.
	ServerInfo_PRE_INITIALIZING ServerInfo_State = 2
	// Server is running the cluster manager initialization callbacks (e.g., RDS).
	ServerInfo_INITIALIZING ServerInfo_State = 3
)

var ServerInfo_State_name = map[int32]string{
	0: "LIVE",
	1: "DRAINING",
	2: "PRE_INITIALIZING",
	3: "INITIALIZING",
}

var ServerInfo_State_value = map[string]int32{
	"LIVE":             0,
	"DRAINING":         1,
	"PRE_INITIALIZING": 2,
	"INITIALIZING":     3,
}

func (x ServerInfo_State) String() string {
	return proto.EnumName(ServerInfo_State_name, int32(x))
}

func (ServerInfo_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ed0f406f9d75bf97, []int{0, 0}
}

type CommandLineOptions_IpVersion int32

const (
	CommandLineOptions_v4 CommandLineOptions_IpVersion = 0
	CommandLineOptions_v6 CommandLineOptions_IpVersion = 1
)

var CommandLineOptions_IpVersion_name = map[int32]string{
	0: "v4",
	1: "v6",
}

var CommandLineOptions_IpVersion_value = map[string]int32{
	"v4": 0,
	"v6": 1,
}

func (x CommandLineOptions_IpVersion) String() string {
	return proto.EnumName(CommandLineOptions_IpVersion_name, int32(x))
}

func (CommandLineOptions_IpVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ed0f406f9d75bf97, []int{1, 0}
}

type CommandLineOptions_Mode int32

const (
	// Validate configs and then serve traffic normally.
	CommandLineOptions_Serve CommandLineOptions_Mode = 0
	// Validate configs and exit.
	CommandLineOptions_Validate CommandLineOptions_Mode = 1
	// Completely load and initialize the config, and then exit without running the listener loop.
	CommandLineOptions_InitOnly CommandLineOptions_Mode = 2
)

var CommandLineOptions_Mode_name = map[int32]string{
	0: "Serve",
	1: "Validate",
	2: "InitOnly",
}

var CommandLineOptions_Mode_value = map[string]int32{
	"Serve":    0,
	"Validate": 1,
	"InitOnly": 2,
}

func (x CommandLineOptions_Mode) String() string {
	return proto.EnumName(CommandLineOptions_Mode_name, int32(x))
}

func (CommandLineOptions_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ed0f406f9d75bf97, []int{1, 1}
}

// Proto representation of the value returned by /server_info, containing
// server version/server status information.
type ServerInfo struct {
	// Server version.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// State of the server.
	State ServerInfo_State `protobuf:"varint,2,opt,name=state,proto3,enum=envoy.admin.v2alpha.ServerInfo_State" json:"state,omitempty"`
	// Uptime since current epoch was started.
	UptimeCurrentEpoch *duration.Duration `protobuf:"bytes,3,opt,name=uptime_current_epoch,json=uptimeCurrentEpoch,proto3" json:"uptime_current_epoch,omitempty"`
	// Uptime since the start of the first epoch.
	UptimeAllEpochs *duration.Duration `protobuf:"bytes,4,opt,name=uptime_all_epochs,json=uptimeAllEpochs,proto3" json:"uptime_all_epochs,omitempty"`
	// Command line options the server is currently running with.
	CommandLineOptions   *CommandLineOptions `protobuf:"bytes,6,opt,name=command_line_options,json=commandLineOptions,proto3" json:"command_line_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed0f406f9d75bf97, []int{0}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerInfo) GetState() ServerInfo_State {
	if m != nil {
		return m.State
	}
	return ServerInfo_LIVE
}

func (m *ServerInfo) GetUptimeCurrentEpoch() *duration.Duration {
	if m != nil {
		return m.UptimeCurrentEpoch
	}
	return nil
}

func (m *ServerInfo) GetUptimeAllEpochs() *duration.Duration {
	if m != nil {
		return m.UptimeAllEpochs
	}
	return nil
}

func (m *ServerInfo) GetCommandLineOptions() *CommandLineOptions {
	if m != nil {
		return m.CommandLineOptions
	}
	return nil
}

type CommandLineOptions struct {
	// See :option:`--base-id` for details.
	BaseId uint64 `protobuf:"varint,1,opt,name=base_id,json=baseId,proto3" json:"base_id,omitempty"`
	// See :option:`--concurrency` for details.
	Concurrency uint32 `protobuf:"varint,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	// See :option:`--config-path` for details.
	ConfigPath string `protobuf:"bytes,3,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`
	// See :option:`--config-yaml` for details.
	ConfigYaml string `protobuf:"bytes,4,opt,name=config_yaml,json=configYaml,proto3" json:"config_yaml,omitempty"`
	// See :option:`--allow-unknown-fields` for details.
	AllowUnknownFields bool `protobuf:"varint,5,opt,name=allow_unknown_fields,json=allowUnknownFields,proto3" json:"allow_unknown_fields,omitempty"`
	// See :option:`--admin-address-path` for details.
	AdminAddressPath string `protobuf:"bytes,6,opt,name=admin_address_path,json=adminAddressPath,proto3" json:"admin_address_path,omitempty"`
	// See :option:`--local-address-ip-version` for details.
	LocalAddressIpVersion CommandLineOptions_IpVersion `protobuf:"varint,7,opt,name=local_address_ip_version,json=localAddressIpVersion,proto3,enum=envoy.admin.v2alpha.CommandLineOptions_IpVersion" json:"local_address_ip_version,omitempty"`
	// See :option:`--log-level` for details.
	LogLevel string `protobuf:"bytes,8,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	// See :option:`--component-log-level` for details.
	ComponentLogLevel string `protobuf:"bytes,9,opt,name=component_log_level,json=componentLogLevel,proto3" json:"component_log_level,omitempty"`
	// See :option:`--log-format` for details.
	LogFormat string `protobuf:"bytes,10,opt,name=log_format,json=logFormat,proto3" json:"log_format,omitempty"`
	// See :option:`--log-path` for details.
	LogPath string `protobuf:"bytes,11,opt,name=log_path,json=logPath,proto3" json:"log_path,omitempty"`
	// See :option:`--hot-restart-version` for details.
	HotRestartVersion bool `protobuf:"varint,12,opt,name=hot_restart_version,json=hotRestartVersion,proto3" json:"hot_restart_version,omitempty"`
	// See :option:`--service-cluster` for details.
	ServiceCluster string `protobuf:"bytes,13,opt,name=service_cluster,json=serviceCluster,proto3" json:"service_cluster,omitempty"`
	// See :option:`--service-node` for details.
	ServiceNode string `protobuf:"bytes,14,opt,name=service_node,json=serviceNode,proto3" json:"service_node,omitempty"`
	// See :option:`--service-zone` for details.
	ServiceZone string `protobuf:"bytes,15,opt,name=service_zone,json=serviceZone,proto3" json:"service_zone,omitempty"`
	// See :option:`--file-flush-interval-msec` for details.
	FileFlushInterval *duration.Duration `protobuf:"bytes,16,opt,name=file_flush_interval,json=fileFlushInterval,proto3" json:"file_flush_interval,omitempty"`
	// See :option:`--drain-time-s` for details.
	DrainTime *duration.Duration `protobuf:"bytes,17,opt,name=drain_time,json=drainTime,proto3" json:"drain_time,omitempty"`
	// See :option:`--parent-shutdown-time-s` for details.
	ParentShutdownTime *duration.Duration `protobuf:"bytes,18,opt,name=parent_shutdown_time,json=parentShutdownTime,proto3" json:"parent_shutdown_time,omitempty"`
	// See :option:`--mode` for details.
	Mode CommandLineOptions_Mode `protobuf:"varint,19,opt,name=mode,proto3,enum=envoy.admin.v2alpha.CommandLineOptions_Mode" json:"mode,omitempty"`
	// See :option:`--max-stats` for details.
	MaxStats uint64 `protobuf:"varint,20,opt,name=max_stats,json=maxStats,proto3" json:"max_stats,omitempty"`
	// See :option:`--max-obj-name-len` for details.
	MaxObjNameLen uint64 `protobuf:"varint,21,opt,name=max_obj_name_len,json=maxObjNameLen,proto3" json:"max_obj_name_len,omitempty"`
	// See :option:`--disable-hot-restart` for details.
	DisableHotRestart bool `protobuf:"varint,22,opt,name=disable_hot_restart,json=disableHotRestart,proto3" json:"disable_hot_restart,omitempty"`
	// See :option:`--enable-mutex-tracing` for details.
	EnableMutexTracing bool `protobuf:"varint,23,opt,name=enable_mutex_tracing,json=enableMutexTracing,proto3" json:"enable_mutex_tracing,omitempty"`
	// See :option:`--restart-epoch` for details.
	RestartEpoch uint32 `protobuf:"varint,24,opt,name=restart_epoch,json=restartEpoch,proto3" json:"restart_epoch,omitempty"`
	// See :option:`--cpuset-threads` for details.
	CpusetThreads        bool     `protobuf:"varint,25,opt,name=cpuset_threads,json=cpusetThreads,proto3" json:"cpuset_threads,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandLineOptions) Reset()         { *m = CommandLineOptions{} }
func (m *CommandLineOptions) String() string { return proto.CompactTextString(m) }
func (*CommandLineOptions) ProtoMessage()    {}
func (*CommandLineOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed0f406f9d75bf97, []int{1}
}

func (m *CommandLineOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandLineOptions.Unmarshal(m, b)
}
func (m *CommandLineOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandLineOptions.Marshal(b, m, deterministic)
}
func (m *CommandLineOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandLineOptions.Merge(m, src)
}
func (m *CommandLineOptions) XXX_Size() int {
	return xxx_messageInfo_CommandLineOptions.Size(m)
}
func (m *CommandLineOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandLineOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CommandLineOptions proto.InternalMessageInfo

func (m *CommandLineOptions) GetBaseId() uint64 {
	if m != nil {
		return m.BaseId
	}
	return 0
}

func (m *CommandLineOptions) GetConcurrency() uint32 {
	if m != nil {
		return m.Concurrency
	}
	return 0
}

func (m *CommandLineOptions) GetConfigPath() string {
	if m != nil {
		return m.ConfigPath
	}
	return ""
}

func (m *CommandLineOptions) GetConfigYaml() string {
	if m != nil {
		return m.ConfigYaml
	}
	return ""
}

func (m *CommandLineOptions) GetAllowUnknownFields() bool {
	if m != nil {
		return m.AllowUnknownFields
	}
	return false
}

func (m *CommandLineOptions) GetAdminAddressPath() string {
	if m != nil {
		return m.AdminAddressPath
	}
	return ""
}

func (m *CommandLineOptions) GetLocalAddressIpVersion() CommandLineOptions_IpVersion {
	if m != nil {
		return m.LocalAddressIpVersion
	}
	return CommandLineOptions_v4
}

func (m *CommandLineOptions) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *CommandLineOptions) GetComponentLogLevel() string {
	if m != nil {
		return m.ComponentLogLevel
	}
	return ""
}

func (m *CommandLineOptions) GetLogFormat() string {
	if m != nil {
		return m.LogFormat
	}
	return ""
}

func (m *CommandLineOptions) GetLogPath() string {
	if m != nil {
		return m.LogPath
	}
	return ""
}

func (m *CommandLineOptions) GetHotRestartVersion() bool {
	if m != nil {
		return m.HotRestartVersion
	}
	return false
}

func (m *CommandLineOptions) GetServiceCluster() string {
	if m != nil {
		return m.ServiceCluster
	}
	return ""
}

func (m *CommandLineOptions) GetServiceNode() string {
	if m != nil {
		return m.ServiceNode
	}
	return ""
}

func (m *CommandLineOptions) GetServiceZone() string {
	if m != nil {
		return m.ServiceZone
	}
	return ""
}

func (m *CommandLineOptions) GetFileFlushInterval() *duration.Duration {
	if m != nil {
		return m.FileFlushInterval
	}
	return nil
}

func (m *CommandLineOptions) GetDrainTime() *duration.Duration {
	if m != nil {
		return m.DrainTime
	}
	return nil
}

func (m *CommandLineOptions) GetParentShutdownTime() *duration.Duration {
	if m != nil {
		return m.ParentShutdownTime
	}
	return nil
}

func (m *CommandLineOptions) GetMode() CommandLineOptions_Mode {
	if m != nil {
		return m.Mode
	}
	return CommandLineOptions_Serve
}

func (m *CommandLineOptions) GetMaxStats() uint64 {
	if m != nil {
		return m.MaxStats
	}
	return 0
}

func (m *CommandLineOptions) GetMaxObjNameLen() uint64 {
	if m != nil {
		return m.MaxObjNameLen
	}
	return 0
}

func (m *CommandLineOptions) GetDisableHotRestart() bool {
	if m != nil {
		return m.DisableHotRestart
	}
	return false
}

func (m *CommandLineOptions) GetEnableMutexTracing() bool {
	if m != nil {
		return m.EnableMutexTracing
	}
	return false
}

func (m *CommandLineOptions) GetRestartEpoch() uint32 {
	if m != nil {
		return m.RestartEpoch
	}
	return 0
}

func (m *CommandLineOptions) GetCpusetThreads() bool {
	if m != nil {
		return m.CpusetThreads
	}
	return false
}

func init() {
	proto.RegisterEnum("envoy.admin.v2alpha.ServerInfo_State", ServerInfo_State_name, ServerInfo_State_value)
	proto.RegisterEnum("envoy.admin.v2alpha.CommandLineOptions_IpVersion", CommandLineOptions_IpVersion_name, CommandLineOptions_IpVersion_value)
	proto.RegisterEnum("envoy.admin.v2alpha.CommandLineOptions_Mode", CommandLineOptions_Mode_name, CommandLineOptions_Mode_value)
	proto.RegisterType((*ServerInfo)(nil), "envoy.admin.v2alpha.ServerInfo")
	proto.RegisterType((*CommandLineOptions)(nil), "envoy.admin.v2alpha.CommandLineOptions")
}

func init() {
	proto.RegisterFile("envoy/admin/v2alpha/server_info.proto", fileDescriptor_ed0f406f9d75bf97)
}

var fileDescriptor_ed0f406f9d75bf97 = []byte{
	// 912 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0xad, 0xf3, 0x69, 0xdf, 0x7c, 0x29, 0x4c, 0xba, 0x32, 0x2b, 0xb6, 0xa5, 0x1e, 0x82, 0xe4,
	0xa1, 0x93, 0xb7, 0x6c, 0x28, 0x06, 0xec, 0x65, 0x69, 0x9a, 0x74, 0xc2, 0x5c, 0x27, 0x50, 0xb2,
	0x00, 0xed, 0x0b, 0x41, 0x4b, 0xb4, 0xcd, 0x8c, 0x22, 0x05, 0x89, 0x72, 0xed, 0xfd, 0x95, 0xfd,
	0xb0, 0xfd, 0x9d, 0x81, 0x97, 0xb2, 0x93, 0xa2, 0x01, 0xd2, 0x27, 0x43, 0xe7, 0x9e, 0x73, 0x48,
	0xde, 0x2f, 0xc3, 0x81, 0xd0, 0x63, 0x33, 0xed, 0xf0, 0x34, 0x93, 0xba, 0x33, 0x3e, 0xe6, 0x2a,
	0x1f, 0xf1, 0x4e, 0x29, 0x8a, 0xb1, 0x28, 0x98, 0xd4, 0x03, 0x13, 0xe6, 0x85, 0xb1, 0x86, 0xec,
	0x20, 0x2d, 0x44, 0x5a, 0x58, 0xd3, 0xbe, 0xfe, 0x76, 0x68, 0xcc, 0x50, 0x89, 0x0e, 0x52, 0xfa,
	0xd5, 0xa0, 0x93, 0x56, 0x05, 0xb7, 0xd2, 0x68, 0x2f, 0x6a, 0xff, 0xbb, 0x08, 0x70, 0x85, 0x56,
	0x91, 0x1e, 0x18, 0x42, 0x61, 0x75, 0x2c, 0x8a, 0x52, 0x1a, 0x4d, 0x1b, 0xfb, 0x8d, 0xa3, 0x56,
	0x3c, 0xfb, 0x24, 0xbf, 0xc1, 0x72, 0x69, 0xb9, 0x15, 0x74, 0x61, 0xbf, 0x71, 0xb4, 0x79, 0x7c,
	0x10, 0x3e, 0x70, 0x5a, 0x78, 0xe7, 0x14, 0x5e, 0x39, 0x72, 0xec, 0x35, 0xe4, 0x4f, 0xd8, 0xad,
	0x72, 0x2b, 0x33, 0xc1, 0x92, 0xaa, 0x28, 0x84, 0xb6, 0x4c, 0xe4, 0x26, 0x19, 0xd1, 0xc5, 0xfd,
	0xc6, 0xd1, 0xda, 0xf1, 0x5e, 0xe8, 0x2f, 0x19, 0xce, 0x2e, 0x19, 0xbe, 0xa9, 0x2f, 0x19, 0x13,
	0x2f, 0x3b, 0xf5, 0xaa, 0x33, 0x27, 0x22, 0x67, 0xb0, 0x5d, 0x9b, 0x71, 0xa5, 0xbc, 0x51, 0x49,
	0x97, 0x1e, 0x73, 0xda, 0xf2, 0x9a, 0x13, 0xa5, 0xd0, 0xa5, 0x24, 0xef, 0x61, 0x37, 0x31, 0x59,
	0xc6, 0x75, 0xca, 0x94, 0xd4, 0x82, 0x99, 0xdc, 0xf1, 0x4a, 0xba, 0x82, 0x4e, 0x87, 0x0f, 0xbe,
	0xef, 0xd4, 0x0b, 0xba, 0x52, 0x8b, 0x0b, 0x4f, 0x8f, 0x49, 0xf2, 0x19, 0xd6, 0x7e, 0x0b, 0xcb,
	0xf8, 0x7c, 0xd2, 0x84, 0xa5, 0x6e, 0x74, 0x73, 0x16, 0x3c, 0x21, 0xeb, 0xd0, 0x7c, 0x13, 0x9f,
	0x44, 0xbd, 0xa8, 0xf7, 0x36, 0x68, 0x90, 0x5d, 0x08, 0x2e, 0xe3, 0x33, 0x16, 0xf5, 0xa2, 0xeb,
	0xe8, 0xa4, 0x1b, 0x7d, 0x70, 0xe8, 0x02, 0x09, 0x60, 0xfd, 0x13, 0x64, 0xb1, 0xfd, 0x5f, 0x0b,
	0xc8, 0xe7, 0x67, 0x92, 0x67, 0xb0, 0xda, 0xe7, 0xa5, 0x60, 0x32, 0xc5, 0x2a, 0x2d, 0xc5, 0x2b,
	0xee, 0x33, 0x4a, 0xc9, 0x3e, 0xac, 0x25, 0x46, 0xfb, 0x1c, 0x27, 0x53, 0x2c, 0xd5, 0x46, 0x7c,
	0x1f, 0x22, 0xdf, 0x21, 0x63, 0x20, 0x87, 0x2c, 0xe7, 0xd6, 0x17, 0xa0, 0x15, 0x83, 0x87, 0x2e,
	0xb9, 0x1d, 0xdd, 0x23, 0x4c, 0x79, 0xa6, 0x30, 0xaf, 0x73, 0xc2, 0x7b, 0x9e, 0x29, 0xf2, 0x23,
	0xec, 0x72, 0xa5, 0xcc, 0x47, 0x56, 0xe9, 0xbf, 0xb5, 0xf9, 0xa8, 0xd9, 0x40, 0x0a, 0x95, 0x96,
	0x74, 0x79, 0xbf, 0x71, 0xd4, 0x8c, 0x09, 0xc6, 0xfe, 0xf2, 0xa1, 0x73, 0x8c, 0x90, 0x97, 0x40,
	0x30, 0x8d, 0x8c, 0xa7, 0x69, 0x21, 0xca, 0xd2, 0x1f, 0xbd, 0x82, 0xce, 0x01, 0x46, 0x4e, 0x7c,
	0x00, 0x2f, 0x70, 0x0b, 0x54, 0x99, 0x84, 0xab, 0x39, 0x5b, 0xe6, 0x6c, 0xd6, 0x93, 0xab, 0xd8,
	0x7b, 0x3f, 0x7d, 0x61, 0x6d, 0xc2, 0x28, 0xbf, 0xf1, 0xc2, 0xf8, 0x29, 0x5a, 0xd6, 0xc7, 0xcc,
	0x61, 0xf2, 0x1c, 0x5a, 0xca, 0x0c, 0x99, 0x12, 0x63, 0xa1, 0x68, 0x13, 0x2f, 0xd4, 0x54, 0x66,
	0xd8, 0x75, 0xdf, 0x24, 0x84, 0x9d, 0xc4, 0x64, 0xb9, 0xd1, 0xae, 0x5f, 0xef, 0x68, 0x2d, 0xa4,
	0x6d, 0xcf, 0x43, 0xdd, 0x19, 0xff, 0x1b, 0x00, 0xc7, 0x1a, 0x98, 0x22, 0xe3, 0x96, 0x02, 0xd2,
	0x9c, 0xfd, 0x39, 0x02, 0x64, 0x0f, 0x9c, 0xb5, 0x7f, 0xfb, 0x9a, 0x9f, 0x2d, 0x65, 0x7c, 0xce,
	0x43, 0xd8, 0x19, 0x19, 0xcb, 0x0a, 0x51, 0x5a, 0x5e, 0xd8, 0xf9, 0x6b, 0xd7, 0x31, 0xa3, 0xdb,
	0x23, 0x63, 0x63, 0x1f, 0x99, 0x5d, 0xfb, 0x10, 0xb6, 0xdc, 0xf8, 0xcb, 0x44, 0xb0, 0x44, 0x55,
	0xa5, 0x15, 0x05, 0xdd, 0x40, 0xc7, 0xcd, 0x1a, 0x3e, 0xf5, 0x28, 0x79, 0x01, 0xeb, 0x33, 0xa2,
	0x36, 0xa9, 0xa0, 0x9b, 0xc8, 0x5a, 0xab, 0xb1, 0x9e, 0x49, 0xc5, 0x7d, 0xca, 0x3f, 0x46, 0x0b,
	0xba, 0xf5, 0x09, 0xe5, 0x83, 0xd1, 0x82, 0x44, 0xb0, 0x33, 0x90, 0x4a, 0xb0, 0x81, 0xaa, 0xca,
	0x11, 0x93, 0xda, 0x8a, 0x62, 0xcc, 0x15, 0x0d, 0x1e, 0x1b, 0xb9, 0x6d, 0xa7, 0x3a, 0x77, 0xa2,
	0xa8, 0xd6, 0x90, 0x5f, 0x01, 0xd2, 0x82, 0x4b, 0xcd, 0xdc, 0x2c, 0xd2, 0xed, 0xc7, 0x1c, 0x5a,
	0x48, 0xbe, 0x96, 0x19, 0xae, 0x90, 0x9c, 0xe3, 0xea, 0x28, 0x47, 0x95, 0x4d, 0x5d, 0xe3, 0xa1,
	0x07, 0x79, 0x74, 0x85, 0x78, 0xd9, 0x55, 0xad, 0x42, 0xb3, 0xdf, 0x61, 0x29, 0x73, 0xf9, 0xd8,
	0xc1, 0x7e, 0x7a, 0xf9, 0xa5, 0xfd, 0xf4, 0xce, 0xa4, 0x22, 0x46, 0xa5, 0xeb, 0x9c, 0x8c, 0x4f,
	0x98, 0x5b, 0x6f, 0x25, 0xdd, 0xc5, 0x21, 0x6c, 0x66, 0x7c, 0xe2, 0xc6, 0xbe, 0x24, 0x87, 0x10,
	0xb8, 0xa0, 0xe9, 0xdf, 0x32, 0xcd, 0x33, 0xc1, 0x94, 0xd0, 0xf4, 0x29, 0x72, 0x36, 0x32, 0x3e,
	0xb9, 0xe8, 0xdf, 0xf6, 0x78, 0x26, 0xba, 0x42, 0xbb, 0xc2, 0xa7, 0xb2, 0xe4, 0x7d, 0x25, 0xd8,
	0xbd, 0x06, 0xa0, 0x5f, 0xf9, 0xc2, 0xd7, 0xa1, 0x3f, 0xe6, 0xf5, 0x77, 0xb3, 0x27, 0x34, 0xd2,
	0xb3, 0xca, 0x8a, 0x09, 0xb3, 0x05, 0x4f, 0xa4, 0x1e, 0xd2, 0x67, 0x7e, 0xf6, 0x7c, 0xec, 0x9d,
	0x0b, 0x5d, 0xfb, 0x08, 0xf9, 0x1e, 0x36, 0x66, 0x6d, 0xe5, 0x57, 0x2e, 0xc5, 0x9d, 0xb0, 0x5e,
	0x83, 0x7e, 0xa3, 0x1e, 0xc0, 0x66, 0x92, 0x57, 0xa5, 0xb0, 0xcc, 0x8e, 0x0a, 0xc1, 0xd3, 0x92,
	0xee, 0xa1, 0xe1, 0x86, 0x47, 0xaf, 0x3d, 0xd8, 0x7e, 0x0e, 0xad, 0xbb, 0xd1, 0x59, 0x81, 0x85,
	0xf1, 0x2f, 0xc1, 0x13, 0xfc, 0x7d, 0x15, 0x34, 0xda, 0x3f, 0xc0, 0x92, 0x4b, 0x0f, 0x69, 0xc1,
	0x32, 0xfe, 0x0b, 0xf8, 0x9d, 0x77, 0xc3, 0x95, 0x4c, 0xb9, 0x15, 0x41, 0xc3, 0x7d, 0x45, 0x5a,
	0xda, 0x0b, 0xad, 0xa6, 0xc1, 0xc2, 0xeb, 0x57, 0xf0, 0x42, 0x1a, 0x9f, 0xf7, 0xbc, 0x30, 0x93,
	0xe9, 0x43, 0x25, 0x78, 0xbd, 0x75, 0xf7, 0x7f, 0x72, 0xe9, 0xea, 0x7a, 0xd9, 0xe8, 0xaf, 0x60,
	0x81, 0x7f, 0xfe, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x37, 0xe0, 0x41, 0x14, 0x07, 0x00, 0x00,
}
