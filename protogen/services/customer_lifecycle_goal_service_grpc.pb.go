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
// source: google/ads/googleads/v15/services/customer_lifecycle_goal_service.proto

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
	CustomerLifecycleService_ConfigureCustomerLifecycleGoals_FullMethodName = "/google.ads.googleads.v15.services.CustomerLifecycleService/ConfigureCustomerLifecycleGoals"
)

// CustomerLifecycleServiceClient is the client API for CustomerLifecycleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerLifecycleServiceClient interface {
	// Process the given customer lifecycle configurations.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CustomerLifecycleGoalConfigError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ConfigureCustomerLifecycleGoals(ctx context.Context, in *ConfigureCustomerLifecycleGoalsRequest, opts ...grpc.CallOption) (*ConfigureCustomerLifecycleGoalsResponse, error)
}

type customerLifecycleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerLifecycleServiceClient(cc grpc.ClientConnInterface) CustomerLifecycleServiceClient {
	return &customerLifecycleServiceClient{cc}
}

func (c *customerLifecycleServiceClient) ConfigureCustomerLifecycleGoals(ctx context.Context, in *ConfigureCustomerLifecycleGoalsRequest, opts ...grpc.CallOption) (*ConfigureCustomerLifecycleGoalsResponse, error) {
	out := new(ConfigureCustomerLifecycleGoalsResponse)
	err := c.cc.Invoke(ctx, CustomerLifecycleService_ConfigureCustomerLifecycleGoals_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerLifecycleServiceServer is the server API for CustomerLifecycleService service.
// All implementations must embed UnimplementedCustomerLifecycleServiceServer
// for forward compatibility
type CustomerLifecycleServiceServer interface {
	// Process the given customer lifecycle configurations.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CustomerLifecycleGoalConfigError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ConfigureCustomerLifecycleGoals(context.Context, *ConfigureCustomerLifecycleGoalsRequest) (*ConfigureCustomerLifecycleGoalsResponse, error)
	mustEmbedUnimplementedCustomerLifecycleServiceServer()
}

// UnimplementedCustomerLifecycleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerLifecycleServiceServer struct {
}

func (UnimplementedCustomerLifecycleServiceServer) ConfigureCustomerLifecycleGoals(context.Context, *ConfigureCustomerLifecycleGoalsRequest) (*ConfigureCustomerLifecycleGoalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureCustomerLifecycleGoals not implemented")
}
func (UnimplementedCustomerLifecycleServiceServer) mustEmbedUnimplementedCustomerLifecycleServiceServer() {
}

// UnsafeCustomerLifecycleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerLifecycleServiceServer will
// result in compilation errors.
type UnsafeCustomerLifecycleServiceServer interface {
	mustEmbedUnimplementedCustomerLifecycleServiceServer()
}

func RegisterCustomerLifecycleServiceServer(s grpc.ServiceRegistrar, srv CustomerLifecycleServiceServer) {
	s.RegisterService(&CustomerLifecycleService_ServiceDesc, srv)
}

func _CustomerLifecycleService_ConfigureCustomerLifecycleGoals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureCustomerLifecycleGoalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerLifecycleServiceServer).ConfigureCustomerLifecycleGoals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerLifecycleService_ConfigureCustomerLifecycleGoals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerLifecycleServiceServer).ConfigureCustomerLifecycleGoals(ctx, req.(*ConfigureCustomerLifecycleGoalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerLifecycleService_ServiceDesc is the grpc.ServiceDesc for CustomerLifecycleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerLifecycleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v15.services.CustomerLifecycleService",
	HandlerType: (*CustomerLifecycleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfigureCustomerLifecycleGoals",
			Handler:    _CustomerLifecycleService_ConfigureCustomerLifecycleGoals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v15/services/customer_lifecycle_goal_service.proto",
}