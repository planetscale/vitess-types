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

// gRPC RPC interface for the internal resharding throttler (go/vt/throttler)
// which is used by vreplication.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: vitess/throttlerservice/v18/throttlerservice.proto

package throttlerservicev18

import (
	v18 "github.com/planetscale/vitess-types/gen/vitess/throttlerdata/v18"
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

var File_vitess_throttlerservice_v18_throttlerservice_proto protoreflect.FileDescriptor

var file_vitess_throttlerservice_v18_throttlerservice_proto_rawDesc = []byte{
	0x0a, 0x32, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x74, 0x68,
	0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x68, 0x72,
	0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x38, 0x1a, 0x2c, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74,
	0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x38, 0x2f, 0x74, 0x68, 0x72, 0x6f,
	0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xe3, 0x04, 0x0a, 0x09, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x12, 0x63, 0x0a,
	0x08, 0x4d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x29, 0x2e, 0x76, 0x69, 0x74, 0x65,
	0x73, 0x73, 0x2e, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x38, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x68,
	0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x38, 0x2e,
	0x4d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x69, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x2b, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74,
	0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x53, 0x65, 0x74, 0x4d,
	0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7b, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x31, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x68, 0x72, 0x6f, 0x74,
	0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x68,
	0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x38, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x84, 0x01, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x34, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x68, 0x72, 0x6f,
	0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x38, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x81, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x38, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e,
	0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x38, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x59, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f,
	0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65,
	0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x38, 0x3b, 0x74, 0x68, 0x72,
	0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x38,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_vitess_throttlerservice_v18_throttlerservice_proto_goTypes = []any{
	(*v18.MaxRatesRequest)(nil),             // 0: vitess.throttlerdata.v18.MaxRatesRequest
	(*v18.SetMaxRateRequest)(nil),           // 1: vitess.throttlerdata.v18.SetMaxRateRequest
	(*v18.GetConfigurationRequest)(nil),     // 2: vitess.throttlerdata.v18.GetConfigurationRequest
	(*v18.UpdateConfigurationRequest)(nil),  // 3: vitess.throttlerdata.v18.UpdateConfigurationRequest
	(*v18.ResetConfigurationRequest)(nil),   // 4: vitess.throttlerdata.v18.ResetConfigurationRequest
	(*v18.MaxRatesResponse)(nil),            // 5: vitess.throttlerdata.v18.MaxRatesResponse
	(*v18.SetMaxRateResponse)(nil),          // 6: vitess.throttlerdata.v18.SetMaxRateResponse
	(*v18.GetConfigurationResponse)(nil),    // 7: vitess.throttlerdata.v18.GetConfigurationResponse
	(*v18.UpdateConfigurationResponse)(nil), // 8: vitess.throttlerdata.v18.UpdateConfigurationResponse
	(*v18.ResetConfigurationResponse)(nil),  // 9: vitess.throttlerdata.v18.ResetConfigurationResponse
}
var file_vitess_throttlerservice_v18_throttlerservice_proto_depIdxs = []int32{
	0, // 0: vitess.throttlerservice.v18.Throttler.MaxRates:input_type -> vitess.throttlerdata.v18.MaxRatesRequest
	1, // 1: vitess.throttlerservice.v18.Throttler.SetMaxRate:input_type -> vitess.throttlerdata.v18.SetMaxRateRequest
	2, // 2: vitess.throttlerservice.v18.Throttler.GetConfiguration:input_type -> vitess.throttlerdata.v18.GetConfigurationRequest
	3, // 3: vitess.throttlerservice.v18.Throttler.UpdateConfiguration:input_type -> vitess.throttlerdata.v18.UpdateConfigurationRequest
	4, // 4: vitess.throttlerservice.v18.Throttler.ResetConfiguration:input_type -> vitess.throttlerdata.v18.ResetConfigurationRequest
	5, // 5: vitess.throttlerservice.v18.Throttler.MaxRates:output_type -> vitess.throttlerdata.v18.MaxRatesResponse
	6, // 6: vitess.throttlerservice.v18.Throttler.SetMaxRate:output_type -> vitess.throttlerdata.v18.SetMaxRateResponse
	7, // 7: vitess.throttlerservice.v18.Throttler.GetConfiguration:output_type -> vitess.throttlerdata.v18.GetConfigurationResponse
	8, // 8: vitess.throttlerservice.v18.Throttler.UpdateConfiguration:output_type -> vitess.throttlerdata.v18.UpdateConfigurationResponse
	9, // 9: vitess.throttlerservice.v18.Throttler.ResetConfiguration:output_type -> vitess.throttlerdata.v18.ResetConfigurationResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vitess_throttlerservice_v18_throttlerservice_proto_init() }
func file_vitess_throttlerservice_v18_throttlerservice_proto_init() {
	if File_vitess_throttlerservice_v18_throttlerservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vitess_throttlerservice_v18_throttlerservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vitess_throttlerservice_v18_throttlerservice_proto_goTypes,
		DependencyIndexes: file_vitess_throttlerservice_v18_throttlerservice_proto_depIdxs,
	}.Build()
	File_vitess_throttlerservice_v18_throttlerservice_proto = out.File
	file_vitess_throttlerservice_v18_throttlerservice_proto_rawDesc = nil
	file_vitess_throttlerservice_v18_throttlerservice_proto_goTypes = nil
	file_vitess_throttlerservice_v18_throttlerservice_proto_depIdxs = nil
}
