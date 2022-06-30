package logic

import (
	"bwcxgdz/dousheng/service/relation/rpc/relation"
	"context"

	"bwcxgdz/dousheng/service/relation/api/internal/svc"
	"bwcxgdz/dousheng/service/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowerListLogic) FollowerList(req *types.FollowerListRequest) (resp *types.FollowerListResponse, err error) {
	_, err = l.svcCtx.RelationRpc.GetFollowerList(l.ctx, &relation.IdRequest{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &types.FollowerListResponse{
		UserList: nil,
	}, nil
}
