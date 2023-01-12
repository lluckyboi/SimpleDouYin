package lg

import (
	"context"
	"go-zero-demo/app/common"
	"go-zero-demo/app/service/user/dao/model"
	"go-zero-demo/app/service/user/rpc/user"
	"log"
	"strconv"

	"go-zero-demo/app/service/user/api/internal/svc"
	"go-zero-demo/app/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	resp = new(types.LoginResponse)
	//验证账号是否存在
	bl := l.svcCtx.RedisDB.SIsMember("user_pnum", req.UserPnum)
	if bl.Val() == false {
		resp.Info = "手机号未注册"
		resp.Status = 201
		return resp, nil
	}
	logx.Info("验证账号是否存在 success")
	//验证账号密码
	rst, err := l.svcCtx.UserClient.Login(l.ctx, &user.LoginReq{
		UserPnum: req.UserPnum,
		PassWord: req.UserPassword,
	})
	resp.Status, _ = strconv.Atoi(rst.Status)
	resp.Info = rst.Info
	if err != nil {
		logx.Error("验证密码错误: ", err)
		return resp, nil
	}
	if rst.Status == "500" || rst.Status == "201" {
		return resp, nil
	}
	logx.Info("验证密码 success")

	//获取用户信息
	rste, err := l.svcCtx.UserClient.GetInfoByPnum(l.ctx, &user.GetInfoByPnumReq{
		UserPnum: req.UserPnum,
	})
	resp.Status, _ = strconv.Atoi(rste.Status)
	resp.Info = rste.Info
	resp.UserRole = rste.UserRole
	if err != nil || rste.Status == "500" {
		logx.Error("获取用户信息错误: ", err)
		return resp, nil
	}
	logx.Info("获取用户信息 success")

	//生成token
	User := model.User{
		UserRole: rste.UserRole,
		UserPnum: rste.UserPnum,
	}
	//生成AccessToken
	AccessToken, err := common.GenAccessToken(User)
	if err != nil {
		logx.Error("生成AccessToken失败:", err)
		resp.Info = "服务器错误"
		resp.Status = 500
		return resp, nil
	}
	resp.AccessToken = AccessToken
	//过期时间 小时
	resp.AccessExpire = common.ExpTD * 24
	//生成RefreshToken
	RefreshToken, err := common.GenRefreshToken(User)
	if err != nil {
		logx.Error("生成RefreshToken失败:", err)
		resp.Info = "服务器错误"
		resp.Status = 500
		return resp, nil
	}
	resp.RefreshToken = RefreshToken

	log.Println("生成Token success")
	return resp, nil
}
