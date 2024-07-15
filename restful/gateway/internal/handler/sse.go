package handler

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"im2/internal/auth"
	"im2/internal/errno"
	"im2/internal/response"
	"im2/restful/gateway/internal/svc"
	"io"
	"net/http"
	"reflect"
	"time"
)

type Message struct {
	Code      string      `json:"code"`
	Data      interface{} `json:"data"`
	Timestamp string      `json:"timestamp"`
}

type Request struct {
	Token string `path:"token,optional" form:"token,optional"`
}

func handleClientDisconnect(r *http.Request) {
	fmt.Println("handleClientDisconnect: ", r.RemoteAddr)
}

func sseHandle(svcCtx *svc.ServiceContext) http.HandlerFunc {
	fmt.Println(svcCtx.Config.Auth.AccessSecret)
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("sseHandle")
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		var req Request
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println("parse error", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if req.Token == "" {
			response.Response(w, nil, errno.NewDefaultError(errno.ErrToken))
			return
		}

		jwtData, err := auth.CheckToken(req.Token, svcCtx.Config.Auth.AccessSecret)
		if err != nil {
			fmt.Println("check token error", err)
			response.Response(w, nil, errno.NewDefaultError(errno.ErrToken))
			return
		}
		//fmt.Println("jwtData: ", jwtData)
		fmt.Println("userId: ", jwtData, reflect.TypeOf(jwtData))

		channel := make(chan Message)

		// 使用一个 goroutine 来生成数据
		go func() {
			defer close(channel)
			for i := 0; ; i++ {
				select {
				case <-r.Context().Done():
					return
				case <-time.After(1 * time.Second):
					channel <- Message{
						Code:      "10000",
						Data:      "ping",
						Timestamp: time.Now().Format(time.RFC3339),
					}
				}
			}

		}()

		// 定义写入器函数
		sseWriter := func(w io.Writer) bool {

			select {
			case <-r.Context().Done():
				handleClientDisconnect(r)
				return false
			case message := <-channel:
				jsonData, _ := json.Marshal(message)
				_, err := fmt.Fprintf(w, "data: %s\n\n", jsonData)
				if err != nil {
					fmt.Printf("Error writing to wr: %v\n", err)
					handleClientDisconnect(r)
					return false
				}
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
				return true
			}

			//message, ok := <-channel
			//if !ok {
			//	return false
			//}
			//_, err := fmt.Fprintf(w, "data: %s\n\n", message)
			//return err == nil
		}

		httpx.Stream(r.Context(), w, sseWriter)
	}
}

func RegisterSSE(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/sse",
				Handler: sseHandle(serverCtx),
			},
		},
	)
}
