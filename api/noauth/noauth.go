package main

import (
	"flag"
	"fmt"
	"github.com/xu756/appserver/api/noauth/internal/config"
	"github.com/xu756/appserver/api/noauth/internal/handler"
	"github.com/xu756/appserver/api/noauth/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/noauth-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("【启动普通api成功】 at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
