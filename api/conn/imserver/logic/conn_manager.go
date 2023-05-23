package logic

import (
	"context"
	"github.com/xu756/appserver/api/conn/imserver/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

var Hostname string
var Heartbeat int
var ImDefault string

type AddConnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddConnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddConnLogic {
	return &AddConnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddConnLogic) AddConn(w http.ResponseWriter, r *http.Request) error {
	//jwtUser, ok := ctxdata.FromContextForJwt(l.ctx)
	//if !ok {
	//	l.Errorf("没有找到JWT参数")
	//	return xerr.NewErrMsg("没有找到JWT参数")
	//}

	//conn, err := l.svcCtx.Upgrader.Upgrade(w, r, nil)
	//if err != nil {
	//	l.Error(err)
	//	return err
	//}
	//
	//ip := exnet.ClientPublicIP(r)
	//if ip == "" {
	//	ip = exnet.ClientIP(r)
	//}
	//wsid := fmt.Sprintf("%s_%s", jwtUser.GetJob(), jwtUser.GetStringID())
	return nil
}
