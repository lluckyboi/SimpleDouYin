package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.SendMsgResp{}, nil
}
