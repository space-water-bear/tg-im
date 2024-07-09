package group

import (
	"net/http"

	"im2/internal/response"
	"im2/restful/gateway/internal/logic/group"
	"im2/restful/gateway/internal/svc"
)

// Get group info
func GroupInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := group.NewGroupInfoLogic(r.Context(), svcCtx)
		resp, err := l.GroupInfo()
		response.Response(w, resp, err)
	}
}
