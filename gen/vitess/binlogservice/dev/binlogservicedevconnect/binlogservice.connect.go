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
// This file contains the UpdateStream service definition, necessary
// to make RPC calls to VtTablet for the binlog protocol, used by
// filtered replication only.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: vitess/binlogservice/dev/binlogservice.proto

package binlogservicedevconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	dev "github.com/planetscale/vitess-types/gen/vitess/binlogdata/dev"
	_ "github.com/planetscale/vitess-types/gen/vitess/binlogservice/dev"
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
	// UpdateStreamName is the fully-qualified name of the UpdateStream service.
	UpdateStreamName = "binlogservice.UpdateStream"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UpdateStreamStreamKeyRangeProcedure is the fully-qualified name of the UpdateStream's
	// StreamKeyRange RPC.
	UpdateStreamStreamKeyRangeProcedure = "/binlogservice.UpdateStream/StreamKeyRange"
	// UpdateStreamStreamTablesProcedure is the fully-qualified name of the UpdateStream's StreamTables
	// RPC.
	UpdateStreamStreamTablesProcedure = "/binlogservice.UpdateStream/StreamTables"
)

// UpdateStreamClient is a client for the binlogservice.UpdateStream service.
type UpdateStreamClient interface {
	// StreamKeyRange returns the binlog transactions related to
	// the specified Keyrange.
	StreamKeyRange(context.Context, *connect.Request[dev.StreamKeyRangeRequest]) (*connect.ServerStreamForClient[dev.StreamKeyRangeResponse], error)
	// StreamTables returns the binlog transactions related to
	// the specified Tables.
	StreamTables(context.Context, *connect.Request[dev.StreamTablesRequest]) (*connect.ServerStreamForClient[dev.StreamTablesResponse], error)
}

// NewUpdateStreamClient constructs a client for the binlogservice.UpdateStream service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUpdateStreamClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UpdateStreamClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &updateStreamClient{
		streamKeyRange: connect.NewClient[dev.StreamKeyRangeRequest, dev.StreamKeyRangeResponse](
			httpClient,
			baseURL+UpdateStreamStreamKeyRangeProcedure,
			opts...,
		),
		streamTables: connect.NewClient[dev.StreamTablesRequest, dev.StreamTablesResponse](
			httpClient,
			baseURL+UpdateStreamStreamTablesProcedure,
			opts...,
		),
	}
}

// updateStreamClient implements UpdateStreamClient.
type updateStreamClient struct {
	streamKeyRange *connect.Client[dev.StreamKeyRangeRequest, dev.StreamKeyRangeResponse]
	streamTables   *connect.Client[dev.StreamTablesRequest, dev.StreamTablesResponse]
}

// StreamKeyRange calls binlogservice.UpdateStream.StreamKeyRange.
func (c *updateStreamClient) StreamKeyRange(ctx context.Context, req *connect.Request[dev.StreamKeyRangeRequest]) (*connect.ServerStreamForClient[dev.StreamKeyRangeResponse], error) {
	return c.streamKeyRange.CallServerStream(ctx, req)
}

// StreamTables calls binlogservice.UpdateStream.StreamTables.
func (c *updateStreamClient) StreamTables(ctx context.Context, req *connect.Request[dev.StreamTablesRequest]) (*connect.ServerStreamForClient[dev.StreamTablesResponse], error) {
	return c.streamTables.CallServerStream(ctx, req)
}

// UpdateStreamHandler is an implementation of the binlogservice.UpdateStream service.
type UpdateStreamHandler interface {
	// StreamKeyRange returns the binlog transactions related to
	// the specified Keyrange.
	StreamKeyRange(context.Context, *connect.Request[dev.StreamKeyRangeRequest], *connect.ServerStream[dev.StreamKeyRangeResponse]) error
	// StreamTables returns the binlog transactions related to
	// the specified Tables.
	StreamTables(context.Context, *connect.Request[dev.StreamTablesRequest], *connect.ServerStream[dev.StreamTablesResponse]) error
}

// NewUpdateStreamHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUpdateStreamHandler(svc UpdateStreamHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	updateStreamStreamKeyRangeHandler := connect.NewServerStreamHandler(
		UpdateStreamStreamKeyRangeProcedure,
		svc.StreamKeyRange,
		opts...,
	)
	updateStreamStreamTablesHandler := connect.NewServerStreamHandler(
		UpdateStreamStreamTablesProcedure,
		svc.StreamTables,
		opts...,
	)
	return "/binlogservice.UpdateStream/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UpdateStreamStreamKeyRangeProcedure:
			updateStreamStreamKeyRangeHandler.ServeHTTP(w, r)
		case UpdateStreamStreamTablesProcedure:
			updateStreamStreamTablesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUpdateStreamHandler returns CodeUnimplemented from all methods.
type UnimplementedUpdateStreamHandler struct{}

func (UnimplementedUpdateStreamHandler) StreamKeyRange(context.Context, *connect.Request[dev.StreamKeyRangeRequest], *connect.ServerStream[dev.StreamKeyRangeResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("binlogservice.UpdateStream.StreamKeyRange is not implemented"))
}

func (UnimplementedUpdateStreamHandler) StreamTables(context.Context, *connect.Request[dev.StreamTablesRequest], *connect.ServerStream[dev.StreamTablesResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("binlogservice.UpdateStream.StreamTables is not implemented"))
}
