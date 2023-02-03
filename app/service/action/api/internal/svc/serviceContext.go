package svc

import (
	"SimpleDouYin/app/service/action/api/internal/config"
	"SimpleDouYin/app/service/action/api/internal/middleware"
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
		LimitMiddleware: middleware.NewLimitMiddleware().Handle,
	}
}
