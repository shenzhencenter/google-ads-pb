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
// source: google/ads/googleads/v19/services/account_budget_proposal_service.proto

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
	AccountBudgetProposalService_MutateAccountBudgetProposal_FullMethodName = "/google.ads.googleads.v19.services.AccountBudgetProposalService/MutateAccountBudgetProposal"
)

// AccountBudgetProposalServiceClient is the client API for AccountBudgetProposalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A service for managing account-level budgets through proposals.
//
// A proposal is a request to create a new budget or make changes to an
// existing one.
//
// Mutates:
// The CREATE operation creates a new proposal.
// UPDATE operations aren't supported.
// The REMOVE operation cancels a pending proposal.
type AccountBudgetProposalServiceClient interface {
	// Creates, updates, or removes account budget proposals.  Operation statuses
	// are returned.
	//
	// List of thrown errors:
	//
	//	[AccountBudgetProposalError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[StringLengthError]()
	MutateAccountBudgetProposal(ctx context.Context, in *MutateAccountBudgetProposalRequest, opts ...grpc.CallOption) (*MutateAccountBudgetProposalResponse, error)
}

type accountBudgetProposalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountBudgetProposalServiceClient(cc grpc.ClientConnInterface) AccountBudgetProposalServiceClient {
	return &accountBudgetProposalServiceClient{cc}
}

func (c *accountBudgetProposalServiceClient) MutateAccountBudgetProposal(ctx context.Context, in *MutateAccountBudgetProposalRequest, opts ...grpc.CallOption) (*MutateAccountBudgetProposalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MutateAccountBudgetProposalResponse)
	err := c.cc.Invoke(ctx, AccountBudgetProposalService_MutateAccountBudgetProposal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountBudgetProposalServiceServer is the server API for AccountBudgetProposalService service.
// All implementations must embed UnimplementedAccountBudgetProposalServiceServer
// for forward compatibility.
//
// A service for managing account-level budgets through proposals.
//
// A proposal is a request to create a new budget or make changes to an
// existing one.
//
// Mutates:
// The CREATE operation creates a new proposal.
// UPDATE operations aren't supported.
// The REMOVE operation cancels a pending proposal.
type AccountBudgetProposalServiceServer interface {
	// Creates, updates, or removes account budget proposals.  Operation statuses
	// are returned.
	//
	// List of thrown errors:
	//
	//	[AccountBudgetProposalError]()
	//	[AuthenticationError]()
	//	[AuthorizationError]()
	//	[DatabaseError]()
	//	[DateError]()
	//	[FieldError]()
	//	[FieldMaskError]()
	//	[HeaderError]()
	//	[InternalError]()
	//	[MutateError]()
	//	[QuotaError]()
	//	[RequestError]()
	//	[StringLengthError]()
	MutateAccountBudgetProposal(context.Context, *MutateAccountBudgetProposalRequest) (*MutateAccountBudgetProposalResponse, error)
	mustEmbedUnimplementedAccountBudgetProposalServiceServer()
}

// UnimplementedAccountBudgetProposalServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccountBudgetProposalServiceServer struct{}

func (UnimplementedAccountBudgetProposalServiceServer) MutateAccountBudgetProposal(context.Context, *MutateAccountBudgetProposalRequest) (*MutateAccountBudgetProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAccountBudgetProposal not implemented")
}
func (UnimplementedAccountBudgetProposalServiceServer) mustEmbedUnimplementedAccountBudgetProposalServiceServer() {
}
func (UnimplementedAccountBudgetProposalServiceServer) testEmbeddedByValue() {}

// UnsafeAccountBudgetProposalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountBudgetProposalServiceServer will
// result in compilation errors.
type UnsafeAccountBudgetProposalServiceServer interface {
	mustEmbedUnimplementedAccountBudgetProposalServiceServer()
}

func RegisterAccountBudgetProposalServiceServer(s grpc.ServiceRegistrar, srv AccountBudgetProposalServiceServer) {
	// If the following call pancis, it indicates UnimplementedAccountBudgetProposalServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AccountBudgetProposalService_ServiceDesc, srv)
}

func _AccountBudgetProposalService_MutateAccountBudgetProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAccountBudgetProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountBudgetProposalServiceServer).MutateAccountBudgetProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountBudgetProposalService_MutateAccountBudgetProposal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountBudgetProposalServiceServer).MutateAccountBudgetProposal(ctx, req.(*MutateAccountBudgetProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountBudgetProposalService_ServiceDesc is the grpc.ServiceDesc for AccountBudgetProposalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountBudgetProposalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v19.services.AccountBudgetProposalService",
	HandlerType: (*AccountBudgetProposalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAccountBudgetProposal",
			Handler:    _AccountBudgetProposalService_MutateAccountBudgetProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v19/services/account_budget_proposal_service.proto",
}
