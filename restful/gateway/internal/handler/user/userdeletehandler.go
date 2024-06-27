package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im2/restful/gateway/internal/logic/user"
	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"
)

// Delete user
func UserDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserDeleteLogic(r.Context(), svcCtx)
		resp, err := l.UserDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
