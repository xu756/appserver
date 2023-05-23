package logic

import (
	"context"
	"github.com/thinkeridea/go-extend/exnet"
	"github.com/xu756/appserver/api/conn/imserver/svc"
	"github.com/xu756/appserver/internal/ctxdata"
	"github.com/xu756/appserver/internal/xerr"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"sync"
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

func (l *AddConnLogic) AddConn(w http.ResponseWriter, r *http.Request, c *websocket.Conn) error {
	jwtUser, ok := ctxdata.FromContextForJwt(l.ctx)
	if !ok {
		l.Errorf("没有找到JWT参数")
		return xerr.NewErrMsg("没有找到JWT参数")
	}
	log.Print("jwtUser:", jwtUser.Job)
	ip := exnet.ClientPublicIP(r)
	if ip == "" {
		ip = exnet.ClientIP(r)
	}
	log.Print("ip:", ip)
	clientLogic := NewClientLogic(l.ctx, l.svcCtx)
	clientLogic.client = &Client{
		imId:   "11",
		conn:   c,
		isOpen: true,
		ip:     "",
		reader: make(chan []byte, 1024),
		write:  make(chan []byte, 1024),
		close:  make(chan bool),
		mutex:  sync.Mutex{},
	}
	clientLogic.read()
	clientLogic.write()
	Hubs.register <- clientLogic.client

	return nil
}
