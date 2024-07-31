package svc

import (
	"github.com/SpectatorNan/gorm-zero/gormc/config/pg"
	"github.com/nsqio/go-nsq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"im2/service/friends/friendsclient"
	"im2/service/message/internal/config"
	"im2/service/message/internal/model"
	"im2/service/user/userclient"
	"log"
)

type ServiceContext struct {
	Config      config.Config
	Redis       *redis.Redis
	UserRpc     userclient.User
	FriendsRpc  friendsclient.Friends
	Model       model.MessagesModel
	NsqProducer *nsq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 数据库
	db, err := pg.Connect(c.Database)
	if err != nil {
		log.Fatal(err)
	}

	// Redis
	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.Redis.Host,
		Type: c.Redis.Type,
		Pass: c.Redis.Pass,
		Tls:  c.Redis.Tls,
	})

	// 消息队列
	nsqProducer, err := nsq.NewProducer(c.Nsq.Host, nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	return &ServiceContext{
		Config:      c,
		Redis:       rds,
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		FriendsRpc:  friendsclient.NewFriends(zrpc.MustNewClient(c.FriendsRpc)),
		Model:       model.NewMessagesModel(db),
		NsqProducer: nsqProducer,
	}
}
