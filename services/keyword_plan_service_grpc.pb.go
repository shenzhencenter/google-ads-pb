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
// source: google/ads/googleads/v19/services/keyword_plan_service.proto

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
	KeywordPlanService_MutateKeywordPlans_FullMethodName = "/google.ads.googleads.v19.services.KeywordPlanService/MutateKeywordPlans"
)

// KeywordPlanServiceClient is the client API for KeywordPlanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage keyword plans.
type KeywordPlanServiceClient interface {
	// Creates, updates, or removes keyword plans. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[KeywordPlanError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[StringLengthError]()
	MutateKeywordPlans(ctx context.Context, in *MutateKeywordPlansRequest, opts ...grpc.CallOption) (*MutateKeywordPlansResponse, error)
}

type keywordPlanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanServiceClient(cc grpc.ClientConnInterface) KeywordPlanServiceClient {
	return &keywordPlanServiceClient{cc}
}

func (c *keywordPlanServiceClient) MutateKeywordPlans(ctx context.Context, in *MutateKeywordPlansRequest, opts ...grpc.CallOption) (*MutateKeywordPlansResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateKeywordPlansResponse)
	err := c.cc.Invoke(ctx, KeywordPlanService_MutateKeywordPlans_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanServiceServer is the server API for KeywordPlanService service.
// All implementations must embed UnimplementedKeywordPlanServiceServer
// for forward compatibility.
//
// Service to manage keyword plans.
type KeywordPlanServiceServer interface {
	// Creates, updates, or removes keyword plans. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[KeywordPlanError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[StringLengthError]()
	MutateKeywordPlans(context.Context, *MutateKeywordPlansRequest) (*MutateKeywordPlansResponse, error)
	mustEmbedUnimplementedKeywordPlanServiceServer()
}

// UnimplementedKeywordPlanServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKeywordPlanServiceServer struct{}

func (UnimplementedKeywordPlanServiceServer) MutateKeywordPlans(context.Context, *MutateKeywordPlansRequest) (*MutateKeywordPlansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateKeywordPlans not implemented")
}
func (UnimplementedKeywordPlanServiceServer) mustEmbedUnimplementedKeywordPlanServiceServer() {}
func (UnimplementedKeywordPlanServiceServer) testEmbeddedByValue()                            {}

// UnsafeKeywordPlanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeywordPlanServiceServer will
// result in compilation errors.
type UnsafeKeywordPlanServiceServer interface {
	mustEmbedUnimplementedKeywordPlanServiceServer()
}

func RegisterKeywordPlanServiceServer(s grpc.ServiceRegistrar, srv KeywordPlanServiceServer) {
	// If the following call pancis, it indicates UnimplementedKeywordPlanServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KeywordPlanService_ServiceDesc, srv)
}

func _KeywordPlanService_MutateKeywordPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanServiceServer).MutateKeywordPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeywordPlanService_MutateKeywordPlans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanServiceServer).MutateKeywordPlans(ctx, req.(*MutateKeywordPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeywordPlanService_ServiceDesc is the grpc.ServiceDesc for KeywordPlanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeywordPlanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.KeywordPlanService",
	HandlerType: (*KeywordPlanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateKeywordPlans",
			Handler:    _KeywordPlanService_MutateKeywordPlans_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/keyword_plan_service.proto",
}
