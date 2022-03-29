// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nexus-config.proto

package config

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

type NexusConfig struct {
	Version              string    `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Cloud                *Cloud    `protobuf:"bytes,2,opt,name=cloud,proto3" json:"cloud,omitempty"`
	Platform             *Platform `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *NexusConfig) Reset()         { *m = NexusConfig{} }
func (m *NexusConfig) String() string { return proto.CompactTextString(m) }
func (*NexusConfig) ProtoMessage()    {}
func (*NexusConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_nexus_config_267e1c975478cf7f, []int{0}
}
func (m *NexusConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NexusConfig.Unmarshal(m, b)
}
func (m *NexusConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NexusConfig.Marshal(b, m, deterministic)
}
func (dst *NexusConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NexusConfig.Merge(dst, src)
}
func (m *NexusConfig) XXX_Size() int {
	return xxx_messageInfo_NexusConfig.Size(m)
}
func (m *NexusConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NexusConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NexusConfig proto.InternalMessageInfo

func (m *NexusConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *NexusConfig) GetCloud() *Cloud {
	if m != nil {
		return m.Cloud
	}
	return nil
}

func (m *NexusConfig) GetPlatform() *Platform {
	if m != nil {
		return m.Platform
	}
	return nil
}

type Cloud struct {
	Provider             string             `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Region               string             `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	Zone                 string             `protobuf:"bytes,3,opt,name=zone,proto3" json:"zone,omitempty"`
	Credentials          *Cloud_Credentials `protobuf:"bytes,4,opt,name=credentials,proto3" json:"credentials,omitempty"`
	StreamName           string             `protobuf:"bytes,5,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Cloud) Reset()         { *m = Cloud{} }
func (m *Cloud) String() string { return proto.CompactTextString(m) }
func (*Cloud) ProtoMessage()    {}
func (*Cloud) Descriptor() ([]byte, []int) {
	return fileDescriptor_nexus_config_267e1c975478cf7f, []int{1}
}
func (m *Cloud) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cloud.Unmarshal(m, b)
}
func (m *Cloud) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cloud.Marshal(b, m, deterministic)
}
func (dst *Cloud) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cloud.Merge(dst, src)
}
func (m *Cloud) XXX_Size() int {
	return xxx_messageInfo_Cloud.Size(m)
}
func (m *Cloud) XXX_DiscardUnknown() {
	xxx_messageInfo_Cloud.DiscardUnknown(m)
}

var xxx_messageInfo_Cloud proto.InternalMessageInfo

func (m *Cloud) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Cloud) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Cloud) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *Cloud) GetCredentials() *Cloud_Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *Cloud) GetStreamName() string {
	if m != nil {
		return m.StreamName
	}
	return ""
}

type Cloud_Credentials struct {
	AccessKey            string   `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey            string   `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	AccountId            string   `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cloud_Credentials) Reset()         { *m = Cloud_Credentials{} }
func (m *Cloud_Credentials) String() string { return proto.CompactTextString(m) }
func (*Cloud_Credentials) ProtoMessage()    {}
func (*Cloud_Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_nexus_config_267e1c975478cf7f, []int{1, 0}
}
func (m *Cloud_Credentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cloud_Credentials.Unmarshal(m, b)
}
func (m *Cloud_Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cloud_Credentials.Marshal(b, m, deterministic)
}
func (dst *Cloud_Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cloud_Credentials.Merge(dst, src)
}
func (m *Cloud_Credentials) XXX_Size() int {
	return xxx_messageInfo_Cloud_Credentials.Size(m)
}
func (m *Cloud_Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Cloud_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Cloud_Credentials proto.InternalMessageInfo

func (m *Cloud_Credentials) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *Cloud_Credentials) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *Cloud_Credentials) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type Platform struct {
	Database             *Platform_Database       `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	MessageBus           *Platform_MessageBusInfo `protobuf:"bytes,2,opt,name=message_bus,json=messageBus,proto3" json:"message_bus,omitempty"`
	Health               *Platform_Health         `protobuf:"bytes,3,opt,name=health,proto3" json:"health,omitempty"`
	Sidecar              *Platform_Sidecar        `protobuf:"bytes,4,opt,name=sidecar,proto3" json:"sidecar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Platform) Reset()         { *m = Platform{} }
func (m *Platform) String() string { return proto.CompactTextString(m) }
func (*Platform) ProtoMessage()    {}
func (*Platform) Descriptor() ([]byte, []int) {
	return fileDescriptor_nexus_config_267e1c975478cf7f, []int{2}
}
func (m *Platform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform.Unmarshal(m, b)
}
func (m *Platform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform.Marshal(b, m, deterministic)
}
func (dst *Platform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform.Merge(dst, src)
}
func (m *Platform) XXX_Size() int {
	return xxx_messageInfo_Platform.Size(m)
}
func (m *Platform) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform.DiscardUnknown(m)
}

var xxx_messageInfo_Platform proto.InternalMessageInfo

func (m *Platform) GetDatabase() *Platform_Database {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *Platform) GetMessageBus() *Platform_MessageBusInfo {
	if m != nil {
		return m.MessageBus
	}
	return nil
}

func (m *Platform) GetHealth() *Platform_Health {
	if m != nil {
		return m.Health
	}
	return nil
}

func (m *Platform) GetSidecar() *Platform_Sidecar {
	if m != nil {
		return m.Sidecar
	}
	return nil
}

type Platform_HealthInfo struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Platform_HealthInfo) Reset()         { *m = Platform_HealthInfo{} }
func (m *Platform_HealthInfo) String() string { return proto.CompactTextString(m) }
func (*Platform_HealthInfo) ProtoMessage()    {}
func (*Platform_HealthInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_nexus_config_267e1c975478cf7f, []int{2, 0}
}
func (m *Platform_HealthInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform_HealthInfo.Unmarshal(m, b)
}
func (m *Platform_HealthInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform_HealthInfo.Marshal(b, m, deterministic)
}
func (dst *Platform_HealthInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform_HealthInfo.Merge(dst, src)
}
func (m *Platform_HealthInfo) XXX_Size() int {
	return xxx_messageInfo_Platform_HealthInfo.Size(m)
}
func (m *Platform_HealthInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform_HealthInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Platform_HealthInfo proto.InternalMessageInfo

func (m *Platform_HealthInfo) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Platform_SidecarInfo struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Platform_SidecarInfo) Reset()         { *m = Platform_SidecarInfo{} }
func (m *Platform_SidecarInfo) String() string { return proto.CompactTextString(m) }
func (*Platform_SidecarInfo) ProtoMessage()    {}
func (*Platform_SidecarInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_nexus_config_267e1c975478cf7f, []int{2, 1}
}
func (m *Platform_SidecarInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform_SidecarInfo.Unmarshal(m, b)
}
func (m *Platform_SidecarInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform_SidecarInfo.Marshal(b, m, deterministic)
}
func (dst *Platform_SidecarInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform_SidecarInfo.Merge(dst, src)
}
func (m *Platform_SidecarInfo) XXX_Size() int {
	return xxx_messageInfo_Platform_SidecarInfo.Size(m)
}
func (m *Platform_SidecarInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform_SidecarInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Platform_SidecarInfo proto.InternalMessageInfo

func (m *Platform_SidecarInfo) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Platform_DatabaseInfo struct {
	// IP or host
	Host                 string                                   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 uint32                                   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	DatabaseType         string                                   `protobuf:"bytes,3,opt,name=database_type,json=databaseType,proto3" json:"database_type,omitempty"`
	Location             string                                   `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Credentials          *Platform_DatabaseInfo_Credentials       `protobuf:"bytes,5,opt,name=credentials,proto3" json:"credentials,omitempty"`
	Name                 string                                   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Capacity             string                                   `protobuf:"bytes,7,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Retention            string                                   `protobuf:"bytes,8,opt,name=retention,proto3" json:"retention,omitempty"`
	HttpPort             uint32                                   `protobuf:"varint,9,opt,name=http_port,json=httpPort,proto3" json:"http_port,omitempty"`
	RequestsMemory       string                                   `protobuf:"bytes,10,opt,name=requests_memory,json=requestsMemory,proto3" json:"requests_memory,omitempty"`
	DMFeatureFlags       string                                   `protobuf:"bytes,11,opt,name=dm_feature_flags,json=dmFeatureFlags,proto3" json:"dm_feature_flags,omitempty"`
	EtcdManagerConfig    *Platform_DatabaseInfo_EtcdManagerConfig `protobuf:"bytes,12,opt,name=etcd_manager_config,json=etcdManagerConfig,proto3" json:"etcd_manager_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *Platform_DatabaseInfo) Reset()         { *m = Platform_DatabaseInfo{} }
func (m *Platform_DatabaseInfo) String() string { return proto.CompactTextString(m) }
func (*Platform_DatabaseInfo) ProtoMessage()    {}
func (*Platform_DatabaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_nexus_config_267e1c975478cf7f, []int{2, 2}
}
func (m *Platform_DatabaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform_DatabaseInfo.Unmarshal(m, b)
}
func (m *Platform_DatabaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform_DatabaseInfo.Marshal(b, m, deterministic)
}
func (dst *Platform_DatabaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform_DatabaseInfo.Merge(dst, src)
}
func (m *Platform_DatabaseInfo) XXX_Size() int {
	return xxx_messageInfo_Platform_DatabaseInfo.Size(m)
}
func (m *Platform_DatabaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform_DatabaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Platform_DatabaseInfo proto.InternalMessageInfo

func (m *Platform_DatabaseInfo) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Platform_DatabaseInfo) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Platform_DatabaseInfo) GetDatabaseType() string {
	if m != nil {
		return m.DatabaseType
	}
	return ""
}

func (m *Platform_DatabaseInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Platform_DatabaseInfo) GetCredentials() *Platform_DatabaseInfo_Credentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *Platform_DatabaseInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Platform_DatabaseInfo) GetCapacity() string {
	if m != nil {
		return m.Capacity
	}
	return ""
}

func (m *Platform_DatabaseInfo) GetRetention() string {
	if m != nil {
		return m.Retention
	}
	return ""
}

func (m *Platform_DatabaseInfo) GetHttpPort() uint32 {
	if m != nil {
		return m.HttpPort
	}
	return 0
}

func (m *Platform_DatabaseInfo) GetRequestsMemory() string {
	if m != nil {
		return m.RequestsMemory
	}
	return ""
}

func (m *Platform_DatabaseInfo) GetDMFeatureFlags() string {
	if m != nil {
		return m.DMFeatureFlags
	}
	return ""
}

func (m *Platform_DatabaseInfo) GetEtcdManagerConfig() *Platform_DatabaseInfo_EtcdManagerConfig {
	if m != nil {
		return m.EtcdManagerConfig
	}
	return nil
}

type Platform_DatabaseInfo_Credentials struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Platform_DatabaseInfo_Credentials) Reset()         { *m = Platform_DatabaseInfo_Credentials{} }
func (m *Platform_DatabaseInfo_Credentials) String() string { return proto.CompactTextString(m) }
func (*Platform_DatabaseInfo_Credentials) ProtoMessage()    {}
func (*Platform_DatabaseInfo_Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_nexus_config_267e1c975478cf7f, []int{2, 2, 0}
}
func (m *Platform_DatabaseInfo_Credentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform_DatabaseInfo_Credentials.Unmarshal(m, b)
}
func (m *Platform_DatabaseInfo_Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform_DatabaseInfo_Credentials.Marshal(b, m, deterministic)
}
func (dst *Platform_DatabaseInfo_Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform_DatabaseInfo_Credentials.Merge(dst, src)
}
func (m *Platform_DatabaseInfo_Credentials) XXX_Size() int {
	return xxx_messageInfo_Platform_DatabaseInfo_Credentials.Size(m)
}
func (m *Platform_DatabaseInfo_Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform_DatabaseInfo_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Platform_DatabaseInfo_Credentials proto.InternalMessageInfo

func (m *Platform_DatabaseInfo_Credentials) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Platform_DatabaseInfo_Credentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Platform_DatabaseInfo_EtcdManagerConfig struct {
	CompactRevisionInterval string   `protobuf:"bytes,1,opt,name=compact_revision_interval,json=compactRevisionInterval,proto3" json:"compact_revision_interval,omitempty"`
	CompactRetryTimer       string   `protobuf:"bytes,2,opt,name=compact_retry_timer,json=compactRetryTimer,proto3" json:"compact_retry_timer,omitempty"`
	CompactPhysical         string   `protobuf:"bytes,3,opt,name=compact_physical,json=compactPhysical,proto3" json:"compact_physical,omitempty"`
	CompactContextTimeout   string   `protobuf:"bytes,4,opt,name=compact_context_timeout,json=compactContextTimeout,proto3" json:"compact_context_timeout,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Platform_DatabaseInfo_EtcdManagerConfig) Reset() {
	*m = Platform_DatabaseInfo_EtcdManagerConfig{}
}
func (m *Platform_DatabaseInfo_EtcdManagerConfig) String() string { return proto.CompactTextString(m) }
func (*Platform_DatabaseInfo_EtcdManagerConfig) ProtoMessage()    {}
func (*Platform_DatabaseInfo_EtcdManagerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_nexus_config_267e1c975478cf7f, []int{2, 2, 1}
}
func (m *Platform_DatabaseInfo_EtcdManagerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform_DatabaseInfo_EtcdManagerConfig.Unmarshal(m, b)
}
func (m *Platform_DatabaseInfo_EtcdManagerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform_DatabaseInfo_EtcdManagerConfig.Marshal(b, m, deterministic)
}
func (dst *Platform_DatabaseInfo_EtcdManagerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform_DatabaseInfo_EtcdManagerConfig.Merge(dst, src)
}
func (m *Platform_DatabaseInfo_EtcdManagerConfig) XXX_Size() int {
	return xxx_messageInfo_Platform_DatabaseInfo_EtcdManagerConfig.Size(m)
}
func (m *Platform_DatabaseInfo_EtcdManagerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform_DatabaseInfo_EtcdManagerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Platform_DatabaseInfo_EtcdManagerConfig proto.InternalMessageInfo

func (m *Platform_DatabaseInfo_EtcdManagerConfig) GetCompactRevisionInterval() string {
	if m != nil {
		return m.CompactRevisionInterval
	}
	return ""
}

func (m *Platform_DatabaseInfo_EtcdManagerConfig) GetCompactRetryTimer() string {
	if m != nil {
		return m.CompactRetryTimer
	}
	return ""
}

func (m *Platform_DatabaseInfo_EtcdManagerConfig) GetCompactPhysical() string {
	if m != nil {
		return m.CompactPhysical
	}
	return ""
}

func (m *Platform_DatabaseInfo_EtcdManagerConfig) GetCompactContextTimeout() string {
	if m != nil {
		return m.CompactContextTimeout
	}
	return ""
}

type Platform_Health struct {
	Server               *Platform_HealthInfo `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Platform_Health) Reset()         { *m = Platform_Health{} }
func (m *Platform_Health) String() string { return proto.CompactTextString(m) }
func (*Platform_Health) ProtoMessage()    {}
func (*Platform_Health) Descriptor() ([]byte, []int) {
	return fileDescriptor_nexus_config_267e1c975478cf7f, []int{2, 3}
}
func (m *Platform_Health) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform_Health.Unmarshal(m, b)
}
func (m *Platform_Health) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform_Health.Marshal(b, m, deterministic)
}
func (dst *Platform_Health) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform_Health.Merge(dst, src)
}
func (m *Platform_Health) XXX_Size() int {
	return xxx_messageInfo_Platform_Health.Size(m)
}
func (m *Platform_Health) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform_Health.DiscardUnknown(m)
}

var xxx_messageInfo_Platform_Health proto.InternalMessageInfo

func (m *Platform_Health) GetServer() *Platform_HealthInfo {
	if m != nil {
		return m.Server
	}
	return nil
}

type Platform_Database struct {
	Graph                *Platform_DatabaseInfo `protobuf:"bytes,1,opt,name=graph,proto3" json:"graph,omitempty"`
	Metrics              *Platform_DatabaseInfo `protobuf:"bytes,2,opt,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Platform_Database) Reset()         { *m = Platform_Database{} }
func (m *Platform_Database) String() string { return proto.CompactTextString(m) }
func (*Platform_Database) ProtoMessage()    {}
func (*Platform_Database) Descriptor() ([]byte, []int) {
	return fileDescriptor_nexus_config_267e1c975478cf7f, []int{2, 4}
}
func (m *Platform_Database) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform_Database.Unmarshal(m, b)
}
func (m *Platform_Database) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform_Database.Marshal(b, m, deterministic)
}
func (dst *Platform_Database) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform_Database.Merge(dst, src)
}
func (m *Platform_Database) XXX_Size() int {
	return xxx_messageInfo_Platform_Database.Size(m)
}
func (m *Platform_Database) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform_Database.DiscardUnknown(m)
}

var xxx_messageInfo_Platform_Database proto.InternalMessageInfo

func (m *Platform_Database) GetGraph() *Platform_DatabaseInfo {
	if m != nil {
		return m.Graph
	}
	return nil
}

func (m *Platform_Database) GetMetrics() *Platform_DatabaseInfo {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type Platform_MessageBusInfo struct {
	Provider             string   `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	PrefixName           string   `protobuf:"bytes,3,opt,name=prefix_name,json=prefixName,proto3" json:"prefix_name,omitempty"`
	QueueName            string   `protobuf:"bytes,4,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty"`
	ClusterName          string   `protobuf:"bytes,5,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Platform_MessageBusInfo) Reset()         { *m = Platform_MessageBusInfo{} }
func (m *Platform_MessageBusInfo) String() string { return proto.CompactTextString(m) }
func (*Platform_MessageBusInfo) ProtoMessage()    {}
func (*Platform_MessageBusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_nexus_config_267e1c975478cf7f, []int{2, 5}
}
func (m *Platform_MessageBusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform_MessageBusInfo.Unmarshal(m, b)
}
func (m *Platform_MessageBusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform_MessageBusInfo.Marshal(b, m, deterministic)
}
func (dst *Platform_MessageBusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform_MessageBusInfo.Merge(dst, src)
}
func (m *Platform_MessageBusInfo) XXX_Size() int {
	return xxx_messageInfo_Platform_MessageBusInfo.Size(m)
}
func (m *Platform_MessageBusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform_MessageBusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Platform_MessageBusInfo proto.InternalMessageInfo

func (m *Platform_MessageBusInfo) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *Platform_MessageBusInfo) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Platform_MessageBusInfo) GetPrefixName() string {
	if m != nil {
		return m.PrefixName
	}
	return ""
}

func (m *Platform_MessageBusInfo) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *Platform_MessageBusInfo) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

type Platform_Sidecar struct {
	Server               *Platform_SidecarInfo `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Platform_Sidecar) Reset()         { *m = Platform_Sidecar{} }
func (m *Platform_Sidecar) String() string { return proto.CompactTextString(m) }
func (*Platform_Sidecar) ProtoMessage()    {}
func (*Platform_Sidecar) Descriptor() ([]byte, []int) {
	return fileDescriptor_nexus_config_267e1c975478cf7f, []int{2, 6}
}
func (m *Platform_Sidecar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Platform_Sidecar.Unmarshal(m, b)
}
func (m *Platform_Sidecar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Platform_Sidecar.Marshal(b, m, deterministic)
}
func (dst *Platform_Sidecar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Platform_Sidecar.Merge(dst, src)
}
func (m *Platform_Sidecar) XXX_Size() int {
	return xxx_messageInfo_Platform_Sidecar.Size(m)
}
func (m *Platform_Sidecar) XXX_DiscardUnknown() {
	xxx_messageInfo_Platform_Sidecar.DiscardUnknown(m)
}

var xxx_messageInfo_Platform_Sidecar proto.InternalMessageInfo

func (m *Platform_Sidecar) GetServer() *Platform_SidecarInfo {
	if m != nil {
		return m.Server
	}
	return nil
}

func init() {
	proto.RegisterType((*NexusConfig)(nil), "config.NexusConfig")
	proto.RegisterType((*Cloud)(nil), "config.Cloud")
	proto.RegisterType((*Cloud_Credentials)(nil), "config.Cloud.Credentials")
	proto.RegisterType((*Platform)(nil), "config.Platform")
	proto.RegisterType((*Platform_HealthInfo)(nil), "config.Platform.HealthInfo")
	proto.RegisterType((*Platform_SidecarInfo)(nil), "config.Platform.SidecarInfo")
	proto.RegisterType((*Platform_DatabaseInfo)(nil), "config.Platform.DatabaseInfo")
	proto.RegisterType((*Platform_DatabaseInfo_Credentials)(nil), "config.Platform.DatabaseInfo.Credentials")
	proto.RegisterType((*Platform_DatabaseInfo_EtcdManagerConfig)(nil), "config.Platform.DatabaseInfo.EtcdManagerConfig")
	proto.RegisterType((*Platform_Health)(nil), "config.Platform.Health")
	proto.RegisterType((*Platform_Database)(nil), "config.Platform.Database")
	proto.RegisterType((*Platform_MessageBusInfo)(nil), "config.Platform.MessageBusInfo")
	proto.RegisterType((*Platform_Sidecar)(nil), "config.Platform.Sidecar")
}

func init() { proto.RegisterFile("nexus-config.proto", fileDescriptor_nexus_config_267e1c975478cf7f) }

var fileDescriptor_nexus_config_267e1c975478cf7f = []byte{
	// 876 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdb, 0x6e, 0xdc, 0x36,
	0x10, 0xc5, 0x6e, 0xf7, 0xa6, 0xd1, 0x3a, 0xb1, 0x19, 0xb4, 0x51, 0x94, 0x18, 0xb1, 0x9d, 0x87,
	0x3a, 0x40, 0xbb, 0x01, 0xec, 0x5e, 0x80, 0x16, 0x41, 0x8b, 0xb8, 0x0d, 0x6a, 0x04, 0x0e, 0x0c,
	0xd5, 0x4f, 0x7d, 0x11, 0x68, 0x6a, 0x76, 0x57, 0xad, 0x24, 0x2a, 0x24, 0xb5, 0xb5, 0xf2, 0x03,
	0xfd, 0x86, 0xfe, 0x46, 0x3f, 0xa9, 0x0f, 0xfd, 0x8e, 0x82, 0x37, 0x79, 0x1d, 0x63, 0x9d, 0x37,
	0xce, 0x39, 0x67, 0x38, 0xc3, 0xe1, 0x70, 0x08, 0xa4, 0xc2, 0xab, 0x46, 0x7e, 0xc9, 0x78, 0x35,
	0xcf, 0x17, 0xb3, 0x5a, 0x70, 0xc5, 0xc9, 0xc8, 0x5a, 0x07, 0xef, 0x21, 0x7c, 0xab, 0xd9, 0x13,
	0x63, 0x92, 0x08, 0xc6, 0x2b, 0x14, 0x32, 0xe7, 0x55, 0xd4, 0xdb, 0xeb, 0x1d, 0x06, 0x89, 0x37,
	0xc9, 0x33, 0x18, 0xb2, 0x82, 0x37, 0x59, 0xd4, 0xdf, 0xeb, 0x1d, 0x86, 0x47, 0x5b, 0x33, 0xb7,
	0xdd, 0x89, 0x06, 0x13, 0xcb, 0x91, 0x2f, 0x60, 0x52, 0x17, 0x54, 0xcd, 0xb9, 0x28, 0xa3, 0x4f,
	0x8c, 0x6e, 0xdb, 0xeb, 0xce, 0x1d, 0x9e, 0x74, 0x8a, 0x83, 0xbf, 0xfb, 0x30, 0x34, 0xee, 0x24,
	0x86, 0x49, 0x2d, 0xf8, 0x2a, 0xcf, 0x50, 0xb8, 0xb8, 0x9d, 0x4d, 0x3e, 0x83, 0x91, 0xc0, 0x85,
	0xce, 0xa8, 0x6f, 0x18, 0x67, 0x11, 0x02, 0x83, 0xf7, 0xbc, 0x42, 0x13, 0x27, 0x48, 0xcc, 0x9a,
	0x7c, 0x0f, 0x21, 0x13, 0x98, 0x61, 0xa5, 0x72, 0x5a, 0xc8, 0x68, 0x60, 0x52, 0x78, 0x74, 0x23,
	0xd5, 0xd9, 0xc9, 0xb5, 0x20, 0x59, 0x57, 0x93, 0xa7, 0x10, 0x4a, 0x25, 0x90, 0x96, 0x69, 0x45,
	0x4b, 0x8c, 0x86, 0x66, 0x5f, 0xb0, 0xd0, 0x5b, 0x5a, 0x62, 0xfc, 0x3b, 0x84, 0x6b, 0xce, 0x64,
	0x17, 0x80, 0x32, 0x86, 0x52, 0xa6, 0x7f, 0x60, 0xeb, 0xd2, 0x0e, 0x2c, 0xf2, 0x06, 0x5b, 0x4d,
	0x4b, 0x64, 0x02, 0x95, 0xa1, 0x6d, 0xee, 0x81, 0x45, 0x1c, 0x4d, 0x19, 0xe3, 0x4d, 0xa5, 0xd2,
	0x3c, 0x73, 0x87, 0x08, 0x1c, 0x72, 0x9a, 0x1d, 0xfc, 0x17, 0xc2, 0xc4, 0x97, 0x8c, 0x7c, 0x0d,
	0x93, 0x8c, 0x2a, 0x7a, 0x49, 0x25, 0x9a, 0x38, 0x6b, 0x67, 0xf2, 0x9a, 0xd9, 0x4f, 0x4e, 0x90,
	0x74, 0x52, 0xf2, 0x23, 0x84, 0x25, 0x4a, 0x49, 0x17, 0x98, 0x5e, 0x36, 0xd2, 0x5d, 0xdc, 0xd3,
	0x5b, 0x9e, 0x67, 0x56, 0xf3, 0xaa, 0x91, 0xa7, 0xd5, 0x9c, 0x27, 0x50, 0x76, 0x36, 0x79, 0x01,
	0xa3, 0x25, 0xd2, 0x42, 0x2d, 0xdd, 0x6d, 0x3e, 0xbc, 0xe5, 0xfc, 0x8b, 0xa1, 0x13, 0x27, 0x23,
	0x47, 0x30, 0x96, 0x79, 0x86, 0x8c, 0x0a, 0x57, 0xfc, 0xe8, 0x96, 0xc7, 0xaf, 0x96, 0x4f, 0xbc,
	0x30, 0xde, 0x03, 0xb0, 0xbb, 0xe8, 0xf0, 0xfa, 0x5a, 0x6b, 0x2e, 0x94, 0x39, 0xe7, 0x56, 0x62,
	0xd6, 0xf1, 0x3e, 0x84, 0xce, 0x6b, 0xa3, 0xe4, 0xaf, 0x11, 0x4c, 0x7d, 0x09, 0xbc, 0x68, 0xc9,
	0xa5, 0x72, 0xf7, 0x62, 0xd6, 0x9d, 0x63, 0xff, 0xda, 0x91, 0x3c, 0x83, 0x2d, 0x5f, 0xb0, 0x54,
	0xb5, 0xb5, 0xef, 0xa7, 0xa9, 0x07, 0x2f, 0xda, 0x1a, 0x75, 0x7f, 0x16, 0x9c, 0x51, 0xa5, 0xbb,
	0x70, 0x60, 0xfb, 0xd3, 0xdb, 0xe4, 0xcd, 0xcd, 0x9e, 0x1b, 0x9a, 0x63, 0x3f, 0xdf, 0x78, 0x3f,
	0x3a, 0xb9, 0xcd, 0x3d, 0x48, 0x60, 0x60, 0x9a, 0x6f, 0x64, 0xb3, 0xd6, 0x6b, 0x1d, 0x9c, 0xd1,
	0x9a, 0xb2, 0x5c, 0xb5, 0xd1, 0xd8, 0x06, 0xf7, 0x36, 0x79, 0x02, 0x81, 0x40, 0xa5, 0xbd, 0x79,
	0x15, 0x4d, 0x6c, 0x13, 0x75, 0x00, 0x79, 0x0c, 0xc1, 0x52, 0xa9, 0x3a, 0x35, 0x87, 0x0e, 0xcc,
	0xa1, 0x27, 0x1a, 0x38, 0xd7, 0x07, 0xff, 0x1c, 0xee, 0x0b, 0x7c, 0xd7, 0xa0, 0x54, 0x32, 0x2d,
	0xb1, 0xe4, 0xa2, 0x8d, 0xc0, 0x6c, 0x70, 0xcf, 0xc3, 0x67, 0x06, 0x25, 0x87, 0xb0, 0x9d, 0x95,
	0xe9, 0x1c, 0xa9, 0x6a, 0x04, 0xa6, 0xf3, 0x82, 0x2e, 0x64, 0x14, 0x5a, 0x65, 0x56, 0xbe, 0xb6,
	0xf0, 0x6b, 0x8d, 0x92, 0x14, 0x1e, 0xa0, 0x62, 0x59, 0x5a, 0xd2, 0x8a, 0x2e, 0x50, 0xa4, 0xb6,
	0x06, 0xd1, 0xd4, 0x94, 0xe4, 0xc5, 0xdd, 0x25, 0xf9, 0x59, 0xb1, 0xec, 0xcc, 0xfa, 0xd9, 0x59,
	0x94, 0xec, 0xe0, 0x87, 0x50, 0xfc, 0xf2, 0xe6, 0x0b, 0x24, 0x30, 0x68, 0x64, 0x37, 0x32, 0xcc,
	0xda, 0x8c, 0x12, 0x2a, 0xe5, 0x9f, 0x5c, 0x64, 0xee, 0xd1, 0x75, 0x76, 0xfc, 0x6f, 0x0f, 0x76,
	0x6e, 0xc5, 0x21, 0xdf, 0xc1, 0x23, 0xc6, 0xcb, 0x9a, 0x32, 0x95, 0x0a, 0x5c, 0xe5, 0x7a, 0xda,
	0xa5, 0x79, 0xa5, 0x50, 0xac, 0x68, 0xe1, 0xb6, 0x7e, 0xe8, 0x04, 0x89, 0xe3, 0x4f, 0x1d, 0x4d,
	0x66, 0xf0, 0xe0, 0xda, 0x57, 0x89, 0x36, 0x55, 0x79, 0x89, 0xc2, 0x05, 0xde, 0xe9, 0xbc, 0x94,
	0x68, 0x2f, 0x34, 0x41, 0x9e, 0xc3, 0xb6, 0xd7, 0xd7, 0xcb, 0x56, 0xe6, 0x8c, 0x16, 0xae, 0xe1,
	0xee, 0x3b, 0xfc, 0xdc, 0xc1, 0xe4, 0x1b, 0xf0, 0x51, 0x75, 0x1d, 0x15, 0x5e, 0x29, 0xb3, 0x39,
	0x6f, 0x94, 0x6b, 0xc1, 0x4f, 0x1d, 0x7d, 0x62, 0xd9, 0x0b, 0x4b, 0xc6, 0x2f, 0x61, 0x64, 0x9f,
	0x13, 0x39, 0x86, 0x91, 0x44, 0xb1, 0x72, 0x05, 0x0a, 0x8f, 0x1e, 0x6f, 0x78, 0xbd, 0xe6, 0xd9,
	0x3b, 0x69, 0x7c, 0x05, 0x13, 0x7f, 0x2f, 0xe4, 0x18, 0x86, 0x0b, 0x41, 0xeb, 0xa5, 0xf3, 0xdf,
	0xbd, 0xf3, 0x06, 0x13, 0xab, 0x25, 0xdf, 0xc2, 0xb8, 0x44, 0x25, 0x72, 0xe6, 0x27, 0xce, 0x47,
	0xdc, 0xbc, 0x3a, 0xfe, 0xa7, 0x07, 0xf7, 0x6e, 0xce, 0xa2, 0x3b, 0xff, 0x85, 0x7d, 0x98, 0xea,
	0x94, 0x73, 0x86, 0x76, 0x5e, 0xdb, 0x9a, 0x87, 0x0e, 0xd3, 0x03, 0x5b, 0x4f, 0xf4, 0x5a, 0xe0,
	0x3c, 0xbf, 0xb2, 0x0a, 0x5b, 0x68, 0xb0, 0x90, 0x11, 0xec, 0x02, 0xbc, 0x6b, 0xb0, 0x71, 0x3b,
	0xd8, 0xb2, 0x06, 0x06, 0x31, 0xf4, 0x3e, 0x4c, 0x59, 0xd1, 0x48, 0x85, 0x62, 0xfd, 0x4b, 0x08,
	0x1d, 0x66, 0xfe, 0x84, 0x1f, 0x60, 0xec, 0x46, 0x13, 0xf9, 0xea, 0x83, 0x72, 0x3f, 0xd9, 0x34,
	0xfa, 0xd6, 0xeb, 0xfd, 0x6a, 0xf0, 0x5b, 0x3f, 0x2b, 0x2f, 0x47, 0xe6, 0x57, 0x3e, 0xfe, 0x3f,
	0x00, 0x00, 0xff, 0xff, 0x98, 0xf7, 0xce, 0xf5, 0xab, 0x07, 0x00, 0x00,
}
