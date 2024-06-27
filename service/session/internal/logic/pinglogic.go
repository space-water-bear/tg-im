package logic

import (
	"context"

	"im2/service/session/internal/svc"
	"im2/service/session/session"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *session.Request) (*session.Response, error) {
	// todo: add your logic here and delete this line

	return &session.Response{}, nil
}
