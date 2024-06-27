package user

import (
	"context"

	"im2/restful/user/internal/svc"
	"im2/restful/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update user info
func NewPatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PatchLogic {
	return &PatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PatchLogic) Patch(req *types.UpdateUserInfoRequest) (resp *types.UserUpdateResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
