package svc

import (
	"github.com/SpectatorNan/gorm-zero/gormc/config/pg"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"im2/service/friends/internal/config"
	"im2/service/friends/internal/model"
	"im2/service/user/userclient"
	"log"
)

type ServiceContext struct {
	Config  config.Config
	Redis   *redis.Redis
	Model   model.FriendsModel
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := pg.Connect(c.Database)
	if err != nil {
		log.Fatal(err)
	}
	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.Redis.Host,
		Type: c.Redis.Type,
		Pass: c.Redis.Pass,
		Tls:  c.Redis.Tls,
	})
	return &ServiceContext{
		Config:  c,
		Redis:   rds,
		Model:   model.NewFriendsModel(db),
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
