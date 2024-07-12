package logic

import (
	"context"
	"errors"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"im2/internal/errno"

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

	if in.UserId == 0 || in.FriendId == 0 {
		return &friends.RemoveFriendResponse{Code: errno.ErrValidation}, nil
	}

	// 判断是否好友,
	data, err := l.svcCtx.Model.GetFriendInfo(l.ctx, in.UserId, in.FriendId)
	if err != nil {
		if errors.Is(err, gormc.ErrNotFound) {
			return &friends.RemoveFriendResponse{Code: errno.ErrFriendNotFound}, nil
		}
		return &friends.RemoveFriendResponse{Code: errno.ErrDatabase}, nil
	}
	// 删除好友，当前用户向
	if err := l.svcCtx.Model.Delete(l.ctx, nil, data.Id); err != nil {
		return &friends.RemoveFriendResponse{Code: errno.ErrDatabase}, nil
	}

	// 删除好友，好友向
	if dt, err := l.svcCtx.Model.GetFriendInfo(l.ctx, in.FriendId, in.UserId); err == nil {
		if err := l.svcCtx.Model.Delete(l.ctx, nil, dt.Id); err != nil {
			return &friends.RemoveFriendResponse{Code: errno.ErrDatabase}, nil
		}
	} else if errors.Is(err, gormc.ErrNotFound) {
		return &friends.RemoveFriendResponse{Code: errno.ErrFriendNotFound}, nil
	} else {
		return &friends.RemoveFriendResponse{Code: errno.ErrDatabase}, nil
	}

	return &friends.RemoveFriendResponse{Code: errno.OK}, nil
}
