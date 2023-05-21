package logic

import (
	"context"

	"github.com/xu756/appserver/app/login/internal/svc"
	"github.com/xu756/appserver/app/login/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MiniLoginByAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMiniLoginByAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MiniLoginByAuthLogic {
	return &MiniLoginByAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MiniLoginByAuthLogic) MiniLoginByAuth(in *pb.MiniAuthReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginResp{}, nil
}
