// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: google/ads/googleads/v17/services/account_link_service.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AccountLinkService_CreateAccountLink_FullMethodName = "/google.ads.googleads.v17.services.AccountLinkService/CreateAccountLink"
	AccountLinkService_MutateAccountLink_FullMethodName = "/google.ads.googleads.v17.services.AccountLinkService/MutateAccountLink"
)

// AccountLinkServiceClient is the client API for AccountLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// This service allows management of links between Google Ads accounts and other
// accounts.
type AccountLinkServiceClient interface {
	// Creates an account link.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[ThirdPartyAppAnalyticsLinkError]()
	CreateAccountLink(ctx context.Context, in *CreateAccountLinkRequest, opts ...grpc.CallOption) (*CreateAccountLinkResponse, error)
	// Creates or removes an account link.
	// From V5, create is not supported through
	// AccountLinkService.MutateAccountLink. Use
	// AccountLinkService.CreateAccountLink instead.
	//
	// List of thrown errors:
	//
	//	[AccountLinkError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateAccountLink(ctx context.Context, in *MutateAccountLinkRequest, opts ...grpc.CallOption) (*MutateAccountLinkResponse, error)
}

type accountLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountLinkServiceClient(cc grpc.ClientConnInterface) AccountLinkServiceClient {
	return &accountLinkServiceClient{cc}
}

func (c *accountLinkServiceClient) CreateAccountLink(ctx context.Context, in *CreateAccountLinkRequest, opts ...grpc.CallOption) (*CreateAccountLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAccountLinkResponse)
	err := c.cc.Invoke(ctx, AccountLinkService_CreateAccountLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountLinkServiceClient) MutateAccountLink(ctx context.Context, in *MutateAccountLinkRequest, opts ...grpc.CallOption) (*MutateAccountLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateAccountLinkResponse)
	err := c.cc.Invoke(ctx, AccountLinkService_MutateAccountLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountLinkServiceServer is the server API for AccountLinkService service.
// All implementations must embed UnimplementedAccountLinkServiceServer
// for forward compatibility.
//
// This service allows management of links between Google Ads accounts and other
// accounts.
type AccountLinkServiceServer interface {
	// Creates an account link.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[ThirdPartyAppAnalyticsLinkError]()
	CreateAccountLink(context.Context, *CreateAccountLinkRequest) (*CreateAccountLinkResponse, error)
	// Creates or removes an account link.
	// From V5, create is not supported through
	// AccountLinkService.MutateAccountLink. Use
	// AccountLinkService.CreateAccountLink instead.
	//
	// List of thrown errors:
	//
	//	[AccountLinkError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateAccountLink(context.Context, *MutateAccountLinkRequest) (*MutateAccountLinkResponse, error)
	mustEmbedUnimplementedAccountLinkServiceServer()
}

// UnimplementedAccountLinkServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccountLinkServiceServer struct{}

func (UnimplementedAccountLinkServiceServer) CreateAccountLink(context.Context, *CreateAccountLinkRequest) (*CreateAccountLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountLink not implemented")
}
func (UnimplementedAccountLinkServiceServer) MutateAccountLink(context.Context, *MutateAccountLinkRequest) (*MutateAccountLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAccountLink not implemented")
}
func (UnimplementedAccountLinkServiceServer) mustEmbedUnimplementedAccountLinkServiceServer() {}
func (UnimplementedAccountLinkServiceServer) testEmbeddedByValue()                            {}

// UnsafeAccountLinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountLinkServiceServer will
// result in compilation errors.
type UnsafeAccountLinkServiceServer interface {
	mustEmbedUnimplementedAccountLinkServiceServer()
}

func RegisterAccountLinkServiceServer(s grpc.ServiceRegistrar, srv AccountLinkServiceServer) {
	// If the following call pancis, it indicates UnimplementedAccountLinkServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AccountLinkService_ServiceDesc, srv)
}

func _AccountLinkService_CreateAccountLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountLinkServiceServer).CreateAccountLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountLinkService_CreateAccountLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountLinkServiceServer).CreateAccountLink(ctx, req.(*CreateAccountLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountLinkService_MutateAccountLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAccountLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountLinkServiceServer).MutateAccountLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountLinkService_MutateAccountLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountLinkServiceServer).MutateAccountLink(ctx, req.(*MutateAccountLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountLinkService_ServiceDesc is the grpc.ServiceDesc for AccountLinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountLinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.AccountLinkService",
	HandlerType: (*AccountLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccountLink",
			Handler:    _AccountLinkService_CreateAccountLink_Handler,
		},
		{
			MethodName: "MutateAccountLink",
			Handler:    _AccountLinkService_MutateAccountLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/account_link_service.proto",
}
