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
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: vitess/queryservice/dev/queryservice.proto

package queryservicedev

import (
	dev1 "github.com/planetscale/vitess-types/gen/vitess/binlogdata/dev"
	dev "github.com/planetscale/vitess-types/gen/vitess/query/dev"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_vitess_queryservice_dev_queryservice_proto protoreflect.FileDescriptor

const file_vitess_queryservice_dev_queryservice_proto_rawDesc = "" +
	"\n" +
	"*vitess/queryservice/dev/queryservice.proto\x12\x17vitess.queryservice.dev\x1a\x1cvitess/query/dev/query.proto\x1a&vitess/binlogdata/dev/binlogdata.proto2\x94\x17\n" +
	"\x05Query\x12P\n" +
	"\aExecute\x12 .vitess.query.dev.ExecuteRequest\x1a!.vitess.query.dev.ExecuteResponse\"\x00\x12d\n" +
	"\rStreamExecute\x12&.vitess.query.dev.StreamExecuteRequest\x1a'.vitess.query.dev.StreamExecuteResponse\"\x000\x01\x12J\n" +
	"\x05Begin\x12\x1e.vitess.query.dev.BeginRequest\x1a\x1f.vitess.query.dev.BeginResponse\"\x00\x12M\n" +
	"\x06Commit\x12\x1f.vitess.query.dev.CommitRequest\x1a .vitess.query.dev.CommitResponse\"\x00\x12S\n" +
	"\bRollback\x12!.vitess.query.dev.RollbackRequest\x1a\".vitess.query.dev.RollbackResponse\"\x00\x12P\n" +
	"\aPrepare\x12 .vitess.query.dev.PrepareRequest\x1a!.vitess.query.dev.PrepareResponse\"\x00\x12e\n" +
	"\x0eCommitPrepared\x12'.vitess.query.dev.CommitPreparedRequest\x1a(.vitess.query.dev.CommitPreparedResponse\"\x00\x12k\n" +
	"\x10RollbackPrepared\x12).vitess.query.dev.RollbackPreparedRequest\x1a*.vitess.query.dev.RollbackPreparedResponse\"\x00\x12n\n" +
	"\x11CreateTransaction\x12*.vitess.query.dev.CreateTransactionRequest\x1a+.vitess.query.dev.CreateTransactionResponse\"\x00\x12\\\n" +
	"\vStartCommit\x12$.vitess.query.dev.StartCommitRequest\x1a%.vitess.query.dev.StartCommitResponse\"\x00\x12\\\n" +
	"\vSetRollback\x12$.vitess.query.dev.SetRollbackRequest\x1a%.vitess.query.dev.SetRollbackResponse\"\x00\x12t\n" +
	"\x13ConcludeTransaction\x12,.vitess.query.dev.ConcludeTransactionRequest\x1a-.vitess.query.dev.ConcludeTransactionResponse\"\x00\x12h\n" +
	"\x0fReadTransaction\x12(.vitess.query.dev.ReadTransactionRequest\x1a).vitess.query.dev.ReadTransactionResponse\"\x00\x12}\n" +
	"\x16UnresolvedTransactions\x12/.vitess.query.dev.UnresolvedTransactionsRequest\x1a0.vitess.query.dev.UnresolvedTransactionsResponse\"\x00\x12_\n" +
	"\fBeginExecute\x12%.vitess.query.dev.BeginExecuteRequest\x1a&.vitess.query.dev.BeginExecuteResponse\"\x00\x12s\n" +
	"\x12BeginStreamExecute\x12+.vitess.query.dev.BeginStreamExecuteRequest\x1a,.vitess.query.dev.BeginStreamExecuteResponse\"\x000\x01\x12d\n" +
	"\rMessageStream\x12&.vitess.query.dev.MessageStreamRequest\x1a'.vitess.query.dev.MessageStreamResponse\"\x000\x01\x12Y\n" +
	"\n" +
	"MessageAck\x12#.vitess.query.dev.MessageAckRequest\x1a$.vitess.query.dev.MessageAckResponse\"\x00\x12e\n" +
	"\x0eReserveExecute\x12'.vitess.query.dev.ReserveExecuteRequest\x1a(.vitess.query.dev.ReserveExecuteResponse\"\x00\x12t\n" +
	"\x13ReserveBeginExecute\x12,.vitess.query.dev.ReserveBeginExecuteRequest\x1a-.vitess.query.dev.ReserveBeginExecuteResponse\"\x00\x12y\n" +
	"\x14ReserveStreamExecute\x12-.vitess.query.dev.ReserveStreamExecuteRequest\x1a..vitess.query.dev.ReserveStreamExecuteResponse\"\x000\x01\x12\x88\x01\n" +
	"\x19ReserveBeginStreamExecute\x122.vitess.query.dev.ReserveBeginStreamExecuteRequest\x1a3.vitess.query.dev.ReserveBeginStreamExecuteResponse\"\x000\x01\x12P\n" +
	"\aRelease\x12 .vitess.query.dev.ReleaseRequest\x1a!.vitess.query.dev.ReleaseResponse\"\x00\x12a\n" +
	"\fStreamHealth\x12%.vitess.query.dev.StreamHealthRequest\x1a&.vitess.query.dev.StreamHealthResponse\"\x000\x01\x12\\\n" +
	"\aVStream\x12%.vitess.binlogdata.dev.VStreamRequest\x1a&.vitess.binlogdata.dev.VStreamResponse\"\x000\x01\x12h\n" +
	"\vVStreamRows\x12).vitess.binlogdata.dev.VStreamRowsRequest\x1a*.vitess.binlogdata.dev.VStreamRowsResponse\"\x000\x01\x12n\n" +
	"\rVStreamTables\x12+.vitess.binlogdata.dev.VStreamTablesRequest\x1a,.vitess.binlogdata.dev.VStreamTablesResponse\"\x000\x01\x12q\n" +
	"\x0eVStreamResults\x12,.vitess.binlogdata.dev.VStreamResultsRequest\x1a-.vitess.binlogdata.dev.VStreamResultsResponse\"\x000\x01\x12X\n" +
	"\tGetSchema\x12\".vitess.query.dev.GetSchemaRequest\x1a#.vitess.query.dev.GetSchemaResponse\"\x000\x01BQZOgithub.com/planetscale/vitess-types/gen/vitess/queryservice/dev;queryservicedevb\x06proto3"

var file_vitess_queryservice_dev_queryservice_proto_goTypes = []any{
	(*dev.ExecuteRequest)(nil),                    // 0: vitess.query.dev.ExecuteRequest
	(*dev.StreamExecuteRequest)(nil),              // 1: vitess.query.dev.StreamExecuteRequest
	(*dev.BeginRequest)(nil),                      // 2: vitess.query.dev.BeginRequest
	(*dev.CommitRequest)(nil),                     // 3: vitess.query.dev.CommitRequest
	(*dev.RollbackRequest)(nil),                   // 4: vitess.query.dev.RollbackRequest
	(*dev.PrepareRequest)(nil),                    // 5: vitess.query.dev.PrepareRequest
	(*dev.CommitPreparedRequest)(nil),             // 6: vitess.query.dev.CommitPreparedRequest
	(*dev.RollbackPreparedRequest)(nil),           // 7: vitess.query.dev.RollbackPreparedRequest
	(*dev.CreateTransactionRequest)(nil),          // 8: vitess.query.dev.CreateTransactionRequest
	(*dev.StartCommitRequest)(nil),                // 9: vitess.query.dev.StartCommitRequest
	(*dev.SetRollbackRequest)(nil),                // 10: vitess.query.dev.SetRollbackRequest
	(*dev.ConcludeTransactionRequest)(nil),        // 11: vitess.query.dev.ConcludeTransactionRequest
	(*dev.ReadTransactionRequest)(nil),            // 12: vitess.query.dev.ReadTransactionRequest
	(*dev.UnresolvedTransactionsRequest)(nil),     // 13: vitess.query.dev.UnresolvedTransactionsRequest
	(*dev.BeginExecuteRequest)(nil),               // 14: vitess.query.dev.BeginExecuteRequest
	(*dev.BeginStreamExecuteRequest)(nil),         // 15: vitess.query.dev.BeginStreamExecuteRequest
	(*dev.MessageStreamRequest)(nil),              // 16: vitess.query.dev.MessageStreamRequest
	(*dev.MessageAckRequest)(nil),                 // 17: vitess.query.dev.MessageAckRequest
	(*dev.ReserveExecuteRequest)(nil),             // 18: vitess.query.dev.ReserveExecuteRequest
	(*dev.ReserveBeginExecuteRequest)(nil),        // 19: vitess.query.dev.ReserveBeginExecuteRequest
	(*dev.ReserveStreamExecuteRequest)(nil),       // 20: vitess.query.dev.ReserveStreamExecuteRequest
	(*dev.ReserveBeginStreamExecuteRequest)(nil),  // 21: vitess.query.dev.ReserveBeginStreamExecuteRequest
	(*dev.ReleaseRequest)(nil),                    // 22: vitess.query.dev.ReleaseRequest
	(*dev.StreamHealthRequest)(nil),               // 23: vitess.query.dev.StreamHealthRequest
	(*dev1.VStreamRequest)(nil),                   // 24: vitess.binlogdata.dev.VStreamRequest
	(*dev1.VStreamRowsRequest)(nil),               // 25: vitess.binlogdata.dev.VStreamRowsRequest
	(*dev1.VStreamTablesRequest)(nil),             // 26: vitess.binlogdata.dev.VStreamTablesRequest
	(*dev1.VStreamResultsRequest)(nil),            // 27: vitess.binlogdata.dev.VStreamResultsRequest
	(*dev.GetSchemaRequest)(nil),                  // 28: vitess.query.dev.GetSchemaRequest
	(*dev.ExecuteResponse)(nil),                   // 29: vitess.query.dev.ExecuteResponse
	(*dev.StreamExecuteResponse)(nil),             // 30: vitess.query.dev.StreamExecuteResponse
	(*dev.BeginResponse)(nil),                     // 31: vitess.query.dev.BeginResponse
	(*dev.CommitResponse)(nil),                    // 32: vitess.query.dev.CommitResponse
	(*dev.RollbackResponse)(nil),                  // 33: vitess.query.dev.RollbackResponse
	(*dev.PrepareResponse)(nil),                   // 34: vitess.query.dev.PrepareResponse
	(*dev.CommitPreparedResponse)(nil),            // 35: vitess.query.dev.CommitPreparedResponse
	(*dev.RollbackPreparedResponse)(nil),          // 36: vitess.query.dev.RollbackPreparedResponse
	(*dev.CreateTransactionResponse)(nil),         // 37: vitess.query.dev.CreateTransactionResponse
	(*dev.StartCommitResponse)(nil),               // 38: vitess.query.dev.StartCommitResponse
	(*dev.SetRollbackResponse)(nil),               // 39: vitess.query.dev.SetRollbackResponse
	(*dev.ConcludeTransactionResponse)(nil),       // 40: vitess.query.dev.ConcludeTransactionResponse
	(*dev.ReadTransactionResponse)(nil),           // 41: vitess.query.dev.ReadTransactionResponse
	(*dev.UnresolvedTransactionsResponse)(nil),    // 42: vitess.query.dev.UnresolvedTransactionsResponse
	(*dev.BeginExecuteResponse)(nil),              // 43: vitess.query.dev.BeginExecuteResponse
	(*dev.BeginStreamExecuteResponse)(nil),        // 44: vitess.query.dev.BeginStreamExecuteResponse
	(*dev.MessageStreamResponse)(nil),             // 45: vitess.query.dev.MessageStreamResponse
	(*dev.MessageAckResponse)(nil),                // 46: vitess.query.dev.MessageAckResponse
	(*dev.ReserveExecuteResponse)(nil),            // 47: vitess.query.dev.ReserveExecuteResponse
	(*dev.ReserveBeginExecuteResponse)(nil),       // 48: vitess.query.dev.ReserveBeginExecuteResponse
	(*dev.ReserveStreamExecuteResponse)(nil),      // 49: vitess.query.dev.ReserveStreamExecuteResponse
	(*dev.ReserveBeginStreamExecuteResponse)(nil), // 50: vitess.query.dev.ReserveBeginStreamExecuteResponse
	(*dev.ReleaseResponse)(nil),                   // 51: vitess.query.dev.ReleaseResponse
	(*dev.StreamHealthResponse)(nil),              // 52: vitess.query.dev.StreamHealthResponse
	(*dev1.VStreamResponse)(nil),                  // 53: vitess.binlogdata.dev.VStreamResponse
	(*dev1.VStreamRowsResponse)(nil),              // 54: vitess.binlogdata.dev.VStreamRowsResponse
	(*dev1.VStreamTablesResponse)(nil),            // 55: vitess.binlogdata.dev.VStreamTablesResponse
	(*dev1.VStreamResultsResponse)(nil),           // 56: vitess.binlogdata.dev.VStreamResultsResponse
	(*dev.GetSchemaResponse)(nil),                 // 57: vitess.query.dev.GetSchemaResponse
}
var file_vitess_queryservice_dev_queryservice_proto_depIdxs = []int32{
	0,  // 0: vitess.queryservice.dev.Query.Execute:input_type -> vitess.query.dev.ExecuteRequest
	1,  // 1: vitess.queryservice.dev.Query.StreamExecute:input_type -> vitess.query.dev.StreamExecuteRequest
	2,  // 2: vitess.queryservice.dev.Query.Begin:input_type -> vitess.query.dev.BeginRequest
	3,  // 3: vitess.queryservice.dev.Query.Commit:input_type -> vitess.query.dev.CommitRequest
	4,  // 4: vitess.queryservice.dev.Query.Rollback:input_type -> vitess.query.dev.RollbackRequest
	5,  // 5: vitess.queryservice.dev.Query.Prepare:input_type -> vitess.query.dev.PrepareRequest
	6,  // 6: vitess.queryservice.dev.Query.CommitPrepared:input_type -> vitess.query.dev.CommitPreparedRequest
	7,  // 7: vitess.queryservice.dev.Query.RollbackPrepared:input_type -> vitess.query.dev.RollbackPreparedRequest
	8,  // 8: vitess.queryservice.dev.Query.CreateTransaction:input_type -> vitess.query.dev.CreateTransactionRequest
	9,  // 9: vitess.queryservice.dev.Query.StartCommit:input_type -> vitess.query.dev.StartCommitRequest
	10, // 10: vitess.queryservice.dev.Query.SetRollback:input_type -> vitess.query.dev.SetRollbackRequest
	11, // 11: vitess.queryservice.dev.Query.ConcludeTransaction:input_type -> vitess.query.dev.ConcludeTransactionRequest
	12, // 12: vitess.queryservice.dev.Query.ReadTransaction:input_type -> vitess.query.dev.ReadTransactionRequest
	13, // 13: vitess.queryservice.dev.Query.UnresolvedTransactions:input_type -> vitess.query.dev.UnresolvedTransactionsRequest
	14, // 14: vitess.queryservice.dev.Query.BeginExecute:input_type -> vitess.query.dev.BeginExecuteRequest
	15, // 15: vitess.queryservice.dev.Query.BeginStreamExecute:input_type -> vitess.query.dev.BeginStreamExecuteRequest
	16, // 16: vitess.queryservice.dev.Query.MessageStream:input_type -> vitess.query.dev.MessageStreamRequest
	17, // 17: vitess.queryservice.dev.Query.MessageAck:input_type -> vitess.query.dev.MessageAckRequest
	18, // 18: vitess.queryservice.dev.Query.ReserveExecute:input_type -> vitess.query.dev.ReserveExecuteRequest
	19, // 19: vitess.queryservice.dev.Query.ReserveBeginExecute:input_type -> vitess.query.dev.ReserveBeginExecuteRequest
	20, // 20: vitess.queryservice.dev.Query.ReserveStreamExecute:input_type -> vitess.query.dev.ReserveStreamExecuteRequest
	21, // 21: vitess.queryservice.dev.Query.ReserveBeginStreamExecute:input_type -> vitess.query.dev.ReserveBeginStreamExecuteRequest
	22, // 22: vitess.queryservice.dev.Query.Release:input_type -> vitess.query.dev.ReleaseRequest
	23, // 23: vitess.queryservice.dev.Query.StreamHealth:input_type -> vitess.query.dev.StreamHealthRequest
	24, // 24: vitess.queryservice.dev.Query.VStream:input_type -> vitess.binlogdata.dev.VStreamRequest
	25, // 25: vitess.queryservice.dev.Query.VStreamRows:input_type -> vitess.binlogdata.dev.VStreamRowsRequest
	26, // 26: vitess.queryservice.dev.Query.VStreamTables:input_type -> vitess.binlogdata.dev.VStreamTablesRequest
	27, // 27: vitess.queryservice.dev.Query.VStreamResults:input_type -> vitess.binlogdata.dev.VStreamResultsRequest
	28, // 28: vitess.queryservice.dev.Query.GetSchema:input_type -> vitess.query.dev.GetSchemaRequest
	29, // 29: vitess.queryservice.dev.Query.Execute:output_type -> vitess.query.dev.ExecuteResponse
	30, // 30: vitess.queryservice.dev.Query.StreamExecute:output_type -> vitess.query.dev.StreamExecuteResponse
	31, // 31: vitess.queryservice.dev.Query.Begin:output_type -> vitess.query.dev.BeginResponse
	32, // 32: vitess.queryservice.dev.Query.Commit:output_type -> vitess.query.dev.CommitResponse
	33, // 33: vitess.queryservice.dev.Query.Rollback:output_type -> vitess.query.dev.RollbackResponse
	34, // 34: vitess.queryservice.dev.Query.Prepare:output_type -> vitess.query.dev.PrepareResponse
	35, // 35: vitess.queryservice.dev.Query.CommitPrepared:output_type -> vitess.query.dev.CommitPreparedResponse
	36, // 36: vitess.queryservice.dev.Query.RollbackPrepared:output_type -> vitess.query.dev.RollbackPreparedResponse
	37, // 37: vitess.queryservice.dev.Query.CreateTransaction:output_type -> vitess.query.dev.CreateTransactionResponse
	38, // 38: vitess.queryservice.dev.Query.StartCommit:output_type -> vitess.query.dev.StartCommitResponse
	39, // 39: vitess.queryservice.dev.Query.SetRollback:output_type -> vitess.query.dev.SetRollbackResponse
	40, // 40: vitess.queryservice.dev.Query.ConcludeTransaction:output_type -> vitess.query.dev.ConcludeTransactionResponse
	41, // 41: vitess.queryservice.dev.Query.ReadTransaction:output_type -> vitess.query.dev.ReadTransactionResponse
	42, // 42: vitess.queryservice.dev.Query.UnresolvedTransactions:output_type -> vitess.query.dev.UnresolvedTransactionsResponse
	43, // 43: vitess.queryservice.dev.Query.BeginExecute:output_type -> vitess.query.dev.BeginExecuteResponse
	44, // 44: vitess.queryservice.dev.Query.BeginStreamExecute:output_type -> vitess.query.dev.BeginStreamExecuteResponse
	45, // 45: vitess.queryservice.dev.Query.MessageStream:output_type -> vitess.query.dev.MessageStreamResponse
	46, // 46: vitess.queryservice.dev.Query.MessageAck:output_type -> vitess.query.dev.MessageAckResponse
	47, // 47: vitess.queryservice.dev.Query.ReserveExecute:output_type -> vitess.query.dev.ReserveExecuteResponse
	48, // 48: vitess.queryservice.dev.Query.ReserveBeginExecute:output_type -> vitess.query.dev.ReserveBeginExecuteResponse
	49, // 49: vitess.queryservice.dev.Query.ReserveStreamExecute:output_type -> vitess.query.dev.ReserveStreamExecuteResponse
	50, // 50: vitess.queryservice.dev.Query.ReserveBeginStreamExecute:output_type -> vitess.query.dev.ReserveBeginStreamExecuteResponse
	51, // 51: vitess.queryservice.dev.Query.Release:output_type -> vitess.query.dev.ReleaseResponse
	52, // 52: vitess.queryservice.dev.Query.StreamHealth:output_type -> vitess.query.dev.StreamHealthResponse
	53, // 53: vitess.queryservice.dev.Query.VStream:output_type -> vitess.binlogdata.dev.VStreamResponse
	54, // 54: vitess.queryservice.dev.Query.VStreamRows:output_type -> vitess.binlogdata.dev.VStreamRowsResponse
	55, // 55: vitess.queryservice.dev.Query.VStreamTables:output_type -> vitess.binlogdata.dev.VStreamTablesResponse
	56, // 56: vitess.queryservice.dev.Query.VStreamResults:output_type -> vitess.binlogdata.dev.VStreamResultsResponse
	57, // 57: vitess.queryservice.dev.Query.GetSchema:output_type -> vitess.query.dev.GetSchemaResponse
	29, // [29:58] is the sub-list for method output_type
	0,  // [0:29] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_vitess_queryservice_dev_queryservice_proto_init() }
func file_vitess_queryservice_dev_queryservice_proto_init() {
	if File_vitess_queryservice_dev_queryservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_vitess_queryservice_dev_queryservice_proto_rawDesc), len(file_vitess_queryservice_dev_queryservice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vitess_queryservice_dev_queryservice_proto_goTypes,
		DependencyIndexes: file_vitess_queryservice_dev_queryservice_proto_depIdxs,
	}.Build()
	File_vitess_queryservice_dev_queryservice_proto = out.File
	file_vitess_queryservice_dev_queryservice_proto_goTypes = nil
	file_vitess_queryservice_dev_queryservice_proto_depIdxs = nil
}
