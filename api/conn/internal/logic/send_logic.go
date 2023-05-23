package logic

import (
	"context"

	"github.com/xu756/appserver/api/conn/internal/svc"
	"github.com/xu756/appserver/api/conn/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Send 其他服务调用发送消息
func (l *SendLogic) Send(in *pb.ImData) (*pb.ImResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ImResp{}, nil
}
