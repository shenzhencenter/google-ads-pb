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
// source: google/ads/googleads/v18/services/ad_group_asset_set_service.proto

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
	AdGroupAssetSetService_MutateAdGroupAssetSets_FullMethodName = "/google.ads.googleads.v18.services.AdGroupAssetSetService/MutateAdGroupAssetSets"
)

// AdGroupAssetSetServiceClient is the client API for AdGroupAssetSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage ad group asset set
type AdGroupAssetSetServiceClient interface {
	// Creates, or removes ad group asset sets. Operation statuses are
	// returned.
	MutateAdGroupAssetSets(ctx context.Context, in *MutateAdGroupAssetSetsRequest, opts ...grpc.CallOption) (*MutateAdGroupAssetSetsResponse, error)
}

type adGroupAssetSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupAssetSetServiceClient(cc grpc.ClientConnInterface) AdGroupAssetSetServiceClient {
	return &adGroupAssetSetServiceClient{cc}
}

func (c *adGroupAssetSetServiceClient) MutateAdGroupAssetSets(ctx context.Context, in *MutateAdGroupAssetSetsRequest, opts ...grpc.CallOption) (*MutateAdGroupAssetSetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateAdGroupAssetSetsResponse)
	err := c.cc.Invoke(ctx, AdGroupAssetSetService_MutateAdGroupAssetSets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAssetSetServiceServer is the server API for AdGroupAssetSetService service.
// All implementations must embed UnimplementedAdGroupAssetSetServiceServer
// for forward compatibility.
//
// Service to manage ad group asset set
type AdGroupAssetSetServiceServer interface {
	// Creates, or removes ad group asset sets. Operation statuses are
	// returned.
	MutateAdGroupAssetSets(context.Context, *MutateAdGroupAssetSetsRequest) (*MutateAdGroupAssetSetsResponse, error)
	mustEmbedUnimplementedAdGroupAssetSetServiceServer()
}

// UnimplementedAdGroupAssetSetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdGroupAssetSetServiceServer struct{}

func (UnimplementedAdGroupAssetSetServiceServer) MutateAdGroupAssetSets(context.Context, *MutateAdGroupAssetSetsRequest) (*MutateAdGroupAssetSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroupAssetSets not implemented")
}
func (UnimplementedAdGroupAssetSetServiceServer) mustEmbedUnimplementedAdGroupAssetSetServiceServer() {
}
func (UnimplementedAdGroupAssetSetServiceServer) testEmbeddedByValue() {}

// UnsafeAdGroupAssetSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupAssetSetServiceServer will
// result in compilation errors.
type UnsafeAdGroupAssetSetServiceServer interface {
	mustEmbedUnimplementedAdGroupAssetSetServiceServer()
}

func RegisterAdGroupAssetSetServiceServer(s grpc.ServiceRegistrar, srv AdGroupAssetSetServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdGroupAssetSetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdGroupAssetSetService_ServiceDesc, srv)
}

func _AdGroupAssetSetService_MutateAdGroupAssetSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupAssetSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAssetSetServiceServer).MutateAdGroupAssetSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdGroupAssetSetService_MutateAdGroupAssetSets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAssetSetServiceServer).MutateAdGroupAssetSets(ctx, req.(*MutateAdGroupAssetSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdGroupAssetSetService_ServiceDesc is the grpc.ServiceDesc for AdGroupAssetSetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdGroupAssetSetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v18.services.AdGroupAssetSetService",
	HandlerType: (*AdGroupAssetSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdGroupAssetSets",
			Handler:    _AdGroupAssetSetService_MutateAdGroupAssetSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v18/services/ad_group_asset_set_service.proto",
}
