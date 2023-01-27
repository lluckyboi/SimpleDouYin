package svc

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/middleware"
	"SimpleDouYin/app/service/video/api/internal/config"
	"SimpleDouYin/app/service/video/rpc/videosv"
	"github.com/go-redis/redis"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	CORSMiddleware  rest.Middleware
	LimitMiddleware rest.Middleware

	JWTMap      *jwt.JWTMap
	VideoClient videosv.VideoSv
	RedisDB     *redis.Client
	GormDB      *gorm.DB
	Minio       *minio.Client
}

func NewServiceContext(c config.Config, JWTMap *jwt.JWTMap) *ServiceContext {
	var ssl bool
	if c.Minio.SSL == 1 {
		ssl = true
	} else {
		ssl = false
	}
	MinioClient, err := minio.New(c.Minio.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.Minio.AcKey, c.Minio.Sec, ""),
		Secure: ssl,
	})
	if err != nil {
		logx.Info("连接至minio错误")
	}

	db, err := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{})
	if err != nil {
		logx.Info("连接至mysql错误")
	}

	return &ServiceContext{
		Config:      c,
		JWTMap:      JWTMap,
		Minio:       MinioClient,
		VideoClient: videosv.NewVideoSv(zrpc.MustNewClient(c.VideoClient)),
		GormDB:      db,
		RedisDB: redis.NewClient(&redis.Options{
			Addr:     c.RedisDB.RHost,
			Password: c.RedisDB.RPass,
			DB:       0, // use default DB
		}),
		CORSMiddleware:  middleware.NewCORSMiddleware().Handle,
		LimitMiddleware: middleware.NewLimitMiddleware(key.LimitKeyUserApi, c.LimitKey.Seconds, c.LimitKey.Quota).Handle,
	}
}
