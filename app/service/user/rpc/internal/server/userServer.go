// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"SimpleDouYin/app/service/user/rpc/internal/logic"
	"SimpleDouYin/app/service/user/rpc/internal/svc"
	"SimpleDouYin/app/service/user/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterRes, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginReps, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) GetInfo(ctx context.Context, in *pb.GetInfoReq) (*pb.GetInfoReps, error) {
	l := logic.NewGetInfoLogic(ctx, s.svcCtx)
	return l.GetInfo(in)
}
