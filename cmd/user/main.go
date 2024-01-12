package main

import (
	"context"
	"flag"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"server/cmd/user/hander"
	"server/common/config"
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
		server.WithErrorHandler(ServerErrorHandler),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
	)
	log.Printf("【user】server start at %s", config.RunData.Addr.UserRpcAddr)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}

func ServerErrorHandler(ctx context.Context, err error) error {
	log.Print("ser")
	return err

}
