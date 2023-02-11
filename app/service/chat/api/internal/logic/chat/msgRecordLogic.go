package chat

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/service/chat/rpc/chat"
	"context"
	"log"
	"strconv"

	"SimpleDouYin/app/service/chat/api/internal/svc"
	"SimpleDouYin/app/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgRecordLogic {
	return &MsgRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MsgRecordLogic) MsgRecord(req *types.MsgRecordReq) (*types.MsgRecordResp, error) {
	resp := new(types.MsgRecordResp)
	//解析token
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = strconv.Itoa(status.ErrFailParseToken)
		resp.StatusMsg = "token解析失败"
		logx.Error(err.Error())
		return resp, nil
	}

	//解析ID
	tuid, err := strconv.ParseInt(req.ToUserId, 10, 64)
	if err != nil {
		resp.StatusCode = strconv.Itoa(status.ErrOfServer)
		resp.StatusMsg = "ID有误或服务器错误"
		return resp, nil
	}

	//调rpc
	Grsp, err := l.svcCtx.ChatClient.MsgRecord(l.ctx, &chat.MsgRecordReq{
		UserId:       claims.UserId,
		TargetUserId: tuid,
	})
	if err != nil {
		resp.StatusCode = strconv.Itoa(status.ErrOfServer)
		resp.StatusMsg = status.InfoErrOfServer
		log.Print("关注rpc错误", err.Error())
		return resp, nil
	}
	resp.StatusCode = "0"
	resp.StatusMsg = Grsp.StatusMsg

	//msg list
	msgList := Grsp.MsgList
	for _, v := range msgList {
		msg := types.Message{
			Id:         v.Id,
			Content:    v.Content,
			CreateTime: v.CreatTime,
			FromUserId: v.FromUserId,
			ToUserId:   v.ToUserId,
		}
		resp.MsgList = append(resp.MsgList, msg)
	}
	return resp, nil
}
