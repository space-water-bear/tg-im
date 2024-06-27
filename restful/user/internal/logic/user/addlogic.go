package user

import (
	"context"

	"im2/restful/user/internal/svc"
	"im2/restful/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Add a new user
func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddUserRequest) (resp *types.UserAddResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
