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

// Update group info
func GroupPatchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateGroupRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, errno.NewParamsError(err))
			return
		}

		l := group.NewGroupPatchLogic(r.Context(), svcCtx)
		resp, err := l.GroupPatch(&req)
		response.Response(w, resp, err)
	}
}
