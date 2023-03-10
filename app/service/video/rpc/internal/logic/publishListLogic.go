package logic

import (
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/service/video/dao/model"
	"SimpleDouYin/app/service/video/rpc/internal/svc"
	"SimpleDouYin/app/service/video/rpc/pb"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"log"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishListLogic) PublishList(in *pb.PublishListReq) (*pb.PublishListResp, error) {
	resp := new(pb.PublishListResp)

	//通过ID查询发布情况
	var publishs []model.Publish
	var pubCount int64
	err := l.svcCtx.GormDB.Where("user_id = ?", in.UserId).
		Find(&publishs).
		Count(&pubCount)
	if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
		log.Println("ID查询出错:", err.Error, " count:", pubCount)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}
	if pubCount == 0 {
		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "成功"
		return resp, nil
	}
	log.Println("publish查询成功,pubCount:", pubCount, "  publishes:", publishs)

	//查询视频
	var VIDS []int64
	var videos []model.Video
	for idx := 0; int64(idx) < pubCount; idx++ {
		VIDS = append(VIDS, publishs[idx].VideoID)
	}
	errr := l.svcCtx.GormDB.Model(&model.Video{}).
		Where("video_id in ?", VIDS).
		Find(&videos)
	if errr.Error != nil && (!errors.Is(errr.Error, gorm.ErrRecordNotFound)) {
		log.Println("查询出错:", errr.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}
	log.Println("video查询成功:", videos)

	//查询作者信息 只有一个
	var users []model.User
	var UIDS []int64
	for idx := 0; int64(idx) < pubCount; idx++ {
		UIDS = append(UIDS, publishs[idx].UserID)
	}

	errr = l.svcCtx.GormDB.Model(&model.User{}).
		Where("user_id in ?", UIDS).
		Order(tool.FiledStringBuild("user_id", UIDS)).
		Find(&users)
	if errr.Error != nil {
		log.Println("查询出错:", errr.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "服务器错误"
		return resp, nil
	}
	log.Println("users查询成功:", users)

	//关注情况
	//user对应follow关系
	var follow []model.Follow
	isF := true
	errr = l.svcCtx.GormDB.
		Where("uid = ? and target_uid = ?", in.UserId, users[0].UserID).
		First(&follow)
	if errr.Error != nil && (!errors.Is(errr.Error, gorm.ErrRecordNotFound)) {
		log.Println("查询出错:", errr.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	} else if errors.Is(errr.Error, gorm.ErrRecordNotFound) {
		isF = false
	}
	log.Println("follow查询成功:", follow)

	//整合结果
	for i := 0; int64(i) < pubCount; i++ {
		var Author pb.Author
		var resVi pb.Video

		//author
		Author.Id = users[0].UserID
		Author.FollowerCount = users[0].FollowCount
		Author.FollowerCount = users[0].FollowerCount
		Author.Name = users[0].Name
		Author.IsFollow = isF

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
	return resp, nil
}
