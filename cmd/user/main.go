package main

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"net"
	"server/cmd/user/hander"
	"server/common/config"
	"server/internal/xerr"
	"server/kitex_gen/user/usersrv"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	addr, err := net.ResolveTCPAddr("tcp", config.RunData.Addr.UserRpcAddr)
	if err != nil {
		klog.Fatal(err)
	}
	svr := usersrv.NewServer(hander.NewUserImpl(),
		server.WithServiceAddr(addr),
		server.WithErrorHandler(xerr.ServerErrorHandler),
	)
	klog.Debugf("【user】server start at %s", config.RunData.Addr.UserRpcAddr)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
