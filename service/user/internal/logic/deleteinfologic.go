package logic

import (
	"context"

	"im2/service/user/internal/svc"
	"im2/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInfoLogic {
	return &DeleteInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteInfoLogic) DeleteInfo(in *user.DeleteInfoRequest) (*user.DeleteInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.DeleteInfoResponse{}, nil
}
