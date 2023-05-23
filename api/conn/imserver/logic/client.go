package logic

import (
	"context"
	"github.com/xu756/appserver/api/conn/imserver/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"nhooyr.io/websocket"
	"sync"
)

type Client struct {
	mutex  sync.Mutex
	imId   string
	conn   *websocket.Conn
	isOpen bool
	ip     string
	reader chan []byte
	write  chan []byte
	close  chan bool
}

type ClientLogic struct {
	logx.Logger
	ctx    context.Context
	client *Client
	svcCtx *svc.ServiceContext
}

func NewClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientLogic {
	return &ClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClientLogic) read() {
	go func() {
		for {
			select {
			case <-l.client.close:
				return
			default:
				_, message, err := l.client.conn.Read(context.TODO())
				if err != nil {
					return
				}
				l.client.write <- message
			}
		}
	}()

}

func (l *ClientLogic) write() {
	go func() {
		for {
			select {
			case <-l.client.close:
				return
			case message := <-l.client.write:
				err := l.client.conn.Write(context.Background(), websocket.MessageBinary, message)
				if err != nil {
					return
				}
			}
		}
	}()
}
