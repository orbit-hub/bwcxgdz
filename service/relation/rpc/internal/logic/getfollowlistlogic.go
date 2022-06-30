package logic

import (
	"context"

	"bwcxgdz/dousheng/service/relation/rpc/internal/svc"
	"bwcxgdz/dousheng/service/relation/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListLogic {
	return &GetFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowListLogic) GetFollowList(in *proto.IdRequest) (*proto.UserIdListResponse, error) {
	list, err := l.svcCtx.RelationModel.GetFollowList(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &proto.UserIdListResponse{
		UserIdList: list,
	}, nil
}
