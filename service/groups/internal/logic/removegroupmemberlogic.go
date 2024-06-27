package logic

import (
	"context"

	"im2/service/groups/groups"
	"im2/service/groups/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveGroupMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveGroupMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveGroupMemberLogic {
	return &RemoveGroupMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveGroupMemberLogic) RemoveGroupMember(in *groups.RemoveGroupMemberRequest) (*groups.RemoveGroupMemberResponse, error) {
	// todo: add your logic here and delete this line

	return &groups.RemoveGroupMemberResponse{}, nil
}
