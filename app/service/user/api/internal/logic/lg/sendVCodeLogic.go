package lg

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/app/service/user/api/internal/svc"
	"go-zero-demo/app/service/user/api/internal/types"
	"go-zero-demo/app/service/user/rpc/user"
)

type SendVCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendVCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVCodeLogic {
	return &SendVCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendVCodeLogic) SendVCode(req *types.SendVCodeRequest) (resp *types.SendVCodeResponse, err error) {
	resp = new(types.SendVCodeResponse)
	res, err := l.svcCtx.UserClient.SendVCode(l.ctx, &user.SendVCodeReq{
		PhoneNumber: req.PhoneNumber,
	})
	logx.Info(resp)
	if err != nil {
		resp.Status = 500
		resp.Info = err.Error()
		return resp, nil
	} else if res.Mes == "次数已达上限" {
		resp.Status = 201
		resp.Info = res.Mes
		return resp, nil
	}
	resp.Status = 200
	resp.Info = "success"
	return resp, nil
}
