package svc

import (
	"github.com/xu756/appserver/api/conn/internal/config"
	"github.com/xu756/appserver/internal/xjwt"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	Jwt         xjwt.JWT
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		RedisClient: redisClient,
		Config:      c,
		Jwt:         c.Jwt,
	}
}
