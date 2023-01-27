package helper

import (
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/snowFlake"
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
	"net/http"
	"path"
	"strconv"
	"strings"
)

// MinioUpload 上传文件到Minio
func MinioUpload(r *http.Request, svcCtx *svc.ServiceContext, w http.ResponseWriter, req *types.PublishRequest) {
	// 获取文件信息
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

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
			return
		}
		if lastTP >= ID {
			httpx.ErrorCtx(r.Context(), w, errors.New("时钟回拨"))
			return
		} else {
			svcCtx.RedisDB.Set(key.RedisVideoIDLastTimeStamp, strconv.FormatInt(snowFlake.GetTimestamp(ID), 10), 0)
		}
	} else if errors.Is(redis.Nil, cmd.Err()) {
		//如果是第一次写入，不用比较
		svcCtx.RedisDB.Set(key.RedisVideoIDLastTimeStamp, strconv.FormatInt(snowFlake.GetTimestamp(ID), 10), 0)
	} else { //报错了
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	//ID没问题，可以向下传递
	req.ID = ID

	// 判断文件是否存在
	b := make([]byte, fileHeader.Size)
	_, err = file.Read(b)
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	hash := fmt.Sprintf("%x", md5.Sum(b))

	//数据库里查
	video := model.Video{}
	db := svcCtx.GormDB.Where("hash = ?", hash).First(&video)
	if db.Error == nil { //如果存在相同的hash 不用再上传 直接入库
		req.PlayUrl = video.PlayURL
		req.CoverUrl = video.CoverURL
		req.Hash = video.Hash
		return
	} else if errors.Is(db.Error, gorm.ErrRecordNotFound) { //如果没有相同的hash 需要上传
		req.Hash = hash

		var builder strings.Builder
		builder.WriteString(strconv.FormatInt(ID, 10))
		builder.WriteString(path.Ext(fileHeader.Filename))
		ObjectName := builder.String()

		_, err = svcCtx.Minio.PutObject(context.Background(), svcCtx.Config.Minio.Buckets[0], ObjectName, file, fileHeader.Size,
			minio.PutObjectOptions{ContentType: "binary/octet-stream"})
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		var bd1 strings.Builder
		bd1.WriteString("http://")
		bd1.WriteString(svcCtx.Config.Minio.EndPoint)
		bd1.WriteString("/")
		bd1.WriteString(svcCtx.Config.Minio.Buckets[0])
		bd1.WriteString("/")
		bd1.WriteString(ObjectName)
		req.PlayUrl = bd1.String()
		req.CoverUrl = "https://typora.fengxiangrui.top/1674827367.png"
		return
	} else { //出错了
		httpx.ErrorCtx(r.Context(), w, db.Error)
		return
	}
}
