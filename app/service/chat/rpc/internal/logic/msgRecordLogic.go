package logic

import (
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/service/chat/dao/model"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"

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
	resp := new(pb.MsgRecordResp)

	var (
		msgs  []model.Message
		msgCt int64
	)
	rst := l.svcCtx.GormDB.Where("(uid = ? and target_uid = ?)||(target_uid = ? and uid = ?)",
		in.UserId, in.TargetUserId, in.TargetUserId, in.UserId).
		Order("create_time desc").
		Find(&msgs).
		Count(&msgCt)
	if rst.Error != nil && !errors.Is(rst.Error, gorm.ErrRecordNotFound) {
		log.Println("聊天记录查询出错:", rst.Error, " count:", msgCt)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}

	//整合结果
	for i := 0; int64(i) < msgCt; i++ {
		msg := &pb.Msg{
			Id:        msgs[i].ID,
			Content:   msgs[i].Content,
			CreatTime: msgs[i].CreateTime.Format("2006-01-02 15-01-05"),
		}
		resp.MsgList = append(resp.MsgList, msg)
	}
	resp.StatusCode = status.SuccessCode
	resp.StatusMsg = "获取聊天记录成功"
	return resp, nil
}
