package imserver

import (
	"fmt"
	"github.com/xu756/appserver/api/conn/imserver/logic"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/xu756/appserver/api/conn/imserver/handler"
	"github.com/xu756/appserver/api/conn/imserver/svc"
	"github.com/xu756/appserver/api/conn/internal/config"
)

func InitServer(c config.Config) {

	ctx := svc.NewServiceContext(c)
	server := http.Server{
		Addr:         ":" + strconv.Itoa(c.ImConfig.Port),
		ReadTimeout:  time.Millisecond * 500,
		WriteTimeout: time.Millisecond * 500,
	}
	handler.RegisterHandlers(&server, ctx)

	fmt.Println("IM服务已启动：", c.ImConfig.Heartbeat)
	defer func(server *http.Server) {
		if err := server.Close(); err != nil {
			logx.Errorf("【IM-Api】程序退出，IM服务关闭失败 %v", err)
		}
	}(&server)
	logic.Hostname, _ = os.Hostname()

	logic.Heartbeat = c.ImConfig.Heartbeat
	err := ctx.RedisClient.Setex("cache:im:server:"+logic.Hostname+logic.ImDefault, logic.Hostname+logic.ImDefault, 60)
	if err != nil {
		logx.Errorf("【IM-Api】设置服务器失败 %v", err)
		return
	}
	go imServer(ctx)
	logic.ImDefault = c.ImConfig.ImDefault
	err = server.ListenAndServe()
	if err != nil {
		logx.Errorf("【IM-Api】服务停止 %v", err)
	}
}

func imServer(svc *svc.ServiceContext) {
	// 每60向数据库写入一次心跳
	ticker := time.NewTicker(time.Second * 55)
	for {
		select {
		case <-ticker.C:
			err := svc.RedisClient.Setex("cache:im:server:"+logic.Hostname+logic.ImDefault, logic.Hostname+logic.ImDefault, 60)
			if err != nil {
				logx.Errorf("【IM-Api】设置服务器失败 %v", err)
				return
			}
		}
	}
}
