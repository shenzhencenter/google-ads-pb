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
// source: google/ads/googleads/v18/services/product_link_invitation_service.proto

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
	ProductLinkInvitationService_CreateProductLinkInvitation_FullMethodName = "/google.ads.googleads.v18.services.ProductLinkInvitationService/CreateProductLinkInvitation"
	ProductLinkInvitationService_UpdateProductLinkInvitation_FullMethodName = "/google.ads.googleads.v18.services.ProductLinkInvitationService/UpdateProductLinkInvitation"
	ProductLinkInvitationService_RemoveProductLinkInvitation_FullMethodName = "/google.ads.googleads.v18.services.ProductLinkInvitationService/RemoveProductLinkInvitation"
)

// ProductLinkInvitationServiceClient is the client API for ProductLinkInvitationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// This service allows management of product link invitations from Google Ads
// accounts to other accounts.
type ProductLinkInvitationServiceClient interface {
	// Creates a product link invitation.
	CreateProductLinkInvitation(ctx context.Context, in *CreateProductLinkInvitationRequest, opts ...grpc.CallOption) (*CreateProductLinkInvitationResponse, error)
	// Update a product link invitation.
	UpdateProductLinkInvitation(ctx context.Context, in *UpdateProductLinkInvitationRequest, opts ...grpc.CallOption) (*UpdateProductLinkInvitationResponse, error)
	// Remove a product link invitation.
	RemoveProductLinkInvitation(ctx context.Context, in *RemoveProductLinkInvitationRequest, opts ...grpc.CallOption) (*RemoveProductLinkInvitationResponse, error)
}

type productLinkInvitationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductLinkInvitationServiceClient(cc grpc.ClientConnInterface) ProductLinkInvitationServiceClient {
	return &productLinkInvitationServiceClient{cc}
}

func (c *productLinkInvitationServiceClient) CreateProductLinkInvitation(ctx context.Context, in *CreateProductLinkInvitationRequest, opts ...grpc.CallOption) (*CreateProductLinkInvitationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProductLinkInvitationResponse)
	err := c.cc.Invoke(ctx, ProductLinkInvitationService_CreateProductLinkInvitation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productLinkInvitationServiceClient) UpdateProductLinkInvitation(ctx context.Context, in *UpdateProductLinkInvitationRequest, opts ...grpc.CallOption) (*UpdateProductLinkInvitationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProductLinkInvitationResponse)
	err := c.cc.Invoke(ctx, ProductLinkInvitationService_UpdateProductLinkInvitation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productLinkInvitationServiceClient) RemoveProductLinkInvitation(ctx context.Context, in *RemoveProductLinkInvitationRequest, opts ...grpc.CallOption) (*RemoveProductLinkInvitationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveProductLinkInvitationResponse)
	err := c.cc.Invoke(ctx, ProductLinkInvitationService_RemoveProductLinkInvitation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductLinkInvitationServiceServer is the server API for ProductLinkInvitationService service.
// All implementations must embed UnimplementedProductLinkInvitationServiceServer
// for forward compatibility.
//
// This service allows management of product link invitations from Google Ads
// accounts to other accounts.
type ProductLinkInvitationServiceServer interface {
	// Creates a product link invitation.
	CreateProductLinkInvitation(context.Context, *CreateProductLinkInvitationRequest) (*CreateProductLinkInvitationResponse, error)
	// Update a product link invitation.
	UpdateProductLinkInvitation(context.Context, *UpdateProductLinkInvitationRequest) (*UpdateProductLinkInvitationResponse, error)
	// Remove a product link invitation.
	RemoveProductLinkInvitation(context.Context, *RemoveProductLinkInvitationRequest) (*RemoveProductLinkInvitationResponse, error)
	mustEmbedUnimplementedProductLinkInvitationServiceServer()
}

// UnimplementedProductLinkInvitationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductLinkInvitationServiceServer struct{}

func (UnimplementedProductLinkInvitationServiceServer) CreateProductLinkInvitation(context.Context, *CreateProductLinkInvitationRequest) (*CreateProductLinkInvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductLinkInvitation not implemented")
}
func (UnimplementedProductLinkInvitationServiceServer) UpdateProductLinkInvitation(context.Context, *UpdateProductLinkInvitationRequest) (*UpdateProductLinkInvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductLinkInvitation not implemented")
}
func (UnimplementedProductLinkInvitationServiceServer) RemoveProductLinkInvitation(context.Context, *RemoveProductLinkInvitationRequest) (*RemoveProductLinkInvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProductLinkInvitation not implemented")
}
func (UnimplementedProductLinkInvitationServiceServer) mustEmbedUnimplementedProductLinkInvitationServiceServer() {
}
func (UnimplementedProductLinkInvitationServiceServer) testEmbeddedByValue() {}

// UnsafeProductLinkInvitationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductLinkInvitationServiceServer will
// result in compilation errors.
type UnsafeProductLinkInvitationServiceServer interface {
	mustEmbedUnimplementedProductLinkInvitationServiceServer()
}

func RegisterProductLinkInvitationServiceServer(s grpc.ServiceRegistrar, srv ProductLinkInvitationServiceServer) {
	// If the following call pancis, it indicates UnimplementedProductLinkInvitationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProductLinkInvitationService_ServiceDesc, srv)
}

func _ProductLinkInvitationService_CreateProductLinkInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductLinkInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductLinkInvitationServiceServer).CreateProductLinkInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductLinkInvitationService_CreateProductLinkInvitation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductLinkInvitationServiceServer).CreateProductLinkInvitation(ctx, req.(*CreateProductLinkInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductLinkInvitationService_UpdateProductLinkInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductLinkInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductLinkInvitationServiceServer).UpdateProductLinkInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductLinkInvitationService_UpdateProductLinkInvitation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductLinkInvitationServiceServer).UpdateProductLinkInvitation(ctx, req.(*UpdateProductLinkInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductLinkInvitationService_RemoveProductLinkInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProductLinkInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductLinkInvitationServiceServer).RemoveProductLinkInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductLinkInvitationService_RemoveProductLinkInvitation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductLinkInvitationServiceServer).RemoveProductLinkInvitation(ctx, req.(*RemoveProductLinkInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductLinkInvitationService_ServiceDesc is the grpc.ServiceDesc for ProductLinkInvitationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductLinkInvitationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v18.services.ProductLinkInvitationService",
	HandlerType: (*ProductLinkInvitationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProductLinkInvitation",
			Handler:    _ProductLinkInvitationService_CreateProductLinkInvitation_Handler,
		},
		{
			MethodName: "UpdateProductLinkInvitation",
			Handler:    _ProductLinkInvitationService_UpdateProductLinkInvitation_Handler,
		},
		{
			MethodName: "RemoveProductLinkInvitation",
			Handler:    _ProductLinkInvitationService_RemoveProductLinkInvitation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v18/services/product_link_invitation_service.proto",
}
