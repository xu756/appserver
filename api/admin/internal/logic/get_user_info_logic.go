package logic

import (
	"context"
	"github.com/xu756/appserver/api/admin/internal/svc"
	"github.com/xu756/appserver/api/admin/internal/types"
	"github.com/xu756/appserver/app/admin/pb"
	"github.com/xu756/appserver/internal/xjwt"
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

	user := xjwt.AuthToInfo(l.ctx.Value("user"))
	info, err := l.svcCtx.AdminRpc.GetUserInfo(l.ctx, &pb.GetUserInfoReq{Id: int64(user.ID)})
	if err != nil {
		return nil, err
	}
	resp = &types.UserInfo{
		Id:     info.Id,
		Name:   info.Name,
		Avatar: info.Avatar,
		Role:   info.Role,
		OpenId: info.OpenId,
	}
	return resp, nil
}
