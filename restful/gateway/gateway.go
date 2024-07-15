package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"im2/internal/errno"
	"im2/internal/response"
	"im2/restful/gateway/internal/config"
	"im2/restful/gateway/internal/handler"
	"im2/restful/gateway/internal/svc"
	"net/http"
)

var configFile = flag.String("f", "restful/gateway/etc/gateway.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		response.Response(w, nil, errno.NewDefaultError(errno.ErrToken))
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	handler.RegisterSSE(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
