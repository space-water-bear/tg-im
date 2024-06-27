package user

import (
	"context"

	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update user info
func NewUserPatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPatchLogic {
	return &UserPatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPatchLogic) UserPatch(req *types.UpdateUserInfoRequest) (resp *types.UpdateUserResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
