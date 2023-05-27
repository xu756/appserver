package logic

import (
	"context"
	"fmt"
	"github.com/xu756/appserver/app/login/internal/svc"
	"github.com/xu756/appserver/app/login/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCaptchaLogic) GetCaptcha(in *pb.Empty) (*pb.CaptchaResp, error) {
	dots, b64, tb64, key, err := l.svcCtx.Captcha.Generate()
	if err != nil {
		panic(err)
		return nil, err
	}
	err = l.svcCtx.RedisClient.SetexCtx(l.ctx, key, fmt.Sprintf("%+v", dots[0]), 60*5)
	if err != nil {
		return nil, err
	}
	logx.Error("【Redis-error】", err)
	return &pb.CaptchaResp{
		Code:  0,
		Key:   key,
		Image: b64,
		Thumb: tb64,
	}, nil
}
