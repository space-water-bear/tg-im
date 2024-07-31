package message

import (
	"net/http"

	"im2/internal/response"
	"im2/restful/gateway/internal/logic/message"
	"im2/restful/gateway/internal/svc"
)

// Send file type message
func MessageFileSendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := message.NewMessageFileSendLogic(r.Context(), svcCtx)
		resp, err := l.MessageFileSend(r)
		response.Response(w, resp, err)
	}
}
