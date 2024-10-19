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
// source: google/ads/googleads/v18/services/conversion_value_rule_service.proto

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
	ConversionValueRuleService_MutateConversionValueRules_FullMethodName = "/google.ads.googleads.v18.services.ConversionValueRuleService/MutateConversionValueRules"
)

// ConversionValueRuleServiceClient is the client API for ConversionValueRuleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage conversion value rules.
type ConversionValueRuleServiceClient interface {
	// Creates, updates, or removes conversion value rules. Operation statuses are
	// returned.
	MutateConversionValueRules(ctx context.Context, in *MutateConversionValueRulesRequest, opts ...grpc.CallOption) (*MutateConversionValueRulesResponse, error)
}

type conversionValueRuleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversionValueRuleServiceClient(cc grpc.ClientConnInterface) ConversionValueRuleServiceClient {
	return &conversionValueRuleServiceClient{cc}
}

func (c *conversionValueRuleServiceClient) MutateConversionValueRules(ctx context.Context, in *MutateConversionValueRulesRequest, opts ...grpc.CallOption) (*MutateConversionValueRulesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateConversionValueRulesResponse)
	err := c.cc.Invoke(ctx, ConversionValueRuleService_MutateConversionValueRules_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionValueRuleServiceServer is the server API for ConversionValueRuleService service.
// All implementations must embed UnimplementedConversionValueRuleServiceServer
// for forward compatibility.
//
// Service to manage conversion value rules.
type ConversionValueRuleServiceServer interface {
	// Creates, updates, or removes conversion value rules. Operation statuses are
	// returned.
	MutateConversionValueRules(context.Context, *MutateConversionValueRulesRequest) (*MutateConversionValueRulesResponse, error)
	mustEmbedUnimplementedConversionValueRuleServiceServer()
}

// UnimplementedConversionValueRuleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConversionValueRuleServiceServer struct{}

func (UnimplementedConversionValueRuleServiceServer) MutateConversionValueRules(context.Context, *MutateConversionValueRulesRequest) (*MutateConversionValueRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateConversionValueRules not implemented")
}
func (UnimplementedConversionValueRuleServiceServer) mustEmbedUnimplementedConversionValueRuleServiceServer() {
}
func (UnimplementedConversionValueRuleServiceServer) testEmbeddedByValue() {}

// UnsafeConversionValueRuleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversionValueRuleServiceServer will
// result in compilation errors.
type UnsafeConversionValueRuleServiceServer interface {
	mustEmbedUnimplementedConversionValueRuleServiceServer()
}

func RegisterConversionValueRuleServiceServer(s grpc.ServiceRegistrar, srv ConversionValueRuleServiceServer) {
	// If the following call pancis, it indicates UnimplementedConversionValueRuleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ConversionValueRuleService_ServiceDesc, srv)
}

func _ConversionValueRuleService_MutateConversionValueRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateConversionValueRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionValueRuleServiceServer).MutateConversionValueRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConversionValueRuleService_MutateConversionValueRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionValueRuleServiceServer).MutateConversionValueRules(ctx, req.(*MutateConversionValueRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConversionValueRuleService_ServiceDesc is the grpc.ServiceDesc for ConversionValueRuleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConversionValueRuleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v18.services.ConversionValueRuleService",
	HandlerType: (*ConversionValueRuleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateConversionValueRules",
			Handler:    _ConversionValueRuleService_MutateConversionValueRules_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v18/services/conversion_value_rule_service.proto",
}
