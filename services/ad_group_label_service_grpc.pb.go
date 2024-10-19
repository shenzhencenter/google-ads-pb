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
// source: google/ads/googleads/v18/services/ad_group_label_service.proto

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
	AdGroupLabelService_MutateAdGroupLabels_FullMethodName = "/google.ads.googleads.v18.services.AdGroupLabelService/MutateAdGroupLabels"
)

// AdGroupLabelServiceClient is the client API for AdGroupLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage labels on ad groups.
type AdGroupLabelServiceClient interface {
	// Creates and removes ad group labels.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[LabelError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateAdGroupLabels(ctx context.Context, in *MutateAdGroupLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupLabelsResponse, error)
}

type adGroupLabelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupLabelServiceClient(cc grpc.ClientConnInterface) AdGroupLabelServiceClient {
	return &adGroupLabelServiceClient{cc}
}

func (c *adGroupLabelServiceClient) MutateAdGroupLabels(ctx context.Context, in *MutateAdGroupLabelsRequest, opts ...grpc.CallOption) (*MutateAdGroupLabelsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateAdGroupLabelsResponse)
	err := c.cc.Invoke(ctx, AdGroupLabelService_MutateAdGroupLabels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupLabelServiceServer is the server API for AdGroupLabelService service.
// All implementations must embed UnimplementedAdGroupLabelServiceServer
// for forward compatibility.
//
// Service to manage labels on ad groups.
type AdGroupLabelServiceServer interface {
	// Creates and removes ad group labels.
	// Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[LabelError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateAdGroupLabels(context.Context, *MutateAdGroupLabelsRequest) (*MutateAdGroupLabelsResponse, error)
	mustEmbedUnimplementedAdGroupLabelServiceServer()
}

// UnimplementedAdGroupLabelServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdGroupLabelServiceServer struct{}

func (UnimplementedAdGroupLabelServiceServer) MutateAdGroupLabels(context.Context, *MutateAdGroupLabelsRequest) (*MutateAdGroupLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAdGroupLabels not implemented")
}
func (UnimplementedAdGroupLabelServiceServer) mustEmbedUnimplementedAdGroupLabelServiceServer() {}
func (UnimplementedAdGroupLabelServiceServer) testEmbeddedByValue()                             {}

// UnsafeAdGroupLabelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdGroupLabelServiceServer will
// result in compilation errors.
type UnsafeAdGroupLabelServiceServer interface {
	mustEmbedUnimplementedAdGroupLabelServiceServer()
}

func RegisterAdGroupLabelServiceServer(s grpc.ServiceRegistrar, srv AdGroupLabelServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdGroupLabelServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdGroupLabelService_ServiceDesc, srv)
}

func _AdGroupLabelService_MutateAdGroupLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupLabelServiceServer).MutateAdGroupLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdGroupLabelService_MutateAdGroupLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupLabelServiceServer).MutateAdGroupLabels(ctx, req.(*MutateAdGroupLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdGroupLabelService_ServiceDesc is the grpc.ServiceDesc for AdGroupLabelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdGroupLabelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v18.services.AdGroupLabelService",
	HandlerType: (*AdGroupLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAdGroupLabels",
			Handler:    _AdGroupLabelService_MutateAdGroupLabels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v18/services/ad_group_label_service.proto",
}
