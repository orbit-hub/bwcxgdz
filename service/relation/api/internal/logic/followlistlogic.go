package logic

import (
	"bwcxgdz/dousheng/service/relation/rpc/relation"
	"context"

	"bwcxgdz/dousheng/service/relation/api/internal/svc"
	"bwcxgdz/dousheng/service/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowListLogic) FollowList(req *types.FollowListRequest) (resp *types.FollowListResponse, err error) {
	_, err = l.svcCtx.RelationRpc.GetFollowList(l.ctx, &relation.IdRequest{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	//for _, id := range ids.UserIdList {
	//
	//}
	return &types.FollowListResponse{
		UserList: nil,
	}, nil
}
