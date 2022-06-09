// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	resources "github.com/shenzhencenter/google-ads-pb/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BiddingStrategySimulationServiceClient is the client API for BiddingStrategySimulationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BiddingStrategySimulationServiceClient interface {
	// Returns the requested bidding strategy simulation in full detail.
	GetBiddingStrategySimulation(ctx context.Context, in *GetBiddingStrategySimulationRequest, opts ...grpc.CallOption) (*resources.BiddingStrategySimulation, error)
}

type biddingStrategySimulationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBiddingStrategySimulationServiceClient(cc grpc.ClientConnInterface) BiddingStrategySimulationServiceClient {
	return &biddingStrategySimulationServiceClient{cc}
}

func (c *biddingStrategySimulationServiceClient) GetBiddingStrategySimulation(ctx context.Context, in *GetBiddingStrategySimulationRequest, opts ...grpc.CallOption) (*resources.BiddingStrategySimulation, error) {
	out := new(resources.BiddingStrategySimulation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.BiddingStrategySimulationService/GetBiddingStrategySimulation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BiddingStrategySimulationServiceServer is the server API for BiddingStrategySimulationService service.
// All implementations must embed UnimplementedBiddingStrategySimulationServiceServer
// for forward compatibility
type BiddingStrategySimulationServiceServer interface {
	// Returns the requested bidding strategy simulation in full detail.
	GetBiddingStrategySimulation(context.Context, *GetBiddingStrategySimulationRequest) (*resources.BiddingStrategySimulation, error)
	mustEmbedUnimplementedBiddingStrategySimulationServiceServer()
}

// UnimplementedBiddingStrategySimulationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBiddingStrategySimulationServiceServer struct {
}

func (UnimplementedBiddingStrategySimulationServiceServer) GetBiddingStrategySimulation(context.Context, *GetBiddingStrategySimulationRequest) (*resources.BiddingStrategySimulation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBiddingStrategySimulation not implemented")
}
func (UnimplementedBiddingStrategySimulationServiceServer) mustEmbedUnimplementedBiddingStrategySimulationServiceServer() {
}

// UnsafeBiddingStrategySimulationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BiddingStrategySimulationServiceServer will
// result in compilation errors.
type UnsafeBiddingStrategySimulationServiceServer interface {
	mustEmbedUnimplementedBiddingStrategySimulationServiceServer()
}

func RegisterBiddingStrategySimulationServiceServer(s grpc.ServiceRegistrar, srv BiddingStrategySimulationServiceServer) {
	s.RegisterService(&BiddingStrategySimulationService_ServiceDesc, srv)
}

func _BiddingStrategySimulationService_GetBiddingStrategySimulation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBiddingStrategySimulationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiddingStrategySimulationServiceServer).GetBiddingStrategySimulation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.BiddingStrategySimulationService/GetBiddingStrategySimulation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiddingStrategySimulationServiceServer).GetBiddingStrategySimulation(ctx, req.(*GetBiddingStrategySimulationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BiddingStrategySimulationService_ServiceDesc is the grpc.ServiceDesc for BiddingStrategySimulationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BiddingStrategySimulationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.BiddingStrategySimulationService",
	HandlerType: (*BiddingStrategySimulationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBiddingStrategySimulation",
			Handler:    _BiddingStrategySimulationService_GetBiddingStrategySimulation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/bidding_strategy_simulation_service.proto",
}