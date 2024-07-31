package config

import (
	"github.com/SpectatorNan/gorm-zero/gormc/config/pg"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	UserRpc    zrpc.RpcClientConf
	FriendsRpc zrpc.RpcClientConf
	Database   pg.PgSql
	Nsq        struct {
		Host string
	}
}
