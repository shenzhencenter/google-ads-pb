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
// source: google/ads/googleads/v19/services/shared_set_service.proto

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
	SharedSetService_MutateSharedSets_FullMethodName = "/google.ads.googleads.v19.services.SharedSetService/MutateSharedSets"
)

// SharedSetServiceClient is the client API for SharedSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage shared sets.
type SharedSetServiceClient interface {
	// Creates, updates, or removes shared sets. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SharedSetError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateSharedSets(ctx context.Context, in *MutateSharedSetsRequest, opts ...grpc.CallOption) (*MutateSharedSetsResponse, error)
}

type sharedSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSharedSetServiceClient(cc grpc.ClientConnInterface) SharedSetServiceClient {
	return &sharedSetServiceClient{cc}
}

func (c *sharedSetServiceClient) MutateSharedSets(ctx context.Context, in *MutateSharedSetsRequest, opts ...grpc.CallOption) (*MutateSharedSetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateSharedSetsResponse)
	err := c.cc.Invoke(ctx, SharedSetService_MutateSharedSets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharedSetServiceServer is the server API for SharedSetService service.
// All implementations must embed UnimplementedSharedSetServiceServer
// for forward compatibility.
//
// Service to manage shared sets.
type SharedSetServiceServer interface {
	// Creates, updates, or removes shared sets. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[IdError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotEmptyError]()
	//	[NullError]()
	//	[OperatorError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	//	[SharedSetError]()
	//	[SizeLimitError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	MutateSharedSets(context.Context, *MutateSharedSetsRequest) (*MutateSharedSetsResponse, error)
	mustEmbedUnimplementedSharedSetServiceServer()
}

// UnimplementedSharedSetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSharedSetServiceServer struct{}

func (UnimplementedSharedSetServiceServer) MutateSharedSets(context.Context, *MutateSharedSetsRequest) (*MutateSharedSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateSharedSets not implemented")
}
func (UnimplementedSharedSetServiceServer) mustEmbedUnimplementedSharedSetServiceServer() {}
func (UnimplementedSharedSetServiceServer) testEmbeddedByValue()                          {}

// UnsafeSharedSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SharedSetServiceServer will
// result in compilation errors.
type UnsafeSharedSetServiceServer interface {
	mustEmbedUnimplementedSharedSetServiceServer()
}

func RegisterSharedSetServiceServer(s grpc.ServiceRegistrar, srv SharedSetServiceServer) {
	// If the following call pancis, it indicates UnimplementedSharedSetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SharedSetService_ServiceDesc, srv)
}

func _SharedSetService_MutateSharedSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateSharedSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedSetServiceServer).MutateSharedSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SharedSetService_MutateSharedSets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedSetServiceServer).MutateSharedSets(ctx, req.(*MutateSharedSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SharedSetService_ServiceDesc is the grpc.ServiceDesc for SharedSetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SharedSetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.SharedSetService",
	HandlerType: (*SharedSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateSharedSets",
			Handler:    _SharedSetService_MutateSharedSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/shared_set_service.proto",
}
