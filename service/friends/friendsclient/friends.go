// Code generated by goctl. DO NOT EDIT.
// Source: friends.proto

package friendsclient

import (
	"context"

	"im2/service/friends/friends"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddFriendRequest       = friends.AddFriendRequest
	AddFriendResponse      = friends.AddFriendResponse
	FriendInfo             = friends.FriendInfo
	GetFriendRequest       = friends.GetFriendRequest
	GetFriendResponse      = friends.GetFriendResponse
	GetFriendsListRequest  = friends.GetFriendsListRequest
	GetFriendsListResponse = friends.GetFriendsListResponse
	RemoveFriendRequest    = friends.RemoveFriendRequest
	RemoveFriendResponse   = friends.RemoveFriendResponse

	Friends interface {
		GetFriendsList(ctx context.Context, in *GetFriendsListRequest, opts ...grpc.CallOption) (*GetFriendsListResponse, error)
		GetFriend(ctx context.Context, in *GetFriendRequest, opts ...grpc.CallOption) (*GetFriendResponse, error)
		AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error)
		RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*RemoveFriendResponse, error)
	}

	defaultFriends struct {
		cli zrpc.Client
	}
)

func NewFriends(cli zrpc.Client) Friends {
	return &defaultFriends{
		cli: cli,
	}
}

func (m *defaultFriends) GetFriendsList(ctx context.Context, in *GetFriendsListRequest, opts ...grpc.CallOption) (*GetFriendsListResponse, error) {
	client := friends.NewFriendsClient(m.cli.Conn())
	return client.GetFriendsList(ctx, in, opts...)
}

func (m *defaultFriends) GetFriend(ctx context.Context, in *GetFriendRequest, opts ...grpc.CallOption) (*GetFriendResponse, error) {
	client := friends.NewFriendsClient(m.cli.Conn())
	return client.GetFriend(ctx, in, opts...)
}

func (m *defaultFriends) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error) {
	client := friends.NewFriendsClient(m.cli.Conn())
	return client.AddFriend(ctx, in, opts...)
}

func (m *defaultFriends) RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*RemoveFriendResponse, error) {
	client := friends.NewFriendsClient(m.cli.Conn())
	return client.RemoveFriend(ctx, in, opts...)
}
