// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AudienceInsightsServiceClient is the client API for AudienceInsightsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AudienceInsightsServiceClient interface {
	// Creates a saved report that can be viewed in the Insights Finder tool.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	GenerateInsightsFinderReport(ctx context.Context, in *GenerateInsightsFinderReportRequest, opts ...grpc.CallOption) (*GenerateInsightsFinderReportResponse, error)
	// Searches for audience attributes that can be used to generate insights.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	ListAudienceInsightsAttributes(ctx context.Context, in *ListAudienceInsightsAttributesRequest, opts ...grpc.CallOption) (*ListAudienceInsightsAttributesResponse, error)
}

type audienceInsightsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAudienceInsightsServiceClient(cc grpc.ClientConnInterface) AudienceInsightsServiceClient {
	return &audienceInsightsServiceClient{cc}
}

func (c *audienceInsightsServiceClient) GenerateInsightsFinderReport(ctx context.Context, in *GenerateInsightsFinderReportRequest, opts ...grpc.CallOption) (*GenerateInsightsFinderReportResponse, error) {
	out := new(GenerateInsightsFinderReportResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.AudienceInsightsService/GenerateInsightsFinderReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *audienceInsightsServiceClient) ListAudienceInsightsAttributes(ctx context.Context, in *ListAudienceInsightsAttributesRequest, opts ...grpc.CallOption) (*ListAudienceInsightsAttributesResponse, error) {
	out := new(ListAudienceInsightsAttributesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v11.services.AudienceInsightsService/ListAudienceInsightsAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AudienceInsightsServiceServer is the server API for AudienceInsightsService service.
// All implementations must embed UnimplementedAudienceInsightsServiceServer
// for forward compatibility
type AudienceInsightsServiceServer interface {
	// Creates a saved report that can be viewed in the Insights Finder tool.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	GenerateInsightsFinderReport(context.Context, *GenerateInsightsFinderReportRequest) (*GenerateInsightsFinderReportResponse, error)
	// Searches for audience attributes that can be used to generate insights.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	ListAudienceInsightsAttributes(context.Context, *ListAudienceInsightsAttributesRequest) (*ListAudienceInsightsAttributesResponse, error)
	mustEmbedUnimplementedAudienceInsightsServiceServer()
}

// UnimplementedAudienceInsightsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAudienceInsightsServiceServer struct {
}

func (UnimplementedAudienceInsightsServiceServer) GenerateInsightsFinderReport(context.Context, *GenerateInsightsFinderReportRequest) (*GenerateInsightsFinderReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateInsightsFinderReport not implemented")
}
func (UnimplementedAudienceInsightsServiceServer) ListAudienceInsightsAttributes(context.Context, *ListAudienceInsightsAttributesRequest) (*ListAudienceInsightsAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAudienceInsightsAttributes not implemented")
}
func (UnimplementedAudienceInsightsServiceServer) mustEmbedUnimplementedAudienceInsightsServiceServer() {
}

// UnsafeAudienceInsightsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AudienceInsightsServiceServer will
// result in compilation errors.
type UnsafeAudienceInsightsServiceServer interface {
	mustEmbedUnimplementedAudienceInsightsServiceServer()
}

func RegisterAudienceInsightsServiceServer(s grpc.ServiceRegistrar, srv AudienceInsightsServiceServer) {
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
		FullMethod: "/google.ads.googleads.v11.services.AudienceInsightsService/GenerateInsightsFinderReport",
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
		FullMethod: "/google.ads.googleads.v11.services.AudienceInsightsService/ListAudienceInsightsAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AudienceInsightsServiceServer).ListAudienceInsightsAttributes(ctx, req.(*ListAudienceInsightsAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AudienceInsightsService_ServiceDesc is the grpc.ServiceDesc for AudienceInsightsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AudienceInsightsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v11.services.AudienceInsightsService",
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v11/services/audience_insights_service.proto",
}