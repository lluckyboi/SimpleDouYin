// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	chat "SimpleDouYin/app/service/chat/api/internal/handler/chat"
	"SimpleDouYin/app/service/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CORSMiddleware, serverCtx.LimitMiddleware, serverCtx.LogPusherMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/douyin/message/action",
					Handler: chat.SendMsgHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/douyin/message/chat",
					Handler: chat.MsgRecordHandler(serverCtx),
				},
			}...,
		),
	)
}
