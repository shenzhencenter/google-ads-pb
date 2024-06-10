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
// source: google/ads/googleads/v17/services/asset_group_signal_service.proto

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
	AssetGroupSignalService_MutateAssetGroupSignals_FullMethodName = "/google.ads.googleads.v17.services.AssetGroupSignalService/MutateAssetGroupSignals"
)

// AssetGroupSignalServiceClient is the client API for AssetGroupSignalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage asset group signal.
type AssetGroupSignalServiceClient interface {
	// Creates or removes asset group signals. Operation statuses are
	// returned.
	MutateAssetGroupSignals(ctx context.Context, in *MutateAssetGroupSignalsRequest, opts ...grpc.CallOption) (*MutateAssetGroupSignalsResponse, error)
}

type assetGroupSignalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetGroupSignalServiceClient(cc grpc.ClientConnInterface) AssetGroupSignalServiceClient {
	return &assetGroupSignalServiceClient{cc}
}

func (c *assetGroupSignalServiceClient) MutateAssetGroupSignals(ctx context.Context, in *MutateAssetGroupSignalsRequest, opts ...grpc.CallOption) (*MutateAssetGroupSignalsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateAssetGroupSignalsResponse)
	err := c.cc.Invoke(ctx, AssetGroupSignalService_MutateAssetGroupSignals_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetGroupSignalServiceServer is the server API for AssetGroupSignalService service.
// All implementations must embed UnimplementedAssetGroupSignalServiceServer
// for forward compatibility
//
// Service to manage asset group signal.
type AssetGroupSignalServiceServer interface {
	// Creates or removes asset group signals. Operation statuses are
	// returned.
	MutateAssetGroupSignals(context.Context, *MutateAssetGroupSignalsRequest) (*MutateAssetGroupSignalsResponse, error)
	mustEmbedUnimplementedAssetGroupSignalServiceServer()
}

// UnimplementedAssetGroupSignalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssetGroupSignalServiceServer struct {
}

func (UnimplementedAssetGroupSignalServiceServer) MutateAssetGroupSignals(context.Context, *MutateAssetGroupSignalsRequest) (*MutateAssetGroupSignalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAssetGroupSignals not implemented")
}
func (UnimplementedAssetGroupSignalServiceServer) mustEmbedUnimplementedAssetGroupSignalServiceServer() {
}

// UnsafeAssetGroupSignalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetGroupSignalServiceServer will
// result in compilation errors.
type UnsafeAssetGroupSignalServiceServer interface {
	mustEmbedUnimplementedAssetGroupSignalServiceServer()
}

func RegisterAssetGroupSignalServiceServer(s grpc.ServiceRegistrar, srv AssetGroupSignalServiceServer) {
	s.RegisterService(&AssetGroupSignalService_ServiceDesc, srv)
}

func _AssetGroupSignalService_MutateAssetGroupSignals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAssetGroupSignalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetGroupSignalServiceServer).MutateAssetGroupSignals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetGroupSignalService_MutateAssetGroupSignals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetGroupSignalServiceServer).MutateAssetGroupSignals(ctx, req.(*MutateAssetGroupSignalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetGroupSignalService_ServiceDesc is the grpc.ServiceDesc for AssetGroupSignalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetGroupSignalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.AssetGroupSignalService",
	HandlerType: (*AssetGroupSignalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAssetGroupSignals",
			Handler:    _AssetGroupSignalService_MutateAssetGroupSignals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/asset_group_signal_service.proto",
}
