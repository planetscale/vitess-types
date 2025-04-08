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
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	dev "github.com/planetscale/vitess-types/gen/vitess/mysqlctl/dev"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

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
	// MysqlCtlApplyBinlogFileProcedure is the fully-qualified name of the MysqlCtl's ApplyBinlogFile
	// RPC.
	MysqlCtlApplyBinlogFileProcedure = "/mysqlctl.MysqlCtl/ApplyBinlogFile"
	// MysqlCtlReadBinlogFilesTimestampsProcedure is the fully-qualified name of the MysqlCtl's
	// ReadBinlogFilesTimestamps RPC.
	MysqlCtlReadBinlogFilesTimestampsProcedure = "/mysqlctl.MysqlCtl/ReadBinlogFilesTimestamps"
	// MysqlCtlReinitConfigProcedure is the fully-qualified name of the MysqlCtl's ReinitConfig RPC.
	MysqlCtlReinitConfigProcedure = "/mysqlctl.MysqlCtl/ReinitConfig"
	// MysqlCtlRefreshConfigProcedure is the fully-qualified name of the MysqlCtl's RefreshConfig RPC.
	MysqlCtlRefreshConfigProcedure = "/mysqlctl.MysqlCtl/RefreshConfig"
	// MysqlCtlVersionStringProcedure is the fully-qualified name of the MysqlCtl's VersionString RPC.
	MysqlCtlVersionStringProcedure = "/mysqlctl.MysqlCtl/VersionString"
	// MysqlCtlHostMetricsProcedure is the fully-qualified name of the MysqlCtl's HostMetrics RPC.
	MysqlCtlHostMetricsProcedure = "/mysqlctl.MysqlCtl/HostMetrics"
)

// MysqlCtlClient is a client for the mysqlctl.MysqlCtl service.
type MysqlCtlClient interface {
	Start(context.Context, *connect.Request[dev.StartRequest]) (*connect.Response[dev.StartResponse], error)
	Shutdown(context.Context, *connect.Request[dev.ShutdownRequest]) (*connect.Response[dev.ShutdownResponse], error)
	RunMysqlUpgrade(context.Context, *connect.Request[dev.RunMysqlUpgradeRequest]) (*connect.Response[dev.RunMysqlUpgradeResponse], error)
	ApplyBinlogFile(context.Context, *connect.Request[dev.ApplyBinlogFileRequest]) (*connect.Response[dev.ApplyBinlogFileResponse], error)
	ReadBinlogFilesTimestamps(context.Context, *connect.Request[dev.ReadBinlogFilesTimestampsRequest]) (*connect.Response[dev.ReadBinlogFilesTimestampsResponse], error)
	ReinitConfig(context.Context, *connect.Request[dev.ReinitConfigRequest]) (*connect.Response[dev.ReinitConfigResponse], error)
	RefreshConfig(context.Context, *connect.Request[dev.RefreshConfigRequest]) (*connect.Response[dev.RefreshConfigResponse], error)
	VersionString(context.Context, *connect.Request[dev.VersionStringRequest]) (*connect.Response[dev.VersionStringResponse], error)
	HostMetrics(context.Context, *connect.Request[dev.HostMetricsRequest]) (*connect.Response[dev.HostMetricsResponse], error)
}

// NewMysqlCtlClient constructs a client for the mysqlctl.MysqlCtl service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMysqlCtlClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) MysqlCtlClient {
	baseURL = strings.TrimRight(baseURL, "/")
	mysqlCtlMethods := dev.File_vitess_mysqlctl_dev_mysqlctl_proto.Services().ByName("MysqlCtl").Methods()
	return &mysqlCtlClient{
		start: connect.NewClient[dev.StartRequest, dev.StartResponse](
			httpClient,
			baseURL+MysqlCtlStartProcedure,
			connect.WithSchema(mysqlCtlMethods.ByName("Start")),
			connect.WithClientOptions(opts...),
		),
		shutdown: connect.NewClient[dev.ShutdownRequest, dev.ShutdownResponse](
			httpClient,
			baseURL+MysqlCtlShutdownProcedure,
			connect.WithSchema(mysqlCtlMethods.ByName("Shutdown")),
			connect.WithClientOptions(opts...),
		),
		runMysqlUpgrade: connect.NewClient[dev.RunMysqlUpgradeRequest, dev.RunMysqlUpgradeResponse](
			httpClient,
			baseURL+MysqlCtlRunMysqlUpgradeProcedure,
			connect.WithSchema(mysqlCtlMethods.ByName("RunMysqlUpgrade")),
			connect.WithClientOptions(opts...),
		),
		applyBinlogFile: connect.NewClient[dev.ApplyBinlogFileRequest, dev.ApplyBinlogFileResponse](
			httpClient,
			baseURL+MysqlCtlApplyBinlogFileProcedure,
			connect.WithSchema(mysqlCtlMethods.ByName("ApplyBinlogFile")),
			connect.WithClientOptions(opts...),
		),
		readBinlogFilesTimestamps: connect.NewClient[dev.ReadBinlogFilesTimestampsRequest, dev.ReadBinlogFilesTimestampsResponse](
			httpClient,
			baseURL+MysqlCtlReadBinlogFilesTimestampsProcedure,
			connect.WithSchema(mysqlCtlMethods.ByName("ReadBinlogFilesTimestamps")),
			connect.WithClientOptions(opts...),
		),
		reinitConfig: connect.NewClient[dev.ReinitConfigRequest, dev.ReinitConfigResponse](
			httpClient,
			baseURL+MysqlCtlReinitConfigProcedure,
			connect.WithSchema(mysqlCtlMethods.ByName("ReinitConfig")),
			connect.WithClientOptions(opts...),
		),
		refreshConfig: connect.NewClient[dev.RefreshConfigRequest, dev.RefreshConfigResponse](
			httpClient,
			baseURL+MysqlCtlRefreshConfigProcedure,
			connect.WithSchema(mysqlCtlMethods.ByName("RefreshConfig")),
			connect.WithClientOptions(opts...),
		),
		versionString: connect.NewClient[dev.VersionStringRequest, dev.VersionStringResponse](
			httpClient,
			baseURL+MysqlCtlVersionStringProcedure,
			connect.WithSchema(mysqlCtlMethods.ByName("VersionString")),
			connect.WithClientOptions(opts...),
		),
		hostMetrics: connect.NewClient[dev.HostMetricsRequest, dev.HostMetricsResponse](
			httpClient,
			baseURL+MysqlCtlHostMetricsProcedure,
			connect.WithSchema(mysqlCtlMethods.ByName("HostMetrics")),
			connect.WithClientOptions(opts...),
		),
	}
}

// mysqlCtlClient implements MysqlCtlClient.
type mysqlCtlClient struct {
	start                     *connect.Client[dev.StartRequest, dev.StartResponse]
	shutdown                  *connect.Client[dev.ShutdownRequest, dev.ShutdownResponse]
	runMysqlUpgrade           *connect.Client[dev.RunMysqlUpgradeRequest, dev.RunMysqlUpgradeResponse]
	applyBinlogFile           *connect.Client[dev.ApplyBinlogFileRequest, dev.ApplyBinlogFileResponse]
	readBinlogFilesTimestamps *connect.Client[dev.ReadBinlogFilesTimestampsRequest, dev.ReadBinlogFilesTimestampsResponse]
	reinitConfig              *connect.Client[dev.ReinitConfigRequest, dev.ReinitConfigResponse]
	refreshConfig             *connect.Client[dev.RefreshConfigRequest, dev.RefreshConfigResponse]
	versionString             *connect.Client[dev.VersionStringRequest, dev.VersionStringResponse]
	hostMetrics               *connect.Client[dev.HostMetricsRequest, dev.HostMetricsResponse]
}

// Start calls mysqlctl.MysqlCtl.Start.
func (c *mysqlCtlClient) Start(ctx context.Context, req *connect.Request[dev.StartRequest]) (*connect.Response[dev.StartResponse], error) {
	return c.start.CallUnary(ctx, req)
}

// Shutdown calls mysqlctl.MysqlCtl.Shutdown.
func (c *mysqlCtlClient) Shutdown(ctx context.Context, req *connect.Request[dev.ShutdownRequest]) (*connect.Response[dev.ShutdownResponse], error) {
	return c.shutdown.CallUnary(ctx, req)
}

// RunMysqlUpgrade calls mysqlctl.MysqlCtl.RunMysqlUpgrade.
func (c *mysqlCtlClient) RunMysqlUpgrade(ctx context.Context, req *connect.Request[dev.RunMysqlUpgradeRequest]) (*connect.Response[dev.RunMysqlUpgradeResponse], error) {
	return c.runMysqlUpgrade.CallUnary(ctx, req)
}

// ApplyBinlogFile calls mysqlctl.MysqlCtl.ApplyBinlogFile.
func (c *mysqlCtlClient) ApplyBinlogFile(ctx context.Context, req *connect.Request[dev.ApplyBinlogFileRequest]) (*connect.Response[dev.ApplyBinlogFileResponse], error) {
	return c.applyBinlogFile.CallUnary(ctx, req)
}

// ReadBinlogFilesTimestamps calls mysqlctl.MysqlCtl.ReadBinlogFilesTimestamps.
func (c *mysqlCtlClient) ReadBinlogFilesTimestamps(ctx context.Context, req *connect.Request[dev.ReadBinlogFilesTimestampsRequest]) (*connect.Response[dev.ReadBinlogFilesTimestampsResponse], error) {
	return c.readBinlogFilesTimestamps.CallUnary(ctx, req)
}

// ReinitConfig calls mysqlctl.MysqlCtl.ReinitConfig.
func (c *mysqlCtlClient) ReinitConfig(ctx context.Context, req *connect.Request[dev.ReinitConfigRequest]) (*connect.Response[dev.ReinitConfigResponse], error) {
	return c.reinitConfig.CallUnary(ctx, req)
}

// RefreshConfig calls mysqlctl.MysqlCtl.RefreshConfig.
func (c *mysqlCtlClient) RefreshConfig(ctx context.Context, req *connect.Request[dev.RefreshConfigRequest]) (*connect.Response[dev.RefreshConfigResponse], error) {
	return c.refreshConfig.CallUnary(ctx, req)
}

// VersionString calls mysqlctl.MysqlCtl.VersionString.
func (c *mysqlCtlClient) VersionString(ctx context.Context, req *connect.Request[dev.VersionStringRequest]) (*connect.Response[dev.VersionStringResponse], error) {
	return c.versionString.CallUnary(ctx, req)
}

// HostMetrics calls mysqlctl.MysqlCtl.HostMetrics.
func (c *mysqlCtlClient) HostMetrics(ctx context.Context, req *connect.Request[dev.HostMetricsRequest]) (*connect.Response[dev.HostMetricsResponse], error) {
	return c.hostMetrics.CallUnary(ctx, req)
}

// MysqlCtlHandler is an implementation of the mysqlctl.MysqlCtl service.
type MysqlCtlHandler interface {
	Start(context.Context, *connect.Request[dev.StartRequest]) (*connect.Response[dev.StartResponse], error)
	Shutdown(context.Context, *connect.Request[dev.ShutdownRequest]) (*connect.Response[dev.ShutdownResponse], error)
	RunMysqlUpgrade(context.Context, *connect.Request[dev.RunMysqlUpgradeRequest]) (*connect.Response[dev.RunMysqlUpgradeResponse], error)
	ApplyBinlogFile(context.Context, *connect.Request[dev.ApplyBinlogFileRequest]) (*connect.Response[dev.ApplyBinlogFileResponse], error)
	ReadBinlogFilesTimestamps(context.Context, *connect.Request[dev.ReadBinlogFilesTimestampsRequest]) (*connect.Response[dev.ReadBinlogFilesTimestampsResponse], error)
	ReinitConfig(context.Context, *connect.Request[dev.ReinitConfigRequest]) (*connect.Response[dev.ReinitConfigResponse], error)
	RefreshConfig(context.Context, *connect.Request[dev.RefreshConfigRequest]) (*connect.Response[dev.RefreshConfigResponse], error)
	VersionString(context.Context, *connect.Request[dev.VersionStringRequest]) (*connect.Response[dev.VersionStringResponse], error)
	HostMetrics(context.Context, *connect.Request[dev.HostMetricsRequest]) (*connect.Response[dev.HostMetricsResponse], error)
}

// NewMysqlCtlHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMysqlCtlHandler(svc MysqlCtlHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	mysqlCtlMethods := dev.File_vitess_mysqlctl_dev_mysqlctl_proto.Services().ByName("MysqlCtl").Methods()
	mysqlCtlStartHandler := connect.NewUnaryHandler(
		MysqlCtlStartProcedure,
		svc.Start,
		connect.WithSchema(mysqlCtlMethods.ByName("Start")),
		connect.WithHandlerOptions(opts...),
	)
	mysqlCtlShutdownHandler := connect.NewUnaryHandler(
		MysqlCtlShutdownProcedure,
		svc.Shutdown,
		connect.WithSchema(mysqlCtlMethods.ByName("Shutdown")),
		connect.WithHandlerOptions(opts...),
	)
	mysqlCtlRunMysqlUpgradeHandler := connect.NewUnaryHandler(
		MysqlCtlRunMysqlUpgradeProcedure,
		svc.RunMysqlUpgrade,
		connect.WithSchema(mysqlCtlMethods.ByName("RunMysqlUpgrade")),
		connect.WithHandlerOptions(opts...),
	)
	mysqlCtlApplyBinlogFileHandler := connect.NewUnaryHandler(
		MysqlCtlApplyBinlogFileProcedure,
		svc.ApplyBinlogFile,
		connect.WithSchema(mysqlCtlMethods.ByName("ApplyBinlogFile")),
		connect.WithHandlerOptions(opts...),
	)
	mysqlCtlReadBinlogFilesTimestampsHandler := connect.NewUnaryHandler(
		MysqlCtlReadBinlogFilesTimestampsProcedure,
		svc.ReadBinlogFilesTimestamps,
		connect.WithSchema(mysqlCtlMethods.ByName("ReadBinlogFilesTimestamps")),
		connect.WithHandlerOptions(opts...),
	)
	mysqlCtlReinitConfigHandler := connect.NewUnaryHandler(
		MysqlCtlReinitConfigProcedure,
		svc.ReinitConfig,
		connect.WithSchema(mysqlCtlMethods.ByName("ReinitConfig")),
		connect.WithHandlerOptions(opts...),
	)
	mysqlCtlRefreshConfigHandler := connect.NewUnaryHandler(
		MysqlCtlRefreshConfigProcedure,
		svc.RefreshConfig,
		connect.WithSchema(mysqlCtlMethods.ByName("RefreshConfig")),
		connect.WithHandlerOptions(opts...),
	)
	mysqlCtlVersionStringHandler := connect.NewUnaryHandler(
		MysqlCtlVersionStringProcedure,
		svc.VersionString,
		connect.WithSchema(mysqlCtlMethods.ByName("VersionString")),
		connect.WithHandlerOptions(opts...),
	)
	mysqlCtlHostMetricsHandler := connect.NewUnaryHandler(
		MysqlCtlHostMetricsProcedure,
		svc.HostMetrics,
		connect.WithSchema(mysqlCtlMethods.ByName("HostMetrics")),
		connect.WithHandlerOptions(opts...),
	)
	return "/mysqlctl.MysqlCtl/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MysqlCtlStartProcedure:
			mysqlCtlStartHandler.ServeHTTP(w, r)
		case MysqlCtlShutdownProcedure:
			mysqlCtlShutdownHandler.ServeHTTP(w, r)
		case MysqlCtlRunMysqlUpgradeProcedure:
			mysqlCtlRunMysqlUpgradeHandler.ServeHTTP(w, r)
		case MysqlCtlApplyBinlogFileProcedure:
			mysqlCtlApplyBinlogFileHandler.ServeHTTP(w, r)
		case MysqlCtlReadBinlogFilesTimestampsProcedure:
			mysqlCtlReadBinlogFilesTimestampsHandler.ServeHTTP(w, r)
		case MysqlCtlReinitConfigProcedure:
			mysqlCtlReinitConfigHandler.ServeHTTP(w, r)
		case MysqlCtlRefreshConfigProcedure:
			mysqlCtlRefreshConfigHandler.ServeHTTP(w, r)
		case MysqlCtlVersionStringProcedure:
			mysqlCtlVersionStringHandler.ServeHTTP(w, r)
		case MysqlCtlHostMetricsProcedure:
			mysqlCtlHostMetricsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMysqlCtlHandler returns CodeUnimplemented from all methods.
type UnimplementedMysqlCtlHandler struct{}

func (UnimplementedMysqlCtlHandler) Start(context.Context, *connect.Request[dev.StartRequest]) (*connect.Response[dev.StartResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.Start is not implemented"))
}

func (UnimplementedMysqlCtlHandler) Shutdown(context.Context, *connect.Request[dev.ShutdownRequest]) (*connect.Response[dev.ShutdownResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.Shutdown is not implemented"))
}

func (UnimplementedMysqlCtlHandler) RunMysqlUpgrade(context.Context, *connect.Request[dev.RunMysqlUpgradeRequest]) (*connect.Response[dev.RunMysqlUpgradeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.RunMysqlUpgrade is not implemented"))
}

func (UnimplementedMysqlCtlHandler) ApplyBinlogFile(context.Context, *connect.Request[dev.ApplyBinlogFileRequest]) (*connect.Response[dev.ApplyBinlogFileResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.ApplyBinlogFile is not implemented"))
}

func (UnimplementedMysqlCtlHandler) ReadBinlogFilesTimestamps(context.Context, *connect.Request[dev.ReadBinlogFilesTimestampsRequest]) (*connect.Response[dev.ReadBinlogFilesTimestampsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.ReadBinlogFilesTimestamps is not implemented"))
}

func (UnimplementedMysqlCtlHandler) ReinitConfig(context.Context, *connect.Request[dev.ReinitConfigRequest]) (*connect.Response[dev.ReinitConfigResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.ReinitConfig is not implemented"))
}

func (UnimplementedMysqlCtlHandler) RefreshConfig(context.Context, *connect.Request[dev.RefreshConfigRequest]) (*connect.Response[dev.RefreshConfigResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.RefreshConfig is not implemented"))
}

func (UnimplementedMysqlCtlHandler) VersionString(context.Context, *connect.Request[dev.VersionStringRequest]) (*connect.Response[dev.VersionStringResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.VersionString is not implemented"))
}

func (UnimplementedMysqlCtlHandler) HostMetrics(context.Context, *connect.Request[dev.HostMetricsRequest]) (*connect.Response[dev.HostMetricsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mysqlctl.MysqlCtl.HostMetrics is not implemented"))
}
