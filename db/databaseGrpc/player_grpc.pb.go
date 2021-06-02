// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package databaseGrpc

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

// PlayerServiceClient is the client API for PlayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerServiceClient interface {
	PlayerFindByPlayerId(ctx context.Context, in *PlayerFindByPlayerIdRequest, opts ...grpc.CallOption) (*PlayerFindByPlayerIdResponse, error)
	PlayerAdd(ctx context.Context, in *PlayerAddRequest, opts ...grpc.CallOption) (*PlayerAddResponse, error)
	PlayerUpdateHighestScoreByPlayerId(ctx context.Context, in *PlayerUpdateHighestScoreByPlayerIdRequest, opts ...grpc.CallOption) (*PlayerUpdateHighestScoreByPlayerIdResponse, error)
	PlayerGetRankByPlayerId(ctx context.Context, in *PlayerGetRankByPlayerIdRequest, opts ...grpc.CallOption) (*PlayerGetRankByPlayerIdResponse, error)
	PlayerGetPlayerRank(ctx context.Context, in *PlayerGetPlayerRankRequest, opts ...grpc.CallOption) (*PlayerGetPlayerRankResponse, error)
}

type playerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerServiceClient(cc grpc.ClientConnInterface) PlayerServiceClient {
	return &playerServiceClient{cc}
}

func (c *playerServiceClient) PlayerFindByPlayerId(ctx context.Context, in *PlayerFindByPlayerIdRequest, opts ...grpc.CallOption) (*PlayerFindByPlayerIdResponse, error) {
	out := new(PlayerFindByPlayerIdResponse)
	err := c.cc.Invoke(ctx, "/databaseGrpc.PlayerService/PlayerFindByPlayerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) PlayerAdd(ctx context.Context, in *PlayerAddRequest, opts ...grpc.CallOption) (*PlayerAddResponse, error) {
	out := new(PlayerAddResponse)
	err := c.cc.Invoke(ctx, "/databaseGrpc.PlayerService/PlayerAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) PlayerUpdateHighestScoreByPlayerId(ctx context.Context, in *PlayerUpdateHighestScoreByPlayerIdRequest, opts ...grpc.CallOption) (*PlayerUpdateHighestScoreByPlayerIdResponse, error) {
	out := new(PlayerUpdateHighestScoreByPlayerIdResponse)
	err := c.cc.Invoke(ctx, "/databaseGrpc.PlayerService/PlayerUpdateHighestScoreByPlayerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) PlayerGetRankByPlayerId(ctx context.Context, in *PlayerGetRankByPlayerIdRequest, opts ...grpc.CallOption) (*PlayerGetRankByPlayerIdResponse, error) {
	out := new(PlayerGetRankByPlayerIdResponse)
	err := c.cc.Invoke(ctx, "/databaseGrpc.PlayerService/PlayerGetRankByPlayerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) PlayerGetPlayerRank(ctx context.Context, in *PlayerGetPlayerRankRequest, opts ...grpc.CallOption) (*PlayerGetPlayerRankResponse, error) {
	out := new(PlayerGetPlayerRankResponse)
	err := c.cc.Invoke(ctx, "/databaseGrpc.PlayerService/PlayerGetPlayerRank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerServiceServer is the server API for PlayerService service.
// All implementations must embed UnimplementedPlayerServiceServer
// for forward compatibility
type PlayerServiceServer interface {
	PlayerFindByPlayerId(context.Context, *PlayerFindByPlayerIdRequest) (*PlayerFindByPlayerIdResponse, error)
	PlayerAdd(context.Context, *PlayerAddRequest) (*PlayerAddResponse, error)
	PlayerUpdateHighestScoreByPlayerId(context.Context, *PlayerUpdateHighestScoreByPlayerIdRequest) (*PlayerUpdateHighestScoreByPlayerIdResponse, error)
	PlayerGetRankByPlayerId(context.Context, *PlayerGetRankByPlayerIdRequest) (*PlayerGetRankByPlayerIdResponse, error)
	PlayerGetPlayerRank(context.Context, *PlayerGetPlayerRankRequest) (*PlayerGetPlayerRankResponse, error)
	mustEmbedUnimplementedPlayerServiceServer()
}

// UnimplementedPlayerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlayerServiceServer struct {
}

func (UnimplementedPlayerServiceServer) PlayerFindByPlayerId(context.Context, *PlayerFindByPlayerIdRequest) (*PlayerFindByPlayerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerFindByPlayerId not implemented")
}
func (UnimplementedPlayerServiceServer) PlayerAdd(context.Context, *PlayerAddRequest) (*PlayerAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerAdd not implemented")
}
func (UnimplementedPlayerServiceServer) PlayerUpdateHighestScoreByPlayerId(context.Context, *PlayerUpdateHighestScoreByPlayerIdRequest) (*PlayerUpdateHighestScoreByPlayerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerUpdateHighestScoreByPlayerId not implemented")
}
func (UnimplementedPlayerServiceServer) PlayerGetRankByPlayerId(context.Context, *PlayerGetRankByPlayerIdRequest) (*PlayerGetRankByPlayerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerGetRankByPlayerId not implemented")
}
func (UnimplementedPlayerServiceServer) PlayerGetPlayerRank(context.Context, *PlayerGetPlayerRankRequest) (*PlayerGetPlayerRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerGetPlayerRank not implemented")
}
func (UnimplementedPlayerServiceServer) mustEmbedUnimplementedPlayerServiceServer() {}

// UnsafePlayerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerServiceServer will
// result in compilation errors.
type UnsafePlayerServiceServer interface {
	mustEmbedUnimplementedPlayerServiceServer()
}

func RegisterPlayerServiceServer(s grpc.ServiceRegistrar, srv PlayerServiceServer) {
	s.RegisterService(&PlayerService_ServiceDesc, srv)
}

func _PlayerService_PlayerFindByPlayerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerFindByPlayerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).PlayerFindByPlayerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseGrpc.PlayerService/PlayerFindByPlayerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).PlayerFindByPlayerId(ctx, req.(*PlayerFindByPlayerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_PlayerAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).PlayerAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseGrpc.PlayerService/PlayerAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).PlayerAdd(ctx, req.(*PlayerAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_PlayerUpdateHighestScoreByPlayerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerUpdateHighestScoreByPlayerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).PlayerUpdateHighestScoreByPlayerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseGrpc.PlayerService/PlayerUpdateHighestScoreByPlayerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).PlayerUpdateHighestScoreByPlayerId(ctx, req.(*PlayerUpdateHighestScoreByPlayerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_PlayerGetRankByPlayerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerGetRankByPlayerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).PlayerGetRankByPlayerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseGrpc.PlayerService/PlayerGetRankByPlayerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).PlayerGetRankByPlayerId(ctx, req.(*PlayerGetRankByPlayerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_PlayerGetPlayerRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerGetPlayerRankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).PlayerGetPlayerRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseGrpc.PlayerService/PlayerGetPlayerRank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).PlayerGetPlayerRank(ctx, req.(*PlayerGetPlayerRankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayerService_ServiceDesc is the grpc.ServiceDesc for PlayerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "databaseGrpc.PlayerService",
	HandlerType: (*PlayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlayerFindByPlayerId",
			Handler:    _PlayerService_PlayerFindByPlayerId_Handler,
		},
		{
			MethodName: "PlayerAdd",
			Handler:    _PlayerService_PlayerAdd_Handler,
		},
		{
			MethodName: "PlayerUpdateHighestScoreByPlayerId",
			Handler:    _PlayerService_PlayerUpdateHighestScoreByPlayerId_Handler,
		},
		{
			MethodName: "PlayerGetRankByPlayerId",
			Handler:    _PlayerService_PlayerGetRankByPlayerId_Handler,
		},
		{
			MethodName: "PlayerGetPlayerRank",
			Handler:    _PlayerService_PlayerGetPlayerRank_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "player.proto",
}
