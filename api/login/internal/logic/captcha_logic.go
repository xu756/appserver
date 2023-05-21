package logic

import (
	"context"
	"github.com/xu756/appserver/app/login/loginrpc"

	"github.com/xu756/appserver/api/login/internal/svc"
	"github.com/xu756/appserver/api/login/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) Captcha() (resp *types.CaptchaResp, err error) {
	captcha, err := l.svcCtx.LoginRpc.GetCaptcha(l.ctx, &loginrpc.Empty{})
	if err != nil {
		return nil, err
	}
	return &types.CaptchaResp{
		Code:  captcha.Code,
		Key:   captcha.Key,
		Image: captcha.Image,
		Thumb: captcha.Thumb,
	}, nil
}
