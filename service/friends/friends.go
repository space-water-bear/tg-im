package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"

	"im2/service/friends/friends"
	"im2/service/friends/internal/config"
	"im2/service/friends/internal/server"
	"im2/service/friends/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "service/friends/etc/friends.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logc.MustSetup(c.Log)
	defer logc.Close()

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		friends.RegisterFriendsServer(grpcServer, server.NewFriendsServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
