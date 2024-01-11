package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"server/common/config"
	"server/internal/middleware"
)

var HttpServer *server.Hertz

func InitRouter() {
	h := server.Default(
		server.WithHostPorts(config.RunData.Addr.ApiAddr),
		server.WithReadBufferSize(1024*1024*100),
		server.WithMaxRequestBodySize(1024*1024*100),
	)
	h.Use(middleware.HertzJwt())
	//router := h.Group("/wxapp")
	//PublicRoute(router.Group("/public"))
	HttpServer = h
}
