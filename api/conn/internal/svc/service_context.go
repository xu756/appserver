package svc

import (
	"github.com/xu756/appserver/api/conn/internal/config"
	"github.com/xu756/appserver/app/im/imrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	ImRpc  imrpc.ImRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ImRpc:  imrpc.NewImRpc(zrpc.MustNewClient(c.ImRpc)),
	}
}
