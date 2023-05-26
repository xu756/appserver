package logic

import (
	"context"
	"github.com/thinkeridea/go-extend/exnet"
	"github.com/xu756/appserver/api/conn/imserver/svc"
	"github.com/xu756/appserver/app/im/imrpc"
	"github.com/xu756/appserver/internal/ctxdata"
	"github.com/xu756/appserver/internal/xerr"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"sync"
	"time"
)

var Hostname string
var Heartbeat int
var ImDefault string

type AddConnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddConnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConnLogic {
	return &AddConnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddConnLogic) AddConn(w http.ResponseWriter, r *http.Request) error {
	jwtUser, ok := ctxdata.FromContextForJwt(l.ctx)
	if !ok {
		l.Errorf("没有找到JWT参数")
		return xerr.MsgError("没有找到JWT参数")
	}
	log.Print("jwtUser:", jwtUser.Job)
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	log.Print("ip:", ip)
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols:         nil,
		InsecureSkipVerify:   false,
		OriginPatterns:       nil,
		CompressionMode:      0,
		CompressionThreshold: 0,
	})
	if err != nil {
		logx.WithContext(r.Context()).Error("【中间件 middleware】升级请求头 error:", err)
		return xerr.MsgError("【中间件 middleware】升级请求头 error:" + err.Error())
	}
	clientLogic := NewClientLogic(l.ctx, l.svcCtx)
	clientLogic.client = &Client{
		imId:   time.Now().String(),
		conn:   c,
		isOpen: true,
		ip:     ip,
		reader: make(chan []byte, 10),
		write:  make(chan []byte, 10),
		close:  make(chan bool, 1),
		mutex:  sync.Mutex{},
	}
	clientLogic.read()
	clientLogic.write()
	Hubs.register <- clientLogic.client
	meta, err := l.svcCtx.ImRpc.Meta(l.ctx, &imrpc.ImMeta{
		DetailType: "connect",
		ImId:       clientLogic.client.imId,
		LoginId:    jwtUser.ID,
		Job:        jwtUser.Job,
		Ip:         ip,
		Issue:      "",
		Version:    "v1",
		ImServer:   Hostname,
	})
	if err != nil {
		l.Errorf("【调用 ImRpc】连接事件 error:", err)
		return err
	}
	if !meta.Ok {
		l.Errorf("【调用 ImRpc】连接事件 error:", err)
		return err
	}
	return nil
}
