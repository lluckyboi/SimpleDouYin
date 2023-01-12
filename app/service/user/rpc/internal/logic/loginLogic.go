package logic

import (
	"context"
	"encoding/base64"
	"go-zero-demo/app/common"
	"go-zero-demo/app/service/user/dao/model"
	"go-zero-demo/app/service/user/rpc/internal/svc"
	"go-zero-demo/app/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginReps, error) {
	resp := new(pb.LoginReps)

	user := model.User{}
	//验证账号密码
	rst := l.svcCtx.GormDB.Where("user_pnum = ?", in.UserPnum).Find(&user)
	if rst.Error != nil {
		resp.Info = "服务器错误"
		resp.Status = "500"
		logx.Error("登录: 获取用户错误:", rst.Error.Error())
		return resp, nil
	}

	//解密
	res, err := base64.StdEncoding.DecodeString(user.UserPassword)
	if rst.Error != nil {
		resp.Info = "服务器错误"
		resp.Status = "500"
		logx.Error("登录错误: 解密错误:", err)
		return resp, nil
	}
	pass := string(common.RSA_Decrypt((res), "private.pem"))
	if in.PassWord != pass {
		resp.Status = "201"
		resp.Info = "密码错误"
		return resp, nil
	}
	resp.Status = "200"
	resp.Info = "登陆成功"
	return resp, nil
}
