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

// This package contains the data structures for the logging service.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: vitess/logutil/v19/logutil.proto

package logutilv19

import (
	v19 "github.com/planetscale/vitess-types/gen/vitess/vttime/v19"
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

// Level is the level of the log messages.
type Level int32

const (
	// The usual logging levels.
	// Should be logged using logging facility.
	Level_INFO    Level = 0
	Level_WARNING Level = 1
	Level_ERROR   Level = 2
	// For messages that may contains non-logging events.
	// Should be logged to console directly.
	Level_CONSOLE Level = 3
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "INFO",
		1: "WARNING",
		2: "ERROR",
		3: "CONSOLE",
	}
	Level_value = map[string]int32{
		"INFO":    0,
		"WARNING": 1,
		"ERROR":   2,
		"CONSOLE": 3,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_vitess_logutil_v19_logutil_proto_enumTypes[0].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_vitess_logutil_v19_logutil_proto_enumTypes[0]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_vitess_logutil_v19_logutil_proto_rawDescGZIP(), []int{0}
}

// Event is a single logging event
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time  *v19.Time `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Level Level     `protobuf:"varint,2,opt,name=level,proto3,enum=vitess.logutil.v19.Level" json:"level,omitempty"`
	File  string    `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Line  int64     `protobuf:"varint,4,opt,name=line,proto3" json:"line,omitempty"`
	Value string    `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vitess_logutil_v19_logutil_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_vitess_logutil_v19_logutil_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_vitess_logutil_v19_logutil_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetTime() *v19.Time {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Event) GetLevel() Level {
	if x != nil {
		return x.Level
	}
	return Level_INFO
}

func (x *Event) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *Event) GetLine() int64 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *Event) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_vitess_logutil_v19_logutil_proto protoreflect.FileDescriptor

var file_vitess_logutil_v19_logutil_proto_rawDesc = []byte{
	0x0a, 0x20, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x75, 0x74, 0x69, 0x6c,
	0x2f, 0x76, 0x31, 0x39, 0x2f, 0x6c, 0x6f, 0x67, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x6c, 0x6f, 0x67, 0x75, 0x74,
	0x69, 0x6c, 0x2e, 0x76, 0x31, 0x39, 0x1a, 0x1e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x74, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x39, 0x2f, 0x76, 0x74, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x2b, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x39, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x2f, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x76,
	0x69, 0x74, 0x65, 0x73, 0x73, 0x2e, 0x6c, 0x6f, 0x67, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x76, 0x31,
	0x39, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x36, 0x0a, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4e, 0x53, 0x4f,
	0x4c, 0x45, 0x10, 0x03, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76,
	0x69, 0x74, 0x65, 0x73, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x76,
	0x31, 0x39, 0x3b, 0x6c, 0x6f, 0x67, 0x75, 0x74, 0x69, 0x6c, 0x76, 0x31, 0x39, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vitess_logutil_v19_logutil_proto_rawDescOnce sync.Once
	file_vitess_logutil_v19_logutil_proto_rawDescData = file_vitess_logutil_v19_logutil_proto_rawDesc
)

func file_vitess_logutil_v19_logutil_proto_rawDescGZIP() []byte {
	file_vitess_logutil_v19_logutil_proto_rawDescOnce.Do(func() {
		file_vitess_logutil_v19_logutil_proto_rawDescData = protoimpl.X.CompressGZIP(file_vitess_logutil_v19_logutil_proto_rawDescData)
	})
	return file_vitess_logutil_v19_logutil_proto_rawDescData
}

var file_vitess_logutil_v19_logutil_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vitess_logutil_v19_logutil_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_vitess_logutil_v19_logutil_proto_goTypes = []any{
	(Level)(0),       // 0: vitess.logutil.v19.Level
	(*Event)(nil),    // 1: vitess.logutil.v19.Event
	(*v19.Time)(nil), // 2: vitess.vttime.v19.Time
}
var file_vitess_logutil_v19_logutil_proto_depIdxs = []int32{
	2, // 0: vitess.logutil.v19.Event.time:type_name -> vitess.vttime.v19.Time
	0, // 1: vitess.logutil.v19.Event.level:type_name -> vitess.logutil.v19.Level
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_vitess_logutil_v19_logutil_proto_init() }
func file_vitess_logutil_v19_logutil_proto_init() {
	if File_vitess_logutil_v19_logutil_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vitess_logutil_v19_logutil_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_vitess_logutil_v19_logutil_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vitess_logutil_v19_logutil_proto_goTypes,
		DependencyIndexes: file_vitess_logutil_v19_logutil_proto_depIdxs,
		EnumInfos:         file_vitess_logutil_v19_logutil_proto_enumTypes,
		MessageInfos:      file_vitess_logutil_v19_logutil_proto_msgTypes,
	}.Build()
	File_vitess_logutil_v19_logutil_proto = out.File
	file_vitess_logutil_v19_logutil_proto_rawDesc = nil
	file_vitess_logutil_v19_logutil_proto_goTypes = nil
	file_vitess_logutil_v19_logutil_proto_depIdxs = nil
}
