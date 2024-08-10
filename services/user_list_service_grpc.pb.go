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
// source: google/ads/googleads/v17/services/user_list_service.proto

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
	UserListService_MutateUserLists_FullMethodName = "/google.ads.googleads.v17.services.UserListService/MutateUserLists"
)

// UserListServiceClient is the client API for UserListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage user lists.
type UserListServiceClient interface {
	// Creates or updates user lists. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotAllowlistedError]()
	//	[NotEmptyError]()
	//	[OperationAccessDeniedError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UserListError]()
	MutateUserLists(ctx context.Context, in *MutateUserListsRequest, opts ...grpc.CallOption) (*MutateUserListsResponse, error)
}

type userListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserListServiceClient(cc grpc.ClientConnInterface) UserListServiceClient {
	return &userListServiceClient{cc}
}

func (c *userListServiceClient) MutateUserLists(ctx context.Context, in *MutateUserListsRequest, opts ...grpc.CallOption) (*MutateUserListsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateUserListsResponse)
	err := c.cc.Invoke(ctx, UserListService_MutateUserLists_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserListServiceServer is the server API for UserListService service.
// All implementations must embed UnimplementedUserListServiceServer
// for forward compatibility.
//
// Service to manage user lists.
type UserListServiceServer interface {
	// Creates or updates user lists. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[DatabaseError]()
	//	[DistinctError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[NewResourceCreationError]()
	//	[NotAllowlistedError]()
	//	[NotEmptyError]()
	//	[OperationAccessDeniedError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	//	[StringFormatError]()
	//	[StringLengthError]()
	//	[UserListError]()
	MutateUserLists(context.Context, *MutateUserListsRequest) (*MutateUserListsResponse, error)
	mustEmbedUnimplementedUserListServiceServer()
}

// UnimplementedUserListServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserListServiceServer struct{}

func (UnimplementedUserListServiceServer) MutateUserLists(context.Context, *MutateUserListsRequest) (*MutateUserListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateUserLists not implemented")
}
func (UnimplementedUserListServiceServer) mustEmbedUnimplementedUserListServiceServer() {}
func (UnimplementedUserListServiceServer) testEmbeddedByValue()                         {}

// UnsafeUserListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserListServiceServer will
// result in compilation errors.
type UnsafeUserListServiceServer interface {
	mustEmbedUnimplementedUserListServiceServer()
}

func RegisterUserListServiceServer(s grpc.ServiceRegistrar, srv UserListServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserListServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserListService_ServiceDesc, srv)
}

func _UserListService_MutateUserLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateUserListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserListServiceServer).MutateUserLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserListService_MutateUserLists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserListServiceServer).MutateUserLists(ctx, req.(*MutateUserListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserListService_ServiceDesc is the grpc.ServiceDesc for UserListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v17.services.UserListService",
	HandlerType: (*UserListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateUserLists",
			Handler:    _UserListService_MutateUserLists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v17/services/user_list_service.proto",
}
