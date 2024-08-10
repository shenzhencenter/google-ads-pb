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
// source: google/ads/googleads/v17/services/remarketing_action_service.proto

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
	RemarketingActionService_MutateRemarketingActions_FullMethodName = "/google.ads.googleads.v17.services.RemarketingActionService/MutateRemarketingActions"
)

// RemarketingActionServiceClient is the client API for RemarketingActionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage remarketing actions.
type RemarketingActionServiceClient interface {
	// Creates or updates remarketing actions. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ConversionActionError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateRemarketingActions(ctx context.Context, in *MutateRemarketingActionsRequest, opts ...grpc.CallOption) (*MutateRemarketingActionsResponse, error)
}

type remarketingActionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRemarketingActionServiceClient(cc grpc.ClientConnInterface) RemarketingActionServiceClient {
	return &remarketingActionServiceClient{cc}
}

func (c *remarketingActionServiceClient) MutateRemarketingActions(ctx context.Context, in *MutateRemarketingActionsRequest, opts ...grpc.CallOption) (*MutateRemarketingActionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateRemarketingActionsResponse)
	err := c.cc.Invoke(ctx, RemarketingActionService_MutateRemarketingActions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemarketingActionServiceServer is the server API for RemarketingActionService service.
// All implementations must embed UnimplementedRemarketingActionServiceServer
// for forward compatibility.
//
// Service to manage remarketing actions.
type RemarketingActionServiceServer interface {
	// Creates or updates remarketing actions. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ConversionActionError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateRemarketingActions(context.Context, *MutateRemarketingActionsRequest) (*MutateRemarketingActionsResponse, error)
	mustEmbedUnimplementedRemarketingActionServiceServer()
}

// UnimplementedRemarketingActionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRemarketingActionServiceServer struct{}

func (UnimplementedRemarketingActionServiceServer) MutateRemarketingActions(context.Context, *MutateRemarketingActionsRequest) (*MutateRemarketingActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateRemarketingActions not implemented")
}
func (UnimplementedRemarketingActionServiceServer) mustEmbedUnimplementedRemarketingActionServiceServer() {
}
func (UnimplementedRemarketingActionServiceServer) testEmbeddedByValue() {}

// UnsafeRemarketingActionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemarketingActionServiceServer will
// result in compilation errors.
type UnsafeRemarketingActionServiceServer interface {
	mustEmbedUnimplementedRemarketingActionServiceServer()
}

func RegisterRemarketingActionServiceServer(s grpc.ServiceRegistrar, srv RemarketingActionServiceServer) {
	// If the following call pancis, it indicates UnimplementedRemarketingActionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RemarketingActionService_ServiceDesc, srv)
}

func _RemarketingActionService_MutateRemarketingActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateRemarketingActionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemarketingActionServiceServer).MutateRemarketingActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemarketingActionService_MutateRemarketingActions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemarketingActionServiceServer).MutateRemarketingActions(ctx, req.(*MutateRemarketingActionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RemarketingActionService_ServiceDesc is the grpc.ServiceDesc for RemarketingActionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemarketingActionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.RemarketingActionService",
	HandlerType: (*RemarketingActionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateRemarketingActions",
			Handler:    _RemarketingActionService_MutateRemarketingActions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/remarketing_action_service.proto",
}
