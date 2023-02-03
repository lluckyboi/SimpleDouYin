package logic

import (
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/service/action/dao/model"
	"context"

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
		//开始事务
		tx := l.svcCtx.GormDB.Begin()
		//创建记录
		if err := tx.Create(&model.Favorite{
			UserID:  in.UserId,
			VideoID: in.VideoId,
		}); err != nil {
			tx.Rollback()
			logx.Info(err)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, nil
		}
		//更新video.favorite_count
		err := tx.Exec("UPDATE video SET favorite_count=favorite_count+1 where video_id = ?", in.VideoId)
		if err != nil {
			tx.Rollback()
			logx.Info(err)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, nil
		}
		//提交事务
		tx.Commit()

		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "点赞成功"
	} else { //否则取消点赞
		//开始事务
		tx := l.svcCtx.GormDB.Begin()
		//删除记录
		if err := tx.Where("user_id = ? and video_id = ?", in.UserId, in.VideoId).
			Delete(&model.Favorite{}); err != nil {
			tx.Rollback()
			logx.Info(err)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, nil
		}
		//更新video.favorite_count
		err := tx.Exec("UPDATE video SET favorite_count=favorite_count-1 where video_id = ?", in.VideoId)
		if err != nil {
			tx.Rollback()
			logx.Info(err)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, nil
		}
		//提交事务
		tx.Commit()

		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "取消点赞成功"
	}
	return resp, nil
}
