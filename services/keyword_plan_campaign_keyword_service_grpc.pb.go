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
// source: google/ads/googleads/v19/services/keyword_plan_campaign_keyword_service.proto

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
	KeywordPlanCampaignKeywordService_MutateKeywordPlanCampaignKeywords_FullMethodName = "/google.ads.googleads.v19.services.KeywordPlanCampaignKeywordService/MutateKeywordPlanCampaignKeywords"
)

// KeywordPlanCampaignKeywordServiceClient is the client API for KeywordPlanCampaignKeywordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage Keyword Plan campaign keywords. KeywordPlanCampaign is
// required to add the campaign keywords. Only negative keywords are supported.
// A maximum of 1000 negative keywords are allowed per plan. This includes both
// campaign negative keywords and ad group negative keywords.
type KeywordPlanCampaignKeywordServiceClient interface {
	// Creates, updates, or removes Keyword Plan campaign keywords. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[KeywordPlanAdGroupKeywordError]()
	//	[KeywordPlanCampaignKeywordError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	MutateKeywordPlanCampaignKeywords(ctx context.Context, in *MutateKeywordPlanCampaignKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanCampaignKeywordsResponse, error)
}

type keywordPlanCampaignKeywordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanCampaignKeywordServiceClient(cc grpc.ClientConnInterface) KeywordPlanCampaignKeywordServiceClient {
	return &keywordPlanCampaignKeywordServiceClient{cc}
}

func (c *keywordPlanCampaignKeywordServiceClient) MutateKeywordPlanCampaignKeywords(ctx context.Context, in *MutateKeywordPlanCampaignKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanCampaignKeywordsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateKeywordPlanCampaignKeywordsResponse)
	err := c.cc.Invoke(ctx, KeywordPlanCampaignKeywordService_MutateKeywordPlanCampaignKeywords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanCampaignKeywordServiceServer is the server API for KeywordPlanCampaignKeywordService service.
// All implementations must embed UnimplementedKeywordPlanCampaignKeywordServiceServer
// for forward compatibility.
//
// Service to manage Keyword Plan campaign keywords. KeywordPlanCampaign is
// required to add the campaign keywords. Only negative keywords are supported.
// A maximum of 1000 negative keywords are allowed per plan. This includes both
// campaign negative keywords and ad group negative keywords.
type KeywordPlanCampaignKeywordServiceServer interface {
	// Creates, updates, or removes Keyword Plan campaign keywords. Operation
	// statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[KeywordPlanAdGroupKeywordError]()
	//	[KeywordPlanCampaignKeywordError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	MutateKeywordPlanCampaignKeywords(context.Context, *MutateKeywordPlanCampaignKeywordsRequest) (*MutateKeywordPlanCampaignKeywordsResponse, error)
	mustEmbedUnimplementedKeywordPlanCampaignKeywordServiceServer()
}

// UnimplementedKeywordPlanCampaignKeywordServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKeywordPlanCampaignKeywordServiceServer struct{}

func (UnimplementedKeywordPlanCampaignKeywordServiceServer) MutateKeywordPlanCampaignKeywords(context.Context, *MutateKeywordPlanCampaignKeywordsRequest) (*MutateKeywordPlanCampaignKeywordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateKeywordPlanCampaignKeywords not implemented")
}
func (UnimplementedKeywordPlanCampaignKeywordServiceServer) mustEmbedUnimplementedKeywordPlanCampaignKeywordServiceServer() {
}
func (UnimplementedKeywordPlanCampaignKeywordServiceServer) testEmbeddedByValue() {}

// UnsafeKeywordPlanCampaignKeywordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeywordPlanCampaignKeywordServiceServer will
// result in compilation errors.
type UnsafeKeywordPlanCampaignKeywordServiceServer interface {
	mustEmbedUnimplementedKeywordPlanCampaignKeywordServiceServer()
}

func RegisterKeywordPlanCampaignKeywordServiceServer(s grpc.ServiceRegistrar, srv KeywordPlanCampaignKeywordServiceServer) {
	// If the following call pancis, it indicates UnimplementedKeywordPlanCampaignKeywordServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KeywordPlanCampaignKeywordService_ServiceDesc, srv)
}

func _KeywordPlanCampaignKeywordService_MutateKeywordPlanCampaignKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanCampaignKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanCampaignKeywordServiceServer).MutateKeywordPlanCampaignKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeywordPlanCampaignKeywordService_MutateKeywordPlanCampaignKeywords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanCampaignKeywordServiceServer).MutateKeywordPlanCampaignKeywords(ctx, req.(*MutateKeywordPlanCampaignKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeywordPlanCampaignKeywordService_ServiceDesc is the grpc.ServiceDesc for KeywordPlanCampaignKeywordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeywordPlanCampaignKeywordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.KeywordPlanCampaignKeywordService",
	HandlerType: (*KeywordPlanCampaignKeywordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateKeywordPlanCampaignKeywords",
			Handler:    _KeywordPlanCampaignKeywordService_MutateKeywordPlanCampaignKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/keyword_plan_campaign_keyword_service.proto",
}
