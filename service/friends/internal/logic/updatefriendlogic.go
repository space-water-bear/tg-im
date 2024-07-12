package logic

import (
	"context"
	"database/sql"
	"fmt"
	"im2/internal/errno"
	"im2/internal/utils"
	"im2/service/friends/internal/model"
	"im2/service/user/userclient"

	"im2/service/friends/friends"
	"im2/service/friends/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFriendLogic {
	return &UpdateFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFriendLogic) UpdateFriend(in *friends.UpdateFriendRequest) (*friends.UpdateFriendResponse, error) {
	/*
	 * 例子：李四(id:4)处理张三(id:3)发起的添加好友请求
	 * 当前UserId 是处理人的ID，FriendId 是发起人的ID
	 * 1. 获取好友请求
	 * 2. 通过好友请求
	 */

	// 通过好友请求
	if in.UserId == 0 || in.FriendId == 0 {
		return &friends.UpdateFriendResponse{Code: errno.ErrValidation}, nil
	}

	// 判断是否已经添加过
	source, err := l.svcCtx.Model.GetFriendInfo(l.ctx, in.UserId, in.FriendId)
	if err == nil && source.Status == 1 {
		return &friends.UpdateFriendResponse{Code: errno.ErrFriendExist}, nil
	} else if err == nil && source.Status == 4 {
		return &friends.UpdateFriendResponse{Code: errno.ErrFriendBlackList}, nil
	}

	// 获取添加好友的记录（这个记录是张三作为发起人的，所以是发起人ID[FriendId]） 注意，通过的时候是好友通过，所以是好友ID，当前用户ID
	data, err := l.svcCtx.Model.GetFriendInfo(l.ctx, in.FriendId, in.UserId)
	if err != nil {
		return &friends.UpdateFriendResponse{Code: errno.ErrFriendRequestNotFound}, nil
	}

	if err := l.svcCtx.Model.PassFriend(l.ctx, data.Id, in.Status); err != nil {
		return &friends.UpdateFriendResponse{Code: errno.ErrDatabase}, nil
	}

	if in.Status == 3 { // 拒绝
		return &friends.UpdateFriendResponse{Code: errno.ErrFriendRequestRejected}, nil
	}

	// 获取自己的信息
	rest, err := l.svcCtx.UserRpc.GetInfo(l.ctx, &userclient.GetInfoRequest{UserId: in.FriendId})
	if err != nil {
		return &friends.UpdateFriendResponse{Code: errno.ErrDatabase}, nil
	}

	// 补充关系，即李四作为用户添加张三为好友
	if err := l.svcCtx.Model.Insert(l.ctx, nil, &model.Friends{
		UserId:         in.UserId,   // 李四
		FriendId:       in.FriendId, // 张三
		Status:         1,           // 因为是通过，所以是1
		FriendNickname: utils.TernaryIF(rest.Nickname != "", rest.Nickname, rest.Username),
		FriendUsername: rest.Username,
		FriendEmail:    sql.NullString{String: rest.Email, Valid: true},
		FriendAvatar:   sql.NullString{String: rest.Avatar, Valid: true},
		Remark:         sql.NullString{String: "", Valid: true},
		RequestStatus:  sql.NullBool{Bool: false, Valid: true},
	}); err != nil {
		return &friends.UpdateFriendResponse{Code: errno.ErrDatabase}, nil
	}

	fmt.Println("添加好友成功")
	return &friends.UpdateFriendResponse{Code: errno.OK}, nil
}
