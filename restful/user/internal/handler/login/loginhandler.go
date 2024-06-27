package login

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im2/internal/response"
	"im2/restful/user/internal/logic/login"
	"im2/restful/user/internal/svc"
	"im2/restful/user/internal/types"
)

// Login
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err)
	}
}
