package svc

import (
	"SimpleDouYin/app/common"
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/middleware"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/service/chat/api/internal/config"
	"SimpleDouYin/app/service/chat/rpc/chat"
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

	JWTMap     *jwt.JWTMap
	ChatClient chat.Chat
	RedisDB    *redis.Client
	GormDB     *gorm.DB
	SensitiveT *tool.SensitiveTrie
}

func NewServiceContext(c config.Config, JWTMap *jwt.JWTMap) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{})
	if err != nil {
		logx.Info("连接至mysql错误")
	}

	return &ServiceContext{
		Config:     c,
		JWTMap:     JWTMap,
		ChatClient: chat.NewChat(zrpc.MustNewClient(c.ChatClient)),
		GormDB:     db,
		RedisDB: redis.NewClient(&redis.Options{
			Addr:     c.RedisDB.RHost,
			Password: c.RedisDB.RPass,
			DB:       0, // use default DB
		}),
		SensitiveT:          tool.NewSensitiveTrie(c.DirtyReplace.Replace).Init(c.DirtyReplace.Words),
		CORSMiddleware:      middleware.NewCORSMiddleware().Handle,
		LimitMiddleware:     middleware.NewLimitMiddleware(key.LimitKeyChatApi, c.LimitKey.Seconds, c.LimitKey.Quota).Handle,
		LogPusherMiddleware: middleware.NewLoggerPusher(common.MsgQHost, common.MsgQUser, common.MsgQPass).WithMsgQ,
	}
}
