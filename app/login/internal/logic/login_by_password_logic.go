package logic

import (
	"context"
	"github.com/xu756/appserver/app/login/internal/svc"
	"github.com/xu756/appserver/app/login/pb"
	"github.com/xu756/appserver/app/model"
	"github.com/xu756/appserver/internal/xerr"
	"github.com/xu756/appserver/internal/xjwt"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPasswordLogic {
	return &LoginByPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByPasswordLogic) LoginByPassword(in *pb.PasswordLoginReq) (*pb.LoginResp, error) {
	var resp = new(pb.LoginResp)
	user, err := l.svcCtx.UserModel.UserLogin(l.ctx, in.Username)
	switch err {
	case model.ErrNotFound:
		return resp, xerr.MsgError("用户不存在")
	case nil:
		if user.Password != in.Password {
			return resp, xerr.MsgError("密码错误")
		}
	default:
		return resp, xerr.DbErr(err)
	}
	accessToken, err := l.svcCtx.Config.Jwt.IssueToken(&xjwt.AuthInfo{
		ID:     uint64(user.Id),
		Job:    "admin",
		Issuer: "admin",
	})
	if err != nil {
		return resp, xerr.StystenError(err)
	}
	resp.AccessExpire = l.svcCtx.Config.Jwt.GetExpire()
	resp.RefreshAfter = l.svcCtx.Config.Jwt.GetMaxRefresh()
	resp.AccessToken = accessToken
	return resp, nil
}
