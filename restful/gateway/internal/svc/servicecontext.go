package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"im2/restful/gateway/internal/config"
	"im2/service/user/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
	Redis   *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Redis: redis.MustNewRedis(redis.RedisConf{
			Host: c.Redis.Host,
			Pass: c.Redis.Pass,
			Type: c.Redis.Type,
			Tls:  c.Redis.Tls,
		}),
	}
}
