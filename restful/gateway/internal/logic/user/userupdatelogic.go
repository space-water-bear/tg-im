package user

import (
	"context"
	"im2/internal/errno"
	"im2/service/user/user"

	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update user info
func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UpdateUserInfoRequest) (resp *types.UpdateUserResponse, err error) {
	rest, err := l.svcCtx.UserRpc.UpdateInfo(l.ctx, &user.UpdateInfoRequest{
		UserId:   req.UserId,
		Email:    req.Email,
		Avatar:   req.Avatar,
		Nickname: req.NickName,
		Status:   req.Status,
	})
	if err != nil {
		return nil, errno.NewDefaultError(errno.InternalServerError)
	} else if rest.Code != 0 {
		return nil, errno.NewDefaultError(rest.Code)
	}

	return &types.UpdateUserResponse{
		Success: true,
	}, nil
}
