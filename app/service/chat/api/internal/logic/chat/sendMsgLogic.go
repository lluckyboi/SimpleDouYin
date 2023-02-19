package chat

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/service/chat/dao/model"
	"SimpleDouYin/app/service/chat/rpc/chat"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"strconv"

	"SimpleDouYin/app/service/chat/api/internal/svc"
	"SimpleDouYin/app/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMsgLogic) SendMsg(req *types.SendMsgReq) (*types.SendMsgResp, error) {
	resp := new(types.SendMsgResp)

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
		resp.StatusMsg = "ID有误或服务器错误"
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
	if err != nil || !act {
		resp.StatusCode = status.ErrUnknownAcType
		resp.StatusMsg = "unknown ActionType"
		return resp, nil
	}

	//敏感词和谐
	_, req.Content = l.svcCtx.SensitiveT.Match(req.Content)

	Grsp, err := l.svcCtx.ChatClient.SendMsg(l.ctx, &chat.SendMsgReq{
		UserId:       claims.UserId,
		TargetUserId: tuid,
		ActionType:   act,
		Content:      req.Content,
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
