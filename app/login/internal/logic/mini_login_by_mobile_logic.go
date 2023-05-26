package logic

import (
	"context"
	"github.com/rs/xid"
	"github.com/xu756/appserver/app/model"
	"github.com/xu756/appserver/internal/xerr"
	"github.com/xu756/appserver/internal/xjwt"
	"time"

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
	var useId int64
	if in.SmsCode != "666666" {
		return nil, xerr.MsgError("验证码错误")
	}
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	switch err {
	case nil:
		useId = user.Id
	case model.ErrNotFound:
		uid := xid.New()
		resp, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
			CreatedTime: time.Now().Unix(),
			UpdatedTime: time.Now().Unix(),
			Mobile:      in.Mobile,
			Username:    "user_" + in.Mobile,
			Password:    "123456",
			OpenId:      "open_" + uid.String(),
			Avatar:      "https://file.xu756.top/avatar.png",
		})
		if err != nil {
			return nil, xerr.DbErr(err)
		}
		useId, err = resp.LastInsertId()
		if err != nil {
			return nil, xerr.DbErr(err)
		}
	default:
		return nil, xerr.DbErr(err)

	}
	accessToken, err := l.svcCtx.Config.Jwt.IssueToken(&xjwt.AuthInfo{
		ID:     uint64(useId),
		Job:    "user",
		Issuer: "mini",
	})
	if err != nil {
		return nil, xerr.StystenError(err)
	}
	res.AccessExpire = l.svcCtx.Config.Jwt.GetExpire()
	res.RefreshAfter = l.svcCtx.Config.Jwt.GetMaxRefresh()
	res.AccessToken = accessToken
	return res, nil
}
