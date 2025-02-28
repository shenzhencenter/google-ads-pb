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
// source: google/ads/googleads/v19/services/keyword_theme_constant_service.proto

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
	KeywordThemeConstantService_SuggestKeywordThemeConstants_FullMethodName = "/google.ads.googleads.v19.services.KeywordThemeConstantService/SuggestKeywordThemeConstants"
)

// KeywordThemeConstantServiceClient is the client API for KeywordThemeConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to fetch Smart Campaign keyword themes.
type KeywordThemeConstantServiceClient interface {
	// Returns KeywordThemeConstant suggestions by keyword themes.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	SuggestKeywordThemeConstants(ctx context.Context, in *SuggestKeywordThemeConstantsRequest, opts ...grpc.CallOption) (*SuggestKeywordThemeConstantsResponse, error)
}

type keywordThemeConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordThemeConstantServiceClient(cc grpc.ClientConnInterface) KeywordThemeConstantServiceClient {
	return &keywordThemeConstantServiceClient{cc}
}

func (c *keywordThemeConstantServiceClient) SuggestKeywordThemeConstants(ctx context.Context, in *SuggestKeywordThemeConstantsRequest, opts ...grpc.CallOption) (*SuggestKeywordThemeConstantsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SuggestKeywordThemeConstantsResponse)
	err := c.cc.Invoke(ctx, KeywordThemeConstantService_SuggestKeywordThemeConstants_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordThemeConstantServiceServer is the server API for KeywordThemeConstantService service.
// All implementations must embed UnimplementedKeywordThemeConstantServiceServer
// for forward compatibility.
//
// Service to fetch Smart Campaign keyword themes.
type KeywordThemeConstantServiceServer interface {
	// Returns KeywordThemeConstant suggestions by keyword themes.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	SuggestKeywordThemeConstants(context.Context, *SuggestKeywordThemeConstantsRequest) (*SuggestKeywordThemeConstantsResponse, error)
	mustEmbedUnimplementedKeywordThemeConstantServiceServer()
}

// UnimplementedKeywordThemeConstantServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKeywordThemeConstantServiceServer struct{}

func (UnimplementedKeywordThemeConstantServiceServer) SuggestKeywordThemeConstants(context.Context, *SuggestKeywordThemeConstantsRequest) (*SuggestKeywordThemeConstantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuggestKeywordThemeConstants not implemented")
}
func (UnimplementedKeywordThemeConstantServiceServer) mustEmbedUnimplementedKeywordThemeConstantServiceServer() {
}
func (UnimplementedKeywordThemeConstantServiceServer) testEmbeddedByValue() {}

// UnsafeKeywordThemeConstantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeywordThemeConstantServiceServer will
// result in compilation errors.
type UnsafeKeywordThemeConstantServiceServer interface {
	mustEmbedUnimplementedKeywordThemeConstantServiceServer()
}

func RegisterKeywordThemeConstantServiceServer(s grpc.ServiceRegistrar, srv KeywordThemeConstantServiceServer) {
	// If the following call pancis, it indicates UnimplementedKeywordThemeConstantServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KeywordThemeConstantService_ServiceDesc, srv)
}

func _KeywordThemeConstantService_SuggestKeywordThemeConstants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestKeywordThemeConstantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordThemeConstantServiceServer).SuggestKeywordThemeConstants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeywordThemeConstantService_SuggestKeywordThemeConstants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordThemeConstantServiceServer).SuggestKeywordThemeConstants(ctx, req.(*SuggestKeywordThemeConstantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeywordThemeConstantService_ServiceDesc is the grpc.ServiceDesc for KeywordThemeConstantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeywordThemeConstantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.KeywordThemeConstantService",
	HandlerType: (*KeywordThemeConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuggestKeywordThemeConstants",
			Handler:    _KeywordThemeConstantService_SuggestKeywordThemeConstants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/keyword_theme_constant_service.proto",
}
