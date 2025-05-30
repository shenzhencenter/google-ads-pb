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
// source: google/ads/googleads/v19/services/shareable_preview_service.proto

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
	ShareablePreviewService_GenerateShareablePreviews_FullMethodName = "/google.ads.googleads.v19.services.ShareablePreviewService/GenerateShareablePreviews"
)

// ShareablePreviewServiceClient is the client API for ShareablePreviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to generate Shareable Previews.
type ShareablePreviewServiceClient interface {
	// Returns the requested Shareable Preview.
	GenerateShareablePreviews(ctx context.Context, in *GenerateShareablePreviewsRequest, opts ...grpc.CallOption) (*GenerateShareablePreviewsResponse, error)
}

type shareablePreviewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShareablePreviewServiceClient(cc grpc.ClientConnInterface) ShareablePreviewServiceClient {
	return &shareablePreviewServiceClient{cc}
}

func (c *shareablePreviewServiceClient) GenerateShareablePreviews(ctx context.Context, in *GenerateShareablePreviewsRequest, opts ...grpc.CallOption) (*GenerateShareablePreviewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateShareablePreviewsResponse)
	err := c.cc.Invoke(ctx, ShareablePreviewService_GenerateShareablePreviews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShareablePreviewServiceServer is the server API for ShareablePreviewService service.
// All implementations must embed UnimplementedShareablePreviewServiceServer
// for forward compatibility.
//
// Service to generate Shareable Previews.
type ShareablePreviewServiceServer interface {
	// Returns the requested Shareable Preview.
	GenerateShareablePreviews(context.Context, *GenerateShareablePreviewsRequest) (*GenerateShareablePreviewsResponse, error)
	mustEmbedUnimplementedShareablePreviewServiceServer()
}

// UnimplementedShareablePreviewServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedShareablePreviewServiceServer struct{}

func (UnimplementedShareablePreviewServiceServer) GenerateShareablePreviews(context.Context, *GenerateShareablePreviewsRequest) (*GenerateShareablePreviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateShareablePreviews not implemented")
}
func (UnimplementedShareablePreviewServiceServer) mustEmbedUnimplementedShareablePreviewServiceServer() {
}
func (UnimplementedShareablePreviewServiceServer) testEmbeddedByValue() {}

// UnsafeShareablePreviewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShareablePreviewServiceServer will
// result in compilation errors.
type UnsafeShareablePreviewServiceServer interface {
	mustEmbedUnimplementedShareablePreviewServiceServer()
}

func RegisterShareablePreviewServiceServer(s grpc.ServiceRegistrar, srv ShareablePreviewServiceServer) {
	// If the following call pancis, it indicates UnimplementedShareablePreviewServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ShareablePreviewService_ServiceDesc, srv)
}

func _ShareablePreviewService_GenerateShareablePreviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateShareablePreviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareablePreviewServiceServer).GenerateShareablePreviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShareablePreviewService_GenerateShareablePreviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareablePreviewServiceServer).GenerateShareablePreviews(ctx, req.(*GenerateShareablePreviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShareablePreviewService_ServiceDesc is the grpc.ServiceDesc for ShareablePreviewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShareablePreviewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.ShareablePreviewService",
	HandlerType: (*ShareablePreviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateShareablePreviews",
			Handler:    _ShareablePreviewService_GenerateShareablePreviews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/shareable_preview_service.proto",
}
