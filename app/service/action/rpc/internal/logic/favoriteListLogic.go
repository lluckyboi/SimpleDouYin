package logic

import (
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/service/action/dao/model"
	"SimpleDouYin/app/service/action/rpc/internal/svc"
	"SimpleDouYin/app/service/action/rpc/pb"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteListLogic) FavoriteList(in *pb.FavoriteListReq) (*pb.FavoriteListResp, error) {
	resp := new(pb.FavoriteListResp)

	//查询点赞列表
	var (
		favorites []model.Favorite
		favCt     int64
	)
	err := l.svcCtx.GormDB.Where("user_id = ?", in.UserId).
		Find(&favorites).
		Count(&favCt)
	if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
		log.Println("ID查询出错:", err.Error, " count:", favCt)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}
	if favCt == 0 {
		resp.StatusCode = status.SuccessCode
		resp.StatusMsg = "成功"
		return resp, nil
	}
	log.Println("favorites查询成功", favCt)

	//查询视频
	var VIDS []int64
	var videos []model.Video
	for idx := 0; int64(idx) < favCt; idx++ {
		VIDS = append(VIDS, favorites[idx].VideoID)
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

	//查询publish
	var publishs []model.Publish
	errr = l.svcCtx.GormDB.Model(&model.Publish{}).
		Where("video_id in ?", VIDS).
		Order(tool.FiledStringBuild("video_id", VIDS)).
		Find(&publishs)
	if errr.Error != nil && (!errors.Is(errr.Error, gorm.ErrRecordNotFound)) {
		log.Println("查publish表错误", errr.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "服务器错误"
		return resp, err.Error
	}
	log.Println("publishs 查询成功", publishs)

	//查询作者
	var (
		usersTp []model.User
		UIDS    []int64
		userCt  int64
	)
	for idx := 0; int64(idx) < favCt; idx++ {
		UIDS = append(UIDS, publishs[idx].UserID)
	}
	errr = l.svcCtx.GormDB.Model(&model.User{}).
		Where("user_id in ?", UIDS).
		Order(tool.FiledStringBuild("user_id", UIDS)).
		Find(&usersTp).
		Count(&userCt)
	if errr.Error != nil {
		log.Println("作者查询出错:", errr.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "服务器错误"
		return resp, err.Error
	}
	log.Println("作者查询成功:", usersTp)
	//补全
	UserIdx := make(map[int64]int)
	for _, val := range favorites {
		for idx, uv := range usersTp {
			if val.UserID == uv.UserID {
				UserIdx[uv.UserID] = idx
			}
		}
	}
	var users []model.User
	for i := 0; int64(i) < favCt; i++ {
		users = append(users, usersTp[UserIdx[favorites[i].UserID]])
	}
	log.Println("users补全成功:", users)

	//user对应follow关系
	var Tfollows []model.Follow
	follows := make([]bool, favCt)
	for i := 0; int64(i) < favCt; i++ {
		follows[i] = false
	}

	errr1 := l.svcCtx.GormDB.
		Where("uid = ? and target_uid in ?", in.CurUser, UIDS).
		Find(&Tfollows)
	if errr1.Error != nil && (!errors.Is(errr1.Error, gorm.ErrRecordNotFound)) {
		log.Println("查询出错:", errr1.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "服务器错误"
		return resp, err.Error
	}
	for _, v := range Tfollows {
		for ku, vu := range UIDS {
			if vu == v.TargetUID {
				follows[ku] = true
			}
		}
	}
	log.Println("follows查询成功:", follows)

	//整合成结果
	for i := 0; int64(i) < favCt; i++ {
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
		resVi.Id = favorites[i].VideoID
		resVi.IsFavorite = true
		resVi.CommentCount = videos[i].CommentCount
		resVi.FavoriteCount = videos[i].FavoriteCount
		resVi.CoverUrl = videos[i].CoverURL
		resVi.PlayUrl = videos[i].PlayURL

		resp.VideoList = append(resp.VideoList, &resVi)
	}
	resp.StatusCode = status.SuccessCode
	resp.StatusMsg = "获取列表成功"
	return resp, nil
}
