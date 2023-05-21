package logic

import (
	"context"
	"fmt"
	"github.com/xu756/appserver/app/login/internal/svc"
	"github.com/xu756/appserver/app/login/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaCompareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCaptchaCompareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaCompareLogic {
	return &CaptchaCompareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CaptchaCompareLogic) CaptchaCompare(in *pb.CaptchaCheckReq) (*pb.CaptchaCheckResp, error) {
	//captcha.CheckPointDistWithPadding()
	fmt.Print("验证码比对逻辑", in.Dots)
	return &pb.CaptchaCheckResp{}, nil
}
