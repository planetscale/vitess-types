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

// This file contains the service definition for making management API
// calls to VtTablet.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: vitess/tabletmanagerservice/v19/tabletmanagerservice.proto

package tabletmanagerservicev19

import (
	v19 "github.com/planetscale/vitess-types/gen/vitess/tabletmanagerdata/v19"
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

var File_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto protoreflect.FileDescriptor

const file_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto_rawDesc = "" +
	"\n" +
	":vitess/tabletmanagerservice/v19/tabletmanagerservice.proto\x12\x1fvitess.tabletmanagerservice.v19\x1a4vitess/tabletmanagerdata/v19/tabletmanagerdata.proto2\xd45\n" +
	"\rTabletManager\x12_\n" +
	"\x04Ping\x12).vitess.tabletmanagerdata.v19.PingRequest\x1a*.vitess.tabletmanagerdata.v19.PingResponse\"\x00\x12b\n" +
	"\x05Sleep\x12*.vitess.tabletmanagerdata.v19.SleepRequest\x1a+.vitess.tabletmanagerdata.v19.SleepResponse\"\x00\x12t\n" +
	"\vExecuteHook\x120.vitess.tabletmanagerdata.v19.ExecuteHookRequest\x1a1.vitess.tabletmanagerdata.v19.ExecuteHookResponse\"\x00\x12n\n" +
	"\tGetSchema\x12..vitess.tabletmanagerdata.v19.GetSchemaRequest\x1a/.vitess.tabletmanagerdata.v19.GetSchemaResponse\"\x00\x12}\n" +
	"\x0eGetPermissions\x123.vitess.tabletmanagerdata.v19.GetPermissionsRequest\x1a4.vitess.tabletmanagerdata.v19.GetPermissionsResponse\"\x00\x12t\n" +
	"\vSetReadOnly\x120.vitess.tabletmanagerdata.v19.SetReadOnlyRequest\x1a1.vitess.tabletmanagerdata.v19.SetReadOnlyResponse\"\x00\x12w\n" +
	"\fSetReadWrite\x121.vitess.tabletmanagerdata.v19.SetReadWriteRequest\x1a2.vitess.tabletmanagerdata.v19.SetReadWriteResponse\"\x00\x12q\n" +
	"\n" +
	"ChangeType\x12/.vitess.tabletmanagerdata.v19.ChangeTypeRequest\x1a0.vitess.tabletmanagerdata.v19.ChangeTypeResponse\"\x00\x12w\n" +
	"\fRefreshState\x121.vitess.tabletmanagerdata.v19.RefreshStateRequest\x1a2.vitess.tabletmanagerdata.v19.RefreshStateResponse\"\x00\x12}\n" +
	"\x0eRunHealthCheck\x123.vitess.tabletmanagerdata.v19.RunHealthCheckRequest\x1a4.vitess.tabletmanagerdata.v19.RunHealthCheckResponse\"\x00\x12w\n" +
	"\fReloadSchema\x121.vitess.tabletmanagerdata.v19.ReloadSchemaRequest\x1a2.vitess.tabletmanagerdata.v19.ReloadSchemaResponse\"\x00\x12\x80\x01\n" +
	"\x0fPreflightSchema\x124.vitess.tabletmanagerdata.v19.PreflightSchemaRequest\x1a5.vitess.tabletmanagerdata.v19.PreflightSchemaResponse\"\x00\x12t\n" +
	"\vApplySchema\x120.vitess.tabletmanagerdata.v19.ApplySchemaRequest\x1a1.vitess.tabletmanagerdata.v19.ApplySchemaResponse\"\x00\x12}\n" +
	"\x0eResetSequences\x123.vitess.tabletmanagerdata.v19.ResetSequencesRequest\x1a4.vitess.tabletmanagerdata.v19.ResetSequencesResponse\"\x00\x12q\n" +
	"\n" +
	"LockTables\x12/.vitess.tabletmanagerdata.v19.LockTablesRequest\x1a0.vitess.tabletmanagerdata.v19.LockTablesResponse\"\x00\x12w\n" +
	"\fUnlockTables\x121.vitess.tabletmanagerdata.v19.UnlockTablesRequest\x1a2.vitess.tabletmanagerdata.v19.UnlockTablesResponse\"\x00\x12w\n" +
	"\fExecuteQuery\x121.vitess.tabletmanagerdata.v19.ExecuteQueryRequest\x1a2.vitess.tabletmanagerdata.v19.ExecuteQueryResponse\"\x00\x12\x86\x01\n" +
	"\x11ExecuteFetchAsDba\x126.vitess.tabletmanagerdata.v19.ExecuteFetchAsDbaRequest\x1a7.vitess.tabletmanagerdata.v19.ExecuteFetchAsDbaResponse\"\x00\x12\x95\x01\n" +
	"\x16ExecuteFetchAsAllPrivs\x12;.vitess.tabletmanagerdata.v19.ExecuteFetchAsAllPrivsRequest\x1a<.vitess.tabletmanagerdata.v19.ExecuteFetchAsAllPrivsResponse\"\x00\x12\x86\x01\n" +
	"\x11ExecuteFetchAsApp\x126.vitess.tabletmanagerdata.v19.ExecuteFetchAsAppRequest\x1a7.vitess.tabletmanagerdata.v19.ExecuteFetchAsAppResponse\"\x00\x12\x86\x01\n" +
	"\x11ReplicationStatus\x126.vitess.tabletmanagerdata.v19.ReplicationStatusRequest\x1a7.vitess.tabletmanagerdata.v19.ReplicationStatusResponse\"\x00\x12z\n" +
	"\rPrimaryStatus\x122.vitess.tabletmanagerdata.v19.PrimaryStatusRequest\x1a3.vitess.tabletmanagerdata.v19.PrimaryStatusResponse\"\x00\x12\x80\x01\n" +
	"\x0fPrimaryPosition\x124.vitess.tabletmanagerdata.v19.PrimaryPositionRequest\x1a5.vitess.tabletmanagerdata.v19.PrimaryPositionResponse\"\x00\x12\x80\x01\n" +
	"\x0fWaitForPosition\x124.vitess.tabletmanagerdata.v19.WaitForPositionRequest\x1a5.vitess.tabletmanagerdata.v19.WaitForPositionResponse\"\x00\x12\x80\x01\n" +
	"\x0fStopReplication\x124.vitess.tabletmanagerdata.v19.StopReplicationRequest\x1a5.vitess.tabletmanagerdata.v19.StopReplicationResponse\"\x00\x12\x95\x01\n" +
	"\x16StopReplicationMinimum\x12;.vitess.tabletmanagerdata.v19.StopReplicationMinimumRequest\x1a<.vitess.tabletmanagerdata.v19.StopReplicationMinimumResponse\"\x00\x12\x83\x01\n" +
	"\x10StartReplication\x125.vitess.tabletmanagerdata.v19.StartReplicationRequest\x1a6.vitess.tabletmanagerdata.v19.StartReplicationResponse\"\x00\x12\xa1\x01\n" +
	"\x1aStartReplicationUntilAfter\x12?.vitess.tabletmanagerdata.v19.StartReplicationUntilAfterRequest\x1a@.vitess.tabletmanagerdata.v19.StartReplicationUntilAfterResponse\"\x00\x12t\n" +
	"\vGetReplicas\x120.vitess.tabletmanagerdata.v19.GetReplicasRequest\x1a1.vitess.tabletmanagerdata.v19.GetReplicasResponse\"\x00\x12\xa1\x01\n" +
	"\x1aCreateVReplicationWorkflow\x12?.vitess.tabletmanagerdata.v19.CreateVReplicationWorkflowRequest\x1a@.vitess.tabletmanagerdata.v19.CreateVReplicationWorkflowResponse\"\x00\x12\xa1\x01\n" +
	"\x1aDeleteVReplicationWorkflow\x12?.vitess.tabletmanagerdata.v19.DeleteVReplicationWorkflowRequest\x1a@.vitess.tabletmanagerdata.v19.DeleteVReplicationWorkflowResponse\"\x00\x12\x9b\x01\n" +
	"\x18ReadVReplicationWorkflow\x12=.vitess.tabletmanagerdata.v19.ReadVReplicationWorkflowRequest\x1a>.vitess.tabletmanagerdata.v19.ReadVReplicationWorkflowResponse\"\x00\x12\x83\x01\n" +
	"\x10VReplicationExec\x125.vitess.tabletmanagerdata.v19.VReplicationExecRequest\x1a6.vitess.tabletmanagerdata.v19.VReplicationExecResponse\"\x00\x12\x95\x01\n" +
	"\x16VReplicationWaitForPos\x12;.vitess.tabletmanagerdata.v19.VReplicationWaitForPosRequest\x1a<.vitess.tabletmanagerdata.v19.VReplicationWaitForPosResponse\"\x00\x12\xa1\x01\n" +
	"\x1aUpdateVReplicationWorkflow\x12?.vitess.tabletmanagerdata.v19.UpdateVReplicationWorkflowRequest\x1a@.vitess.tabletmanagerdata.v19.UpdateVReplicationWorkflowResponse\"\x00\x12b\n" +
	"\x05VDiff\x12*.vitess.tabletmanagerdata.v19.VDiffRequest\x1a+.vitess.tabletmanagerdata.v19.VDiffResponse\"\x00\x12\x83\x01\n" +
	"\x10ResetReplication\x125.vitess.tabletmanagerdata.v19.ResetReplicationRequest\x1a6.vitess.tabletmanagerdata.v19.ResetReplicationResponse\"\x00\x12t\n" +
	"\vInitPrimary\x120.vitess.tabletmanagerdata.v19.InitPrimaryRequest\x1a1.vitess.tabletmanagerdata.v19.InitPrimaryResponse\"\x00\x12\x98\x01\n" +
	"\x17PopulateReparentJournal\x12<.vitess.tabletmanagerdata.v19.PopulateReparentJournalRequest\x1a=.vitess.tabletmanagerdata.v19.PopulateReparentJournalResponse\"\x00\x12t\n" +
	"\vInitReplica\x120.vitess.tabletmanagerdata.v19.InitReplicaRequest\x1a1.vitess.tabletmanagerdata.v19.InitReplicaResponse\"\x00\x12z\n" +
	"\rDemotePrimary\x122.vitess.tabletmanagerdata.v19.DemotePrimaryRequest\x1a3.vitess.tabletmanagerdata.v19.DemotePrimaryResponse\"\x00\x12\x86\x01\n" +
	"\x11UndoDemotePrimary\x126.vitess.tabletmanagerdata.v19.UndoDemotePrimaryRequest\x1a7.vitess.tabletmanagerdata.v19.UndoDemotePrimaryResponse\"\x00\x12\x89\x01\n" +
	"\x12ReplicaWasPromoted\x127.vitess.tabletmanagerdata.v19.ReplicaWasPromotedRequest\x1a8.vitess.tabletmanagerdata.v19.ReplicaWasPromotedResponse\"\x00\x12\xa1\x01\n" +
	"\x1aResetReplicationParameters\x12?.vitess.tabletmanagerdata.v19.ResetReplicationParametersRequest\x1a@.vitess.tabletmanagerdata.v19.ResetReplicationParametersResponse\"\x00\x12q\n" +
	"\n" +
	"FullStatus\x12/.vitess.tabletmanagerdata.v19.FullStatusRequest\x1a0.vitess.tabletmanagerdata.v19.FullStatusResponse\"\x00\x12\x8f\x01\n" +
	"\x14SetReplicationSource\x129.vitess.tabletmanagerdata.v19.SetReplicationSourceRequest\x1a:.vitess.tabletmanagerdata.v19.SetReplicationSourceResponse\"\x00\x12\x8c\x01\n" +
	"\x13ReplicaWasRestarted\x128.vitess.tabletmanagerdata.v19.ReplicaWasRestartedRequest\x1a9.vitess.tabletmanagerdata.v19.ReplicaWasRestartedResponse\"\x00\x12\xa4\x01\n" +
	"\x1bStopReplicationAndGetStatus\x12@.vitess.tabletmanagerdata.v19.StopReplicationAndGetStatusRequest\x1aA.vitess.tabletmanagerdata.v19.StopReplicationAndGetStatusResponse\"\x00\x12}\n" +
	"\x0ePromoteReplica\x123.vitess.tabletmanagerdata.v19.PromoteReplicaRequest\x1a4.vitess.tabletmanagerdata.v19.PromoteReplicaResponse\"\x00\x12g\n" +
	"\x06Backup\x12+.vitess.tabletmanagerdata.v19.BackupRequest\x1a,.vitess.tabletmanagerdata.v19.BackupResponse\"\x000\x01\x12\x88\x01\n" +
	"\x11RestoreFromBackup\x126.vitess.tabletmanagerdata.v19.RestoreFromBackupRequest\x1a7.vitess.tabletmanagerdata.v19.RestoreFromBackupResponse\"\x000\x01\x12}\n" +
	"\x0eCheckThrottler\x123.vitess.tabletmanagerdata.v19.CheckThrottlerRequest\x1a4.vitess.tabletmanagerdata.v19.CheckThrottlerResponse\"\x00BaZ_github.com/planetscale/vitess-types/gen/vitess/tabletmanagerservice/v19;tabletmanagerservicev19b\x06proto3"

var file_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto_goTypes = []any{
	(*v19.PingRequest)(nil),                         // 0: vitess.tabletmanagerdata.v19.PingRequest
	(*v19.SleepRequest)(nil),                        // 1: vitess.tabletmanagerdata.v19.SleepRequest
	(*v19.ExecuteHookRequest)(nil),                  // 2: vitess.tabletmanagerdata.v19.ExecuteHookRequest
	(*v19.GetSchemaRequest)(nil),                    // 3: vitess.tabletmanagerdata.v19.GetSchemaRequest
	(*v19.GetPermissionsRequest)(nil),               // 4: vitess.tabletmanagerdata.v19.GetPermissionsRequest
	(*v19.SetReadOnlyRequest)(nil),                  // 5: vitess.tabletmanagerdata.v19.SetReadOnlyRequest
	(*v19.SetReadWriteRequest)(nil),                 // 6: vitess.tabletmanagerdata.v19.SetReadWriteRequest
	(*v19.ChangeTypeRequest)(nil),                   // 7: vitess.tabletmanagerdata.v19.ChangeTypeRequest
	(*v19.RefreshStateRequest)(nil),                 // 8: vitess.tabletmanagerdata.v19.RefreshStateRequest
	(*v19.RunHealthCheckRequest)(nil),               // 9: vitess.tabletmanagerdata.v19.RunHealthCheckRequest
	(*v19.ReloadSchemaRequest)(nil),                 // 10: vitess.tabletmanagerdata.v19.ReloadSchemaRequest
	(*v19.PreflightSchemaRequest)(nil),              // 11: vitess.tabletmanagerdata.v19.PreflightSchemaRequest
	(*v19.ApplySchemaRequest)(nil),                  // 12: vitess.tabletmanagerdata.v19.ApplySchemaRequest
	(*v19.ResetSequencesRequest)(nil),               // 13: vitess.tabletmanagerdata.v19.ResetSequencesRequest
	(*v19.LockTablesRequest)(nil),                   // 14: vitess.tabletmanagerdata.v19.LockTablesRequest
	(*v19.UnlockTablesRequest)(nil),                 // 15: vitess.tabletmanagerdata.v19.UnlockTablesRequest
	(*v19.ExecuteQueryRequest)(nil),                 // 16: vitess.tabletmanagerdata.v19.ExecuteQueryRequest
	(*v19.ExecuteFetchAsDbaRequest)(nil),            // 17: vitess.tabletmanagerdata.v19.ExecuteFetchAsDbaRequest
	(*v19.ExecuteFetchAsAllPrivsRequest)(nil),       // 18: vitess.tabletmanagerdata.v19.ExecuteFetchAsAllPrivsRequest
	(*v19.ExecuteFetchAsAppRequest)(nil),            // 19: vitess.tabletmanagerdata.v19.ExecuteFetchAsAppRequest
	(*v19.ReplicationStatusRequest)(nil),            // 20: vitess.tabletmanagerdata.v19.ReplicationStatusRequest
	(*v19.PrimaryStatusRequest)(nil),                // 21: vitess.tabletmanagerdata.v19.PrimaryStatusRequest
	(*v19.PrimaryPositionRequest)(nil),              // 22: vitess.tabletmanagerdata.v19.PrimaryPositionRequest
	(*v19.WaitForPositionRequest)(nil),              // 23: vitess.tabletmanagerdata.v19.WaitForPositionRequest
	(*v19.StopReplicationRequest)(nil),              // 24: vitess.tabletmanagerdata.v19.StopReplicationRequest
	(*v19.StopReplicationMinimumRequest)(nil),       // 25: vitess.tabletmanagerdata.v19.StopReplicationMinimumRequest
	(*v19.StartReplicationRequest)(nil),             // 26: vitess.tabletmanagerdata.v19.StartReplicationRequest
	(*v19.StartReplicationUntilAfterRequest)(nil),   // 27: vitess.tabletmanagerdata.v19.StartReplicationUntilAfterRequest
	(*v19.GetReplicasRequest)(nil),                  // 28: vitess.tabletmanagerdata.v19.GetReplicasRequest
	(*v19.CreateVReplicationWorkflowRequest)(nil),   // 29: vitess.tabletmanagerdata.v19.CreateVReplicationWorkflowRequest
	(*v19.DeleteVReplicationWorkflowRequest)(nil),   // 30: vitess.tabletmanagerdata.v19.DeleteVReplicationWorkflowRequest
	(*v19.ReadVReplicationWorkflowRequest)(nil),     // 31: vitess.tabletmanagerdata.v19.ReadVReplicationWorkflowRequest
	(*v19.VReplicationExecRequest)(nil),             // 32: vitess.tabletmanagerdata.v19.VReplicationExecRequest
	(*v19.VReplicationWaitForPosRequest)(nil),       // 33: vitess.tabletmanagerdata.v19.VReplicationWaitForPosRequest
	(*v19.UpdateVReplicationWorkflowRequest)(nil),   // 34: vitess.tabletmanagerdata.v19.UpdateVReplicationWorkflowRequest
	(*v19.VDiffRequest)(nil),                        // 35: vitess.tabletmanagerdata.v19.VDiffRequest
	(*v19.ResetReplicationRequest)(nil),             // 36: vitess.tabletmanagerdata.v19.ResetReplicationRequest
	(*v19.InitPrimaryRequest)(nil),                  // 37: vitess.tabletmanagerdata.v19.InitPrimaryRequest
	(*v19.PopulateReparentJournalRequest)(nil),      // 38: vitess.tabletmanagerdata.v19.PopulateReparentJournalRequest
	(*v19.InitReplicaRequest)(nil),                  // 39: vitess.tabletmanagerdata.v19.InitReplicaRequest
	(*v19.DemotePrimaryRequest)(nil),                // 40: vitess.tabletmanagerdata.v19.DemotePrimaryRequest
	(*v19.UndoDemotePrimaryRequest)(nil),            // 41: vitess.tabletmanagerdata.v19.UndoDemotePrimaryRequest
	(*v19.ReplicaWasPromotedRequest)(nil),           // 42: vitess.tabletmanagerdata.v19.ReplicaWasPromotedRequest
	(*v19.ResetReplicationParametersRequest)(nil),   // 43: vitess.tabletmanagerdata.v19.ResetReplicationParametersRequest
	(*v19.FullStatusRequest)(nil),                   // 44: vitess.tabletmanagerdata.v19.FullStatusRequest
	(*v19.SetReplicationSourceRequest)(nil),         // 45: vitess.tabletmanagerdata.v19.SetReplicationSourceRequest
	(*v19.ReplicaWasRestartedRequest)(nil),          // 46: vitess.tabletmanagerdata.v19.ReplicaWasRestartedRequest
	(*v19.StopReplicationAndGetStatusRequest)(nil),  // 47: vitess.tabletmanagerdata.v19.StopReplicationAndGetStatusRequest
	(*v19.PromoteReplicaRequest)(nil),               // 48: vitess.tabletmanagerdata.v19.PromoteReplicaRequest
	(*v19.BackupRequest)(nil),                       // 49: vitess.tabletmanagerdata.v19.BackupRequest
	(*v19.RestoreFromBackupRequest)(nil),            // 50: vitess.tabletmanagerdata.v19.RestoreFromBackupRequest
	(*v19.CheckThrottlerRequest)(nil),               // 51: vitess.tabletmanagerdata.v19.CheckThrottlerRequest
	(*v19.PingResponse)(nil),                        // 52: vitess.tabletmanagerdata.v19.PingResponse
	(*v19.SleepResponse)(nil),                       // 53: vitess.tabletmanagerdata.v19.SleepResponse
	(*v19.ExecuteHookResponse)(nil),                 // 54: vitess.tabletmanagerdata.v19.ExecuteHookResponse
	(*v19.GetSchemaResponse)(nil),                   // 55: vitess.tabletmanagerdata.v19.GetSchemaResponse
	(*v19.GetPermissionsResponse)(nil),              // 56: vitess.tabletmanagerdata.v19.GetPermissionsResponse
	(*v19.SetReadOnlyResponse)(nil),                 // 57: vitess.tabletmanagerdata.v19.SetReadOnlyResponse
	(*v19.SetReadWriteResponse)(nil),                // 58: vitess.tabletmanagerdata.v19.SetReadWriteResponse
	(*v19.ChangeTypeResponse)(nil),                  // 59: vitess.tabletmanagerdata.v19.ChangeTypeResponse
	(*v19.RefreshStateResponse)(nil),                // 60: vitess.tabletmanagerdata.v19.RefreshStateResponse
	(*v19.RunHealthCheckResponse)(nil),              // 61: vitess.tabletmanagerdata.v19.RunHealthCheckResponse
	(*v19.ReloadSchemaResponse)(nil),                // 62: vitess.tabletmanagerdata.v19.ReloadSchemaResponse
	(*v19.PreflightSchemaResponse)(nil),             // 63: vitess.tabletmanagerdata.v19.PreflightSchemaResponse
	(*v19.ApplySchemaResponse)(nil),                 // 64: vitess.tabletmanagerdata.v19.ApplySchemaResponse
	(*v19.ResetSequencesResponse)(nil),              // 65: vitess.tabletmanagerdata.v19.ResetSequencesResponse
	(*v19.LockTablesResponse)(nil),                  // 66: vitess.tabletmanagerdata.v19.LockTablesResponse
	(*v19.UnlockTablesResponse)(nil),                // 67: vitess.tabletmanagerdata.v19.UnlockTablesResponse
	(*v19.ExecuteQueryResponse)(nil),                // 68: vitess.tabletmanagerdata.v19.ExecuteQueryResponse
	(*v19.ExecuteFetchAsDbaResponse)(nil),           // 69: vitess.tabletmanagerdata.v19.ExecuteFetchAsDbaResponse
	(*v19.ExecuteFetchAsAllPrivsResponse)(nil),      // 70: vitess.tabletmanagerdata.v19.ExecuteFetchAsAllPrivsResponse
	(*v19.ExecuteFetchAsAppResponse)(nil),           // 71: vitess.tabletmanagerdata.v19.ExecuteFetchAsAppResponse
	(*v19.ReplicationStatusResponse)(nil),           // 72: vitess.tabletmanagerdata.v19.ReplicationStatusResponse
	(*v19.PrimaryStatusResponse)(nil),               // 73: vitess.tabletmanagerdata.v19.PrimaryStatusResponse
	(*v19.PrimaryPositionResponse)(nil),             // 74: vitess.tabletmanagerdata.v19.PrimaryPositionResponse
	(*v19.WaitForPositionResponse)(nil),             // 75: vitess.tabletmanagerdata.v19.WaitForPositionResponse
	(*v19.StopReplicationResponse)(nil),             // 76: vitess.tabletmanagerdata.v19.StopReplicationResponse
	(*v19.StopReplicationMinimumResponse)(nil),      // 77: vitess.tabletmanagerdata.v19.StopReplicationMinimumResponse
	(*v19.StartReplicationResponse)(nil),            // 78: vitess.tabletmanagerdata.v19.StartReplicationResponse
	(*v19.StartReplicationUntilAfterResponse)(nil),  // 79: vitess.tabletmanagerdata.v19.StartReplicationUntilAfterResponse
	(*v19.GetReplicasResponse)(nil),                 // 80: vitess.tabletmanagerdata.v19.GetReplicasResponse
	(*v19.CreateVReplicationWorkflowResponse)(nil),  // 81: vitess.tabletmanagerdata.v19.CreateVReplicationWorkflowResponse
	(*v19.DeleteVReplicationWorkflowResponse)(nil),  // 82: vitess.tabletmanagerdata.v19.DeleteVReplicationWorkflowResponse
	(*v19.ReadVReplicationWorkflowResponse)(nil),    // 83: vitess.tabletmanagerdata.v19.ReadVReplicationWorkflowResponse
	(*v19.VReplicationExecResponse)(nil),            // 84: vitess.tabletmanagerdata.v19.VReplicationExecResponse
	(*v19.VReplicationWaitForPosResponse)(nil),      // 85: vitess.tabletmanagerdata.v19.VReplicationWaitForPosResponse
	(*v19.UpdateVReplicationWorkflowResponse)(nil),  // 86: vitess.tabletmanagerdata.v19.UpdateVReplicationWorkflowResponse
	(*v19.VDiffResponse)(nil),                       // 87: vitess.tabletmanagerdata.v19.VDiffResponse
	(*v19.ResetReplicationResponse)(nil),            // 88: vitess.tabletmanagerdata.v19.ResetReplicationResponse
	(*v19.InitPrimaryResponse)(nil),                 // 89: vitess.tabletmanagerdata.v19.InitPrimaryResponse
	(*v19.PopulateReparentJournalResponse)(nil),     // 90: vitess.tabletmanagerdata.v19.PopulateReparentJournalResponse
	(*v19.InitReplicaResponse)(nil),                 // 91: vitess.tabletmanagerdata.v19.InitReplicaResponse
	(*v19.DemotePrimaryResponse)(nil),               // 92: vitess.tabletmanagerdata.v19.DemotePrimaryResponse
	(*v19.UndoDemotePrimaryResponse)(nil),           // 93: vitess.tabletmanagerdata.v19.UndoDemotePrimaryResponse
	(*v19.ReplicaWasPromotedResponse)(nil),          // 94: vitess.tabletmanagerdata.v19.ReplicaWasPromotedResponse
	(*v19.ResetReplicationParametersResponse)(nil),  // 95: vitess.tabletmanagerdata.v19.ResetReplicationParametersResponse
	(*v19.FullStatusResponse)(nil),                  // 96: vitess.tabletmanagerdata.v19.FullStatusResponse
	(*v19.SetReplicationSourceResponse)(nil),        // 97: vitess.tabletmanagerdata.v19.SetReplicationSourceResponse
	(*v19.ReplicaWasRestartedResponse)(nil),         // 98: vitess.tabletmanagerdata.v19.ReplicaWasRestartedResponse
	(*v19.StopReplicationAndGetStatusResponse)(nil), // 99: vitess.tabletmanagerdata.v19.StopReplicationAndGetStatusResponse
	(*v19.PromoteReplicaResponse)(nil),              // 100: vitess.tabletmanagerdata.v19.PromoteReplicaResponse
	(*v19.BackupResponse)(nil),                      // 101: vitess.tabletmanagerdata.v19.BackupResponse
	(*v19.RestoreFromBackupResponse)(nil),           // 102: vitess.tabletmanagerdata.v19.RestoreFromBackupResponse
	(*v19.CheckThrottlerResponse)(nil),              // 103: vitess.tabletmanagerdata.v19.CheckThrottlerResponse
}
var file_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto_depIdxs = []int32{
	0,   // 0: vitess.tabletmanagerservice.v19.TabletManager.Ping:input_type -> vitess.tabletmanagerdata.v19.PingRequest
	1,   // 1: vitess.tabletmanagerservice.v19.TabletManager.Sleep:input_type -> vitess.tabletmanagerdata.v19.SleepRequest
	2,   // 2: vitess.tabletmanagerservice.v19.TabletManager.ExecuteHook:input_type -> vitess.tabletmanagerdata.v19.ExecuteHookRequest
	3,   // 3: vitess.tabletmanagerservice.v19.TabletManager.GetSchema:input_type -> vitess.tabletmanagerdata.v19.GetSchemaRequest
	4,   // 4: vitess.tabletmanagerservice.v19.TabletManager.GetPermissions:input_type -> vitess.tabletmanagerdata.v19.GetPermissionsRequest
	5,   // 5: vitess.tabletmanagerservice.v19.TabletManager.SetReadOnly:input_type -> vitess.tabletmanagerdata.v19.SetReadOnlyRequest
	6,   // 6: vitess.tabletmanagerservice.v19.TabletManager.SetReadWrite:input_type -> vitess.tabletmanagerdata.v19.SetReadWriteRequest
	7,   // 7: vitess.tabletmanagerservice.v19.TabletManager.ChangeType:input_type -> vitess.tabletmanagerdata.v19.ChangeTypeRequest
	8,   // 8: vitess.tabletmanagerservice.v19.TabletManager.RefreshState:input_type -> vitess.tabletmanagerdata.v19.RefreshStateRequest
	9,   // 9: vitess.tabletmanagerservice.v19.TabletManager.RunHealthCheck:input_type -> vitess.tabletmanagerdata.v19.RunHealthCheckRequest
	10,  // 10: vitess.tabletmanagerservice.v19.TabletManager.ReloadSchema:input_type -> vitess.tabletmanagerdata.v19.ReloadSchemaRequest
	11,  // 11: vitess.tabletmanagerservice.v19.TabletManager.PreflightSchema:input_type -> vitess.tabletmanagerdata.v19.PreflightSchemaRequest
	12,  // 12: vitess.tabletmanagerservice.v19.TabletManager.ApplySchema:input_type -> vitess.tabletmanagerdata.v19.ApplySchemaRequest
	13,  // 13: vitess.tabletmanagerservice.v19.TabletManager.ResetSequences:input_type -> vitess.tabletmanagerdata.v19.ResetSequencesRequest
	14,  // 14: vitess.tabletmanagerservice.v19.TabletManager.LockTables:input_type -> vitess.tabletmanagerdata.v19.LockTablesRequest
	15,  // 15: vitess.tabletmanagerservice.v19.TabletManager.UnlockTables:input_type -> vitess.tabletmanagerdata.v19.UnlockTablesRequest
	16,  // 16: vitess.tabletmanagerservice.v19.TabletManager.ExecuteQuery:input_type -> vitess.tabletmanagerdata.v19.ExecuteQueryRequest
	17,  // 17: vitess.tabletmanagerservice.v19.TabletManager.ExecuteFetchAsDba:input_type -> vitess.tabletmanagerdata.v19.ExecuteFetchAsDbaRequest
	18,  // 18: vitess.tabletmanagerservice.v19.TabletManager.ExecuteFetchAsAllPrivs:input_type -> vitess.tabletmanagerdata.v19.ExecuteFetchAsAllPrivsRequest
	19,  // 19: vitess.tabletmanagerservice.v19.TabletManager.ExecuteFetchAsApp:input_type -> vitess.tabletmanagerdata.v19.ExecuteFetchAsAppRequest
	20,  // 20: vitess.tabletmanagerservice.v19.TabletManager.ReplicationStatus:input_type -> vitess.tabletmanagerdata.v19.ReplicationStatusRequest
	21,  // 21: vitess.tabletmanagerservice.v19.TabletManager.PrimaryStatus:input_type -> vitess.tabletmanagerdata.v19.PrimaryStatusRequest
	22,  // 22: vitess.tabletmanagerservice.v19.TabletManager.PrimaryPosition:input_type -> vitess.tabletmanagerdata.v19.PrimaryPositionRequest
	23,  // 23: vitess.tabletmanagerservice.v19.TabletManager.WaitForPosition:input_type -> vitess.tabletmanagerdata.v19.WaitForPositionRequest
	24,  // 24: vitess.tabletmanagerservice.v19.TabletManager.StopReplication:input_type -> vitess.tabletmanagerdata.v19.StopReplicationRequest
	25,  // 25: vitess.tabletmanagerservice.v19.TabletManager.StopReplicationMinimum:input_type -> vitess.tabletmanagerdata.v19.StopReplicationMinimumRequest
	26,  // 26: vitess.tabletmanagerservice.v19.TabletManager.StartReplication:input_type -> vitess.tabletmanagerdata.v19.StartReplicationRequest
	27,  // 27: vitess.tabletmanagerservice.v19.TabletManager.StartReplicationUntilAfter:input_type -> vitess.tabletmanagerdata.v19.StartReplicationUntilAfterRequest
	28,  // 28: vitess.tabletmanagerservice.v19.TabletManager.GetReplicas:input_type -> vitess.tabletmanagerdata.v19.GetReplicasRequest
	29,  // 29: vitess.tabletmanagerservice.v19.TabletManager.CreateVReplicationWorkflow:input_type -> vitess.tabletmanagerdata.v19.CreateVReplicationWorkflowRequest
	30,  // 30: vitess.tabletmanagerservice.v19.TabletManager.DeleteVReplicationWorkflow:input_type -> vitess.tabletmanagerdata.v19.DeleteVReplicationWorkflowRequest
	31,  // 31: vitess.tabletmanagerservice.v19.TabletManager.ReadVReplicationWorkflow:input_type -> vitess.tabletmanagerdata.v19.ReadVReplicationWorkflowRequest
	32,  // 32: vitess.tabletmanagerservice.v19.TabletManager.VReplicationExec:input_type -> vitess.tabletmanagerdata.v19.VReplicationExecRequest
	33,  // 33: vitess.tabletmanagerservice.v19.TabletManager.VReplicationWaitForPos:input_type -> vitess.tabletmanagerdata.v19.VReplicationWaitForPosRequest
	34,  // 34: vitess.tabletmanagerservice.v19.TabletManager.UpdateVReplicationWorkflow:input_type -> vitess.tabletmanagerdata.v19.UpdateVReplicationWorkflowRequest
	35,  // 35: vitess.tabletmanagerservice.v19.TabletManager.VDiff:input_type -> vitess.tabletmanagerdata.v19.VDiffRequest
	36,  // 36: vitess.tabletmanagerservice.v19.TabletManager.ResetReplication:input_type -> vitess.tabletmanagerdata.v19.ResetReplicationRequest
	37,  // 37: vitess.tabletmanagerservice.v19.TabletManager.InitPrimary:input_type -> vitess.tabletmanagerdata.v19.InitPrimaryRequest
	38,  // 38: vitess.tabletmanagerservice.v19.TabletManager.PopulateReparentJournal:input_type -> vitess.tabletmanagerdata.v19.PopulateReparentJournalRequest
	39,  // 39: vitess.tabletmanagerservice.v19.TabletManager.InitReplica:input_type -> vitess.tabletmanagerdata.v19.InitReplicaRequest
	40,  // 40: vitess.tabletmanagerservice.v19.TabletManager.DemotePrimary:input_type -> vitess.tabletmanagerdata.v19.DemotePrimaryRequest
	41,  // 41: vitess.tabletmanagerservice.v19.TabletManager.UndoDemotePrimary:input_type -> vitess.tabletmanagerdata.v19.UndoDemotePrimaryRequest
	42,  // 42: vitess.tabletmanagerservice.v19.TabletManager.ReplicaWasPromoted:input_type -> vitess.tabletmanagerdata.v19.ReplicaWasPromotedRequest
	43,  // 43: vitess.tabletmanagerservice.v19.TabletManager.ResetReplicationParameters:input_type -> vitess.tabletmanagerdata.v19.ResetReplicationParametersRequest
	44,  // 44: vitess.tabletmanagerservice.v19.TabletManager.FullStatus:input_type -> vitess.tabletmanagerdata.v19.FullStatusRequest
	45,  // 45: vitess.tabletmanagerservice.v19.TabletManager.SetReplicationSource:input_type -> vitess.tabletmanagerdata.v19.SetReplicationSourceRequest
	46,  // 46: vitess.tabletmanagerservice.v19.TabletManager.ReplicaWasRestarted:input_type -> vitess.tabletmanagerdata.v19.ReplicaWasRestartedRequest
	47,  // 47: vitess.tabletmanagerservice.v19.TabletManager.StopReplicationAndGetStatus:input_type -> vitess.tabletmanagerdata.v19.StopReplicationAndGetStatusRequest
	48,  // 48: vitess.tabletmanagerservice.v19.TabletManager.PromoteReplica:input_type -> vitess.tabletmanagerdata.v19.PromoteReplicaRequest
	49,  // 49: vitess.tabletmanagerservice.v19.TabletManager.Backup:input_type -> vitess.tabletmanagerdata.v19.BackupRequest
	50,  // 50: vitess.tabletmanagerservice.v19.TabletManager.RestoreFromBackup:input_type -> vitess.tabletmanagerdata.v19.RestoreFromBackupRequest
	51,  // 51: vitess.tabletmanagerservice.v19.TabletManager.CheckThrottler:input_type -> vitess.tabletmanagerdata.v19.CheckThrottlerRequest
	52,  // 52: vitess.tabletmanagerservice.v19.TabletManager.Ping:output_type -> vitess.tabletmanagerdata.v19.PingResponse
	53,  // 53: vitess.tabletmanagerservice.v19.TabletManager.Sleep:output_type -> vitess.tabletmanagerdata.v19.SleepResponse
	54,  // 54: vitess.tabletmanagerservice.v19.TabletManager.ExecuteHook:output_type -> vitess.tabletmanagerdata.v19.ExecuteHookResponse
	55,  // 55: vitess.tabletmanagerservice.v19.TabletManager.GetSchema:output_type -> vitess.tabletmanagerdata.v19.GetSchemaResponse
	56,  // 56: vitess.tabletmanagerservice.v19.TabletManager.GetPermissions:output_type -> vitess.tabletmanagerdata.v19.GetPermissionsResponse
	57,  // 57: vitess.tabletmanagerservice.v19.TabletManager.SetReadOnly:output_type -> vitess.tabletmanagerdata.v19.SetReadOnlyResponse
	58,  // 58: vitess.tabletmanagerservice.v19.TabletManager.SetReadWrite:output_type -> vitess.tabletmanagerdata.v19.SetReadWriteResponse
	59,  // 59: vitess.tabletmanagerservice.v19.TabletManager.ChangeType:output_type -> vitess.tabletmanagerdata.v19.ChangeTypeResponse
	60,  // 60: vitess.tabletmanagerservice.v19.TabletManager.RefreshState:output_type -> vitess.tabletmanagerdata.v19.RefreshStateResponse
	61,  // 61: vitess.tabletmanagerservice.v19.TabletManager.RunHealthCheck:output_type -> vitess.tabletmanagerdata.v19.RunHealthCheckResponse
	62,  // 62: vitess.tabletmanagerservice.v19.TabletManager.ReloadSchema:output_type -> vitess.tabletmanagerdata.v19.ReloadSchemaResponse
	63,  // 63: vitess.tabletmanagerservice.v19.TabletManager.PreflightSchema:output_type -> vitess.tabletmanagerdata.v19.PreflightSchemaResponse
	64,  // 64: vitess.tabletmanagerservice.v19.TabletManager.ApplySchema:output_type -> vitess.tabletmanagerdata.v19.ApplySchemaResponse
	65,  // 65: vitess.tabletmanagerservice.v19.TabletManager.ResetSequences:output_type -> vitess.tabletmanagerdata.v19.ResetSequencesResponse
	66,  // 66: vitess.tabletmanagerservice.v19.TabletManager.LockTables:output_type -> vitess.tabletmanagerdata.v19.LockTablesResponse
	67,  // 67: vitess.tabletmanagerservice.v19.TabletManager.UnlockTables:output_type -> vitess.tabletmanagerdata.v19.UnlockTablesResponse
	68,  // 68: vitess.tabletmanagerservice.v19.TabletManager.ExecuteQuery:output_type -> vitess.tabletmanagerdata.v19.ExecuteQueryResponse
	69,  // 69: vitess.tabletmanagerservice.v19.TabletManager.ExecuteFetchAsDba:output_type -> vitess.tabletmanagerdata.v19.ExecuteFetchAsDbaResponse
	70,  // 70: vitess.tabletmanagerservice.v19.TabletManager.ExecuteFetchAsAllPrivs:output_type -> vitess.tabletmanagerdata.v19.ExecuteFetchAsAllPrivsResponse
	71,  // 71: vitess.tabletmanagerservice.v19.TabletManager.ExecuteFetchAsApp:output_type -> vitess.tabletmanagerdata.v19.ExecuteFetchAsAppResponse
	72,  // 72: vitess.tabletmanagerservice.v19.TabletManager.ReplicationStatus:output_type -> vitess.tabletmanagerdata.v19.ReplicationStatusResponse
	73,  // 73: vitess.tabletmanagerservice.v19.TabletManager.PrimaryStatus:output_type -> vitess.tabletmanagerdata.v19.PrimaryStatusResponse
	74,  // 74: vitess.tabletmanagerservice.v19.TabletManager.PrimaryPosition:output_type -> vitess.tabletmanagerdata.v19.PrimaryPositionResponse
	75,  // 75: vitess.tabletmanagerservice.v19.TabletManager.WaitForPosition:output_type -> vitess.tabletmanagerdata.v19.WaitForPositionResponse
	76,  // 76: vitess.tabletmanagerservice.v19.TabletManager.StopReplication:output_type -> vitess.tabletmanagerdata.v19.StopReplicationResponse
	77,  // 77: vitess.tabletmanagerservice.v19.TabletManager.StopReplicationMinimum:output_type -> vitess.tabletmanagerdata.v19.StopReplicationMinimumResponse
	78,  // 78: vitess.tabletmanagerservice.v19.TabletManager.StartReplication:output_type -> vitess.tabletmanagerdata.v19.StartReplicationResponse
	79,  // 79: vitess.tabletmanagerservice.v19.TabletManager.StartReplicationUntilAfter:output_type -> vitess.tabletmanagerdata.v19.StartReplicationUntilAfterResponse
	80,  // 80: vitess.tabletmanagerservice.v19.TabletManager.GetReplicas:output_type -> vitess.tabletmanagerdata.v19.GetReplicasResponse
	81,  // 81: vitess.tabletmanagerservice.v19.TabletManager.CreateVReplicationWorkflow:output_type -> vitess.tabletmanagerdata.v19.CreateVReplicationWorkflowResponse
	82,  // 82: vitess.tabletmanagerservice.v19.TabletManager.DeleteVReplicationWorkflow:output_type -> vitess.tabletmanagerdata.v19.DeleteVReplicationWorkflowResponse
	83,  // 83: vitess.tabletmanagerservice.v19.TabletManager.ReadVReplicationWorkflow:output_type -> vitess.tabletmanagerdata.v19.ReadVReplicationWorkflowResponse
	84,  // 84: vitess.tabletmanagerservice.v19.TabletManager.VReplicationExec:output_type -> vitess.tabletmanagerdata.v19.VReplicationExecResponse
	85,  // 85: vitess.tabletmanagerservice.v19.TabletManager.VReplicationWaitForPos:output_type -> vitess.tabletmanagerdata.v19.VReplicationWaitForPosResponse
	86,  // 86: vitess.tabletmanagerservice.v19.TabletManager.UpdateVReplicationWorkflow:output_type -> vitess.tabletmanagerdata.v19.UpdateVReplicationWorkflowResponse
	87,  // 87: vitess.tabletmanagerservice.v19.TabletManager.VDiff:output_type -> vitess.tabletmanagerdata.v19.VDiffResponse
	88,  // 88: vitess.tabletmanagerservice.v19.TabletManager.ResetReplication:output_type -> vitess.tabletmanagerdata.v19.ResetReplicationResponse
	89,  // 89: vitess.tabletmanagerservice.v19.TabletManager.InitPrimary:output_type -> vitess.tabletmanagerdata.v19.InitPrimaryResponse
	90,  // 90: vitess.tabletmanagerservice.v19.TabletManager.PopulateReparentJournal:output_type -> vitess.tabletmanagerdata.v19.PopulateReparentJournalResponse
	91,  // 91: vitess.tabletmanagerservice.v19.TabletManager.InitReplica:output_type -> vitess.tabletmanagerdata.v19.InitReplicaResponse
	92,  // 92: vitess.tabletmanagerservice.v19.TabletManager.DemotePrimary:output_type -> vitess.tabletmanagerdata.v19.DemotePrimaryResponse
	93,  // 93: vitess.tabletmanagerservice.v19.TabletManager.UndoDemotePrimary:output_type -> vitess.tabletmanagerdata.v19.UndoDemotePrimaryResponse
	94,  // 94: vitess.tabletmanagerservice.v19.TabletManager.ReplicaWasPromoted:output_type -> vitess.tabletmanagerdata.v19.ReplicaWasPromotedResponse
	95,  // 95: vitess.tabletmanagerservice.v19.TabletManager.ResetReplicationParameters:output_type -> vitess.tabletmanagerdata.v19.ResetReplicationParametersResponse
	96,  // 96: vitess.tabletmanagerservice.v19.TabletManager.FullStatus:output_type -> vitess.tabletmanagerdata.v19.FullStatusResponse
	97,  // 97: vitess.tabletmanagerservice.v19.TabletManager.SetReplicationSource:output_type -> vitess.tabletmanagerdata.v19.SetReplicationSourceResponse
	98,  // 98: vitess.tabletmanagerservice.v19.TabletManager.ReplicaWasRestarted:output_type -> vitess.tabletmanagerdata.v19.ReplicaWasRestartedResponse
	99,  // 99: vitess.tabletmanagerservice.v19.TabletManager.StopReplicationAndGetStatus:output_type -> vitess.tabletmanagerdata.v19.StopReplicationAndGetStatusResponse
	100, // 100: vitess.tabletmanagerservice.v19.TabletManager.PromoteReplica:output_type -> vitess.tabletmanagerdata.v19.PromoteReplicaResponse
	101, // 101: vitess.tabletmanagerservice.v19.TabletManager.Backup:output_type -> vitess.tabletmanagerdata.v19.BackupResponse
	102, // 102: vitess.tabletmanagerservice.v19.TabletManager.RestoreFromBackup:output_type -> vitess.tabletmanagerdata.v19.RestoreFromBackupResponse
	103, // 103: vitess.tabletmanagerservice.v19.TabletManager.CheckThrottler:output_type -> vitess.tabletmanagerdata.v19.CheckThrottlerResponse
	52,  // [52:104] is the sub-list for method output_type
	0,   // [0:52] is the sub-list for method input_type
	0,   // [0:0] is the sub-list for extension type_name
	0,   // [0:0] is the sub-list for extension extendee
	0,   // [0:0] is the sub-list for field type_name
}

func init() { file_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto_init() }
func file_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto_init() {
	if File_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto_rawDesc), len(file_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto_goTypes,
		DependencyIndexes: file_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto_depIdxs,
	}.Build()
	File_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto = out.File
	file_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto_goTypes = nil
	file_vitess_tabletmanagerservice_v19_tabletmanagerservice_proto_depIdxs = nil
}
