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
// gRPC RPC interface for the internal resharding throttler (go/vt/throttler)
// which is used by vreplication.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: vitess/throttlerservice/dev/throttlerservice.proto

package throttlerservicedevconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	dev "github.com/planetscale/vitess-types/gen/vitess/throttlerdata/dev"
	_ "github.com/planetscale/vitess-types/gen/vitess/throttlerservice/dev"
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
	// ThrottlerName is the fully-qualified name of the Throttler service.
	ThrottlerName = "vitess.throttlerservice.dev.Throttler"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ThrottlerMaxRatesProcedure is the fully-qualified name of the Throttler's MaxRates RPC.
	ThrottlerMaxRatesProcedure = "/vitess.throttlerservice.dev.Throttler/MaxRates"
	// ThrottlerSetMaxRateProcedure is the fully-qualified name of the Throttler's SetMaxRate RPC.
	ThrottlerSetMaxRateProcedure = "/vitess.throttlerservice.dev.Throttler/SetMaxRate"
	// ThrottlerGetConfigurationProcedure is the fully-qualified name of the Throttler's
	// GetConfiguration RPC.
	ThrottlerGetConfigurationProcedure = "/vitess.throttlerservice.dev.Throttler/GetConfiguration"
	// ThrottlerUpdateConfigurationProcedure is the fully-qualified name of the Throttler's
	// UpdateConfiguration RPC.
	ThrottlerUpdateConfigurationProcedure = "/vitess.throttlerservice.dev.Throttler/UpdateConfiguration"
	// ThrottlerResetConfigurationProcedure is the fully-qualified name of the Throttler's
	// ResetConfiguration RPC.
	ThrottlerResetConfigurationProcedure = "/vitess.throttlerservice.dev.Throttler/ResetConfiguration"
)

// ThrottlerClient is a client for the vitess.throttlerservice.dev.Throttler service.
type ThrottlerClient interface {
	// MaxRates returns the current max rate for each throttler of the process.
	MaxRates(context.Context, *connect_go.Request[dev.MaxRatesRequest]) (*connect_go.Response[dev.MaxRatesResponse], error)
	// SetMaxRate allows to change the current max rate for all throttlers
	// of the process.
	SetMaxRate(context.Context, *connect_go.Request[dev.SetMaxRateRequest]) (*connect_go.Response[dev.SetMaxRateResponse], error)
	// GetConfiguration returns the configuration of the MaxReplicationlag module
	// for the given throttler or all throttlers if "throttler_name" is empty.
	GetConfiguration(context.Context, *connect_go.Request[dev.GetConfigurationRequest]) (*connect_go.Response[dev.GetConfigurationResponse], error)
	// UpdateConfiguration (partially) updates the configuration of the
	// MaxReplicationlag module for the given throttler or all throttlers if
	// "throttler_name" is empty.
	// If "copy_zero_values" is true, fields with zero values will be copied
	// as well.
	UpdateConfiguration(context.Context, *connect_go.Request[dev.UpdateConfigurationRequest]) (*connect_go.Response[dev.UpdateConfigurationResponse], error)
	// ResetConfiguration resets the configuration of the MaxReplicationlag module
	// to the initial configuration for the given throttler or all throttlers if
	// "throttler_name" is empty.
	ResetConfiguration(context.Context, *connect_go.Request[dev.ResetConfigurationRequest]) (*connect_go.Response[dev.ResetConfigurationResponse], error)
}

// NewThrottlerClient constructs a client for the vitess.throttlerservice.dev.Throttler service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewThrottlerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ThrottlerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &throttlerClient{
		maxRates: connect_go.NewClient[dev.MaxRatesRequest, dev.MaxRatesResponse](
			httpClient,
			baseURL+ThrottlerMaxRatesProcedure,
			opts...,
		),
		setMaxRate: connect_go.NewClient[dev.SetMaxRateRequest, dev.SetMaxRateResponse](
			httpClient,
			baseURL+ThrottlerSetMaxRateProcedure,
			opts...,
		),
		getConfiguration: connect_go.NewClient[dev.GetConfigurationRequest, dev.GetConfigurationResponse](
			httpClient,
			baseURL+ThrottlerGetConfigurationProcedure,
			opts...,
		),
		updateConfiguration: connect_go.NewClient[dev.UpdateConfigurationRequest, dev.UpdateConfigurationResponse](
			httpClient,
			baseURL+ThrottlerUpdateConfigurationProcedure,
			opts...,
		),
		resetConfiguration: connect_go.NewClient[dev.ResetConfigurationRequest, dev.ResetConfigurationResponse](
			httpClient,
			baseURL+ThrottlerResetConfigurationProcedure,
			opts...,
		),
	}
}

// throttlerClient implements ThrottlerClient.
type throttlerClient struct {
	maxRates            *connect_go.Client[dev.MaxRatesRequest, dev.MaxRatesResponse]
	setMaxRate          *connect_go.Client[dev.SetMaxRateRequest, dev.SetMaxRateResponse]
	getConfiguration    *connect_go.Client[dev.GetConfigurationRequest, dev.GetConfigurationResponse]
	updateConfiguration *connect_go.Client[dev.UpdateConfigurationRequest, dev.UpdateConfigurationResponse]
	resetConfiguration  *connect_go.Client[dev.ResetConfigurationRequest, dev.ResetConfigurationResponse]
}

// MaxRates calls vitess.throttlerservice.dev.Throttler.MaxRates.
func (c *throttlerClient) MaxRates(ctx context.Context, req *connect_go.Request[dev.MaxRatesRequest]) (*connect_go.Response[dev.MaxRatesResponse], error) {
	return c.maxRates.CallUnary(ctx, req)
}

// SetMaxRate calls vitess.throttlerservice.dev.Throttler.SetMaxRate.
func (c *throttlerClient) SetMaxRate(ctx context.Context, req *connect_go.Request[dev.SetMaxRateRequest]) (*connect_go.Response[dev.SetMaxRateResponse], error) {
	return c.setMaxRate.CallUnary(ctx, req)
}

// GetConfiguration calls vitess.throttlerservice.dev.Throttler.GetConfiguration.
func (c *throttlerClient) GetConfiguration(ctx context.Context, req *connect_go.Request[dev.GetConfigurationRequest]) (*connect_go.Response[dev.GetConfigurationResponse], error) {
	return c.getConfiguration.CallUnary(ctx, req)
}

// UpdateConfiguration calls vitess.throttlerservice.dev.Throttler.UpdateConfiguration.
func (c *throttlerClient) UpdateConfiguration(ctx context.Context, req *connect_go.Request[dev.UpdateConfigurationRequest]) (*connect_go.Response[dev.UpdateConfigurationResponse], error) {
	return c.updateConfiguration.CallUnary(ctx, req)
}

// ResetConfiguration calls vitess.throttlerservice.dev.Throttler.ResetConfiguration.
func (c *throttlerClient) ResetConfiguration(ctx context.Context, req *connect_go.Request[dev.ResetConfigurationRequest]) (*connect_go.Response[dev.ResetConfigurationResponse], error) {
	return c.resetConfiguration.CallUnary(ctx, req)
}

// ThrottlerHandler is an implementation of the vitess.throttlerservice.dev.Throttler service.
type ThrottlerHandler interface {
	// MaxRates returns the current max rate for each throttler of the process.
	MaxRates(context.Context, *connect_go.Request[dev.MaxRatesRequest]) (*connect_go.Response[dev.MaxRatesResponse], error)
	// SetMaxRate allows to change the current max rate for all throttlers
	// of the process.
	SetMaxRate(context.Context, *connect_go.Request[dev.SetMaxRateRequest]) (*connect_go.Response[dev.SetMaxRateResponse], error)
	// GetConfiguration returns the configuration of the MaxReplicationlag module
	// for the given throttler or all throttlers if "throttler_name" is empty.
	GetConfiguration(context.Context, *connect_go.Request[dev.GetConfigurationRequest]) (*connect_go.Response[dev.GetConfigurationResponse], error)
	// UpdateConfiguration (partially) updates the configuration of the
	// MaxReplicationlag module for the given throttler or all throttlers if
	// "throttler_name" is empty.
	// If "copy_zero_values" is true, fields with zero values will be copied
	// as well.
	UpdateConfiguration(context.Context, *connect_go.Request[dev.UpdateConfigurationRequest]) (*connect_go.Response[dev.UpdateConfigurationResponse], error)
	// ResetConfiguration resets the configuration of the MaxReplicationlag module
	// to the initial configuration for the given throttler or all throttlers if
	// "throttler_name" is empty.
	ResetConfiguration(context.Context, *connect_go.Request[dev.ResetConfigurationRequest]) (*connect_go.Response[dev.ResetConfigurationResponse], error)
}

// NewThrottlerHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewThrottlerHandler(svc ThrottlerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ThrottlerMaxRatesProcedure, connect_go.NewUnaryHandler(
		ThrottlerMaxRatesProcedure,
		svc.MaxRates,
		opts...,
	))
	mux.Handle(ThrottlerSetMaxRateProcedure, connect_go.NewUnaryHandler(
		ThrottlerSetMaxRateProcedure,
		svc.SetMaxRate,
		opts...,
	))
	mux.Handle(ThrottlerGetConfigurationProcedure, connect_go.NewUnaryHandler(
		ThrottlerGetConfigurationProcedure,
		svc.GetConfiguration,
		opts...,
	))
	mux.Handle(ThrottlerUpdateConfigurationProcedure, connect_go.NewUnaryHandler(
		ThrottlerUpdateConfigurationProcedure,
		svc.UpdateConfiguration,
		opts...,
	))
	mux.Handle(ThrottlerResetConfigurationProcedure, connect_go.NewUnaryHandler(
		ThrottlerResetConfigurationProcedure,
		svc.ResetConfiguration,
		opts...,
	))
	return "/vitess.throttlerservice.dev.Throttler/", mux
}

// UnimplementedThrottlerHandler returns CodeUnimplemented from all methods.
type UnimplementedThrottlerHandler struct{}

func (UnimplementedThrottlerHandler) MaxRates(context.Context, *connect_go.Request[dev.MaxRatesRequest]) (*connect_go.Response[dev.MaxRatesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.throttlerservice.dev.Throttler.MaxRates is not implemented"))
}

func (UnimplementedThrottlerHandler) SetMaxRate(context.Context, *connect_go.Request[dev.SetMaxRateRequest]) (*connect_go.Response[dev.SetMaxRateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.throttlerservice.dev.Throttler.SetMaxRate is not implemented"))
}

func (UnimplementedThrottlerHandler) GetConfiguration(context.Context, *connect_go.Request[dev.GetConfigurationRequest]) (*connect_go.Response[dev.GetConfigurationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.throttlerservice.dev.Throttler.GetConfiguration is not implemented"))
}

func (UnimplementedThrottlerHandler) UpdateConfiguration(context.Context, *connect_go.Request[dev.UpdateConfigurationRequest]) (*connect_go.Response[dev.UpdateConfigurationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.throttlerservice.dev.Throttler.UpdateConfiguration is not implemented"))
}

func (UnimplementedThrottlerHandler) ResetConfiguration(context.Context, *connect_go.Request[dev.ResetConfigurationRequest]) (*connect_go.Response[dev.ResetConfigurationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("vitess.throttlerservice.dev.Throttler.ResetConfiguration is not implemented"))
}
