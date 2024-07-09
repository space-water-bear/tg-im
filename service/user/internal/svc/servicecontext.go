package svc

import (
	//_ "github.com/lib/pq"
	"github.com/SpectatorNan/gorm-zero/gormc/config/pg"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"im2/service/user/internal/config"
	"im2/service/user/internal/model"
	"log"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
	Model  model.UsersModel
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
		Config: c,
		Redis:  rds,
		Model:  model.NewUsersModel(db),
	}
}
