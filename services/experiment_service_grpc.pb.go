// Copyright 2025 Google LLC
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
// source: google/ads/googleads/v19/services/experiment_service.proto

package services

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExperimentService_MutateExperiments_FullMethodName         = "/google.ads.googleads.v19.services.ExperimentService/MutateExperiments"
	ExperimentService_EndExperiment_FullMethodName             = "/google.ads.googleads.v19.services.ExperimentService/EndExperiment"
	ExperimentService_ListExperimentAsyncErrors_FullMethodName = "/google.ads.googleads.v19.services.ExperimentService/ListExperimentAsyncErrors"
	ExperimentService_GraduateExperiment_FullMethodName        = "/google.ads.googleads.v19.services.ExperimentService/GraduateExperiment"
	ExperimentService_ScheduleExperiment_FullMethodName        = "/google.ads.googleads.v19.services.ExperimentService/ScheduleExperiment"
	ExperimentService_PromoteExperiment_FullMethodName         = "/google.ads.googleads.v19.services.ExperimentService/PromoteExperiment"
)

// ExperimentServiceClient is the client API for ExperimentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage experiments.
type ExperimentServiceClient interface {
	// Creates, updates, or removes experiments. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ExperimentError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateExperiments(ctx context.Context, in *MutateExperimentsRequest, opts ...grpc.CallOption) (*MutateExperimentsResponse, error)
	// Immediately ends an experiment, changing the experiment's scheduled
	// end date and without waiting for end of day. End date is updated to be the
	// time of the request.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ExperimentError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	EndExperiment(ctx context.Context, in *EndExperimentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Returns all errors that occurred during the last Experiment update (either
	// scheduling or promotion).
	// Supports standard list paging.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ListExperimentAsyncErrors(ctx context.Context, in *ListExperimentAsyncErrorsRequest, opts ...grpc.CallOption) (*ListExperimentAsyncErrorsResponse, error)
	// Graduates an experiment to a full campaign.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ExperimentError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	GraduateExperiment(ctx context.Context, in *GraduateExperimentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Schedule an experiment. The in design campaign
	// will be converted into a real campaign (called the experiment campaign)
	// that will begin serving ads if successfully created.
	//
	// The experiment is scheduled immediately with status INITIALIZING.
	// This method returns a long running operation that tracks the forking of the
	// in design campaign. If the forking fails, a list of errors can be retrieved
	// using the ListExperimentAsyncErrors method. The operation's
	// metadata will be a string containing the resource name of the created
	// experiment.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ExperimentError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DateRangeError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	ScheduleExperiment(ctx context.Context, in *ScheduleExperimentRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Promotes the trial campaign thus applying changes in the trial campaign
	// to the base campaign.
	// This method returns a long running operation that tracks the promotion of
	// the experiment campaign. If it fails, a list of errors can be retrieved
	// using the ListExperimentAsyncErrors method. The operation's
	// metadata will be a string containing the resource name of the created
	// experiment.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ExperimentError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	PromoteExperiment(ctx context.Context, in *PromoteExperimentRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type experimentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExperimentServiceClient(cc grpc.ClientConnInterface) ExperimentServiceClient {
	return &experimentServiceClient{cc}
}

func (c *experimentServiceClient) MutateExperiments(ctx context.Context, in *MutateExperimentsRequest, opts ...grpc.CallOption) (*MutateExperimentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateExperimentsResponse)
	err := c.cc.Invoke(ctx, ExperimentService_MutateExperiments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentServiceClient) EndExperiment(ctx context.Context, in *EndExperimentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ExperimentService_EndExperiment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentServiceClient) ListExperimentAsyncErrors(ctx context.Context, in *ListExperimentAsyncErrorsRequest, opts ...grpc.CallOption) (*ListExperimentAsyncErrorsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListExperimentAsyncErrorsResponse)
	err := c.cc.Invoke(ctx, ExperimentService_ListExperimentAsyncErrors_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentServiceClient) GraduateExperiment(ctx context.Context, in *GraduateExperimentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ExperimentService_GraduateExperiment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentServiceClient) ScheduleExperiment(ctx context.Context, in *ScheduleExperimentRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ExperimentService_ScheduleExperiment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *experimentServiceClient) PromoteExperiment(ctx context.Context, in *PromoteExperimentRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, ExperimentService_PromoteExperiment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExperimentServiceServer is the server API for ExperimentService service.
// All implementations must embed UnimplementedExperimentServiceServer
// for forward compatibility.
//
// Service to manage experiments.
type ExperimentServiceServer interface {
	// Creates, updates, or removes experiments. Operation statuses are returned.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ExperimentError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	MutateExperiments(context.Context, *MutateExperimentsRequest) (*MutateExperimentsResponse, error)
	// Immediately ends an experiment, changing the experiment's scheduled
	// end date and without waiting for end of day. End date is updated to be the
	// time of the request.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ExperimentError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	EndExperiment(context.Context, *EndExperimentRequest) (*emptypb.Empty, error)
	// Returns all errors that occurred during the last Experiment update (either
	// scheduling or promotion).
	// Supports standard list paging.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ListExperimentAsyncErrors(context.Context, *ListExperimentAsyncErrorsRequest) (*ListExperimentAsyncErrorsResponse, error)
	// Graduates an experiment to a full campaign.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ExperimentError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	GraduateExperiment(context.Context, *GraduateExperimentRequest) (*emptypb.Empty, error)
	// Schedule an experiment. The in design campaign
	// will be converted into a real campaign (called the experiment campaign)
	// that will begin serving ads if successfully created.
	//
	// The experiment is scheduled immediately with status INITIALIZING.
	// This method returns a long running operation that tracks the forking of the
	// in design campaign. If the forking fails, a list of errors can be retrieved
	// using the ListExperimentAsyncErrors method. The operation's
	// metadata will be a string containing the resource name of the created
	// experiment.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ExperimentError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[DateRangeError]()
	//	[FieldError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RangeError]()
	//	[RequestError]()
	ScheduleExperiment(context.Context, *ScheduleExperimentRequest) (*longrunningpb.Operation, error)
	// Promotes the trial campaign thus applying changes in the trial campaign
	// to the base campaign.
	// This method returns a long running operation that tracks the promotion of
	// the experiment campaign. If it fails, a list of errors can be retrieved
	// using the ListExperimentAsyncErrors method. The operation's
	// metadata will be a string containing the resource name of the created
	// experiment.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[ExperimentError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	PromoteExperiment(context.Context, *PromoteExperimentRequest) (*longrunningpb.Operation, error)
	mustEmbedUnimplementedExperimentServiceServer()
}

// UnimplementedExperimentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExperimentServiceServer struct{}

func (UnimplementedExperimentServiceServer) MutateExperiments(context.Context, *MutateExperimentsRequest) (*MutateExperimentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateExperiments not implemented")
}
func (UnimplementedExperimentServiceServer) EndExperiment(context.Context, *EndExperimentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndExperiment not implemented")
}
func (UnimplementedExperimentServiceServer) ListExperimentAsyncErrors(context.Context, *ListExperimentAsyncErrorsRequest) (*ListExperimentAsyncErrorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExperimentAsyncErrors not implemented")
}
func (UnimplementedExperimentServiceServer) GraduateExperiment(context.Context, *GraduateExperimentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GraduateExperiment not implemented")
}
func (UnimplementedExperimentServiceServer) ScheduleExperiment(context.Context, *ScheduleExperimentRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleExperiment not implemented")
}
func (UnimplementedExperimentServiceServer) PromoteExperiment(context.Context, *PromoteExperimentRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromoteExperiment not implemented")
}
func (UnimplementedExperimentServiceServer) mustEmbedUnimplementedExperimentServiceServer() {}
func (UnimplementedExperimentServiceServer) testEmbeddedByValue()                           {}

// UnsafeExperimentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExperimentServiceServer will
// result in compilation errors.
type UnsafeExperimentServiceServer interface {
	mustEmbedUnimplementedExperimentServiceServer()
}

func RegisterExperimentServiceServer(s grpc.ServiceRegistrar, srv ExperimentServiceServer) {
	// If the following call pancis, it indicates UnimplementedExperimentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExperimentService_ServiceDesc, srv)
}

func _ExperimentService_MutateExperiments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateExperimentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentServiceServer).MutateExperiments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperimentService_MutateExperiments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentServiceServer).MutateExperiments(ctx, req.(*MutateExperimentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentService_EndExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndExperimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentServiceServer).EndExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperimentService_EndExperiment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentServiceServer).EndExperiment(ctx, req.(*EndExperimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentService_ListExperimentAsyncErrors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListExperimentAsyncErrorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentServiceServer).ListExperimentAsyncErrors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperimentService_ListExperimentAsyncErrors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentServiceServer).ListExperimentAsyncErrors(ctx, req.(*ListExperimentAsyncErrorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentService_GraduateExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraduateExperimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentServiceServer).GraduateExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperimentService_GraduateExperiment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentServiceServer).GraduateExperiment(ctx, req.(*GraduateExperimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentService_ScheduleExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleExperimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentServiceServer).ScheduleExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperimentService_ScheduleExperiment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentServiceServer).ScheduleExperiment(ctx, req.(*ScheduleExperimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExperimentService_PromoteExperiment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoteExperimentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExperimentServiceServer).PromoteExperiment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExperimentService_PromoteExperiment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExperimentServiceServer).PromoteExperiment(ctx, req.(*PromoteExperimentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExperimentService_ServiceDesc is the grpc.ServiceDesc for ExperimentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExperimentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.ExperimentService",
	HandlerType: (*ExperimentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateExperiments",
			Handler:    _ExperimentService_MutateExperiments_Handler,
		},
		{
			MethodName: "EndExperiment",
			Handler:    _ExperimentService_EndExperiment_Handler,
		},
		{
			MethodName: "ListExperimentAsyncErrors",
			Handler:    _ExperimentService_ListExperimentAsyncErrors_Handler,
		},
		{
			MethodName: "GraduateExperiment",
			Handler:    _ExperimentService_GraduateExperiment_Handler,
		},
		{
			MethodName: "ScheduleExperiment",
			Handler:    _ExperimentService_ScheduleExperiment_Handler,
		},
		{
			MethodName: "PromoteExperiment",
			Handler:    _ExperimentService_PromoteExperiment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/experiment_service.proto",
}
