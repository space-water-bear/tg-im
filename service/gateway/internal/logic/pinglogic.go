package logic

import (
	"context"

	"im2/service/gateway/gateway"
	"im2/service/gateway/internal/svc"

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

func (l *PingLogic) Ping(in *gateway.Request) (*gateway.Response, error) {
	// todo: add your logic here and delete this line

	return &gateway.Response{}, nil
}
