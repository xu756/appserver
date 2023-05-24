package logic

import (
	"context"
	"github.com/xu756/appserver/api/login/internal/svc"
	"github.com/xu756/appserver/api/login/internal/types"
	"github.com/xu756/appserver/app/login/loginrpc"
	"github.com/xu756/appserver/internal/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByMobileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByMobileLogic {
	return &LoginByMobileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *LoginByMobileLogic) LoginByMobile(req *types.LoginByMobileReq) (resp *types.LoginResp, err error) {
	res, err := l.svcCtx.LoginRpc.MiniLoginByMobile(l.ctx, &loginrpc.LoginReq{
		Mobile:  req.Mobile,
		SmsCode: req.SmsCode,
	})
	if err != nil {
		return nil, xerr.StystenError(err)
	}
	return &types.LoginResp{
		AccessToken:  res.AccessToken,
		AccessExpire: res.AccessExpire,
		RefreshAfter: res.RefreshAfter,
	}, err
}
