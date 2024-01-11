package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"log"
	"server/cmd/admin/rpc"
	"server/internal/result"
	"server/internal/xerr"
	"server/kitex_gen/user"
)

func LoginRouter(r *route.RouterGroup) {
	r.POST("/account", loginByPassword)
	r.POST("/mobile", loginByMobile)
	r.POST("/captcha", sendCaptcha)

}

func loginByPassword(ctx context.Context, c *app.RequestContext) {
	var req LoginByPasswordReq
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpErrorCode(c, xerr.Param)
		return
	}
	res, err := rpc.UserClient.LoginByPassword(ctx, &user.LoginByPasswordReq{
		Username:  req.Username,
		Password:  req.Password,
		SessionId: req.SessionId,
	})
	if err != nil {
		log.Print(err)
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res)
}

func loginByMobile(ctx context.Context, c *app.RequestContext) {
	var req LoginByMobileReq
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpError(c, xerr.ParamErr())
		return
	}
	res, err := rpc.UserClient.LoginByMobile(ctx, &user.LoginByMobileReq{
		Mobile:    req.Mobile,
		Captcha:   req.Captcha,
		SessionId: req.SessionId,
	})
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res)

}

func sendCaptcha(ctx context.Context, c *app.RequestContext) {
	var req SendCaptchaReq
	if err := c.BindAndValidate(&req); err != nil {
		result.HttpError(c, xerr.ParamErr())
		return
	}
	res, err := rpc.UserClient.SendCaptcha(ctx, &user.SendCaptchaReq{
		Mobile:    req.Mobile,
		SessionId: req.SessionId,
	})
	if err != nil {
		result.HttpError(c, err)
		return
	}
	result.HttpSuccess(c, res.Success)
}
