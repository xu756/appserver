package main

import (
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"server/cmd/user/hander"
	"server/common/config"
	"server/internal/middleware"
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
		server.WithErrorHandler(middleware.ServerErrorHandler),
		//server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
	)
	log.Printf("【user】server start at %s", config.RunData.Addr.UserRpcAddr)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
