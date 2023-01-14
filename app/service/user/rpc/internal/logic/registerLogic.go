package logic

import (
	"SimpleDouYin/app/common"
	"SimpleDouYin/app/service/user/dao/model"
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
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

	//雪花算法生成id
	sk, err := common.NewSnowflake(1, 1)
	if err != nil {
		logx.Error(err)
		return
	}
	User.UserID = sk.NextVal()

	//比较时间戳，如果和上次的更大，说明时钟回拨
	cmd := l.svcCtx.Redis.Get("userid_last_timestamp")
	switch cmd.Err() {
	//如果是第一次写入，不用比较
	case redis.Nil:
		{
			l.svcCtx.Redis.Set("userid_last_timestamp", strconv.FormatInt(common.GetTimestamp(User.UserID), 10), 0)
		}
	//如果没有错误，开始比较
	case nil:
		{
			lastTP, err := strconv.ParseInt(l.svcCtx.Redis.Get("userid_last_timestamp").String(), 10, 64)
			if err != nil {
				logx.Error("时间戳解析错误：", err)
				return
			}
			if lastTP >= User.UserID {
				logx.Error("可能出现时钟回拨")
				return
			} else {
				l.svcCtx.Redis.Set("userid_last_timestamp", strconv.FormatInt(common.GetTimestamp(User.UserID), 10), 0)
			}
		}
	//报错了 返回
	default:
		{
			logx.Error("获取时间戳错误:", cmd.Err())
		}
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

	//用户名写入缓存
	l.svcCtx.Redis.SAdd("username", in.Username)
	regs.StatusCode = 200
	regs.StatusMsg = "注册成功"
	regs.UserId = User.UserID
	return regs, nil
}
