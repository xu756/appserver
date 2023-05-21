package svc

import (
	"github.com/wenlng/go-captcha/captcha"
	"github.com/xu756/appserver/app/login/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	Captcha     *captcha.Captcha
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient, err := redis.NewRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	capt := captcha.NewCaptcha()
	// todo 验证码配置 https://github.com/wenlng/go-captcha/blob/master/README_zh.md

	return &ServiceContext{
		Config:      c,
		RedisClient: redisClient,
		Captcha:     capt,
	}
}
