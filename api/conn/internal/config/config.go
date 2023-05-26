package config

import (
	"github.com/xu756/appserver/internal/xjwt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type IMConfig struct {
	Port      int `json:"port"`
	Heartbeat int `json:"heartbeat"`
}

type Config struct {
	zrpc.RpcServerConf
	ImConfig  IMConfig `json:"ImConfig"`
	RedisConf redis.RedisConf
	Jwt       xjwt.JWT `json:"jwt"`
	ImRpc     zrpc.RpcClientConf
}
