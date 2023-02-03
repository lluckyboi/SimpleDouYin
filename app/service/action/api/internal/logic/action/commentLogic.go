package action

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/service/action/rpc/action"
	"context"
	"log"
	"strconv"

	"SimpleDouYin/app/service/action/api/internal/svc"
	"SimpleDouYin/app/service/action/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLogic) Comment(req *types.CommentReq) (*types.CommentResp, error) {
	resp := new(types.CommentResp)
	//解析token
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = status.ErrFailParseToken
		resp.StatusMsg = "token解析失败"
		logx.Error(err.Error())
		return resp, nil
	}

	//校验ActionType 0-add-true 1-delete-false
	act, err := tool.AcTypeStringToBool(req.ActionType)
	if err != nil {
		resp.StatusCode = status.ErrUnknownAcType
		resp.StatusMsg = "unknown ActionType"
		return resp, nil
	}

	//解析ID
	vid, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = "VID有误或服务器错误"
		return resp, nil
	}

	var cid int64 = 0
	if !act {
		cid, err = strconv.ParseInt(req.CommentId, 10, 64)
		if err != nil {
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = "VID有误或服务器错误"
			return resp, nil
		}
	} else {
		//长度校验
		if !tool.CommentLengthCheck(req.CommentText) {
			resp.StatusCode = status.ErrLengthErr
			resp.StatusMsg = "长度有误"
			return resp, nil
		}
	}

	//rpc
	Grsp, err := l.svcCtx.ActionClient.Comment(l.ctx, &action.CommentReq{
		UserId:      claims.UserId,
		VideoId:     vid,
		ActionType:  act,
		CommentText: req.CommentText,
		CommentId:   cid,
	})
	if err != nil {
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		log.Println("rpc err:", err)
		return resp, nil
	}
	resp.StatusCode = Grsp.StatusCode
	resp.StatusMsg = Grsp.StatusMsg
	log.Println("comment rpc成功")

	//resp.comment
	cm := Grsp.Comment
	user := types.Author{
		Id:            cm.User.Id,
		Name:          cm.User.Name,
		FollowCount:   cm.User.FollowerCount,
		FollowerCount: cm.User.FollowCount,
		IsFollow:      cm.User.IsFollow,
	}
	log.Println(user)
	comment := types.Comment{
		Id:         cm.Id,
		User:       user,
		Content:    cm.Content,
		CreateDate: cm.CreateDate,
	}
	log.Print("last")
	resp.Comment = comment
	return resp, nil
}
