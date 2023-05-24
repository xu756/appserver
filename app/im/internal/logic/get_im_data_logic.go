package logic

import (
	"context"

	"github.com/xu756/appserver/app/im/internal/svc"
	"github.com/xu756/appserver/app/im/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetImDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetImDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetImDataLogic {
	return &GetImDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetImData 服务端主动推送
func (l *GetImDataLogic) GetImData(in *pb.ImData) (*pb.ImResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ImResp{}, nil
}
