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

// This file contains the Vitess workflow management related data
// structures.  They are used to store / retrieve state from topology
// server.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: workflow/workflow.proto

package workflow

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

// WorkflowState describes the state of a workflow.
// This constant should match the Node object described in
// web/vtctld2/src/app/workflows/node.ts as it is exposed as JSON to
// the Angular 2 web app.
type WorkflowState int32

const (
	WorkflowState_NotStarted WorkflowState = 0
	WorkflowState_Running    WorkflowState = 1
	WorkflowState_Done       WorkflowState = 2
)

// Enum value maps for WorkflowState.
var (
	WorkflowState_name = map[int32]string{
		0: "NotStarted",
		1: "Running",
		2: "Done",
	}
	WorkflowState_value = map[string]int32{
		"NotStarted": 0,
		"Running":    1,
		"Done":       2,
	}
)

func (x WorkflowState) Enum() *WorkflowState {
	p := new(WorkflowState)
	*p = x
	return p
}

func (x WorkflowState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkflowState) Descriptor() protoreflect.EnumDescriptor {
	return file_workflow_workflow_proto_enumTypes[0].Descriptor()
}

func (WorkflowState) Type() protoreflect.EnumType {
	return &file_workflow_workflow_proto_enumTypes[0]
}

func (x WorkflowState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkflowState.Descriptor instead.
func (WorkflowState) EnumDescriptor() ([]byte, []int) {
	return file_workflow_workflow_proto_rawDescGZIP(), []int{0}
}

type TaskState int32

const (
	TaskState_TaskNotStarted TaskState = 0
	TaskState_TaskRunning    TaskState = 1
	TaskState_TaskDone       TaskState = 2
)

// Enum value maps for TaskState.
var (
	TaskState_name = map[int32]string{
		0: "TaskNotStarted",
		1: "TaskRunning",
		2: "TaskDone",
	}
	TaskState_value = map[string]int32{
		"TaskNotStarted": 0,
		"TaskRunning":    1,
		"TaskDone":       2,
	}
)

func (x TaskState) Enum() *TaskState {
	p := new(TaskState)
	*p = x
	return p
}

func (x TaskState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskState) Descriptor() protoreflect.EnumDescriptor {
	return file_workflow_workflow_proto_enumTypes[1].Descriptor()
}

func (TaskState) Type() protoreflect.EnumType {
	return &file_workflow_workflow_proto_enumTypes[1]
}

func (x TaskState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskState.Descriptor instead.
func (TaskState) EnumDescriptor() ([]byte, []int) {
	return file_workflow_workflow_proto_rawDescGZIP(), []int{1}
}

// Workflow is the persisted state of a long-running workflow.
type Workflow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// uuid is set when the workflow is created, and immutable after
	// that.
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// factory_name is set with the name of the factory that created the
	// job (and can also restart it). It is set at creation time, and
	// immutable after that.
	FactoryName string `protobuf:"bytes,2,opt,name=factory_name,json=factoryName,proto3" json:"factory_name,omitempty"`
	// name is the display name of the workflow.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// state describes the state of the job. A job is created as
	// NotStarted, then the Workflow Manager picks it up and starts it,
	// switching it to Running (and populating 'start_time').  The
	// workflow can then fail over to a new Workflow Manager is
	// necessary, and still be in Running state.  When done, it goes to
	// Done, 'end_time' is populated, and 'error' is set if there was an
	// error.
	State WorkflowState `protobuf:"varint,4,opt,name=state,proto3,enum=workflow.WorkflowState" json:"state,omitempty"`
	// data is workflow-specific stored data. It is usually a binary
	// proto-encoded data structure. It can vary throughout the
	// execution of the workflow.  It will not change after the workflow
	// is Done.
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	// error is set if the job finished with an error. This field only
	// makes sense if 'state' is Done.
	Error string `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	// start_time is set when the workflow manager starts a workflow for
	// the first time. This field only makes sense if 'state' is Running
	// or Done.
	StartTime int64 `protobuf:"varint,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// end_time is set when the workflow is finished.
	// This field only makes sense if 'state' is Done.
	EndTime int64 `protobuf:"varint,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// create_time is set when the workflow is created.
	CreateTime int64 `protobuf:"varint,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Workflow) Reset() {
	*x = Workflow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_workflow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workflow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workflow) ProtoMessage() {}

func (x *Workflow) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_workflow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workflow.ProtoReflect.Descriptor instead.
func (*Workflow) Descriptor() ([]byte, []int) {
	return file_workflow_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *Workflow) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Workflow) GetFactoryName() string {
	if x != nil {
		return x.FactoryName
	}
	return ""
}

func (x *Workflow) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Workflow) GetState() WorkflowState {
	if x != nil {
		return x.State
	}
	return WorkflowState_NotStarted
}

func (x *Workflow) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Workflow) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Workflow) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Workflow) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *Workflow) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type WorkflowCheckpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// code_version is used to detect incompabilities between the version of the
	// running workflow and the one which wrote the checkpoint. If they don't
	// match, the workflow must not continue. The author of workflow must update
	// this variable in their implementation when incompabilities are introduced.
	CodeVersion int32 `protobuf:"varint,1,opt,name=code_version,json=codeVersion,proto3" json:"code_version,omitempty"`
	// Task is the data structure that stores the execution status and the
	// attributes of a task.
	Tasks map[string]*Task `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// settings includes workflow specific data, e.g. the resharding workflow
	// would store the source shards and destination shards.
	Settings map[string]string `protobuf:"bytes,3,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WorkflowCheckpoint) Reset() {
	*x = WorkflowCheckpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_workflow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowCheckpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowCheckpoint) ProtoMessage() {}

func (x *WorkflowCheckpoint) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_workflow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowCheckpoint.ProtoReflect.Descriptor instead.
func (*WorkflowCheckpoint) Descriptor() ([]byte, []int) {
	return file_workflow_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowCheckpoint) GetCodeVersion() int32 {
	if x != nil {
		return x.CodeVersion
	}
	return 0
}

func (x *WorkflowCheckpoint) GetTasks() map[string]*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *WorkflowCheckpoint) GetSettings() map[string]string {
	if x != nil {
		return x.Settings
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State TaskState `protobuf:"varint,2,opt,name=state,proto3,enum=workflow.TaskState" json:"state,omitempty"`
	// attributes includes the parameters the task needs.
	Attributes map[string]string `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Error      string            `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workflow_workflow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_workflow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_workflow_workflow_proto_rawDescGZIP(), []int{2}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetState() TaskState {
	if x != nil {
		return x.State
	}
	return TaskState_TaskNotStarted
}

func (x *Task) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Task) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_workflow_workflow_proto protoreflect.FileDescriptor

var file_workflow_workflow_proto_rawDesc = []byte{
	0x0a, 0x17, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x22, 0x89, 0x02, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0xc5, 0x02, 0x0a, 0x12, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f,
	0x64, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x46, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x1a, 0x48, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd6, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x29, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x2a, 0x36, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x02, 0x2a, 0x3e, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x4e, 0x6f, 0x74,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x61,
	0x73, 0x6b, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x02, 0x42, 0x9a, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x34, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0xa2, 0x02, 0x03, 0x57, 0x58, 0x58, 0xaa, 0x02, 0x08,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0xca, 0x02, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0xe2, 0x02, 0x14, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workflow_workflow_proto_rawDescOnce sync.Once
	file_workflow_workflow_proto_rawDescData = file_workflow_workflow_proto_rawDesc
)

func file_workflow_workflow_proto_rawDescGZIP() []byte {
	file_workflow_workflow_proto_rawDescOnce.Do(func() {
		file_workflow_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_workflow_workflow_proto_rawDescData)
	})
	return file_workflow_workflow_proto_rawDescData
}

var file_workflow_workflow_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_workflow_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_workflow_workflow_proto_goTypes = []interface{}{
	(WorkflowState)(0),         // 0: workflow.WorkflowState
	(TaskState)(0),             // 1: workflow.TaskState
	(*Workflow)(nil),           // 2: workflow.Workflow
	(*WorkflowCheckpoint)(nil), // 3: workflow.WorkflowCheckpoint
	(*Task)(nil),               // 4: workflow.Task
	nil,                        // 5: workflow.WorkflowCheckpoint.TasksEntry
	nil,                        // 6: workflow.WorkflowCheckpoint.SettingsEntry
	nil,                        // 7: workflow.Task.AttributesEntry
}
var file_workflow_workflow_proto_depIdxs = []int32{
	0, // 0: workflow.Workflow.state:type_name -> workflow.WorkflowState
	5, // 1: workflow.WorkflowCheckpoint.tasks:type_name -> workflow.WorkflowCheckpoint.TasksEntry
	6, // 2: workflow.WorkflowCheckpoint.settings:type_name -> workflow.WorkflowCheckpoint.SettingsEntry
	1, // 3: workflow.Task.state:type_name -> workflow.TaskState
	7, // 4: workflow.Task.attributes:type_name -> workflow.Task.AttributesEntry
	4, // 5: workflow.WorkflowCheckpoint.TasksEntry.value:type_name -> workflow.Task
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_workflow_workflow_proto_init() }
func file_workflow_workflow_proto_init() {
	if File_workflow_workflow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workflow_workflow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workflow); i {
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
		file_workflow_workflow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowCheckpoint); i {
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
		file_workflow_workflow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
			RawDescriptor: file_workflow_workflow_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_workflow_workflow_proto_goTypes,
		DependencyIndexes: file_workflow_workflow_proto_depIdxs,
		EnumInfos:         file_workflow_workflow_proto_enumTypes,
		MessageInfos:      file_workflow_workflow_proto_msgTypes,
	}.Build()
	File_workflow_workflow_proto = out.File
	file_workflow_workflow_proto_rawDesc = nil
	file_workflow_workflow_proto_goTypes = nil
	file_workflow_workflow_proto_depIdxs = nil
}
