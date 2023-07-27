// Copyright 2023 Google LLC
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
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: google/ads/googleads/v14/services/geo_target_constant_service.proto

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

const (
	GeoTargetConstantService_SuggestGeoTargetConstants_FullMethodName = "/google.ads.googleads.v14.services.GeoTargetConstantService/SuggestGeoTargetConstants"
)

// GeoTargetConstantServiceClient is the client API for GeoTargetConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeoTargetConstantServiceClient interface {
	// Returns GeoTargetConstant suggestions by location name or by resource name.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[GeoTargetConstantSuggestionError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	SuggestGeoTargetConstants(ctx context.Context, in *SuggestGeoTargetConstantsRequest, opts ...grpc.CallOption) (*SuggestGeoTargetConstantsResponse, error)
}

type geoTargetConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoTargetConstantServiceClient(cc grpc.ClientConnInterface) GeoTargetConstantServiceClient {
	return &geoTargetConstantServiceClient{cc}
}

func (c *geoTargetConstantServiceClient) SuggestGeoTargetConstants(ctx context.Context, in *SuggestGeoTargetConstantsRequest, opts ...grpc.CallOption) (*SuggestGeoTargetConstantsResponse, error) {
	out := new(SuggestGeoTargetConstantsResponse)
	err := c.cc.Invoke(ctx, GeoTargetConstantService_SuggestGeoTargetConstants_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoTargetConstantServiceServer is the server API for GeoTargetConstantService service.
// All implementations must embed UnimplementedGeoTargetConstantServiceServer
// for forward compatibility
type GeoTargetConstantServiceServer interface {
	// Returns GeoTargetConstant suggestions by location name or by resource name.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[GeoTargetConstantSuggestionError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	SuggestGeoTargetConstants(context.Context, *SuggestGeoTargetConstantsRequest) (*SuggestGeoTargetConstantsResponse, error)
	mustEmbedUnimplementedGeoTargetConstantServiceServer()
}

// UnimplementedGeoTargetConstantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGeoTargetConstantServiceServer struct {
}

func (UnimplementedGeoTargetConstantServiceServer) SuggestGeoTargetConstants(context.Context, *SuggestGeoTargetConstantsRequest) (*SuggestGeoTargetConstantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestGeoTargetConstants not implemented")
}
func (UnimplementedGeoTargetConstantServiceServer) mustEmbedUnimplementedGeoTargetConstantServiceServer() {
}

// UnsafeGeoTargetConstantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeoTargetConstantServiceServer will
// result in compilation errors.
type UnsafeGeoTargetConstantServiceServer interface {
	mustEmbedUnimplementedGeoTargetConstantServiceServer()
}

func RegisterGeoTargetConstantServiceServer(s grpc.ServiceRegistrar, srv GeoTargetConstantServiceServer) {
	s.RegisterService(&GeoTargetConstantService_ServiceDesc, srv)
}

func _GeoTargetConstantService_SuggestGeoTargetConstants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestGeoTargetConstantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoTargetConstantServiceServer).SuggestGeoTargetConstants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GeoTargetConstantService_SuggestGeoTargetConstants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoTargetConstantServiceServer).SuggestGeoTargetConstants(ctx, req.(*SuggestGeoTargetConstantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GeoTargetConstantService_ServiceDesc is the grpc.ServiceDesc for GeoTargetConstantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeoTargetConstantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v14.services.GeoTargetConstantService",
	HandlerType: (*GeoTargetConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuggestGeoTargetConstants",
			Handler:    _GeoTargetConstantService_SuggestGeoTargetConstants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v14/services/geo_target_constant_service.proto",
}
