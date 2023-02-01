package logic

import (
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/service/video/dao/model"
	"context"
	"log"
	"time"

	"SimpleDouYin/app/service/video/rpc/internal/svc"
	"SimpleDouYin/app/service/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishLogic) Publish(in *pb.PublishReq) (*pb.PublishResp, error) {
	resp := new(pb.PublishResp)
	//加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		logx.Info(err)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}
	//解析时间
	t, err := time.Parse("2006-01-02T15:04:05Z07:00", time.Now().In(loc).Format("2006-01-02T15:04:05Z07:00"))
	if err != nil {
		logx.Info(err)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}
	log.Print("时间:", t)
	publish := &model.Publish{
		PublishTime: t,
		Title:       in.Title,
		UserID:      in.UserId,
		VideoID:     in.VideoId,
	}
	video := &model.Video{
		VideoID:  in.VideoId,
		PlayURL:  in.VideoUrl,
		CoverURL: in.CoverUrl,
		Hash:     in.Hash,
	}

	//开始事务
	tx := l.svcCtx.GormDB.Begin()
	if err := tx.Create(&video).Error; err != nil {
		tx.Rollback()
		logx.Info(err)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}
	if err := tx.Create(&publish).Error; err != nil {
		tx.Rollback()
		logx.Info(err)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}
	tx.Commit()

	resp.StatusCode = status.SuccessCode
	resp.StatusMsg = "上传视频成功"
	return resp, nil
}
