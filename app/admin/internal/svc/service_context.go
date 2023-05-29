package svc

import (
	"github.com/xu756/appserver/app/admin/internal/config"
	"github.com/xu756/appserver/app/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	UserModel   model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DbSource)
	redisClient, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:      c,
		RedisClient: redisClient,
		UserModel:   model.NewUserModel(conn, c.Cache),
	}
}
