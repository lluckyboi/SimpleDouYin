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

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, error error) {
	resp = new(types.RegisterResponse)
	//长度校验
	if !(tool.LengthCheck(req.UserName)) ||
		!(tool.LengthCheck(req.PassWord)) {
		resp.StatusCode = status.ErrLengthErr
		resp.StatusMsg = "长度错误"
		return resp, nil
	}

	//哈希取模
	suf := tool.Hash_Mode(req.UserName, key.RedisHashMod)
	//先查询用户名是否已注册
	bl := l.svcCtx.RedisDB.SIsMember(key.RedisUserNameCacheKey+suf, req.UserName)
	if bl.Val() == true {
		resp.StatusMsg = "用户名已存在"
		resp.StatusCode = status.ErrAlreadyHaveUser
		return resp, nil
	}
	//调rpc入库
	RgRes, err := l.svcCtx.UserClient.Register(l.ctx, &user.RegisterReq{
		Username: req.UserName,
		Password: req.PassWord,
	})
	if err != nil {
		logx.Error("注册入库错误：", err)
		resp.StatusCode = status.ErrOfServer
		resp.StatusCode = status.ErrAlreadyHaveUser
		return resp, nil
	}

	//生成token
	user := model.User{UserID: RgRes.UserId}
	token, err := jwt.GenAccessToken(user)
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		logx.Error("生成token错误：", err)
		return
	}

	resp.StatusMsg = RgRes.StatusMsg
	resp.StatusCode = RgRes.StatusCode
	resp.UserId = RgRes.UserId
	resp.Token = token
	return resp, nil
}
