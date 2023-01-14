package logic

import (
	"SimpleDouYin/app/common"
	"SimpleDouYin/app/service/user/dao/model"
	"context"
	"time"

	"SimpleDouYin/app/service/user/rpc/internal/svc"
	"SimpleDouYin/app/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (regs *pb.RegisterRes, err error) {
	User := model.User{
		Username: in.Username,
		Name:     "用户" + time.Now().String(),
	}

	//RSA 密码公钥加密
	User.Password = common.RSA_Encrypt([]byte(in.Password), l.svcCtx.Config.Sec.SecPub)

	//入库
	ds := l.svcCtx.GormDB.Create(&User)
	if ds.Error != nil {
		regs.StatusMsg = "服务器错误"
		regs.StatusCode = 500
		logx.Error("user Register 密码入库错误:", ds.Error)
		return regs, nil
	}

	//查id 未使用连接池,直接查
	var id []int64
	l.svcCtx.GormDB.Raw("select Last_INSERT_ID() as id").Pluck("id", &id)

	//用户名写入缓存
	l.svcCtx.Redis.SAdd("username", in.Username)
	regs.StatusCode = 200
	regs.StatusMsg = "注册成功"
	regs.UserId = id[0]
	return regs, nil
}
