package logic

import (
	"context"
	"github.com/xu756/appserver/app/admin/internal/svc"
	"github.com/xu756/appserver/app/admin/pb"
	"github.com/xu756/appserver/app/model"
	"github.com/xu756/appserver/internal/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.UserInfo, error) {
	var resp = new(pb.UserInfo)
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
		resp.Id = user.Id
		resp.Name = user.Username
		resp.Avatar = user.Avatar
		resp.Role = user.Role
		resp.OpenId = user.OpenId
	case model.ErrNotFound:
		return resp, model.ErrNotFound
	default:
		return resp, xerr.DbErr(err)

	}
	return resp, nil
}
