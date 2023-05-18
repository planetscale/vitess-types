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

// This file contains the service VtTablet exposes for queries.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: vitess/queryservice/v15/queryservice.proto

package queryservicev15

import (
	v151 "github.com/planetscale/vitess-types/gen/vitess/binlogdata/v15"
	v15 "github.com/planetscale/vitess-types/gen/vitess/query/v15"
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

var File_vitess_queryservice_v15_queryservice_proto protoreflect.FileDescriptor

var file_vitess_queryservice_v15_queryservice_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x76, 0x69,
	0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x35, 0x1a, 0x1c, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x62, 0x69, 0x6e, 0x6c,
	0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x35, 0x2f, 0x62, 0x69, 0x6e, 0x6c, 0x6f,
	0x67, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcb, 0x14, 0x0a, 0x05,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x50, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x12, 0x20, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x35, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x35, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4a, 0x0a,
	0x05, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x1e, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x12, 0x1f, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x21, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73,
	0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a,
	0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x20, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x50, 0x72, 0x65, 0x70,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x69, 0x74,
	0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x50, 0x72,
	0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x65, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65,
	0x64, 0x12, 0x27, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x35, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x76, 0x69, 0x74,
	0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x10, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x12, 0x29, 0x2e, 0x76, 0x69, 0x74,
	0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x6f,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x12, 0x24, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5c, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x12, 0x24, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x35, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x74, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5f, 0x0a, 0x0c, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12,
	0x25, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x35, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x73, 0x0a, 0x12, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x64, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x26, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x59, 0x0a, 0x0a, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x23, 0x2e, 0x76, 0x69, 0x74, 0x65,
	0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x35, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x74, 0x0a,
	0x13, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x42,
	0x65, 0x67, 0x69, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x42, 0x65, 0x67,
	0x69, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x76, 0x69,
	0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x76, 0x69, 0x74,
	0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x88,
	0x01, 0x0a, 0x19, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x76,
	0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x33, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x42, 0x65, 0x67, 0x69, 0x6e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x50, 0x0a, 0x07, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x12, 0x20, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0c, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x25, 0x2e, 0x76, 0x69,
	0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5c,
	0x0a, 0x07, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x25, 0x2e, 0x76, 0x69, 0x74, 0x65,
	0x73, 0x73, 0x2e, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x35, 0x2e, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x68, 0x0a, 0x0b,
	0x56, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x29, 0x2e, 0x76, 0x69,
	0x74, 0x65, 0x73, 0x73, 0x2e, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x35, 0x2e, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e,
	0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x56,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x71, 0x0a, 0x0e, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x2c, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x35,
	0x2e, 0x56, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e,
	0x62, 0x69, 0x6e, 0x6c, 0x6f, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x35, 0x2e, 0x56,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x35, 0x3b, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0x35, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_vitess_queryservice_v15_queryservice_proto_goTypes = []interface{}{
	(*v15.ExecuteRequest)(nil),                    // 0: vitess.query.v15.ExecuteRequest
	(*v15.StreamExecuteRequest)(nil),              // 1: vitess.query.v15.StreamExecuteRequest
	(*v15.BeginRequest)(nil),                      // 2: vitess.query.v15.BeginRequest
	(*v15.CommitRequest)(nil),                     // 3: vitess.query.v15.CommitRequest
	(*v15.RollbackRequest)(nil),                   // 4: vitess.query.v15.RollbackRequest
	(*v15.PrepareRequest)(nil),                    // 5: vitess.query.v15.PrepareRequest
	(*v15.CommitPreparedRequest)(nil),             // 6: vitess.query.v15.CommitPreparedRequest
	(*v15.RollbackPreparedRequest)(nil),           // 7: vitess.query.v15.RollbackPreparedRequest
	(*v15.CreateTransactionRequest)(nil),          // 8: vitess.query.v15.CreateTransactionRequest
	(*v15.StartCommitRequest)(nil),                // 9: vitess.query.v15.StartCommitRequest
	(*v15.SetRollbackRequest)(nil),                // 10: vitess.query.v15.SetRollbackRequest
	(*v15.ConcludeTransactionRequest)(nil),        // 11: vitess.query.v15.ConcludeTransactionRequest
	(*v15.ReadTransactionRequest)(nil),            // 12: vitess.query.v15.ReadTransactionRequest
	(*v15.BeginExecuteRequest)(nil),               // 13: vitess.query.v15.BeginExecuteRequest
	(*v15.BeginStreamExecuteRequest)(nil),         // 14: vitess.query.v15.BeginStreamExecuteRequest
	(*v15.MessageStreamRequest)(nil),              // 15: vitess.query.v15.MessageStreamRequest
	(*v15.MessageAckRequest)(nil),                 // 16: vitess.query.v15.MessageAckRequest
	(*v15.ReserveExecuteRequest)(nil),             // 17: vitess.query.v15.ReserveExecuteRequest
	(*v15.ReserveBeginExecuteRequest)(nil),        // 18: vitess.query.v15.ReserveBeginExecuteRequest
	(*v15.ReserveStreamExecuteRequest)(nil),       // 19: vitess.query.v15.ReserveStreamExecuteRequest
	(*v15.ReserveBeginStreamExecuteRequest)(nil),  // 20: vitess.query.v15.ReserveBeginStreamExecuteRequest
	(*v15.ReleaseRequest)(nil),                    // 21: vitess.query.v15.ReleaseRequest
	(*v15.StreamHealthRequest)(nil),               // 22: vitess.query.v15.StreamHealthRequest
	(*v151.VStreamRequest)(nil),                   // 23: vitess.binlogdata.v15.VStreamRequest
	(*v151.VStreamRowsRequest)(nil),               // 24: vitess.binlogdata.v15.VStreamRowsRequest
	(*v151.VStreamResultsRequest)(nil),            // 25: vitess.binlogdata.v15.VStreamResultsRequest
	(*v15.ExecuteResponse)(nil),                   // 26: vitess.query.v15.ExecuteResponse
	(*v15.StreamExecuteResponse)(nil),             // 27: vitess.query.v15.StreamExecuteResponse
	(*v15.BeginResponse)(nil),                     // 28: vitess.query.v15.BeginResponse
	(*v15.CommitResponse)(nil),                    // 29: vitess.query.v15.CommitResponse
	(*v15.RollbackResponse)(nil),                  // 30: vitess.query.v15.RollbackResponse
	(*v15.PrepareResponse)(nil),                   // 31: vitess.query.v15.PrepareResponse
	(*v15.CommitPreparedResponse)(nil),            // 32: vitess.query.v15.CommitPreparedResponse
	(*v15.RollbackPreparedResponse)(nil),          // 33: vitess.query.v15.RollbackPreparedResponse
	(*v15.CreateTransactionResponse)(nil),         // 34: vitess.query.v15.CreateTransactionResponse
	(*v15.StartCommitResponse)(nil),               // 35: vitess.query.v15.StartCommitResponse
	(*v15.SetRollbackResponse)(nil),               // 36: vitess.query.v15.SetRollbackResponse
	(*v15.ConcludeTransactionResponse)(nil),       // 37: vitess.query.v15.ConcludeTransactionResponse
	(*v15.ReadTransactionResponse)(nil),           // 38: vitess.query.v15.ReadTransactionResponse
	(*v15.BeginExecuteResponse)(nil),              // 39: vitess.query.v15.BeginExecuteResponse
	(*v15.BeginStreamExecuteResponse)(nil),        // 40: vitess.query.v15.BeginStreamExecuteResponse
	(*v15.MessageStreamResponse)(nil),             // 41: vitess.query.v15.MessageStreamResponse
	(*v15.MessageAckResponse)(nil),                // 42: vitess.query.v15.MessageAckResponse
	(*v15.ReserveExecuteResponse)(nil),            // 43: vitess.query.v15.ReserveExecuteResponse
	(*v15.ReserveBeginExecuteResponse)(nil),       // 44: vitess.query.v15.ReserveBeginExecuteResponse
	(*v15.ReserveStreamExecuteResponse)(nil),      // 45: vitess.query.v15.ReserveStreamExecuteResponse
	(*v15.ReserveBeginStreamExecuteResponse)(nil), // 46: vitess.query.v15.ReserveBeginStreamExecuteResponse
	(*v15.ReleaseResponse)(nil),                   // 47: vitess.query.v15.ReleaseResponse
	(*v15.StreamHealthResponse)(nil),              // 48: vitess.query.v15.StreamHealthResponse
	(*v151.VStreamResponse)(nil),                  // 49: vitess.binlogdata.v15.VStreamResponse
	(*v151.VStreamRowsResponse)(nil),              // 50: vitess.binlogdata.v15.VStreamRowsResponse
	(*v151.VStreamResultsResponse)(nil),           // 51: vitess.binlogdata.v15.VStreamResultsResponse
}
var file_vitess_queryservice_v15_queryservice_proto_depIdxs = []int32{
	0,  // 0: vitess.queryservice.v15.Query.Execute:input_type -> vitess.query.v15.ExecuteRequest
	1,  // 1: vitess.queryservice.v15.Query.StreamExecute:input_type -> vitess.query.v15.StreamExecuteRequest
	2,  // 2: vitess.queryservice.v15.Query.Begin:input_type -> vitess.query.v15.BeginRequest
	3,  // 3: vitess.queryservice.v15.Query.Commit:input_type -> vitess.query.v15.CommitRequest
	4,  // 4: vitess.queryservice.v15.Query.Rollback:input_type -> vitess.query.v15.RollbackRequest
	5,  // 5: vitess.queryservice.v15.Query.Prepare:input_type -> vitess.query.v15.PrepareRequest
	6,  // 6: vitess.queryservice.v15.Query.CommitPrepared:input_type -> vitess.query.v15.CommitPreparedRequest
	7,  // 7: vitess.queryservice.v15.Query.RollbackPrepared:input_type -> vitess.query.v15.RollbackPreparedRequest
	8,  // 8: vitess.queryservice.v15.Query.CreateTransaction:input_type -> vitess.query.v15.CreateTransactionRequest
	9,  // 9: vitess.queryservice.v15.Query.StartCommit:input_type -> vitess.query.v15.StartCommitRequest
	10, // 10: vitess.queryservice.v15.Query.SetRollback:input_type -> vitess.query.v15.SetRollbackRequest
	11, // 11: vitess.queryservice.v15.Query.ConcludeTransaction:input_type -> vitess.query.v15.ConcludeTransactionRequest
	12, // 12: vitess.queryservice.v15.Query.ReadTransaction:input_type -> vitess.query.v15.ReadTransactionRequest
	13, // 13: vitess.queryservice.v15.Query.BeginExecute:input_type -> vitess.query.v15.BeginExecuteRequest
	14, // 14: vitess.queryservice.v15.Query.BeginStreamExecute:input_type -> vitess.query.v15.BeginStreamExecuteRequest
	15, // 15: vitess.queryservice.v15.Query.MessageStream:input_type -> vitess.query.v15.MessageStreamRequest
	16, // 16: vitess.queryservice.v15.Query.MessageAck:input_type -> vitess.query.v15.MessageAckRequest
	17, // 17: vitess.queryservice.v15.Query.ReserveExecute:input_type -> vitess.query.v15.ReserveExecuteRequest
	18, // 18: vitess.queryservice.v15.Query.ReserveBeginExecute:input_type -> vitess.query.v15.ReserveBeginExecuteRequest
	19, // 19: vitess.queryservice.v15.Query.ReserveStreamExecute:input_type -> vitess.query.v15.ReserveStreamExecuteRequest
	20, // 20: vitess.queryservice.v15.Query.ReserveBeginStreamExecute:input_type -> vitess.query.v15.ReserveBeginStreamExecuteRequest
	21, // 21: vitess.queryservice.v15.Query.Release:input_type -> vitess.query.v15.ReleaseRequest
	22, // 22: vitess.queryservice.v15.Query.StreamHealth:input_type -> vitess.query.v15.StreamHealthRequest
	23, // 23: vitess.queryservice.v15.Query.VStream:input_type -> vitess.binlogdata.v15.VStreamRequest
	24, // 24: vitess.queryservice.v15.Query.VStreamRows:input_type -> vitess.binlogdata.v15.VStreamRowsRequest
	25, // 25: vitess.queryservice.v15.Query.VStreamResults:input_type -> vitess.binlogdata.v15.VStreamResultsRequest
	26, // 26: vitess.queryservice.v15.Query.Execute:output_type -> vitess.query.v15.ExecuteResponse
	27, // 27: vitess.queryservice.v15.Query.StreamExecute:output_type -> vitess.query.v15.StreamExecuteResponse
	28, // 28: vitess.queryservice.v15.Query.Begin:output_type -> vitess.query.v15.BeginResponse
	29, // 29: vitess.queryservice.v15.Query.Commit:output_type -> vitess.query.v15.CommitResponse
	30, // 30: vitess.queryservice.v15.Query.Rollback:output_type -> vitess.query.v15.RollbackResponse
	31, // 31: vitess.queryservice.v15.Query.Prepare:output_type -> vitess.query.v15.PrepareResponse
	32, // 32: vitess.queryservice.v15.Query.CommitPrepared:output_type -> vitess.query.v15.CommitPreparedResponse
	33, // 33: vitess.queryservice.v15.Query.RollbackPrepared:output_type -> vitess.query.v15.RollbackPreparedResponse
	34, // 34: vitess.queryservice.v15.Query.CreateTransaction:output_type -> vitess.query.v15.CreateTransactionResponse
	35, // 35: vitess.queryservice.v15.Query.StartCommit:output_type -> vitess.query.v15.StartCommitResponse
	36, // 36: vitess.queryservice.v15.Query.SetRollback:output_type -> vitess.query.v15.SetRollbackResponse
	37, // 37: vitess.queryservice.v15.Query.ConcludeTransaction:output_type -> vitess.query.v15.ConcludeTransactionResponse
	38, // 38: vitess.queryservice.v15.Query.ReadTransaction:output_type -> vitess.query.v15.ReadTransactionResponse
	39, // 39: vitess.queryservice.v15.Query.BeginExecute:output_type -> vitess.query.v15.BeginExecuteResponse
	40, // 40: vitess.queryservice.v15.Query.BeginStreamExecute:output_type -> vitess.query.v15.BeginStreamExecuteResponse
	41, // 41: vitess.queryservice.v15.Query.MessageStream:output_type -> vitess.query.v15.MessageStreamResponse
	42, // 42: vitess.queryservice.v15.Query.MessageAck:output_type -> vitess.query.v15.MessageAckResponse
	43, // 43: vitess.queryservice.v15.Query.ReserveExecute:output_type -> vitess.query.v15.ReserveExecuteResponse
	44, // 44: vitess.queryservice.v15.Query.ReserveBeginExecute:output_type -> vitess.query.v15.ReserveBeginExecuteResponse
	45, // 45: vitess.queryservice.v15.Query.ReserveStreamExecute:output_type -> vitess.query.v15.ReserveStreamExecuteResponse
	46, // 46: vitess.queryservice.v15.Query.ReserveBeginStreamExecute:output_type -> vitess.query.v15.ReserveBeginStreamExecuteResponse
	47, // 47: vitess.queryservice.v15.Query.Release:output_type -> vitess.query.v15.ReleaseResponse
	48, // 48: vitess.queryservice.v15.Query.StreamHealth:output_type -> vitess.query.v15.StreamHealthResponse
	49, // 49: vitess.queryservice.v15.Query.VStream:output_type -> vitess.binlogdata.v15.VStreamResponse
	50, // 50: vitess.queryservice.v15.Query.VStreamRows:output_type -> vitess.binlogdata.v15.VStreamRowsResponse
	51, // 51: vitess.queryservice.v15.Query.VStreamResults:output_type -> vitess.binlogdata.v15.VStreamResultsResponse
	26, // [26:52] is the sub-list for method output_type
	0,  // [0:26] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_vitess_queryservice_v15_queryservice_proto_init() }
func file_vitess_queryservice_v15_queryservice_proto_init() {
	if File_vitess_queryservice_v15_queryservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vitess_queryservice_v15_queryservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vitess_queryservice_v15_queryservice_proto_goTypes,
		DependencyIndexes: file_vitess_queryservice_v15_queryservice_proto_depIdxs,
	}.Build()
	File_vitess_queryservice_v15_queryservice_proto = out.File
	file_vitess_queryservice_v15_queryservice_proto_rawDesc = nil
	file_vitess_queryservice_v15_queryservice_proto_goTypes = nil
	file_vitess_queryservice_v15_queryservice_proto_depIdxs = nil
}