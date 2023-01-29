package video

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/service/video/rpc/videosv"
	"context"
	"log"
	"time"

	"SimpleDouYin/app/service/video/api/internal/svc"
	"SimpleDouYin/app/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedRequest) (*types.FeedResponse, error) {
	resp := new(types.FeedResponse)
	rpcReq := &videosv.FeedReq{}

	if req.Token != "" {
		//解析token
		claims, err := jwt.ParseToken(req.Token)
		if err != nil {
			resp.StatusCode = status.ErrFailParseToken
			resp.StatusMsg = "token解析失败"
			logx.Error(err.Error())
			return resp, nil
		}
		rpcReq.UserId = claims.UserId
	} else {
		rpcReq.UserId = -1
	}

	//检查时间戳
	if req.LastTime == "" || req.LastTime == "0" { //为空 默认当前时间
		rpcReq.LastTime = time.Now().Format("2006-01-02T15:04:05")
	} else { //检查时间戳格式
		log.Println(req.LastTime)
		parse, err := time.Parse("2006-01-02T15:04:05", req.LastTime)
		if err != nil {
			resp.StatusCode = status.ErrParseTime
			resp.StatusMsg = "时间戳格式错误"
			return resp, nil
		}
		log.Println("解析时间戳成功:", parse)
		rpcReq.LastTime = parse.Format("2006-01-02T15:04:05")
	}

	//调用rpc
	Grsp, err := l.svcCtx.VideoClient.Feed(l.ctx, rpcReq)
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		logx.Error(err.Error())
		return resp, nil
	}
	log.Print("feed rpc 成功")
	resp.StatusCode = Grsp.StatusCode
	resp.StatusMsg = Grsp.StatusMsg
	resp.NextTime = Grsp.NextTime

	VideoList := Grsp.VideoList

	for _, v := range VideoList {
		author := types.Author{
			Id:            v.Author.Id,
			Name:          v.Author.Name,
			FollowCount:   v.Author.FollowCount,
			FollowerCount: v.Author.FollowerCount,
			IsFollow:      v.Author.IsFollow,
		}
		video := types.Video{
			Id:            v.Id,
			Author:        author,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		}
		resp.VideoList = append(resp.VideoList, video)
	}
	log.Println("videolist:", len(resp.VideoList))
	return resp, nil
}
