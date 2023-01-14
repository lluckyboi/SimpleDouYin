package logic

import (
	"SimpleDouYin/app/common"
	"SimpleDouYin/app/service/user/dao/model"
	"context"
	"net/http"

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

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginReps, error) {
	resp := new(pb.LoginReps)
	user := model.User{}
	err := l.svcCtx.GormDB.Where("username = ?", in.Username).First(&user)
	if err.Error != nil {
		resp.StatusCode = common.ErrOfServer
		resp.StatusMsg = common.InfoErrOfServer
		logx.Error("登录:查询用户名错误:", err)
		return resp, nil
	}
	if common.RSA_Encrypt([]byte(in.Password), l.svcCtx.Config.Sec.SecPub) != user.Password {
		resp.StatusCode = common.ErrWrongPassword
		resp.StatusMsg = "密码错误"
		return resp, nil
	}

	resp.StatusCode = http.StatusOK
	resp.StatusMsg = "登录成功"
	resp.UserId = user.UserID
	return resp, nil
}
