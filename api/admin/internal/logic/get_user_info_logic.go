package logic

import (
	"context"

	"github.com/xu756/appserver/api/admin/internal/svc"
	"github.com/xu756/appserver/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.EmptyReq) (resp *types.UserInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
