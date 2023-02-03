package action

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/service/action/rpc/action"
	"context"
	"log"
	"strconv"

	"SimpleDouYin/app/service/action/api/internal/svc"
	"SimpleDouYin/app/service/action/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteListReq) (*types.FavoriteListResp, error) {
	resp := new(types.FavoriteListResp)
	//解析token
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = strconv.Itoa(status.ErrFailParseToken)
		resp.StatusMsg = "token解析失败"
		logx.Error(err.Error())
		return resp, nil
	}

	//解析ID
	uid, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		resp.StatusCode = strconv.Itoa(status.ErrOfServer)
		resp.StatusMsg = "UID有误或服务器错误"
		return resp, nil
	}

	//rpc
	Grsp, err := l.svcCtx.ActionClient.FavoriteList(l.ctx, &action.FavoriteListReq{
		UserId: uid,
	})
	if err != nil {
		resp.StatusCode = strconv.Itoa(status.ErrOfServer)
		resp.StatusMsg = status.InfoErrOfServer
		logx.Error(err.Error())
		return resp, nil
	}
	resp.StatusCode = strconv.Itoa(int(Grsp.StatusCode))
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
	log.Print(resp)
	return resp, nil
}
