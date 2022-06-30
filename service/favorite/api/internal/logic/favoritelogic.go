package logic

import (
	"bwcxgdz/dousheng/service/favorite/rpc/favorite"
	"context"
	"encoding/json"

	"bwcxgdz/dousheng/service/favorite/api/internal/svc"
	"bwcxgdz/dousheng/service/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteLogic) Favorite(req *types.FavoriteRequest) (resp *types.FavoriteResponse, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	_, err = l.svcCtx.FavoriteRpc.FavoriteAction(l.ctx, &favorite.ActionRequest{
		Id:         uid,
		VideoID:    req.VideoId,
		ActionType: req.ActionType,
	})
	if err != nil {
		return nil, err
	}
	return &types.FavoriteResponse{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}
