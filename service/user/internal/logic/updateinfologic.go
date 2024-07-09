package logic

import (
	"context"
	"database/sql"
	"im2/internal/errno"
	"im2/service/user/internal/svc"
	"im2/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateInfoLogic) UpdateInfo(in *user.UpdateInfoRequest) (*user.UpdateInfoResponse, error) {

	if in.UserId == 0 {
		return &user.UpdateInfoResponse{
			Code: errno.ErrUserNotFound,
		}, nil
	}

	u, err := l.svcCtx.Model.FindOne(l.ctx, in.UserId)
	if err != nil {
		return &user.UpdateInfoResponse{
			Code: errno.ErrUserNotFound,
		}, nil
	}

	u.Nickname = sql.NullString{String: in.Nickname, Valid: in.Nickname != ""}
	u.Avatar = sql.NullString{String: in.Avatar, Valid: in.Avatar != ""}

	if in.Email != "" {
		u.Email = in.Email
	}
	if in.Status != 0 {
		u.Status = sql.NullInt64{Int64: int64(in.Status), Valid: in.Status != 0}
	}

	err = l.svcCtx.Model.Update(l.ctx, nil, u)
	if err != nil {
		return &user.UpdateInfoResponse{
			Code: errno.ErrDatabase,
		}, nil
	}

	return &user.UpdateInfoResponse{Code: errno.OK}, nil
}
