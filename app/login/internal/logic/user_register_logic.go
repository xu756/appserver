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

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserRegister 注册
func (l *UserRegisterLogic) UserRegister(in *pb.Register) (*pb.LoginResp, error) {
	var res = new(pb.LoginResp)

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
		return res, xerr.DbErr(err)
	}
	useId, err := resp.LastInsertId()
	if err != nil {
		return res, xerr.DbErr(err)
	}
	accessToken, err := l.svcCtx.Config.Jwt.IssueToken(&xjwt.AuthInfo{
		ID:     uint64(useId),
		Job:    "user",
		Issuer: "register",
	})
	if err != nil {
		return res, xerr.StystenError(err)
	}
	res.AccessExpire = l.svcCtx.Config.Jwt.GetExpire()
	res.RefreshAfter = l.svcCtx.Config.Jwt.GetMaxRefresh()
	res.AccessToken = accessToken
	return res, nil
}
