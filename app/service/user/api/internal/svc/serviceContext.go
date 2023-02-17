package svc

import (
	"SimpleDouYin/app/common"
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/middleware"
	"SimpleDouYin/app/service/user/api/internal/config"
	"SimpleDouYin/app/service/user/rpc/user"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	JWT                 rest.Middleware
	CORSMiddleware      rest.Middleware
	LimitMiddleware     rest.Middleware
	LogPusherMiddleware rest.Middleware

	JWTMap     *jwt.JWTMap
	UserClient user.User
	RedisDB    *redis.Client
	GormDB     *gorm.DB
}

func NewServiceContext(c config.Config, JWTMap *jwt.JWTMap) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{})
	if err != nil {
		logx.Error("gorm open err:", err)
	}
	return &ServiceContext{
		Config:              c,
		GormDB:              db,
		UserClient:          user.NewUser(zrpc.MustNewClient(c.UserClient)),
		JWTMap:              JWTMap,
		CORSMiddleware:      middleware.NewCORSMiddleware().Handle,
		LimitMiddleware:     middleware.NewLimitMiddleware(key.LimitKeyUserApi, c.LimitKey.Seconds, c.LimitKey.Quota).Handle,
		LogPusherMiddleware: middleware.NewLoggerPusher(common.MsgQHost, common.MsgQUser, common.MsgQPass).WithMsgQ,
		RedisDB: redis.NewClient(&redis.Options{
			Addr:     c.RedisDB.RHost,
			Password: c.RedisDB.RPass,
			DB:       0, // use default DB
		}),
	}
}
