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

// This file contains useful data structures for RPCs in Vitess.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: vitess/vtrpc/v21/vtrpc.proto

// option java_package="io.vitess.proto";

package vtrpcv21

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

// Code represents canonical error codes. The names, numbers and comments
// must match the ones defined by grpc (0-16):
//
//	https://godoc.org/google.golang.org/grpc/codes.
//
// 17+ are custom codes
type Code int32

const (
	// OK is returned on success.
	Code_OK Code = 0
	// CANCELED indicates the operation was cancelled (typically by the caller).
	Code_CANCELED Code = 1
	// UNKNOWN error. An example of where this error may be returned is
	// if a Status value received from another address space belongs to
	// an error-space that is not known in this address space. Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	Code_UNKNOWN Code = 2
	// INVALID_ARGUMENT indicates client specified an invalid argument.
	// Note that this differs from FAILED_PRECONDITION. It indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	Code_INVALID_ARGUMENT Code = 3
	// DEADLINE_EXCEEDED means operation expired before completion.
	// For operations that change the state of the system, this error may be
	// returned even if the operation has completed successfully. For
	// example, a successful response from a server could have been delayed
	// long enough for the deadline to expire.
	Code_DEADLINE_EXCEEDED Code = 4
	// NOT_FOUND means some requested entity (e.g., file or directory) was
	// not found.
	Code_NOT_FOUND Code = 5
	// ALREADY_EXISTS means an attempt to create an entity failed because one
	// already exists.
	Code_ALREADY_EXISTS Code = 6
	// PERMISSION_DENIED indicates the caller does not have permission to
	// execute the specified operation. It must not be used for rejections
	// caused by exhausting some resource (use RESOURCE_EXHAUSTED
	// instead for those errors).  It must not be
	// used if the caller cannot be identified (use Unauthenticated
	// instead for those errors).
	Code_PERMISSION_DENIED Code = 7
	// RESOURCE_EXHAUSTED indicates some resource has been exhausted, perhaps
	// a per-user quota, or perhaps the entire file system is out of space.
	Code_RESOURCE_EXHAUSTED Code = 8
	// FAILED_PRECONDITION indicates operation was rejected because the
	// system is not in a state required for the operation's execution.
	// For example, directory to be deleted may be non-empty, an rmdir
	// operation is applied to a non-directory, etc.
	//
	// A litmus test that may help a service implementor in deciding
	// between FAILED_PRECONDITION, ABORTED, and UNAVAILABLE:
	//
	//	(a) Use UNAVAILABLE if the client can retry just the failing call.
	//	(b) Use ABORTED if the client should retry at a higher-level
	//	    (e.g., restarting a read-modify-write sequence).
	//	(c) Use FAILED_PRECONDITION if the client should not retry until
	//	    the system state has been explicitly fixed.  E.g., if an "rmdir"
	//	    fails because the directory is non-empty, FAILED_PRECONDITION
	//	    should be returned since the client should not retry unless
	//	    they have first fixed up the directory by deleting files from it.
	//	(d) Use FAILED_PRECONDITION if the client performs conditional
	//	    REST Get/Update/Delete on a resource and the resource on the
	//	    server does not match the condition. E.g., conflicting
	//	    read-modify-write on the same resource.
	Code_FAILED_PRECONDITION Code = 9
	// ABORTED indicates the operation was aborted, typically due to a
	// concurrency issue like sequencer check failures, transaction aborts,
	// etc.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Code_ABORTED Code = 10
	// OUT_OF_RANGE means operation was attempted past the valid range.
	// E.g., seeking or reading past end of file.
	//
	// Unlike INVALID_ARGUMENT, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate INVALID_ARGUMENT if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// OUT_OF_RANGE if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between FAILED_PRECONDITION and
	// OUT_OF_RANGE.  We recommend using OUT_OF_RANGE (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an OUT_OF_RANGE error to detect when
	// they are done.
	Code_OUT_OF_RANGE Code = 11
	// UNIMPLEMENTED indicates operation is not implemented or not
	// supported/enabled in this service.
	Code_UNIMPLEMENTED Code = 12
	// INTERNAL errors. Means some invariants expected by underlying
	// system has been broken.  If you see one of these errors,
	// something is very broken.
	Code_INTERNAL Code = 13
	// UNAVAILABLE indicates the service is currently unavailable.
	// This is a most likely a transient condition and may be corrected
	// by retrying with a backoff.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Code_UNAVAILABLE Code = 14
	// DATA_LOSS indicates unrecoverable data loss or corruption.
	Code_DATA_LOSS Code = 15
	// UNAUTHENTICATED indicates the request does not have valid
	// authentication credentials for the operation.
	Code_UNAUTHENTICATED Code = 16
	// CLUSTER_EVENT indicates that a cluster operation might be in effect
	Code_CLUSTER_EVENT Code = 17
	// Topo server connection is read-only
	Code_READ_ONLY Code = 18
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:  "OK",
		1:  "CANCELED",
		2:  "UNKNOWN",
		3:  "INVALID_ARGUMENT",
		4:  "DEADLINE_EXCEEDED",
		5:  "NOT_FOUND",
		6:  "ALREADY_EXISTS",
		7:  "PERMISSION_DENIED",
		8:  "RESOURCE_EXHAUSTED",
		9:  "FAILED_PRECONDITION",
		10: "ABORTED",
		11: "OUT_OF_RANGE",
		12: "UNIMPLEMENTED",
		13: "INTERNAL",
		14: "UNAVAILABLE",
		15: "DATA_LOSS",
		16: "UNAUTHENTICATED",
		17: "CLUSTER_EVENT",
		18: "READ_ONLY",
	}
	Code_value = map[string]int32{
		"OK":                  0,
		"CANCELED":            1,
		"UNKNOWN":             2,
		"INVALID_ARGUMENT":    3,
		"DEADLINE_EXCEEDED":   4,
		"NOT_FOUND":           5,
		"ALREADY_EXISTS":      6,
		"PERMISSION_DENIED":   7,
		"RESOURCE_EXHAUSTED":  8,
		"FAILED_PRECONDITION": 9,
		"ABORTED":             10,
		"OUT_OF_RANGE":        11,
		"UNIMPLEMENTED":       12,
		"INTERNAL":            13,
		"UNAVAILABLE":         14,
		"DATA_LOSS":           15,
		"UNAUTHENTICATED":     16,
		"CLUSTER_EVENT":       17,
		"READ_ONLY":           18,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_vitess_vtrpc_v21_vtrpc_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_vitess_vtrpc_v21_vtrpc_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_vitess_vtrpc_v21_vtrpc_proto_rawDescGZIP(), []int{0}
}

// CallerID is passed along RPCs to identify the originating client
// for a request. It is not meant to be secure, but only
// informational.  The client can put whatever info they want in these
// fields, and they will be trusted by the servers. The fields will
// just be used for logging purposes, and to easily find a client.
// VtGate propagates it to VtTablet, and VtTablet may use this
// information for monitoring purposes, to display on dashboards, or
// for denying access to tables during a migration.
type CallerID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// principal is the effective user identifier. It is usually filled in
	// with whoever made the request to the appserver, if the request
	// came from an automated job or another system component.
	// If the request comes directly from the Internet, or if the Vitess client
	// takes action on its own accord, it is okay for this field to be absent.
	Principal string `protobuf:"bytes,1,opt,name=principal,proto3" json:"principal,omitempty"`
	// component describes the running process of the effective caller.
	// It can for instance be the hostname:port of the servlet initiating the
	// database call, or the container engine ID used by the servlet.
	Component string `protobuf:"bytes,2,opt,name=component,proto3" json:"component,omitempty"`
	// subcomponent describes a component inisde the immediate caller which
	// is responsible for generating is request. Suggested values are a
	// servlet name or an API endpoint name.
	Subcomponent string `protobuf:"bytes,3,opt,name=subcomponent,proto3" json:"subcomponent,omitempty"`
	// set of security groups that should be assigned to this caller.
	Groups []string `protobuf:"bytes,4,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *CallerID) Reset() {
	*x = CallerID{}
	mi := &file_vitess_vtrpc_v21_vtrpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CallerID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallerID) ProtoMessage() {}

func (x *CallerID) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_vtrpc_v21_vtrpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallerID.ProtoReflect.Descriptor instead.
func (*CallerID) Descriptor() ([]byte, []int) {
	return file_vitess_vtrpc_v21_vtrpc_proto_rawDescGZIP(), []int{0}
}

func (x *CallerID) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *CallerID) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *CallerID) GetSubcomponent() string {
	if x != nil {
		return x.Subcomponent
	}
	return ""
}

func (x *CallerID) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

// RPCError is an application-level error structure returned by
// VtTablet (and passed along by VtGate if appropriate).
// We use this so the clients don't have to parse the error messages,
// but instead can depend on the value of the code.
type RPCError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Code    Code   `protobuf:"varint,3,opt,name=code,proto3,enum=vitess.vtrpc.v21.Code" json:"code,omitempty"`
}

func (x *RPCError) Reset() {
	*x = RPCError{}
	mi := &file_vitess_vtrpc_v21_vtrpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RPCError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCError) ProtoMessage() {}

func (x *RPCError) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_vtrpc_v21_vtrpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCError.ProtoReflect.Descriptor instead.
func (*RPCError) Descriptor() ([]byte, []int) {
	return file_vitess_vtrpc_v21_vtrpc_proto_rawDescGZIP(), []int{1}
}

func (x *RPCError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RPCError) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

var File_vitess_vtrpc_v21_vtrpc_proto protoreflect.FileDescriptor

var file_vitess_vtrpc_v21_vtrpc_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x76,
	0x32, 0x31, 0x2f, 0x76, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x32, 0x31,
	0x22, 0x82, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x62,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x75, 0x62, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x63, 0x0a, 0x08, 0x52, 0x50, 0x43, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x76, 0x69, 0x74, 0x65,
	0x73, 0x73, 0x2e, 0x76, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x32, 0x31, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x0b, 0x6c,
	0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2a, 0xd8, 0x02, 0x0a, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11,
	0x44, 0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45,
	0x44, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44,
	0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58,
	0x49, 0x53, 0x54, 0x53, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x07, 0x12, 0x16, 0x0a,
	0x12, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x58, 0x48, 0x41, 0x55, 0x53,
	0x54, 0x45, 0x44, 0x10, 0x08, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f,
	0x50, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x0b,
	0x0a, 0x07, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x4f,
	0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x0b, 0x12, 0x11, 0x0a,
	0x0d, 0x55, 0x4e, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x0c,
	0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x0d, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x0e, 0x12,
	0x0d, 0x0a, 0x09, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4c, 0x4f, 0x53, 0x53, 0x10, 0x0f, 0x12, 0x13,
	0x0a, 0x0f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x10, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x10, 0x11, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x4f,
	0x4e, 0x4c, 0x59, 0x10, 0x12, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f,
	0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x74, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x32,
	0x31, 0x3b, 0x76, 0x74, 0x72, 0x70, 0x63, 0x76, 0x32, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_vitess_vtrpc_v21_vtrpc_proto_rawDescOnce sync.Once
	file_vitess_vtrpc_v21_vtrpc_proto_rawDescData = file_vitess_vtrpc_v21_vtrpc_proto_rawDesc
)

func file_vitess_vtrpc_v21_vtrpc_proto_rawDescGZIP() []byte {
	file_vitess_vtrpc_v21_vtrpc_proto_rawDescOnce.Do(func() {
		file_vitess_vtrpc_v21_vtrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_vitess_vtrpc_v21_vtrpc_proto_rawDescData)
	})
	return file_vitess_vtrpc_v21_vtrpc_proto_rawDescData
}

var file_vitess_vtrpc_v21_vtrpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vitess_vtrpc_v21_vtrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vitess_vtrpc_v21_vtrpc_proto_goTypes = []any{
	(Code)(0),        // 0: vitess.vtrpc.v21.Code
	(*CallerID)(nil), // 1: vitess.vtrpc.v21.CallerID
	(*RPCError)(nil), // 2: vitess.vtrpc.v21.RPCError
}
var file_vitess_vtrpc_v21_vtrpc_proto_depIdxs = []int32{
	0, // 0: vitess.vtrpc.v21.RPCError.code:type_name -> vitess.vtrpc.v21.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vitess_vtrpc_v21_vtrpc_proto_init() }
func file_vitess_vtrpc_v21_vtrpc_proto_init() {
	if File_vitess_vtrpc_v21_vtrpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vitess_vtrpc_v21_vtrpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vitess_vtrpc_v21_vtrpc_proto_goTypes,
		DependencyIndexes: file_vitess_vtrpc_v21_vtrpc_proto_depIdxs,
		EnumInfos:         file_vitess_vtrpc_v21_vtrpc_proto_enumTypes,
		MessageInfos:      file_vitess_vtrpc_v21_vtrpc_proto_msgTypes,
	}.Build()
	File_vitess_vtrpc_v21_vtrpc_proto = out.File
	file_vitess_vtrpc_v21_vtrpc_proto_rawDesc = nil
	file_vitess_vtrpc_v21_vtrpc_proto_goTypes = nil
	file_vitess_vtrpc_v21_vtrpc_proto_depIdxs = nil
}
