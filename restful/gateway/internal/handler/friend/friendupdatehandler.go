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

// Update friend info
func FriendUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateFriendRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, errno.NewParamsError(err))
			return
		}

		l := friend.NewFriendUpdateLogic(r.Context(), svcCtx)
		resp, err := l.FriendUpdate(&req)
		response.Response(w, resp, err)
	}
}
