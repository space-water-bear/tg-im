package logic

import (
	"context"
	"im2/internal/errno"
	"im2/service/user/user"

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

	if in.FriendId == 0 || in.UserId == 0 {
		return &friends.GetFriendResponse{
			Code: errno.ErrValidation,
		}, nil
	}

	// 是否有关系
	data, err := l.svcCtx.Model.GetFriendInfo(l.ctx, in.UserId, in.FriendId)
	if err != nil {
		return &friends.GetFriendResponse{
			Code: errno.ErrDatabase,
		}, nil
	}

	rest, err := l.svcCtx.UserRpc.GetInfo(l.ctx, &user.GetInfoRequest{
		UserId: data.UserId,
	})

	if err != nil {
		return &friends.GetFriendResponse{
			Code: errno.InternalServerError,
		}, nil
	} else if rest.Code != errno.OK {
		return &friends.GetFriendResponse{
			Code: rest.Code,
		}, nil
	}

	return &friends.GetFriendResponse{
		Code: errno.OK,
		Friend: &friends.FriendInfo{
			FriendId: data.FriendId,
			Nickname: data.FriendNickname,
			Avatar:   data.FriendAvatar.String,
			Username: rest.Username,
			Email:    rest.Email,
		},
	}, nil
}
