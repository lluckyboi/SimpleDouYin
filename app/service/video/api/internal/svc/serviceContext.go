package svc

import (
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/middleware"
	"SimpleDouYin/app/service/video/api/internal/config"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config          config.Config
	CORSMiddleware  rest.Middleware
	LimitMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		CORSMiddleware:  middleware.NewCORSMiddleware().Handle,
		LimitMiddleware: middleware.NewLimitMiddleware(key.LimitKeyVideoApi, c.LimitKey.Seconds, c.LimitKey.Quota).Handle,
	}
}
