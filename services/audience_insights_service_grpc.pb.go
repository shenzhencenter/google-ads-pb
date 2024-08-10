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
// source: google/ads/googleads/v17/services/audience_insights_service.proto

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
	AudienceInsightsService_GenerateInsightsFinderReport_FullMethodName        = "/google.ads.googleads.v17.services.AudienceInsightsService/GenerateInsightsFinderReport"
	AudienceInsightsService_ListAudienceInsightsAttributes_FullMethodName      = "/google.ads.googleads.v17.services.AudienceInsightsService/ListAudienceInsightsAttributes"
	AudienceInsightsService_ListInsightsEligibleDates_FullMethodName           = "/google.ads.googleads.v17.services.AudienceInsightsService/ListInsightsEligibleDates"
	AudienceInsightsService_GenerateAudienceCompositionInsights_FullMethodName = "/google.ads.googleads.v17.services.AudienceInsightsService/GenerateAudienceCompositionInsights"
	AudienceInsightsService_GenerateSuggestedTargetingInsights_FullMethodName  = "/google.ads.googleads.v17.services.AudienceInsightsService/GenerateSuggestedTargetingInsights"
	AudienceInsightsService_GenerateAudienceOverlapInsights_FullMethodName     = "/google.ads.googleads.v17.services.AudienceInsightsService/GenerateAudienceOverlapInsights"
)

// AudienceInsightsServiceClient is the client API for AudienceInsightsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Audience Insights Service helps users find information about groups of
// people and how they can be reached with Google Ads. Accessible to
// allowlisted customers only.
type AudienceInsightsServiceClient interface {
	// Creates a saved report that can be viewed in the Insights Finder tool.
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
	//	[RequestError]()
	GenerateInsightsFinderReport(ctx context.Context, in *GenerateInsightsFinderReportRequest, opts ...grpc.CallOption) (*GenerateInsightsFinderReportResponse, error)
	// Searches for audience attributes that can be used to generate insights.
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
	//	[RequestError]()
	ListAudienceInsightsAttributes(ctx context.Context, in *ListAudienceInsightsAttributesRequest, opts ...grpc.CallOption) (*ListAudienceInsightsAttributesResponse, error)
	// Lists date ranges for which audience insights data can be requested.
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
	//	[RequestError]()
	ListInsightsEligibleDates(ctx context.Context, in *ListInsightsEligibleDatesRequest, opts ...grpc.CallOption) (*ListInsightsEligibleDatesResponse, error)
	// Returns a collection of attributes that are represented in an audience of
	// interest, with metrics that compare each attribute's share of the audience
	// with its share of a baseline audience.
	//
	// List of thrown errors:
	//
	//	[AudienceInsightsError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	GenerateAudienceCompositionInsights(ctx context.Context, in *GenerateAudienceCompositionInsightsRequest, opts ...grpc.CallOption) (*GenerateAudienceCompositionInsightsResponse, error)
	// Returns a collection of targeting insights (e.g. targetable audiences) that
	// are relevant to the requested audience.
	//
	// List of thrown errors:
	//
	//	[AudienceInsightsError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	GenerateSuggestedTargetingInsights(ctx context.Context, in *GenerateSuggestedTargetingInsightsRequest, opts ...grpc.CallOption) (*GenerateSuggestedTargetingInsightsResponse, error)
	// Returns a collection of audience attributes along with estimates of the
	// overlap between their potential YouTube reach and that of a given input
	// attribute.
	//
	// List of thrown errors:
	//
	//	[AudienceInsightsError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	GenerateAudienceOverlapInsights(ctx context.Context, in *GenerateAudienceOverlapInsightsRequest, opts ...grpc.CallOption) (*GenerateAudienceOverlapInsightsResponse, error)
}

type audienceInsightsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAudienceInsightsServiceClient(cc grpc.ClientConnInterface) AudienceInsightsServiceClient {
	return &audienceInsightsServiceClient{cc}
}

func (c *audienceInsightsServiceClient) GenerateInsightsFinderReport(ctx context.Context, in *GenerateInsightsFinderReportRequest, opts ...grpc.CallOption) (*GenerateInsightsFinderReportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateInsightsFinderReportResponse)
	err := c.cc.Invoke(ctx, AudienceInsightsService_GenerateInsightsFinderReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceInsightsServiceClient) ListAudienceInsightsAttributes(ctx context.Context, in *ListAudienceInsightsAttributesRequest, opts ...grpc.CallOption) (*ListAudienceInsightsAttributesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAudienceInsightsAttributesResponse)
	err := c.cc.Invoke(ctx, AudienceInsightsService_ListAudienceInsightsAttributes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceInsightsServiceClient) ListInsightsEligibleDates(ctx context.Context, in *ListInsightsEligibleDatesRequest, opts ...grpc.CallOption) (*ListInsightsEligibleDatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListInsightsEligibleDatesResponse)
	err := c.cc.Invoke(ctx, AudienceInsightsService_ListInsightsEligibleDates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceInsightsServiceClient) GenerateAudienceCompositionInsights(ctx context.Context, in *GenerateAudienceCompositionInsightsRequest, opts ...grpc.CallOption) (*GenerateAudienceCompositionInsightsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateAudienceCompositionInsightsResponse)
	err := c.cc.Invoke(ctx, AudienceInsightsService_GenerateAudienceCompositionInsights_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceInsightsServiceClient) GenerateSuggestedTargetingInsights(ctx context.Context, in *GenerateSuggestedTargetingInsightsRequest, opts ...grpc.CallOption) (*GenerateSuggestedTargetingInsightsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateSuggestedTargetingInsightsResponse)
	err := c.cc.Invoke(ctx, AudienceInsightsService_GenerateSuggestedTargetingInsights_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceInsightsServiceClient) GenerateAudienceOverlapInsights(ctx context.Context, in *GenerateAudienceOverlapInsightsRequest, opts ...grpc.CallOption) (*GenerateAudienceOverlapInsightsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateAudienceOverlapInsightsResponse)
	err := c.cc.Invoke(ctx, AudienceInsightsService_GenerateAudienceOverlapInsights_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudienceInsightsServiceServer is the server API for AudienceInsightsService service.
// All implementations must embed UnimplementedAudienceInsightsServiceServer
// for forward compatibility.
//
// Audience Insights Service helps users find information about groups of
// people and how they can be reached with Google Ads. Accessible to
// allowlisted customers only.
type AudienceInsightsServiceServer interface {
	// Creates a saved report that can be viewed in the Insights Finder tool.
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
	//	[RequestError]()
	GenerateInsightsFinderReport(context.Context, *GenerateInsightsFinderReportRequest) (*GenerateInsightsFinderReportResponse, error)
	// Searches for audience attributes that can be used to generate insights.
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
	//	[RequestError]()
	ListAudienceInsightsAttributes(context.Context, *ListAudienceInsightsAttributesRequest) (*ListAudienceInsightsAttributesResponse, error)
	// Lists date ranges for which audience insights data can be requested.
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
	//	[RequestError]()
	ListInsightsEligibleDates(context.Context, *ListInsightsEligibleDatesRequest) (*ListInsightsEligibleDatesResponse, error)
	// Returns a collection of attributes that are represented in an audience of
	// interest, with metrics that compare each attribute's share of the audience
	// with its share of a baseline audience.
	//
	// List of thrown errors:
	//
	//	[AudienceInsightsError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	GenerateAudienceCompositionInsights(context.Context, *GenerateAudienceCompositionInsightsRequest) (*GenerateAudienceCompositionInsightsResponse, error)
	// Returns a collection of targeting insights (e.g. targetable audiences) that
	// are relevant to the requested audience.
	//
	// List of thrown errors:
	//
	//	[AudienceInsightsError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	GenerateSuggestedTargetingInsights(context.Context, *GenerateSuggestedTargetingInsightsRequest) (*GenerateSuggestedTargetingInsightsResponse, error)
	// Returns a collection of audience attributes along with estimates of the
	// overlap between their potential YouTube reach and that of a given input
	// attribute.
	//
	// List of thrown errors:
	//
	//	[AudienceInsightsError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	GenerateAudienceOverlapInsights(context.Context, *GenerateAudienceOverlapInsightsRequest) (*GenerateAudienceOverlapInsightsResponse, error)
	mustEmbedUnimplementedAudienceInsightsServiceServer()
}

// UnimplementedAudienceInsightsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAudienceInsightsServiceServer struct{}

func (UnimplementedAudienceInsightsServiceServer) GenerateInsightsFinderReport(context.Context, *GenerateInsightsFinderReportRequest) (*GenerateInsightsFinderReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateInsightsFinderReport not implemented")
}
func (UnimplementedAudienceInsightsServiceServer) ListAudienceInsightsAttributes(context.Context, *ListAudienceInsightsAttributesRequest) (*ListAudienceInsightsAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAudienceInsightsAttributes not implemented")
}
func (UnimplementedAudienceInsightsServiceServer) ListInsightsEligibleDates(context.Context, *ListInsightsEligibleDatesRequest) (*ListInsightsEligibleDatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInsightsEligibleDates not implemented")
}
func (UnimplementedAudienceInsightsServiceServer) GenerateAudienceCompositionInsights(context.Context, *GenerateAudienceCompositionInsightsRequest) (*GenerateAudienceCompositionInsightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAudienceCompositionInsights not implemented")
}
func (UnimplementedAudienceInsightsServiceServer) GenerateSuggestedTargetingInsights(context.Context, *GenerateSuggestedTargetingInsightsRequest) (*GenerateSuggestedTargetingInsightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateSuggestedTargetingInsights not implemented")
}
func (UnimplementedAudienceInsightsServiceServer) GenerateAudienceOverlapInsights(context.Context, *GenerateAudienceOverlapInsightsRequest) (*GenerateAudienceOverlapInsightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateAudienceOverlapInsights not implemented")
}
func (UnimplementedAudienceInsightsServiceServer) mustEmbedUnimplementedAudienceInsightsServiceServer() {
}
func (UnimplementedAudienceInsightsServiceServer) testEmbeddedByValue() {}

// UnsafeAudienceInsightsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudienceInsightsServiceServer will
// result in compilation errors.
type UnsafeAudienceInsightsServiceServer interface {
	mustEmbedUnimplementedAudienceInsightsServiceServer()
}

func RegisterAudienceInsightsServiceServer(s grpc.ServiceRegistrar, srv AudienceInsightsServiceServer) {
	// If the following call pancis, it indicates UnimplementedAudienceInsightsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AudienceInsightsService_ServiceDesc, srv)
}

func _AudienceInsightsService_GenerateInsightsFinderReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateInsightsFinderReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceInsightsServiceServer).GenerateInsightsFinderReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudienceInsightsService_GenerateInsightsFinderReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceInsightsServiceServer).GenerateInsightsFinderReport(ctx, req.(*GenerateInsightsFinderReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudienceInsightsService_ListAudienceInsightsAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAudienceInsightsAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceInsightsServiceServer).ListAudienceInsightsAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudienceInsightsService_ListAudienceInsightsAttributes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceInsightsServiceServer).ListAudienceInsightsAttributes(ctx, req.(*ListAudienceInsightsAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudienceInsightsService_ListInsightsEligibleDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInsightsEligibleDatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceInsightsServiceServer).ListInsightsEligibleDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudienceInsightsService_ListInsightsEligibleDates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceInsightsServiceServer).ListInsightsEligibleDates(ctx, req.(*ListInsightsEligibleDatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudienceInsightsService_GenerateAudienceCompositionInsights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAudienceCompositionInsightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceInsightsServiceServer).GenerateAudienceCompositionInsights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudienceInsightsService_GenerateAudienceCompositionInsights_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceInsightsServiceServer).GenerateAudienceCompositionInsights(ctx, req.(*GenerateAudienceCompositionInsightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudienceInsightsService_GenerateSuggestedTargetingInsights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateSuggestedTargetingInsightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceInsightsServiceServer).GenerateSuggestedTargetingInsights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudienceInsightsService_GenerateSuggestedTargetingInsights_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceInsightsServiceServer).GenerateSuggestedTargetingInsights(ctx, req.(*GenerateSuggestedTargetingInsightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AudienceInsightsService_GenerateAudienceOverlapInsights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAudienceOverlapInsightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AudienceInsightsServiceServer).GenerateAudienceOverlapInsights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AudienceInsightsService_GenerateAudienceOverlapInsights_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceInsightsServiceServer).GenerateAudienceOverlapInsights(ctx, req.(*GenerateAudienceOverlapInsightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AudienceInsightsService_ServiceDesc is the grpc.ServiceDesc for AudienceInsightsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudienceInsightsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.AudienceInsightsService",
	HandlerType: (*AudienceInsightsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateInsightsFinderReport",
			Handler:    _AudienceInsightsService_GenerateInsightsFinderReport_Handler,
		},
		{
			MethodName: "ListAudienceInsightsAttributes",
			Handler:    _AudienceInsightsService_ListAudienceInsightsAttributes_Handler,
		},
		{
			MethodName: "ListInsightsEligibleDates",
			Handler:    _AudienceInsightsService_ListInsightsEligibleDates_Handler,
		},
		{
			MethodName: "GenerateAudienceCompositionInsights",
			Handler:    _AudienceInsightsService_GenerateAudienceCompositionInsights_Handler,
		},
		{
			MethodName: "GenerateSuggestedTargetingInsights",
			Handler:    _AudienceInsightsService_GenerateSuggestedTargetingInsights_Handler,
		},
		{
			MethodName: "GenerateAudienceOverlapInsights",
			Handler:    _AudienceInsightsService_GenerateAudienceOverlapInsights_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/audience_insights_service.proto",
}
