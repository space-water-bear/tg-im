package user

import (
	"context"
	"fmt"
	"im2/internal/errno"
	"im2/service/user/user"

	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete user
func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req *types.DeleteUserRequest) (resp *types.DeleteUserResponse, err error) {

	fmt.Println(req)
	rest, err := l.svcCtx.UserRpc.DeleteInfo(l.ctx, &user.DeleteInfoRequest{
		UserId: req.ID,
	})
	if err != nil {
		return nil, errno.NewDefaultError(errno.InternalServerError)
	} else if rest.Code != 0 {
		return nil, errno.NewDefaultError(rest.Code)
	}

	return &types.DeleteUserResponse{
		Success: true,
	}, nil
}
