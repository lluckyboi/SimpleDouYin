package logic

import (
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/service/chat/dao/model"
	"context"
	"time"

	"SimpleDouYin/app/service/chat/rpc/internal/svc"
	"SimpleDouYin/app/service/chat/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMsgLogic) SendMsg(in *pb.SendMsgReq) (*pb.SendMsgResp, error) {
	resp := new(pb.SendMsgResp)

	//构造入库
	msg := &model.Message{
		Content:    in.Content,
		CreateTime: time.Now(),
		UID:        in.UserId,
		TargetUID:  in.TargetUserId,
	}
	rst := l.svcCtx.GormDB.Create(msg)
	if rst.Error != nil {
		logx.Info("创建message失败", rst.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, rst.Error
	}

	resp.StatusCode = status.SuccessCode
	resp.StatusMsg = "发送成功"
	return resp, nil
}
