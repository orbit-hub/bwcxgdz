package logic

import (
	"bwcxgdz/dousheng/service/favorite/model"
	"bwcxgdz/dousheng/service/favorite/rpc/internal/svc"
	"bwcxgdz/dousheng/service/favorite/rpc/proto"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteActionLogic) FavoriteAction(in *proto.ActionRequest) (*proto.ActionResponse, error) {
	favorite := model.Favorite{
		Id:        in.Id,
		VideoId:   in.VideoID,
		CreatedAt: time.Now(),
	}
	if in.ActionType == 1 {
		_, err := l.svcCtx.FavoriteModel.Insert(l.ctx, &favorite)
		if err != nil {
			return nil, err
		}
	} else {
		err := l.svcCtx.FavoriteModel.DeleteByUserIdAndVideoId(l.ctx, favorite.Id, favorite.VideoId)
		if err != nil {
			return nil, err
		}
	}
	return &proto.ActionResponse{
		Success: true,
	}, nil
}
