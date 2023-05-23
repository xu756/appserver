package logic

import (
	"context"

	"github.com/xu756/appserver/api/conn/internal/svc"
	"github.com/xu756/appserver/api/conn/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BroadcastLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBroadcastLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BroadcastLogic {
	return &BroadcastLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Broadcast 其他服务调用广播消息
func (l *BroadcastLogic) Broadcast(in *pb.ImData) (*pb.ImResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ImResp{}, nil
}
