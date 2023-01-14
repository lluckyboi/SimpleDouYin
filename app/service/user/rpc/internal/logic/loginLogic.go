package logic

import (
	"SimpleDouYin/app/common"
	"SimpleDouYin/app/service/user/dao/model"
	"context"

	"SimpleDouYin/app/service/user/rpc/internal/svc"
	"SimpleDouYin/app/service/user/rpc/pb"

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

func (l *LoginLogic) Login(in *pb.LoginReq) (resp *pb.LoginReps, error error) {
	user := model.User{}
	err := l.svcCtx.GormDB.Where("username = ?", in.Username).First(&user)
	if err.Error != nil {
		logx.Error("登录:查询用户名错误:", err)
		return
	}
	if common.RSA_Encrypt([]byte(in.Password), l.svcCtx.Config.Sec.SecPub) != user.Password {
		resp.StatusCode = 2003
		resp.StatusMsg = "密码错误"
	}

	resp.StatusCode = 200
	resp.StatusMsg = "登录成功"
	resp.UserId = user.UserID
	return resp, nil
}
