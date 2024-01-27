package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"server/common/config"
	"server/internal/middleware"
	"server/kitex_gen/user/usersrv"
)

var UserClient usersrv.Client

func InitUserClient(destService string) {
	s, err := usersrv.NewClient(destService,
		client.WithHostPorts(config.RunData.Rpc.UserRpcAddr),
		client.WithErrorHandler(middleware.ClientErrorHandler),
	)
	if err != nil {
		hlog.Debug("【user】client init error", err)
		panic(err)
	}
	UserClient = s
}
