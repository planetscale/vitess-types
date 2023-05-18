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
// calls to mysqlctld.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: vitess/mysqlctl/v15/mysqlctl.proto

package mysqlctlv15connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v15 "github.com/planetscale/vitess-types/gen/vitess/mysqlctl/v15"
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
	// MysqlCtlName is the fully-qualified name of the MysqlCtl service.
	MysqlCtlName = "vitess.mysqlctl.v15.MysqlCtl"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MysqlCtlStartProcedure is the fully-qualified name of the MysqlCtl's Start RPC.
	MysqlCtlStartProcedure = "/vitess.mysqlctl.v15.MysqlCtl/Start"
	// MysqlCtlShutdownProcedure is the fully-qualified name of the MysqlCtl's Shutdown RPC.
	MysqlCtlShutdownProcedure = "/vitess.mysqlctl.v15.MysqlCtl/Shutdown"
	// MysqlCtlRunMysqlUpgradeProcedure is the fully-qualified name of the MysqlCtl's RunMysqlUpgrade
	// RPC.
	MysqlCtlRunMysqlUpgradeProcedure = "/vitess.mysqlctl.v15.MysqlCtl/RunMysqlUpgrade"
	// MysqlCtlReinitConfigProcedure is the fully-qualified name of the MysqlCtl's ReinitConfig RPC.
	MysqlCtlReinitConfigProcedure = "/vitess.mysqlctl.v15.MysqlCtl/ReinitConfig"
	// MysqlCtlRefreshConfigProcedure is the fully-qualified name of the MysqlCtl's RefreshConfig RPC.
	MysqlCtlRefreshConfigProcedure = "/vitess.mysqlctl.v15.MysqlCtl/RefreshConfig"
)

// MysqlCtlClient is a client for the vitess.mysqlctl.v15.MysqlCtl service.
type MysqlCtlClient interface {
	Start(context.Context, *connect_go.Request[v15.StartRequest]) (*connect_go.Response[v15.StartResponse], error)
	Shutdown(context.Context, *connect_go.Request[v15.ShutdownRequest]) (*connect_go.Response[v15.ShutdownResponse], error)
	RunMysqlUpgrade(context.Context, *connect_go.Request[v15.RunMysqlUpgradeRequest]) (*connect_go.Response[v15.RunMysqlUpgradeResponse], error)
	ReinitConfig(context.Context, *connect_go.Request[v15.ReinitConfigRequest]) (*connect_go.Response[v15.ReinitConfigResponse], error)
	RefreshConfig(context.Context, *connect_go.Request[v15.RefreshConfigRequest]) (*connect_go.Response[v15.RefreshConfigResponse], error)
}

// NewMysqlCtlClient constructs a client for the vitess.mysqlctl.v15.MysqlCtl service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMysqlCtlClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) MysqlCtlClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &mysqlCtlClient{
		start: connect_go.NewClient[v15.StartRequest, v15.StartResponse](
			httpClient,
			baseURL+MysqlCtlStartProcedure,
			opts...,
		),
		shutdown: connect_go.NewClient[v15.ShutdownRequest, v15.ShutdownResponse](
			httpClient,
			baseURL+MysqlCtlShutdownProcedure,
			opts...,
		),
		runMysqlUpgrade: connect_go.NewClient[v15.RunMysqlUpgradeRequest, v15.RunMysqlUpgradeResponse](
			httpClient,
			baseURL+MysqlCtlRunMysqlUpgradeProcedure,
			opts...,
		),
		reinitConfig: connect_go.NewClient[v15.ReinitConfigRequest, v15.ReinitConfigResponse](
			httpClient,
			baseURL+MysqlCtlReinitConfigProcedure,
			opts...,
		),
		refreshConfig: connect_go.NewClient[v15.RefreshConfigRequest, v15.RefreshConfigResponse](
			httpClient,
			baseURL+MysqlCtlRefreshConfigProcedure,
			opts...,
		),
	}
}

// mysqlCtlClient implements MysqlCtlClient.
type mysqlCtlClient struct {
	start           *connect_go.Client[v15.StartRequest, v15.StartResponse]
	shutdown        *connect_go.Client[v15.ShutdownRequest, v15.ShutdownResponse]
	runMysqlUpgrade *connect_go.Client[v15.RunMysqlUpgradeRequest, v15.RunMysqlUpgradeResponse]
	reinitConfig    *connect_go.Client[v15.ReinitConfigRequest, v15.ReinitConfigResponse]
	refreshConfig   *connect_go.Client[v15.RefreshConfigRequest, v15.RefreshConfigResponse]
}

// Start calls vitess.mysqlctl.v15.MysqlCtl.Start.
func (c *mysqlCtlClient) Start(ctx context.Context, req *connect_go.Request[v15.StartRequest]) (*connect_go.Response[v15.StartResponse], error) {
	return c.start.CallUnary(ctx, req)
}

// Shutdown calls vitess.mysqlctl.v15.MysqlCtl.Shutdown.
func (c *mysqlCtlClient) Shutdown(ctx context.Context, req *connect_go.Request[v15.ShutdownRequest]) (*connect_go.Response[v15.ShutdownResponse], error) {
	return c.shutdown.CallUnary(ctx, req)
}

// RunMysqlUpgrade calls vitess.mysqlctl.v15.MysqlCtl.RunMysqlUpgrade.
func (c *mysqlCtlClient) RunMysqlUpgrade(ctx context.Context, req *connect_go.Request[v15.RunMysqlUpgradeRequest]) (*connect_go.Response[v15.RunMysqlUpgradeResponse], error) {
	return c.runMysqlUpgrade.CallUnary(ctx, req)
}

// ReinitConfig calls vitess.mysqlctl.v15.MysqlCtl.ReinitConfig.
func (c *mysqlCtlClient) ReinitConfig(ctx context.Context, req *connect_go.Request[v15.ReinitConfigRequest]) (*connect_go.Response[v15.ReinitConfigResponse], error) {
	return c.reinitConfig.CallUnary(ctx, req)
}

// RefreshConfig calls vitess.mysqlctl.v15.MysqlCtl.RefreshConfig.
func (c *mysqlCtlClient) RefreshConfig(ctx context.Context, req *connect_go.Request[v15.RefreshConfigRequest]) (*connect_go.Response[v15.RefreshConfigResponse], error) {
	return c.refreshConfig.CallUnary(ctx, req)
}

// MysqlCtlHandler is an implementation of the vitess.mysqlctl.v15.MysqlCtl service.
type MysqlCtlHandler interface {
	Start(context.Context, *connect_go.Request[v15.StartRequest]) (*connect_go.Response[v15.StartResponse], error)
	Shutdown(context.Context, *connect_go.Request[v15.ShutdownRequest]) (*connect_go.Response[v15.ShutdownResponse], error)
	RunMysqlUpgrade(context.Context, *connect_go.Request[v15.RunMysqlUpgradeRequest]) (*connect_go.Response[v15.RunMysqlUpgradeResponse], error)
	ReinitConfig(context.Context, *connect_go.Request[v15.ReinitConfigRequest]) (*connect_go.Response[v15.ReinitConfigResponse], error)
	RefreshConfig(context.Context, *connect_go.Request[v15.RefreshConfigRequest]) (*connect_go.Response[v15.RefreshConfigResponse], error)
}

// NewMysqlCtlHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMysqlCtlHandler(svc MysqlCtlHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(MysqlCtlStartProcedure, connect_go.NewUnaryHandler(
		MysqlCtlStartProcedure,
		svc.Start,
		opts...,
	))
	mux.Handle(MysqlCtlShutdownProcedure, connect_go.NewUnaryHandler(
		MysqlCtlShutdownProcedure,
		svc.Shutdown,
		opts...,
	))
	mux.Handle(MysqlCtlRunMysqlUpgradeProcedure, connect_go.NewUnaryHandler(
		MysqlCtlRunMysqlUpgradeProcedure,
		svc.RunMysqlUpgrade,
		opts...,
	))
	mux.Handle(MysqlCtlReinitConfigProcedure, connect_go.NewUnaryHandler(
		MysqlCtlReinitConfigProcedure,
		svc.ReinitConfig,
		opts...,
	))
	mux.Handle(MysqlCtlRefreshConfigProcedure, connect_go.NewUnaryHandler(
		MysqlCtlRefreshConfigProcedure,
		svc.RefreshConfig,
		opts...,
	))
	return "/vitess.mysqlctl.v15.MysqlCtl/", mux
}

// UnimplementedMysqlCtlHandler returns CodeUnimplemented from all methods.
type UnimplementedMysqlCtlHandler struct{}

func (UnimplementedMysqlCtlHandler) Start(context.Context, *connect_go.Request[v15.StartRequest]) (*connect_go.Response[v15.StartResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.mysqlctl.v15.MysqlCtl.Start is not implemented"))
}

func (UnimplementedMysqlCtlHandler) Shutdown(context.Context, *connect_go.Request[v15.ShutdownRequest]) (*connect_go.Response[v15.ShutdownResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.mysqlctl.v15.MysqlCtl.Shutdown is not implemented"))
}

func (UnimplementedMysqlCtlHandler) RunMysqlUpgrade(context.Context, *connect_go.Request[v15.RunMysqlUpgradeRequest]) (*connect_go.Response[v15.RunMysqlUpgradeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.mysqlctl.v15.MysqlCtl.RunMysqlUpgrade is not implemented"))
}

func (UnimplementedMysqlCtlHandler) ReinitConfig(context.Context, *connect_go.Request[v15.ReinitConfigRequest]) (*connect_go.Response[v15.ReinitConfigResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.mysqlctl.v15.MysqlCtl.ReinitConfig is not implemented"))
}

func (UnimplementedMysqlCtlHandler) RefreshConfig(context.Context, *connect_go.Request[v15.RefreshConfigRequest]) (*connect_go.Response[v15.RefreshConfigResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.mysqlctl.v15.MysqlCtl.RefreshConfig is not implemented"))
}
