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
// source: google/ads/googleads/v17/services/local_services_lead_service.proto

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	LocalServicesLeadService_AppendLeadConversation_FullMethodName = "/google.ads.googleads.v17.services.LocalServicesLeadService/AppendLeadConversation"
)

// LocalServicesLeadServiceClient is the client API for LocalServicesLeadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// This service allows management of LocalServicesLead resources.
type LocalServicesLeadServiceClient interface {
	// RPC to append Local Services Lead Conversation resources to Local Services
	// Lead resources.
	AppendLeadConversation(ctx context.Context, in *AppendLeadConversationRequest, opts ...grpc.CallOption) (*AppendLeadConversationResponse, error)
}

type localServicesLeadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalServicesLeadServiceClient(cc grpc.ClientConnInterface) LocalServicesLeadServiceClient {
	return &localServicesLeadServiceClient{cc}
}

func (c *localServicesLeadServiceClient) AppendLeadConversation(ctx context.Context, in *AppendLeadConversationRequest, opts ...grpc.CallOption) (*AppendLeadConversationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AppendLeadConversationResponse)
	err := c.cc.Invoke(ctx, LocalServicesLeadService_AppendLeadConversation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocalServicesLeadServiceServer is the server API for LocalServicesLeadService service.
// All implementations must embed UnimplementedLocalServicesLeadServiceServer
// for forward compatibility
//
// This service allows management of LocalServicesLead resources.
type LocalServicesLeadServiceServer interface {
	// RPC to append Local Services Lead Conversation resources to Local Services
	// Lead resources.
	AppendLeadConversation(context.Context, *AppendLeadConversationRequest) (*AppendLeadConversationResponse, error)
	mustEmbedUnimplementedLocalServicesLeadServiceServer()
}

// UnimplementedLocalServicesLeadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocalServicesLeadServiceServer struct {
}

func (UnimplementedLocalServicesLeadServiceServer) AppendLeadConversation(context.Context, *AppendLeadConversationRequest) (*AppendLeadConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendLeadConversation not implemented")
}
func (UnimplementedLocalServicesLeadServiceServer) mustEmbedUnimplementedLocalServicesLeadServiceServer() {
}

// UnsafeLocalServicesLeadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocalServicesLeadServiceServer will
// result in compilation errors.
type UnsafeLocalServicesLeadServiceServer interface {
	mustEmbedUnimplementedLocalServicesLeadServiceServer()
}

func RegisterLocalServicesLeadServiceServer(s grpc.ServiceRegistrar, srv LocalServicesLeadServiceServer) {
	s.RegisterService(&LocalServicesLeadService_ServiceDesc, srv)
}

func _LocalServicesLeadService_AppendLeadConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendLeadConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalServicesLeadServiceServer).AppendLeadConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocalServicesLeadService_AppendLeadConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalServicesLeadServiceServer).AppendLeadConversation(ctx, req.(*AppendLeadConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocalServicesLeadService_ServiceDesc is the grpc.ServiceDesc for LocalServicesLeadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocalServicesLeadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.LocalServicesLeadService",
	HandlerType: (*LocalServicesLeadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendLeadConversation",
			Handler:    _LocalServicesLeadService_AppendLeadConversation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/local_services_lead_service.proto",
}
