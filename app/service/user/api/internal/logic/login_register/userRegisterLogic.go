package login_register

import (
	"SimpleDouYin/app/service/user/api/internal/svc"
	"SimpleDouYin/app/service/user/api/internal/types"
	"SimpleDouYin/app/service/user/rpc/user"
	"context"

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
		Username: req.UserName,
		Password: req.PassWord,
	})
	resp.StatusMsg = RgRes.StatusMsg
	resp.StatusCode = RgRes.StatusCode
	return resp, nil
}
