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
// source: google/ads/googleads/v17/services/identity_verification_service.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	IdentityVerificationService_StartIdentityVerification_FullMethodName = "/google.ads.googleads.v17.services.IdentityVerificationService/StartIdentityVerification"
	IdentityVerificationService_GetIdentityVerification_FullMethodName   = "/google.ads.googleads.v17.services.IdentityVerificationService/GetIdentityVerification"
)

// IdentityVerificationServiceClient is the client API for IdentityVerificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A service for managing Identity Verification Service.
type IdentityVerificationServiceClient interface {
	// Starts Identity Verification for a given verification program type.
	//
	//	Statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	StartIdentityVerification(ctx context.Context, in *StartIdentityVerificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Returns Identity Verification information.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GetIdentityVerification(ctx context.Context, in *GetIdentityVerificationRequest, opts ...grpc.CallOption) (*GetIdentityVerificationResponse, error)
}

type identityVerificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityVerificationServiceClient(cc grpc.ClientConnInterface) IdentityVerificationServiceClient {
	return &identityVerificationServiceClient{cc}
}

func (c *identityVerificationServiceClient) StartIdentityVerification(ctx context.Context, in *StartIdentityVerificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IdentityVerificationService_StartIdentityVerification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityVerificationServiceClient) GetIdentityVerification(ctx context.Context, in *GetIdentityVerificationRequest, opts ...grpc.CallOption) (*GetIdentityVerificationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIdentityVerificationResponse)
	err := c.cc.Invoke(ctx, IdentityVerificationService_GetIdentityVerification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityVerificationServiceServer is the server API for IdentityVerificationService service.
// All implementations must embed UnimplementedIdentityVerificationServiceServer
// for forward compatibility
//
// A service for managing Identity Verification Service.
type IdentityVerificationServiceServer interface {
	// Starts Identity Verification for a given verification program type.
	//
	//	Statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	StartIdentityVerification(context.Context, *StartIdentityVerificationRequest) (*emptypb.Empty, error)
	// Returns Identity Verification information.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	GetIdentityVerification(context.Context, *GetIdentityVerificationRequest) (*GetIdentityVerificationResponse, error)
	mustEmbedUnimplementedIdentityVerificationServiceServer()
}

// UnimplementedIdentityVerificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityVerificationServiceServer struct {
}

func (UnimplementedIdentityVerificationServiceServer) StartIdentityVerification(context.Context, *StartIdentityVerificationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartIdentityVerification not implemented")
}
func (UnimplementedIdentityVerificationServiceServer) GetIdentityVerification(context.Context, *GetIdentityVerificationRequest) (*GetIdentityVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentityVerification not implemented")
}
func (UnimplementedIdentityVerificationServiceServer) mustEmbedUnimplementedIdentityVerificationServiceServer() {
}

// UnsafeIdentityVerificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityVerificationServiceServer will
// result in compilation errors.
type UnsafeIdentityVerificationServiceServer interface {
	mustEmbedUnimplementedIdentityVerificationServiceServer()
}

func RegisterIdentityVerificationServiceServer(s grpc.ServiceRegistrar, srv IdentityVerificationServiceServer) {
	s.RegisterService(&IdentityVerificationService_ServiceDesc, srv)
}

func _IdentityVerificationService_StartIdentityVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartIdentityVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityVerificationServiceServer).StartIdentityVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityVerificationService_StartIdentityVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityVerificationServiceServer).StartIdentityVerification(ctx, req.(*StartIdentityVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityVerificationService_GetIdentityVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdentityVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityVerificationServiceServer).GetIdentityVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityVerificationService_GetIdentityVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityVerificationServiceServer).GetIdentityVerification(ctx, req.(*GetIdentityVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityVerificationService_ServiceDesc is the grpc.ServiceDesc for IdentityVerificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityVerificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.IdentityVerificationService",
	HandlerType: (*IdentityVerificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartIdentityVerification",
			Handler:    _IdentityVerificationService_StartIdentityVerification_Handler,
		},
		{
			MethodName: "GetIdentityVerification",
			Handler:    _IdentityVerificationService_GetIdentityVerification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/identity_verification_service.proto",
}
