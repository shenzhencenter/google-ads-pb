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
// source: google/ads/googleads/v19/services/reach_plan_service.proto

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
	ReachPlanService_GenerateConversionRates_FullMethodName = "/google.ads.googleads.v19.services.ReachPlanService/GenerateConversionRates"
	ReachPlanService_ListPlannableLocations_FullMethodName  = "/google.ads.googleads.v19.services.ReachPlanService/ListPlannableLocations"
	ReachPlanService_ListPlannableProducts_FullMethodName   = "/google.ads.googleads.v19.services.ReachPlanService/ListPlannableProducts"
	ReachPlanService_GenerateReachForecast_FullMethodName   = "/google.ads.googleads.v19.services.ReachPlanService/GenerateReachForecast"
)

// ReachPlanServiceClient is the client API for ReachPlanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Reach Plan Service gives users information about audience size that can
// be reached through advertisement on YouTube. In particular,
// GenerateReachForecast provides estimated number of people of specified
// demographics that can be reached by an ad in a given market by a campaign of
// certain duration with a defined budget.
type ReachPlanServiceClient interface {
	// Returns a collection of conversion rate suggestions for supported plannable
	// products.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GenerateConversionRates(ctx context.Context, in *GenerateConversionRatesRequest, opts ...grpc.CallOption) (*GenerateConversionRatesResponse, error)
	// Returns the list of plannable locations (for example, countries).
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ListPlannableLocations(ctx context.Context, in *ListPlannableLocationsRequest, opts ...grpc.CallOption) (*ListPlannableLocationsResponse, error)
	// Returns the list of per-location plannable YouTube ad formats with allowed
	// targeting.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ListPlannableProducts(ctx context.Context, in *ListPlannableProductsRequest, opts ...grpc.CallOption) (*ListPlannableProductsResponse, error)
	// Generates a reach forecast for a given targeting / product mix.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[ReachPlanError]()
	//	[RequestError]()
	GenerateReachForecast(ctx context.Context, in *GenerateReachForecastRequest, opts ...grpc.CallOption) (*GenerateReachForecastResponse, error)
}

type reachPlanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReachPlanServiceClient(cc grpc.ClientConnInterface) ReachPlanServiceClient {
	return &reachPlanServiceClient{cc}
}

func (c *reachPlanServiceClient) GenerateConversionRates(ctx context.Context, in *GenerateConversionRatesRequest, opts ...grpc.CallOption) (*GenerateConversionRatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateConversionRatesResponse)
	err := c.cc.Invoke(ctx, ReachPlanService_GenerateConversionRates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reachPlanServiceClient) ListPlannableLocations(ctx context.Context, in *ListPlannableLocationsRequest, opts ...grpc.CallOption) (*ListPlannableLocationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPlannableLocationsResponse)
	err := c.cc.Invoke(ctx, ReachPlanService_ListPlannableLocations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reachPlanServiceClient) ListPlannableProducts(ctx context.Context, in *ListPlannableProductsRequest, opts ...grpc.CallOption) (*ListPlannableProductsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPlannableProductsResponse)
	err := c.cc.Invoke(ctx, ReachPlanService_ListPlannableProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reachPlanServiceClient) GenerateReachForecast(ctx context.Context, in *GenerateReachForecastRequest, opts ...grpc.CallOption) (*GenerateReachForecastResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateReachForecastResponse)
	err := c.cc.Invoke(ctx, ReachPlanService_GenerateReachForecast_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReachPlanServiceServer is the server API for ReachPlanService service.
// All implementations must embed UnimplementedReachPlanServiceServer
// for forward compatibility.
//
// Reach Plan Service gives users information about audience size that can
// be reached through advertisement on YouTube. In particular,
// GenerateReachForecast provides estimated number of people of specified
// demographics that can be reached by an ad in a given market by a campaign of
// certain duration with a defined budget.
type ReachPlanServiceServer interface {
	// Returns a collection of conversion rate suggestions for supported plannable
	// products.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GenerateConversionRates(context.Context, *GenerateConversionRatesRequest) (*GenerateConversionRatesResponse, error)
	// Returns the list of plannable locations (for example, countries).
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ListPlannableLocations(context.Context, *ListPlannableLocationsRequest) (*ListPlannableLocationsResponse, error)
	// Returns the list of per-location plannable YouTube ad formats with allowed
	// targeting.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ListPlannableProducts(context.Context, *ListPlannableProductsRequest) (*ListPlannableProductsResponse, error)
	// Generates a reach forecast for a given targeting / product mix.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[ReachPlanError]()
	//	[RequestError]()
	GenerateReachForecast(context.Context, *GenerateReachForecastRequest) (*GenerateReachForecastResponse, error)
	mustEmbedUnimplementedReachPlanServiceServer()
}

// UnimplementedReachPlanServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReachPlanServiceServer struct{}

func (UnimplementedReachPlanServiceServer) GenerateConversionRates(context.Context, *GenerateConversionRatesRequest) (*GenerateConversionRatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateConversionRates not implemented")
}
func (UnimplementedReachPlanServiceServer) ListPlannableLocations(context.Context, *ListPlannableLocationsRequest) (*ListPlannableLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlannableLocations not implemented")
}
func (UnimplementedReachPlanServiceServer) ListPlannableProducts(context.Context, *ListPlannableProductsRequest) (*ListPlannableProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlannableProducts not implemented")
}
func (UnimplementedReachPlanServiceServer) GenerateReachForecast(context.Context, *GenerateReachForecastRequest) (*GenerateReachForecastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateReachForecast not implemented")
}
func (UnimplementedReachPlanServiceServer) mustEmbedUnimplementedReachPlanServiceServer() {}
func (UnimplementedReachPlanServiceServer) testEmbeddedByValue()                          {}

// UnsafeReachPlanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReachPlanServiceServer will
// result in compilation errors.
type UnsafeReachPlanServiceServer interface {
	mustEmbedUnimplementedReachPlanServiceServer()
}

func RegisterReachPlanServiceServer(s grpc.ServiceRegistrar, srv ReachPlanServiceServer) {
	// If the following call pancis, it indicates UnimplementedReachPlanServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReachPlanService_ServiceDesc, srv)
}

func _ReachPlanService_GenerateConversionRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateConversionRatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReachPlanServiceServer).GenerateConversionRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReachPlanService_GenerateConversionRates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReachPlanServiceServer).GenerateConversionRates(ctx, req.(*GenerateConversionRatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReachPlanService_ListPlannableLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPlannableLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReachPlanServiceServer).ListPlannableLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReachPlanService_ListPlannableLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReachPlanServiceServer).ListPlannableLocations(ctx, req.(*ListPlannableLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReachPlanService_ListPlannableProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPlannableProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReachPlanServiceServer).ListPlannableProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReachPlanService_ListPlannableProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReachPlanServiceServer).ListPlannableProducts(ctx, req.(*ListPlannableProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReachPlanService_GenerateReachForecast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateReachForecastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReachPlanServiceServer).GenerateReachForecast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReachPlanService_GenerateReachForecast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReachPlanServiceServer).GenerateReachForecast(ctx, req.(*GenerateReachForecastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReachPlanService_ServiceDesc is the grpc.ServiceDesc for ReachPlanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReachPlanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.ReachPlanService",
	HandlerType: (*ReachPlanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateConversionRates",
			Handler:    _ReachPlanService_GenerateConversionRates_Handler,
		},
		{
			MethodName: "ListPlannableLocations",
			Handler:    _ReachPlanService_ListPlannableLocations_Handler,
		},
		{
			MethodName: "ListPlannableProducts",
			Handler:    _ReachPlanService_ListPlannableProducts_Handler,
		},
		{
			MethodName: "GenerateReachForecast",
			Handler:    _ReachPlanService_GenerateReachForecast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/reach_plan_service.proto",
}
