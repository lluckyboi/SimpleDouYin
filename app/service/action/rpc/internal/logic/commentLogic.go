package logic

import (
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/service/action/dao/model"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"

	"SimpleDouYin/app/service/action/rpc/internal/svc"
	"SimpleDouYin/app/service/action/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentLogic) Comment(in *pb.CommentReq) (*pb.CommentResp, error) {
	resp := new(pb.CommentResp)

	if in.ActionType { //发布评论
		//开始事务
		tx := l.svcCtx.GormDB.Begin()
		//插入记录
		if err := l.svcCtx.GormDB.Create(&model.Comment{
			UserID:     in.UserId,
			Content:    in.CommentText,
			CreateDate: time.Now(),
		}); err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			log.Println("发布评论查询出错:", err.Error)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, nil
		}
		//更新video.comment_count
		err := tx.Exec("UPDATE video SET comment_count=comment_count+1 where video_id = ?", in.VideoId)
		if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			logx.Info(err)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, nil
		}
		//提交事务
		tx.Commit()

		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "发布评论成功"
	} else { //删除评论
		//开始事务
		tx := l.svcCtx.GormDB.Begin()
		//删除记录
		if err := l.svcCtx.GormDB.
			Where("comment_id = ?", in.CommentId).
			Delete(&model.Favorite{}); err.Error != nil &&
			!errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			logx.Info(err)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, nil
		}
		//更新video.comment_count
		err := tx.Exec("UPDATE video SET comment_count=comment_count-1 where video_id = ?", in.VideoId)
		if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			logx.Info(err)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, nil
		}
		//提交事务
		tx.Commit()

		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "删除评论成功"
	}
	return resp, nil
}
