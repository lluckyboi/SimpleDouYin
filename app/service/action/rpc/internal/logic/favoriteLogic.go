package logic

import (
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/service/action/dao/model"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"

	"SimpleDouYin/app/service/action/rpc/internal/svc"
	"SimpleDouYin/app/service/action/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteLogic) Favorite(in *pb.FavoriteReq) (*pb.FavoriteResp, error) {
	resp := new(pb.FavoriteResp)
	//如果是点赞
	if in.ActionType {
		//检查是否已经点赞
		var fa model.Favorite
		err := l.svcCtx.GormDB.Where("user_id = ? and video_id = ?", in.UserId, in.VideoId).
			First(&fa)
		if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			logx.Info("检查点赞失败:", err.Error)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, err.Error
		} else if !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			resp.StatusCode = status.ErrAlreadyFav
			resp.StatusMsg = "已经点赞了"
			return resp, nil
		}

		log.Print("开始点赞事务")
		//开始事务
		tx := l.svcCtx.GormDB.Begin()
		//创建记录
		if err := tx.Create(&model.Favorite{
			UserID:  in.UserId,
			VideoID: in.VideoId,
		}); err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			logx.Info("创建失败", err)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, err.Error
		}
		log.Println("创建记录成功")
		//更新video.favorite_count
		err = tx.Exec("UPDATE video SET favorite_count=favorite_count+1 where video_id = ?", in.VideoId)
		if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			logx.Info(err)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, err.Error
		}
		log.Println("更新成功")
		//提交事务
		tx.Commit()

		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "点赞成功"
	} else { //否则取消点赞
		//开始事务
		tx := l.svcCtx.GormDB.Begin()
		//删除记录
		if err := tx.Where("user_id = ? and video_id = ?", in.UserId, in.VideoId).
			Delete(&model.Favorite{}); err.Error != nil &&
			!errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			log.Println("删除记录失败:", err.Error)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, err.Error
		}
		//更新video.favorite_count
		err := tx.Exec("UPDATE video SET favorite_count=favorite_count-1 where video_id = ?", in.VideoId)
		if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			log.Println("更新评论-1失败:", err.Error)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, err.Error
		}
		//提交事务
		tx.Commit()

		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "取消点赞成功"
	}
	log.Println(resp)
	return resp, nil
}
