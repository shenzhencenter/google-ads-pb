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

var newUserListCustomerTypeClientHook clientHook

// UserListCustomerTypeCallOptions contains the retry settings for each method of UserListCustomerTypeClient.
type UserListCustomerTypeCallOptions struct {
	MutateUserListCustomerTypes []gax.CallOption
}

func defaultUserListCustomerTypeGRPCClientOptions() []option.ClientOption {
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

func defaultUserListCustomerTypeCallOptions() *UserListCustomerTypeCallOptions {
	return &UserListCustomerTypeCallOptions{
		MutateUserListCustomerTypes: []gax.CallOption{
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

// internalUserListCustomerTypeClient is an interface that defines the methods available from Google Ads API.
type internalUserListCustomerTypeClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateUserListCustomerTypes(context.Context, *servicespb.MutateUserListCustomerTypesRequest, ...gax.CallOption) (*servicespb.MutateUserListCustomerTypesResponse, error)
}

// UserListCustomerTypeClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage user list customer types.
type UserListCustomerTypeClient struct {
	// The internal transport-dependent client.
	internalClient internalUserListCustomerTypeClient

	// The call options for this service.
	CallOptions *UserListCustomerTypeCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *UserListCustomerTypeClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *UserListCustomerTypeClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *UserListCustomerTypeClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateUserListCustomerTypes attach or remove user list customer types. Operation statuses
// are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// UserListCustomerTypeError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *UserListCustomerTypeClient) MutateUserListCustomerTypes(ctx context.Context, req *servicespb.MutateUserListCustomerTypesRequest, opts ...gax.CallOption) (*servicespb.MutateUserListCustomerTypesResponse, error) {
	return c.internalClient.MutateUserListCustomerTypes(ctx, req, opts...)
}

// userListCustomerTypeGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type userListCustomerTypeGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing UserListCustomerTypeClient
	CallOptions **UserListCustomerTypeCallOptions

	// The gRPC API client.
	userListCustomerTypeClient servicespb.UserListCustomerTypeServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewUserListCustomerTypeClient creates a new user list customer type service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage user list customer types.
func NewUserListCustomerTypeClient(ctx context.Context, opts ...option.ClientOption) (*UserListCustomerTypeClient, error) {
	clientOpts := defaultUserListCustomerTypeGRPCClientOptions()
	if newUserListCustomerTypeClientHook != nil {
		hookOpts, err := newUserListCustomerTypeClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := UserListCustomerTypeClient{CallOptions: defaultUserListCustomerTypeCallOptions()}

	c := &userListCustomerTypeGRPCClient{
		connPool:                   connPool,
		userListCustomerTypeClient: servicespb.NewUserListCustomerTypeServiceClient(connPool),
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
func (c *userListCustomerTypeGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *userListCustomerTypeGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *userListCustomerTypeGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *userListCustomerTypeGRPCClient) MutateUserListCustomerTypes(ctx context.Context, req *servicespb.MutateUserListCustomerTypesRequest, opts ...gax.CallOption) (*servicespb.MutateUserListCustomerTypesResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateUserListCustomerTypes[0:len((*c.CallOptions).MutateUserListCustomerTypes):len((*c.CallOptions).MutateUserListCustomerTypes)], opts...)
	var resp *servicespb.MutateUserListCustomerTypesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.userListCustomerTypeClient.MutateUserListCustomerTypes(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}