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
// source: google/ads/googleads/v19/services/campaign_asset_service.proto

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
	CampaignAssetService_MutateCampaignAssets_FullMethodName = "/google.ads.googleads.v19.services.CampaignAssetService/MutateCampaignAssets"
)

// CampaignAssetServiceClient is the client API for CampaignAssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage campaign assets.
type CampaignAssetServiceClient interface {
	// Creates, updates, or removes campaign assets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AssetLinkError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ContextError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NotAllowlistedError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateCampaignAssets(ctx context.Context, in *MutateCampaignAssetsRequest, opts ...grpc.CallOption) (*MutateCampaignAssetsResponse, error)
}

type campaignAssetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignAssetServiceClient(cc grpc.ClientConnInterface) CampaignAssetServiceClient {
	return &campaignAssetServiceClient{cc}
}

func (c *campaignAssetServiceClient) MutateCampaignAssets(ctx context.Context, in *MutateCampaignAssetsRequest, opts ...grpc.CallOption) (*MutateCampaignAssetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateCampaignAssetsResponse)
	err := c.cc.Invoke(ctx, CampaignAssetService_MutateCampaignAssets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignAssetServiceServer is the server API for CampaignAssetService service.
// All implementations must embed UnimplementedCampaignAssetServiceServer
// for forward compatibility.
//
// Service to manage campaign assets.
type CampaignAssetServiceServer interface {
	// Creates, updates, or removes campaign assets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AssetLinkError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ContextError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NotAllowlistedError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateCampaignAssets(context.Context, *MutateCampaignAssetsRequest) (*MutateCampaignAssetsResponse, error)
	mustEmbedUnimplementedCampaignAssetServiceServer()
}

// UnimplementedCampaignAssetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCampaignAssetServiceServer struct{}

func (UnimplementedCampaignAssetServiceServer) MutateCampaignAssets(context.Context, *MutateCampaignAssetsRequest) (*MutateCampaignAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignAssets not implemented")
}
func (UnimplementedCampaignAssetServiceServer) mustEmbedUnimplementedCampaignAssetServiceServer() {}
func (UnimplementedCampaignAssetServiceServer) testEmbeddedByValue()                              {}

// UnsafeCampaignAssetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignAssetServiceServer will
// result in compilation errors.
type UnsafeCampaignAssetServiceServer interface {
	mustEmbedUnimplementedCampaignAssetServiceServer()
}

func RegisterCampaignAssetServiceServer(s grpc.ServiceRegistrar, srv CampaignAssetServiceServer) {
	// If the following call pancis, it indicates UnimplementedCampaignAssetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CampaignAssetService_ServiceDesc, srv)
}

func _CampaignAssetService_MutateCampaignAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignAssetServiceServer).MutateCampaignAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampaignAssetService_MutateCampaignAssets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignAssetServiceServer).MutateCampaignAssets(ctx, req.(*MutateCampaignAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignAssetService_ServiceDesc is the grpc.ServiceDesc for CampaignAssetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignAssetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.CampaignAssetService",
	HandlerType: (*CampaignAssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignAssets",
			Handler:    _CampaignAssetService_MutateCampaignAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/campaign_asset_service.proto",
}
