package logic

import (
	"context"

	"im2/service/session/internal/svc"
	"im2/service/session/session"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSessionLogic {
	return &CreateSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSessionLogic) CreateSession(in *session.CreateSessionRequest) (*session.CreateSessionResponse, error) {
	// todo: add your logic here and delete this line

	return &session.CreateSessionResponse{}, nil
}
