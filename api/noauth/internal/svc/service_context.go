package svc

import (
	"github.com/xu756/appserver/api/noauth/internal/config"
	"github.com/xu756/appserver/app/login/loginrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	LoginRpc loginrpc.LoginRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		LoginRpc: loginrpc.NewLoginRpc(zrpc.MustNewClient(c.LoginRpc)),
	}
}
