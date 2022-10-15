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

// gRPC RPC interface for the internal resharding throttler (go/vt/throttler)
// which is used by vreplication.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: throttlerservice/throttlerservice.proto

package throttlerservice

import (
	throttlerdata "github.com/planetscale/vitess-types/gen/vitess/dev/throttlerdata"
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

var File_throttlerservice_throttlerservice_proto protoreflect.FileDescriptor

var file_throttlerservice_throttlerservice_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x68, 0x72, 0x6f, 0x74,
	0x74, 0x6c, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x21, 0x74, 0x68, 0x72,
	0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x74, 0x68, 0x72, 0x6f, 0x74,
	0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf3,
	0x03, 0x0a, 0x09, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x08,
	0x4d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x74, 0x68, 0x72, 0x6f, 0x74,
	0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x68, 0x72, 0x6f, 0x74,
	0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a, 0x53,
	0x65, 0x74, 0x4d, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x74, 0x68, 0x72, 0x6f,
	0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x78,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x68,
	0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x65, 0x74, 0x4d,
	0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x65, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74,
	0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29,
	0x2e, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x74, 0x68, 0x72, 0x6f,
	0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e,
	0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74,
	0x6c, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0xd2, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x72,
	0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x15, 0x54,
	0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76,
	0x69, 0x74, 0x65, 0x73, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x2f, 0x74, 0x68, 0x72, 0x6f, 0x74,
	0x74, 0x6c, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x03, 0x54, 0x58,
	0x58, 0xaa, 0x02, 0x10, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0xca, 0x02, 0x10, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x1c, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74,
	0x6c, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x54, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_throttlerservice_throttlerservice_proto_goTypes = []interface{}{
	(*throttlerdata.MaxRatesRequest)(nil),             // 0: throttlerdata.MaxRatesRequest
	(*throttlerdata.SetMaxRateRequest)(nil),           // 1: throttlerdata.SetMaxRateRequest
	(*throttlerdata.GetConfigurationRequest)(nil),     // 2: throttlerdata.GetConfigurationRequest
	(*throttlerdata.UpdateConfigurationRequest)(nil),  // 3: throttlerdata.UpdateConfigurationRequest
	(*throttlerdata.ResetConfigurationRequest)(nil),   // 4: throttlerdata.ResetConfigurationRequest
	(*throttlerdata.MaxRatesResponse)(nil),            // 5: throttlerdata.MaxRatesResponse
	(*throttlerdata.SetMaxRateResponse)(nil),          // 6: throttlerdata.SetMaxRateResponse
	(*throttlerdata.GetConfigurationResponse)(nil),    // 7: throttlerdata.GetConfigurationResponse
	(*throttlerdata.UpdateConfigurationResponse)(nil), // 8: throttlerdata.UpdateConfigurationResponse
	(*throttlerdata.ResetConfigurationResponse)(nil),  // 9: throttlerdata.ResetConfigurationResponse
}
var file_throttlerservice_throttlerservice_proto_depIdxs = []int32{
	0, // 0: throttlerservice.Throttler.MaxRates:input_type -> throttlerdata.MaxRatesRequest
	1, // 1: throttlerservice.Throttler.SetMaxRate:input_type -> throttlerdata.SetMaxRateRequest
	2, // 2: throttlerservice.Throttler.GetConfiguration:input_type -> throttlerdata.GetConfigurationRequest
	3, // 3: throttlerservice.Throttler.UpdateConfiguration:input_type -> throttlerdata.UpdateConfigurationRequest
	4, // 4: throttlerservice.Throttler.ResetConfiguration:input_type -> throttlerdata.ResetConfigurationRequest
	5, // 5: throttlerservice.Throttler.MaxRates:output_type -> throttlerdata.MaxRatesResponse
	6, // 6: throttlerservice.Throttler.SetMaxRate:output_type -> throttlerdata.SetMaxRateResponse
	7, // 7: throttlerservice.Throttler.GetConfiguration:output_type -> throttlerdata.GetConfigurationResponse
	8, // 8: throttlerservice.Throttler.UpdateConfiguration:output_type -> throttlerdata.UpdateConfigurationResponse
	9, // 9: throttlerservice.Throttler.ResetConfiguration:output_type -> throttlerdata.ResetConfigurationResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_throttlerservice_throttlerservice_proto_init() }
func file_throttlerservice_throttlerservice_proto_init() {
	if File_throttlerservice_throttlerservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_throttlerservice_throttlerservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_throttlerservice_throttlerservice_proto_goTypes,
		DependencyIndexes: file_throttlerservice_throttlerservice_proto_depIdxs,
	}.Build()
	File_throttlerservice_throttlerservice_proto = out.File
	file_throttlerservice_throttlerservice_proto_rawDesc = nil
	file_throttlerservice_throttlerservice_proto_goTypes = nil
	file_throttlerservice_throttlerservice_proto_depIdxs = nil
}
