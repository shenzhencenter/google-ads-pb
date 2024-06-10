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

var newProductLinkInvitationClientHook clientHook

// ProductLinkInvitationCallOptions contains the retry settings for each method of ProductLinkInvitationClient.
type ProductLinkInvitationCallOptions struct {
	CreateProductLinkInvitation []gax.CallOption
	UpdateProductLinkInvitation []gax.CallOption
	RemoveProductLinkInvitation []gax.CallOption
}

func defaultProductLinkInvitationGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("googleads.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultProductLinkInvitationCallOptions() *ProductLinkInvitationCallOptions {
	return &ProductLinkInvitationCallOptions{
		CreateProductLinkInvitation: []gax.CallOption{
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
		UpdateProductLinkInvitation: []gax.CallOption{
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
		RemoveProductLinkInvitation: []gax.CallOption{
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

// internalProductLinkInvitationClient is an interface that defines the methods available from Google Ads API.
type internalProductLinkInvitationClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateProductLinkInvitation(context.Context, *servicespb.CreateProductLinkInvitationRequest, ...gax.CallOption) (*servicespb.CreateProductLinkInvitationResponse, error)
	UpdateProductLinkInvitation(context.Context, *servicespb.UpdateProductLinkInvitationRequest, ...gax.CallOption) (*servicespb.UpdateProductLinkInvitationResponse, error)
	RemoveProductLinkInvitation(context.Context, *servicespb.RemoveProductLinkInvitationRequest, ...gax.CallOption) (*servicespb.RemoveProductLinkInvitationResponse, error)
}

// ProductLinkInvitationClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// This service allows management of product link invitations from Google Ads
// accounts to other accounts.
type ProductLinkInvitationClient struct {
	// The internal transport-dependent client.
	internalClient internalProductLinkInvitationClient

	// The call options for this service.
	CallOptions *ProductLinkInvitationCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ProductLinkInvitationClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ProductLinkInvitationClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ProductLinkInvitationClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateProductLinkInvitation creates a product link invitation.
func (c *ProductLinkInvitationClient) CreateProductLinkInvitation(ctx context.Context, req *servicespb.CreateProductLinkInvitationRequest, opts ...gax.CallOption) (*servicespb.CreateProductLinkInvitationResponse, error) {
	return c.internalClient.CreateProductLinkInvitation(ctx, req, opts...)
}

// UpdateProductLinkInvitation update a product link invitation.
func (c *ProductLinkInvitationClient) UpdateProductLinkInvitation(ctx context.Context, req *servicespb.UpdateProductLinkInvitationRequest, opts ...gax.CallOption) (*servicespb.UpdateProductLinkInvitationResponse, error) {
	return c.internalClient.UpdateProductLinkInvitation(ctx, req, opts...)
}

// RemoveProductLinkInvitation remove a product link invitation.
func (c *ProductLinkInvitationClient) RemoveProductLinkInvitation(ctx context.Context, req *servicespb.RemoveProductLinkInvitationRequest, opts ...gax.CallOption) (*servicespb.RemoveProductLinkInvitationResponse, error) {
	return c.internalClient.RemoveProductLinkInvitation(ctx, req, opts...)
}

// productLinkInvitationGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type productLinkInvitationGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing ProductLinkInvitationClient
	CallOptions **ProductLinkInvitationCallOptions

	// The gRPC API client.
	productLinkInvitationClient servicespb.ProductLinkInvitationServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewProductLinkInvitationClient creates a new product link invitation service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// This service allows management of product link invitations from Google Ads
// accounts to other accounts.
func NewProductLinkInvitationClient(ctx context.Context, opts ...option.ClientOption) (*ProductLinkInvitationClient, error) {
	clientOpts := defaultProductLinkInvitationGRPCClientOptions()
	if newProductLinkInvitationClientHook != nil {
		hookOpts, err := newProductLinkInvitationClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ProductLinkInvitationClient{CallOptions: defaultProductLinkInvitationCallOptions()}

	c := &productLinkInvitationGRPCClient{
		connPool:                    connPool,
		productLinkInvitationClient: servicespb.NewProductLinkInvitationServiceClient(connPool),
		CallOptions:                 &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *productLinkInvitationGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *productLinkInvitationGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *productLinkInvitationGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *productLinkInvitationGRPCClient) CreateProductLinkInvitation(ctx context.Context, req *servicespb.CreateProductLinkInvitationRequest, opts ...gax.CallOption) (*servicespb.CreateProductLinkInvitationResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateProductLinkInvitation[0:len((*c.CallOptions).CreateProductLinkInvitation):len((*c.CallOptions).CreateProductLinkInvitation)], opts...)
	var resp *servicespb.CreateProductLinkInvitationResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productLinkInvitationClient.CreateProductLinkInvitation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *productLinkInvitationGRPCClient) UpdateProductLinkInvitation(ctx context.Context, req *servicespb.UpdateProductLinkInvitationRequest, opts ...gax.CallOption) (*servicespb.UpdateProductLinkInvitationResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateProductLinkInvitation[0:len((*c.CallOptions).UpdateProductLinkInvitation):len((*c.CallOptions).UpdateProductLinkInvitation)], opts...)
	var resp *servicespb.UpdateProductLinkInvitationResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productLinkInvitationClient.UpdateProductLinkInvitation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *productLinkInvitationGRPCClient) RemoveProductLinkInvitation(ctx context.Context, req *servicespb.RemoveProductLinkInvitationRequest, opts ...gax.CallOption) (*servicespb.RemoveProductLinkInvitationResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).RemoveProductLinkInvitation[0:len((*c.CallOptions).RemoveProductLinkInvitation):len((*c.CallOptions).RemoveProductLinkInvitation)], opts...)
	var resp *servicespb.RemoveProductLinkInvitationResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.productLinkInvitationClient.RemoveProductLinkInvitation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
