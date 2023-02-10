// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: binlogservice/binlogservice.proto

package binlogserviceconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	binlogdata "github.com/planetscale/vitess-types/gen/vitess/v15/binlogdata"
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
	// UpdateStreamName is the fully-qualified name of the UpdateStream service.
	UpdateStreamName = "binlogservice.UpdateStream"
)

// UpdateStreamClient is a client for the binlogservice.UpdateStream service.
type UpdateStreamClient interface {
	// StreamKeyRange returns the binlog transactions related to
	// the specified Keyrange.
	StreamKeyRange(context.Context, *connect_go.Request[binlogdata.StreamKeyRangeRequest]) (*connect_go.ServerStreamForClient[binlogdata.StreamKeyRangeResponse], error)
	// StreamTables returns the binlog transactions related to
	// the specified Tables.
	StreamTables(context.Context, *connect_go.Request[binlogdata.StreamTablesRequest]) (*connect_go.ServerStreamForClient[binlogdata.StreamTablesResponse], error)
}

// NewUpdateStreamClient constructs a client for the binlogservice.UpdateStream service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUpdateStreamClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UpdateStreamClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &updateStreamClient{
		streamKeyRange: connect_go.NewClient[binlogdata.StreamKeyRangeRequest, binlogdata.StreamKeyRangeResponse](
			httpClient,
			baseURL+"/binlogservice.UpdateStream/StreamKeyRange",
			opts...,
		),
		streamTables: connect_go.NewClient[binlogdata.StreamTablesRequest, binlogdata.StreamTablesResponse](
			httpClient,
			baseURL+"/binlogservice.UpdateStream/StreamTables",
			opts...,
		),
	}
}

// updateStreamClient implements UpdateStreamClient.
type updateStreamClient struct {
	streamKeyRange *connect_go.Client[binlogdata.StreamKeyRangeRequest, binlogdata.StreamKeyRangeResponse]
	streamTables   *connect_go.Client[binlogdata.StreamTablesRequest, binlogdata.StreamTablesResponse]
}

// StreamKeyRange calls binlogservice.UpdateStream.StreamKeyRange.
func (c *updateStreamClient) StreamKeyRange(ctx context.Context, req *connect_go.Request[binlogdata.StreamKeyRangeRequest]) (*connect_go.ServerStreamForClient[binlogdata.StreamKeyRangeResponse], error) {
	return c.streamKeyRange.CallServerStream(ctx, req)
}

// StreamTables calls binlogservice.UpdateStream.StreamTables.
func (c *updateStreamClient) StreamTables(ctx context.Context, req *connect_go.Request[binlogdata.StreamTablesRequest]) (*connect_go.ServerStreamForClient[binlogdata.StreamTablesResponse], error) {
	return c.streamTables.CallServerStream(ctx, req)
}

// UpdateStreamHandler is an implementation of the binlogservice.UpdateStream service.
type UpdateStreamHandler interface {
	// StreamKeyRange returns the binlog transactions related to
	// the specified Keyrange.
	StreamKeyRange(context.Context, *connect_go.Request[binlogdata.StreamKeyRangeRequest], *connect_go.ServerStream[binlogdata.StreamKeyRangeResponse]) error
	// StreamTables returns the binlog transactions related to
	// the specified Tables.
	StreamTables(context.Context, *connect_go.Request[binlogdata.StreamTablesRequest], *connect_go.ServerStream[binlogdata.StreamTablesResponse]) error
}

// NewUpdateStreamHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUpdateStreamHandler(svc UpdateStreamHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/binlogservice.UpdateStream/StreamKeyRange", connect_go.NewServerStreamHandler(
		"/binlogservice.UpdateStream/StreamKeyRange",
		svc.StreamKeyRange,
		opts...,
	))
	mux.Handle("/binlogservice.UpdateStream/StreamTables", connect_go.NewServerStreamHandler(
		"/binlogservice.UpdateStream/StreamTables",
		svc.StreamTables,
		opts...,
	))
	return "/binlogservice.UpdateStream/", mux
}

// UnimplementedUpdateStreamHandler returns CodeUnimplemented from all methods.
type UnimplementedUpdateStreamHandler struct{}

func (UnimplementedUpdateStreamHandler) StreamKeyRange(context.Context, *connect_go.Request[binlogdata.StreamKeyRangeRequest], *connect_go.ServerStream[binlogdata.StreamKeyRangeResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("binlogservice.UpdateStream.StreamKeyRange is not implemented"))
}

func (UnimplementedUpdateStreamHandler) StreamTables(context.Context, *connect_go.Request[binlogdata.StreamTablesRequest], *connect_go.ServerStream[binlogdata.StreamTablesResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("binlogservice.UpdateStream.StreamTables is not implemented"))
}