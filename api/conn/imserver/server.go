package imserver

import (
	"fmt"
	"github.com/xu756/appserver/api/conn/imserver/logic"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
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

	fmt.Println("【IM-Api】Api at: 0.0.0.0", c.ImConfig.Port)
	defer func(server *http.Server) {
		if err := server.Close(); err != nil {
			logx.Errorf("【IM-Api】程序退出，IM服务关闭失败 %v", err)
		}
	}(&server)
	logic.Hostname = c.Etcd.Key
	logic.Heartbeat = c.ImConfig.Heartbeat
	err := ctx.RedisClient.Setex("cache:im:server:"+logic.Hostname, logic.Hostname, 60)
	if err != nil {
		logx.Errorf("【IM-Api】设置服务器失败 %v", err)
		return
	}
	go imServer(ctx)
	go logic.Hubs.Run()
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
			err := svc.RedisClient.Setex("cache:im:server:"+logic.Hostname, logic.Hostname, 60)
			if err != nil {
				logx.Errorf("【IM-Api】设置服务器失败 %v", err)
				return
			}
		}
	}
}
