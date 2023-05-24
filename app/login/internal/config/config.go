package config

import (
	"github.com/xu756/appserver/internal/xjwt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConf redis.RedisConf
	Jwt       xjwt.JWT
	Cache     cache.CacheConf
	DbSource  string
}
