package login

import (
	"context"
	"im2/internal/errno"
	"im2/service/user/userclient"

	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

	//fmt.Println("req", req)
	rest, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	//fmt.Println("rest", rest)

	if err != nil {
		return nil, errno.NewDefaultError(errno.InternalServerError)
	} else if rest.Code != 0 {
		return nil, errno.NewDefaultError(rest.Code)
	}

	return &types.LoginResponse{
		UserId: rest.UserId,
		Token:  rest.Token,
		Expire: rest.Expire,
	}, nil

}
