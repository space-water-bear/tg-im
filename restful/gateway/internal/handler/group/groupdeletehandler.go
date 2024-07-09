package group

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im2/internal/errno"
	"im2/internal/response"
	"im2/restful/gateway/internal/logic/group"
	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"
)

// Delete group
func GroupDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteGroupRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, errno.NewParamsError(err))
			return
		}

		l := group.NewGroupDeleteLogic(r.Context(), svcCtx)
		resp, err := l.GroupDelete(&req)
		response.Response(w, resp, err)
	}
}
