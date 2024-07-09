package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im2/internal/errno"
	"im2/internal/response"
	"im2/restful/gateway/internal/logic/user"
	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"
)

// Update user info
func UserPatchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, errno.NewParamsError(err))
			return
		}

		l := user.NewUserPatchLogic(r.Context(), svcCtx)
		resp, err := l.UserPatch(&req)
		response.Response(w, resp, err)
	}
}
