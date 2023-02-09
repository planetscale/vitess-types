// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: vtgateservice/vtgateservice.proto

package vtgateserviceconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	vtgate "github.com/planetscale/vitess-types/gen/vitess/dev/vtgate"
	_ "github.com/planetscale/vitess-types/gen/vitess/dev/vtgateservice"
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
	VitessName = "vtgateservice.Vitess"
)

// VitessClient is a client for the vtgateservice.Vitess service.
type VitessClient interface {
	// Execute tries to route the query to the right shard.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// API group: v3
	Execute(context.Context, *connect_go.Request[vtgate.ExecuteRequest]) (*connect_go.Response[vtgate.ExecuteResponse], error)
	// ExecuteBatch tries to route the list of queries on the right shards.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// API group: v3
	ExecuteBatch(context.Context, *connect_go.Request[vtgate.ExecuteBatchRequest]) (*connect_go.Response[vtgate.ExecuteBatchResponse], error)
	// StreamExecute executes a streaming query based on shards.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// Use this method if the query returns a large number of rows.
	// API group: v3
	StreamExecute(context.Context, *connect_go.Request[vtgate.StreamExecuteRequest]) (*connect_go.ServerStreamForClient[vtgate.StreamExecuteResponse], error)
	// ResolveTransaction resolves a transaction.
	// API group: Transactions
	ResolveTransaction(context.Context, *connect_go.Request[vtgate.ResolveTransactionRequest]) (*connect_go.Response[vtgate.ResolveTransactionResponse], error)
	// VStream streams binlog events from the requested sources.
	VStream(context.Context, *connect_go.Request[vtgate.VStreamRequest]) (*connect_go.ServerStreamForClient[vtgate.VStreamResponse], error)
	// Prepare is used by the MySQL server plugin as part of supporting prepared statements.
	Prepare(context.Context, *connect_go.Request[vtgate.PrepareRequest]) (*connect_go.Response[vtgate.PrepareResponse], error)
	// CloseSession closes the session, rolling back any implicit transactions.
	// This has the same effect as if a "rollback" statement was executed,
	// but does not affect the query statistics.
	CloseSession(context.Context, *connect_go.Request[vtgate.CloseSessionRequest]) (*connect_go.Response[vtgate.CloseSessionResponse], error)
}

// NewVitessClient constructs a client for the vtgateservice.Vitess service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewVitessClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) VitessClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &vitessClient{
		execute: connect_go.NewClient[vtgate.ExecuteRequest, vtgate.ExecuteResponse](
			httpClient,
			baseURL+"/vtgateservice.Vitess/Execute",
			opts...,
		),
		executeBatch: connect_go.NewClient[vtgate.ExecuteBatchRequest, vtgate.ExecuteBatchResponse](
			httpClient,
			baseURL+"/vtgateservice.Vitess/ExecuteBatch",
			opts...,
		),
		streamExecute: connect_go.NewClient[vtgate.StreamExecuteRequest, vtgate.StreamExecuteResponse](
			httpClient,
			baseURL+"/vtgateservice.Vitess/StreamExecute",
			opts...,
		),
		resolveTransaction: connect_go.NewClient[vtgate.ResolveTransactionRequest, vtgate.ResolveTransactionResponse](
			httpClient,
			baseURL+"/vtgateservice.Vitess/ResolveTransaction",
			opts...,
		),
		vStream: connect_go.NewClient[vtgate.VStreamRequest, vtgate.VStreamResponse](
			httpClient,
			baseURL+"/vtgateservice.Vitess/VStream",
			opts...,
		),
		prepare: connect_go.NewClient[vtgate.PrepareRequest, vtgate.PrepareResponse](
			httpClient,
			baseURL+"/vtgateservice.Vitess/Prepare",
			opts...,
		),
		closeSession: connect_go.NewClient[vtgate.CloseSessionRequest, vtgate.CloseSessionResponse](
			httpClient,
			baseURL+"/vtgateservice.Vitess/CloseSession",
			opts...,
		),
	}
}

// vitessClient implements VitessClient.
type vitessClient struct {
	execute            *connect_go.Client[vtgate.ExecuteRequest, vtgate.ExecuteResponse]
	executeBatch       *connect_go.Client[vtgate.ExecuteBatchRequest, vtgate.ExecuteBatchResponse]
	streamExecute      *connect_go.Client[vtgate.StreamExecuteRequest, vtgate.StreamExecuteResponse]
	resolveTransaction *connect_go.Client[vtgate.ResolveTransactionRequest, vtgate.ResolveTransactionResponse]
	vStream            *connect_go.Client[vtgate.VStreamRequest, vtgate.VStreamResponse]
	prepare            *connect_go.Client[vtgate.PrepareRequest, vtgate.PrepareResponse]
	closeSession       *connect_go.Client[vtgate.CloseSessionRequest, vtgate.CloseSessionResponse]
}

// Execute calls vtgateservice.Vitess.Execute.
func (c *vitessClient) Execute(ctx context.Context, req *connect_go.Request[vtgate.ExecuteRequest]) (*connect_go.Response[vtgate.ExecuteResponse], error) {
	return c.execute.CallUnary(ctx, req)
}

// ExecuteBatch calls vtgateservice.Vitess.ExecuteBatch.
func (c *vitessClient) ExecuteBatch(ctx context.Context, req *connect_go.Request[vtgate.ExecuteBatchRequest]) (*connect_go.Response[vtgate.ExecuteBatchResponse], error) {
	return c.executeBatch.CallUnary(ctx, req)
}

// StreamExecute calls vtgateservice.Vitess.StreamExecute.
func (c *vitessClient) StreamExecute(ctx context.Context, req *connect_go.Request[vtgate.StreamExecuteRequest]) (*connect_go.ServerStreamForClient[vtgate.StreamExecuteResponse], error) {
	return c.streamExecute.CallServerStream(ctx, req)
}

// ResolveTransaction calls vtgateservice.Vitess.ResolveTransaction.
func (c *vitessClient) ResolveTransaction(ctx context.Context, req *connect_go.Request[vtgate.ResolveTransactionRequest]) (*connect_go.Response[vtgate.ResolveTransactionResponse], error) {
	return c.resolveTransaction.CallUnary(ctx, req)
}

// VStream calls vtgateservice.Vitess.VStream.
func (c *vitessClient) VStream(ctx context.Context, req *connect_go.Request[vtgate.VStreamRequest]) (*connect_go.ServerStreamForClient[vtgate.VStreamResponse], error) {
	return c.vStream.CallServerStream(ctx, req)
}

// Prepare calls vtgateservice.Vitess.Prepare.
func (c *vitessClient) Prepare(ctx context.Context, req *connect_go.Request[vtgate.PrepareRequest]) (*connect_go.Response[vtgate.PrepareResponse], error) {
	return c.prepare.CallUnary(ctx, req)
}

// CloseSession calls vtgateservice.Vitess.CloseSession.
func (c *vitessClient) CloseSession(ctx context.Context, req *connect_go.Request[vtgate.CloseSessionRequest]) (*connect_go.Response[vtgate.CloseSessionResponse], error) {
	return c.closeSession.CallUnary(ctx, req)
}

// VitessHandler is an implementation of the vtgateservice.Vitess service.
type VitessHandler interface {
	// Execute tries to route the query to the right shard.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// API group: v3
	Execute(context.Context, *connect_go.Request[vtgate.ExecuteRequest]) (*connect_go.Response[vtgate.ExecuteResponse], error)
	// ExecuteBatch tries to route the list of queries on the right shards.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// API group: v3
	ExecuteBatch(context.Context, *connect_go.Request[vtgate.ExecuteBatchRequest]) (*connect_go.Response[vtgate.ExecuteBatchResponse], error)
	// StreamExecute executes a streaming query based on shards.
	// It depends on the query and bind variables to provide enough
	// information in conjunction with the vindexes to route the query.
	// Use this method if the query returns a large number of rows.
	// API group: v3
	StreamExecute(context.Context, *connect_go.Request[vtgate.StreamExecuteRequest], *connect_go.ServerStream[vtgate.StreamExecuteResponse]) error
	// ResolveTransaction resolves a transaction.
	// API group: Transactions
	ResolveTransaction(context.Context, *connect_go.Request[vtgate.ResolveTransactionRequest]) (*connect_go.Response[vtgate.ResolveTransactionResponse], error)
	// VStream streams binlog events from the requested sources.
	VStream(context.Context, *connect_go.Request[vtgate.VStreamRequest], *connect_go.ServerStream[vtgate.VStreamResponse]) error
	// Prepare is used by the MySQL server plugin as part of supporting prepared statements.
	Prepare(context.Context, *connect_go.Request[vtgate.PrepareRequest]) (*connect_go.Response[vtgate.PrepareResponse], error)
	// CloseSession closes the session, rolling back any implicit transactions.
	// This has the same effect as if a "rollback" statement was executed,
	// but does not affect the query statistics.
	CloseSession(context.Context, *connect_go.Request[vtgate.CloseSessionRequest]) (*connect_go.Response[vtgate.CloseSessionResponse], error)
}

// NewVitessHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewVitessHandler(svc VitessHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/vtgateservice.Vitess/Execute", connect_go.NewUnaryHandler(
		"/vtgateservice.Vitess/Execute",
		svc.Execute,
		opts...,
	))
	mux.Handle("/vtgateservice.Vitess/ExecuteBatch", connect_go.NewUnaryHandler(
		"/vtgateservice.Vitess/ExecuteBatch",
		svc.ExecuteBatch,
		opts...,
	))
	mux.Handle("/vtgateservice.Vitess/StreamExecute", connect_go.NewServerStreamHandler(
		"/vtgateservice.Vitess/StreamExecute",
		svc.StreamExecute,
		opts...,
	))
	mux.Handle("/vtgateservice.Vitess/ResolveTransaction", connect_go.NewUnaryHandler(
		"/vtgateservice.Vitess/ResolveTransaction",
		svc.ResolveTransaction,
		opts...,
	))
	mux.Handle("/vtgateservice.Vitess/VStream", connect_go.NewServerStreamHandler(
		"/vtgateservice.Vitess/VStream",
		svc.VStream,
		opts...,
	))
	mux.Handle("/vtgateservice.Vitess/Prepare", connect_go.NewUnaryHandler(
		"/vtgateservice.Vitess/Prepare",
		svc.Prepare,
		opts...,
	))
	mux.Handle("/vtgateservice.Vitess/CloseSession", connect_go.NewUnaryHandler(
		"/vtgateservice.Vitess/CloseSession",
		svc.CloseSession,
		opts...,
	))
	return "/vtgateservice.Vitess/", mux
}

// UnimplementedVitessHandler returns CodeUnimplemented from all methods.
type UnimplementedVitessHandler struct{}

func (UnimplementedVitessHandler) Execute(context.Context, *connect_go.Request[vtgate.ExecuteRequest]) (*connect_go.Response[vtgate.ExecuteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vtgateservice.Vitess.Execute is not implemented"))
}

func (UnimplementedVitessHandler) ExecuteBatch(context.Context, *connect_go.Request[vtgate.ExecuteBatchRequest]) (*connect_go.Response[vtgate.ExecuteBatchResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vtgateservice.Vitess.ExecuteBatch is not implemented"))
}

func (UnimplementedVitessHandler) StreamExecute(context.Context, *connect_go.Request[vtgate.StreamExecuteRequest], *connect_go.ServerStream[vtgate.StreamExecuteResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vtgateservice.Vitess.StreamExecute is not implemented"))
}

func (UnimplementedVitessHandler) ResolveTransaction(context.Context, *connect_go.Request[vtgate.ResolveTransactionRequest]) (*connect_go.Response[vtgate.ResolveTransactionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vtgateservice.Vitess.ResolveTransaction is not implemented"))
}

func (UnimplementedVitessHandler) VStream(context.Context, *connect_go.Request[vtgate.VStreamRequest], *connect_go.ServerStream[vtgate.VStreamResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vtgateservice.Vitess.VStream is not implemented"))
}

func (UnimplementedVitessHandler) Prepare(context.Context, *connect_go.Request[vtgate.PrepareRequest]) (*connect_go.Response[vtgate.PrepareResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vtgateservice.Vitess.Prepare is not implemented"))
}

func (UnimplementedVitessHandler) CloseSession(context.Context, *connect_go.Request[vtgate.CloseSessionRequest]) (*connect_go.Response[vtgate.CloseSessionResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vtgateservice.Vitess.CloseSession is not implemented"))
}
