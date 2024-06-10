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
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.24.4
// source: google/ads/googleads/v17/services/customer_asset_set_service.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CustomerAssetSetService_MutateCustomerAssetSets_FullMethodName = "/google.ads.googleads.v17.services.CustomerAssetSetService/MutateCustomerAssetSets"
)

// CustomerAssetSetServiceClient is the client API for CustomerAssetSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage customer asset set
type CustomerAssetSetServiceClient interface {
	// Creates, or removes customer asset sets. Operation statuses are
	// returned.
	MutateCustomerAssetSets(ctx context.Context, in *MutateCustomerAssetSetsRequest, opts ...grpc.CallOption) (*MutateCustomerAssetSetsResponse, error)
}

type customerAssetSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerAssetSetServiceClient(cc grpc.ClientConnInterface) CustomerAssetSetServiceClient {
	return &customerAssetSetServiceClient{cc}
}

func (c *customerAssetSetServiceClient) MutateCustomerAssetSets(ctx context.Context, in *MutateCustomerAssetSetsRequest, opts ...grpc.CallOption) (*MutateCustomerAssetSetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateCustomerAssetSetsResponse)
	err := c.cc.Invoke(ctx, CustomerAssetSetService_MutateCustomerAssetSets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerAssetSetServiceServer is the server API for CustomerAssetSetService service.
// All implementations must embed UnimplementedCustomerAssetSetServiceServer
// for forward compatibility
//
// Service to manage customer asset set
type CustomerAssetSetServiceServer interface {
	// Creates, or removes customer asset sets. Operation statuses are
	// returned.
	MutateCustomerAssetSets(context.Context, *MutateCustomerAssetSetsRequest) (*MutateCustomerAssetSetsResponse, error)
	mustEmbedUnimplementedCustomerAssetSetServiceServer()
}

// UnimplementedCustomerAssetSetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerAssetSetServiceServer struct {
}

func (UnimplementedCustomerAssetSetServiceServer) MutateCustomerAssetSets(context.Context, *MutateCustomerAssetSetsRequest) (*MutateCustomerAssetSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerAssetSets not implemented")
}
func (UnimplementedCustomerAssetSetServiceServer) mustEmbedUnimplementedCustomerAssetSetServiceServer() {
}

// UnsafeCustomerAssetSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerAssetSetServiceServer will
// result in compilation errors.
type UnsafeCustomerAssetSetServiceServer interface {
	mustEmbedUnimplementedCustomerAssetSetServiceServer()
}

func RegisterCustomerAssetSetServiceServer(s grpc.ServiceRegistrar, srv CustomerAssetSetServiceServer) {
	s.RegisterService(&CustomerAssetSetService_ServiceDesc, srv)
}

func _CustomerAssetSetService_MutateCustomerAssetSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerAssetSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAssetSetServiceServer).MutateCustomerAssetSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerAssetSetService_MutateCustomerAssetSets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAssetSetServiceServer).MutateCustomerAssetSets(ctx, req.(*MutateCustomerAssetSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerAssetSetService_ServiceDesc is the grpc.ServiceDesc for CustomerAssetSetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerAssetSetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.CustomerAssetSetService",
	HandlerType: (*CustomerAssetSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCustomerAssetSets",
			Handler:    _CustomerAssetSetService_MutateCustomerAssetSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/customer_asset_set_service.proto",
}
