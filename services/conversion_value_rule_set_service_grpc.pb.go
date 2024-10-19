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
// source: google/ads/googleads/v18/services/conversion_value_rule_set_service.proto

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
	ConversionValueRuleSetService_MutateConversionValueRuleSets_FullMethodName = "/google.ads.googleads.v18.services.ConversionValueRuleSetService/MutateConversionValueRuleSets"
)

// ConversionValueRuleSetServiceClient is the client API for ConversionValueRuleSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage conversion value rule sets.
type ConversionValueRuleSetServiceClient interface {
	// Creates, updates or removes conversion value rule sets. Operation statuses
	// are returned.
	MutateConversionValueRuleSets(ctx context.Context, in *MutateConversionValueRuleSetsRequest, opts ...grpc.CallOption) (*MutateConversionValueRuleSetsResponse, error)
}

type conversionValueRuleSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversionValueRuleSetServiceClient(cc grpc.ClientConnInterface) ConversionValueRuleSetServiceClient {
	return &conversionValueRuleSetServiceClient{cc}
}

func (c *conversionValueRuleSetServiceClient) MutateConversionValueRuleSets(ctx context.Context, in *MutateConversionValueRuleSetsRequest, opts ...grpc.CallOption) (*MutateConversionValueRuleSetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateConversionValueRuleSetsResponse)
	err := c.cc.Invoke(ctx, ConversionValueRuleSetService_MutateConversionValueRuleSets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionValueRuleSetServiceServer is the server API for ConversionValueRuleSetService service.
// All implementations must embed UnimplementedConversionValueRuleSetServiceServer
// for forward compatibility.
//
// Service to manage conversion value rule sets.
type ConversionValueRuleSetServiceServer interface {
	// Creates, updates or removes conversion value rule sets. Operation statuses
	// are returned.
	MutateConversionValueRuleSets(context.Context, *MutateConversionValueRuleSetsRequest) (*MutateConversionValueRuleSetsResponse, error)
	mustEmbedUnimplementedConversionValueRuleSetServiceServer()
}

// UnimplementedConversionValueRuleSetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConversionValueRuleSetServiceServer struct{}

func (UnimplementedConversionValueRuleSetServiceServer) MutateConversionValueRuleSets(context.Context, *MutateConversionValueRuleSetsRequest) (*MutateConversionValueRuleSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateConversionValueRuleSets not implemented")
}
func (UnimplementedConversionValueRuleSetServiceServer) mustEmbedUnimplementedConversionValueRuleSetServiceServer() {
}
func (UnimplementedConversionValueRuleSetServiceServer) testEmbeddedByValue() {}

// UnsafeConversionValueRuleSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversionValueRuleSetServiceServer will
// result in compilation errors.
type UnsafeConversionValueRuleSetServiceServer interface {
	mustEmbedUnimplementedConversionValueRuleSetServiceServer()
}

func RegisterConversionValueRuleSetServiceServer(s grpc.ServiceRegistrar, srv ConversionValueRuleSetServiceServer) {
	// If the following call pancis, it indicates UnimplementedConversionValueRuleSetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ConversionValueRuleSetService_ServiceDesc, srv)
}

func _ConversionValueRuleSetService_MutateConversionValueRuleSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateConversionValueRuleSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionValueRuleSetServiceServer).MutateConversionValueRuleSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConversionValueRuleSetService_MutateConversionValueRuleSets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionValueRuleSetServiceServer).MutateConversionValueRuleSets(ctx, req.(*MutateConversionValueRuleSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConversionValueRuleSetService_ServiceDesc is the grpc.ServiceDesc for ConversionValueRuleSetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConversionValueRuleSetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v18.services.ConversionValueRuleSetService",
	HandlerType: (*ConversionValueRuleSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateConversionValueRuleSets",
			Handler:    _ConversionValueRuleSetService_MutateConversionValueRuleSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v18/services/conversion_value_rule_set_service.proto",
}
