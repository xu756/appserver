package main

import (
	"flag"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"server/cmd/admin/router"
	"server/common/config"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	hlog.Debugf("【 AdminApi 】addr on %s", config.RunData.Addr.AdminApiAddr)
	router.InitRouter()
	err := router.HttpServer.Run()
	if err != nil {
		hlog.Debug(err.Error())
		return
	}
}
