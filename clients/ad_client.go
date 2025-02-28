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

var newAdClientHook clientHook

// AdCallOptions contains the retry settings for each method of AdClient.
type AdCallOptions struct {
	MutateAds []gax.CallOption
}

func defaultAdGRPCClientOptions() []option.ClientOption {
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

func defaultAdCallOptions() *AdCallOptions {
	return &AdCallOptions{
		MutateAds: []gax.CallOption{
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

// internalAdClient is an interface that defines the methods available from Google Ads API.
type internalAdClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateAds(context.Context, *servicespb.MutateAdsRequest, ...gax.CallOption) (*servicespb.MutateAdsResponse, error)
}

// AdClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage ads.
type AdClient struct {
	// The internal transport-dependent client.
	internalClient internalAdClient

	// The call options for this service.
	CallOptions *AdCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AdClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateAds updates ads. Operation statuses are returned. Updating ads is not supported
// for TextAd, ExpandedDynamicSearchAd, GmailAd and ImageAd.
//
// List of thrown errors:
// AdCustomizerError (at )
// AdError (at )
// AdSharingError (at )
// AdxError (at )
// AssetError (at )
// AssetLinkError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// DatabaseError (at )
// DateError (at )
// DistinctError (at )
// FeedAttributeReferenceError (at )
// FieldError (at )
// FieldMaskError (at )
// FunctionError (at )
// FunctionParsingError (at )
// HeaderError (at )
// IdError (at )
// ImageError (at )
// InternalError (at )
// ListOperationError (at )
// MediaBundleError (at )
// MediaFileError (at )
// MutateError (at )
// NewResourceCreationError (at )
// NotEmptyError (at )
// NullError (at )
// OperatorError (at )
// PolicyFindingError (at )
// PolicyViolationError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
// UrlFieldError (at )
func (c *AdClient) MutateAds(ctx context.Context, req *servicespb.MutateAdsRequest, opts ...gax.CallOption) (*servicespb.MutateAdsResponse, error) {
	return c.internalClient.MutateAds(ctx, req, opts...)
}

// adGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AdClient
	CallOptions **AdCallOptions

	// The gRPC API client.
	adClient servicespb.AdServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewAdClient creates a new ad service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage ads.
func NewAdClient(ctx context.Context, opts ...option.ClientOption) (*AdClient, error) {
	clientOpts := defaultAdGRPCClientOptions()
	if newAdClientHook != nil {
		hookOpts, err := newAdClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AdClient{CallOptions: defaultAdCallOptions()}

	c := &adGRPCClient{
		connPool:    connPool,
		adClient:    servicespb.NewAdServiceClient(connPool),
		CallOptions: &client.CallOptions,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *adGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adGRPCClient) MutateAds(ctx context.Context, req *servicespb.MutateAdsRequest, opts ...gax.CallOption) (*servicespb.MutateAdsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateAds[0:len((*c.CallOptions).MutateAds):len((*c.CallOptions).MutateAds)], opts...)
	var resp *servicespb.MutateAdsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.adClient.MutateAds, req, settings.GRPC, c.logger, "MutateAds")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
