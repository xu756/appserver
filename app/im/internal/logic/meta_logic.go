package logic

import (
	"context"

	"github.com/xu756/appserver/app/im/internal/svc"
	"github.com/xu756/appserver/app/im/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MetaLogic {
	return &MetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Meta 元事件 连接 断开 状态更新 解密错误
func (l *MetaLogic) Meta(in *pb.ImMeta) (*pb.ImResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ImResp{}, nil
}
