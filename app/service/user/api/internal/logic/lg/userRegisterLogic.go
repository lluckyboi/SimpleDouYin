package lg

import (
	"context"
	"go-zero-demo/app/service/user/rpc/user"
	"strconv"

	"go-zero-demo/app/service/user/api/internal/svc"
	"go-zero-demo/app/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	resp = new(types.RegisterResponse)
	//调rpc入库
	RgRes, err := l.svcCtx.UserClient.Register(l.ctx, &user.RegisterReq{
		UserTrueName: req.UserTrueName,
		UserNickName: req.UserNickName,
		Sex:          req.Sex,
		UserSchool:   req.UserSchool,
		UserStno:     req.UserStno,
		UserRole:     req.UserRole,
		UserPnum:     req.UserPnum,
		UserPassword: req.UserPassword,
		VCode:        req.VCode,
	})
	resp.Info = RgRes.Msg
	resp.Status, _ = strconv.Atoi(RgRes.Status)
	if err != nil {
		logx.Error(err)
		resp.Status = 500
		return resp, nil
	}
	return resp, nil
}
