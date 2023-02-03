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
	resp.Comment = &pb.Comment{
		Id:         0,
		User:       &pb.Author{},
		Content:    "",
		CreateDate: "",
	}

	if in.ActionType { //发布评论
		//开始事务
		tx := l.svcCtx.GormDB.Begin()
		//插入记录
		comment := &model.Comment{
			UserID:     in.UserId,
			Content:    in.CommentText,
			CreateDate: time.Now(),
		}
		if err := l.svcCtx.GormDB.Create(comment); err.Error != nil &&
			!errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			log.Println("发布评论查询出错:", err.Error)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, err.Error
		}
		//更新video.comment_count
		err := tx.Exec("UPDATE video SET comment_count=comment_count+1 where video_id = ?", in.VideoId)
		if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			log.Println("更新评论数+1错误:", err.Error)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, err.Error
		}
		//提交事务
		tx.Commit()

		//获取用户
		var user model.User
		err = l.svcCtx.GormDB.Where("user_id = ?", in.UserId).First(&user)
		if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			log.Println("获取用户错误:", err.Error)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, err.Error
		}

		author := &pb.Author{
			Id:            user.UserID,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      false,
		}

		//写入结果
		resCom := &pb.Comment{
			Id:         comment.CommentID,
			User:       author,
			Content:    comment.Content,
			CreateDate: comment.CreateDate.Format("01-02"),
		}

		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "发布评论成功"
		resp.Comment = resCom
	} else { //删除评论
		//先查询是否已经删除
		var com model.Comment
		err := l.svcCtx.GormDB.Where("comment_id = ?", in.CommentId).First(&com)
		if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			log.Println("查找comment_id出错:", err.Error)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, err.Error
		} else if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			resp.StatusCode = status.ErrAlreadyDel
			resp.StatusMsg = "该评论已经被删除"
			return resp, nil
		}
		//开始事务
		tx := l.svcCtx.GormDB.Begin()
		//删除记录
		if err := l.svcCtx.GormDB.
			Where("comment_id = ?", in.CommentId).
			Delete(&model.Comment{}); err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			log.Println("删除记录出错:", err.Error)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, nil
		} else if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			resp.StatusCode = status.ErrNotFind
			resp.StatusMsg = "comment_id有误"
			return resp, nil
		}
		//更新video.comment_count
		err = tx.Exec("UPDATE video SET comment_count=comment_count-1 where video_id = ?", in.VideoId)
		if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			tx.Rollback()
			log.Println("更新评论数-1出错:", err.Error)
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, nil
		}
		//提交事务
		tx.Commit()

		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "删除评论成功"

		author := &pb.Author{
			Id:            0,
			Name:          "",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}
		//写入结果
		resCom := &pb.Comment{
			Id:         0,
			User:       author,
			Content:    "",
			CreateDate: "",
		}
		resp.Comment = resCom
	}
	return resp, nil
}
