package logic

import (
	"context"

	"github.com/xu756/appserver/app/login/internal/svc"
	"github.com/xu756/appserver/app/login/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MiniLoginByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMiniLoginByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MiniLoginByMobileLogic {
	return &MiniLoginByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MiniLoginByMobileLogic) MiniLoginByMobile(in *pb.LoginReq) (*pb.LoginResp, error) {
	var res = new(pb.LoginResp)
	res.AccessExpire = 100
	res.RefreshAfter = 100
	res.AccessToken = "123456"
	return res, nil
}
