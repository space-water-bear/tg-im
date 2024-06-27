package svc

import (
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"im2/service/user/internal/config"
	"im2/service/user/internal/model"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redis.Redis
	Model  model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewSqlConn("postgres", c.DataSource)
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
