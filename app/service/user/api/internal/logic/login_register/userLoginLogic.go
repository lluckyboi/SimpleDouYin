package login_register

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/service/user/api/internal/svc"
	"SimpleDouYin/app/service/user/api/internal/types"
	"SimpleDouYin/app/service/user/dao/model"
	"SimpleDouYin/app/service/user/rpc/user"
	"context"
	"errors"
	"gorm.io/gorm"

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
	if !(tool.LengthCheck(req.UserName)) ||
		!(tool.LengthCheck(req.PassWord)) {
		resp.StatusCode = status.ErrLengthErr
		resp.StatusMsg = "长度错误"
		return resp, nil
	}
	//哈希取模
	suf := tool.Hash_Mode(req.UserName, key.RedisHashMod)
	//验证用户名是否存在
	bl := l.svcCtx.RedisDB.SIsMember(key.RedisUserNameCacheKey+suf, req.UserName)
	if bl.Val() == false {
		res := l.svcCtx.GormDB.Where("username = ?", req.UserName).First(&model.User{})
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			resp.StatusMsg = "用户名不存在"
			resp.StatusCode = status.ErrNoSuchUser
			return resp, nil
		} else if res.Error != nil {
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			logx.Error("校验用户名错误：", res.Error)
			return resp, nil
		}
		//回写缓存
		l.svcCtx.RedisDB.SAdd(key.RedisUserNameCacheKey+suf, req.UserName)
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
	if rst.StatusCode == status.ErrOfServer || rst.StatusCode == status.ErrWrongPassword {
		return resp, nil
	}
	logx.Info("验证密码 success")

	//token
	user := model.User{UserID: rst.UserId}
	resp.Token, err = jwt.GenAccessToken(user)
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		logx.Error("生成token错误：", err)
		return
	}
	return
}
