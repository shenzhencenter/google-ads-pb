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
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.24.4
// source: google/ads/googleads/v17/services/customer_service.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CustomerService_MutateCustomer_FullMethodName          = "/google.ads.googleads.v17.services.CustomerService/MutateCustomer"
	CustomerService_ListAccessibleCustomers_FullMethodName = "/google.ads.googleads.v17.services.CustomerService/ListAccessibleCustomers"
	CustomerService_CreateCustomerClient_FullMethodName    = "/google.ads.googleads.v17.services.CustomerService/CreateCustomerClient"
)

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage customers.
type CustomerServiceClient interface {
	// Updates a customer. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[UrlFieldError]()
	MutateCustomer(ctx context.Context, in *MutateCustomerRequest, opts ...grpc.CallOption) (*MutateCustomerResponse, error)
	// Returns resource names of customers directly accessible by the
	// user authenticating the call.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ListAccessibleCustomers(ctx context.Context, in *ListAccessibleCustomersRequest, opts ...grpc.CallOption) (*ListAccessibleCustomersResponse, error)
	// Creates a new client under manager. The new client customer is returned.
	//
	// List of thrown errors:
	//
	//	[AccessInvitationError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CurrencyCodeError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[ManagerLinkError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[StringLengthError]()
	//	[TimeZoneError]()
	CreateCustomerClient(ctx context.Context, in *CreateCustomerClientRequest, opts ...grpc.CallOption) (*CreateCustomerClientResponse, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) MutateCustomer(ctx context.Context, in *MutateCustomerRequest, opts ...grpc.CallOption) (*MutateCustomerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateCustomerResponse)
	err := c.cc.Invoke(ctx, CustomerService_MutateCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) ListAccessibleCustomers(ctx context.Context, in *ListAccessibleCustomersRequest, opts ...grpc.CallOption) (*ListAccessibleCustomersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAccessibleCustomersResponse)
	err := c.cc.Invoke(ctx, CustomerService_ListAccessibleCustomers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) CreateCustomerClient(ctx context.Context, in *CreateCustomerClientRequest, opts ...grpc.CallOption) (*CreateCustomerClientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCustomerClientResponse)
	err := c.cc.Invoke(ctx, CustomerService_CreateCustomerClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations must embed UnimplementedCustomerServiceServer
// for forward compatibility
//
// Service to manage customers.
type CustomerServiceServer interface {
	// Updates a customer. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[UrlFieldError]()
	MutateCustomer(context.Context, *MutateCustomerRequest) (*MutateCustomerResponse, error)
	// Returns resource names of customers directly accessible by the
	// user authenticating the call.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ListAccessibleCustomers(context.Context, *ListAccessibleCustomersRequest) (*ListAccessibleCustomersResponse, error)
	// Creates a new client under manager. The new client customer is returned.
	//
	// List of thrown errors:
	//
	//	[AccessInvitationError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CurrencyCodeError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[ManagerLinkError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[StringLengthError]()
	//	[TimeZoneError]()
	CreateCustomerClient(context.Context, *CreateCustomerClientRequest) (*CreateCustomerClientResponse, error)
	mustEmbedUnimplementedCustomerServiceServer()
}

// UnimplementedCustomerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (UnimplementedCustomerServiceServer) MutateCustomer(context.Context, *MutateCustomerRequest) (*MutateCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) ListAccessibleCustomers(context.Context, *ListAccessibleCustomersRequest) (*ListAccessibleCustomersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessibleCustomers not implemented")
}
func (UnimplementedCustomerServiceServer) CreateCustomerClient(context.Context, *CreateCustomerClientRequest) (*CreateCustomerClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomerClient not implemented")
}
func (UnimplementedCustomerServiceServer) mustEmbedUnimplementedCustomerServiceServer() {}

// UnsafeCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceServer will
// result in compilation errors.
type UnsafeCustomerServiceServer interface {
	mustEmbedUnimplementedCustomerServiceServer()
}

func RegisterCustomerServiceServer(s grpc.ServiceRegistrar, srv CustomerServiceServer) {
	s.RegisterService(&CustomerService_ServiceDesc, srv)
}

func _CustomerService_MutateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).MutateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_MutateCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).MutateCustomer(ctx, req.(*MutateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_ListAccessibleCustomers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccessibleCustomersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).ListAccessibleCustomers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_ListAccessibleCustomers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).ListAccessibleCustomers(ctx, req.(*ListAccessibleCustomersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_CreateCustomerClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).CreateCustomerClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_CreateCustomerClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).CreateCustomerClient(ctx, req.(*CreateCustomerClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomer",
			Handler:    _CustomerService_MutateCustomer_Handler,
		},
		{
			MethodName: "ListAccessibleCustomers",
			Handler:    _CustomerService_ListAccessibleCustomers_Handler,
		},
		{
			MethodName: "CreateCustomerClient",
			Handler:    _CustomerService_CreateCustomerClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/customer_service.proto",
}
