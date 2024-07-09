package user

import (
	"context"
	"im2/internal/errno"
	"im2/service/user/user"
	"strconv"

	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get user info
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// 将字符串转换成int64
	id, _ := strconv.ParseInt(req.ID, 10, 64)

	rest, err := l.svcCtx.UserRpc.GetInfo(l.ctx, &user.GetInfoRequest{
		UserId: id,
	})

	if err != nil {
		return nil, errno.NewDefaultError(errno.InternalServerError)
	} else if rest.Code != 0 {
		return nil, errno.NewDefaultError(rest.Code)
	}

	return &types.UserInfoResponse{
		NickName: rest.Nickname,
		UserId:   rest.UserId,
		Username: rest.Username,
	}, nil
}
