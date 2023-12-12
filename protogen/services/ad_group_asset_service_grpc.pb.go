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
// source: google/ads/googleads/v15/services/ad_group_asset_service.proto

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
	AdGroupAssetService_MutateAdGroupAssets_FullMethodName = "/google.ads.googleads.v15.services.AdGroupAssetService/MutateAdGroupAssets"
)

// AdGroupAssetServiceClient is the client API for AdGroupAssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdGroupAssetServiceClient interface {
	// Creates, updates, or removes ad group assets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AssetLinkError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ContextError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NotAllowlistedError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateAdGroupAssets(ctx context.Context, in *MutateAdGroupAssetsRequest, opts ...grpc.CallOption) (*MutateAdGroupAssetsResponse, error)
}

type adGroupAssetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupAssetServiceClient(cc grpc.ClientConnInterface) AdGroupAssetServiceClient {
	return &adGroupAssetServiceClient{cc}
}

func (c *adGroupAssetServiceClient) MutateAdGroupAssets(ctx context.Context, in *MutateAdGroupAssetsRequest, opts ...grpc.CallOption) (*MutateAdGroupAssetsResponse, error) {
	out := new(MutateAdGroupAssetsResponse)
	err := c.cc.Invoke(ctx, AdGroupAssetService_MutateAdGroupAssets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAssetServiceServer is the server API for AdGroupAssetService service.
// All implementations must embed UnimplementedAdGroupAssetServiceServer
// for forward compatibility
type AdGroupAssetServiceServer interface {
	// Creates, updates, or removes ad group assets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AssetLinkError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ContextError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NotAllowlistedError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateAdGroupAssets(context.Context, *MutateAdGroupAssetsRequest) (*MutateAdGroupAssetsResponse, error)
	mustEmbedUnimplementedAdGroupAssetServiceServer()
}

// UnimplementedAdGroupAssetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdGroupAssetServiceServer struct {
}

func (UnimplementedAdGroupAssetServiceServer) MutateAdGroupAssets(context.Context, *MutateAdGroupAssetsRequest) (*MutateAdGroupAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroupAssets not implemented")
}
func (UnimplementedAdGroupAssetServiceServer) mustEmbedUnimplementedAdGroupAssetServiceServer() {}

// UnsafeAdGroupAssetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupAssetServiceServer will
// result in compilation errors.
type UnsafeAdGroupAssetServiceServer interface {
	mustEmbedUnimplementedAdGroupAssetServiceServer()
}

func RegisterAdGroupAssetServiceServer(s grpc.ServiceRegistrar, srv AdGroupAssetServiceServer) {
	s.RegisterService(&AdGroupAssetService_ServiceDesc, srv)
}

func _AdGroupAssetService_MutateAdGroupAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAssetServiceServer).MutateAdGroupAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdGroupAssetService_MutateAdGroupAssets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAssetServiceServer).MutateAdGroupAssets(ctx, req.(*MutateAdGroupAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdGroupAssetService_ServiceDesc is the grpc.ServiceDesc for AdGroupAssetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdGroupAssetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.AdGroupAssetService",
	HandlerType: (*AdGroupAssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdGroupAssets",
			Handler:    _AdGroupAssetService_MutateAdGroupAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/ad_group_asset_service.proto",
}