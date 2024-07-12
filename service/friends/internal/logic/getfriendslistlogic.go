package logic

import (
	"context"
	"im2/internal/errno"
	"im2/service/friends/internal/model"

	"im2/service/friends/friends"
	"im2/service/friends/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendsListLogic {
	return &GetFriendsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendsListLogic) GetFriendsList(in *friends.GetFriendsListRequest) (*friends.GetFriendsListResponse, error) {

	friend, total, err := l.svcCtx.Model.FindListByPage(l.ctx, model.Friends{
		UserId:         in.UserId,
		FriendUsername: in.Username,
		FriendNickname: in.Nickname,
	}, in.Page, in.Size)
	if err != nil {
		return &friends.GetFriendsListResponse{Code: errno.ErrDatabase}, nil
	}

	data := make([]*friends.FriendInfo, 0)
	for _, v := range friend {
		data = append(data, &friends.FriendInfo{
			FriendId: v.FriendId,
			Nickname: v.FriendNickname,
			Username: v.FriendUsername,
			Email:    v.FriendEmail.String,
			Avatar:   v.FriendAvatar.String,
		})
	}

	return &friends.GetFriendsListResponse{
		Code:  errno.OK,
		Rows:  data,
		Total: int32(total),
	}, nil
}
