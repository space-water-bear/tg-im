package login

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"im2/internal/errno"
	"im2/restful/user/internal/svc"
	"im2/restful/user/internal/types"
	urpc "im2/service/user/user"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Login
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {

	rest, err := l.svcCtx.UserRpc.Login(l.ctx, &urpc.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, errno.NewDefaultError(errno.InternalServerError)
	} else if rest.Code != 0 {
		return nil, errno.NewDefaultError(rest.Code)
	}

	return &types.LoginResponse{Token: rest.Token, UserId: rest.UserId}, nil
}
