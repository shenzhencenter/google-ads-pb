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

var newAdGroupAdClientHook clientHook

// AdGroupAdCallOptions contains the retry settings for each method of AdGroupAdClient.
type AdGroupAdCallOptions struct {
	MutateAdGroupAds                 []gax.CallOption
	RemoveAutomaticallyCreatedAssets []gax.CallOption
}

func defaultAdGroupAdGRPCClientOptions() []option.ClientOption {
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

func defaultAdGroupAdCallOptions() *AdGroupAdCallOptions {
	return &AdGroupAdCallOptions{
		MutateAdGroupAds: []gax.CallOption{
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
		RemoveAutomaticallyCreatedAssets: []gax.CallOption{
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

// internalAdGroupAdClient is an interface that defines the methods available from Google Ads API.
type internalAdGroupAdClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateAdGroupAds(context.Context, *servicespb.MutateAdGroupAdsRequest, ...gax.CallOption) (*servicespb.MutateAdGroupAdsResponse, error)
	RemoveAutomaticallyCreatedAssets(context.Context, *servicespb.RemoveAutomaticallyCreatedAssetsRequest, ...gax.CallOption) error
}

// AdGroupAdClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage ads in an ad group.
type AdGroupAdClient struct {
	// The internal transport-dependent client.
	internalClient internalAdGroupAdClient

	// The call options for this service.
	CallOptions *AdGroupAdCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdGroupAdClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdGroupAdClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AdGroupAdClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateAdGroupAds creates, updates, or removes ads. Operation statuses are returned.
//
// List of thrown errors:
// AdCustomizerError (at )
// AdError (at )
// AdGroupAdError (at )
// AdSharingError (at )
// AdxError (at )
// AssetError (at )
// AssetLinkError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// ContextError (at )
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
// OperationAccessDeniedError (at )
// OperatorError (at )
// PolicyFindingError (at )
// PolicyValidationParameterError (at )
// PolicyViolationError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// ResourceCountLimitExceededError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
// UrlFieldError (at )
func (c *AdGroupAdClient) MutateAdGroupAds(ctx context.Context, req *servicespb.MutateAdGroupAdsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupAdsResponse, error) {
	return c.internalClient.MutateAdGroupAds(ctx, req, opts...)
}

// RemoveAutomaticallyCreatedAssets remove automatically created assets from an ad.
//
// List of thrown errors:
// AdError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// AutomaticallyCreatedAssetRemovalError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *AdGroupAdClient) RemoveAutomaticallyCreatedAssets(ctx context.Context, req *servicespb.RemoveAutomaticallyCreatedAssetsRequest, opts ...gax.CallOption) error {
	return c.internalClient.RemoveAutomaticallyCreatedAssets(ctx, req, opts...)
}

// adGroupAdGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adGroupAdGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AdGroupAdClient
	CallOptions **AdGroupAdCallOptions

	// The gRPC API client.
	adGroupAdClient servicespb.AdGroupAdServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewAdGroupAdClient creates a new ad group ad service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage ads in an ad group.
func NewAdGroupAdClient(ctx context.Context, opts ...option.ClientOption) (*AdGroupAdClient, error) {
	clientOpts := defaultAdGroupAdGRPCClientOptions()
	if newAdGroupAdClientHook != nil {
		hookOpts, err := newAdGroupAdClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AdGroupAdClient{CallOptions: defaultAdGroupAdCallOptions()}

	c := &adGroupAdGRPCClient{
		connPool:        connPool,
		adGroupAdClient: servicespb.NewAdGroupAdServiceClient(connPool),
		CallOptions:     &client.CallOptions,
		logger:          internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *adGroupAdGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adGroupAdGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adGroupAdGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adGroupAdGRPCClient) MutateAdGroupAds(ctx context.Context, req *servicespb.MutateAdGroupAdsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupAdsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateAdGroupAds[0:len((*c.CallOptions).MutateAdGroupAds):len((*c.CallOptions).MutateAdGroupAds)], opts...)
	var resp *servicespb.MutateAdGroupAdsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.adGroupAdClient.MutateAdGroupAds, req, settings.GRPC, c.logger, "MutateAdGroupAds")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *adGroupAdGRPCClient) RemoveAutomaticallyCreatedAssets(ctx context.Context, req *servicespb.RemoveAutomaticallyCreatedAssetsRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "ad_group_ad", url.QueryEscape(req.GetAdGroupAd()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RemoveAutomaticallyCreatedAssets[0:len((*c.CallOptions).RemoveAutomaticallyCreatedAssets):len((*c.CallOptions).RemoveAutomaticallyCreatedAssets)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.adGroupAdClient.RemoveAutomaticallyCreatedAssets, req, settings.GRPC, c.logger, "RemoveAutomaticallyCreatedAssets")
		return err
	}, opts...)
	return err
}
