package logic

import (
	"context"

	"github.com/xu756/appserver/api/noauth/internal/svc"
	"github.com/xu756/appserver/api/noauth/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MiniAuthLogic {
	return &MiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MiniAuthLogic) MiniAuth(req *types.MiniAuthReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
