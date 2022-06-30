package logic

import (
	"bwcxgdz/dousheng/service/relation/model"
	"context"
	"time"

	"bwcxgdz/dousheng/service/relation/rpc/internal/svc"
	"bwcxgdz/dousheng/service/relation/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RelationActionLogic) RelationAction(in *proto.RelationActionRequest) (*proto.ActionResponse, error) {
	relation := model.Relation{
		UserId:    in.Id,
		CreatedAt: time.Now(),
		Follow:    in.FollowId,
	}
	if in.ActionType == 1 {
		_, err := l.svcCtx.RelationModel.Insert(l.ctx, &relation)
		if err != nil {
			return nil, err
		}
	} else {
		err := l.svcCtx.RelationModel.DeleteRelation(l.ctx, relation.UserId, relation.Follow)
		if err != nil {
			return nil, err
		}
	}

	return &proto.ActionResponse{
		Success: true,
	}, nil
}
