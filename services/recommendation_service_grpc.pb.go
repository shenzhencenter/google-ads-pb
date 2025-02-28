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
// source: google/ads/googleads/v19/services/recommendation_service.proto

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
	RecommendationService_ApplyRecommendation_FullMethodName     = "/google.ads.googleads.v19.services.RecommendationService/ApplyRecommendation"
	RecommendationService_DismissRecommendation_FullMethodName   = "/google.ads.googleads.v19.services.RecommendationService/DismissRecommendation"
	RecommendationService_GenerateRecommendations_FullMethodName = "/google.ads.googleads.v19.services.RecommendationService/GenerateRecommendations"
)

// RecommendationServiceClient is the client API for RecommendationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage recommendations.
type RecommendationServiceClient interface {
	// Applies given recommendations with corresponding apply parameters.
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
	//	[RecommendationError]()
	//	[RequestError]()
	//	[UrlFieldError]()
	ApplyRecommendation(ctx context.Context, in *ApplyRecommendationRequest, opts ...grpc.CallOption) (*ApplyRecommendationResponse, error)
	// Dismisses given recommendations.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RecommendationError]()
	//	[RequestError]()
	DismissRecommendation(ctx context.Context, in *DismissRecommendationRequest, opts ...grpc.CallOption) (*DismissRecommendationResponse, error)
	// Generates Recommendations based off the requested recommendation_types.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RecommendationError]()
	//	[RequestError]()
	GenerateRecommendations(ctx context.Context, in *GenerateRecommendationsRequest, opts ...grpc.CallOption) (*GenerateRecommendationsResponse, error)
}

type recommendationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommendationServiceClient(cc grpc.ClientConnInterface) RecommendationServiceClient {
	return &recommendationServiceClient{cc}
}

func (c *recommendationServiceClient) ApplyRecommendation(ctx context.Context, in *ApplyRecommendationRequest, opts ...grpc.CallOption) (*ApplyRecommendationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApplyRecommendationResponse)
	err := c.cc.Invoke(ctx, RecommendationService_ApplyRecommendation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceClient) DismissRecommendation(ctx context.Context, in *DismissRecommendationRequest, opts ...grpc.CallOption) (*DismissRecommendationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DismissRecommendationResponse)
	err := c.cc.Invoke(ctx, RecommendationService_DismissRecommendation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recommendationServiceClient) GenerateRecommendations(ctx context.Context, in *GenerateRecommendationsRequest, opts ...grpc.CallOption) (*GenerateRecommendationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateRecommendationsResponse)
	err := c.cc.Invoke(ctx, RecommendationService_GenerateRecommendations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommendationServiceServer is the server API for RecommendationService service.
// All implementations must embed UnimplementedRecommendationServiceServer
// for forward compatibility.
//
// Service to manage recommendations.
type RecommendationServiceServer interface {
	// Applies given recommendations with corresponding apply parameters.
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
	//	[RecommendationError]()
	//	[RequestError]()
	//	[UrlFieldError]()
	ApplyRecommendation(context.Context, *ApplyRecommendationRequest) (*ApplyRecommendationResponse, error)
	// Dismisses given recommendations.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RecommendationError]()
	//	[RequestError]()
	DismissRecommendation(context.Context, *DismissRecommendationRequest) (*DismissRecommendationResponse, error)
	// Generates Recommendations based off the requested recommendation_types.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RecommendationError]()
	//	[RequestError]()
	GenerateRecommendations(context.Context, *GenerateRecommendationsRequest) (*GenerateRecommendationsResponse, error)
	mustEmbedUnimplementedRecommendationServiceServer()
}

// UnimplementedRecommendationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRecommendationServiceServer struct{}

func (UnimplementedRecommendationServiceServer) ApplyRecommendation(context.Context, *ApplyRecommendationRequest) (*ApplyRecommendationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyRecommendation not implemented")
}
func (UnimplementedRecommendationServiceServer) DismissRecommendation(context.Context, *DismissRecommendationRequest) (*DismissRecommendationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DismissRecommendation not implemented")
}
func (UnimplementedRecommendationServiceServer) GenerateRecommendations(context.Context, *GenerateRecommendationsRequest) (*GenerateRecommendationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateRecommendations not implemented")
}
func (UnimplementedRecommendationServiceServer) mustEmbedUnimplementedRecommendationServiceServer() {}
func (UnimplementedRecommendationServiceServer) testEmbeddedByValue()                               {}

// UnsafeRecommendationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecommendationServiceServer will
// result in compilation errors.
type UnsafeRecommendationServiceServer interface {
	mustEmbedUnimplementedRecommendationServiceServer()
}

func RegisterRecommendationServiceServer(s grpc.ServiceRegistrar, srv RecommendationServiceServer) {
	// If the following call pancis, it indicates UnimplementedRecommendationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RecommendationService_ServiceDesc, srv)
}

func _RecommendationService_ApplyRecommendation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRecommendationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceServer).ApplyRecommendation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecommendationService_ApplyRecommendation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceServer).ApplyRecommendation(ctx, req.(*ApplyRecommendationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationService_DismissRecommendation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DismissRecommendationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceServer).DismissRecommendation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecommendationService_DismissRecommendation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceServer).DismissRecommendation(ctx, req.(*DismissRecommendationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecommendationService_GenerateRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommendationServiceServer).GenerateRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecommendationService_GenerateRecommendations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommendationServiceServer).GenerateRecommendations(ctx, req.(*GenerateRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecommendationService_ServiceDesc is the grpc.ServiceDesc for RecommendationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecommendationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.RecommendationService",
	HandlerType: (*RecommendationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyRecommendation",
			Handler:    _RecommendationService_ApplyRecommendation_Handler,
		},
		{
			MethodName: "DismissRecommendation",
			Handler:    _RecommendationService_DismissRecommendation_Handler,
		},
		{
			MethodName: "GenerateRecommendations",
			Handler:    _RecommendationService_GenerateRecommendations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/recommendation_service.proto",
}
