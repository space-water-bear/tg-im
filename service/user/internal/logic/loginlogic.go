package logic

import (
	"context"
	"im2/internal/auth"
	"im2/internal/errno"
	"im2/internal/utils"

	"im2/service/user/internal/svc"
	"im2/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {

	u, err := l.svcCtx.Model.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return &user.LoginResponse{Code: errno.ErrUserNotFound}, nil
	}

	if err := utils.Compare(u.Password, in.Password); err != nil {
		return &user.LoginResponse{Code: errno.ErrPasswordIncorrect}, nil
	}

	token, err := auth.GenerateToken(u.Id, u.Username, l.svcCtx.Config.AuthToken.AccessSecret, l.svcCtx.Config.AuthToken.AccessExpire)
	if err != nil {
		return &user.LoginResponse{Code: errno.ErrToken}, nil
	}

	return &user.LoginResponse{Token: token, UserId: u.Id}, nil
}
