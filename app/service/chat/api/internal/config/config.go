package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	ChatClient zrpc.RpcClientConf
	RedisDB    struct {
		RHost string
		RPass string
	}
	LimitKey struct {
		//窗口大小
		Seconds int
		//请求上限
		Quota int
	}
	DB struct {
		DataSource string
	}
}
