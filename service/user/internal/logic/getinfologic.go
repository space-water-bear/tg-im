package logic

import (
	"context"
	"im2/internal/errno"

	"im2/service/user/internal/svc"
	"im2/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoLogic) GetInfo(in *user.GetInfoRequest) (*user.GetInfoResponse, error) {

	if in.UserId == 0 {
		return &user.GetInfoResponse{Code: errno.ErrUserNotFound}, nil
	}

	u, err := l.svcCtx.Model.FindOne(l.ctx, in.UserId)
	if err != nil {
		return &user.GetInfoResponse{Code: errno.ErrUserNotFound}, nil
	}

	return &user.GetInfoResponse{
		Code:     errno.OK,
		UserId:   u.Id,
		Username: u.Username,
		Nickname: u.Nickname.String,
		Email:    u.Email,
		Ip:       u.Ip.String,
	}, nil

}
