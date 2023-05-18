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

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: vitess/vtgateservice/dev/vtgateservice.proto

// option java_package="io.vitess.proto.grpc";

package vtgateservicedevconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	dev "github.com/planetscale/vitess-types/gen/vitess/vtgate/dev"
	_ "github.com/planetscale/vitess-types/gen/vitess/vtgateservice/dev"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// VitessName is the fully-qualified name of the Vitess service.
	VitessName = "vitess.vtgateservice.dev.Vitess"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// VitessExecuteProcedure is the fully-qualified name of the Vitess's Execute RPC.
	VitessExecuteProcedure = "/vitess.vtgateservice.dev.Vitess/Execute"
	// VitessExecuteBatchProcedure is the fully-qualified name of the Vitess's ExecuteBatch RPC.
	VitessExecuteBatchProcedure = "/vitess.vtgateservice.dev.Vitess/ExecuteBatch"
	// VitessStreamExecuteProcedure is the fully-qualified name of the Vitess's StreamExecute RPC.
	VitessStreamExecuteProcedure = "/vitess.vtgateservice.dev.Vitess/StreamExecute"
	// VitessResolveTransactionProcedure is the fully-qualified name of the Vitess's ResolveTransaction
	// RPC.
	VitessResolveTransactionProcedure = "/vitess.vtgateservice.dev.Vitess/ResolveTransaction"
	// VitessVStreamProcedure is the fully-qualified name of the Vitess's VStream RPC.
	VitessVStreamProcedure = "/vitess.vtgateservice.dev.Vitess/VStream"
	// VitessPrepareProcedure is the fully-qualified name of the Vitess's Prepare RPC.
	VitessPrepareProcedure = "/vitess.vtgateservice.dev.Vitess/Prepare"
	// VitessCloseSessionProcedure is the fully-qualified name of the Vitess's CloseSession RPC.
	VitessCloseSessionProcedure = "/vitess.vtgateservice.dev.Vitess/CloseSession"
)

// VitessClient is a client for the vitess.vtgateservice.dev.Vitess service.
type VitessClient interface {
	// Execute tries to route the query to the right shard.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// API group: v3
	Execute(context.Context, *connect_go.Request[dev.ExecuteRequest]) (*connect_go.Response[dev.ExecuteResponse], error)
	// ExecuteBatch tries to route the list of queries on the right shards.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// API group: v3
	ExecuteBatch(context.Context, *connect_go.Request[dev.ExecuteBatchRequest]) (*connect_go.Response[dev.ExecuteBatchResponse], error)
	// StreamExecute executes a streaming query based on shards.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// Use this method if the query returns a large number of rows.
	// API group: v3
	StreamExecute(context.Context, *connect_go.Request[dev.StreamExecuteRequest]) (*connect_go.ServerStreamForClient[dev.StreamExecuteResponse], error)
	// ResolveTransaction resolves a transaction.
	// API group: Transactions
	ResolveTransaction(context.Context, *connect_go.Request[dev.ResolveTransactionRequest]) (*connect_go.Response[dev.ResolveTransactionResponse], error)
	// VStream streams binlog events from the requested sources.
	VStream(context.Context, *connect_go.Request[dev.VStreamRequest]) (*connect_go.ServerStreamForClient[dev.VStreamResponse], error)
	// Prepare is used by the MySQL server plugin as part of supporting prepared statements.
	Prepare(context.Context, *connect_go.Request[dev.PrepareRequest]) (*connect_go.Response[dev.PrepareResponse], error)
	// CloseSession closes the session, rolling back any implicit transactions.
	// This has the same effect as if a "rollback" statement was executed,
	// but does not affect the query statistics.
	CloseSession(context.Context, *connect_go.Request[dev.CloseSessionRequest]) (*connect_go.Response[dev.CloseSessionResponse], error)
}

// NewVitessClient constructs a client for the vitess.vtgateservice.dev.Vitess service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewVitessClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) VitessClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &vitessClient{
		execute: connect_go.NewClient[dev.ExecuteRequest, dev.ExecuteResponse](
			httpClient,
			baseURL+VitessExecuteProcedure,
			opts...,
		),
		executeBatch: connect_go.NewClient[dev.ExecuteBatchRequest, dev.ExecuteBatchResponse](
			httpClient,
			baseURL+VitessExecuteBatchProcedure,
			opts...,
		),
		streamExecute: connect_go.NewClient[dev.StreamExecuteRequest, dev.StreamExecuteResponse](
			httpClient,
			baseURL+VitessStreamExecuteProcedure,
			opts...,
		),
		resolveTransaction: connect_go.NewClient[dev.ResolveTransactionRequest, dev.ResolveTransactionResponse](
			httpClient,
			baseURL+VitessResolveTransactionProcedure,
			opts...,
		),
		vStream: connect_go.NewClient[dev.VStreamRequest, dev.VStreamResponse](
			httpClient,
			baseURL+VitessVStreamProcedure,
			opts...,
		),
		prepare: connect_go.NewClient[dev.PrepareRequest, dev.PrepareResponse](
			httpClient,
			baseURL+VitessPrepareProcedure,
			opts...,
		),
		closeSession: connect_go.NewClient[dev.CloseSessionRequest, dev.CloseSessionResponse](
			httpClient,
			baseURL+VitessCloseSessionProcedure,
			opts...,
		),
	}
}

// vitessClient implements VitessClient.
type vitessClient struct {
	execute            *connect_go.Client[dev.ExecuteRequest, dev.ExecuteResponse]
	executeBatch       *connect_go.Client[dev.ExecuteBatchRequest, dev.ExecuteBatchResponse]
	streamExecute      *connect_go.Client[dev.StreamExecuteRequest, dev.StreamExecuteResponse]
	resolveTransaction *connect_go.Client[dev.ResolveTransactionRequest, dev.ResolveTransactionResponse]
	vStream            *connect_go.Client[dev.VStreamRequest, dev.VStreamResponse]
	prepare            *connect_go.Client[dev.PrepareRequest, dev.PrepareResponse]
	closeSession       *connect_go.Client[dev.CloseSessionRequest, dev.CloseSessionResponse]
}

// Execute calls vitess.vtgateservice.dev.Vitess.Execute.
func (c *vitessClient) Execute(ctx context.Context, req *connect_go.Request[dev.ExecuteRequest]) (*connect_go.Response[dev.ExecuteResponse], error) {
	return c.execute.CallUnary(ctx, req)
}

// ExecuteBatch calls vitess.vtgateservice.dev.Vitess.ExecuteBatch.
func (c *vitessClient) ExecuteBatch(ctx context.Context, req *connect_go.Request[dev.ExecuteBatchRequest]) (*connect_go.Response[dev.ExecuteBatchResponse], error) {
	return c.executeBatch.CallUnary(ctx, req)
}

// StreamExecute calls vitess.vtgateservice.dev.Vitess.StreamExecute.
func (c *vitessClient) StreamExecute(ctx context.Context, req *connect_go.Request[dev.StreamExecuteRequest]) (*connect_go.ServerStreamForClient[dev.StreamExecuteResponse], error) {
	return c.streamExecute.CallServerStream(ctx, req)
}

// ResolveTransaction calls vitess.vtgateservice.dev.Vitess.ResolveTransaction.
func (c *vitessClient) ResolveTransaction(ctx context.Context, req *connect_go.Request[dev.ResolveTransactionRequest]) (*connect_go.Response[dev.ResolveTransactionResponse], error) {
	return c.resolveTransaction.CallUnary(ctx, req)
}

// VStream calls vitess.vtgateservice.dev.Vitess.VStream.
func (c *vitessClient) VStream(ctx context.Context, req *connect_go.Request[dev.VStreamRequest]) (*connect_go.ServerStreamForClient[dev.VStreamResponse], error) {
	return c.vStream.CallServerStream(ctx, req)
}

// Prepare calls vitess.vtgateservice.dev.Vitess.Prepare.
func (c *vitessClient) Prepare(ctx context.Context, req *connect_go.Request[dev.PrepareRequest]) (*connect_go.Response[dev.PrepareResponse], error) {
	return c.prepare.CallUnary(ctx, req)
}

// CloseSession calls vitess.vtgateservice.dev.Vitess.CloseSession.
func (c *vitessClient) CloseSession(ctx context.Context, req *connect_go.Request[dev.CloseSessionRequest]) (*connect_go.Response[dev.CloseSessionResponse], error) {
	return c.closeSession.CallUnary(ctx, req)
}

// VitessHandler is an implementation of the vitess.vtgateservice.dev.Vitess service.
type VitessHandler interface {
	// Execute tries to route the query to the right shard.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// API group: v3
	Execute(context.Context, *connect_go.Request[dev.ExecuteRequest]) (*connect_go.Response[dev.ExecuteResponse], error)
	// ExecuteBatch tries to route the list of queries on the right shards.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// API group: v3
	ExecuteBatch(context.Context, *connect_go.Request[dev.ExecuteBatchRequest]) (*connect_go.Response[dev.ExecuteBatchResponse], error)
	// StreamExecute executes a streaming query based on shards.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// Use this method if the query returns a large number of rows.
	// API group: v3
	StreamExecute(context.Context, *connect_go.Request[dev.StreamExecuteRequest], *connect_go.ServerStream[dev.StreamExecuteResponse]) error
	// ResolveTransaction resolves a transaction.
	// API group: Transactions
	ResolveTransaction(context.Context, *connect_go.Request[dev.ResolveTransactionRequest]) (*connect_go.Response[dev.ResolveTransactionResponse], error)
	// VStream streams binlog events from the requested sources.
	VStream(context.Context, *connect_go.Request[dev.VStreamRequest], *connect_go.ServerStream[dev.VStreamResponse]) error
	// Prepare is used by the MySQL server plugin as part of supporting prepared statements.
	Prepare(context.Context, *connect_go.Request[dev.PrepareRequest]) (*connect_go.Response[dev.PrepareResponse], error)
	// CloseSession closes the session, rolling back any implicit transactions.
	// This has the same effect as if a "rollback" statement was executed,
	// but does not affect the query statistics.
	CloseSession(context.Context, *connect_go.Request[dev.CloseSessionRequest]) (*connect_go.Response[dev.CloseSessionResponse], error)
}

// NewVitessHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewVitessHandler(svc VitessHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(VitessExecuteProcedure, connect_go.NewUnaryHandler(
		VitessExecuteProcedure,
		svc.Execute,
		opts...,
	))
	mux.Handle(VitessExecuteBatchProcedure, connect_go.NewUnaryHandler(
		VitessExecuteBatchProcedure,
		svc.ExecuteBatch,
		opts...,
	))
	mux.Handle(VitessStreamExecuteProcedure, connect_go.NewServerStreamHandler(
		VitessStreamExecuteProcedure,
		svc.StreamExecute,
		opts...,
	))
	mux.Handle(VitessResolveTransactionProcedure, connect_go.NewUnaryHandler(
		VitessResolveTransactionProcedure,
		svc.ResolveTransaction,
		opts...,
	))
	mux.Handle(VitessVStreamProcedure, connect_go.NewServerStreamHandler(
		VitessVStreamProcedure,
		svc.VStream,
		opts...,
	))
	mux.Handle(VitessPrepareProcedure, connect_go.NewUnaryHandler(
		VitessPrepareProcedure,
		svc.Prepare,
		opts...,
	))
	mux.Handle(VitessCloseSessionProcedure, connect_go.NewUnaryHandler(
		VitessCloseSessionProcedure,
		svc.CloseSession,
		opts...,
	))
	return "/vitess.vtgateservice.dev.Vitess/", mux
}

// UnimplementedVitessHandler returns CodeUnimplemented from all methods.
type UnimplementedVitessHandler struct{}

func (UnimplementedVitessHandler) Execute(context.Context, *connect_go.Request[dev.ExecuteRequest]) (*connect_go.Response[dev.ExecuteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.vtgateservice.dev.Vitess.Execute is not implemented"))
}

func (UnimplementedVitessHandler) ExecuteBatch(context.Context, *connect_go.Request[dev.ExecuteBatchRequest]) (*connect_go.Response[dev.ExecuteBatchResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.vtgateservice.dev.Vitess.ExecuteBatch is not implemented"))
}

func (UnimplementedVitessHandler) StreamExecute(context.Context, *connect_go.Request[dev.StreamExecuteRequest], *connect_go.ServerStream[dev.StreamExecuteResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.vtgateservice.dev.Vitess.StreamExecute is not implemented"))
}

func (UnimplementedVitessHandler) ResolveTransaction(context.Context, *connect_go.Request[dev.ResolveTransactionRequest]) (*connect_go.Response[dev.ResolveTransactionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.vtgateservice.dev.Vitess.ResolveTransaction is not implemented"))
}

func (UnimplementedVitessHandler) VStream(context.Context, *connect_go.Request[dev.VStreamRequest], *connect_go.ServerStream[dev.VStreamResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.vtgateservice.dev.Vitess.VStream is not implemented"))
}

func (UnimplementedVitessHandler) Prepare(context.Context, *connect_go.Request[dev.PrepareRequest]) (*connect_go.Response[dev.PrepareResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.vtgateservice.dev.Vitess.Prepare is not implemented"))
}

func (UnimplementedVitessHandler) CloseSession(context.Context, *connect_go.Request[dev.CloseSessionRequest]) (*connect_go.Response[dev.CloseSessionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.vtgateservice.dev.Vitess.CloseSession is not implemented"))
}