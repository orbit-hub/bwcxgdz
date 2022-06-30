package logic

import (
	"bwcxgdz/dousheng/service/relation/rpc/relation"
	"context"
	"encoding/json"

	"bwcxgdz/dousheng/service/relation/api/internal/svc"
	"bwcxgdz/dousheng/service/relation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationLogic {
	return &RelationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationLogic) Relation(req *types.RelationRequest) (resp *types.RelationResponse, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	_, err = l.svcCtx.RelationRpc.RelationAction(l.ctx, &relation.RelationActionRequest{
		Id:         uid,
		FollowId:   req.ToUserId,
		ActionType: req.ActionType,
	})
	if err != nil {
		return nil, err
	}
	return &types.RelationResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}
