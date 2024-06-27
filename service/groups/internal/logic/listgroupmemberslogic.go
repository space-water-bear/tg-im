package logic

import (
	"context"

	"im2/service/groups/groups"
	"im2/service/groups/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListGroupMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListGroupMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListGroupMembersLogic {
	return &ListGroupMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListGroupMembersLogic) ListGroupMembers(in *groups.ListGroupMembersRequest) (*groups.ListGroupMembersResponse, error) {
	// todo: add your logic here and delete this line

	return &groups.ListGroupMembersResponse{}, nil
}
