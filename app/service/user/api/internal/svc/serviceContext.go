package svc

import (
	"SimpleDouYin/app/common"
	"SimpleDouYin/app/common/middleware"
	"SimpleDouYin/app/service/user/api/internal/config"
	"SimpleDouYin/app/service/user/rpc/user"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	JWT             rest.Middleware
	CORSMiddleware  rest.Middleware
	LimitMiddleware rest.Middleware

	JWTMap     *common.JWTMap
	UserClient user.User
	RedisDB    *redis.Client
}

func NewServiceContext(c config.Config, JWTMap *common.JWTMap) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		UserClient:      user.NewUser(zrpc.MustNewClient(c.UserClient)),
		JWTMap:          JWTMap,
		CORSMiddleware:  middleware.NewCORSMiddleware().Handle,
		LimitMiddleware: middleware.NewLimitMiddleware(middleware.KeyUserApi).Handle,
		RedisDB: redis.NewClient(&redis.Options{
			Addr:     c.RedisDB.RHost,
			Password: c.RedisDB.RPass,
			DB:       0, // use default DB
		}),
	}
}
