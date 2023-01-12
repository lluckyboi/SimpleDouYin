package token

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/app/common"
	"go-zero-demo/app/service/user/api/internal/svc"
	"go-zero-demo/app/service/user/api/internal/types"
	"go-zero-demo/app/service/user/rpc/user"
	"strconv"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoResponse, err error) {
	resp = new(types.GetUserInfoResponse)
	//拿到token中信息
	userInfo := l.svcCtx.JWTMap.MustGet("UserInfo").(common.TSInfo)
	//如果是RefreshToken 返回
	if userInfo.IsRef == true {
		resp.Status = 201
		resp.Info = "非AccessToken"
		return resp, nil
	}

	//获取用户信息
	rste, err := l.svcCtx.UserClient.GetInfoByPnum(l.ctx, &user.GetInfoByPnumReq{UserPnum: userInfo.UserPnum})
	//拷贝
	err1 := copier.Copy(&resp, &rste)
	//类型不一致的手动拷贝
	resp.Status, _ = strconv.Atoi(rste.Status)

	if err1 != nil {
		return resp, err
	}
	if err != nil || rste.Status == "500" {
		logx.Error("获取用户信息错误: ", err)
		return resp, nil
	}
	return resp, nil
}
