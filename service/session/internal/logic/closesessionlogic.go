package logic

import (
	"context"

	"im2/service/session/internal/svc"
	"im2/service/session/session"

	"github.com/zeromicro/go-zero/core/logx"
)

type CloseSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCloseSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseSessionLogic {
	return &CloseSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CloseSessionLogic) CloseSession(in *session.CloseSessionRequest) (*session.CloseSessionResponse, error) {
	// todo: add your logic here and delete this line

	return &session.CloseSessionResponse{}, nil
}
