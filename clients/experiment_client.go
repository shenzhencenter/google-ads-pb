// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package clients

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicespb "github.com/shenzhencenter/google-ads-pb/services"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newExperimentClientHook clientHook

// ExperimentCallOptions contains the retry settings for each method of ExperimentClient.
type ExperimentCallOptions struct {
	MutateExperiments []gax.CallOption
	EndExperiment []gax.CallOption
	ListExperimentAsyncErrors []gax.CallOption
	GraduateExperiment []gax.CallOption
	ScheduleExperiment []gax.CallOption
	PromoteExperiment []gax.CallOption
}

func defaultExperimentGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultExperimentCallOptions() *ExperimentCallOptions {
	return &ExperimentCallOptions{
		MutateExperiments: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		EndExperiment: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListExperimentAsyncErrors: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GraduateExperiment: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ScheduleExperiment: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		PromoteExperiment: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalExperimentClient is an interface that defines the methods availaible from Google Ads API.
type internalExperimentClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateExperiments(context.Context, *servicespb.MutateExperimentsRequest, ...gax.CallOption) (*servicespb.MutateExperimentsResponse, error)
	EndExperiment(context.Context, *servicespb.EndExperimentRequest, ...gax.CallOption) error
	ListExperimentAsyncErrors(context.Context, *servicespb.ListExperimentAsyncErrorsRequest, ...gax.CallOption) *StatusIterator
	GraduateExperiment(context.Context, *servicespb.GraduateExperimentRequest, ...gax.CallOption) error
	ScheduleExperiment(context.Context, *servicespb.ScheduleExperimentRequest, ...gax.CallOption) (*ScheduleExperimentOperation, error)
	ScheduleExperimentOperation(name string) *ScheduleExperimentOperation
	PromoteExperiment(context.Context, *servicespb.PromoteExperimentRequest, ...gax.CallOption) (*PromoteExperimentOperation, error)
	PromoteExperimentOperation(name string) *PromoteExperimentOperation
}

// ExperimentClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage experiments.
type ExperimentClient struct {
	// The internal transport-dependent client.
	internalClient internalExperimentClient

	// The call options for this service.
	CallOptions *ExperimentCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ExperimentClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ExperimentClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ExperimentClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateExperiments creates, updates, or removes experiments. Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// ExperimentError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *ExperimentClient) MutateExperiments(ctx context.Context, req *servicespb.MutateExperimentsRequest, opts ...gax.CallOption) (*servicespb.MutateExperimentsResponse, error) {
	return c.internalClient.MutateExperiments(ctx, req, opts...)
}

// EndExperiment immediately ends an experiment, changing the experiment’s scheduled
// end date and without waiting for end of day. End date is updated to be the
// time of the request.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// ExperimentError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *ExperimentClient) EndExperiment(ctx context.Context, req *servicespb.EndExperimentRequest, opts ...gax.CallOption) error {
	return c.internalClient.EndExperiment(ctx, req, opts...)
}

// ListExperimentAsyncErrors returns all errors that occurred during the last Experiment update (either
// scheduling or promotion).
// Supports standard list paging.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *ExperimentClient) ListExperimentAsyncErrors(ctx context.Context, req *servicespb.ListExperimentAsyncErrorsRequest, opts ...gax.CallOption) *StatusIterator {
	return c.internalClient.ListExperimentAsyncErrors(ctx, req, opts...)
}

// GraduateExperiment graduates an experiment to a full campaign.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// ExperimentError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *ExperimentClient) GraduateExperiment(ctx context.Context, req *servicespb.GraduateExperimentRequest, opts ...gax.CallOption) error {
	return c.internalClient.GraduateExperiment(ctx, req, opts...)
}

// ScheduleExperiment schedule an experiment. The in design campaign
// will be converted into a real campaign (called the experiment campaign)
// that will begin serving ads if successfully created.
//
// The experiment is scheduled immediately with status INITIALIZING.
// This method returns a long running operation that tracks the forking of the
// in design campaign. If the forking fails, a list of errors can be retrieved
// using the ListExperimentAsyncErrors method. The operation’s
// metadata will be a string containing the resource name of the created
// experiment.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// ExperimentError (at )
// DatabaseError (at )
// DateError (at )
// DateRangeError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
func (c *ExperimentClient) ScheduleExperiment(ctx context.Context, req *servicespb.ScheduleExperimentRequest, opts ...gax.CallOption) (*ScheduleExperimentOperation, error) {
	return c.internalClient.ScheduleExperiment(ctx, req, opts...)
}

// ScheduleExperimentOperation returns a new ScheduleExperimentOperation from a given name.
// The name must be that of a previously created ScheduleExperimentOperation, possibly from a different process.
func (c *ExperimentClient) ScheduleExperimentOperation(name string) *ScheduleExperimentOperation {
	return c.internalClient.ScheduleExperimentOperation(name)
}

// PromoteExperiment promotes the trial campaign thus applying changes in the trial campaign
// to the base campaign.
// This method returns a long running operation that tracks the promotion of
// the experiment campaign. If it fails, a list of errors can be retrieved
// using the ListExperimentAsyncErrors method. The operation’s
// metadata will be a string containing the resource name of the created
// experiment.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// ExperimentError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *ExperimentClient) PromoteExperiment(ctx context.Context, req *servicespb.PromoteExperimentRequest, opts ...gax.CallOption) (*PromoteExperimentOperation, error) {
	return c.internalClient.PromoteExperiment(ctx, req, opts...)
}

// PromoteExperimentOperation returns a new PromoteExperimentOperation from a given name.
// The name must be that of a previously created PromoteExperimentOperation, possibly from a different process.
func (c *ExperimentClient) PromoteExperimentOperation(name string) *PromoteExperimentOperation {
	return c.internalClient.PromoteExperimentOperation(name)
}

// experimentGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type experimentGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ExperimentClient
	CallOptions **ExperimentCallOptions

	// The gRPC API client.
	experimentClient servicespb.ExperimentServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewExperimentClient creates a new experiment service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage experiments.
func NewExperimentClient(ctx context.Context, opts ...option.ClientOption) (*ExperimentClient, error) {
	clientOpts := defaultExperimentGRPCClientOptions()
	if newExperimentClientHook != nil {
		hookOpts, err := newExperimentClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ExperimentClient{CallOptions: defaultExperimentCallOptions()}

	c := &experimentGRPCClient{
		connPool:    connPool,
		disableDeadlines: disableDeadlines,
		experimentClient: servicespb.NewExperimentServiceClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *experimentGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *experimentGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *experimentGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *experimentGRPCClient) MutateExperiments(ctx context.Context, req *servicespb.MutateExperimentsRequest, opts ...gax.CallOption) (*servicespb.MutateExperimentsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateExperiments[0:len((*c.CallOptions).MutateExperiments):len((*c.CallOptions).MutateExperiments)], opts...)
	var resp *servicespb.MutateExperimentsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.experimentClient.MutateExperiments(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *experimentGRPCClient) EndExperiment(ctx context.Context, req *servicespb.EndExperimentRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "experiment", url.QueryEscape(req.GetExperiment())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).EndExperiment[0:len((*c.CallOptions).EndExperiment):len((*c.CallOptions).EndExperiment)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.experimentClient.EndExperiment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *experimentGRPCClient) ListExperimentAsyncErrors(ctx context.Context, req *servicespb.ListExperimentAsyncErrorsRequest, opts ...gax.CallOption) *StatusIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListExperimentAsyncErrors[0:len((*c.CallOptions).ListExperimentAsyncErrors):len((*c.CallOptions).ListExperimentAsyncErrors)], opts...)
	it := &StatusIterator{}
	req = proto.Clone(req).(*servicespb.ListExperimentAsyncErrorsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*statuspb.Status, string, error) {
		resp := &servicespb.ListExperimentAsyncErrorsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.experimentClient.ListExperimentAsyncErrors(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetErrors(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *experimentGRPCClient) GraduateExperiment(ctx context.Context, req *servicespb.GraduateExperimentRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "experiment", url.QueryEscape(req.GetExperiment())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GraduateExperiment[0:len((*c.CallOptions).GraduateExperiment):len((*c.CallOptions).GraduateExperiment)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.experimentClient.GraduateExperiment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *experimentGRPCClient) ScheduleExperiment(ctx context.Context, req *servicespb.ScheduleExperimentRequest, opts ...gax.CallOption) (*ScheduleExperimentOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ScheduleExperiment[0:len((*c.CallOptions).ScheduleExperiment):len((*c.CallOptions).ScheduleExperiment)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.experimentClient.ScheduleExperiment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ScheduleExperimentOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *experimentGRPCClient) PromoteExperiment(ctx context.Context, req *servicespb.PromoteExperimentRequest, opts ...gax.CallOption) (*PromoteExperimentOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 14400000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).PromoteExperiment[0:len((*c.CallOptions).PromoteExperiment):len((*c.CallOptions).PromoteExperiment)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.experimentClient.PromoteExperiment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &PromoteExperimentOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// PromoteExperimentOperation manages a long-running operation from PromoteExperiment.
type PromoteExperimentOperation struct {
	lro *longrunning.Operation
}

// PromoteExperimentOperation returns a new PromoteExperimentOperation from a given name.
// The name must be that of a previously created PromoteExperimentOperation, possibly from a different process.
func (c *experimentGRPCClient) PromoteExperimentOperation(name string) *PromoteExperimentOperation {
	return &PromoteExperimentOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *PromoteExperimentOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *PromoteExperimentOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *PromoteExperimentOperation) Metadata() (*servicespb.PromoteExperimentMetadata, error) {
	var meta servicespb.PromoteExperimentMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *PromoteExperimentOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *PromoteExperimentOperation) Name() string {
	return op.lro.Name()
}

// ScheduleExperimentOperation manages a long-running operation from ScheduleExperiment.
type ScheduleExperimentOperation struct {
	lro *longrunning.Operation
}

// ScheduleExperimentOperation returns a new ScheduleExperimentOperation from a given name.
// The name must be that of a previously created ScheduleExperimentOperation, possibly from a different process.
func (c *experimentGRPCClient) ScheduleExperimentOperation(name string) *ScheduleExperimentOperation {
	return &ScheduleExperimentOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ScheduleExperimentOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *ScheduleExperimentOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *ScheduleExperimentOperation) Metadata() (*servicespb.ScheduleExperimentMetadata, error) {
	var meta servicespb.ScheduleExperimentMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ScheduleExperimentOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ScheduleExperimentOperation) Name() string {
	return op.lro.Name()
}
