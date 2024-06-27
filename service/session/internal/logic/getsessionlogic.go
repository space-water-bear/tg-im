package logic

import (
	"context"

	"im2/service/session/internal/svc"
	"im2/service/session/session"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSessionLogic {
	return &GetSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSessionLogic) GetSession(in *session.GetSessionRequest) (*session.SessionInfo, error) {
	// todo: add your logic here and delete this line

	return &session.SessionInfo{}, nil
}
