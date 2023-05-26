package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConf redis.RedisConf
	Cache     cache.CacheConf
	DbSource  string
	ConnRpc   zrpc.RpcClientConf
}
