package friend

import (
	"net/http"

	"im2/internal/response"
	"im2/restful/gateway/internal/logic/friend"
	"im2/restful/gateway/internal/svc"
)

// Get friend info
func FriendInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := friend.NewFriendInfoLogic(r.Context(), svcCtx)
		resp, err := l.FriendInfo()
		response.Response(w, resp, err)
	}
}
