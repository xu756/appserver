package config

import (
	"github.com/xu756/appserver/internal/xjwt"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	JWT xjwt.JWT
}
