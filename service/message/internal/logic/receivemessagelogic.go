package logic

import (
	"context"

	"im2/service/message/internal/svc"
	"im2/service/message/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiveMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReceiveMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiveMessageLogic {
	return &ReceiveMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReceiveMessageLogic) ReceiveMessage(in *message.ReceiveMessageRequest, stream message.Message_ReceiveMessageServer) error {

	return nil
}
