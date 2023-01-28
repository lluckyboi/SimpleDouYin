package main

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/key"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"

	"SimpleDouYin/app/service/video/api/internal/config"
	"SimpleDouYin/app/service/video/api/internal/handler"
	"SimpleDouYin/app/service/video/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/video.yaml", "the config file")

func main() {
	flag.Parse()
	logx.DisableStat()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	//有时第三方服务耗时较大 比如rpc默认两秒邮件有时候会超时 设置video-api超时时间15s video-rpc5秒
	c.Timeout = int64(15 * time.Second)
	c.VideoClient.Timeout = int64(5 * time.Second)
	//限制文件大小
	c.MaxBytes = key.MAXBytes
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//初始化JWTMap
	JWTMap := jwt.JWTMap{Keys: make(map[string]interface{})}

	ctx := svc.NewServiceContext(c, &JWTMap)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
