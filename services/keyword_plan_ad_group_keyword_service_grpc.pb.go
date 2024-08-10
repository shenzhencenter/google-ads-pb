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
// source: google/ads/googleads/v17/services/keyword_plan_ad_group_keyword_service.proto

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
	KeywordPlanAdGroupKeywordService_MutateKeywordPlanAdGroupKeywords_FullMethodName = "/google.ads.googleads.v17.services.KeywordPlanAdGroupKeywordService/MutateKeywordPlanAdGroupKeywords"
)

// KeywordPlanAdGroupKeywordServiceClient is the client API for KeywordPlanAdGroupKeywordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage Keyword Plan ad group keywords. KeywordPlanAdGroup is
// required to add ad group keywords. Positive and negative keywords are
// supported. A maximum of 10,000 positive keywords are allowed per keyword
// plan. A maximum of 1,000 negative keywords are allower per keyword plan. This
// includes campaign negative keywords and ad group negative keywords.
type KeywordPlanAdGroupKeywordServiceClient interface {
	// Creates, updates, or removes Keyword Plan ad group keywords. Operation
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
	//	[KeywordPlanError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	MutateKeywordPlanAdGroupKeywords(ctx context.Context, in *MutateKeywordPlanAdGroupKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanAdGroupKeywordsResponse, error)
}

type keywordPlanAdGroupKeywordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanAdGroupKeywordServiceClient(cc grpc.ClientConnInterface) KeywordPlanAdGroupKeywordServiceClient {
	return &keywordPlanAdGroupKeywordServiceClient{cc}
}

func (c *keywordPlanAdGroupKeywordServiceClient) MutateKeywordPlanAdGroupKeywords(ctx context.Context, in *MutateKeywordPlanAdGroupKeywordsRequest, opts ...grpc.CallOption) (*MutateKeywordPlanAdGroupKeywordsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateKeywordPlanAdGroupKeywordsResponse)
	err := c.cc.Invoke(ctx, KeywordPlanAdGroupKeywordService_MutateKeywordPlanAdGroupKeywords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanAdGroupKeywordServiceServer is the server API for KeywordPlanAdGroupKeywordService service.
// All implementations must embed UnimplementedKeywordPlanAdGroupKeywordServiceServer
// for forward compatibility.
//
// Service to manage Keyword Plan ad group keywords. KeywordPlanAdGroup is
// required to add ad group keywords. Positive and negative keywords are
// supported. A maximum of 10,000 positive keywords are allowed per keyword
// plan. A maximum of 1,000 negative keywords are allower per keyword plan. This
// includes campaign negative keywords and ad group negative keywords.
type KeywordPlanAdGroupKeywordServiceServer interface {
	// Creates, updates, or removes Keyword Plan ad group keywords. Operation
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
	//	[KeywordPlanError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	MutateKeywordPlanAdGroupKeywords(context.Context, *MutateKeywordPlanAdGroupKeywordsRequest) (*MutateKeywordPlanAdGroupKeywordsResponse, error)
	mustEmbedUnimplementedKeywordPlanAdGroupKeywordServiceServer()
}

// UnimplementedKeywordPlanAdGroupKeywordServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKeywordPlanAdGroupKeywordServiceServer struct{}

func (UnimplementedKeywordPlanAdGroupKeywordServiceServer) MutateKeywordPlanAdGroupKeywords(context.Context, *MutateKeywordPlanAdGroupKeywordsRequest) (*MutateKeywordPlanAdGroupKeywordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateKeywordPlanAdGroupKeywords not implemented")
}
func (UnimplementedKeywordPlanAdGroupKeywordServiceServer) mustEmbedUnimplementedKeywordPlanAdGroupKeywordServiceServer() {
}
func (UnimplementedKeywordPlanAdGroupKeywordServiceServer) testEmbeddedByValue() {}

// UnsafeKeywordPlanAdGroupKeywordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeywordPlanAdGroupKeywordServiceServer will
// result in compilation errors.
type UnsafeKeywordPlanAdGroupKeywordServiceServer interface {
	mustEmbedUnimplementedKeywordPlanAdGroupKeywordServiceServer()
}

func RegisterKeywordPlanAdGroupKeywordServiceServer(s grpc.ServiceRegistrar, srv KeywordPlanAdGroupKeywordServiceServer) {
	// If the following call pancis, it indicates UnimplementedKeywordPlanAdGroupKeywordServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KeywordPlanAdGroupKeywordService_ServiceDesc, srv)
}

func _KeywordPlanAdGroupKeywordService_MutateKeywordPlanAdGroupKeywords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlanAdGroupKeywordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanAdGroupKeywordServiceServer).MutateKeywordPlanAdGroupKeywords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeywordPlanAdGroupKeywordService_MutateKeywordPlanAdGroupKeywords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanAdGroupKeywordServiceServer).MutateKeywordPlanAdGroupKeywords(ctx, req.(*MutateKeywordPlanAdGroupKeywordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeywordPlanAdGroupKeywordService_ServiceDesc is the grpc.ServiceDesc for KeywordPlanAdGroupKeywordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeywordPlanAdGroupKeywordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.KeywordPlanAdGroupKeywordService",
	HandlerType: (*KeywordPlanAdGroupKeywordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateKeywordPlanAdGroupKeywords",
			Handler:    _KeywordPlanAdGroupKeywordService_MutateKeywordPlanAdGroupKeywords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/keyword_plan_ad_group_keyword_service.proto",
}
