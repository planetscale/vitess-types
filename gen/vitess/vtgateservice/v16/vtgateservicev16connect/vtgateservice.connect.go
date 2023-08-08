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
// Source: vitess/vtgateservice/v16/vtgateservice.proto

// option java_package="io.vitess.proto.grpc";

package vtgateservicev16connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v16 "github.com/planetscale/vitess-types/gen/vitess/vtgate/v16"
	_ "github.com/planetscale/vitess-types/gen/vitess/vtgateservice/v16"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// VitessName is the fully-qualified name of the Vitess service.
	VitessName = "vtgateservice.Vitess"
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
	VitessExecuteProcedure = "/vtgateservice.Vitess/Execute"
	// VitessExecuteBatchProcedure is the fully-qualified name of the Vitess's ExecuteBatch RPC.
	VitessExecuteBatchProcedure = "/vtgateservice.Vitess/ExecuteBatch"
	// VitessStreamExecuteProcedure is the fully-qualified name of the Vitess's StreamExecute RPC.
	VitessStreamExecuteProcedure = "/vtgateservice.Vitess/StreamExecute"
	// VitessResolveTransactionProcedure is the fully-qualified name of the Vitess's ResolveTransaction
	// RPC.
	VitessResolveTransactionProcedure = "/vtgateservice.Vitess/ResolveTransaction"
	// VitessVStreamProcedure is the fully-qualified name of the Vitess's VStream RPC.
	VitessVStreamProcedure = "/vtgateservice.Vitess/VStream"
	// VitessPrepareProcedure is the fully-qualified name of the Vitess's Prepare RPC.
	VitessPrepareProcedure = "/vtgateservice.Vitess/Prepare"
	// VitessCloseSessionProcedure is the fully-qualified name of the Vitess's CloseSession RPC.
	VitessCloseSessionProcedure = "/vtgateservice.Vitess/CloseSession"
)

// VitessClient is a client for the vtgateservice.Vitess service.
type VitessClient interface {
	// Execute tries to route the query to the right shard.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// API group: v3
	Execute(context.Context, *connect.Request[v16.ExecuteRequest]) (*connect.Response[v16.ExecuteResponse], error)
	// ExecuteBatch tries to route the list of queries on the right shards.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// API group: v3
	ExecuteBatch(context.Context, *connect.Request[v16.ExecuteBatchRequest]) (*connect.Response[v16.ExecuteBatchResponse], error)
	// StreamExecute executes a streaming query based on shards.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// Use this method if the query returns a large number of rows.
	// API group: v3
	StreamExecute(context.Context, *connect.Request[v16.StreamExecuteRequest]) (*connect.ServerStreamForClient[v16.StreamExecuteResponse], error)
	// ResolveTransaction resolves a transaction.
	// API group: Transactions
	ResolveTransaction(context.Context, *connect.Request[v16.ResolveTransactionRequest]) (*connect.Response[v16.ResolveTransactionResponse], error)
	// VStream streams binlog events from the requested sources.
	VStream(context.Context, *connect.Request[v16.VStreamRequest]) (*connect.ServerStreamForClient[v16.VStreamResponse], error)
	// Prepare is used by the MySQL server plugin as part of supporting prepared statements.
	Prepare(context.Context, *connect.Request[v16.PrepareRequest]) (*connect.Response[v16.PrepareResponse], error)
	// CloseSession closes the session, rolling back any implicit transactions.
	// This has the same effect as if a "rollback" statement was executed,
	// but does not affect the query statistics.
	CloseSession(context.Context, *connect.Request[v16.CloseSessionRequest]) (*connect.Response[v16.CloseSessionResponse], error)
}

// NewVitessClient constructs a client for the vtgateservice.Vitess service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewVitessClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) VitessClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &vitessClient{
		execute: connect.NewClient[v16.ExecuteRequest, v16.ExecuteResponse](
			httpClient,
			baseURL+VitessExecuteProcedure,
			opts...,
		),
		executeBatch: connect.NewClient[v16.ExecuteBatchRequest, v16.ExecuteBatchResponse](
			httpClient,
			baseURL+VitessExecuteBatchProcedure,
			opts...,
		),
		streamExecute: connect.NewClient[v16.StreamExecuteRequest, v16.StreamExecuteResponse](
			httpClient,
			baseURL+VitessStreamExecuteProcedure,
			opts...,
		),
		resolveTransaction: connect.NewClient[v16.ResolveTransactionRequest, v16.ResolveTransactionResponse](
			httpClient,
			baseURL+VitessResolveTransactionProcedure,
			opts...,
		),
		vStream: connect.NewClient[v16.VStreamRequest, v16.VStreamResponse](
			httpClient,
			baseURL+VitessVStreamProcedure,
			opts...,
		),
		prepare: connect.NewClient[v16.PrepareRequest, v16.PrepareResponse](
			httpClient,
			baseURL+VitessPrepareProcedure,
			opts...,
		),
		closeSession: connect.NewClient[v16.CloseSessionRequest, v16.CloseSessionResponse](
			httpClient,
			baseURL+VitessCloseSessionProcedure,
			opts...,
		),
	}
}

// vitessClient implements VitessClient.
type vitessClient struct {
	execute            *connect.Client[v16.ExecuteRequest, v16.ExecuteResponse]
	executeBatch       *connect.Client[v16.ExecuteBatchRequest, v16.ExecuteBatchResponse]
	streamExecute      *connect.Client[v16.StreamExecuteRequest, v16.StreamExecuteResponse]
	resolveTransaction *connect.Client[v16.ResolveTransactionRequest, v16.ResolveTransactionResponse]
	vStream            *connect.Client[v16.VStreamRequest, v16.VStreamResponse]
	prepare            *connect.Client[v16.PrepareRequest, v16.PrepareResponse]
	closeSession       *connect.Client[v16.CloseSessionRequest, v16.CloseSessionResponse]
}

// Execute calls vtgateservice.Vitess.Execute.
func (c *vitessClient) Execute(ctx context.Context, req *connect.Request[v16.ExecuteRequest]) (*connect.Response[v16.ExecuteResponse], error) {
	return c.execute.CallUnary(ctx, req)
}

// ExecuteBatch calls vtgateservice.Vitess.ExecuteBatch.
func (c *vitessClient) ExecuteBatch(ctx context.Context, req *connect.Request[v16.ExecuteBatchRequest]) (*connect.Response[v16.ExecuteBatchResponse], error) {
	return c.executeBatch.CallUnary(ctx, req)
}

// StreamExecute calls vtgateservice.Vitess.StreamExecute.
func (c *vitessClient) StreamExecute(ctx context.Context, req *connect.Request[v16.StreamExecuteRequest]) (*connect.ServerStreamForClient[v16.StreamExecuteResponse], error) {
	return c.streamExecute.CallServerStream(ctx, req)
}

// ResolveTransaction calls vtgateservice.Vitess.ResolveTransaction.
func (c *vitessClient) ResolveTransaction(ctx context.Context, req *connect.Request[v16.ResolveTransactionRequest]) (*connect.Response[v16.ResolveTransactionResponse], error) {
	return c.resolveTransaction.CallUnary(ctx, req)
}

// VStream calls vtgateservice.Vitess.VStream.
func (c *vitessClient) VStream(ctx context.Context, req *connect.Request[v16.VStreamRequest]) (*connect.ServerStreamForClient[v16.VStreamResponse], error) {
	return c.vStream.CallServerStream(ctx, req)
}

// Prepare calls vtgateservice.Vitess.Prepare.
func (c *vitessClient) Prepare(ctx context.Context, req *connect.Request[v16.PrepareRequest]) (*connect.Response[v16.PrepareResponse], error) {
	return c.prepare.CallUnary(ctx, req)
}

// CloseSession calls vtgateservice.Vitess.CloseSession.
func (c *vitessClient) CloseSession(ctx context.Context, req *connect.Request[v16.CloseSessionRequest]) (*connect.Response[v16.CloseSessionResponse], error) {
	return c.closeSession.CallUnary(ctx, req)
}

// VitessHandler is an implementation of the vtgateservice.Vitess service.
type VitessHandler interface {
	// Execute tries to route the query to the right shard.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// API group: v3
	Execute(context.Context, *connect.Request[v16.ExecuteRequest]) (*connect.Response[v16.ExecuteResponse], error)
	// ExecuteBatch tries to route the list of queries on the right shards.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// API group: v3
	ExecuteBatch(context.Context, *connect.Request[v16.ExecuteBatchRequest]) (*connect.Response[v16.ExecuteBatchResponse], error)
	// StreamExecute executes a streaming query based on shards.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// Use this method if the query returns a large number of rows.
	// API group: v3
	StreamExecute(context.Context, *connect.Request[v16.StreamExecuteRequest], *connect.ServerStream[v16.StreamExecuteResponse]) error
	// ResolveTransaction resolves a transaction.
	// API group: Transactions
	ResolveTransaction(context.Context, *connect.Request[v16.ResolveTransactionRequest]) (*connect.Response[v16.ResolveTransactionResponse], error)
	// VStream streams binlog events from the requested sources.
	VStream(context.Context, *connect.Request[v16.VStreamRequest], *connect.ServerStream[v16.VStreamResponse]) error
	// Prepare is used by the MySQL server plugin as part of supporting prepared statements.
	Prepare(context.Context, *connect.Request[v16.PrepareRequest]) (*connect.Response[v16.PrepareResponse], error)
	// CloseSession closes the session, rolling back any implicit transactions.
	// This has the same effect as if a "rollback" statement was executed,
	// but does not affect the query statistics.
	CloseSession(context.Context, *connect.Request[v16.CloseSessionRequest]) (*connect.Response[v16.CloseSessionResponse], error)
}

// NewVitessHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewVitessHandler(svc VitessHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	vitessExecuteHandler := connect.NewUnaryHandler(
		VitessExecuteProcedure,
		svc.Execute,
		opts...,
	)
	vitessExecuteBatchHandler := connect.NewUnaryHandler(
		VitessExecuteBatchProcedure,
		svc.ExecuteBatch,
		opts...,
	)
	vitessStreamExecuteHandler := connect.NewServerStreamHandler(
		VitessStreamExecuteProcedure,
		svc.StreamExecute,
		opts...,
	)
	vitessResolveTransactionHandler := connect.NewUnaryHandler(
		VitessResolveTransactionProcedure,
		svc.ResolveTransaction,
		opts...,
	)
	vitessVStreamHandler := connect.NewServerStreamHandler(
		VitessVStreamProcedure,
		svc.VStream,
		opts...,
	)
	vitessPrepareHandler := connect.NewUnaryHandler(
		VitessPrepareProcedure,
		svc.Prepare,
		opts...,
	)
	vitessCloseSessionHandler := connect.NewUnaryHandler(
		VitessCloseSessionProcedure,
		svc.CloseSession,
		opts...,
	)
	return "/vtgateservice.Vitess/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case VitessExecuteProcedure:
			vitessExecuteHandler.ServeHTTP(w, r)
		case VitessExecuteBatchProcedure:
			vitessExecuteBatchHandler.ServeHTTP(w, r)
		case VitessStreamExecuteProcedure:
			vitessStreamExecuteHandler.ServeHTTP(w, r)
		case VitessResolveTransactionProcedure:
			vitessResolveTransactionHandler.ServeHTTP(w, r)
		case VitessVStreamProcedure:
			vitessVStreamHandler.ServeHTTP(w, r)
		case VitessPrepareProcedure:
			vitessPrepareHandler.ServeHTTP(w, r)
		case VitessCloseSessionProcedure:
			vitessCloseSessionHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedVitessHandler returns CodeUnimplemented from all methods.
type UnimplementedVitessHandler struct{}

func (UnimplementedVitessHandler) Execute(context.Context, *connect.Request[v16.ExecuteRequest]) (*connect.Response[v16.ExecuteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("vtgateservice.Vitess.Execute is not implemented"))
}

func (UnimplementedVitessHandler) ExecuteBatch(context.Context, *connect.Request[v16.ExecuteBatchRequest]) (*connect.Response[v16.ExecuteBatchResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("vtgateservice.Vitess.ExecuteBatch is not implemented"))
}

func (UnimplementedVitessHandler) StreamExecute(context.Context, *connect.Request[v16.StreamExecuteRequest], *connect.ServerStream[v16.StreamExecuteResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("vtgateservice.Vitess.StreamExecute is not implemented"))
}

func (UnimplementedVitessHandler) ResolveTransaction(context.Context, *connect.Request[v16.ResolveTransactionRequest]) (*connect.Response[v16.ResolveTransactionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("vtgateservice.Vitess.ResolveTransaction is not implemented"))
}

func (UnimplementedVitessHandler) VStream(context.Context, *connect.Request[v16.VStreamRequest], *connect.ServerStream[v16.VStreamResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("vtgateservice.Vitess.VStream is not implemented"))
}

func (UnimplementedVitessHandler) Prepare(context.Context, *connect.Request[v16.PrepareRequest]) (*connect.Response[v16.PrepareResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("vtgateservice.Vitess.Prepare is not implemented"))
}

func (UnimplementedVitessHandler) CloseSession(context.Context, *connect.Request[v16.CloseSessionRequest]) (*connect.Response[v16.CloseSessionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("vtgateservice.Vitess.CloseSession is not implemented"))
}
