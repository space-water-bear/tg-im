package logic

import (
	"context"

	"im2/service/message/internal/svc"
	"im2/service/message/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *message.SendMessageRequest) (*message.SendMessageResponse, error) {
	// todo: add your logic here and delete this line

	return &message.SendMessageResponse{}, nil
}
