package group

import (
	"context"

	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update group info
func NewGroupPatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPatchLogic {
	return &GroupPatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupPatchLogic) GroupPatch(req *types.UpdateGroupRequest) (resp *types.UpdateGroupResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
