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
// source: google/ads/googleads/v19/services/user_data_service.proto

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
	UserDataService_UploadUserData_FullMethodName = "/google.ads.googleads.v19.services.UserDataService/UploadUserData"
)

// UserDataServiceClient is the client API for UserDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage user data uploads.
// Any uploads made to a Customer Match list through this service will be
// eligible for matching as per the customer matching process. See
// https://support.google.com/google-ads/answer/7474263. However, the uploads
// made through this service will not be visible under the 'Segment members'
// section for the Customer Match List in the Google Ads UI.
type UserDataServiceClient interface {
	// Uploads the given user data.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[OfflineUserDataJobError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[UserDataError]()
	UploadUserData(ctx context.Context, in *UploadUserDataRequest, opts ...grpc.CallOption) (*UploadUserDataResponse, error)
}

type userDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserDataServiceClient(cc grpc.ClientConnInterface) UserDataServiceClient {
	return &userDataServiceClient{cc}
}

func (c *userDataServiceClient) UploadUserData(ctx context.Context, in *UploadUserDataRequest, opts ...grpc.CallOption) (*UploadUserDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadUserDataResponse)
	err := c.cc.Invoke(ctx, UserDataService_UploadUserData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserDataServiceServer is the server API for UserDataService service.
// All implementations must embed UnimplementedUserDataServiceServer
// for forward compatibility.
//
// Service to manage user data uploads.
// Any uploads made to a Customer Match list through this service will be
// eligible for matching as per the customer matching process. See
// https://support.google.com/google-ads/answer/7474263. However, the uploads
// made through this service will not be visible under the 'Segment members'
// section for the Customer Match List in the Google Ads UI.
type UserDataServiceServer interface {
	// Uploads the given user data.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[CollectionSizeError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[OfflineUserDataJobError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[UserDataError]()
	UploadUserData(context.Context, *UploadUserDataRequest) (*UploadUserDataResponse, error)
	mustEmbedUnimplementedUserDataServiceServer()
}

// UnimplementedUserDataServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserDataServiceServer struct{}

func (UnimplementedUserDataServiceServer) UploadUserData(context.Context, *UploadUserDataRequest) (*UploadUserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadUserData not implemented")
}
func (UnimplementedUserDataServiceServer) mustEmbedUnimplementedUserDataServiceServer() {}
func (UnimplementedUserDataServiceServer) testEmbeddedByValue()                         {}

// UnsafeUserDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserDataServiceServer will
// result in compilation errors.
type UnsafeUserDataServiceServer interface {
	mustEmbedUnimplementedUserDataServiceServer()
}

func RegisterUserDataServiceServer(s grpc.ServiceRegistrar, srv UserDataServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserDataServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserDataService_ServiceDesc, srv)
}

func _UserDataService_UploadUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServiceServer).UploadUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserDataService_UploadUserData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServiceServer).UploadUserData(ctx, req.(*UploadUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserDataService_ServiceDesc is the grpc.ServiceDesc for UserDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.UserDataService",
	HandlerType: (*UserDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadUserData",
			Handler:    _UserDataService_UploadUserData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/user_data_service.proto",
}
