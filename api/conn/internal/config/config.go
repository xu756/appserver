package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type IMConfig struct {
	Port      int    `json:"port"`
	Heartbeat int    `json:"heartbeat"`
	ImDefault string `json:"im_default"`
}

type Config struct {
	zrpc.RpcServerConf
	ImConfig  IMConfig `json:"ImConfig"`
	RedisConf redis.RedisConf
}
