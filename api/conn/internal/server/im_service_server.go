// Code generated by goctl. DO NOT EDIT.
// Source: conn.proto

package server

import (
	"context"

	"github.com/xu756/appserver/api/conn/internal/logic"
	"github.com/xu756/appserver/api/conn/internal/svc"
	"github.com/xu756/appserver/api/conn/pb"
)

type ImServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedImServiceServer
}

func NewImServiceServer(svcCtx *svc.ServiceContext) *ImServiceServer {
	return &ImServiceServer{
		svcCtx: svcCtx,
	}
}

// 其他服务调用发送消息
func (s *ImServiceServer) Send(ctx context.Context, in *pb.ImData) (*pb.ImResp, error) {
	l := logic.NewSendLogic(ctx, s.svcCtx)
	return l.Send(in)
}

// 其他服务调用广播消息
func (s *ImServiceServer) Broadcast(ctx context.Context, in *pb.ImData) (*pb.ImResp, error) {
	l := logic.NewBroadcastLogic(ctx, s.svcCtx)
	return l.Broadcast(in)
}