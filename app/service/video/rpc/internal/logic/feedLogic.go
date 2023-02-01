package logic

import (
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/service/video/dao/model"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"SimpleDouYin/app/service/video/rpc/internal/svc"
	"SimpleDouYin/app/service/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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

	//解析时间戳
	parse, err := time.Parse("2006-01-02T15:04:05", in.LastTime)
	if err != nil {
		log.Println(in.LastTime)
		resp.StatusCode = status.ErrParseTime
		resp.StatusMsg = "时间戳格式错误"
		return resp, nil
	}
	log.Print("解析时间戳完成", in.LastTime, "pares:", parse)

	//先去publish找到发布时间符合条件的
	var count int64
	errr := l.svcCtx.GormDB.Model(&model.Publish{}).
		Where("publish_time < ?", parse.Format("2006-01-02 15:04:05")).
		Order("publish_time desc").
		Limit(key.FeedNum).
		Find(&publishs).
		Count(&count)
	if errr.Error != nil && (!errors.Is(errr.Error, gorm.ErrRecordNotFound)) {
		log.Println("查publish表错误", errr.Error)
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "服务器错误"
		return resp, err
	}
	log.Println("publish查询成功:", publishs, count)
	//没有了
	if count == 0 {
		resp.StatusCode = http.StatusOK
		resp.StatusMsg = "成功"
		return resp, err
	}

	//查询对应user和video
	//查询对应user
	var UIDS []int64
	var strb strings.Builder
	var userCt int64
	strb.WriteString("FIELD(user_id")
	for idx := 0; int64(idx) < count; idx++ {
		UIDS = append(UIDS, publishs[idx].UserID)
		strb.WriteString(",")
		strb.WriteString(strconv.FormatInt(publishs[idx].VideoID, 10))
	}
	strb.WriteString(")")

	errr = l.svcCtx.GormDB.Model(&model.User{}).
		Where("user_id in ?", UIDS).
		Order(strb.String()).
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
	uidx := 0
	var users []model.User
	users = append(users, usersTp[uidx])
	for i := 1; int64(i) < count; i++ {
		if publishs[i].UserID == usersTp[uidx].UserID {
			users = append(users, usersTp[uidx])
		} else if int64(uidx+1) >= userCt {
			break
		} else {
			uidx++
		}
	}

	//user对应follow关系
	var Tfollows []model.Follow
	follows := make([]bool, count)
	for i := 0; int64(i) < count; i++ {
		follows[i] = false
	}
	//构造字符串
	strb.Reset()
	strb.WriteString("FIELD(uid")
	for idx := 0; int64(idx) < count; idx++ {
		UIDS = append(UIDS, publishs[idx].UserID)
		strb.WriteString(",")
		strb.WriteString(strconv.FormatInt(publishs[idx].VideoID, 10))
	}
	strb.WriteString(")")
	errr1 := l.svcCtx.GormDB.
		Where("uid = ? and target_uid in ?", in.UserId, UIDS).
		Order(strb.String()).
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
	strb.Reset()
	strb.WriteString("FIELD(video_id")
	for idx := 0; int64(idx) < count; idx++ {
		VIDS = append(VIDS, publishs[idx].VideoID)
		strb.WriteString(",")
		strb.WriteString(strconv.FormatInt(publishs[idx].VideoID, 10))
	}
	strb.WriteString(")")

	errr = l.svcCtx.GormDB.Model(&model.Video{}).
		Where("video_id in ?", VIDS).
		Order(strb.String()).
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
	resp.NextTime = publishs[count-1].PublishTime.Format("2006-01-02T15:04:05")
	return resp, nil
}
