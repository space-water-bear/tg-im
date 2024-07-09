package logic

import (
	"context"
	"im2/internal/errno"

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
	if in.UserId == 0 {
		return &user.DeleteInfoResponse{
			Code: errno.ErrUserNotFound,
		}, nil
	}

	err := l.svcCtx.Model.Delete(l.ctx, nil, in.UserId)
	if err != nil {
		return &user.DeleteInfoResponse{
			Code: errno.ErrDatabase,
		}, nil
	}

	return &user.DeleteInfoResponse{Code: errno.OK}, nil
}
