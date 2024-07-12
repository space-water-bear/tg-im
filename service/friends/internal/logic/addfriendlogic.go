package logic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"im2/internal/errno"
	"im2/internal/utils"
	"im2/service/friends/friends"
	"im2/service/friends/internal/model"
	"im2/service/friends/internal/svc"
	"im2/service/user/user"

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
	/*
	 * 例子，张三(id:3)对李四(id:4)发起添加好友请求，
	 * 1. 判断是否已经添加过
	 * 2. 判断好友ID是否存在
	 */

	if in.UserId == 0 || in.FriendId == 0 {
		return &friends.AddFriendResponse{Code: errno.ErrValidation}, nil
	}

	// 判断是否已经添加过
	data, err := l.svcCtx.Model.GetFriendInfo(l.ctx, in.UserId, in.FriendId)
	if err != nil {
		if !errors.Is(err, gormc.ErrNotFound) {
			logx.WithContext(l.ctx).Errorf("GetFriendInfo error: %v", err)
			return &friends.AddFriendResponse{Code: errno.ErrDatabase}, nil
		}
	}

	if data != nil && data.Status == 1 { // 是否已经是好友
		return &friends.AddFriendResponse{Code: errno.ErrFriendExist}, nil
	} else if data != nil && data.RequestStatus.Bool { // 发送过添加请求
		return &friends.AddFriendResponse{Code: errno.ErrFriendRequestExist}, nil
	}

	// 好友ID是否存在
	rest, err := l.svcCtx.UserRpc.GetInfo(l.ctx, &user.GetInfoRequest{UserId: in.FriendId})
	if err != nil {
		return &friends.AddFriendResponse{
			Code: errno.InternalServerError,
		}, nil
	} else if rest.Code != errno.OK {
		return &friends.AddFriendResponse{
			Code: rest.Code,
		}, nil
	}

	// 当前用户添加好友
	err = l.svcCtx.Model.Insert(l.ctx, nil, &model.Friends{
		UserId:         in.UserId,   // 张三
		FriendId:       in.FriendId, // 李四
		Status:         2,           // 等待确认
		FriendNickname: utils.TernaryIF(rest.Nickname != "", rest.Nickname, rest.Username),
		FriendUsername: rest.Username,
		FriendEmail:    sql.NullString{String: rest.Email, Valid: true},
		FriendAvatar:   sql.NullString{String: rest.Avatar, Valid: true},
		Remark:         sql.NullString{String: in.Remark, Valid: true},
		RequestStatus:  sql.NullBool{Bool: true, Valid: true},
	})
	if err != nil {
		return &friends.AddFriendResponse{
			Code: errno.ErrDatabase,
		}, nil
	}

	return &friends.AddFriendResponse{Code: errno.OK}, nil
}
