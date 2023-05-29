package svc

import (
	"github.com/xu756/appserver/api/admin/internal/config"
	"github.com/xu756/appserver/api/admin/internal/middleware"
	"github.com/xu756/appserver/internal/xjwt"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware
	Jwt    xjwt.JWT
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Auth:   middleware.NewAuthMiddleware().Handle,
		Jwt:    c.JWT,
	}
}
