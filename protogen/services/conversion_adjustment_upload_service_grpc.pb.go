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
// - protoc             v3.21.1
// source: google/ads/googleads/v15/services/conversion_adjustment_upload_service.proto

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
	ConversionAdjustmentUploadService_UploadConversionAdjustments_FullMethodName = "/google.ads.googleads.v15.services.ConversionAdjustmentUploadService/UploadConversionAdjustments"
)

// ConversionAdjustmentUploadServiceClient is the client API for ConversionAdjustmentUploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConversionAdjustmentUploadServiceClient interface {
	// Processes the given conversion adjustments.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[PartialFailureError]()
	//	[QuotaError]()
	//	[RequestError]()
	UploadConversionAdjustments(ctx context.Context, in *UploadConversionAdjustmentsRequest, opts ...grpc.CallOption) (*UploadConversionAdjustmentsResponse, error)
}

type conversionAdjustmentUploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversionAdjustmentUploadServiceClient(cc grpc.ClientConnInterface) ConversionAdjustmentUploadServiceClient {
	return &conversionAdjustmentUploadServiceClient{cc}
}

func (c *conversionAdjustmentUploadServiceClient) UploadConversionAdjustments(ctx context.Context, in *UploadConversionAdjustmentsRequest, opts ...grpc.CallOption) (*UploadConversionAdjustmentsResponse, error) {
	out := new(UploadConversionAdjustmentsResponse)
	err := c.cc.Invoke(ctx, ConversionAdjustmentUploadService_UploadConversionAdjustments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionAdjustmentUploadServiceServer is the server API for ConversionAdjustmentUploadService service.
// All implementations must embed UnimplementedConversionAdjustmentUploadServiceServer
// for forward compatibility
type ConversionAdjustmentUploadServiceServer interface {
	// Processes the given conversion adjustments.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[PartialFailureError]()
	//	[QuotaError]()
	//	[RequestError]()
	UploadConversionAdjustments(context.Context, *UploadConversionAdjustmentsRequest) (*UploadConversionAdjustmentsResponse, error)
	mustEmbedUnimplementedConversionAdjustmentUploadServiceServer()
}

// UnimplementedConversionAdjustmentUploadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConversionAdjustmentUploadServiceServer struct {
}

func (UnimplementedConversionAdjustmentUploadServiceServer) UploadConversionAdjustments(context.Context, *UploadConversionAdjustmentsRequest) (*UploadConversionAdjustmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadConversionAdjustments not implemented")
}
func (UnimplementedConversionAdjustmentUploadServiceServer) mustEmbedUnimplementedConversionAdjustmentUploadServiceServer() {
}

// UnsafeConversionAdjustmentUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversionAdjustmentUploadServiceServer will
// result in compilation errors.
type UnsafeConversionAdjustmentUploadServiceServer interface {
	mustEmbedUnimplementedConversionAdjustmentUploadServiceServer()
}

func RegisterConversionAdjustmentUploadServiceServer(s grpc.ServiceRegistrar, srv ConversionAdjustmentUploadServiceServer) {
	s.RegisterService(&ConversionAdjustmentUploadService_ServiceDesc, srv)
}

func _ConversionAdjustmentUploadService_UploadConversionAdjustments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadConversionAdjustmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionAdjustmentUploadServiceServer).UploadConversionAdjustments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConversionAdjustmentUploadService_UploadConversionAdjustments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionAdjustmentUploadServiceServer).UploadConversionAdjustments(ctx, req.(*UploadConversionAdjustmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConversionAdjustmentUploadService_ServiceDesc is the grpc.ServiceDesc for ConversionAdjustmentUploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConversionAdjustmentUploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.ConversionAdjustmentUploadService",
	HandlerType: (*ConversionAdjustmentUploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadConversionAdjustments",
			Handler:    _ConversionAdjustmentUploadService_UploadConversionAdjustments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/conversion_adjustment_upload_service.proto",
}