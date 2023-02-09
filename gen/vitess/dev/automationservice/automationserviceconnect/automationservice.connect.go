// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: automationservice/automationservice.proto

package automationserviceconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	automation "github.com/planetscale/vitess-types/gen/vitess/dev/automation"
	_ "github.com/planetscale/vitess-types/gen/vitess/dev/automationservice"
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
	// AutomationName is the fully-qualified name of the Automation service.
	AutomationName = "automationservice.Automation"
)

// AutomationClient is a client for the automationservice.Automation service.
type AutomationClient interface {
	// Start a cluster operation.
	EnqueueClusterOperation(context.Context, *connect_go.Request[automation.EnqueueClusterOperationRequest]) (*connect_go.Response[automation.EnqueueClusterOperationResponse], error)
	// TODO(mberlin): Polling this is bad. Implement a subscribe mechanism to wait for changes?
	// Get all details of an active cluster operation.
	GetClusterOperationDetails(context.Context, *connect_go.Request[automation.GetClusterOperationDetailsRequest]) (*connect_go.Response[automation.GetClusterOperationDetailsResponse], error)
}

// NewAutomationClient constructs a client for the automationservice.Automation service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAutomationClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AutomationClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &automationClient{
		enqueueClusterOperation: connect_go.NewClient[automation.EnqueueClusterOperationRequest, automation.EnqueueClusterOperationResponse](
			httpClient,
			baseURL+"/automationservice.Automation/EnqueueClusterOperation",
			opts...,
		),
		getClusterOperationDetails: connect_go.NewClient[automation.GetClusterOperationDetailsRequest, automation.GetClusterOperationDetailsResponse](
			httpClient,
			baseURL+"/automationservice.Automation/GetClusterOperationDetails",
			opts...,
		),
	}
}

// automationClient implements AutomationClient.
type automationClient struct {
	enqueueClusterOperation    *connect_go.Client[automation.EnqueueClusterOperationRequest, automation.EnqueueClusterOperationResponse]
	getClusterOperationDetails *connect_go.Client[automation.GetClusterOperationDetailsRequest, automation.GetClusterOperationDetailsResponse]
}

// EnqueueClusterOperation calls automationservice.Automation.EnqueueClusterOperation.
func (c *automationClient) EnqueueClusterOperation(ctx context.Context, req *connect_go.Request[automation.EnqueueClusterOperationRequest]) (*connect_go.Response[automation.EnqueueClusterOperationResponse], error) {
	return c.enqueueClusterOperation.CallUnary(ctx, req)
}

// GetClusterOperationDetails calls automationservice.Automation.GetClusterOperationDetails.
func (c *automationClient) GetClusterOperationDetails(ctx context.Context, req *connect_go.Request[automation.GetClusterOperationDetailsRequest]) (*connect_go.Response[automation.GetClusterOperationDetailsResponse], error) {
	return c.getClusterOperationDetails.CallUnary(ctx, req)
}

// AutomationHandler is an implementation of the automationservice.Automation service.
type AutomationHandler interface {
	// Start a cluster operation.
	EnqueueClusterOperation(context.Context, *connect_go.Request[automation.EnqueueClusterOperationRequest]) (*connect_go.Response[automation.EnqueueClusterOperationResponse], error)
	// TODO(mberlin): Polling this is bad. Implement a subscribe mechanism to wait for changes?
	// Get all details of an active cluster operation.
	GetClusterOperationDetails(context.Context, *connect_go.Request[automation.GetClusterOperationDetailsRequest]) (*connect_go.Response[automation.GetClusterOperationDetailsResponse], error)
}

// NewAutomationHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAutomationHandler(svc AutomationHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/automationservice.Automation/EnqueueClusterOperation", connect_go.NewUnaryHandler(
		"/automationservice.Automation/EnqueueClusterOperation",
		svc.EnqueueClusterOperation,
		opts...,
	))
	mux.Handle("/automationservice.Automation/GetClusterOperationDetails", connect_go.NewUnaryHandler(
		"/automationservice.Automation/GetClusterOperationDetails",
		svc.GetClusterOperationDetails,
		opts...,
	))
	return "/automationservice.Automation/", mux
}

// UnimplementedAutomationHandler returns CodeUnimplemented from all methods.
type UnimplementedAutomationHandler struct{}

func (UnimplementedAutomationHandler) EnqueueClusterOperation(context.Context, *connect_go.Request[automation.EnqueueClusterOperationRequest]) (*connect_go.Response[automation.EnqueueClusterOperationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("automationservice.Automation.EnqueueClusterOperation is not implemented"))
}

func (UnimplementedAutomationHandler) GetClusterOperationDetails(context.Context, *connect_go.Request[automation.GetClusterOperationDetailsRequest]) (*connect_go.Response[automation.GetClusterOperationDetailsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("automationservice.Automation.GetClusterOperationDetails is not implemented"))
}
