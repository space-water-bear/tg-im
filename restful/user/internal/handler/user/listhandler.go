package user

import (
	"net/http"

	_ "github.com/zeromicro/go-zero/rest/httpx"
	"im2/internal/response"
	"im2/restful/user/internal/logic/user"
	"im2/restful/user/internal/svc"
)

// Get user list
func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()
		response.Response(w, resp, err)
	}
}
