package svc

import (
	"github.com/xu756/appserver/api/admin/internal/config"
	"github.com/xu756/appserver/app/admin/adminrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	AdminRpc adminrpc.AdminRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		AdminRpc: adminrpc.NewAdminRpc(zrpc.MustNewClient(c.AdminRpc)),
	}
}
