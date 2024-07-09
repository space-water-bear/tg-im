package response

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"im2/internal/errno"
	"net/http"
)

type Body struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body

	var errBody *errno.Errno
	errors.As(err, &errBody)

	fmt.Println("Response resp", resp, errBody)
	if err != nil {
		body.Code = errBody.Code
		body.Msg = errBody.Error()
	} else {
		body.Msg = "OK"
		body.Data = resp
	}
	fmt.Println("Response body", body)
	httpx.OkJson(w, body)
}
