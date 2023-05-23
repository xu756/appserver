package logic

import (
	"context"
	"github.com/xu756/appserver/internal/xjwt"

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
	accessToken, err := l.svcCtx.Config.Jwt.IssueToken(&xjwt.AuthInfo{
		ID:  uint64(123),
		Job: "mini_user",
	})
	if err != nil {
		logx.Error("【 login-rpc】生成jwt", err)
		return nil, err
	}
	res.AccessExpire = l.svcCtx.Config.Jwt.GetExpire()
	res.RefreshAfter = l.svcCtx.Config.Jwt.GetMaxRefresh()
	res.AccessToken = accessToken
	return res, nil
}
