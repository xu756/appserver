package logic

import (
	"context"
	"github.com/xu756/appserver/app/login/loginrpc"

	"github.com/xu756/appserver/api/login/internal/svc"
	"github.com/xu756/appserver/api/login/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaCompareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaCompareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaCompareLogic {
	return &CaptchaCompareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaCompareLogic) CaptchaCompare(req *types.CaptchaReq) (resp *types.CaptchaCompareResp, err error) {
	res, err := l.svcCtx.LoginRpc.CaptchaCompare(l.ctx, &loginrpc.CaptchaCheckReq{
		Dots: req.Dots,
		Key:  req.Key,
	})
	if err != nil {
		return nil, err
	}
	return &types.CaptchaCompareResp{
		Ok: res.Result,
	}, nil
}
