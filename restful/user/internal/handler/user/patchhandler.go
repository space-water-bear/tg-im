package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im2/internal/response"
	"im2/restful/user/internal/logic/user"
	"im2/restful/user/internal/svc"
	"im2/restful/user/internal/types"
)

// Update user info
func PatchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewPatchLogic(r.Context(), svcCtx)
		resp, err := l.Patch(&req)
		response.Response(w, resp, err)
	}
}
