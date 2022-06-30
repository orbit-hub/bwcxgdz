package logic

import (
	"context"

	"bwcxgdz/dousheng/service/relation/rpc/internal/svc"
	"bwcxgdz/dousheng/service/relation/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerListLogic) GetFollowerList(in *proto.IdRequest) (*proto.UserIdListResponse, error) {

	followerList, err := l.svcCtx.RelationModel.GetFollowerList(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &proto.UserIdListResponse{UserIdList: followerList}, nil
}
