// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: throttlerservice/throttlerservice.proto

package throttlerserviceconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	throttlerdata "github.com/planetscale/vitess-types/gen/vitess/dev/throttlerdata"
	_ "github.com/planetscale/vitess-types/gen/vitess/dev/throttlerservice"
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
	ThrottlerName = "throttlerservice.Throttler"
)

// ThrottlerClient is a client for the throttlerservice.Throttler service.
type ThrottlerClient interface {
	// MaxRates returns the current max rate for each throttler of the process.
	MaxRates(context.Context, *connect_go.Request[throttlerdata.MaxRatesRequest]) (*connect_go.Response[throttlerdata.MaxRatesResponse], error)
	// SetMaxRate allows to change the current max rate for all throttlers
	// of the process.
	SetMaxRate(context.Context, *connect_go.Request[throttlerdata.SetMaxRateRequest]) (*connect_go.Response[throttlerdata.SetMaxRateResponse], error)
	// GetConfiguration returns the configuration of the MaxReplicationlag module
	// for the given throttler or all throttlers if "throttler_name" is empty.
	GetConfiguration(context.Context, *connect_go.Request[throttlerdata.GetConfigurationRequest]) (*connect_go.Response[throttlerdata.GetConfigurationResponse], error)
	// UpdateConfiguration (partially) updates the configuration of the
	// MaxReplicationlag module for the given throttler or all throttlers if
	// "throttler_name" is empty.
	// If "copy_zero_values" is true, fields with zero values will be copied
	// as well.
	UpdateConfiguration(context.Context, *connect_go.Request[throttlerdata.UpdateConfigurationRequest]) (*connect_go.Response[throttlerdata.UpdateConfigurationResponse], error)
	// ResetConfiguration resets the configuration of the MaxReplicationlag module
	// to the initial configuration for the given throttler or all throttlers if
	// "throttler_name" is empty.
	ResetConfiguration(context.Context, *connect_go.Request[throttlerdata.ResetConfigurationRequest]) (*connect_go.Response[throttlerdata.ResetConfigurationResponse], error)
}

// NewThrottlerClient constructs a client for the throttlerservice.Throttler service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewThrottlerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ThrottlerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &throttlerClient{
		maxRates: connect_go.NewClient[throttlerdata.MaxRatesRequest, throttlerdata.MaxRatesResponse](
			httpClient,
			baseURL+"/throttlerservice.Throttler/MaxRates",
			opts...,
		),
		setMaxRate: connect_go.NewClient[throttlerdata.SetMaxRateRequest, throttlerdata.SetMaxRateResponse](
			httpClient,
			baseURL+"/throttlerservice.Throttler/SetMaxRate",
			opts...,
		),
		getConfiguration: connect_go.NewClient[throttlerdata.GetConfigurationRequest, throttlerdata.GetConfigurationResponse](
			httpClient,
			baseURL+"/throttlerservice.Throttler/GetConfiguration",
			opts...,
		),
		updateConfiguration: connect_go.NewClient[throttlerdata.UpdateConfigurationRequest, throttlerdata.UpdateConfigurationResponse](
			httpClient,
			baseURL+"/throttlerservice.Throttler/UpdateConfiguration",
			opts...,
		),
		resetConfiguration: connect_go.NewClient[throttlerdata.ResetConfigurationRequest, throttlerdata.ResetConfigurationResponse](
			httpClient,
			baseURL+"/throttlerservice.Throttler/ResetConfiguration",
			opts...,
		),
	}
}

// throttlerClient implements ThrottlerClient.
type throttlerClient struct {
	maxRates            *connect_go.Client[throttlerdata.MaxRatesRequest, throttlerdata.MaxRatesResponse]
	setMaxRate          *connect_go.Client[throttlerdata.SetMaxRateRequest, throttlerdata.SetMaxRateResponse]
	getConfiguration    *connect_go.Client[throttlerdata.GetConfigurationRequest, throttlerdata.GetConfigurationResponse]
	updateConfiguration *connect_go.Client[throttlerdata.UpdateConfigurationRequest, throttlerdata.UpdateConfigurationResponse]
	resetConfiguration  *connect_go.Client[throttlerdata.ResetConfigurationRequest, throttlerdata.ResetConfigurationResponse]
}

// MaxRates calls throttlerservice.Throttler.MaxRates.
func (c *throttlerClient) MaxRates(ctx context.Context, req *connect_go.Request[throttlerdata.MaxRatesRequest]) (*connect_go.Response[throttlerdata.MaxRatesResponse], error) {
	return c.maxRates.CallUnary(ctx, req)
}

// SetMaxRate calls throttlerservice.Throttler.SetMaxRate.
func (c *throttlerClient) SetMaxRate(ctx context.Context, req *connect_go.Request[throttlerdata.SetMaxRateRequest]) (*connect_go.Response[throttlerdata.SetMaxRateResponse], error) {
	return c.setMaxRate.CallUnary(ctx, req)
}

// GetConfiguration calls throttlerservice.Throttler.GetConfiguration.
func (c *throttlerClient) GetConfiguration(ctx context.Context, req *connect_go.Request[throttlerdata.GetConfigurationRequest]) (*connect_go.Response[throttlerdata.GetConfigurationResponse], error) {
	return c.getConfiguration.CallUnary(ctx, req)
}

// UpdateConfiguration calls throttlerservice.Throttler.UpdateConfiguration.
func (c *throttlerClient) UpdateConfiguration(ctx context.Context, req *connect_go.Request[throttlerdata.UpdateConfigurationRequest]) (*connect_go.Response[throttlerdata.UpdateConfigurationResponse], error) {
	return c.updateConfiguration.CallUnary(ctx, req)
}

// ResetConfiguration calls throttlerservice.Throttler.ResetConfiguration.
func (c *throttlerClient) ResetConfiguration(ctx context.Context, req *connect_go.Request[throttlerdata.ResetConfigurationRequest]) (*connect_go.Response[throttlerdata.ResetConfigurationResponse], error) {
	return c.resetConfiguration.CallUnary(ctx, req)
}

// ThrottlerHandler is an implementation of the throttlerservice.Throttler service.
type ThrottlerHandler interface {
	// MaxRates returns the current max rate for each throttler of the process.
	MaxRates(context.Context, *connect_go.Request[throttlerdata.MaxRatesRequest]) (*connect_go.Response[throttlerdata.MaxRatesResponse], error)
	// SetMaxRate allows to change the current max rate for all throttlers
	// of the process.
	SetMaxRate(context.Context, *connect_go.Request[throttlerdata.SetMaxRateRequest]) (*connect_go.Response[throttlerdata.SetMaxRateResponse], error)
	// GetConfiguration returns the configuration of the MaxReplicationlag module
	// for the given throttler or all throttlers if "throttler_name" is empty.
	GetConfiguration(context.Context, *connect_go.Request[throttlerdata.GetConfigurationRequest]) (*connect_go.Response[throttlerdata.GetConfigurationResponse], error)
	// UpdateConfiguration (partially) updates the configuration of the
	// MaxReplicationlag module for the given throttler or all throttlers if
	// "throttler_name" is empty.
	// If "copy_zero_values" is true, fields with zero values will be copied
	// as well.
	UpdateConfiguration(context.Context, *connect_go.Request[throttlerdata.UpdateConfigurationRequest]) (*connect_go.Response[throttlerdata.UpdateConfigurationResponse], error)
	// ResetConfiguration resets the configuration of the MaxReplicationlag module
	// to the initial configuration for the given throttler or all throttlers if
	// "throttler_name" is empty.
	ResetConfiguration(context.Context, *connect_go.Request[throttlerdata.ResetConfigurationRequest]) (*connect_go.Response[throttlerdata.ResetConfigurationResponse], error)
}

// NewThrottlerHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewThrottlerHandler(svc ThrottlerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/throttlerservice.Throttler/MaxRates", connect_go.NewUnaryHandler(
		"/throttlerservice.Throttler/MaxRates",
		svc.MaxRates,
		opts...,
	))
	mux.Handle("/throttlerservice.Throttler/SetMaxRate", connect_go.NewUnaryHandler(
		"/throttlerservice.Throttler/SetMaxRate",
		svc.SetMaxRate,
		opts...,
	))
	mux.Handle("/throttlerservice.Throttler/GetConfiguration", connect_go.NewUnaryHandler(
		"/throttlerservice.Throttler/GetConfiguration",
		svc.GetConfiguration,
		opts...,
	))
	mux.Handle("/throttlerservice.Throttler/UpdateConfiguration", connect_go.NewUnaryHandler(
		"/throttlerservice.Throttler/UpdateConfiguration",
		svc.UpdateConfiguration,
		opts...,
	))
	mux.Handle("/throttlerservice.Throttler/ResetConfiguration", connect_go.NewUnaryHandler(
		"/throttlerservice.Throttler/ResetConfiguration",
		svc.ResetConfiguration,
		opts...,
	))
	return "/throttlerservice.Throttler/", mux
}

// UnimplementedThrottlerHandler returns CodeUnimplemented from all methods.
type UnimplementedThrottlerHandler struct{}

func (UnimplementedThrottlerHandler) MaxRates(context.Context, *connect_go.Request[throttlerdata.MaxRatesRequest]) (*connect_go.Response[throttlerdata.MaxRatesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("throttlerservice.Throttler.MaxRates is not implemented"))
}

func (UnimplementedThrottlerHandler) SetMaxRate(context.Context, *connect_go.Request[throttlerdata.SetMaxRateRequest]) (*connect_go.Response[throttlerdata.SetMaxRateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("throttlerservice.Throttler.SetMaxRate is not implemented"))
}

func (UnimplementedThrottlerHandler) GetConfiguration(context.Context, *connect_go.Request[throttlerdata.GetConfigurationRequest]) (*connect_go.Response[throttlerdata.GetConfigurationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("throttlerservice.Throttler.GetConfiguration is not implemented"))
}

func (UnimplementedThrottlerHandler) UpdateConfiguration(context.Context, *connect_go.Request[throttlerdata.UpdateConfigurationRequest]) (*connect_go.Response[throttlerdata.UpdateConfigurationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("throttlerservice.Throttler.UpdateConfiguration is not implemented"))
}

func (UnimplementedThrottlerHandler) ResetConfiguration(context.Context, *connect_go.Request[throttlerdata.ResetConfigurationRequest]) (*connect_go.Response[throttlerdata.ResetConfigurationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("throttlerservice.Throttler.ResetConfiguration is not implemented"))
}
