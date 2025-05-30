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
// source: google/ads/googleads/v19/services/asset_group_listing_group_filter_service.proto

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
	AssetGroupListingGroupFilterService_MutateAssetGroupListingGroupFilters_FullMethodName = "/google.ads.googleads.v19.services.AssetGroupListingGroupFilterService/MutateAssetGroupListingGroupFilters"
)

// AssetGroupListingGroupFilterServiceClient is the client API for AssetGroupListingGroupFilterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage asset group listing group filter.
type AssetGroupListingGroupFilterServiceClient interface {
	// Creates, updates or removes asset group listing group filters. Operation
	// statuses are returned.
	MutateAssetGroupListingGroupFilters(ctx context.Context, in *MutateAssetGroupListingGroupFiltersRequest, opts ...grpc.CallOption) (*MutateAssetGroupListingGroupFiltersResponse, error)
}

type assetGroupListingGroupFilterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetGroupListingGroupFilterServiceClient(cc grpc.ClientConnInterface) AssetGroupListingGroupFilterServiceClient {
	return &assetGroupListingGroupFilterServiceClient{cc}
}

func (c *assetGroupListingGroupFilterServiceClient) MutateAssetGroupListingGroupFilters(ctx context.Context, in *MutateAssetGroupListingGroupFiltersRequest, opts ...grpc.CallOption) (*MutateAssetGroupListingGroupFiltersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateAssetGroupListingGroupFiltersResponse)
	err := c.cc.Invoke(ctx, AssetGroupListingGroupFilterService_MutateAssetGroupListingGroupFilters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetGroupListingGroupFilterServiceServer is the server API for AssetGroupListingGroupFilterService service.
// All implementations must embed UnimplementedAssetGroupListingGroupFilterServiceServer
// for forward compatibility.
//
// Service to manage asset group listing group filter.
type AssetGroupListingGroupFilterServiceServer interface {
	// Creates, updates or removes asset group listing group filters. Operation
	// statuses are returned.
	MutateAssetGroupListingGroupFilters(context.Context, *MutateAssetGroupListingGroupFiltersRequest) (*MutateAssetGroupListingGroupFiltersResponse, error)
	mustEmbedUnimplementedAssetGroupListingGroupFilterServiceServer()
}

// UnimplementedAssetGroupListingGroupFilterServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAssetGroupListingGroupFilterServiceServer struct{}

func (UnimplementedAssetGroupListingGroupFilterServiceServer) MutateAssetGroupListingGroupFilters(context.Context, *MutateAssetGroupListingGroupFiltersRequest) (*MutateAssetGroupListingGroupFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAssetGroupListingGroupFilters not implemented")
}
func (UnimplementedAssetGroupListingGroupFilterServiceServer) mustEmbedUnimplementedAssetGroupListingGroupFilterServiceServer() {
}
func (UnimplementedAssetGroupListingGroupFilterServiceServer) testEmbeddedByValue() {}

// UnsafeAssetGroupListingGroupFilterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetGroupListingGroupFilterServiceServer will
// result in compilation errors.
type UnsafeAssetGroupListingGroupFilterServiceServer interface {
	mustEmbedUnimplementedAssetGroupListingGroupFilterServiceServer()
}

func RegisterAssetGroupListingGroupFilterServiceServer(s grpc.ServiceRegistrar, srv AssetGroupListingGroupFilterServiceServer) {
	// If the following call pancis, it indicates UnimplementedAssetGroupListingGroupFilterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AssetGroupListingGroupFilterService_ServiceDesc, srv)
}

func _AssetGroupListingGroupFilterService_MutateAssetGroupListingGroupFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAssetGroupListingGroupFiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetGroupListingGroupFilterServiceServer).MutateAssetGroupListingGroupFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetGroupListingGroupFilterService_MutateAssetGroupListingGroupFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetGroupListingGroupFilterServiceServer).MutateAssetGroupListingGroupFilters(ctx, req.(*MutateAssetGroupListingGroupFiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetGroupListingGroupFilterService_ServiceDesc is the grpc.ServiceDesc for AssetGroupListingGroupFilterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetGroupListingGroupFilterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.AssetGroupListingGroupFilterService",
	HandlerType: (*AssetGroupListingGroupFilterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAssetGroupListingGroupFilters",
			Handler:    _AssetGroupListingGroupFilterService_MutateAssetGroupListingGroupFilters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/asset_group_listing_group_filter_service.proto",
}
