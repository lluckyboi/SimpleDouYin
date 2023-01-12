package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserLoginRegisterClient zrpc.RpcClientConf
	RedisDB                 struct {
		RHost string
		RPass string
	}
}
