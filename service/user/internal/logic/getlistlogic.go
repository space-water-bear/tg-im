package logic

import (
	"context"
	"database/sql"
	"im2/internal/errno"
	"im2/service/user/internal/model"
	"im2/service/user/internal/svc"
	"im2/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListLogic) GetList(in *user.GetListRequest) (*user.GetListResponse, error) {

	//condition := map[string]interface{}{
	//	"username": in.Username,
	//	"nickname": in.Nickname,
	//	"email":    in.Email,
	//	"status":   in.Status,
	//}

	condition := model.Users{
		Username: in.Username,
		Nickname: sql.NullString{String: in.Nickname, Valid: in.Nickname != ""},
		Email:    in.Email,
		Status:   sql.NullInt64{Int64: int64(in.Status), Valid: in.Status != 0},
	}

	users, total, err := l.svcCtx.Model.FindListByPage(l.ctx, condition, in.Page, in.Size)
	if err != nil {
		return &user.GetListResponse{Code: errno.ErrDatabase}, nil
	}

	data := make([]*user.GetInfoResponse, 0)
	for _, dt := range users {
		data = append(data, &user.GetInfoResponse{
			UserId:   dt.Id,
			Username: dt.Username,
			Nickname: dt.Nickname.String,
			Avatar:   dt.Avatar.String,
			Email:    dt.Email,
		})
	}

	return &user.GetListResponse{
		Code:  errno.OK,
		Total: int32(total),
		Rows:  data,
	}, nil
}
