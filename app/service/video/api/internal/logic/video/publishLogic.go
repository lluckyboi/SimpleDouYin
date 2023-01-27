package video

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/service/video/api/internal/svc"
	"SimpleDouYin/app/service/video/api/internal/types"
	"SimpleDouYin/app/service/video/rpc/videosv"
	"context"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishRequest) (*types.PublishResponse, error) {
	resp := new(types.PublishResponse)
	//解析token
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = status.ErrFailParseToken
		logx.Error(err.Error())
		return resp, nil
	}

	//rpc入库
	Grsp, err := l.svcCtx.VideoClient.Publish(l.ctx, &videosv.PublishReq{
		UserId:   claims.UserId,
		Title:    req.Title,
		VideoUrl: req.PlayUrl,
		VideoId:  req.ID,
		Hash:     req.Hash,
		CoverUrl: req.CoverUrl,
	})
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		logx.Error(err.Error())
		return resp, nil
	}
	log.Print("video publish rpc成功")
	resp.StatusCode = Grsp.StatusCode
	resp.StatusMsg = Grsp.StatusMsg
	return resp, nil
}
