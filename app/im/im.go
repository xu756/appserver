package main

import (
	"flag"
	"fmt"

	"github.com/xu756/appserver/app/im/internal/config"
	"github.com/xu756/appserver/app/im/internal/server"
	"github.com/xu756/appserver/app/im/internal/svc"
	"github.com/xu756/appserver/app/im/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterImLoginServer(grpcServer, server.NewImLoginServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("【启动Im Rpc】 at %s...\n", c.ListenOn)
	s.Start()
}
