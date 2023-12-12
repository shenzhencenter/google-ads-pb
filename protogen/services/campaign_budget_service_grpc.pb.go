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
// - protoc             v3.21.1
// source: google/ads/googleads/v15/services/campaign_budget_service.proto

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
	CampaignBudgetService_MutateCampaignBudgets_FullMethodName = "/google.ads.googleads.v15.services.CampaignBudgetService/MutateCampaignBudgets"
)

// CampaignBudgetServiceClient is the client API for CampaignBudgetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampaignBudgetServiceClient interface {
	// Creates, updates, or removes campaign budgets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignBudgetError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[OperationAccessDeniedError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[StringLengthError]()
	MutateCampaignBudgets(ctx context.Context, in *MutateCampaignBudgetsRequest, opts ...grpc.CallOption) (*MutateCampaignBudgetsResponse, error)
}

type campaignBudgetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignBudgetServiceClient(cc grpc.ClientConnInterface) CampaignBudgetServiceClient {
	return &campaignBudgetServiceClient{cc}
}

func (c *campaignBudgetServiceClient) MutateCampaignBudgets(ctx context.Context, in *MutateCampaignBudgetsRequest, opts ...grpc.CallOption) (*MutateCampaignBudgetsResponse, error) {
	out := new(MutateCampaignBudgetsResponse)
	err := c.cc.Invoke(ctx, CampaignBudgetService_MutateCampaignBudgets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignBudgetServiceServer is the server API for CampaignBudgetService service.
// All implementations must embed UnimplementedCampaignBudgetServiceServer
// for forward compatibility
type CampaignBudgetServiceServer interface {
	// Creates, updates, or removes campaign budgets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CampaignBudgetError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[OperationAccessDeniedError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[StringLengthError]()
	MutateCampaignBudgets(context.Context, *MutateCampaignBudgetsRequest) (*MutateCampaignBudgetsResponse, error)
	mustEmbedUnimplementedCampaignBudgetServiceServer()
}

// UnimplementedCampaignBudgetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCampaignBudgetServiceServer struct {
}

func (UnimplementedCampaignBudgetServiceServer) MutateCampaignBudgets(context.Context, *MutateCampaignBudgetsRequest) (*MutateCampaignBudgetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCampaignBudgets not implemented")
}
func (UnimplementedCampaignBudgetServiceServer) mustEmbedUnimplementedCampaignBudgetServiceServer() {}

// UnsafeCampaignBudgetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampaignBudgetServiceServer will
// result in compilation errors.
type UnsafeCampaignBudgetServiceServer interface {
	mustEmbedUnimplementedCampaignBudgetServiceServer()
}

func RegisterCampaignBudgetServiceServer(s grpc.ServiceRegistrar, srv CampaignBudgetServiceServer) {
	s.RegisterService(&CampaignBudgetService_ServiceDesc, srv)
}

func _CampaignBudgetService_MutateCampaignBudgets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignBudgetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignBudgetServiceServer).MutateCampaignBudgets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CampaignBudgetService_MutateCampaignBudgets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignBudgetServiceServer).MutateCampaignBudgets(ctx, req.(*MutateCampaignBudgetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CampaignBudgetService_ServiceDesc is the grpc.ServiceDesc for CampaignBudgetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CampaignBudgetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.CampaignBudgetService",
	HandlerType: (*CampaignBudgetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateCampaignBudgets",
			Handler:    _CampaignBudgetService_MutateCampaignBudgets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/campaign_budget_service.proto",
}