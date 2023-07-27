// Copyright 2023 Google LLC
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

	gax "github.com/googleapis/gax-go/v2"
	servicespb "github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newCustomizerAttributeClientHook clientHook

// CustomizerAttributeCallOptions contains the retry settings for each method of CustomizerAttributeClient.
type CustomizerAttributeCallOptions struct {
	MutateCustomizerAttributes []gax.CallOption
}

func defaultCustomizerAttributeGRPCClientOptions() []option.ClientOption {
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

func defaultCustomizerAttributeCallOptions() *CustomizerAttributeCallOptions {
	return &CustomizerAttributeCallOptions{
		MutateCustomizerAttributes: []gax.CallOption{
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

// internalCustomizerAttributeClient is an interface that defines the methods available from Google Ads API.
type internalCustomizerAttributeClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateCustomizerAttributes(context.Context, *servicespb.MutateCustomizerAttributesRequest, ...gax.CallOption) (*servicespb.MutateCustomizerAttributesResponse, error)
}

// CustomizerAttributeClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage customizer attribute
type CustomizerAttributeClient struct {
	// The internal transport-dependent client.
	internalClient internalCustomizerAttributeClient

	// The call options for this service.
	CallOptions *CustomizerAttributeCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CustomizerAttributeClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CustomizerAttributeClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CustomizerAttributeClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateCustomizerAttributes creates, updates or removes customizer attributes. Operation statuses are
// returned.
func (c *CustomizerAttributeClient) MutateCustomizerAttributes(ctx context.Context, req *servicespb.MutateCustomizerAttributesRequest, opts ...gax.CallOption) (*servicespb.MutateCustomizerAttributesResponse, error) {
	return c.internalClient.MutateCustomizerAttributes(ctx, req, opts...)
}

// customizerAttributeGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type customizerAttributeGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CustomizerAttributeClient
	CallOptions **CustomizerAttributeCallOptions

	// The gRPC API client.
	customizerAttributeClient servicespb.CustomizerAttributeServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCustomizerAttributeClient creates a new customizer attribute service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage customizer attribute
func NewCustomizerAttributeClient(ctx context.Context, opts ...option.ClientOption) (*CustomizerAttributeClient, error) {
	clientOpts := defaultCustomizerAttributeGRPCClientOptions()
	if newCustomizerAttributeClientHook != nil {
		hookOpts, err := newCustomizerAttributeClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CustomizerAttributeClient{CallOptions: defaultCustomizerAttributeCallOptions()}

	c := &customizerAttributeGRPCClient{
		connPool:                  connPool,
		customizerAttributeClient: servicespb.NewCustomizerAttributeServiceClient(connPool),
		CallOptions:               &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *customizerAttributeGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *customizerAttributeGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *customizerAttributeGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *customizerAttributeGRPCClient) MutateCustomizerAttributes(ctx context.Context, req *servicespb.MutateCustomizerAttributesRequest, opts ...gax.CallOption) (*servicespb.MutateCustomizerAttributesResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateCustomizerAttributes[0:len((*c.CallOptions).MutateCustomizerAttributes):len((*c.CallOptions).MutateCustomizerAttributes)], opts...)
	var resp *servicespb.MutateCustomizerAttributesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customizerAttributeClient.MutateCustomizerAttributes(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
