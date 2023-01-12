package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-demo/app/service/user/dao/model"
	"go-zero-demo/app/service/user/rpc/internal/svc"
	"go-zero-demo/app/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoByPnumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoByPnumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoByPnumLogic {
	return &GetInfoByPnumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoByPnumLogic) GetInfoByPnum(in *pb.GetInfoByPnumReq) (*pb.GetInfoByPnumReps, error) {
	resp := new(pb.GetInfoByPnumReps)

	//查询 在缓存中确认user_pnum存在后再调用
	user := model.User{}
	rst := l.svcCtx.GormDB.Where("user_pnum = ?", in.UserPnum).Find(&user)
	if rst.Error != nil {
		resp.Info = "服务器错误"
		resp.Status = "500"
		logx.Error("获取用户信息错误: 获取用户错误:", rst.Error.Error())
		return resp, nil
	}
	//拷贝
	err := copier.Copy(&resp, &user)
	if err != nil {
		resp.Info = "服务器错误"
		resp.Status = "500"
		logx.Error("用户数据拷贝错误: ", err)
		return resp, nil
	}

	resp.Info = "success"
	resp.Status = "200"
	return resp, nil
}
