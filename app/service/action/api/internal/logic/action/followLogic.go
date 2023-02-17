package action

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/service/action/dao/model"
	"SimpleDouYin/app/service/action/rpc/action"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"strconv"

	"SimpleDouYin/app/service/action/api/internal/svc"
	"SimpleDouYin/app/service/action/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) (*types.FollowResp, error) {
	resp := new(types.FollowResp)

	//解析token
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = status.ErrFailParseToken
		resp.StatusMsg = "token解析失败"
		logx.Error(err.Error())
		return resp, nil
	}

	//解析ID
	tuid, err := strconv.ParseInt(req.ToUserId, 10, 64)
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "VID有误或服务器错误"
		return resp, nil
	}

	//哈希取模
	sufId := tool.Hash_Mode(req.ToUserId, key.RedisHashMod)
	//查询id是否存在
	bl := l.svcCtx.RedisDB.SIsMember(key.RedisUserIdCacheKey+sufId, tuid)
	if bl.Val() == false {
		res := l.svcCtx.GormDB.Where("user_id = ?", tuid).First(&model.User{})
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
		l.svcCtx.RedisDB.SAdd(key.RedisUserIdCacheKey+sufId, tuid)
	}

	//校验ActionType 1-关注-true 2-取消-false
	act, err := tool.AcTypeStringToBool(req.ActionType)
	if err != nil {
		resp.StatusCode = status.ErrUnknownAcType
		resp.StatusMsg = "unknown ActionType"
		return resp, nil
	}

	Grsp, err := l.svcCtx.ActionClient.
		Follow(l.ctx, &action.FollowReq{
			UserId:       claims.UserId,
			TargetUserId: tuid,
			ActionType:   act,
		})
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		log.Print("关注rpc错误", err.Error())
		return resp, nil
	}
	resp.StatusCode = Grsp.StatusCode
	resp.StatusMsg = Grsp.StatusMsg
	return resp, nil
}
