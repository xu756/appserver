package rpc

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"log"
	"server/common/config"
	"server/internal/xerr"
	"server/kitex_gen/user/usersrv"
)

var UserClient usersrv.Client

func InitUserClient(destService string) {
	s, err := usersrv.NewClient(destService,
		client.WithHostPorts(config.RunData.Rpc.UserRpcAddr),
		client.WithErrorHandler(ClientErrorHandler),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)
	if err != nil {
		hlog.Debug("【user】client init error", err)
		panic(err)
	}
	UserClient = s
}

func ClientErrorHandler(ctx context.Context, err error) error {
	// todo 错误统一处理
	//bizErr, isBizErr := kerrors.FromBizStatusError(err)
	if bizErr, isBizErr := kerrors.FromBizStatusError(err); isBizErr {
		log.Print(bizErr.Error())
		return xerr.NewErr(bizErr.BizStatusCode(), bizErr.BizMessage())
	}
	return err
}
