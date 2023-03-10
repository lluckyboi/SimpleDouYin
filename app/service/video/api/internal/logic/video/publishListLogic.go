package video

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/service/video/dao/model"
	"SimpleDouYin/app/service/video/rpc/videosv"
	"context"
	"errors"
	"gorm.io/gorm"
	"strconv"

	"SimpleDouYin/app/service/video/api/internal/svc"
	"SimpleDouYin/app/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.PublishListRequest) (*types.PublishListResponse, error) {
	resp := new(types.PublishListResponse)
	//解析token
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = status.ErrFailParseToken
		resp.StatusMsg = "token解析失败"
		logx.Error(err.Error())
		return resp, nil
	}
	//解析ID
	uid, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "UID有误或解析错误"
		return resp, nil
	}
	//哈希取模
	sufId := tool.Hash_Mode(req.UserId, key.RedisHashMod)
	//查询id是否存在
	bl := l.svcCtx.RedisDB.SIsMember(key.RedisUserIdCacheKey+sufId, uid)
	if bl.Val() == false {
		res := l.svcCtx.GormDB.Where("user_id = ?", uid).First(&model.User{})
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			resp.StatusMsg = "用户ID不存在"
			resp.StatusCode = status.ErrNoSuchUser
			return resp, nil
		} else if res.Error != nil {
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			logx.Error("校验用户ID错误：", res.Error)
			return resp, nil
		}
		//回写缓存
		l.svcCtx.RedisDB.SAdd(key.RedisUserIdCacheKey+sufId, uid)
	}

	Grsp, err := l.svcCtx.VideoClient.PublishList(l.ctx, &videosv.PublishListReq{
		UserId: uid,
	})
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		logx.Error(err.Error())
		return resp, nil
	}

	resp.StatusCode = Grsp.StatusCode
	resp.StatusMsg = Grsp.StatusMsg

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
	return resp, nil
}
