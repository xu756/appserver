package svc

import (
	"github.com/xu756/appserver/api/conn/imservice"
	"github.com/xu756/appserver/app/im/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	ConnRpc     imservice.ImService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		ConnRpc: imservice.NewImService(zrpc.MustNewClient(c.ConnRpc)),
	}
}
