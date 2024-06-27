package logic

import (
	"context"

	"im2/service/friends/friends"
	"im2/service/friends/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveFriendLogic {
	return &RemoveFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveFriendLogic) RemoveFriend(in *friends.RemoveFriendRequest) (*friends.RemoveFriendResponse, error) {
	// todo: add your logic here and delete this line

	return &friends.RemoveFriendResponse{}, nil
}
