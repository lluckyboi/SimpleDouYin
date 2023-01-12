package logic

import (
	"context"
	"errors"
	"go-zero-demo/app/service/user/rpc/helper"
	"go-zero-demo/app/service/user/rpc/internal/svc"
	"go-zero-demo/app/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendVCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendVCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVCodeLogic {
	return &SendVCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendVCodeLogic) SendVCode(in *pb.SendVCodeReq) (*pb.SendVCodeRes, error) {
	res := new(pb.SendVCodeRes)
	err, b := helper.SendMail(l.svcCtx.Redis, in.PhoneNumber)
	if err != nil {
		res.Mes = "发送失败"
		logx.Error(err)
		return res, errors.New("连接超时，请重试")
	} else if b != true {
		res.Mes = "次数已达上限"
		return res, nil
	}
	res.Mes = "发送成功"
	return res, nil
}
