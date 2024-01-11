package hander

import (
	"context"
	"server/kitex_gen/user"
)

func (u UserImpl) LoginByPassword(ctx context.Context, req *user.LoginByPasswordReq) (res *user.LoginRes, err error) {
	//TODO implement me
	panic("implement me")
}

func (u UserImpl) LoginByMobile(ctx context.Context, req *user.LoginByMobileReq) (res *user.LoginRes, err error) {
	//TODO implement me
	panic("implement me")
}
