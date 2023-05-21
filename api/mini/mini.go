package main

import (
	"flag"
	"fmt"

	"github.com/xu756/appserver/api/mini/internal/config"
	"github.com/xu756/appserver/api/mini/internal/handler"
	"github.com/xu756/appserver/api/mini/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/mini-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("【启动小程序api成功】 at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
