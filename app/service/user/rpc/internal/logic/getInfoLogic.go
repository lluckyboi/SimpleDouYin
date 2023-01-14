package logic

import (
	"SimpleDouYin/app/common"
	"SimpleDouYin/app/service/user/dao/model"
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"

	"SimpleDouYin/app/service/user/rpc/internal/svc"
	"SimpleDouYin/app/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoLogic) GetInfo(in *pb.GetInfoReq) (*pb.GetInfoReps, error) {
	resp := new(pb.GetInfoReps)
	User := model.User{}
	Follow := model.Follow{}
	isf := true

	//先查询目标id信息
	db := l.svcCtx.GormDB.Where("user_id = ?", in.TargetId).First(&User)
	if db.Error != nil {
		resp.StatusCode = common.ErrOfServer
		resp.StatusMsg = common.InfoErrOfServer
		return resp, nil
	}

	//再查询是否关注
	db = l.svcCtx.GormDB.Where("uid = ? and target_id = ?", in.UserId, in.TargetId).First(&Follow)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			isf = false
		} else {
			resp.StatusCode = common.ErrOfServer
			resp.StatusMsg = common.InfoErrOfServer
			return resp, nil
		}
	}

	resp.StatusCode = http.StatusOK
	resp.StatusMsg = "查询成功"
	resp.User = &pb.UserInfo{
		UserId:        User.UserID,
		Name:          User.Name,
		FollowCount:   User.FollowCount,
		FollowerCount: User.FollowerCount,
		IsFollow:      isf,
	}
	return resp, nil
}
