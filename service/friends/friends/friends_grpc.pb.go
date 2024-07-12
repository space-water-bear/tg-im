// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: friends.proto

package friends

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
	Friends_GetFriendsList_FullMethodName = "/friends.Friends/GetFriendsList"
	Friends_GetFriend_FullMethodName      = "/friends.Friends/GetFriend"
	Friends_AddFriend_FullMethodName      = "/friends.Friends/AddFriend"
	Friends_UpdateFriend_FullMethodName   = "/friends.Friends/UpdateFriend"
	Friends_RemoveFriend_FullMethodName   = "/friends.Friends/RemoveFriend"
)

// FriendsClient is the client API for Friends service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendsClient interface {
	GetFriendsList(ctx context.Context, in *GetFriendsListRequest, opts ...grpc.CallOption) (*GetFriendsListResponse, error)
	GetFriend(ctx context.Context, in *GetFriendRequest, opts ...grpc.CallOption) (*GetFriendResponse, error)
	AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error)
	UpdateFriend(ctx context.Context, in *UpdateFriendRequest, opts ...grpc.CallOption) (*UpdateFriendResponse, error)
	RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*RemoveFriendResponse, error)
}

type friendsClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendsClient(cc grpc.ClientConnInterface) FriendsClient {
	return &friendsClient{cc}
}

func (c *friendsClient) GetFriendsList(ctx context.Context, in *GetFriendsListRequest, opts ...grpc.CallOption) (*GetFriendsListResponse, error) {
	out := new(GetFriendsListResponse)
	err := c.cc.Invoke(ctx, Friends_GetFriendsList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsClient) GetFriend(ctx context.Context, in *GetFriendRequest, opts ...grpc.CallOption) (*GetFriendResponse, error) {
	out := new(GetFriendResponse)
	err := c.cc.Invoke(ctx, Friends_GetFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsClient) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error) {
	out := new(AddFriendResponse)
	err := c.cc.Invoke(ctx, Friends_AddFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsClient) UpdateFriend(ctx context.Context, in *UpdateFriendRequest, opts ...grpc.CallOption) (*UpdateFriendResponse, error) {
	out := new(UpdateFriendResponse)
	err := c.cc.Invoke(ctx, Friends_UpdateFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendsClient) RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*RemoveFriendResponse, error) {
	out := new(RemoveFriendResponse)
	err := c.cc.Invoke(ctx, Friends_RemoveFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendsServer is the server API for Friends service.
// All implementations must embed UnimplementedFriendsServer
// for forward compatibility
type FriendsServer interface {
	GetFriendsList(context.Context, *GetFriendsListRequest) (*GetFriendsListResponse, error)
	GetFriend(context.Context, *GetFriendRequest) (*GetFriendResponse, error)
	AddFriend(context.Context, *AddFriendRequest) (*AddFriendResponse, error)
	UpdateFriend(context.Context, *UpdateFriendRequest) (*UpdateFriendResponse, error)
	RemoveFriend(context.Context, *RemoveFriendRequest) (*RemoveFriendResponse, error)
	mustEmbedUnimplementedFriendsServer()
}

// UnimplementedFriendsServer must be embedded to have forward compatible implementations.
type UnimplementedFriendsServer struct {
}

func (UnimplementedFriendsServer) GetFriendsList(context.Context, *GetFriendsListRequest) (*GetFriendsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendsList not implemented")
}
func (UnimplementedFriendsServer) GetFriend(context.Context, *GetFriendRequest) (*GetFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriend not implemented")
}
func (UnimplementedFriendsServer) AddFriend(context.Context, *AddFriendRequest) (*AddFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedFriendsServer) UpdateFriend(context.Context, *UpdateFriendRequest) (*UpdateFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFriend not implemented")
}
func (UnimplementedFriendsServer) RemoveFriend(context.Context, *RemoveFriendRequest) (*RemoveFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFriend not implemented")
}
func (UnimplementedFriendsServer) mustEmbedUnimplementedFriendsServer() {}

// UnsafeFriendsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendsServer will
// result in compilation errors.
type UnsafeFriendsServer interface {
	mustEmbedUnimplementedFriendsServer()
}

func RegisterFriendsServer(s grpc.ServiceRegistrar, srv FriendsServer) {
	s.RegisterService(&Friends_ServiceDesc, srv)
}

func _Friends_GetFriendsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).GetFriendsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friends_GetFriendsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).GetFriendsList(ctx, req.(*GetFriendsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friends_GetFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).GetFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friends_GetFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).GetFriend(ctx, req.(*GetFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friends_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friends_AddFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).AddFriend(ctx, req.(*AddFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friends_UpdateFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).UpdateFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friends_UpdateFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).UpdateFriend(ctx, req.(*UpdateFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Friends_RemoveFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendsServer).RemoveFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Friends_RemoveFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendsServer).RemoveFriend(ctx, req.(*RemoveFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Friends_ServiceDesc is the grpc.ServiceDesc for Friends service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Friends_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "friends.Friends",
	HandlerType: (*FriendsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFriendsList",
			Handler:    _Friends_GetFriendsList_Handler,
		},
		{
			MethodName: "GetFriend",
			Handler:    _Friends_GetFriend_Handler,
		},
		{
			MethodName: "AddFriend",
			Handler:    _Friends_AddFriend_Handler,
		},
		{
			MethodName: "UpdateFriend",
			Handler:    _Friends_UpdateFriend_Handler,
		},
		{
			MethodName: "RemoveFriend",
			Handler:    _Friends_RemoveFriend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "friends.proto",
}
