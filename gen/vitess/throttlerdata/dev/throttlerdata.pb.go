//
//Copyright 2019 The Vitess Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Data structures for the throttler RPC interface.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: vitess/throttlerdata/dev/throttlerdata.proto

package throttlerdatadev

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MaxRatesRequest is the payload for the MaxRates RPC.
type MaxRatesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MaxRatesRequest) Reset() {
	*x = MaxRatesRequest{}
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MaxRatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxRatesRequest) ProtoMessage() {}

func (x *MaxRatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxRatesRequest.ProtoReflect.Descriptor instead.
func (*MaxRatesRequest) Descriptor() ([]byte, []int) {
	return file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescGZIP(), []int{0}
}

// MaxRatesResponse is returned by the MaxRates RPC.
type MaxRatesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// max_rates returns the max rate for each throttler. It's keyed by the
	// throttler name.
	Rates         map[string]int64 `protobuf:"bytes,1,rep,name=rates,proto3" json:"rates,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MaxRatesResponse) Reset() {
	*x = MaxRatesResponse{}
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MaxRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxRatesResponse) ProtoMessage() {}

func (x *MaxRatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxRatesResponse.ProtoReflect.Descriptor instead.
func (*MaxRatesResponse) Descriptor() ([]byte, []int) {
	return file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescGZIP(), []int{1}
}

func (x *MaxRatesResponse) GetRates() map[string]int64 {
	if x != nil {
		return x.Rates
	}
	return nil
}

// SetMaxRateRequest is the payload for the SetMaxRate RPC.
type SetMaxRateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rate          int64                  `protobuf:"varint,1,opt,name=rate,proto3" json:"rate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetMaxRateRequest) Reset() {
	*x = SetMaxRateRequest{}
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetMaxRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMaxRateRequest) ProtoMessage() {}

func (x *SetMaxRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMaxRateRequest.ProtoReflect.Descriptor instead.
func (*SetMaxRateRequest) Descriptor() ([]byte, []int) {
	return file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescGZIP(), []int{2}
}

func (x *SetMaxRateRequest) GetRate() int64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

// SetMaxRateResponse is returned by the SetMaxRate RPC.
type SetMaxRateResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// names is the list of throttler names which were updated.
	Names         []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetMaxRateResponse) Reset() {
	*x = SetMaxRateResponse{}
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetMaxRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMaxRateResponse) ProtoMessage() {}

func (x *SetMaxRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMaxRateResponse.ProtoReflect.Descriptor instead.
func (*SetMaxRateResponse) Descriptor() ([]byte, []int) {
	return file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescGZIP(), []int{3}
}

func (x *SetMaxRateResponse) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

// Configuration holds the configuration parameters for the
// MaxReplicationLagModule which adaptively adjusts the throttling rate based on
// the observed replication lag across all replicas.
type Configuration struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// target_replication_lag_sec is the replication lag (in seconds) the
	// MaxReplicationLagModule tries to aim for.
	// If it is within the target, it tries to increase the throttler
	// rate, otherwise it will lower it based on an educated guess of the
	// replica's throughput.
	TargetReplicationLagSec int64 `protobuf:"varint,1,opt,name=target_replication_lag_sec,json=targetReplicationLagSec,proto3" json:"target_replication_lag_sec,omitempty"`
	// max_replication_lag_sec is meant as a last resort.
	// By default, the module tries to find out the system maximum capacity while
	// trying to keep the replication lag around "target_replication_lag_sec".
	// Usually, we'll wait min_duration_between_(increases|decreases)_sec to see
	// the effect of a throttler rate change on the replication lag.
	// But if the lag goes above this field's value we will go into an "emergency"
	// state and throttle more aggressively (see "emergency_decrease" below).
	// This is the only way to ensure that the system will recover.
	MaxReplicationLagSec int64 `protobuf:"varint,2,opt,name=max_replication_lag_sec,json=maxReplicationLagSec,proto3" json:"max_replication_lag_sec,omitempty"`
	// initial_rate is the rate at which the module will start.
	InitialRate int64 `protobuf:"varint,3,opt,name=initial_rate,json=initialRate,proto3" json:"initial_rate,omitempty"`
	// max_increase defines by how much we will increase the rate
	// e.g. 0.05 increases the rate by 5% while 1.0 by 100%.
	// Note that any increase will let the system wait for at least
	// (1 / MaxIncrease) seconds. If we wait for shorter periods of time, we
	// won't notice if the rate increase also increases the replication lag.
	// (If the system was already at its maximum capacity (e.g. 1k QPS) and we
	// increase the rate by e.g. 5% to 1050 QPS, it will take 20 seconds until
	// 1000 extra queries are buffered and the lag increases by 1 second.)
	MaxIncrease float64 `protobuf:"fixed64,4,opt,name=max_increase,json=maxIncrease,proto3" json:"max_increase,omitempty"`
	// emergency_decrease defines by how much we will decrease the current rate
	// if the observed replication lag is above "max_replication_lag_sec".
	// E.g. 0.50 decreases the current rate by 50%.
	EmergencyDecrease float64 `protobuf:"fixed64,5,opt,name=emergency_decrease,json=emergencyDecrease,proto3" json:"emergency_decrease,omitempty"`
	// min_duration_between_increases_sec specifies how long we'll wait at least
	// for the last rate increase to have an effect on the system.
	MinDurationBetweenIncreasesSec int64 `protobuf:"varint,6,opt,name=min_duration_between_increases_sec,json=minDurationBetweenIncreasesSec,proto3" json:"min_duration_between_increases_sec,omitempty"`
	// max_duration_between_increases_sec specifies how long we'll wait at most
	// for the last rate increase to have an effect on the system.
	MaxDurationBetweenIncreasesSec int64 `protobuf:"varint,7,opt,name=max_duration_between_increases_sec,json=maxDurationBetweenIncreasesSec,proto3" json:"max_duration_between_increases_sec,omitempty"`
	// min_duration_between_decreases_sec specifies how long we'll wait at least
	// for the last rate decrease to have an effect on the system.
	MinDurationBetweenDecreasesSec int64 `protobuf:"varint,8,opt,name=min_duration_between_decreases_sec,json=minDurationBetweenDecreasesSec,proto3" json:"min_duration_between_decreases_sec,omitempty"`
	// spread_backlog_across_sec is used when we set the throttler rate after
	// we guessed the rate of a replica and determined its backlog.
	// For example, at a guessed rate of 100 QPS and a lag of 10s, the replica has
	// a backlog of 1000 queries.
	// When we set the new, decreased throttler rate, we factor in how long it
	// will take the replica to go through the backlog (in addition to new
	// requests). This field specifies over which timespan we plan to spread this.
	// For example, for a backlog of 1000 queries spread over 5s means that we
	// have to further reduce the rate by 200 QPS or the backlog will not be
	// processed within the 5 seconds.
	SpreadBacklogAcrossSec int64 `protobuf:"varint,9,opt,name=spread_backlog_across_sec,json=spreadBacklogAcrossSec,proto3" json:"spread_backlog_across_sec,omitempty"`
	// ignore_n_slowest_replicas will ignore replication lag updates from the
	// N slowest REPLICA tablets. Under certain circumstances, replicas are still
	// considered e.g. a) if the lag is at most max_replication_lag_sec, b) there
	// are less than N+1 replicas or c) the lag increased on each replica such
	// that all replicas were ignored in a row.
	IgnoreNSlowestReplicas int32 `protobuf:"varint,10,opt,name=ignore_n_slowest_replicas,json=ignoreNSlowestReplicas,proto3" json:"ignore_n_slowest_replicas,omitempty"`
	// ignore_n_slowest_rdonlys does the same thing as ignore_n_slowest_replicas
	// but for RDONLY tablets. Note that these two settings are independent.
	IgnoreNSlowestRdonlys int32 `protobuf:"varint,11,opt,name=ignore_n_slowest_rdonlys,json=ignoreNSlowestRdonlys,proto3" json:"ignore_n_slowest_rdonlys,omitempty"`
	// age_bad_rate_after_sec is the duration after which an unchanged bad rate
	// will "age out" and increase by "bad_rate_increase".
	// Bad rates are tracked by the code in memory.go and serve as an upper bound
	// for future rate changes. This ensures that the adaptive throttler does not
	// try known too high (bad) rates over and over again.
	// To avoid that temporary degradations permanently reduce the maximum rate,
	// a stable bad rate "ages out" after "age_bad_rate_after_sec".
	AgeBadRateAfterSec int64 `protobuf:"varint,12,opt,name=age_bad_rate_after_sec,json=ageBadRateAfterSec,proto3" json:"age_bad_rate_after_sec,omitempty"`
	// bad_rate_increase defines the percentage by which a bad rate will be
	// increased when it's aging out.
	BadRateIncrease float64 `protobuf:"fixed64,13,opt,name=bad_rate_increase,json=badRateIncrease,proto3" json:"bad_rate_increase,omitempty"`
	// max_rate_approach_threshold is the fraction of the current rate limit that the actual
	// rate must exceed for the throttler to increase the limit when the replication lag
	// is below target_replication_lag_sec. For example, assuming the actual replication lag
	// is below target_replication_lag_sec, if the current rate limit is 100, then the actual
	// rate must exceed 100*max_rate_approach_threshold for the throttler to increase the current
	// limit.
	MaxRateApproachThreshold float64 `protobuf:"fixed64,14,opt,name=max_rate_approach_threshold,json=maxRateApproachThreshold,proto3" json:"max_rate_approach_threshold,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescGZIP(), []int{4}
}

func (x *Configuration) GetTargetReplicationLagSec() int64 {
	if x != nil {
		return x.TargetReplicationLagSec
	}
	return 0
}

func (x *Configuration) GetMaxReplicationLagSec() int64 {
	if x != nil {
		return x.MaxReplicationLagSec
	}
	return 0
}

func (x *Configuration) GetInitialRate() int64 {
	if x != nil {
		return x.InitialRate
	}
	return 0
}

func (x *Configuration) GetMaxIncrease() float64 {
	if x != nil {
		return x.MaxIncrease
	}
	return 0
}

func (x *Configuration) GetEmergencyDecrease() float64 {
	if x != nil {
		return x.EmergencyDecrease
	}
	return 0
}

func (x *Configuration) GetMinDurationBetweenIncreasesSec() int64 {
	if x != nil {
		return x.MinDurationBetweenIncreasesSec
	}
	return 0
}

func (x *Configuration) GetMaxDurationBetweenIncreasesSec() int64 {
	if x != nil {
		return x.MaxDurationBetweenIncreasesSec
	}
	return 0
}

func (x *Configuration) GetMinDurationBetweenDecreasesSec() int64 {
	if x != nil {
		return x.MinDurationBetweenDecreasesSec
	}
	return 0
}

func (x *Configuration) GetSpreadBacklogAcrossSec() int64 {
	if x != nil {
		return x.SpreadBacklogAcrossSec
	}
	return 0
}

func (x *Configuration) GetIgnoreNSlowestReplicas() int32 {
	if x != nil {
		return x.IgnoreNSlowestReplicas
	}
	return 0
}

func (x *Configuration) GetIgnoreNSlowestRdonlys() int32 {
	if x != nil {
		return x.IgnoreNSlowestRdonlys
	}
	return 0
}

func (x *Configuration) GetAgeBadRateAfterSec() int64 {
	if x != nil {
		return x.AgeBadRateAfterSec
	}
	return 0
}

func (x *Configuration) GetBadRateIncrease() float64 {
	if x != nil {
		return x.BadRateIncrease
	}
	return 0
}

func (x *Configuration) GetMaxRateApproachThreshold() float64 {
	if x != nil {
		return x.MaxRateApproachThreshold
	}
	return 0
}

// GetConfigurationRequest is the payload for the GetConfiguration RPC.
type GetConfigurationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// throttler_name specifies which throttler to select. If empty, all active
	// throttlers will be selected.
	ThrottlerName string `protobuf:"bytes,1,opt,name=throttler_name,json=throttlerName,proto3" json:"throttler_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetConfigurationRequest) Reset() {
	*x = GetConfigurationRequest{}
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationRequest) ProtoMessage() {}

func (x *GetConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigurationRequest.ProtoReflect.Descriptor instead.
func (*GetConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescGZIP(), []int{5}
}

func (x *GetConfigurationRequest) GetThrottlerName() string {
	if x != nil {
		return x.ThrottlerName
	}
	return ""
}

// GetConfigurationResponse is returned by the GetConfiguration RPC.
type GetConfigurationResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// max_rates returns the configurations for each throttler.
	// It's keyed by the throttler name.
	Configurations map[string]*Configuration `protobuf:"bytes,1,rep,name=configurations,proto3" json:"configurations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetConfigurationResponse) Reset() {
	*x = GetConfigurationResponse{}
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationResponse) ProtoMessage() {}

func (x *GetConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigurationResponse.ProtoReflect.Descriptor instead.
func (*GetConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescGZIP(), []int{6}
}

func (x *GetConfigurationResponse) GetConfigurations() map[string]*Configuration {
	if x != nil {
		return x.Configurations
	}
	return nil
}

// UpdateConfigurationRequest is the payload for the UpdateConfiguration RPC.
type UpdateConfigurationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// throttler_name specifies which throttler to update. If empty, all active
	// throttlers will be updated.
	ThrottlerName string `protobuf:"bytes,1,opt,name=throttler_name,json=throttlerName,proto3" json:"throttler_name,omitempty"`
	// configuration is the new (partial) configuration.
	Configuration *Configuration `protobuf:"bytes,2,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// copy_zero_values specifies whether fields with zero values should be copied
	// as well.
	CopyZeroValues bool `protobuf:"varint,3,opt,name=copy_zero_values,json=copyZeroValues,proto3" json:"copy_zero_values,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UpdateConfigurationRequest) Reset() {
	*x = UpdateConfigurationRequest{}
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigurationRequest) ProtoMessage() {}

func (x *UpdateConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigurationRequest.ProtoReflect.Descriptor instead.
func (*UpdateConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateConfigurationRequest) GetThrottlerName() string {
	if x != nil {
		return x.ThrottlerName
	}
	return ""
}

func (x *UpdateConfigurationRequest) GetConfiguration() *Configuration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *UpdateConfigurationRequest) GetCopyZeroValues() bool {
	if x != nil {
		return x.CopyZeroValues
	}
	return false
}

// UpdateConfigurationResponse is returned by the UpdateConfiguration RPC.
type UpdateConfigurationResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// names is the list of throttler names which were updated.
	Names         []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateConfigurationResponse) Reset() {
	*x = UpdateConfigurationResponse{}
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigurationResponse) ProtoMessage() {}

func (x *UpdateConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigurationResponse.ProtoReflect.Descriptor instead.
func (*UpdateConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateConfigurationResponse) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

// ResetConfigurationRequest is the payload for the ResetConfiguration RPC.
type ResetConfigurationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// throttler_name specifies which throttler to reset. If empty, all active
	// throttlers will be reset.
	ThrottlerName string `protobuf:"bytes,1,opt,name=throttler_name,json=throttlerName,proto3" json:"throttler_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResetConfigurationRequest) Reset() {
	*x = ResetConfigurationRequest{}
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResetConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetConfigurationRequest) ProtoMessage() {}

func (x *ResetConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetConfigurationRequest.ProtoReflect.Descriptor instead.
func (*ResetConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescGZIP(), []int{9}
}

func (x *ResetConfigurationRequest) GetThrottlerName() string {
	if x != nil {
		return x.ThrottlerName
	}
	return ""
}

// ResetConfigurationResponse is returned by the ResetConfiguration RPC.
type ResetConfigurationResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// names is the list of throttler names which were updated.
	Names         []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResetConfigurationResponse) Reset() {
	*x = ResetConfigurationResponse{}
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResetConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetConfigurationResponse) ProtoMessage() {}

func (x *ResetConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetConfigurationResponse.ProtoReflect.Descriptor instead.
func (*ResetConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescGZIP(), []int{10}
}

func (x *ResetConfigurationResponse) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

var File_vitess_throttlerdata_dev_throttlerdata_proto protoreflect.FileDescriptor

const file_vitess_throttlerdata_dev_throttlerdata_proto_rawDesc = "" +
	"\n" +
	",vitess/throttlerdata/dev/throttlerdata.proto\x12\x18vitess.throttlerdata.dev\"\x11\n" +
	"\x0fMaxRatesRequest\"\x99\x01\n" +
	"\x10MaxRatesResponse\x12K\n" +
	"\x05rates\x18\x01 \x03(\v25.vitess.throttlerdata.dev.MaxRatesResponse.RatesEntryR\x05rates\x1a8\n" +
	"\n" +
	"RatesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x03R\x05value:\x028\x01\"'\n" +
	"\x11SetMaxRateRequest\x12\x12\n" +
	"\x04rate\x18\x01 \x01(\x03R\x04rate\"*\n" +
	"\x12SetMaxRateResponse\x12\x14\n" +
	"\x05names\x18\x01 \x03(\tR\x05names\"\xaa\x06\n" +
	"\rConfiguration\x12;\n" +
	"\x1atarget_replication_lag_sec\x18\x01 \x01(\x03R\x17targetReplicationLagSec\x125\n" +
	"\x17max_replication_lag_sec\x18\x02 \x01(\x03R\x14maxReplicationLagSec\x12!\n" +
	"\finitial_rate\x18\x03 \x01(\x03R\vinitialRate\x12!\n" +
	"\fmax_increase\x18\x04 \x01(\x01R\vmaxIncrease\x12-\n" +
	"\x12emergency_decrease\x18\x05 \x01(\x01R\x11emergencyDecrease\x12J\n" +
	"\"min_duration_between_increases_sec\x18\x06 \x01(\x03R\x1eminDurationBetweenIncreasesSec\x12J\n" +
	"\"max_duration_between_increases_sec\x18\a \x01(\x03R\x1emaxDurationBetweenIncreasesSec\x12J\n" +
	"\"min_duration_between_decreases_sec\x18\b \x01(\x03R\x1eminDurationBetweenDecreasesSec\x129\n" +
	"\x19spread_backlog_across_sec\x18\t \x01(\x03R\x16spreadBacklogAcrossSec\x129\n" +
	"\x19ignore_n_slowest_replicas\x18\n" +
	" \x01(\x05R\x16ignoreNSlowestReplicas\x127\n" +
	"\x18ignore_n_slowest_rdonlys\x18\v \x01(\x05R\x15ignoreNSlowestRdonlys\x122\n" +
	"\x16age_bad_rate_after_sec\x18\f \x01(\x03R\x12ageBadRateAfterSec\x12*\n" +
	"\x11bad_rate_increase\x18\r \x01(\x01R\x0fbadRateIncrease\x12=\n" +
	"\x1bmax_rate_approach_threshold\x18\x0e \x01(\x01R\x18maxRateApproachThreshold\"@\n" +
	"\x17GetConfigurationRequest\x12%\n" +
	"\x0ethrottler_name\x18\x01 \x01(\tR\rthrottlerName\"\xf6\x01\n" +
	"\x18GetConfigurationResponse\x12n\n" +
	"\x0econfigurations\x18\x01 \x03(\v2F.vitess.throttlerdata.dev.GetConfigurationResponse.ConfigurationsEntryR\x0econfigurations\x1aj\n" +
	"\x13ConfigurationsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12=\n" +
	"\x05value\x18\x02 \x01(\v2'.vitess.throttlerdata.dev.ConfigurationR\x05value:\x028\x01\"\xbc\x01\n" +
	"\x1aUpdateConfigurationRequest\x12%\n" +
	"\x0ethrottler_name\x18\x01 \x01(\tR\rthrottlerName\x12M\n" +
	"\rconfiguration\x18\x02 \x01(\v2'.vitess.throttlerdata.dev.ConfigurationR\rconfiguration\x12(\n" +
	"\x10copy_zero_values\x18\x03 \x01(\bR\x0ecopyZeroValues\"3\n" +
	"\x1bUpdateConfigurationResponse\x12\x14\n" +
	"\x05names\x18\x01 \x03(\tR\x05names\"B\n" +
	"\x19ResetConfigurationRequest\x12%\n" +
	"\x0ethrottler_name\x18\x01 \x01(\tR\rthrottlerName\"2\n" +
	"\x1aResetConfigurationResponse\x12\x14\n" +
	"\x05names\x18\x01 \x03(\tR\x05namesBSZQgithub.com/planetscale/vitess-types/gen/vitess/throttlerdata/dev;throttlerdatadevb\x06proto3"

var (
	file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescOnce sync.Once
	file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescData []byte
)

func file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescGZIP() []byte {
	file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescOnce.Do(func() {
		file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_vitess_throttlerdata_dev_throttlerdata_proto_rawDesc), len(file_vitess_throttlerdata_dev_throttlerdata_proto_rawDesc)))
	})
	return file_vitess_throttlerdata_dev_throttlerdata_proto_rawDescData
}

var file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_vitess_throttlerdata_dev_throttlerdata_proto_goTypes = []any{
	(*MaxRatesRequest)(nil),             // 0: vitess.throttlerdata.dev.MaxRatesRequest
	(*MaxRatesResponse)(nil),            // 1: vitess.throttlerdata.dev.MaxRatesResponse
	(*SetMaxRateRequest)(nil),           // 2: vitess.throttlerdata.dev.SetMaxRateRequest
	(*SetMaxRateResponse)(nil),          // 3: vitess.throttlerdata.dev.SetMaxRateResponse
	(*Configuration)(nil),               // 4: vitess.throttlerdata.dev.Configuration
	(*GetConfigurationRequest)(nil),     // 5: vitess.throttlerdata.dev.GetConfigurationRequest
	(*GetConfigurationResponse)(nil),    // 6: vitess.throttlerdata.dev.GetConfigurationResponse
	(*UpdateConfigurationRequest)(nil),  // 7: vitess.throttlerdata.dev.UpdateConfigurationRequest
	(*UpdateConfigurationResponse)(nil), // 8: vitess.throttlerdata.dev.UpdateConfigurationResponse
	(*ResetConfigurationRequest)(nil),   // 9: vitess.throttlerdata.dev.ResetConfigurationRequest
	(*ResetConfigurationResponse)(nil),  // 10: vitess.throttlerdata.dev.ResetConfigurationResponse
	nil,                                 // 11: vitess.throttlerdata.dev.MaxRatesResponse.RatesEntry
	nil,                                 // 12: vitess.throttlerdata.dev.GetConfigurationResponse.ConfigurationsEntry
}
var file_vitess_throttlerdata_dev_throttlerdata_proto_depIdxs = []int32{
	11, // 0: vitess.throttlerdata.dev.MaxRatesResponse.rates:type_name -> vitess.throttlerdata.dev.MaxRatesResponse.RatesEntry
	12, // 1: vitess.throttlerdata.dev.GetConfigurationResponse.configurations:type_name -> vitess.throttlerdata.dev.GetConfigurationResponse.ConfigurationsEntry
	4,  // 2: vitess.throttlerdata.dev.UpdateConfigurationRequest.configuration:type_name -> vitess.throttlerdata.dev.Configuration
	4,  // 3: vitess.throttlerdata.dev.GetConfigurationResponse.ConfigurationsEntry.value:type_name -> vitess.throttlerdata.dev.Configuration
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_vitess_throttlerdata_dev_throttlerdata_proto_init() }
func file_vitess_throttlerdata_dev_throttlerdata_proto_init() {
	if File_vitess_throttlerdata_dev_throttlerdata_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_vitess_throttlerdata_dev_throttlerdata_proto_rawDesc), len(file_vitess_throttlerdata_dev_throttlerdata_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vitess_throttlerdata_dev_throttlerdata_proto_goTypes,
		DependencyIndexes: file_vitess_throttlerdata_dev_throttlerdata_proto_depIdxs,
		MessageInfos:      file_vitess_throttlerdata_dev_throttlerdata_proto_msgTypes,
	}.Build()
	File_vitess_throttlerdata_dev_throttlerdata_proto = out.File
	file_vitess_throttlerdata_dev_throttlerdata_proto_goTypes = nil
	file_vitess_throttlerdata_dev_throttlerdata_proto_depIdxs = nil
}
