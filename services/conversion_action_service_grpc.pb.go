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
// source: google/ads/googleads/v19/services/conversion_action_service.proto

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
	ConversionActionService_MutateConversionActions_FullMethodName = "/google.ads.googleads.v19.services.ConversionActionService/MutateConversionActions"
)

// ConversionActionServiceClient is the client API for ConversionActionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage conversion actions.
type ConversionActionServiceClient interface {
	// Creates, updates or removes conversion actions. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ConversionActionError]()
	//	[CurrencyCodeError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[StringLengthError]()
	MutateConversionActions(ctx context.Context, in *MutateConversionActionsRequest, opts ...grpc.CallOption) (*MutateConversionActionsResponse, error)
}

type conversionActionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversionActionServiceClient(cc grpc.ClientConnInterface) ConversionActionServiceClient {
	return &conversionActionServiceClient{cc}
}

func (c *conversionActionServiceClient) MutateConversionActions(ctx context.Context, in *MutateConversionActionsRequest, opts ...grpc.CallOption) (*MutateConversionActionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateConversionActionsResponse)
	err := c.cc.Invoke(ctx, ConversionActionService_MutateConversionActions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionActionServiceServer is the server API for ConversionActionService service.
// All implementations must embed UnimplementedConversionActionServiceServer
// for forward compatibility.
//
// Service to manage conversion actions.
type ConversionActionServiceServer interface {
	// Creates, updates or removes conversion actions. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ConversionActionError]()
	//	[CurrencyCodeError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[StringLengthError]()
	MutateConversionActions(context.Context, *MutateConversionActionsRequest) (*MutateConversionActionsResponse, error)
	mustEmbedUnimplementedConversionActionServiceServer()
}

// UnimplementedConversionActionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConversionActionServiceServer struct{}

func (UnimplementedConversionActionServiceServer) MutateConversionActions(context.Context, *MutateConversionActionsRequest) (*MutateConversionActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateConversionActions not implemented")
}
func (UnimplementedConversionActionServiceServer) mustEmbedUnimplementedConversionActionServiceServer() {
}
func (UnimplementedConversionActionServiceServer) testEmbeddedByValue() {}

// UnsafeConversionActionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversionActionServiceServer will
// result in compilation errors.
type UnsafeConversionActionServiceServer interface {
	mustEmbedUnimplementedConversionActionServiceServer()
}

func RegisterConversionActionServiceServer(s grpc.ServiceRegistrar, srv ConversionActionServiceServer) {
	// If the following call pancis, it indicates UnimplementedConversionActionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ConversionActionService_ServiceDesc, srv)
}

func _ConversionActionService_MutateConversionActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateConversionActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionActionServiceServer).MutateConversionActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConversionActionService_MutateConversionActions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionActionServiceServer).MutateConversionActions(ctx, req.(*MutateConversionActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConversionActionService_ServiceDesc is the grpc.ServiceDesc for ConversionActionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConversionActionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.ConversionActionService",
	HandlerType: (*ConversionActionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateConversionActions",
			Handler:    _ConversionActionService_MutateConversionActions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/conversion_action_service.proto",
}
