package logic

import (
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/sec"
	"SimpleDouYin/app/common/snowFlake"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/service/user/dao/model"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"strconv"
	"strings"
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

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterRes, error) {
	regs := new(pb.RegisterRes)
	User := model.User{
		Username: in.Username,
		Name:     "用户" + time.Now().Format("2006-01-02T15:04:05"),
	}

	//雪花算法生成id
	sk, err := snowFlake.NewSnowflake(1, 1)
	if err != nil {
		logx.Error(err)
		return regs, nil
	}
	User.UserID = sk.NextVal()

	//比较时间戳，如果和上次的更大，说明时钟回拨
	cmd := l.svcCtx.Redis.Get("userid_last_timestamp")
	if cmd.Err() == nil {
		//正常，开始比较
		//redis结果处理
		_, val, _ := strings.Cut(l.svcCtx.Redis.Get(key.RedisUserIDLastTimeStamp).String(), tool.RedisStrBuilder(key.RedisUserIDLastTimeStamp))
		logx.Info(val)
		//解析
		lastTP, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			regs.StatusCode = status.ErrOfServer
			regs.StatusMsg = status.InfoErrOfServer
			logx.Error("时间戳解析错误：", err)
			return regs, nil
		}
		if lastTP >= User.UserID {
			regs.StatusCode = status.ErrOfServer
			regs.StatusMsg = status.InfoErrOfServer
			logx.Error("可能出现时钟回拨")
			return regs, nil
		} else {
			l.svcCtx.Redis.Set(key.RedisUserIDLastTimeStamp, strconv.FormatInt(snowFlake.GetTimestamp(User.UserID), 10), 0)
		}
	} else if errors.Is(redis.Nil, cmd.Err()) {
		//如果是第一次写入，不用比较
		l.svcCtx.Redis.Set(key.RedisUserIDLastTimeStamp, strconv.FormatInt(snowFlake.GetTimestamp(User.UserID), 10), 0)
	} else { //报错了
		regs.StatusCode = status.ErrOfServer
		regs.StatusMsg = status.InfoErrOfServer
		logx.Error("获取时间戳错误:", cmd.Err())
		return regs, nil
	}

	//三重DES加密
	if User.Password, err = sec.TripleDesEncrypt(in.Password, l.svcCtx.Config.Sec.DESKey, l.svcCtx.Config.Sec.DESIv); err != nil {
		regs.StatusCode = status.ErrOfServer
		regs.StatusMsg = status.InfoErrOfServer
		logx.Error("三重DES加密错误:", cmd.Err())
		return regs, nil
	}
	//入库
	ds := l.svcCtx.GormDB.Create(&User)
	if ds.Error != nil {
		regs.StatusMsg = "服务器错误"
		regs.StatusCode = status.ErrOfServer
		logx.Error("user Register 密码入库错误:", ds.Error)
		return regs, nil
	}
	//用户名和id写入缓存
	l.svcCtx.Redis.SAdd("username", in.Username)
	l.svcCtx.Redis.SAdd("user_id", User.UserID)

	regs.StatusCode = http.StatusOK
	regs.StatusMsg = "注册成功"
	regs.UserId = User.UserID
	return regs, nil
}
