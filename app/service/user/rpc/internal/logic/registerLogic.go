package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/app/common"
	"go-zero-demo/app/service/user/dao/model"
	"go-zero-demo/app/service/user/rpc/internal/svc"
	"go-zero-demo/app/service/user/rpc/pb"
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

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterRes, error) {
	res := new(pb.RegisterRes)
	//校验验证码
	vc := l.svcCtx.Redis.Get(in.UserPnum + in.VCode).String()
	cidx := 0
	for {
		n := len(vc[:])
		if vc[cidx] != ':' {
			cidx++
			if cidx > n {
				break
			}
			continue
		}
		cidx += 2
		break
	}
	if vc[cidx:] != in.VCode {
		res.Msg = "验证码错误"
		res.Status = "201"
		return res, nil
	}

	//查询是否注册
	bl := l.svcCtx.Redis.SIsMember("user_pnum", in.UserPnum)
	if bl.Val() == true {
		res.Msg = "手机号已经注册过了"
		res.Status = "201"
		return res, nil
	}

	User := model.User{
		UserTrueName: in.UserTrueName,
		UserNickName: in.UserNickName,
		Sex:          in.Sex,
		UserSchool:   in.UserSchool,
		UserStno:     in.UserStno,
		UserRole:     in.UserRole,
		UserPnum:     in.UserPnum,
		UserPassword: in.UserPassword,
	}

	//RAS 密码公钥加密
	User.UserPassword = common.RSA_Encrypt([]byte(in.UserPassword), "public.pem")
	//入库
	ds := l.svcCtx.GormDB.Create(&User)
	if ds.Error != nil {
		res.Msg = "服务器错误"
		res.Status = "500"
		return res, nil
	}

	//写入缓存 集合
	l.svcCtx.Redis.SAdd("user_pnum", in.UserPnum)
	res.Msg = "注册成功"
	res.Status = "200"
	return res, nil
}
