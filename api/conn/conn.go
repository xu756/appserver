package main

import (
	"flag"
	"fmt"
	"github.com/xu756/appserver/api/conn/imserver"

	"github.com/xu756/appserver/api/conn/internal/config"
	"github.com/xu756/appserver/api/conn/internal/server"
	"github.com/xu756/appserver/api/conn/internal/svc"
	"github.com/xu756/appserver/api/conn/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/conn.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	imserver.InitServer(c)
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterImServiceServer(grpcServer, server.NewImServiceServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}