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
// source: google/ads/googleads/v18/services/user_list_customer_type_service.proto

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
	UserListCustomerTypeService_MutateUserListCustomerTypes_FullMethodName = "/google.ads.googleads.v18.services.UserListCustomerTypeService/MutateUserListCustomerTypes"
)

// UserListCustomerTypeServiceClient is the client API for UserListCustomerTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage user list customer types.
type UserListCustomerTypeServiceClient interface {
	// Attach or remove user list customer types. Operation statuses
	// are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[UserListCustomerTypeError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateUserListCustomerTypes(ctx context.Context, in *MutateUserListCustomerTypesRequest, opts ...grpc.CallOption) (*MutateUserListCustomerTypesResponse, error)
}

type userListCustomerTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserListCustomerTypeServiceClient(cc grpc.ClientConnInterface) UserListCustomerTypeServiceClient {
	return &userListCustomerTypeServiceClient{cc}
}

func (c *userListCustomerTypeServiceClient) MutateUserListCustomerTypes(ctx context.Context, in *MutateUserListCustomerTypesRequest, opts ...grpc.CallOption) (*MutateUserListCustomerTypesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateUserListCustomerTypesResponse)
	err := c.cc.Invoke(ctx, UserListCustomerTypeService_MutateUserListCustomerTypes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserListCustomerTypeServiceServer is the server API for UserListCustomerTypeService service.
// All implementations must embed UnimplementedUserListCustomerTypeServiceServer
// for forward compatibility.
//
// Service to manage user list customer types.
type UserListCustomerTypeServiceServer interface {
	// Attach or remove user list customer types. Operation statuses
	// are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[UserListCustomerTypeError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateUserListCustomerTypes(context.Context, *MutateUserListCustomerTypesRequest) (*MutateUserListCustomerTypesResponse, error)
	mustEmbedUnimplementedUserListCustomerTypeServiceServer()
}

// UnimplementedUserListCustomerTypeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserListCustomerTypeServiceServer struct{}

func (UnimplementedUserListCustomerTypeServiceServer) MutateUserListCustomerTypes(context.Context, *MutateUserListCustomerTypesRequest) (*MutateUserListCustomerTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateUserListCustomerTypes not implemented")
}
func (UnimplementedUserListCustomerTypeServiceServer) mustEmbedUnimplementedUserListCustomerTypeServiceServer() {
}
func (UnimplementedUserListCustomerTypeServiceServer) testEmbeddedByValue() {}

// UnsafeUserListCustomerTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserListCustomerTypeServiceServer will
// result in compilation errors.
type UnsafeUserListCustomerTypeServiceServer interface {
	mustEmbedUnimplementedUserListCustomerTypeServiceServer()
}

func RegisterUserListCustomerTypeServiceServer(s grpc.ServiceRegistrar, srv UserListCustomerTypeServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserListCustomerTypeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserListCustomerTypeService_ServiceDesc, srv)
}

func _UserListCustomerTypeService_MutateUserListCustomerTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateUserListCustomerTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserListCustomerTypeServiceServer).MutateUserListCustomerTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserListCustomerTypeService_MutateUserListCustomerTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserListCustomerTypeServiceServer).MutateUserListCustomerTypes(ctx, req.(*MutateUserListCustomerTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserListCustomerTypeService_ServiceDesc is the grpc.ServiceDesc for UserListCustomerTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserListCustomerTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v18.services.UserListCustomerTypeService",
	HandlerType: (*UserListCustomerTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateUserListCustomerTypes",
			Handler:    _UserListCustomerTypeService_MutateUserListCustomerTypes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v18/services/user_list_customer_type_service.proto",
}
