package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	AuthToken struct {
		AccessSecret string
		AccessExpire int64
	}
	// db
	DataSource string
}
