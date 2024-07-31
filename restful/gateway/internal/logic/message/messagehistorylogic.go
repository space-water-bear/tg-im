package message

import (
	"context"

	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get message history
func NewMessageHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageHistoryLogic {
	return &MessageHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageHistoryLogic) MessageHistory(req *types.MessageHistoryListRequest) (resp *types.MessageHistoryListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
