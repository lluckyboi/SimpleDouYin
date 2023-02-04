package logic

import (
	"context"

	"SimpleDouYin/app/service/chat/rpc/internal/svc"
	"SimpleDouYin/app/service/chat/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMsgRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MsgRecordLogic {
	return &MsgRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MsgRecordLogic) MsgRecord(in *pb.MsgRecordReq) (*pb.MsgRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.MsgRecordResp{}, nil
}
