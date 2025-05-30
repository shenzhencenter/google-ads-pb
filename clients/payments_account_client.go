// Copyright 2025 Google LLC
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
	"log/slog"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	servicespb "github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newPaymentsAccountClientHook clientHook

// PaymentsAccountCallOptions contains the retry settings for each method of PaymentsAccountClient.
type PaymentsAccountCallOptions struct {
	ListPaymentsAccounts []gax.CallOption
}

func defaultPaymentsAccountGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("googleads.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPaymentsAccountCallOptions() *PaymentsAccountCallOptions {
	return &PaymentsAccountCallOptions{
		ListPaymentsAccounts: []gax.CallOption{
			gax.WithTimeout(14400000 * time.Millisecond),
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

// internalPaymentsAccountClient is an interface that defines the methods available from Google Ads API.
type internalPaymentsAccountClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListPaymentsAccounts(context.Context, *servicespb.ListPaymentsAccountsRequest, ...gax.CallOption) (*servicespb.ListPaymentsAccountsResponse, error)
}

// PaymentsAccountClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to provide payments accounts that can be used to set up consolidated
// billing.
type PaymentsAccountClient struct {
	// The internal transport-dependent client.
	internalClient internalPaymentsAccountClient

	// The call options for this service.
	CallOptions *PaymentsAccountCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PaymentsAccountClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PaymentsAccountClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PaymentsAccountClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListPaymentsAccounts returns all payments accounts associated with all managers
// between the login customer ID and specified serving customer in the
// hierarchy, inclusive.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// PaymentsAccountError (at )
// QuotaError (at )
// RequestError (at )
func (c *PaymentsAccountClient) ListPaymentsAccounts(ctx context.Context, req *servicespb.ListPaymentsAccountsRequest, opts ...gax.CallOption) (*servicespb.ListPaymentsAccountsResponse, error) {
	return c.internalClient.ListPaymentsAccounts(ctx, req, opts...)
}

// paymentsAccountGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type paymentsAccountGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing PaymentsAccountClient
	CallOptions **PaymentsAccountCallOptions

	// The gRPC API client.
	paymentsAccountClient servicespb.PaymentsAccountServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewPaymentsAccountClient creates a new payments account service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to provide payments accounts that can be used to set up consolidated
// billing.
func NewPaymentsAccountClient(ctx context.Context, opts ...option.ClientOption) (*PaymentsAccountClient, error) {
	clientOpts := defaultPaymentsAccountGRPCClientOptions()
	if newPaymentsAccountClientHook != nil {
		hookOpts, err := newPaymentsAccountClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PaymentsAccountClient{CallOptions: defaultPaymentsAccountCallOptions()}

	c := &paymentsAccountGRPCClient{
		connPool:              connPool,
		paymentsAccountClient: servicespb.NewPaymentsAccountServiceClient(connPool),
		CallOptions:           &client.CallOptions,
		logger:                internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *paymentsAccountGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *paymentsAccountGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *paymentsAccountGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *paymentsAccountGRPCClient) ListPaymentsAccounts(ctx context.Context, req *servicespb.ListPaymentsAccountsRequest, opts ...gax.CallOption) (*servicespb.ListPaymentsAccountsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListPaymentsAccounts[0:len((*c.CallOptions).ListPaymentsAccounts):len((*c.CallOptions).ListPaymentsAccounts)], opts...)
	var resp *servicespb.ListPaymentsAccountsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.paymentsAccountClient.ListPaymentsAccounts, req, settings.GRPC, c.logger, "ListPaymentsAccounts")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
