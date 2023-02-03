package logic

import (
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/dao/model"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"

	"SimpleDouYin/app/service/action/rpc/internal/svc"
	"SimpleDouYin/app/service/action/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentListLogic) CommentList(in *pb.CommentListReq) (*pb.CommentListResp, error) {
	resp := new(pb.CommentListResp)

	//根据video_id查询评论列表
	var (
		comments []model.Comment
		comCt    int64
	)
	if err := l.svcCtx.GormDB.
		Where("video_id = ?", in.VideoId).
		Find(&comments).
		Count(&comCt); err.Error != nil &&
		!errors.Is(err.Error, gorm.ErrRecordNotFound) {
		log.Println("comment查询出错:", err.Error, " count:", comCt)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}
	if comCt == 0 {
		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "成功"
		return resp, nil
	}

	//查询评论者信息
	var (
		UIDS    []int64
		usersTp []model.User
	)
	for idx := 0; int64(idx) < comCt; idx++ {
		UIDS = append(UIDS, comments[idx].UserID)
	}
	errr := l.svcCtx.GormDB.Model(&model.User{}).
		Where("user_id in ?", UIDS).
		Order(tool.FiledStringBuild("user_id", UIDS)).
		Find(&usersTp)
	if errr.Error != nil && (!errors.Is(errr.Error, gorm.ErrRecordNotFound)) {
		log.Println("评论者查询出错:", errr.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}
	log.Println("评论者查询成功:", usersTp)
	//补全
	UserIdx := make(map[int64]int)
	for _, val := range comments {
		for idx, uv := range usersTp {
			if val.UserID == uv.UserID {
				UserIdx[uv.UserID] = idx
			}
		}
	}
	var users []model.User
	for i := 0; int64(i) < comCt; i++ {
		users = append(users, usersTp[UserIdx[comments[i].UserID]])
	}
	log.Println("评论者补全成功:", users)

	//user对应follow关系
	var Tfollows []model.Follow
	follows := make([]bool, comCt)
	for i := 0; int64(i) < comCt; i++ {
		follows[i] = false
	}

	errr1 := l.svcCtx.GormDB.
		Where("uid = ? and target_uid in ?", in.UserId, UIDS).
		Find(&Tfollows)
	if errr1.Error != nil && (!errors.Is(errr1.Error, gorm.ErrRecordNotFound)) {
		log.Println("查询出错:", errr1.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "服务器错误"
		return resp, nil
	}
	for _, v := range Tfollows {
		for ku, vu := range UIDS {
			if vu == v.TargetUID {
				follows[ku] = true
			}
		}
	}
	log.Println("follows查询成功:", follows)

	//整合结果
	for i := 0; int64(i) < comCt; i++ {
		var Author pb.Author
		var resCt pb.Comment

		//author
		Author.Id = users[i].UserID
		Author.FollowerCount = users[i].FollowCount
		Author.FollowerCount = users[i].FollowerCount
		Author.Name = users[i].Name
		Author.IsFollow = follows[i]

		//comment
		resCt.User = &Author
		resCt.Id = comments[i].CommentID
		resCt.Content = comments[i].Content
		resCt.CreateDate = comments[i].CreateDate.Format("01-02")

		resp.CommentList = append(resp.CommentList, &resCt)
	}
	resp.StatusCode = status.SuccessCode
	resp.StatusMsg = "获取列表成功"
	return resp, nil
}
