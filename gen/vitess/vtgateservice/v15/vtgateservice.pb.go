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

// Service definition for vtgateservice.
// This is the main entry point to Vitess.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: vitess/vtgateservice/v15/vtgateservice.proto

// option java_package="io.vitess.proto.grpc";

package vtgateservicev15

import (
	v15 "github.com/planetscale/vitess-types/gen/vitess/vtgate/v15"
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

var File_vitess_vtgateservice_v15_vtgateservice_proto protoreflect.FileDescriptor

var file_vitess_vtgateservice_v15_vtgateservice_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x76, 0x74, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x35, 0x1a, 0x1e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73,
	0x2f, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x76, 0x74, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa9, 0x05, 0x0a, 0x06, 0x56, 0x69, 0x74,
	0x65, 0x73, 0x73, 0x12, 0x52, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x21,
	0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x35, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0c, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x26, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73,
	0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x35, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x0d, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x76, 0x69,
	0x74, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x35, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x74,
	0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x73, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e,
	0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x07, 0x56, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x21, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x74, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x76,
	0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x52, 0x0a,
	0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x21, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x69,
	0x74, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x35, 0x2e,
	0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x61, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x26, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x76, 0x69, 0x74, 0x65,
	0x73, 0x73, 0x2e, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76,
	0x69, 0x74, 0x65, 0x73, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x35, 0x3b, 0x76, 0x74, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x35, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_vitess_vtgateservice_v15_vtgateservice_proto_goTypes = []interface{}{
	(*v15.ExecuteRequest)(nil),             // 0: vitess.vtgate.v15.ExecuteRequest
	(*v15.ExecuteBatchRequest)(nil),        // 1: vitess.vtgate.v15.ExecuteBatchRequest
	(*v15.StreamExecuteRequest)(nil),       // 2: vitess.vtgate.v15.StreamExecuteRequest
	(*v15.ResolveTransactionRequest)(nil),  // 3: vitess.vtgate.v15.ResolveTransactionRequest
	(*v15.VStreamRequest)(nil),             // 4: vitess.vtgate.v15.VStreamRequest
	(*v15.PrepareRequest)(nil),             // 5: vitess.vtgate.v15.PrepareRequest
	(*v15.CloseSessionRequest)(nil),        // 6: vitess.vtgate.v15.CloseSessionRequest
	(*v15.ExecuteResponse)(nil),            // 7: vitess.vtgate.v15.ExecuteResponse
	(*v15.ExecuteBatchResponse)(nil),       // 8: vitess.vtgate.v15.ExecuteBatchResponse
	(*v15.StreamExecuteResponse)(nil),      // 9: vitess.vtgate.v15.StreamExecuteResponse
	(*v15.ResolveTransactionResponse)(nil), // 10: vitess.vtgate.v15.ResolveTransactionResponse
	(*v15.VStreamResponse)(nil),            // 11: vitess.vtgate.v15.VStreamResponse
	(*v15.PrepareResponse)(nil),            // 12: vitess.vtgate.v15.PrepareResponse
	(*v15.CloseSessionResponse)(nil),       // 13: vitess.vtgate.v15.CloseSessionResponse
}
var file_vitess_vtgateservice_v15_vtgateservice_proto_depIdxs = []int32{
	0,  // 0: vitess.vtgateservice.v15.Vitess.Execute:input_type -> vitess.vtgate.v15.ExecuteRequest
	1,  // 1: vitess.vtgateservice.v15.Vitess.ExecuteBatch:input_type -> vitess.vtgate.v15.ExecuteBatchRequest
	2,  // 2: vitess.vtgateservice.v15.Vitess.StreamExecute:input_type -> vitess.vtgate.v15.StreamExecuteRequest
	3,  // 3: vitess.vtgateservice.v15.Vitess.ResolveTransaction:input_type -> vitess.vtgate.v15.ResolveTransactionRequest
	4,  // 4: vitess.vtgateservice.v15.Vitess.VStream:input_type -> vitess.vtgate.v15.VStreamRequest
	5,  // 5: vitess.vtgateservice.v15.Vitess.Prepare:input_type -> vitess.vtgate.v15.PrepareRequest
	6,  // 6: vitess.vtgateservice.v15.Vitess.CloseSession:input_type -> vitess.vtgate.v15.CloseSessionRequest
	7,  // 7: vitess.vtgateservice.v15.Vitess.Execute:output_type -> vitess.vtgate.v15.ExecuteResponse
	8,  // 8: vitess.vtgateservice.v15.Vitess.ExecuteBatch:output_type -> vitess.vtgate.v15.ExecuteBatchResponse
	9,  // 9: vitess.vtgateservice.v15.Vitess.StreamExecute:output_type -> vitess.vtgate.v15.StreamExecuteResponse
	10, // 10: vitess.vtgateservice.v15.Vitess.ResolveTransaction:output_type -> vitess.vtgate.v15.ResolveTransactionResponse
	11, // 11: vitess.vtgateservice.v15.Vitess.VStream:output_type -> vitess.vtgate.v15.VStreamResponse
	12, // 12: vitess.vtgateservice.v15.Vitess.Prepare:output_type -> vitess.vtgate.v15.PrepareResponse
	13, // 13: vitess.vtgateservice.v15.Vitess.CloseSession:output_type -> vitess.vtgate.v15.CloseSessionResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_vitess_vtgateservice_v15_vtgateservice_proto_init() }
func file_vitess_vtgateservice_v15_vtgateservice_proto_init() {
	if File_vitess_vtgateservice_v15_vtgateservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vitess_vtgateservice_v15_vtgateservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vitess_vtgateservice_v15_vtgateservice_proto_goTypes,
		DependencyIndexes: file_vitess_vtgateservice_v15_vtgateservice_proto_depIdxs,
	}.Build()
	File_vitess_vtgateservice_v15_vtgateservice_proto = out.File
	file_vitess_vtgateservice_v15_vtgateservice_proto_rawDesc = nil
	file_vitess_vtgateservice_v15_vtgateservice_proto_goTypes = nil
	file_vitess_vtgateservice_v15_vtgateservice_proto_depIdxs = nil
}
