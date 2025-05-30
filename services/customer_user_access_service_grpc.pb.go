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
// source: google/ads/googleads/v19/services/customer_user_access_service.proto

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
	CustomerUserAccessService_MutateCustomerUserAccess_FullMethodName = "/google.ads.googleads.v19.services.CustomerUserAccessService/MutateCustomerUserAccess"
)

// CustomerUserAccessServiceClient is the client API for CustomerUserAccessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// This service manages the permissions of a user on a given customer.
type CustomerUserAccessServiceClient interface {
	// Updates, removes permission of a user on a given customer. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CustomerUserAccessError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateCustomerUserAccess(ctx context.Context, in *MutateCustomerUserAccessRequest, opts ...grpc.CallOption) (*MutateCustomerUserAccessResponse, error)
}

type customerUserAccessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerUserAccessServiceClient(cc grpc.ClientConnInterface) CustomerUserAccessServiceClient {
	return &customerUserAccessServiceClient{cc}
}

func (c *customerUserAccessServiceClient) MutateCustomerUserAccess(ctx context.Context, in *MutateCustomerUserAccessRequest, opts ...grpc.CallOption) (*MutateCustomerUserAccessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateCustomerUserAccessResponse)
	err := c.cc.Invoke(ctx, CustomerUserAccessService_MutateCustomerUserAccess_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerUserAccessServiceServer is the server API for CustomerUserAccessService service.
// All implementations must embed UnimplementedCustomerUserAccessServiceServer
// for forward compatibility.
//
// This service manages the permissions of a user on a given customer.
type CustomerUserAccessServiceServer interface {
	// Updates, removes permission of a user on a given customer. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CustomerUserAccessError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateCustomerUserAccess(context.Context, *MutateCustomerUserAccessRequest) (*MutateCustomerUserAccessResponse, error)
	mustEmbedUnimplementedCustomerUserAccessServiceServer()
}

// UnimplementedCustomerUserAccessServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCustomerUserAccessServiceServer struct{}

func (UnimplementedCustomerUserAccessServiceServer) MutateCustomerUserAccess(context.Context, *MutateCustomerUserAccessRequest) (*MutateCustomerUserAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerUserAccess not implemented")
}
func (UnimplementedCustomerUserAccessServiceServer) mustEmbedUnimplementedCustomerUserAccessServiceServer() {
}
func (UnimplementedCustomerUserAccessServiceServer) testEmbeddedByValue() {}

// UnsafeCustomerUserAccessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerUserAccessServiceServer will
// result in compilation errors.
type UnsafeCustomerUserAccessServiceServer interface {
	mustEmbedUnimplementedCustomerUserAccessServiceServer()
}

func RegisterCustomerUserAccessServiceServer(s grpc.ServiceRegistrar, srv CustomerUserAccessServiceServer) {
	// If the following call pancis, it indicates UnimplementedCustomerUserAccessServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CustomerUserAccessService_ServiceDesc, srv)
}

func _CustomerUserAccessService_MutateCustomerUserAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerUserAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerUserAccessServiceServer).MutateCustomerUserAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerUserAccessService_MutateCustomerUserAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerUserAccessServiceServer).MutateCustomerUserAccess(ctx, req.(*MutateCustomerUserAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerUserAccessService_ServiceDesc is the grpc.ServiceDesc for CustomerUserAccessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerUserAccessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.CustomerUserAccessService",
	HandlerType: (*CustomerUserAccessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomerUserAccess",
			Handler:    _CustomerUserAccessService_MutateCustomerUserAccess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/customer_user_access_service.proto",
}
