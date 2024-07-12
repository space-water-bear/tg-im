package friend

import (
	"context"
	"encoding/json"
	"fmt"
	"im2/internal/errno"
	"im2/service/friends/friendsclient"

	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update friend info
func NewFriendUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendUpdateLogic {
	return &FriendUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendUpdateLogic) FriendUpdate(req *types.UpdateFriendRequest) (resp *types.UpdateFriendResponse, err error) {
	userIdJson := l.ctx.Value("id")
	userId, _ := userIdJson.(json.Number).Int64()

	rest, err := l.svcCtx.FriendsRpc.UpdateFriend(l.ctx, &friendsclient.UpdateFriendRequest{
		UserId:   userId,
		FriendId: req.FriendId,
		Status:   req.Status,
	})

	if err != nil {
		fmt.Println("err: ", err)
		return nil, errno.NewDefaultError(errno.InternalServerError)
	} else if rest.Code != errno.OK {
		return nil, errno.NewDefaultError(rest.Code)
	}

	return &types.UpdateFriendResponse{
		Success: true,
	}, nil
}
