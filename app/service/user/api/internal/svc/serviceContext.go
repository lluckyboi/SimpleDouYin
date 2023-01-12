package svc

import (
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/app/common"
	"go-zero-demo/app/common/middleware"
	"go-zero-demo/app/service/user/api/internal/config"
	"go-zero-demo/app/service/user/rpc/user"
)

type ServiceContext struct {
	Config     config.Config
	JWTMap     *common.JWTMap
	JWT        rest.Middleware
	UserClient user.User
	RedisDB    *redis.Client
}

func NewServiceContext(c config.Config, JWTMap *common.JWTMap) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserClient: user.NewUser(zrpc.MustNewClient(c.UserLoginRegisterClient)),
		JWT:        middleware.NewJWTMiddleware(JWTMap).Handle,
		JWTMap:     JWTMap,
		RedisDB: redis.NewClient(&redis.Options{
			Addr:     c.RedisDB.RHost,
			Password: c.RedisDB.RPass,
			DB:       0, // use default DB
		}),
	}
}
