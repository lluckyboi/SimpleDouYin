package svc

import (
	"SimpleDouYin/app/service/user/api/internal/config"
	"SimpleDouYin/app/service/user/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	JWT    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		JWT:    middleware.NewJWTMiddleware().Handle,
	}
}
