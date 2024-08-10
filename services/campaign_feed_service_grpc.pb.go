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
// source: google/ads/googleads/v17/services/campaign_feed_service.proto

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
	CampaignFeedService_MutateCampaignFeeds_FullMethodName = "/google.ads.googleads.v17.services.CampaignFeedService/MutateCampaignFeeds"
)

// CampaignFeedServiceClient is the client API for CampaignFeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage campaign feeds.
type CampaignFeedServiceClient interface {
	// Creates, updates, or removes campaign feeds. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignFeedError]()
	//	[CollectionSizeError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FunctionError]()
	//	[FunctionParsingError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
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
	MutateCampaignFeeds(ctx context.Context, in *MutateCampaignFeedsRequest, opts ...grpc.CallOption) (*MutateCampaignFeedsResponse, error)
}

type campaignFeedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignFeedServiceClient(cc grpc.ClientConnInterface) CampaignFeedServiceClient {
	return &campaignFeedServiceClient{cc}
}

func (c *campaignFeedServiceClient) MutateCampaignFeeds(ctx context.Context, in *MutateCampaignFeedsRequest, opts ...grpc.CallOption) (*MutateCampaignFeedsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateCampaignFeedsResponse)
	err := c.cc.Invoke(ctx, CampaignFeedService_MutateCampaignFeeds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignFeedServiceServer is the server API for CampaignFeedService service.
// All implementations must embed UnimplementedCampaignFeedServiceServer
// for forward compatibility.
//
// Service to manage campaign feeds.
type CampaignFeedServiceServer interface {
	// Creates, updates, or removes campaign feeds. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignFeedError]()
	//	[CollectionSizeError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FunctionError]()
	//	[FunctionParsingError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
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
	MutateCampaignFeeds(context.Context, *MutateCampaignFeedsRequest) (*MutateCampaignFeedsResponse, error)
	mustEmbedUnimplementedCampaignFeedServiceServer()
}

// UnimplementedCampaignFeedServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCampaignFeedServiceServer struct{}

func (UnimplementedCampaignFeedServiceServer) MutateCampaignFeeds(context.Context, *MutateCampaignFeedsRequest) (*MutateCampaignFeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignFeeds not implemented")
}
func (UnimplementedCampaignFeedServiceServer) mustEmbedUnimplementedCampaignFeedServiceServer() {}
func (UnimplementedCampaignFeedServiceServer) testEmbeddedByValue()                             {}

// UnsafeCampaignFeedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignFeedServiceServer will
// result in compilation errors.
type UnsafeCampaignFeedServiceServer interface {
	mustEmbedUnimplementedCampaignFeedServiceServer()
}

func RegisterCampaignFeedServiceServer(s grpc.ServiceRegistrar, srv CampaignFeedServiceServer) {
	// If the following call pancis, it indicates UnimplementedCampaignFeedServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CampaignFeedService_ServiceDesc, srv)
}

func _CampaignFeedService_MutateCampaignFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignFeedServiceServer).MutateCampaignFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampaignFeedService_MutateCampaignFeeds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignFeedServiceServer).MutateCampaignFeeds(ctx, req.(*MutateCampaignFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignFeedService_ServiceDesc is the grpc.ServiceDesc for CampaignFeedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignFeedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.CampaignFeedService",
	HandlerType: (*CampaignFeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignFeeds",
			Handler:    _CampaignFeedService_MutateCampaignFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/campaign_feed_service.proto",
}
