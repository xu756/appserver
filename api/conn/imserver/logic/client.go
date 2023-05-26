package logic

import (
	"context"
	"github.com/xu756/appserver/api/conn/imserver/svc"
	"github.com/xu756/appserver/app/im/pb"
	"github.com/zeromicro/go-zero/core/logx"
	"nhooyr.io/websocket"
	"sync"
)

type Client struct {
	mutex     sync.Mutex
	imId      string
	conn      *websocket.Conn
	isOpen    bool
	ip        string
	reader    chan []byte
	write     chan []byte
	close     chan bool
	closeOnce sync.Once
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
				_, message, err := l.client.conn.Read(context.Background())
				if err != nil {
					l.closeClient()
					return
				}
				message, err = aESDecrypt(message)
				if err != nil {
					logx.Error("【解密失败】", err)
					l.closeClient()
					return
				}
				Hubs.Broadcast <- message
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
				err := l.client.conn.Write(context.Background(), websocket.MessageBinary, aESEncrypt(message))
				if err != nil {
					l.closeClient()
					return
				}
			}
		}
	}()
}

func (l *ClientLogic) closeClient() {
	l.client.closeOnce.Do(func() {
		l.client.mutex.Lock()
		defer l.client.mutex.Unlock()
		if l.client.isOpen {
			l.client.isOpen = false
			Hubs.unregister <- l.client
			l.svcCtx.ImRpc.Meta(context.Background(), &pb.ImMeta{
				DetailType: "disconnect",
				ImId:       l.client.imId,
				LoginId:    0,
				Job:        "",
				Ip:         "",
				Issue:      "",
				Version:    "",
				ImServer:   "",
				Data:       nil,
			})
			l.client.conn.Close(websocket.StatusNormalClosure, "Connection closed")

		}
	})
}
