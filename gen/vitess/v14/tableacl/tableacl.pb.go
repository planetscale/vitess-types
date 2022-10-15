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

// Table ACL proto definitions.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tableacl/tableacl.proto

package tableacl

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

// TableGroupSpec defines ACLs for a group of tables.
type TableGroupSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// either tables or a table name prefixes (if it ends in a %)
	TableNamesOrPrefixes []string `protobuf:"bytes,2,rep,name=table_names_or_prefixes,json=tableNamesOrPrefixes,proto3" json:"table_names_or_prefixes,omitempty"`
	Readers              []string `protobuf:"bytes,3,rep,name=readers,proto3" json:"readers,omitempty"`
	Writers              []string `protobuf:"bytes,4,rep,name=writers,proto3" json:"writers,omitempty"`
	Admins               []string `protobuf:"bytes,5,rep,name=admins,proto3" json:"admins,omitempty"`
}

func (x *TableGroupSpec) Reset() {
	*x = TableGroupSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableacl_tableacl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableGroupSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableGroupSpec) ProtoMessage() {}

func (x *TableGroupSpec) ProtoReflect() protoreflect.Message {
	mi := &file_tableacl_tableacl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableGroupSpec.ProtoReflect.Descriptor instead.
func (*TableGroupSpec) Descriptor() ([]byte, []int) {
	return file_tableacl_tableacl_proto_rawDescGZIP(), []int{0}
}

func (x *TableGroupSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TableGroupSpec) GetTableNamesOrPrefixes() []string {
	if x != nil {
		return x.TableNamesOrPrefixes
	}
	return nil
}

func (x *TableGroupSpec) GetReaders() []string {
	if x != nil {
		return x.Readers
	}
	return nil
}

func (x *TableGroupSpec) GetWriters() []string {
	if x != nil {
		return x.Writers
	}
	return nil
}

func (x *TableGroupSpec) GetAdmins() []string {
	if x != nil {
		return x.Admins
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableGroups []*TableGroupSpec `protobuf:"bytes,1,rep,name=table_groups,json=tableGroups,proto3" json:"table_groups,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tableacl_tableacl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_tableacl_tableacl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_tableacl_tableacl_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetTableGroups() []*TableGroupSpec {
	if x != nil {
		return x.TableGroups
	}
	return nil
}

var File_tableacl_tableacl_proto protoreflect.FileDescriptor

var file_tableacl_tableacl_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x63, 0x6c, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x61, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x61, 0x63, 0x6c, 0x22, 0xa7, 0x01, 0x0a, 0x0e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x4f, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x22, 0x45, 0x0a,
	0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x63, 0x6c, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x42, 0x9a, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x61, 0x63, 0x6c, 0x42, 0x0d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x63, 0x6c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76,
	0x69, 0x74, 0x65, 0x73, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x61, 0x63, 0x6c, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x61, 0x63, 0x6c, 0xca, 0x02, 0x08, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x63, 0x6c, 0xe2,
	0x02, 0x14, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x63, 0x6c, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x63,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tableacl_tableacl_proto_rawDescOnce sync.Once
	file_tableacl_tableacl_proto_rawDescData = file_tableacl_tableacl_proto_rawDesc
)

func file_tableacl_tableacl_proto_rawDescGZIP() []byte {
	file_tableacl_tableacl_proto_rawDescOnce.Do(func() {
		file_tableacl_tableacl_proto_rawDescData = protoimpl.X.CompressGZIP(file_tableacl_tableacl_proto_rawDescData)
	})
	return file_tableacl_tableacl_proto_rawDescData
}

var file_tableacl_tableacl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tableacl_tableacl_proto_goTypes = []interface{}{
	(*TableGroupSpec)(nil), // 0: tableacl.TableGroupSpec
	(*Config)(nil),         // 1: tableacl.Config
}
var file_tableacl_tableacl_proto_depIdxs = []int32{
	0, // 0: tableacl.Config.table_groups:type_name -> tableacl.TableGroupSpec
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tableacl_tableacl_proto_init() }
func file_tableacl_tableacl_proto_init() {
	if File_tableacl_tableacl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tableacl_tableacl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableGroupSpec); i {
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
		file_tableacl_tableacl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_tableacl_tableacl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tableacl_tableacl_proto_goTypes,
		DependencyIndexes: file_tableacl_tableacl_proto_depIdxs,
		MessageInfos:      file_tableacl_tableacl_proto_msgTypes,
	}.Build()
	File_tableacl_tableacl_proto = out.File
	file_tableacl_tableacl_proto_rawDesc = nil
	file_tableacl_tableacl_proto_goTypes = nil
	file_tableacl_tableacl_proto_depIdxs = nil
}
