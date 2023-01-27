package svc

import (
	"SimpleDouYin/app/service/video/rpc/internal/config"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	GormDB *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{})
	if err != nil {
		logx.Error("gorm open err:", err)
	}
	return &ServiceContext{
		Config: c,
		GormDB: db,
		Redis: redis.NewClient(&redis.Options{
			Addr:     c.RedisDB.RHost,
			Password: c.RedisDB.RPass,
			DB:       0, // use default DB
		}),
	}
}
