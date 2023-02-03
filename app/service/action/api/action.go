package main

import (
	"SimpleDouYin/app/common/jwt"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"

	"SimpleDouYin/app/service/action/api/internal/config"
	"SimpleDouYin/app/service/action/api/internal/handler"
	"SimpleDouYin/app/service/action/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/action_conf.yaml", "the config file")

func main() {
	flag.Parse()
	logx.DisableStat()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	//有时第三方服务耗时较大 比如rpc默认两秒邮件有时候会超时 设置api超时时间30s rpc15秒
	c.Timeout = int64(30 * time.Second)
	c.ActionClient.Timeout = int64(15 * time.Second)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	//初始化JWTMap
	JWTMap := jwt.JWTMap{Keys: make(map[string]interface{})}

	ctx := svc.NewServiceContext(c, &JWTMap)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
