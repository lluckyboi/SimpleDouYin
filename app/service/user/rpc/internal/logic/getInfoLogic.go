package logic

import (
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/dao/model"
	"SimpleDouYin/app/service/user/rpc/internal/svc"
	"SimpleDouYin/app/service/user/rpc/pb"
	"context"
	"errors"
	"gorm.io/gorm"

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
		resp.StatusCode = status.ErrOfServer
		resp.StatusMsg = status.InfoErrOfServer
		return resp, nil
	}

	//再查询是否关注
	db = l.svcCtx.GormDB.Where("uid = ? and target_uid = ?", in.UserId, in.TargetId).First(&Follow)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			isf = false
		} else {
			resp.StatusCode = status.ErrOfServer
			resp.StatusMsg = status.InfoErrOfServer
			return resp, nil
		}
	}

	resp.StatusCode = status.SuccessCode
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
