//
//Copyright 2019 The Vitess Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// This file defines the replication related structures we use.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: replicationdata/replicationdata.proto

package replicationdata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// StopReplicationMode is used to provide controls over how replication is stopped.
type StopReplicationMode int32

const (
	StopReplicationMode_IOANDSQLTHREAD StopReplicationMode = 0
	StopReplicationMode_IOTHREADONLY   StopReplicationMode = 1
)

// Enum value maps for StopReplicationMode.
var (
	StopReplicationMode_name = map[int32]string{
		0: "IOANDSQLTHREAD",
		1: "IOTHREADONLY",
	}
	StopReplicationMode_value = map[string]int32{
		"IOANDSQLTHREAD": 0,
		"IOTHREADONLY":   1,
	}
)

func (x StopReplicationMode) Enum() *StopReplicationMode {
	p := new(StopReplicationMode)
	*p = x
	return p
}

func (x StopReplicationMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StopReplicationMode) Descriptor() protoreflect.EnumDescriptor {
	return file_replicationdata_replicationdata_proto_enumTypes[0].Descriptor()
}

func (StopReplicationMode) Type() protoreflect.EnumType {
	return &file_replicationdata_replicationdata_proto_enumTypes[0]
}

func (x StopReplicationMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StopReplicationMode.Descriptor instead.
func (StopReplicationMode) EnumDescriptor() ([]byte, []int) {
	return file_replicationdata_replicationdata_proto_rawDescGZIP(), []int{0}
}

// Status is the replication status for MySQL/MariaDB/File-based. Returned by a
// flavor-specific command and parsed into a Position and fields.
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position              string `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	ReplicationLagSeconds uint32 `protobuf:"varint,4,opt,name=replication_lag_seconds,json=replicationLagSeconds,proto3" json:"replication_lag_seconds,omitempty"`
	SourceHost            string `protobuf:"bytes,5,opt,name=source_host,json=sourceHost,proto3" json:"source_host,omitempty"`
	SourcePort            int32  `protobuf:"varint,6,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	ConnectRetry          int32  `protobuf:"varint,7,opt,name=connect_retry,json=connectRetry,proto3" json:"connect_retry,omitempty"`
	// RelayLogPosition will be empty for flavors that do not support returning the full GTIDSet from the relay log, such as MariaDB.
	RelayLogPosition                       string `protobuf:"bytes,8,opt,name=relay_log_position,json=relayLogPosition,proto3" json:"relay_log_position,omitempty"`
	FilePosition                           string `protobuf:"bytes,9,opt,name=file_position,json=filePosition,proto3" json:"file_position,omitempty"`
	RelayLogSourceBinlogEquivalentPosition string `protobuf:"bytes,10,opt,name=relay_log_source_binlog_equivalent_position,json=relayLogSourceBinlogEquivalentPosition,proto3" json:"relay_log_source_binlog_equivalent_position,omitempty"`
	SourceServerId                         uint32 `protobuf:"varint,11,opt,name=source_server_id,json=sourceServerId,proto3" json:"source_server_id,omitempty"`
	SourceUuid                             string `protobuf:"bytes,12,opt,name=source_uuid,json=sourceUuid,proto3" json:"source_uuid,omitempty"`
	IoState                                int32  `protobuf:"varint,13,opt,name=io_state,json=ioState,proto3" json:"io_state,omitempty"`
	LastIoError                            string `protobuf:"bytes,14,opt,name=last_io_error,json=lastIoError,proto3" json:"last_io_error,omitempty"`
	SqlState                               int32  `protobuf:"varint,15,opt,name=sql_state,json=sqlState,proto3" json:"sql_state,omitempty"`
	LastSqlError                           string `protobuf:"bytes,16,opt,name=last_sql_error,json=lastSqlError,proto3" json:"last_sql_error,omitempty"`
	RelayLogFilePosition                   string `protobuf:"bytes,17,opt,name=relay_log_file_position,json=relayLogFilePosition,proto3" json:"relay_log_file_position,omitempty"`
	SourceUser                             string `protobuf:"bytes,18,opt,name=source_user,json=sourceUser,proto3" json:"source_user,omitempty"`
	SqlDelay                               uint32 `protobuf:"varint,19,opt,name=sql_delay,json=sqlDelay,proto3" json:"sql_delay,omitempty"`
	AutoPosition                           bool   `protobuf:"varint,20,opt,name=auto_position,json=autoPosition,proto3" json:"auto_position,omitempty"`
	UsingGtid                              bool   `protobuf:"varint,21,opt,name=using_gtid,json=usingGtid,proto3" json:"using_gtid,omitempty"`
	HasReplicationFilters                  bool   `protobuf:"varint,22,opt,name=has_replication_filters,json=hasReplicationFilters,proto3" json:"has_replication_filters,omitempty"`
	SslAllowed                             bool   `protobuf:"varint,23,opt,name=ssl_allowed,json=sslAllowed,proto3" json:"ssl_allowed,omitempty"`
	ReplicationLagUnknown                  bool   `protobuf:"varint,24,opt,name=replication_lag_unknown,json=replicationLagUnknown,proto3" json:"replication_lag_unknown,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replicationdata_replicationdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_replicationdata_replicationdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_replicationdata_replicationdata_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Status) GetReplicationLagSeconds() uint32 {
	if x != nil {
		return x.ReplicationLagSeconds
	}
	return 0
}

func (x *Status) GetSourceHost() string {
	if x != nil {
		return x.SourceHost
	}
	return ""
}

func (x *Status) GetSourcePort() int32 {
	if x != nil {
		return x.SourcePort
	}
	return 0
}

func (x *Status) GetConnectRetry() int32 {
	if x != nil {
		return x.ConnectRetry
	}
	return 0
}

func (x *Status) GetRelayLogPosition() string {
	if x != nil {
		return x.RelayLogPosition
	}
	return ""
}

func (x *Status) GetFilePosition() string {
	if x != nil {
		return x.FilePosition
	}
	return ""
}

func (x *Status) GetRelayLogSourceBinlogEquivalentPosition() string {
	if x != nil {
		return x.RelayLogSourceBinlogEquivalentPosition
	}
	return ""
}

func (x *Status) GetSourceServerId() uint32 {
	if x != nil {
		return x.SourceServerId
	}
	return 0
}

func (x *Status) GetSourceUuid() string {
	if x != nil {
		return x.SourceUuid
	}
	return ""
}

func (x *Status) GetIoState() int32 {
	if x != nil {
		return x.IoState
	}
	return 0
}

func (x *Status) GetLastIoError() string {
	if x != nil {
		return x.LastIoError
	}
	return ""
}

func (x *Status) GetSqlState() int32 {
	if x != nil {
		return x.SqlState
	}
	return 0
}

func (x *Status) GetLastSqlError() string {
	if x != nil {
		return x.LastSqlError
	}
	return ""
}

func (x *Status) GetRelayLogFilePosition() string {
	if x != nil {
		return x.RelayLogFilePosition
	}
	return ""
}

func (x *Status) GetSourceUser() string {
	if x != nil {
		return x.SourceUser
	}
	return ""
}

func (x *Status) GetSqlDelay() uint32 {
	if x != nil {
		return x.SqlDelay
	}
	return 0
}

func (x *Status) GetAutoPosition() bool {
	if x != nil {
		return x.AutoPosition
	}
	return false
}

func (x *Status) GetUsingGtid() bool {
	if x != nil {
		return x.UsingGtid
	}
	return false
}

func (x *Status) GetHasReplicationFilters() bool {
	if x != nil {
		return x.HasReplicationFilters
	}
	return false
}

func (x *Status) GetSslAllowed() bool {
	if x != nil {
		return x.SslAllowed
	}
	return false
}

func (x *Status) GetReplicationLagUnknown() bool {
	if x != nil {
		return x.ReplicationLagUnknown
	}
	return false
}

// StopReplicationStatus represents the replication status before calling StopReplication, and the replication status collected immediately after
// calling StopReplication.
type StopReplicationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Before *Status `protobuf:"bytes,1,opt,name=before,proto3" json:"before,omitempty"`
	After  *Status `protobuf:"bytes,2,opt,name=after,proto3" json:"after,omitempty"`
}

func (x *StopReplicationStatus) Reset() {
	*x = StopReplicationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replicationdata_replicationdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopReplicationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopReplicationStatus) ProtoMessage() {}

func (x *StopReplicationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_replicationdata_replicationdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopReplicationStatus.ProtoReflect.Descriptor instead.
func (*StopReplicationStatus) Descriptor() ([]byte, []int) {
	return file_replicationdata_replicationdata_proto_rawDescGZIP(), []int{1}
}

func (x *StopReplicationStatus) GetBefore() *Status {
	if x != nil {
		return x.Before
	}
	return nil
}

func (x *StopReplicationStatus) GetAfter() *Status {
	if x != nil {
		return x.After
	}
	return nil
}

// PrimaryStatus is the replication status for a MySQL primary (returned by 'show master status').
type PrimaryStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position     string `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	FilePosition string `protobuf:"bytes,2,opt,name=file_position,json=filePosition,proto3" json:"file_position,omitempty"`
}

func (x *PrimaryStatus) Reset() {
	*x = PrimaryStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replicationdata_replicationdata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimaryStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimaryStatus) ProtoMessage() {}

func (x *PrimaryStatus) ProtoReflect() protoreflect.Message {
	mi := &file_replicationdata_replicationdata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimaryStatus.ProtoReflect.Descriptor instead.
func (*PrimaryStatus) Descriptor() ([]byte, []int) {
	return file_replicationdata_replicationdata_proto_rawDescGZIP(), []int{2}
}

func (x *PrimaryStatus) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *PrimaryStatus) GetFilePosition() string {
	if x != nil {
		return x.FilePosition
	}
	return ""
}

// FullStatus contains the full status of MySQL including the replication information, semi-sync information, GTID information among others
type FullStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId                    uint32         `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ServerUuid                  string         `protobuf:"bytes,2,opt,name=server_uuid,json=serverUuid,proto3" json:"server_uuid,omitempty"`
	ReplicationStatus           *Status        `protobuf:"bytes,3,opt,name=replication_status,json=replicationStatus,proto3" json:"replication_status,omitempty"`
	PrimaryStatus               *PrimaryStatus `protobuf:"bytes,4,opt,name=primary_status,json=primaryStatus,proto3" json:"primary_status,omitempty"`
	GtidPurged                  string         `protobuf:"bytes,5,opt,name=gtid_purged,json=gtidPurged,proto3" json:"gtid_purged,omitempty"`
	Version                     string         `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	VersionComment              string         `protobuf:"bytes,7,opt,name=version_comment,json=versionComment,proto3" json:"version_comment,omitempty"`
	ReadOnly                    bool           `protobuf:"varint,8,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	GtidMode                    string         `protobuf:"bytes,9,opt,name=gtid_mode,json=gtidMode,proto3" json:"gtid_mode,omitempty"`
	BinlogFormat                string         `protobuf:"bytes,10,opt,name=binlog_format,json=binlogFormat,proto3" json:"binlog_format,omitempty"`
	BinlogRowImage              string         `protobuf:"bytes,11,opt,name=binlog_row_image,json=binlogRowImage,proto3" json:"binlog_row_image,omitempty"`
	LogBinEnabled               bool           `protobuf:"varint,12,opt,name=log_bin_enabled,json=logBinEnabled,proto3" json:"log_bin_enabled,omitempty"`
	LogReplicaUpdates           bool           `protobuf:"varint,13,opt,name=log_replica_updates,json=logReplicaUpdates,proto3" json:"log_replica_updates,omitempty"`
	SemiSyncPrimaryEnabled      bool           `protobuf:"varint,14,opt,name=semi_sync_primary_enabled,json=semiSyncPrimaryEnabled,proto3" json:"semi_sync_primary_enabled,omitempty"`
	SemiSyncReplicaEnabled      bool           `protobuf:"varint,15,opt,name=semi_sync_replica_enabled,json=semiSyncReplicaEnabled,proto3" json:"semi_sync_replica_enabled,omitempty"`
	SemiSyncPrimaryStatus       bool           `protobuf:"varint,16,opt,name=semi_sync_primary_status,json=semiSyncPrimaryStatus,proto3" json:"semi_sync_primary_status,omitempty"`
	SemiSyncReplicaStatus       bool           `protobuf:"varint,17,opt,name=semi_sync_replica_status,json=semiSyncReplicaStatus,proto3" json:"semi_sync_replica_status,omitempty"`
	SemiSyncPrimaryClients      uint32         `protobuf:"varint,18,opt,name=semi_sync_primary_clients,json=semiSyncPrimaryClients,proto3" json:"semi_sync_primary_clients,omitempty"`
	SemiSyncPrimaryTimeout      uint64         `protobuf:"varint,19,opt,name=semi_sync_primary_timeout,json=semiSyncPrimaryTimeout,proto3" json:"semi_sync_primary_timeout,omitempty"`
	SemiSyncWaitForReplicaCount uint32         `protobuf:"varint,20,opt,name=semi_sync_wait_for_replica_count,json=semiSyncWaitForReplicaCount,proto3" json:"semi_sync_wait_for_replica_count,omitempty"`
}

func (x *FullStatus) Reset() {
	*x = FullStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_replicationdata_replicationdata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullStatus) ProtoMessage() {}

func (x *FullStatus) ProtoReflect() protoreflect.Message {
	mi := &file_replicationdata_replicationdata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullStatus.ProtoReflect.Descriptor instead.
func (*FullStatus) Descriptor() ([]byte, []int) {
	return file_replicationdata_replicationdata_proto_rawDescGZIP(), []int{3}
}

func (x *FullStatus) GetServerId() uint32 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *FullStatus) GetServerUuid() string {
	if x != nil {
		return x.ServerUuid
	}
	return ""
}

func (x *FullStatus) GetReplicationStatus() *Status {
	if x != nil {
		return x.ReplicationStatus
	}
	return nil
}

func (x *FullStatus) GetPrimaryStatus() *PrimaryStatus {
	if x != nil {
		return x.PrimaryStatus
	}
	return nil
}

func (x *FullStatus) GetGtidPurged() string {
	if x != nil {
		return x.GtidPurged
	}
	return ""
}

func (x *FullStatus) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *FullStatus) GetVersionComment() string {
	if x != nil {
		return x.VersionComment
	}
	return ""
}

func (x *FullStatus) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

func (x *FullStatus) GetGtidMode() string {
	if x != nil {
		return x.GtidMode
	}
	return ""
}

func (x *FullStatus) GetBinlogFormat() string {
	if x != nil {
		return x.BinlogFormat
	}
	return ""
}

func (x *FullStatus) GetBinlogRowImage() string {
	if x != nil {
		return x.BinlogRowImage
	}
	return ""
}

func (x *FullStatus) GetLogBinEnabled() bool {
	if x != nil {
		return x.LogBinEnabled
	}
	return false
}

func (x *FullStatus) GetLogReplicaUpdates() bool {
	if x != nil {
		return x.LogReplicaUpdates
	}
	return false
}

func (x *FullStatus) GetSemiSyncPrimaryEnabled() bool {
	if x != nil {
		return x.SemiSyncPrimaryEnabled
	}
	return false
}

func (x *FullStatus) GetSemiSyncReplicaEnabled() bool {
	if x != nil {
		return x.SemiSyncReplicaEnabled
	}
	return false
}

func (x *FullStatus) GetSemiSyncPrimaryStatus() bool {
	if x != nil {
		return x.SemiSyncPrimaryStatus
	}
	return false
}

func (x *FullStatus) GetSemiSyncReplicaStatus() bool {
	if x != nil {
		return x.SemiSyncReplicaStatus
	}
	return false
}

func (x *FullStatus) GetSemiSyncPrimaryClients() uint32 {
	if x != nil {
		return x.SemiSyncPrimaryClients
	}
	return 0
}

func (x *FullStatus) GetSemiSyncPrimaryTimeout() uint64 {
	if x != nil {
		return x.SemiSyncPrimaryTimeout
	}
	return 0
}

func (x *FullStatus) GetSemiSyncWaitForReplicaCount() uint32 {
	if x != nil {
		return x.SemiSyncWaitForReplicaCount
	}
	return 0
}

var File_replicationdata_replicationdata_proto protoreflect.FileDescriptor

var file_replicationdata_replicationdata_proto_rawDesc = []byte{
	0x0a, 0x25, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x22, 0x96, 0x07, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x36, 0x0a, 0x17, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c,
	0x61, 0x67, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x15, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x67,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x74, 0x72, 0x79, 0x12, 0x2c,
	0x0a, 0x12, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x4c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x5b, 0x0a, 0x2b, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x71, 0x75,
	0x69, 0x76, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x26, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x4c, 0x6f, 0x67,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x45, 0x71, 0x75, 0x69,
	0x76, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28,
	0x0a, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6f, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x6f, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6f, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x49, 0x6f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x71, 0x6c, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x71, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x71,
	0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x53, 0x71, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x35, 0x0a, 0x17, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x71, 0x6c, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x71, 0x6c, 0x44, 0x65, 0x6c, 0x61, 0x79,
	0x12, 0x23, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x67,
	0x74, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x75, 0x73, 0x69, 0x6e, 0x67,
	0x47, 0x74, 0x69, 0x64, 0x12, 0x36, 0x0a, 0x17, 0x68, 0x61, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x68, 0x61, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x73, 0x6c, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x73, 0x73, 0x6c, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x36, 0x0a,
	0x17, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x67,
	0x5f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x67, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10,
	0x04, 0x22, 0x77, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x62, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x05, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x0d, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc3, 0x07, 0x0a,
	0x0a, 0x46, 0x75, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x46, 0x0a, 0x12, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x11,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x45, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x74, 0x69, 0x64,
	0x5f, 0x70, 0x75, 0x72, 0x67, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67,
	0x74, 0x69, 0x64, 0x50, 0x75, 0x72, 0x67, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x74, 0x69,
	0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x74,
	0x69, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67,
	0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62,
	0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x62,
	0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x6f, 0x77, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x52, 0x6f, 0x77,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x5f, 0x62, 0x69, 0x6e,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x6c, 0x6f, 0x67, 0x42, 0x69, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2e, 0x0a,
	0x13, 0x6c, 0x6f, 0x67, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x39, 0x0a,
	0x19, 0x73, 0x65, 0x6d, 0x69, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x16, 0x73, 0x65, 0x6d, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x19, 0x73, 0x65, 0x6d, 0x69,
	0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x73, 0x65, 0x6d,
	0x69, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x73, 0x65, 0x6d, 0x69, 0x5f, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x73, 0x65, 0x6d, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x18,
	0x73, 0x65, 0x6d, 0x69, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15,
	0x73, 0x65, 0x6d, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x73, 0x65, 0x6d, 0x69, 0x5f, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x73, 0x65, 0x6d, 0x69, 0x53, 0x79,
	0x6e, 0x63, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x39, 0x0a, 0x19, 0x73, 0x65, 0x6d, 0x69, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x16, 0x73, 0x65, 0x6d, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x45, 0x0a, 0x20, 0x73,
	0x65, 0x6d, 0x69, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x66, 0x6f,
	0x72, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1b, 0x73, 0x65, 0x6d, 0x69, 0x53, 0x79, 0x6e, 0x63, 0x57,
	0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x2a, 0x3b, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x4f, 0x41,
	0x4e, 0x44, 0x53, 0x51, 0x4c, 0x54, 0x48, 0x52, 0x45, 0x41, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x49, 0x4f, 0x54, 0x48, 0x52, 0x45, 0x41, 0x44, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x01, 0x42,
	0xcb, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x42, 0x14, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2d, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f,
	0x76, 0x31, 0x35, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64,
	0x61, 0x74, 0x61, 0xa2, 0x02, 0x03, 0x52, 0x58, 0x58, 0xaa, 0x02, 0x0f, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0xca, 0x02, 0x0f, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0xe2, 0x02, 0x1b,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_replicationdata_replicationdata_proto_rawDescOnce sync.Once
	file_replicationdata_replicationdata_proto_rawDescData = file_replicationdata_replicationdata_proto_rawDesc
)

func file_replicationdata_replicationdata_proto_rawDescGZIP() []byte {
	file_replicationdata_replicationdata_proto_rawDescOnce.Do(func() {
		file_replicationdata_replicationdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_replicationdata_replicationdata_proto_rawDescData)
	})
	return file_replicationdata_replicationdata_proto_rawDescData
}

var file_replicationdata_replicationdata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_replicationdata_replicationdata_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_replicationdata_replicationdata_proto_goTypes = []interface{}{
	(StopReplicationMode)(0),      // 0: replicationdata.StopReplicationMode
	(*Status)(nil),                // 1: replicationdata.Status
	(*StopReplicationStatus)(nil), // 2: replicationdata.StopReplicationStatus
	(*PrimaryStatus)(nil),         // 3: replicationdata.PrimaryStatus
	(*FullStatus)(nil),            // 4: replicationdata.FullStatus
}
var file_replicationdata_replicationdata_proto_depIdxs = []int32{
	1, // 0: replicationdata.StopReplicationStatus.before:type_name -> replicationdata.Status
	1, // 1: replicationdata.StopReplicationStatus.after:type_name -> replicationdata.Status
	1, // 2: replicationdata.FullStatus.replication_status:type_name -> replicationdata.Status
	3, // 3: replicationdata.FullStatus.primary_status:type_name -> replicationdata.PrimaryStatus
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_replicationdata_replicationdata_proto_init() }
func file_replicationdata_replicationdata_proto_init() {
	if File_replicationdata_replicationdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_replicationdata_replicationdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_replicationdata_replicationdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopReplicationStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_replicationdata_replicationdata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimaryStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_replicationdata_replicationdata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_replicationdata_replicationdata_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_replicationdata_replicationdata_proto_goTypes,
		DependencyIndexes: file_replicationdata_replicationdata_proto_depIdxs,
		EnumInfos:         file_replicationdata_replicationdata_proto_enumTypes,
		MessageInfos:      file_replicationdata_replicationdata_proto_msgTypes,
	}.Build()
	File_replicationdata_replicationdata_proto = out.File
	file_replicationdata_replicationdata_proto_rawDesc = nil
	file_replicationdata_replicationdata_proto_goTypes = nil
	file_replicationdata_replicationdata_proto_depIdxs = nil
}