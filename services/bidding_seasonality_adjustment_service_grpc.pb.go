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
// - protoc             v4.24.4
// source: google/ads/googleads/v15/services/bidding_seasonality_adjustment_service.proto

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
	BiddingSeasonalityAdjustmentService_MutateBiddingSeasonalityAdjustments_FullMethodName = "/google.ads.googleads.v15.services.BiddingSeasonalityAdjustmentService/MutateBiddingSeasonalityAdjustments"
)

// BiddingSeasonalityAdjustmentServiceClient is the client API for BiddingSeasonalityAdjustmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BiddingSeasonalityAdjustmentServiceClient interface {
	// Creates, updates, or removes seasonality adjustments.
	// Operation statuses are returned.
	MutateBiddingSeasonalityAdjustments(ctx context.Context, in *MutateBiddingSeasonalityAdjustmentsRequest, opts ...grpc.CallOption) (*MutateBiddingSeasonalityAdjustmentsResponse, error)
}

type biddingSeasonalityAdjustmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBiddingSeasonalityAdjustmentServiceClient(cc grpc.ClientConnInterface) BiddingSeasonalityAdjustmentServiceClient {
	return &biddingSeasonalityAdjustmentServiceClient{cc}
}

func (c *biddingSeasonalityAdjustmentServiceClient) MutateBiddingSeasonalityAdjustments(ctx context.Context, in *MutateBiddingSeasonalityAdjustmentsRequest, opts ...grpc.CallOption) (*MutateBiddingSeasonalityAdjustmentsResponse, error) {
	out := new(MutateBiddingSeasonalityAdjustmentsResponse)
	err := c.cc.Invoke(ctx, BiddingSeasonalityAdjustmentService_MutateBiddingSeasonalityAdjustments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BiddingSeasonalityAdjustmentServiceServer is the server API for BiddingSeasonalityAdjustmentService service.
// All implementations must embed UnimplementedBiddingSeasonalityAdjustmentServiceServer
// for forward compatibility
type BiddingSeasonalityAdjustmentServiceServer interface {
	// Creates, updates, or removes seasonality adjustments.
	// Operation statuses are returned.
	MutateBiddingSeasonalityAdjustments(context.Context, *MutateBiddingSeasonalityAdjustmentsRequest) (*MutateBiddingSeasonalityAdjustmentsResponse, error)
	mustEmbedUnimplementedBiddingSeasonalityAdjustmentServiceServer()
}

// UnimplementedBiddingSeasonalityAdjustmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBiddingSeasonalityAdjustmentServiceServer struct {
}

func (UnimplementedBiddingSeasonalityAdjustmentServiceServer) MutateBiddingSeasonalityAdjustments(context.Context, *MutateBiddingSeasonalityAdjustmentsRequest) (*MutateBiddingSeasonalityAdjustmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateBiddingSeasonalityAdjustments not implemented")
}
func (UnimplementedBiddingSeasonalityAdjustmentServiceServer) mustEmbedUnimplementedBiddingSeasonalityAdjustmentServiceServer() {
}

// UnsafeBiddingSeasonalityAdjustmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BiddingSeasonalityAdjustmentServiceServer will
// result in compilation errors.
type UnsafeBiddingSeasonalityAdjustmentServiceServer interface {
	mustEmbedUnimplementedBiddingSeasonalityAdjustmentServiceServer()
}

func RegisterBiddingSeasonalityAdjustmentServiceServer(s grpc.ServiceRegistrar, srv BiddingSeasonalityAdjustmentServiceServer) {
	s.RegisterService(&BiddingSeasonalityAdjustmentService_ServiceDesc, srv)
}

func _BiddingSeasonalityAdjustmentService_MutateBiddingSeasonalityAdjustments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBiddingSeasonalityAdjustmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiddingSeasonalityAdjustmentServiceServer).MutateBiddingSeasonalityAdjustments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BiddingSeasonalityAdjustmentService_MutateBiddingSeasonalityAdjustments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiddingSeasonalityAdjustmentServiceServer).MutateBiddingSeasonalityAdjustments(ctx, req.(*MutateBiddingSeasonalityAdjustmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BiddingSeasonalityAdjustmentService_ServiceDesc is the grpc.ServiceDesc for BiddingSeasonalityAdjustmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BiddingSeasonalityAdjustmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.BiddingSeasonalityAdjustmentService",
	HandlerType: (*BiddingSeasonalityAdjustmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateBiddingSeasonalityAdjustments",
			Handler:    _BiddingSeasonalityAdjustmentService_MutateBiddingSeasonalityAdjustments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/bidding_seasonality_adjustment_service.proto",
}
