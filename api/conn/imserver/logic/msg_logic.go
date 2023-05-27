package logic

import (
	"github.com/xu756/appserver/immsg"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *ClientLogic) msgLogic() {
	msg, err := immsg.ToMsg(<-l.client.reader)
	if err != nil {
		logx.Error("【解析消息失败】", err)
		l.closeClient()
	}
	Hubs.Broadcast <- msg.ToBytes()
}
