package logic

import (
	"SimpleDouYin/app/common/sec"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/dao/model"
	"SimpleDouYin/app/service/user/rpc/internal/svc"
	"SimpleDouYin/app/service/user/rpc/pb"
	"context"
	"log"

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
	err := l.svcCtx.GormDB.Where("username = ?", in.Username).First(&user)
	if err.Error != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		logx.Error("登录:查询用户名错误:", err)
		return resp, nil
	}

	//解密
	pass, err2 := sec.TripleDesDecrypt(user.Password, l.svcCtx.Config.Sec.DESKey, l.svcCtx.Config.Sec.DESIv)
	if err2 != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		logx.Error("登录:base64解码错误:", err2)
		return resp, nil
	}
	//比较
	if pass != in.Password {
		log.Print(pass)
		resp.StatusCode = status.ErrWrongPassword
		resp.StatusMsg = "密码错误"
		return resp, nil
	}

	resp.StatusCode = status.SuccessCode
	resp.StatusMsg = "登录成功"
	resp.UserId = user.UserID
	return resp, nil
}
