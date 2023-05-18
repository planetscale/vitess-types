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

// This file contains the UpdateStream service definition, necessary
// to make RPC calls to VtTablet for the binlog protocol, used by
// filtered replication only.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: vitess/binlogservice/v16/binlogservice.proto

package binlogservicev16

import (
	v16 "github.com/planetscale/vitess-types/gen/vitess/binlogdata/v16"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_vitess_binlogservice_v16_binlogservice_proto protoreflect.FileDescriptor

var file_vitess_binlogservice_v16_binlogservice_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x36, 0x2f, 0x62, 0x69, 0x6e, 0x6c, 0x6f,
	0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x36, 0x1a, 0x26, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73,
	0x2f, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x36, 0x2f,
	0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xee, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x71, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4b, 0x65, 0x79, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x2c, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x62, 0x69, 0x6e,
	0x6c, 0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x62, 0x69, 0x6e, 0x6c, 0x6f,
	0x67, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4b, 0x65, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x6b, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x12, 0x2a, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x62, 0x69,
	0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x36, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x69, 0x74, 0x65,
	0x73, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x69, 0x74,
	0x65, 0x73, 0x73, 0x2f, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x36, 0x3b, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x76, 0x31, 0x36, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_vitess_binlogservice_v16_binlogservice_proto_goTypes = []interface{}{
	(*v16.StreamKeyRangeRequest)(nil),  // 0: vitess.binlogdata.v16.StreamKeyRangeRequest
	(*v16.StreamTablesRequest)(nil),    // 1: vitess.binlogdata.v16.StreamTablesRequest
	(*v16.StreamKeyRangeResponse)(nil), // 2: vitess.binlogdata.v16.StreamKeyRangeResponse
	(*v16.StreamTablesResponse)(nil),   // 3: vitess.binlogdata.v16.StreamTablesResponse
}
var file_vitess_binlogservice_v16_binlogservice_proto_depIdxs = []int32{
	0, // 0: vitess.binlogservice.v16.UpdateStream.StreamKeyRange:input_type -> vitess.binlogdata.v16.StreamKeyRangeRequest
	1, // 1: vitess.binlogservice.v16.UpdateStream.StreamTables:input_type -> vitess.binlogdata.v16.StreamTablesRequest
	2, // 2: vitess.binlogservice.v16.UpdateStream.StreamKeyRange:output_type -> vitess.binlogdata.v16.StreamKeyRangeResponse
	3, // 3: vitess.binlogservice.v16.UpdateStream.StreamTables:output_type -> vitess.binlogdata.v16.StreamTablesResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vitess_binlogservice_v16_binlogservice_proto_init() }
func file_vitess_binlogservice_v16_binlogservice_proto_init() {
	if File_vitess_binlogservice_v16_binlogservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vitess_binlogservice_v16_binlogservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vitess_binlogservice_v16_binlogservice_proto_goTypes,
		DependencyIndexes: file_vitess_binlogservice_v16_binlogservice_proto_depIdxs,
	}.Build()
	File_vitess_binlogservice_v16_binlogservice_proto = out.File
	file_vitess_binlogservice_v16_binlogservice_proto_rawDesc = nil
	file_vitess_binlogservice_v16_binlogservice_proto_goTypes = nil
	file_vitess_binlogservice_v16_binlogservice_proto_depIdxs = nil
}
