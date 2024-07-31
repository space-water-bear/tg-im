package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"im2/internal/store"
	"im2/restful/gateway/internal/config"
	"im2/service/friends/friendsclient"
	"im2/service/message/messageclient"
	"im2/service/user/userclient"
)

type ServiceContext struct {
	Config     config.Config
	Redis      *redis.Redis
	UserRpc    userclient.User
	FriendsRpc friendsclient.Friends
	MessageRpc messageclient.Message
	Minio      *store.MinioClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.Redis.Host,
		Pass: c.Redis.Pass,
		Type: c.Redis.Type,
		Tls:  c.Redis.Tls,
	})
	// 文件服务
	minioClient := store.NewMinio(store.MinioOption{
		Endpoint:        c.Minio.Endpoint,
		AccessKeyID:     c.Minio.AccessKeyID,
		SecretAccessKey: c.Minio.SecretAccessKey,
		UseSSL:          c.Minio.UseSSL,
	}, c.Minio.Bucket)

	return &ServiceContext{
		Config:     c,
		Redis:      rds,
		UserRpc:    userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		FriendsRpc: friendsclient.NewFriends(zrpc.MustNewClient(c.FriendsRpc)),
		MessageRpc: messageclient.NewMessage(zrpc.MustNewClient(c.MessageRpc)),
		Minio:      minioClient,
	}
}
