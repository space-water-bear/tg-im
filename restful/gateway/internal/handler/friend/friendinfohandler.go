package friend

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im2/restful/gateway/internal/logic/friend"
	"im2/restful/gateway/internal/svc"
)

// Get friend info
func FriendInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := friend.NewFriendInfoLogic(r.Context(), svcCtx)
		resp, err := l.FriendInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
