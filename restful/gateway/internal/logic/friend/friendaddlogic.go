package friend

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"im2/internal/errno"
	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"
	"im2/service/friends/friendsclient"
)

type FriendAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Add a new friend
func NewFriendAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendAddLogic {
	return &FriendAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendAddLogic) FriendAdd(req *types.AddFriendRequest) (resp *types.AddFriendResponse, err error) {

	userIdJson := l.ctx.Value("id")
	userId, _ := userIdJson.(json.Number).Int64()
	//fmt.Println("userId", userId)
	//fmt.Println("type: ", reflect.TypeOf(userId))

	rest, err := l.svcCtx.FriendsRpc.AddFriend(l.ctx, &friendsclient.AddFriendRequest{
		FriendId: req.FriendId,
		Remark:   req.Remark,
		UserId:   userId,
	})

	if err != nil {
		return nil, errno.NewDefaultError(errno.InternalServerError)
	} else if rest.Code != errno.OK {
		return nil, errno.NewDefaultError(rest.Code)
	}

	//fmt.Println(rest)

	return &types.AddFriendResponse{
		Success: true,
	}, nil
}
