package logic

import (
	"context"

	"SimpleDouYin/app/service/user/rpc/internal/svc"
	"SimpleDouYin/app/service/user/rpc/pb"

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

func (l *GetInfoByPnumLogic) GetInfoByPnum(in *pb.GetInfoReq) (*pb.GetInfoReps, error) {
	// todo: add your logic here and delete this line

	return &pb.GetInfoReps{}, nil
}
