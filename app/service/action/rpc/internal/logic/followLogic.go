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

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowLogic) Follow(in *pb.FollowReq) (*pb.FollowResp, error) {
	resp := new(pb.FollowResp)
	//判断操作
	if in.ActionType {

		//检查是否已经关注
		var hfo model.Follow
		err := l.svcCtx.GormDB.
			Where("uid = ? and target_uid = ?", in.UserId, in.TargetUserId).
			First(&hfo)
		if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			logx.Info("检查关注失败:", err.Error)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, err.Error
		} else if !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			resp.StatusCode = status.ErrAlreadyFo
			resp.StatusMsg = "已经关注"
			return resp, nil
		}

		log.Print("开始关注事务")

		//开始事务
		tx := l.svcCtx.GormDB.Begin()

		//创建记录
		if err := tx.Create(&model.Follow{
			UID:       in.UserId,
			TargetUID: in.TargetUserId,
		}); err != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			logx.Info("创建失败", err)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, err.Error
		}
		log.Println("创建记录成功")

		//更新user.follower_count
		err = tx.Exec("UPDATE user SET follower_count=user.follower_count+1 where user_id = ?", in.UserId)
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
		log.Println("事务结束")

		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "关注成功"
	} else { //否则取消关注
		//开始事务
		tx := l.svcCtx.GormDB.Begin()

		//删除记录
		if err := tx.Where("uid = ? and target_uid = ?", in.UserId, in.TargetUserId).
			Delete(&model.Favorite{}); err.Error != nil &&
			!errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			log.Println("删除记录失败:", err.Error)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, err.Error
		}

		//更新user.follower_count
		err := tx.Exec("UPDATE user SET follower_count=user.follower_count-1 where user_id = ?", in.UserId)
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
		log.Println("事务结束")

		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "取消点赞成功"
	}
	return resp, nil
}
