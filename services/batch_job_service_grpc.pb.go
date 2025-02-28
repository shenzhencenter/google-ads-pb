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
// source: google/ads/googleads/v19/services/batch_job_service.proto

package services

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
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
	BatchJobService_MutateBatchJob_FullMethodName        = "/google.ads.googleads.v19.services.BatchJobService/MutateBatchJob"
	BatchJobService_ListBatchJobResults_FullMethodName   = "/google.ads.googleads.v19.services.BatchJobService/ListBatchJobResults"
	BatchJobService_RunBatchJob_FullMethodName           = "/google.ads.googleads.v19.services.BatchJobService/RunBatchJob"
	BatchJobService_AddBatchJobOperations_FullMethodName = "/google.ads.googleads.v19.services.BatchJobService/AddBatchJobOperations"
)

// BatchJobServiceClient is the client API for BatchJobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service to manage batch jobs.
type BatchJobServiceClient interface {
	// Mutates a batch job.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	MutateBatchJob(ctx context.Context, in *MutateBatchJobRequest, opts ...grpc.CallOption) (*MutateBatchJobResponse, error)
	// Returns the results of the batch job. The job must be done.
	// Supports standard list paging.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BatchJobError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ListBatchJobResults(ctx context.Context, in *ListBatchJobResultsRequest, opts ...grpc.CallOption) (*ListBatchJobResultsResponse, error)
	// Runs the batch job.
	//
	// The Operation.metadata field type is BatchJobMetadata. When finished, the
	// long running operation will not contain errors or a response. Instead, use
	// ListBatchJobResults to get the results of the job.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BatchJobError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	RunBatchJob(ctx context.Context, in *RunBatchJobRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Add operations to the batch job.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BatchJobError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	AddBatchJobOperations(ctx context.Context, in *AddBatchJobOperationsRequest, opts ...grpc.CallOption) (*AddBatchJobOperationsResponse, error)
}

type batchJobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBatchJobServiceClient(cc grpc.ClientConnInterface) BatchJobServiceClient {
	return &batchJobServiceClient{cc}
}

func (c *batchJobServiceClient) MutateBatchJob(ctx context.Context, in *MutateBatchJobRequest, opts ...grpc.CallOption) (*MutateBatchJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateBatchJobResponse)
	err := c.cc.Invoke(ctx, BatchJobService_MutateBatchJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchJobServiceClient) ListBatchJobResults(ctx context.Context, in *ListBatchJobResultsRequest, opts ...grpc.CallOption) (*ListBatchJobResultsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBatchJobResultsResponse)
	err := c.cc.Invoke(ctx, BatchJobService_ListBatchJobResults_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchJobServiceClient) RunBatchJob(ctx context.Context, in *RunBatchJobRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, BatchJobService_RunBatchJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchJobServiceClient) AddBatchJobOperations(ctx context.Context, in *AddBatchJobOperationsRequest, opts ...grpc.CallOption) (*AddBatchJobOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddBatchJobOperationsResponse)
	err := c.cc.Invoke(ctx, BatchJobService_AddBatchJobOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BatchJobServiceServer is the server API for BatchJobService service.
// All implementations must embed UnimplementedBatchJobServiceServer
// for forward compatibility.
//
// Service to manage batch jobs.
type BatchJobServiceServer interface {
	// Mutates a batch job.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	MutateBatchJob(context.Context, *MutateBatchJobRequest) (*MutateBatchJobResponse, error)
	// Returns the results of the batch job. The job must be done.
	// Supports standard list paging.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BatchJobError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	ListBatchJobResults(context.Context, *ListBatchJobResultsRequest) (*ListBatchJobResultsResponse, error)
	// Runs the batch job.
	//
	// The Operation.metadata field type is BatchJobMetadata. When finished, the
	// long running operation will not contain errors or a response. Instead, use
	// ListBatchJobResults to get the results of the job.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BatchJobError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	RunBatchJob(context.Context, *RunBatchJobRequest) (*longrunningpb.Operation, error)
	// Add operations to the batch job.
	//
	// List of thrown errors:
	//
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[BatchJobError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[ResourceCountLimitExceededError]()
	AddBatchJobOperations(context.Context, *AddBatchJobOperationsRequest) (*AddBatchJobOperationsResponse, error)
	mustEmbedUnimplementedBatchJobServiceServer()
}

// UnimplementedBatchJobServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBatchJobServiceServer struct{}

func (UnimplementedBatchJobServiceServer) MutateBatchJob(context.Context, *MutateBatchJobRequest) (*MutateBatchJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateBatchJob not implemented")
}
func (UnimplementedBatchJobServiceServer) ListBatchJobResults(context.Context, *ListBatchJobResultsRequest) (*ListBatchJobResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBatchJobResults not implemented")
}
func (UnimplementedBatchJobServiceServer) RunBatchJob(context.Context, *RunBatchJobRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunBatchJob not implemented")
}
func (UnimplementedBatchJobServiceServer) AddBatchJobOperations(context.Context, *AddBatchJobOperationsRequest) (*AddBatchJobOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBatchJobOperations not implemented")
}
func (UnimplementedBatchJobServiceServer) mustEmbedUnimplementedBatchJobServiceServer() {}
func (UnimplementedBatchJobServiceServer) testEmbeddedByValue()                         {}

// UnsafeBatchJobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BatchJobServiceServer will
// result in compilation errors.
type UnsafeBatchJobServiceServer interface {
	mustEmbedUnimplementedBatchJobServiceServer()
}

func RegisterBatchJobServiceServer(s grpc.ServiceRegistrar, srv BatchJobServiceServer) {
	// If the following call pancis, it indicates UnimplementedBatchJobServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BatchJobService_ServiceDesc, srv)
}

func _BatchJobService_MutateBatchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBatchJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchJobServiceServer).MutateBatchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BatchJobService_MutateBatchJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchJobServiceServer).MutateBatchJob(ctx, req.(*MutateBatchJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchJobService_ListBatchJobResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBatchJobResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchJobServiceServer).ListBatchJobResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BatchJobService_ListBatchJobResults_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchJobServiceServer).ListBatchJobResults(ctx, req.(*ListBatchJobResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchJobService_RunBatchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunBatchJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchJobServiceServer).RunBatchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BatchJobService_RunBatchJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchJobServiceServer).RunBatchJob(ctx, req.(*RunBatchJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchJobService_AddBatchJobOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBatchJobOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchJobServiceServer).AddBatchJobOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BatchJobService_AddBatchJobOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchJobServiceServer).AddBatchJobOperations(ctx, req.(*AddBatchJobOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BatchJobService_ServiceDesc is the grpc.ServiceDesc for BatchJobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BatchJobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.BatchJobService",
	HandlerType: (*BatchJobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateBatchJob",
			Handler:    _BatchJobService_MutateBatchJob_Handler,
		},
		{
			MethodName: "ListBatchJobResults",
			Handler:    _BatchJobService_ListBatchJobResults_Handler,
		},
		{
			MethodName: "RunBatchJob",
			Handler:    _BatchJobService_RunBatchJob_Handler,
		},
		{
			MethodName: "AddBatchJobOperations",
			Handler:    _BatchJobService_AddBatchJobOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/batch_job_service.proto",
}
