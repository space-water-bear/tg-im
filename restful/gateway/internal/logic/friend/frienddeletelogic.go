package friend

import (
	"context"
	"encoding/json"
	"im2/internal/errno"
	"im2/service/friends/friendsclient"

	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete friend
func NewFriendDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendDeleteLogic {
	return &FriendDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendDeleteLogic) FriendDelete(req *types.DeleteFriendRequest) (resp *types.FriendDeleteResponse, err error) {
	userIdJson := l.ctx.Value("id")
	userId, _ := userIdJson.(json.Number).Int64()

	rest, err := l.svcCtx.FriendsRpc.RemoveFriend(l.ctx, &friendsclient.RemoveFriendRequest{
		UserId:   userId,
		FriendId: req.FriendId,
	})

	if err != nil {
		return nil, errno.NewDefaultError(errno.InternalServerError)
	} else if rest.Code != errno.OK {
		return nil, errno.NewDefaultError(rest.Code)
	}

	return &types.FriendDeleteResponse{
		Success: true,
	}, nil
}
