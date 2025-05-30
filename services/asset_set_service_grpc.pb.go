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
// source: google/ads/googleads/v19/services/asset_set_service.proto

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
	AssetSetService_MutateAssetSets_FullMethodName = "/google.ads.googleads.v19.services.AssetSetService/MutateAssetSets"
)

// AssetSetServiceClient is the client API for AssetSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage asset set
type AssetSetServiceClient interface {
	// Creates, updates or removes asset sets. Operation statuses are
	// returned.
	MutateAssetSets(ctx context.Context, in *MutateAssetSetsRequest, opts ...grpc.CallOption) (*MutateAssetSetsResponse, error)
}

type assetSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetSetServiceClient(cc grpc.ClientConnInterface) AssetSetServiceClient {
	return &assetSetServiceClient{cc}
}

func (c *assetSetServiceClient) MutateAssetSets(ctx context.Context, in *MutateAssetSetsRequest, opts ...grpc.CallOption) (*MutateAssetSetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateAssetSetsResponse)
	err := c.cc.Invoke(ctx, AssetSetService_MutateAssetSets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetSetServiceServer is the server API for AssetSetService service.
// All implementations must embed UnimplementedAssetSetServiceServer
// for forward compatibility.
//
// Service to manage asset set
type AssetSetServiceServer interface {
	// Creates, updates or removes asset sets. Operation statuses are
	// returned.
	MutateAssetSets(context.Context, *MutateAssetSetsRequest) (*MutateAssetSetsResponse, error)
	mustEmbedUnimplementedAssetSetServiceServer()
}

// UnimplementedAssetSetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAssetSetServiceServer struct{}

func (UnimplementedAssetSetServiceServer) MutateAssetSets(context.Context, *MutateAssetSetsRequest) (*MutateAssetSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAssetSets not implemented")
}
func (UnimplementedAssetSetServiceServer) mustEmbedUnimplementedAssetSetServiceServer() {}
func (UnimplementedAssetSetServiceServer) testEmbeddedByValue()                         {}

// UnsafeAssetSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetSetServiceServer will
// result in compilation errors.
type UnsafeAssetSetServiceServer interface {
	mustEmbedUnimplementedAssetSetServiceServer()
}

func RegisterAssetSetServiceServer(s grpc.ServiceRegistrar, srv AssetSetServiceServer) {
	// If the following call pancis, it indicates UnimplementedAssetSetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AssetSetService_ServiceDesc, srv)
}

func _AssetSetService_MutateAssetSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAssetSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetSetServiceServer).MutateAssetSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetSetService_MutateAssetSets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetSetServiceServer).MutateAssetSets(ctx, req.(*MutateAssetSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetSetService_ServiceDesc is the grpc.ServiceDesc for AssetSetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetSetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.AssetSetService",
	HandlerType: (*AssetSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAssetSets",
			Handler:    _AssetSetService_MutateAssetSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/asset_set_service.proto",
}
