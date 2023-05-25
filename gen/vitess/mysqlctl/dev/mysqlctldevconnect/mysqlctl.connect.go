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
// Source: vitess/mysqlctl/dev/mysqlctl.proto

package mysqlctldevconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	dev "github.com/planetscale/vitess-types/gen/vitess/mysqlctl/dev"
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
	MysqlCtlName = "mysqlctl.MysqlCtl"
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
	MysqlCtlStartProcedure = "/mysqlctl.MysqlCtl/Start"
	// MysqlCtlShutdownProcedure is the fully-qualified name of the MysqlCtl's Shutdown RPC.
	MysqlCtlShutdownProcedure = "/mysqlctl.MysqlCtl/Shutdown"
	// MysqlCtlRunMysqlUpgradeProcedure is the fully-qualified name of the MysqlCtl's RunMysqlUpgrade
	// RPC.
	MysqlCtlRunMysqlUpgradeProcedure = "/mysqlctl.MysqlCtl/RunMysqlUpgrade"
	// MysqlCtlReinitConfigProcedure is the fully-qualified name of the MysqlCtl's ReinitConfig RPC.
	MysqlCtlReinitConfigProcedure = "/mysqlctl.MysqlCtl/ReinitConfig"
	// MysqlCtlRefreshConfigProcedure is the fully-qualified name of the MysqlCtl's RefreshConfig RPC.
	MysqlCtlRefreshConfigProcedure = "/mysqlctl.MysqlCtl/RefreshConfig"
)

// MysqlCtlClient is a client for the mysqlctl.MysqlCtl service.
type MysqlCtlClient interface {
	Start(context.Context, *connect_go.Request[dev.StartRequest]) (*connect_go.Response[dev.StartResponse], error)
	Shutdown(context.Context, *connect_go.Request[dev.ShutdownRequest]) (*connect_go.Response[dev.ShutdownResponse], error)
	RunMysqlUpgrade(context.Context, *connect_go.Request[dev.RunMysqlUpgradeRequest]) (*connect_go.Response[dev.RunMysqlUpgradeResponse], error)
	ReinitConfig(context.Context, *connect_go.Request[dev.ReinitConfigRequest]) (*connect_go.Response[dev.ReinitConfigResponse], error)
	RefreshConfig(context.Context, *connect_go.Request[dev.RefreshConfigRequest]) (*connect_go.Response[dev.RefreshConfigResponse], error)
}

// NewMysqlCtlClient constructs a client for the mysqlctl.MysqlCtl service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMysqlCtlClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) MysqlCtlClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &mysqlCtlClient{
		start: connect_go.NewClient[dev.StartRequest, dev.StartResponse](
			httpClient,
			baseURL+MysqlCtlStartProcedure,
			opts...,
		),
		shutdown: connect_go.NewClient[dev.ShutdownRequest, dev.ShutdownResponse](
			httpClient,
			baseURL+MysqlCtlShutdownProcedure,
			opts...,
		),
		runMysqlUpgrade: connect_go.NewClient[dev.RunMysqlUpgradeRequest, dev.RunMysqlUpgradeResponse](
			httpClient,
			baseURL+MysqlCtlRunMysqlUpgradeProcedure,
			opts...,
		),
		reinitConfig: connect_go.NewClient[dev.ReinitConfigRequest, dev.ReinitConfigResponse](
			httpClient,
			baseURL+MysqlCtlReinitConfigProcedure,
			opts...,
		),
		refreshConfig: connect_go.NewClient[dev.RefreshConfigRequest, dev.RefreshConfigResponse](
			httpClient,
			baseURL+MysqlCtlRefreshConfigProcedure,
			opts...,
		),
	}
}

// mysqlCtlClient implements MysqlCtlClient.
type mysqlCtlClient struct {
	start           *connect_go.Client[dev.StartRequest, dev.StartResponse]
	shutdown        *connect_go.Client[dev.ShutdownRequest, dev.ShutdownResponse]
	runMysqlUpgrade *connect_go.Client[dev.RunMysqlUpgradeRequest, dev.RunMysqlUpgradeResponse]
	reinitConfig    *connect_go.Client[dev.ReinitConfigRequest, dev.ReinitConfigResponse]
	refreshConfig   *connect_go.Client[dev.RefreshConfigRequest, dev.RefreshConfigResponse]
}

// Start calls mysqlctl.MysqlCtl.Start.
func (c *mysqlCtlClient) Start(ctx context.Context, req *connect_go.Request[dev.StartRequest]) (*connect_go.Response[dev.StartResponse], error) {
	return c.start.CallUnary(ctx, req)
}

// Shutdown calls mysqlctl.MysqlCtl.Shutdown.
func (c *mysqlCtlClient) Shutdown(ctx context.Context, req *connect_go.Request[dev.ShutdownRequest]) (*connect_go.Response[dev.ShutdownResponse], error) {
	return c.shutdown.CallUnary(ctx, req)
}

// RunMysqlUpgrade calls mysqlctl.MysqlCtl.RunMysqlUpgrade.
func (c *mysqlCtlClient) RunMysqlUpgrade(ctx context.Context, req *connect_go.Request[dev.RunMysqlUpgradeRequest]) (*connect_go.Response[dev.RunMysqlUpgradeResponse], error) {
	return c.runMysqlUpgrade.CallUnary(ctx, req)
}

// ReinitConfig calls mysqlctl.MysqlCtl.ReinitConfig.
func (c *mysqlCtlClient) ReinitConfig(ctx context.Context, req *connect_go.Request[dev.ReinitConfigRequest]) (*connect_go.Response[dev.ReinitConfigResponse], error) {
	return c.reinitConfig.CallUnary(ctx, req)
}

// RefreshConfig calls mysqlctl.MysqlCtl.RefreshConfig.
func (c *mysqlCtlClient) RefreshConfig(ctx context.Context, req *connect_go.Request[dev.RefreshConfigRequest]) (*connect_go.Response[dev.RefreshConfigResponse], error) {
	return c.refreshConfig.CallUnary(ctx, req)
}

// MysqlCtlHandler is an implementation of the mysqlctl.MysqlCtl service.
type MysqlCtlHandler interface {
	Start(context.Context, *connect_go.Request[dev.StartRequest]) (*connect_go.Response[dev.StartResponse], error)
	Shutdown(context.Context, *connect_go.Request[dev.ShutdownRequest]) (*connect_go.Response[dev.ShutdownResponse], error)
	RunMysqlUpgrade(context.Context, *connect_go.Request[dev.RunMysqlUpgradeRequest]) (*connect_go.Response[dev.RunMysqlUpgradeResponse], error)
	ReinitConfig(context.Context, *connect_go.Request[dev.ReinitConfigRequest]) (*connect_go.Response[dev.ReinitConfigResponse], error)
	RefreshConfig(context.Context, *connect_go.Request[dev.RefreshConfigRequest]) (*connect_go.Response[dev.RefreshConfigResponse], error)
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
	return "/mysqlctl.MysqlCtl/", mux
}

// UnimplementedMysqlCtlHandler returns CodeUnimplemented from all methods.
type UnimplementedMysqlCtlHandler struct{}

func (UnimplementedMysqlCtlHandler) Start(context.Context, *connect_go.Request[dev.StartRequest]) (*connect_go.Response[dev.StartResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.Start is not implemented"))
}

func (UnimplementedMysqlCtlHandler) Shutdown(context.Context, *connect_go.Request[dev.ShutdownRequest]) (*connect_go.Response[dev.ShutdownResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.Shutdown is not implemented"))
}

func (UnimplementedMysqlCtlHandler) RunMysqlUpgrade(context.Context, *connect_go.Request[dev.RunMysqlUpgradeRequest]) (*connect_go.Response[dev.RunMysqlUpgradeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.RunMysqlUpgrade is not implemented"))
}

func (UnimplementedMysqlCtlHandler) ReinitConfig(context.Context, *connect_go.Request[dev.ReinitConfigRequest]) (*connect_go.Response[dev.ReinitConfigResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.ReinitConfig is not implemented"))
}

func (UnimplementedMysqlCtlHandler) RefreshConfig(context.Context, *connect_go.Request[dev.RefreshConfigRequest]) (*connect_go.Response[dev.RefreshConfigResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.RefreshConfig is not implemented"))
}
