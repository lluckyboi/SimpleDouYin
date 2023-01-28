package helper

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/snowFlake"
	"SimpleDouYin/app/common/status"
	"SimpleDouYin/app/common/tool"
	"SimpleDouYin/app/service/video/api/internal/svc"
	"SimpleDouYin/app/service/video/api/internal/types"
	"SimpleDouYin/app/service/video/dao/model"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// MinioUpload 上传文件到Minio
func MinioUpload(r *http.Request, svcCtx *svc.ServiceContext, w http.ResponseWriter, req *types.PublishRequest) error {

	//解析token，检查是否合法
	_, err := jwt.ParseToken(req.Token)
	if err != nil {
		log.Println("token解析失败", err)
		httpx.OkJson(w, &types.PublishResponse{
			StatusCode: status.ErrFailParseToken,
			StatusMsg:  "无效的token",
		})
		return err
	}

	// 获取文件信息
	file, fileHeader, err := r.FormFile("data")
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return err
	}
	log.Print("获取文件信息成功:", fileHeader.Size)

	//雪花算法生成ID
	sf, err := snowFlake.NewSnowflake(1, 2)
	ID := sf.NextVal()
	//比较时间戳，检查是否发生时钟回拨
	cmd := svcCtx.RedisDB.Get(key.RedisVideoIDLastTimeStamp)
	if cmd.Err() == nil {
		//正常，开始比较
		//redis结果处理
		_, val, _ := strings.Cut(svcCtx.RedisDB.Get(key.RedisVideoIDLastTimeStamp).String(), tool.RedisStrBuilder(key.RedisVideoIDLastTimeStamp))
		logx.Info(val)
		//解析
		lastTP, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return err
		}
		if lastTP >= ID {
			httpx.ErrorCtx(r.Context(), w, errors.New("时钟回拨"))
			return err
		} else {
			svcCtx.RedisDB.Set(key.RedisVideoIDLastTimeStamp, strconv.FormatInt(snowFlake.GetTimestamp(ID), 10), 0)
		}
	} else if errors.Is(redis.Nil, cmd.Err()) || cmd.Err().Error() == "redis: nil" {
		//如果是第一次写入，不用比较
		svcCtx.RedisDB.Set(key.RedisVideoIDLastTimeStamp, strconv.FormatInt(snowFlake.GetTimestamp(ID), 10), 0)
	} else { //报错了
		httpx.ErrorCtx(r.Context(), w, err)
		return err
	}
	//ID没问题，可以向下传递
	req.ID = ID
	log.Println("生成ID成功", ID)

	// 判断文件是否存在
	b := make([]byte, fileHeader.Size)
	_, err = file.Read(b)
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return err
	}
	hash := fmt.Sprintf("%x", md5.Sum(b))
	//偏移值归零
	_, err = file.Seek(0, 0)
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return err
	}

	//数据库里查
	video := model.Video{}
	db := svcCtx.GormDB.Where("hash = ?", hash).First(&video)
	if db.Error == nil { //如果存在相同的hash 不用再上传 直接入库
		log.Println("存在相同的hash 不用再上传 直接入库")
		req.PlayUrl = video.PlayURL
		req.CoverUrl = video.CoverURL
		req.Hash = video.Hash
		return nil
	} else if errors.Is(db.Error, gorm.ErrRecordNotFound) || db.Error.Error() == "record not found" { //如果没有相同的hash 需要上传
		log.Println("没有相同的hash 需要上传")
		req.Hash = hash

		var builder strings.Builder
		builder.WriteString(strconv.FormatInt(ID, 10))
		builder.WriteString(path.Ext(fileHeader.Filename))
		ObjectName := builder.String()

		_, err = svcCtx.Minio.PutObject(context.Background(),
			svcCtx.Config.Minio.Buckets,
			ObjectName, file, fileHeader.Size,
			minio.PutObjectOptions{ContentType: "binary/octet-stream"})
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			log.Println("上传出错", err.Error())
			return err
		}

		var bd1 strings.Builder
		bd1.WriteString("http://")
		bd1.WriteString(svcCtx.Config.Minio.EndPoint)
		bd1.WriteString("/")
		bd1.WriteString(svcCtx.Config.Minio.Buckets)
		bd1.WriteString("/")
		bd1.WriteString(ObjectName)
		req.PlayUrl = bd1.String()
		req.CoverUrl = "https://typora.fengxiangrui.top/1674827367.png"
		return nil
	} else { //出错了
		httpx.ErrorCtx(r.Context(), w, db.Error)
		fmt.Println("出错了", db.Error.Error())
		return db.Error
	}
}
