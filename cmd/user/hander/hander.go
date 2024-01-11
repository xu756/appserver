package hander

import (
	"context"
	"server/common/db"
	"server/internal/xjwt"
	"server/kitex_gen/user"
)

type UserImpl struct {
	Model db.Model
	Jwt   *xjwt.Jwt
}

func (u UserImpl) SendCaptcha(ctx context.Context, req *user.SendCaptchaReq) (res *user.SendCaptchaRes, err error) {
	//TODO implement me
	panic("implement me")
}

func NewUserImpl() *UserImpl {
	return &UserImpl{
		Model: db.NewModel(),
		Jwt:   xjwt.NewJwt(),
	}
}
