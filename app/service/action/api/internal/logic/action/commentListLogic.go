package action

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/service/action/rpc/action"
	"context"
	"strconv"

	"SimpleDouYin/app/service/action/api/internal/svc"
	"SimpleDouYin/app/service/action/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListReq) (*types.CommentListResp, error) {
	resp := new(types.CommentListResp)
	//解析token
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = status.ErrFailParseToken
		resp.StatusMsg = "token解析失败"
		logx.Error(err.Error())
		return resp, nil
	}

	//解析ID
	vid, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "UID有误或服务器错误"
		return resp, nil
	}

	//rpc
	Grsp, err := l.svcCtx.ActionClient.CommentList(l.ctx, &action.CommentListReq{
		UserId:  claims.UserId,
		VideoId: vid,
	})
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		logx.Error(err.Error())
		return resp, nil
	}
	resp.StatusCode = Grsp.StatusCode
	resp.StatusMsg = Grsp.StatusMsg

	ComList := Grsp.CommentList
	for _, v := range ComList {
		author := types.Author{
			Id:            v.User.Id,
			Name:          v.User.Name,
			FollowCount:   v.User.FollowCount,
			FollowerCount: v.User.FollowerCount,
			IsFollow:      v.User.IsFollow,
		}
		com := types.Comment{
			Id:         v.Id,
			User:       author,
			Content:    v.Content,
			CreateDate: v.CreateDate,
		}
		resp.CommentList = append(resp.CommentList, com)
	}
	return resp, nil
}
