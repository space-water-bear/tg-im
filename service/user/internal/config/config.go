package config

import (
	"github.com/SpectatorNan/gorm-zero/gormc/config/pg"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AuthToken struct {
		AccessSecret string
		AccessExpire int64
	}
	// db
	Database pg.PgSql
}
