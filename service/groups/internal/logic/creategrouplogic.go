package logic

import (
	"context"

	"im2/service/groups/groups"
	"im2/service/groups/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGroupLogic) CreateGroup(in *groups.CreateGroupRequest) (*groups.CreateGroupResponse, error) {
	// todo: add your logic here and delete this line

	return &groups.CreateGroupResponse{}, nil
}
