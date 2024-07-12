package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"im2/restful/gateway/internal/config"
	"im2/service/friends/friendsclient"
	"im2/service/user/userclient"
)

type ServiceContext struct {
	Config     config.Config
	Redis      *redis.Redis
	UserRpc    userclient.User
	FriendsRpc friendsclient.Friends
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.Redis.Host,
		Pass: c.Redis.Pass,
		Type: c.Redis.Type,
		Tls:  c.Redis.Tls,
	})
	return &ServiceContext{
		Config:     c,
		Redis:      rds,
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		FriendsRpc: friendsclient.NewFriends(zrpc.MustNewClient(c.FriendsRpc)),
	}
}
