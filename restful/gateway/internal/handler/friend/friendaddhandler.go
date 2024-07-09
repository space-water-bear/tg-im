package friend

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im2/internal/errno"
	"im2/internal/response"
	"im2/restful/gateway/internal/logic/friend"
	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"
)

// Add a new friend
func FriendAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddFriendRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, errno.NewParamsError(err))
			return
		}

		l := friend.NewFriendAddLogic(r.Context(), svcCtx)
		resp, err := l.FriendAdd(&req)
		response.Response(w, resp, err)
	}
}
