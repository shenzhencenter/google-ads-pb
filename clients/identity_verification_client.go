// Copyright 2024 Google LLC
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
)

var newIdentityVerificationClientHook clientHook

// IdentityVerificationCallOptions contains the retry settings for each method of IdentityVerificationClient.
type IdentityVerificationCallOptions struct {
	StartIdentityVerification []gax.CallOption
	GetIdentityVerification   []gax.CallOption
}

func defaultIdentityVerificationGRPCClientOptions() []option.ClientOption {
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

func defaultIdentityVerificationCallOptions() *IdentityVerificationCallOptions {
	return &IdentityVerificationCallOptions{
		StartIdentityVerification: []gax.CallOption{
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
		GetIdentityVerification: []gax.CallOption{
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

// internalIdentityVerificationClient is an interface that defines the methods available from Google Ads API.
type internalIdentityVerificationClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	StartIdentityVerification(context.Context, *servicespb.StartIdentityVerificationRequest, ...gax.CallOption) error
	GetIdentityVerification(context.Context, *servicespb.GetIdentityVerificationRequest, ...gax.CallOption) (*servicespb.GetIdentityVerificationResponse, error)
}

// IdentityVerificationClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for managing Identity Verification Service.
type IdentityVerificationClient struct {
	// The internal transport-dependent client.
	internalClient internalIdentityVerificationClient

	// The call options for this service.
	CallOptions *IdentityVerificationCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *IdentityVerificationClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *IdentityVerificationClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *IdentityVerificationClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// StartIdentityVerification starts Identity Verification for a given verification program type.
// Statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *IdentityVerificationClient) StartIdentityVerification(ctx context.Context, req *servicespb.StartIdentityVerificationRequest, opts ...gax.CallOption) error {
	return c.internalClient.StartIdentityVerification(ctx, req, opts...)
}

// GetIdentityVerification returns Identity Verification information.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *IdentityVerificationClient) GetIdentityVerification(ctx context.Context, req *servicespb.GetIdentityVerificationRequest, opts ...gax.CallOption) (*servicespb.GetIdentityVerificationResponse, error) {
	return c.internalClient.GetIdentityVerification(ctx, req, opts...)
}

// identityVerificationGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type identityVerificationGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing IdentityVerificationClient
	CallOptions **IdentityVerificationCallOptions

	// The gRPC API client.
	identityVerificationClient servicespb.IdentityVerificationServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewIdentityVerificationClient creates a new identity verification service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for managing Identity Verification Service.
func NewIdentityVerificationClient(ctx context.Context, opts ...option.ClientOption) (*IdentityVerificationClient, error) {
	clientOpts := defaultIdentityVerificationGRPCClientOptions()
	if newIdentityVerificationClientHook != nil {
		hookOpts, err := newIdentityVerificationClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := IdentityVerificationClient{CallOptions: defaultIdentityVerificationCallOptions()}

	c := &identityVerificationGRPCClient{
		connPool:                   connPool,
		identityVerificationClient: servicespb.NewIdentityVerificationServiceClient(connPool),
		CallOptions:                &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *identityVerificationGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *identityVerificationGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *identityVerificationGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *identityVerificationGRPCClient) StartIdentityVerification(ctx context.Context, req *servicespb.StartIdentityVerificationRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).StartIdentityVerification[0:len((*c.CallOptions).StartIdentityVerification):len((*c.CallOptions).StartIdentityVerification)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.identityVerificationClient.StartIdentityVerification(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *identityVerificationGRPCClient) GetIdentityVerification(ctx context.Context, req *servicespb.GetIdentityVerificationRequest, opts ...gax.CallOption) (*servicespb.GetIdentityVerificationResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetIdentityVerification[0:len((*c.CallOptions).GetIdentityVerification):len((*c.CallOptions).GetIdentityVerification)], opts...)
	var resp *servicespb.GetIdentityVerificationResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.identityVerificationClient.GetIdentityVerification(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
