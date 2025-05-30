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
// source: google/ads/googleads/v19/services/billing_setup_service.proto

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
	BillingSetupService_MutateBillingSetup_FullMethodName = "/google.ads.googleads.v19.services.BillingSetupService/MutateBillingSetup"
)

// BillingSetupServiceClient is the client API for BillingSetupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A service for designating the business entity responsible for accrued costs.
//
// A billing setup is associated with a payments account.  Billing-related
// activity for all billing setups associated with a particular payments account
// will appear on a single invoice generated monthly.
//
// Mutates:
// The REMOVE operation cancels a pending billing setup.
// The CREATE operation creates a new billing setup.
type BillingSetupServiceClient interface {
	// Creates a billing setup, or cancels an existing billing setup.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BillingSetupError]()
	//	[DateError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateBillingSetup(ctx context.Context, in *MutateBillingSetupRequest, opts ...grpc.CallOption) (*MutateBillingSetupResponse, error)
}

type billingSetupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingSetupServiceClient(cc grpc.ClientConnInterface) BillingSetupServiceClient {
	return &billingSetupServiceClient{cc}
}

func (c *billingSetupServiceClient) MutateBillingSetup(ctx context.Context, in *MutateBillingSetupRequest, opts ...grpc.CallOption) (*MutateBillingSetupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateBillingSetupResponse)
	err := c.cc.Invoke(ctx, BillingSetupService_MutateBillingSetup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingSetupServiceServer is the server API for BillingSetupService service.
// All implementations must embed UnimplementedBillingSetupServiceServer
// for forward compatibility.
//
// A service for designating the business entity responsible for accrued costs.
//
// A billing setup is associated with a payments account.  Billing-related
// activity for all billing setups associated with a particular payments account
// will appear on a single invoice generated monthly.
//
// Mutates:
// The REMOVE operation cancels a pending billing setup.
// The CREATE operation creates a new billing setup.
type BillingSetupServiceServer interface {
	// Creates a billing setup, or cancels an existing billing setup.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BillingSetupError]()
	//	[DateError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateBillingSetup(context.Context, *MutateBillingSetupRequest) (*MutateBillingSetupResponse, error)
	mustEmbedUnimplementedBillingSetupServiceServer()
}

// UnimplementedBillingSetupServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBillingSetupServiceServer struct{}

func (UnimplementedBillingSetupServiceServer) MutateBillingSetup(context.Context, *MutateBillingSetupRequest) (*MutateBillingSetupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateBillingSetup not implemented")
}
func (UnimplementedBillingSetupServiceServer) mustEmbedUnimplementedBillingSetupServiceServer() {}
func (UnimplementedBillingSetupServiceServer) testEmbeddedByValue()                             {}

// UnsafeBillingSetupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillingSetupServiceServer will
// result in compilation errors.
type UnsafeBillingSetupServiceServer interface {
	mustEmbedUnimplementedBillingSetupServiceServer()
}

func RegisterBillingSetupServiceServer(s grpc.ServiceRegistrar, srv BillingSetupServiceServer) {
	// If the following call pancis, it indicates UnimplementedBillingSetupServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BillingSetupService_ServiceDesc, srv)
}

func _BillingSetupService_MutateBillingSetup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBillingSetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingSetupServiceServer).MutateBillingSetup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingSetupService_MutateBillingSetup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingSetupServiceServer).MutateBillingSetup(ctx, req.(*MutateBillingSetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BillingSetupService_ServiceDesc is the grpc.ServiceDesc for BillingSetupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BillingSetupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.BillingSetupService",
	HandlerType: (*BillingSetupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateBillingSetup",
			Handler:    _BillingSetupService_MutateBillingSetup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/billing_setup_service.proto",
}
