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
// source: google/ads/googleads/v17/services/bidding_strategy_service.proto

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
	BiddingStrategyService_MutateBiddingStrategies_FullMethodName = "/google.ads.googleads.v17.services.BiddingStrategyService/MutateBiddingStrategies"
)

// BiddingStrategyServiceClient is the client API for BiddingStrategyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage bidding strategies.
type BiddingStrategyServiceClient interface {
	// Creates, updates, or removes bidding strategies. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AdxError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BiddingError]()
	//	[BiddingStrategyError]()
	//	[ContextError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateBiddingStrategies(ctx context.Context, in *MutateBiddingStrategiesRequest, opts ...grpc.CallOption) (*MutateBiddingStrategiesResponse, error)
}

type biddingStrategyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBiddingStrategyServiceClient(cc grpc.ClientConnInterface) BiddingStrategyServiceClient {
	return &biddingStrategyServiceClient{cc}
}

func (c *biddingStrategyServiceClient) MutateBiddingStrategies(ctx context.Context, in *MutateBiddingStrategiesRequest, opts ...grpc.CallOption) (*MutateBiddingStrategiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateBiddingStrategiesResponse)
	err := c.cc.Invoke(ctx, BiddingStrategyService_MutateBiddingStrategies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BiddingStrategyServiceServer is the server API for BiddingStrategyService service.
// All implementations must embed UnimplementedBiddingStrategyServiceServer
// for forward compatibility
//
// Service to manage bidding strategies.
type BiddingStrategyServiceServer interface {
	// Creates, updates, or removes bidding strategies. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AdxError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BiddingError]()
	//	[BiddingStrategyError]()
	//	[ContextError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperationAccessDeniedError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateBiddingStrategies(context.Context, *MutateBiddingStrategiesRequest) (*MutateBiddingStrategiesResponse, error)
	mustEmbedUnimplementedBiddingStrategyServiceServer()
}

// UnimplementedBiddingStrategyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBiddingStrategyServiceServer struct {
}

func (UnimplementedBiddingStrategyServiceServer) MutateBiddingStrategies(context.Context, *MutateBiddingStrategiesRequest) (*MutateBiddingStrategiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateBiddingStrategies not implemented")
}
func (UnimplementedBiddingStrategyServiceServer) mustEmbedUnimplementedBiddingStrategyServiceServer() {
}

// UnsafeBiddingStrategyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BiddingStrategyServiceServer will
// result in compilation errors.
type UnsafeBiddingStrategyServiceServer interface {
	mustEmbedUnimplementedBiddingStrategyServiceServer()
}

func RegisterBiddingStrategyServiceServer(s grpc.ServiceRegistrar, srv BiddingStrategyServiceServer) {
	s.RegisterService(&BiddingStrategyService_ServiceDesc, srv)
}

func _BiddingStrategyService_MutateBiddingStrategies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBiddingStrategiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiddingStrategyServiceServer).MutateBiddingStrategies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BiddingStrategyService_MutateBiddingStrategies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiddingStrategyServiceServer).MutateBiddingStrategies(ctx, req.(*MutateBiddingStrategiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BiddingStrategyService_ServiceDesc is the grpc.ServiceDesc for BiddingStrategyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BiddingStrategyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.BiddingStrategyService",
	HandlerType: (*BiddingStrategyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateBiddingStrategies",
			Handler:    _BiddingStrategyService_MutateBiddingStrategies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/bidding_strategy_service.proto",
}
