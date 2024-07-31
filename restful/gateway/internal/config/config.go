package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Host string
		Type string
		Pass string
		Tls  bool
	}
	NSQ struct {
		Host string
	}
	Minio struct {
		Endpoint        string
		AccessKeyID     string
		SecretAccessKey string
		Bucket          string
		UseSSL          bool
	}

	UserRpc    zrpc.RpcClientConf
	FriendsRpc zrpc.RpcClientConf
	MessageRpc zrpc.RpcClientConf
}
