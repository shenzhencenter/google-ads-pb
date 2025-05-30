// Copyright 2025 Google LLC
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
// source: google/ads/googleads/v19/services/customer_client_link_service.proto

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
	CustomerClientLinkService_MutateCustomerClientLink_FullMethodName = "/google.ads.googleads.v19.services.CustomerClientLinkService/MutateCustomerClientLink"
)

// CustomerClientLinkServiceClient is the client API for CustomerClientLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage customer client links.
type CustomerClientLinkServiceClient interface {
	// Creates or updates a customer client link. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[ManagerLinkError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateCustomerClientLink(ctx context.Context, in *MutateCustomerClientLinkRequest, opts ...grpc.CallOption) (*MutateCustomerClientLinkResponse, error)
}

type customerClientLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClientLinkServiceClient(cc grpc.ClientConnInterface) CustomerClientLinkServiceClient {
	return &customerClientLinkServiceClient{cc}
}

func (c *customerClientLinkServiceClient) MutateCustomerClientLink(ctx context.Context, in *MutateCustomerClientLinkRequest, opts ...grpc.CallOption) (*MutateCustomerClientLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateCustomerClientLinkResponse)
	err := c.cc.Invoke(ctx, CustomerClientLinkService_MutateCustomerClientLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerClientLinkServiceServer is the server API for CustomerClientLinkService service.
// All implementations must embed UnimplementedCustomerClientLinkServiceServer
// for forward compatibility.
//
// Service to manage customer client links.
type CustomerClientLinkServiceServer interface {
	// Creates or updates a customer client link. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[ManagerLinkError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateCustomerClientLink(context.Context, *MutateCustomerClientLinkRequest) (*MutateCustomerClientLinkResponse, error)
	mustEmbedUnimplementedCustomerClientLinkServiceServer()
}

// UnimplementedCustomerClientLinkServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCustomerClientLinkServiceServer struct{}

func (UnimplementedCustomerClientLinkServiceServer) MutateCustomerClientLink(context.Context, *MutateCustomerClientLinkRequest) (*MutateCustomerClientLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerClientLink not implemented")
}
func (UnimplementedCustomerClientLinkServiceServer) mustEmbedUnimplementedCustomerClientLinkServiceServer() {
}
func (UnimplementedCustomerClientLinkServiceServer) testEmbeddedByValue() {}

// UnsafeCustomerClientLinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerClientLinkServiceServer will
// result in compilation errors.
type UnsafeCustomerClientLinkServiceServer interface {
	mustEmbedUnimplementedCustomerClientLinkServiceServer()
}

func RegisterCustomerClientLinkServiceServer(s grpc.ServiceRegistrar, srv CustomerClientLinkServiceServer) {
	// If the following call pancis, it indicates UnimplementedCustomerClientLinkServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CustomerClientLinkService_ServiceDesc, srv)
}

func _CustomerClientLinkService_MutateCustomerClientLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerClientLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientLinkServiceServer).MutateCustomerClientLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerClientLinkService_MutateCustomerClientLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientLinkServiceServer).MutateCustomerClientLink(ctx, req.(*MutateCustomerClientLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerClientLinkService_ServiceDesc is the grpc.ServiceDesc for CustomerClientLinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerClientLinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.CustomerClientLinkService",
	HandlerType: (*CustomerClientLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomerClientLink",
			Handler:    _CustomerClientLinkService_MutateCustomerClientLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/customer_client_link_service.proto",
}
