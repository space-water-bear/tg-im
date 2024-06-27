package logic

import (
	"context"

	"im2/service/friends/friends"
	"im2/service/friends/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFriendLogic) AddFriend(in *friends.AddFriendRequest) (*friends.AddFriendResponse, error) {
	// todo: add your logic here and delete this line

	return &friends.AddFriendResponse{}, nil
}
