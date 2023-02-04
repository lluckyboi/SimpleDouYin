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

type FollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowListLogic) FollowList(req *types.FollowListReq) (*types.FollowListResp, error) {
	resp := new(types.FollowListResp)

	//解析token
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = strconv.Itoa(status.ErrFailParseToken)
		resp.StatusMsg = "token解析失败"
		logx.Error(err.Error())
		return resp, nil
	}

	//解析ID
	uid, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		resp.StatusCode = strconv.Itoa(status.ErrOfServer)
		resp.StatusMsg = "UID有误或服务器错误"
		return resp, nil
	}

	//rpc
	Grsp, err := l.svcCtx.ActionClient.FollowList(l.ctx, &action.FollowListReq{
		UserId:  uid,
		CurUser: claims.UserId,
	})
	if err != nil {
		resp.StatusCode = strconv.Itoa(status.ErrOfServer)
		resp.StatusMsg = status.InfoErrOfServer
		logx.Error(err.Error())
		return resp, nil
	}

	resp.StatusCode = strconv.Itoa(int(Grsp.StatusCode))
	resp.StatusMsg = Grsp.StatusMsg

	//userList 赋值
	UserList := Grsp.UserList
	for _, v := range UserList {
		user := types.Author{
			Id:            v.Id,
			Name:          v.Name,
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowerCount,
			IsFollow:      v.IsFollow,
		}
		resp.UserList = append(resp.UserList, user)
	}
	return resp, nil
}
