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
// Protobuf service for the automation framework.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: vitess/automationservice/v15/automationservice.proto

package automationservicev15connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v15 "github.com/planetscale/vitess-types/gen/vitess/automation/v15"
	_ "github.com/planetscale/vitess-types/gen/vitess/automationservice/v15"
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
	// AutomationName is the fully-qualified name of the Automation service.
	AutomationName = "automationservice.Automation"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AutomationEnqueueClusterOperationProcedure is the fully-qualified name of the Automation's
	// EnqueueClusterOperation RPC.
	AutomationEnqueueClusterOperationProcedure = "/automationservice.Automation/EnqueueClusterOperation"
	// AutomationGetClusterOperationDetailsProcedure is the fully-qualified name of the Automation's
	// GetClusterOperationDetails RPC.
	AutomationGetClusterOperationDetailsProcedure = "/automationservice.Automation/GetClusterOperationDetails"
)

// AutomationClient is a client for the automationservice.Automation service.
type AutomationClient interface {
	// Start a cluster operation.
	EnqueueClusterOperation(context.Context, *connect.Request[v15.EnqueueClusterOperationRequest]) (*connect.Response[v15.EnqueueClusterOperationResponse], error)
	// TODO(mberlin): Polling this is bad. Implement a subscribe mechanism to wait for changes?
	// Get all details of an active cluster operation.
	GetClusterOperationDetails(context.Context, *connect.Request[v15.GetClusterOperationDetailsRequest]) (*connect.Response[v15.GetClusterOperationDetailsResponse], error)
}

// NewAutomationClient constructs a client for the automationservice.Automation service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAutomationClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) AutomationClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &automationClient{
		enqueueClusterOperation: connect.NewClient[v15.EnqueueClusterOperationRequest, v15.EnqueueClusterOperationResponse](
			httpClient,
			baseURL+AutomationEnqueueClusterOperationProcedure,
			opts...,
		),
		getClusterOperationDetails: connect.NewClient[v15.GetClusterOperationDetailsRequest, v15.GetClusterOperationDetailsResponse](
			httpClient,
			baseURL+AutomationGetClusterOperationDetailsProcedure,
			opts...,
		),
	}
}

// automationClient implements AutomationClient.
type automationClient struct {
	enqueueClusterOperation    *connect.Client[v15.EnqueueClusterOperationRequest, v15.EnqueueClusterOperationResponse]
	getClusterOperationDetails *connect.Client[v15.GetClusterOperationDetailsRequest, v15.GetClusterOperationDetailsResponse]
}

// EnqueueClusterOperation calls automationservice.Automation.EnqueueClusterOperation.
func (c *automationClient) EnqueueClusterOperation(ctx context.Context, req *connect.Request[v15.EnqueueClusterOperationRequest]) (*connect.Response[v15.EnqueueClusterOperationResponse], error) {
	return c.enqueueClusterOperation.CallUnary(ctx, req)
}

// GetClusterOperationDetails calls
// automationservice.Automation.GetClusterOperationDetails.
func (c *automationClient) GetClusterOperationDetails(ctx context.Context, req *connect.Request[v15.GetClusterOperationDetailsRequest]) (*connect.Response[v15.GetClusterOperationDetailsResponse], error) {
	return c.getClusterOperationDetails.CallUnary(ctx, req)
}

// AutomationHandler is an implementation of the automationservice.Automation service.
type AutomationHandler interface {
	// Start a cluster operation.
	EnqueueClusterOperation(context.Context, *connect.Request[v15.EnqueueClusterOperationRequest]) (*connect.Response[v15.EnqueueClusterOperationResponse], error)
	// TODO(mberlin): Polling this is bad. Implement a subscribe mechanism to wait for changes?
	// Get all details of an active cluster operation.
	GetClusterOperationDetails(context.Context, *connect.Request[v15.GetClusterOperationDetailsRequest]) (*connect.Response[v15.GetClusterOperationDetailsResponse], error)
}

// NewAutomationHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAutomationHandler(svc AutomationHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	automationEnqueueClusterOperationHandler := connect.NewUnaryHandler(
		AutomationEnqueueClusterOperationProcedure,
		svc.EnqueueClusterOperation,
		opts...,
	)
	automationGetClusterOperationDetailsHandler := connect.NewUnaryHandler(
		AutomationGetClusterOperationDetailsProcedure,
		svc.GetClusterOperationDetails,
		opts...,
	)
	return "/automationservice.Automation/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AutomationEnqueueClusterOperationProcedure:
			automationEnqueueClusterOperationHandler.ServeHTTP(w, r)
		case AutomationGetClusterOperationDetailsProcedure:
			automationGetClusterOperationDetailsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAutomationHandler returns CodeUnimplemented from all methods.
type UnimplementedAutomationHandler struct{}

func (UnimplementedAutomationHandler) EnqueueClusterOperation(context.Context, *connect.Request[v15.EnqueueClusterOperationRequest]) (*connect.Response[v15.EnqueueClusterOperationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("automationservice.Automation.EnqueueClusterOperation is not implemented"))
}

func (UnimplementedAutomationHandler) GetClusterOperationDetails(context.Context, *connect.Request[v15.GetClusterOperationDetailsRequest]) (*connect.Response[v15.GetClusterOperationDetailsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("automationservice.Automation.GetClusterOperationDetails is not implemented"))
}
