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

// Data structures for the vtworker RPC interface.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: vtworkerdata/vtworkerdata.proto

package vtworkerdata

import (
	logutil "github.com/planetscale/vitess-types/gen/vitess/v14/logutil"
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

// ExecuteVtworkerCommandRequest is the payload for ExecuteVtworkerCommand.
type ExecuteVtworkerCommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Args []string `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *ExecuteVtworkerCommandRequest) Reset() {
	*x = ExecuteVtworkerCommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vtworkerdata_vtworkerdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteVtworkerCommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteVtworkerCommandRequest) ProtoMessage() {}

func (x *ExecuteVtworkerCommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vtworkerdata_vtworkerdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteVtworkerCommandRequest.ProtoReflect.Descriptor instead.
func (*ExecuteVtworkerCommandRequest) Descriptor() ([]byte, []int) {
	return file_vtworkerdata_vtworkerdata_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteVtworkerCommandRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

// ExecuteVtworkerCommandResponse is streamed back by ExecuteVtworkerCommand.
type ExecuteVtworkerCommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *logutil.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *ExecuteVtworkerCommandResponse) Reset() {
	*x = ExecuteVtworkerCommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vtworkerdata_vtworkerdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteVtworkerCommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteVtworkerCommandResponse) ProtoMessage() {}

func (x *ExecuteVtworkerCommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vtworkerdata_vtworkerdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteVtworkerCommandResponse.ProtoReflect.Descriptor instead.
func (*ExecuteVtworkerCommandResponse) Descriptor() ([]byte, []int) {
	return file_vtworkerdata_vtworkerdata_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteVtworkerCommandResponse) GetEvent() *logutil.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_vtworkerdata_vtworkerdata_proto protoreflect.FileDescriptor

var file_vtworkerdata_vtworkerdata_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x76, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x76, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x15, 0x6c, 0x6f, 0x67, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x6c, 0x6f, 0x67, 0x75, 0x74, 0x69, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x1d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x56, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0x46, 0x0a, 0x1e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x56, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c,
	0x6f, 0x67, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0xb6, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x42, 0x11, 0x56, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2d, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31,
	0x34, 0x2f, 0x76, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0xa2, 0x02,
	0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x56, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x64,
	0x61, 0x74, 0x61, 0xca, 0x02, 0x0c, 0x56, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x64, 0x61,
	0x74, 0x61, 0xe2, 0x02, 0x18, 0x56, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x74,
	0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c,
	0x56, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vtworkerdata_vtworkerdata_proto_rawDescOnce sync.Once
	file_vtworkerdata_vtworkerdata_proto_rawDescData = file_vtworkerdata_vtworkerdata_proto_rawDesc
)

func file_vtworkerdata_vtworkerdata_proto_rawDescGZIP() []byte {
	file_vtworkerdata_vtworkerdata_proto_rawDescOnce.Do(func() {
		file_vtworkerdata_vtworkerdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_vtworkerdata_vtworkerdata_proto_rawDescData)
	})
	return file_vtworkerdata_vtworkerdata_proto_rawDescData
}

var file_vtworkerdata_vtworkerdata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vtworkerdata_vtworkerdata_proto_goTypes = []interface{}{
	(*ExecuteVtworkerCommandRequest)(nil),  // 0: vtworkerdata.ExecuteVtworkerCommandRequest
	(*ExecuteVtworkerCommandResponse)(nil), // 1: vtworkerdata.ExecuteVtworkerCommandResponse
	(*logutil.Event)(nil),                  // 2: logutil.Event
}
var file_vtworkerdata_vtworkerdata_proto_depIdxs = []int32{
	2, // 0: vtworkerdata.ExecuteVtworkerCommandResponse.event:type_name -> logutil.Event
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vtworkerdata_vtworkerdata_proto_init() }
func file_vtworkerdata_vtworkerdata_proto_init() {
	if File_vtworkerdata_vtworkerdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vtworkerdata_vtworkerdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteVtworkerCommandRequest); i {
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
		file_vtworkerdata_vtworkerdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteVtworkerCommandResponse); i {
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
			RawDescriptor: file_vtworkerdata_vtworkerdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vtworkerdata_vtworkerdata_proto_goTypes,
		DependencyIndexes: file_vtworkerdata_vtworkerdata_proto_depIdxs,
		MessageInfos:      file_vtworkerdata_vtworkerdata_proto_msgTypes,
	}.Build()
	File_vtworkerdata_vtworkerdata_proto = out.File
	file_vtworkerdata_vtworkerdata_proto_rawDesc = nil
	file_vtworkerdata_vtworkerdata_proto_goTypes = nil
	file_vtworkerdata_vtworkerdata_proto_depIdxs = nil
}
