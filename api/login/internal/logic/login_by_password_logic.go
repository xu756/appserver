package logic

import (
	"context"
	"github.com/xu756/appserver/api/login/internal/svc"
	"github.com/xu756/appserver/api/login/internal/types"
	"github.com/xu756/appserver/app/login/loginrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPasswordLogic {
	return &LoginByPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByPasswordLogic) LoginByPassword(req *types.LoginByPasswordReq) (resp *types.LoginResp, err error) {
	res, err := l.svcCtx.LoginRpc.LoginByPassword(l.ctx, &loginrpc.PasswordLoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		AccessToken:  res.AccessToken,
		AccessExpire: res.AccessExpire,
		RefreshAfter: res.RefreshAfter,
	}, err
}
