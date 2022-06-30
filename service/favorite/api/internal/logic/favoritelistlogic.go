package logic

import (
	"bwcxgdz/dousheng/service/favorite/api/pack"
	"bwcxgdz/dousheng/service/favorite/rpc/favorite"
	"bwcxgdz/dousheng/service/video/rpc/video"
	"context"

	"bwcxgdz/dousheng/service/favorite/api/internal/svc"
	"bwcxgdz/dousheng/service/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteListRequest) (resp *types.VideoListResponse, err error) {
	list, err := l.svcCtx.FavoriteRpc.GetFavoriteVideoListById(l.ctx, &favorite.IdRequest{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	infos := make([]*types.VideoInfo, 0)
	for _, id := range list.VideoIdList {
		byId, err := l.svcCtx.VideoRpc.GetVideoInfoById(l.ctx, &video.IdRequest{
			Id: id,
		})
		if err != nil {
			return nil, err
		}
		videoInfo := pack.VideoInfo(byId)
		infos = append(infos, videoInfo)
	}

	return &types.VideoListResponse{
		StatusCode:    0,
		StatusMsg:     "success",
		VideoInfoList: infos,
	}, nil
}
