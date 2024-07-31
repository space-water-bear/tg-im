package message

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im2/internal/errno"
	"im2/internal/response"
	"im2/restful/gateway/internal/logic/message"
	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"
)

// Get message history
func MessageHistoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageHistoryListRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, errno.NewParamsError(err))
			return
		}

		l := message.NewMessageHistoryLogic(r.Context(), svcCtx)
		resp, err := l.MessageHistory(&req)
		response.Response(w, resp, err)
	}
}
