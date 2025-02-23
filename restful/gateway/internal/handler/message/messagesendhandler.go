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

// Send message
func MessageSendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, errno.NewParamsError(err))
			return
		}

		l := message.NewMessageSendLogic(r.Context(), svcCtx)
		resp, err := l.MessageSend(&req)
		response.Response(w, resp, err)
	}
}
