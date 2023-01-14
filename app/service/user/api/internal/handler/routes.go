// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	login_register "SimpleDouYin/app/service/user/api/internal/handler/login_register"
	user_info "SimpleDouYin/app/service/user/api/internal/handler/user_info"
	"SimpleDouYin/app/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/douyin/user/register",
				Handler: login_register.UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/user/login",
				Handler: login_register.UserLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/douyin/user",
				Handler: user_info.GetUserInfoHandler(serverCtx),
			},
		},
	)
}
