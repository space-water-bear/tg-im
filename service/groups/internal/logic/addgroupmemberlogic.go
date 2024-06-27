package logic

import (
	"context"

	"im2/service/groups/groups"
	"im2/service/groups/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGroupMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupMemberLogic {
	return &AddGroupMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddGroupMemberLogic) AddGroupMember(in *groups.AddGroupMemberRequest) (*groups.AddGroupMemberResponse, error) {
	// todo: add your logic here and delete this line

	return &groups.AddGroupMemberResponse{}, nil
}
