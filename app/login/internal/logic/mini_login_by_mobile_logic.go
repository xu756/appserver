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
	var useId int64
	if in.SmsCode != "666666" {
		return nil, xerr.MsgError("验证码错误")
	}
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	switch err {
	case nil:
		useId = user.Id
	case model.ErrNotFound:
		return NewUserRegisterLogic(l.ctx, l.svcCtx).UserRegister(&pb.Register{
			Mobile: in.Mobile,
		})
	default:
		return res, xerr.DbErr(err)

	}
	accessToken, err := l.svcCtx.Config.Jwt.IssueToken(&xjwt.AuthInfo{
		ID:     uint64(useId),
		Job:    "user",
		Issuer: "mini",
	})
	if err != nil {
		return res, xerr.StystenError(err)
	}
	res.AccessExpire = l.svcCtx.Config.Jwt.GetExpire()
	res.RefreshAfter = l.svcCtx.Config.Jwt.GetMaxRefresh()
	res.AccessToken = accessToken
	return res, nil
}
