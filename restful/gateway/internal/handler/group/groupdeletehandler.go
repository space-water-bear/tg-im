package group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im2/restful/gateway/internal/logic/group"
	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"
)

// Delete group
func GroupDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteGroupRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewGroupDeleteLogic(r.Context(), svcCtx)
		resp, err := l.GroupDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
