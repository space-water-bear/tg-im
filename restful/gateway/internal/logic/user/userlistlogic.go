package user

import (
	"context"
	"im2/internal/errno"
	"im2/service/user/user"

	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get user list
func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListRequest) (resp *types.UserListResponse, err error) {
	rest, err := l.svcCtx.UserRpc.GetList(l.ctx, &user.GetListRequest{
		Page:     req.Page,
		Size:     req.Size,
		Username: req.Username,
		Nickname: req.NickName,
	})

	if err != nil {
		return nil, errno.NewDefaultError(errno.InternalServerError)
	} else if rest.Code != 0 {
		return nil, errno.NewDefaultError(rest.Code)
	}

	rows := make([]types.UserInfoResponse, 0)
	for _, v := range rest.Rows {
		rows = append(rows, types.UserInfoResponse{
			UserId:   v.UserId,
			Username: v.Username,
			NickName: v.Nickname,
		})
	}

	return &types.UserListResponse{
		Rows:  rows,
		Total: rest.Total,
	}, nil
}
