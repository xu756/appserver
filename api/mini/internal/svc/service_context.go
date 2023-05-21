package svc

import (
	"github.com/xu756/appserver/api/mini/internal/config"
	"github.com/xu756/appserver/api/mini/internal/middleware"
	"github.com/xu756/appserver/app/login/loginrpc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	MiniMidHandler rest.Middleware
	LoginRpc       loginrpc.LoginRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		MiniMidHandler: middleware.NewMiniMidHandlerMiddleware().Handle,
		LoginRpc:       loginrpc.NewLoginRpc(zrpc.MustNewClient(c.LoginRpc)),
	}
}
