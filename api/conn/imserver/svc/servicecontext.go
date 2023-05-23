package svc

import (
	"github.com/xu756/appserver/api/conn/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		RedisClient: redisClient,
		Config:      c,
	}
}
