package svc

import (
	"SimpleDouYin/app/common"
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/middleware"
	"SimpleDouYin/app/service/action/api/internal/config"
	"SimpleDouYin/app/service/action/rpc/action"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	CORSMiddleware      rest.Middleware
	LimitMiddleware     rest.Middleware
	LogPusherMiddleware rest.Middleware

	JWTMap       *jwt.JWTMap
	ActionClient action.Action
	RedisDB      *redis.Client
	GormDB       *gorm.DB
}

func NewServiceContext(c config.Config, JWTMap *jwt.JWTMap) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{})
	if err != nil {
		logx.Info("连接至mysql错误")
	}

	return &ServiceContext{
		Config:       c,
		JWTMap:       JWTMap,
		ActionClient: action.NewAction(zrpc.MustNewClient(c.ActionClient)),
		GormDB:       db,
		RedisDB: redis.NewClient(&redis.Options{
			Addr:     c.RedisDB.RHost,
			Password: c.RedisDB.RPass,
			DB:       0, // use default DB
		}),
		CORSMiddleware:      middleware.NewCORSMiddleware().Handle,
		LimitMiddleware:     middleware.NewLimitMiddleware(key.LimitKeyActionApi, c.LimitKey.Seconds, c.LimitKey.Quota).Handle,
		LogPusherMiddleware: middleware.NewLoggerPusher(common.MsgQHost, common.MsgQUser, common.MsgQPass).WithMsgQ,
	}
}
