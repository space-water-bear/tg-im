package logic

import (
	"context"

	"im2/service/friends/friends"
	"im2/service/friends/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendLogic {
	return &GetFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendLogic) GetFriend(in *friends.GetFriendRequest) (*friends.GetFriendResponse, error) {
	// todo: add your logic here and delete this line

	return &friends.GetFriendResponse{}, nil
}
