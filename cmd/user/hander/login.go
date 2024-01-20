package hander

import (
	"context"
	"server/common/config"
	"server/internal/xerr"
	"server/kitex_gen/user"
)

func (u UserImpl) LoginByPassword(ctx context.Context, req *user.LoginByPasswordReq) (res *user.LoginRes, err error) {
	userInfo, err := u.Model.FindUserByUsername(ctx, req.Username)
	//log.Print(kerrors.FromBizStatusError(err))
	if err != nil {
		return nil, err
	}
	if userInfo.Password != req.Password {
		return nil, xerr.ErrMsg(xerr.UserPasswordErr)
	}
	res.Token = "111"

	res.Expire = config.RunData.JwtConfig.Expire
	return res, nil
}

func (u UserImpl) LoginByMobile(ctx context.Context, req *user.LoginByMobileReq) (res *user.LoginRes, err error) {
	userInfo, err := u.Model.FindUserByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	if req.Captcha == "1234" {
		return nil, xerr.ErrMsg(xerr.CaptchaNotExist)
	}
	res.Token = userInfo.UUID.String()
	res.Expire = config.RunData.JwtConfig.Expire
	return res, nil

}
