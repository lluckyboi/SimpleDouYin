package user_info

import (
	"SimpleDouYin/app/common"
	"SimpleDouYin/app/service/user/rpc/user"
	"context"
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

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {
	//解析token
	claims, err := common.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = common.ErrFailParseToken
		logx.Error(err)
		return
	}
	//转换id类型
	UserId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		resp.StatusCode = common.ErrOfServer
		resp.StatusMsg = common.InfoErrOfServer
		return
	}

	//查询id是否存在
	bl := l.svcCtx.RedisDB.SIsMember(common.RedisUserIdCacheKey, UserId)
	if bl.Val() == false {
		resp.StatusCode = common.ErrNoSuchUser
		resp.StatusMsg = "无效的id"
		return
	}

	//调用rpc查询
	GRsp, err := l.svcCtx.UserClient.GetInfo(l.ctx, &user.GetInfoReq{
		UserId:   claims.UserId,
		TargetId: UserId,
	})
	if err != nil {
		resp.StatusCode = common.ErrOfServer
		resp.StatusMsg = common.InfoErrOfServer
		logx.Error(err)
		return
	}

	//返回结果
	resp.StatusCode = GRsp.GetStatusCode()
	resp.StatusMsg = GRsp.GetStatusMsg()
	resp.User = types.User{
		Id:            UserId,
		Name:          GRsp.User.Name,
		FollowCount:   GRsp.User.FollowerCount,
		FollowerCount: GRsp.User.FollowerCount,
		IsFollow:      GRsp.User.IsFollow,
	}
	return
}
