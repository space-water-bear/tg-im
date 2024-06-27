package login

import (
	"context"
	"im2/internal/errno"
	"im2/restful/user/internal/svc"
	"im2/restful/user/internal/types"
	urpc "im2/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Register a new user
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {

	rest, err := l.svcCtx.UserRpc.Register(l.ctx, &urpc.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Password,
	})
	if err != nil {
		return nil, errno.NewDefaultError(errno.InternalServerError)
	} else if rest.Code != 0 {
		return nil, errno.NewDefaultError(rest.Code)
	}

	return &types.RegisterResponse{Success: true}, nil
}
