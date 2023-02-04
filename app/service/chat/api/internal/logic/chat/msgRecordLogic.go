package chat

import (
	"context"

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

func (l *MsgRecordLogic) MsgRecord(req *types.MsgRecordReq) (resp *types.MsgRecordResp, err error) {
	// todo: add your logic here and delete this line

	return
}
