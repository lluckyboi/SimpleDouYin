package token

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/app/common"
	"go-zero-demo/app/service/user/api/internal/svc"
	"go-zero-demo/app/service/user/api/internal/types"
	"go-zero-demo/app/service/user/dao/model"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken() (resp *types.RefreshTokenResponse, err error) {
	//log.Println("start RefreshToken success")
	resp = new(types.RefreshTokenResponse)
	//拿到解析后的值
	userInfo := l.svcCtx.JWTMap.MustGet("UserInfo").(common.TSInfo)
	//log.Println("解析 success")

	//如果不是RefreshToken 返回
	if userInfo.IsRef == false {
		resp.Status = 201
		resp.Info = "非RefreshToken"
		return resp, nil
	}
	//log.Println(userInfo,"success")

	//生成新的Token
	user := model.User{UserRole: userInfo.UserRole, UserPnum: userInfo.UserPnum}
	AccessToken, err := common.GenAccessToken(user)
	if err != nil {
		logx.Error("RefreshToken错误: 签发新AccessToken错误: ", err)
		return resp, err
	}
	//log.Println("new token success",AccessToken)

	resp.Status = 200
	resp.Info = "success"
	resp.AccessToken = AccessToken
	resp.AccessExpire = common.ExpTD * 24
	return resp, nil
}
