package logic

import (
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/dao/model"
	"SimpleDouYin/app/service/video/rpc/internal/svc"
	"SimpleDouYin/app/service/video/rpc/pb"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedLogic) Feed(in *pb.FeedReq) (*pb.FeedResp, error) {
	resp := new(pb.FeedResp)

	var publishs []model.Publish
	videos := make([]model.Video, key.FeedNum)
	var usersTp []model.User

	//先去publish找到发布时间符合条件的
	var count int64
	errr := l.svcCtx.GormDB.Model(&model.Publish{}).
		Where("publish_time < ?", in.LastTime).
		Order("publish_time desc").
		Limit(key.FeedNum).
		Find(&publishs).
		Count(&count)
	if errr.Error != nil && (!errors.Is(errr.Error, gorm.ErrRecordNotFound)) {
		log.Println("查publish表错误", errr.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "服务器错误"
		return resp, nil
	}
	log.Println("publish查询成功:", publishs, count)
	//没有了
	if count == 0 {
		resp.StatusCode = http.StatusOK
		resp.StatusMsg = "成功"
		return resp, nil
	}

	//查询对应user和video
	//查询对应user
	var UIDS []int64
	var userCt int64
	for idx := 0; int64(idx) < count; idx++ {
		UIDS = append(UIDS, publishs[idx].UserID)
	}

	errr = l.svcCtx.GormDB.Model(&model.User{}).
		Where("user_id in ?", UIDS).
		Order(tool.FiledStringBuild("user_id", UIDS)).
		Find(&usersTp).
		Count(&userCt)
	if errr.Error != nil && (!errors.Is(errr.Error, gorm.ErrRecordNotFound)) {
		log.Println("查询出错:", errr.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "服务器错误"
		return resp, nil
	}
	log.Println("users查询成功:", usersTp)
	//user补全
	//构造map 标记对应user在userTp位置
	UserIdx := make(map[int64]int)
	for _, val := range publishs {
		for idx, uv := range usersTp {
			if val.UserID == uv.UserID {
				UserIdx[uv.UserID] = idx
			}
		}
	}

	var users []model.User
	for i := 0; int64(i) < count; i++ {
		users = append(users, usersTp[UserIdx[publishs[i].UserID]])
	}
	log.Println("users补全成功:", users)

	//user对应follow关系
	var Tfollows []model.Follow
	follows := make([]bool, count)
	for i := 0; int64(i) < count; i++ {
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

	//对应video
	var VIDS []int64
	for idx := 0; int64(idx) < count; idx++ {
		VIDS = append(VIDS, publishs[idx].VideoID)
	}

	errr = l.svcCtx.GormDB.Model(&model.Video{}).
		Where("video_id in ?", VIDS).
		Order(tool.FiledStringBuild("video_id", VIDS)).
		Find(&videos)
	if errr.Error != nil && (!errors.Is(errr.Error, gorm.ErrRecordNotFound)) {
		log.Println("查询出错:", errr.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "服务器错误"
		return resp, nil
	}
	log.Println("video查询成功:", videos)

	//todo favorite

	//整合成结果
	for i := 0; int64(i) < count; i++ {
		var Author pb.Author
		var resVi pb.Video

		//author
		Author.Id = users[i].UserID
		Author.FollowerCount = users[i].FollowCount
		Author.FollowerCount = users[i].FollowerCount
		Author.Name = users[i].Name
		Author.IsFollow = follows[i]

		//video
		resVi.Author = &Author
		resVi.Title = publishs[i].Title
		resVi.Id = publishs[i].VideoID
		resVi.IsFavorite = false
		resVi.CommentCount = videos[i].CommentCount
		resVi.FavoriteCount = videos[i].FavoriteCount
		resVi.CoverUrl = videos[i].CoverURL
		resVi.PlayUrl = videos[i].PlayURL

		resp.VideoList = append(resp.VideoList, &resVi)
	}

	resp.StatusCode = status.SuccessCode
	resp.StatusMsg = "成功"
	resp.NextTime = publishs[count-1].PublishTime
	return resp, nil
}
