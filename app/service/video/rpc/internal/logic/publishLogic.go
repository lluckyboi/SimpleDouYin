package logic

import (
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/service/video/dao/model"
	"context"
	"net/http"
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

	t, err := time.Parse("2006-01-02T15:04:05", time.Now().Format("2006-01-02T15:04:05"))
	if err != nil {
		logx.Info(err)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}

	publish := &model.Publish{
		PublishTime: t,
		Title:       in.Title,
		UserID:      in.UserId,
		VideoID:     in.VideoId,
	}
	video := &model.Video{
		VideoID:  in.VideoId,
		PlayURL:  in.CoverUrl,
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

	resp.StatusCode = http.StatusOK
	resp.StatusMsg = "上传视频成功"
	return resp, nil
}
