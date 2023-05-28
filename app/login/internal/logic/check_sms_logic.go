package logic

import (
	"context"

	"github.com/xu756/appserver/app/login/internal/svc"
	"github.com/xu756/appserver/app/login/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckSmsLogic {
	return &CheckSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证短信验证码
func (l *CheckSmsLogic) CheckSms(in *pb.CheckSmsReq) (*pb.CaptchaCheckResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CaptchaCheckResp{}, nil
}
