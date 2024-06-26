// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: proto/api/recommender.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RecommenderService_GetModules_FullMethodName = "/recommender.RecommenderService/GetModules"
)

// RecommenderServiceClient is the client API for RecommenderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecommenderServiceClient interface {
	GetModules(ctx context.Context, in *GetModulesRequest, opts ...grpc.CallOption) (*GetModulesResponse, error)
}

type recommenderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecommenderServiceClient(cc grpc.ClientConnInterface) RecommenderServiceClient {
	return &recommenderServiceClient{cc}
}

func (c *recommenderServiceClient) GetModules(ctx context.Context, in *GetModulesRequest, opts ...grpc.CallOption) (*GetModulesResponse, error) {
	out := new(GetModulesResponse)
	err := c.cc.Invoke(ctx, RecommenderService_GetModules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecommenderServiceServer is the server API for RecommenderService service.
// All implementations must embed UnimplementedRecommenderServiceServer
// for forward compatibility
type RecommenderServiceServer interface {
	GetModules(context.Context, *GetModulesRequest) (*GetModulesResponse, error)
	mustEmbedUnimplementedRecommenderServiceServer()
}

// UnimplementedRecommenderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecommenderServiceServer struct {
}

func (UnimplementedRecommenderServiceServer) GetModules(context.Context, *GetModulesRequest) (*GetModulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModules not implemented")
}
func (UnimplementedRecommenderServiceServer) mustEmbedUnimplementedRecommenderServiceServer() {}

// UnsafeRecommenderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecommenderServiceServer will
// result in compilation errors.
type UnsafeRecommenderServiceServer interface {
	mustEmbedUnimplementedRecommenderServiceServer()
}

func RegisterRecommenderServiceServer(s grpc.ServiceRegistrar, srv RecommenderServiceServer) {
	s.RegisterService(&RecommenderService_ServiceDesc, srv)
}

func _RecommenderService_GetModules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecommenderServiceServer).GetModules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RecommenderService_GetModules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecommenderServiceServer).GetModules(ctx, req.(*GetModulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RecommenderService_ServiceDesc is the grpc.ServiceDesc for RecommenderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecommenderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "recommender.RecommenderService",
	HandlerType: (*RecommenderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModules",
			Handler:    _RecommenderService_GetModules_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api/recommender.proto",
}
