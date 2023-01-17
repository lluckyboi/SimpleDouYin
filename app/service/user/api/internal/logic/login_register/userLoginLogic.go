package login_register

import (
	"SimpleDouYin/app/common"
	"SimpleDouYin/app/service/user/api/internal/svc"
	"SimpleDouYin/app/service/user/api/internal/types"
	"SimpleDouYin/app/service/user/dao/model"
	"SimpleDouYin/app/service/user/rpc/user"
	"context"

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
	//长度校验
	if !(common.LengthCheck(req.UserName)) ||
		!(common.LengthCheck(req.PassWord)) {
		resp.StatusCode = common.ErrLengthErr
		resp.StatusMsg = "长度错误"
		return resp, nil
	}

	//验证用户名是否存在
	bl := l.svcCtx.RedisDB.SIsMember(common.RedisUserNameCacheKey, req.UserName)
	if bl.Val() == false {
		resp.StatusMsg = "用户名不存在"
		resp.StatusCode = common.ErrNoSuchUser
		return resp, nil
	}
	logx.Info("验证用户名是否存在 success")

	//登录
	rst, err := l.svcCtx.UserClient.Login(l.ctx, &user.LoginReq{
		Username: req.UserName,
		Password: req.PassWord,
	})
	resp.StatusCode = rst.StatusCode
	resp.StatusMsg = rst.StatusMsg
	resp.UserId = rst.UserId
	if err != nil {
		logx.Error("验证密码错误: ", err)
		return resp, nil
	}
	if rst.StatusCode == common.ErrOfServer || rst.StatusCode == common.ErrWrongPassword {
		return resp, nil
	}
	logx.Info("验证密码 success")

	//token
	user := model.User{UserID: rst.UserId}
	resp.Token, err = common.GenAccessToken(user)
	if err != nil {
		resp.StatusCode = common.ErrOfServer
		resp.StatusMsg = common.InfoErrOfServer
		logx.Error("生成token错误：", err)
		return
	}
	return
}
