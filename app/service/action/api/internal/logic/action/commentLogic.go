package action

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/service/action/rpc/action"
	"context"
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

func (l *CommentLogic) Comment(req *types.CommentReq) (resp *types.CommentResp, err error) {
	//解析token
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = status.ErrFailParseToken
		resp.StatusMsg = "token解析失败"
		logx.Error(err.Error())
		return resp, nil
	}

	//校验ActionType
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
	if act {
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
		logx.Error(err.Error())
		return resp, nil
	}
	resp.StatusCode = Grsp.StatusCode
	resp.StatusMsg = Grsp.StatusMsg
	return resp, nil
}
