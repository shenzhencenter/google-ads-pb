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
// source: google/ads/googleads/v18/services/campaign_criterion_service.proto

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
	CampaignCriterionService_MutateCampaignCriteria_FullMethodName = "/google.ads.googleads.v18.services.CampaignCriterionService/MutateCampaignCriteria"
)

// CampaignCriterionServiceClient is the client API for CampaignCriterionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage campaign criteria.
type CampaignCriterionServiceClient interface {
	// Creates, updates, or removes criteria. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AdxError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignCriterionError]()
	//	[CollectionSizeError]()
	//	[ContextError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[FunctionError]()
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
	//	[RegionCodeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateCampaignCriteria(ctx context.Context, in *MutateCampaignCriteriaRequest, opts ...grpc.CallOption) (*MutateCampaignCriteriaResponse, error)
}

type campaignCriterionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignCriterionServiceClient(cc grpc.ClientConnInterface) CampaignCriterionServiceClient {
	return &campaignCriterionServiceClient{cc}
}

func (c *campaignCriterionServiceClient) MutateCampaignCriteria(ctx context.Context, in *MutateCampaignCriteriaRequest, opts ...grpc.CallOption) (*MutateCampaignCriteriaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateCampaignCriteriaResponse)
	err := c.cc.Invoke(ctx, CampaignCriterionService_MutateCampaignCriteria_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignCriterionServiceServer is the server API for CampaignCriterionService service.
// All implementations must embed UnimplementedCampaignCriterionServiceServer
// for forward compatibility.
//
// Service to manage campaign criteria.
type CampaignCriterionServiceServer interface {
	// Creates, updates, or removes criteria. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AdxError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignCriterionError]()
	//	[CollectionSizeError]()
	//	[ContextError]()
	//	[CriterionError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[FunctionError]()
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
	//	[RegionCodeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateCampaignCriteria(context.Context, *MutateCampaignCriteriaRequest) (*MutateCampaignCriteriaResponse, error)
	mustEmbedUnimplementedCampaignCriterionServiceServer()
}

// UnimplementedCampaignCriterionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCampaignCriterionServiceServer struct{}

func (UnimplementedCampaignCriterionServiceServer) MutateCampaignCriteria(context.Context, *MutateCampaignCriteriaRequest) (*MutateCampaignCriteriaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignCriteria not implemented")
}
func (UnimplementedCampaignCriterionServiceServer) mustEmbedUnimplementedCampaignCriterionServiceServer() {
}
func (UnimplementedCampaignCriterionServiceServer) testEmbeddedByValue() {}

// UnsafeCampaignCriterionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignCriterionServiceServer will
// result in compilation errors.
type UnsafeCampaignCriterionServiceServer interface {
	mustEmbedUnimplementedCampaignCriterionServiceServer()
}

func RegisterCampaignCriterionServiceServer(s grpc.ServiceRegistrar, srv CampaignCriterionServiceServer) {
	// If the following call pancis, it indicates UnimplementedCampaignCriterionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CampaignCriterionService_ServiceDesc, srv)
}

func _CampaignCriterionService_MutateCampaignCriteria_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignCriteriaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignCriterionServiceServer).MutateCampaignCriteria(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampaignCriterionService_MutateCampaignCriteria_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignCriterionServiceServer).MutateCampaignCriteria(ctx, req.(*MutateCampaignCriteriaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignCriterionService_ServiceDesc is the grpc.ServiceDesc for CampaignCriterionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignCriterionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v18.services.CampaignCriterionService",
	HandlerType: (*CampaignCriterionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignCriteria",
			Handler:    _CampaignCriterionService_MutateCampaignCriteria_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v18/services/campaign_criterion_service.proto",
}
