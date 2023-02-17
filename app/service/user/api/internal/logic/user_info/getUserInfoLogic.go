package user_info

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/service/user/dao/model"
	"SimpleDouYin/app/service/user/rpc/user"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"strconv"

	"SimpleDouYin/app/service/user/api/internal/svc"
	"SimpleDouYin/app/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (*types.GetUserInfoResponse, error) {
	resp := new(types.GetUserInfoResponse)
	//解析token
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = status.ErrFailParseToken
		logx.Error(err.Error())
		return resp, nil
	}

	//转换id类型
	UserId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}

	//哈希取模
	sufId := tool.Hash_Mode(req.UserId, key.RedisHashMod)
	//查询id是否存在
	bl := l.svcCtx.RedisDB.SIsMember(key.RedisUserIdCacheKey+sufId, UserId)
	if bl.Val() == false {
		res := l.svcCtx.GormDB.Where("user_id = ?", UserId).First(&model.User{})
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			resp.StatusMsg = "用户ID不存在"
			resp.StatusCode = status.ErrNoSuchUser
			return resp, nil
		} else if res.Error != nil {
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			logx.Error("校验用户ID错误：", res.Error)
			return resp, nil
		}
		//回写缓存
		l.svcCtx.RedisDB.SAdd(key.RedisUserIdCacheKey+sufId, UserId)
	}

	//调用rpc查询
	GRsp, err := l.svcCtx.UserClient.GetInfo(l.ctx, &user.GetInfoReq{
		UserId:   claims.UserId,
		TargetId: UserId,
	})
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		logx.Error(err.Error())
		return resp, nil
	}
	log.Print("user info rpc成功")
	//返回结果
	resp.StatusCode = GRsp.StatusCode
	resp.StatusMsg = GRsp.StatusMsg
	resp.User = types.User{
		Id:            UserId,
		Name:          GRsp.User.Name,
		FollowCount:   GRsp.User.FollowerCount,
		FollowerCount: GRsp.User.FollowerCount,
		IsFollow:      GRsp.User.IsFollow,
	}
	return resp, nil
}
