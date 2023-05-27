package logic

import (
	"context"

	"github.com/xu756/appserver/app/login/internal/svc"
	"github.com/xu756/appserver/app/login/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PasswordLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPasswordLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PasswordLoginLogic {
	return &PasswordLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PasswordLoginLogic) PasswordLogin(in *pb.PasswordLoginReq) (*pb.LoginResp, error) {
	var resp = new(pb.LoginResp)

	return resp, nil
}
