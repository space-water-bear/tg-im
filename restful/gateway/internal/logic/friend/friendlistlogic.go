package friend

import (
	"context"
	"encoding/json"
	"im2/internal/errno"
	"im2/service/friends/friends"

	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get friend list
func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListRequest) (resp *types.FriendListResponse, err error) {

	userIdJson := l.ctx.Value("id")
	userId, _ := userIdJson.(json.Number).Int64()

	rest, err := l.svcCtx.FriendsRpc.GetFriendsList(l.ctx, &friends.GetFriendsListRequest{
		Page:     req.Page,
		Size:     req.Size,
		Username: req.Search,
		Nickname: req.Search,
		UserId:   userId,
	})

	if err != nil {
		return nil, errno.NewDefaultError(errno.InternalServerError)
	} else if rest.Code != errno.OK {
		return nil, errno.NewDefaultError(rest.Code)
	}

	rows := make([]types.FriendInfoResponse, 0)
	for _, v := range rest.Rows {
		rows = append(rows, types.FriendInfoResponse{
			FriendId: v.FriendId,
			Username: v.Username,
			Nickname: v.Nickname,
			Avatar:   v.Avatar,
			Email:    v.Email,
		})
	}

	return &types.FriendListResponse{
		Rows:  rows,
		Total: rest.Total,
	}, nil
}
