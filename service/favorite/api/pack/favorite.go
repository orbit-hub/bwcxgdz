package pack

import (
	"bwcxgdz/dousheng/service/favorite/api/internal/types"
	"bwcxgdz/dousheng/service/video/rpc/proto"
)

func VideoInfo(v *proto.VideoInfoResponse) *types.VideoInfo {
	if v == nil {
		return nil
	}

	return &types.VideoInfo{
		Id:           v.Id,
		PlayUrl:      v.PlayUrl,
		CoverUrl:     v.CoverUrl,
		FavoriteCunt: v.FavoriteCount,
		CommentCount: v.CommentCount,
		IsFavorite:   true,
		Title:        v.Title,
	}
}

func VideoInfos(vs []*proto.VideoInfoResponse) []*types.VideoInfo {
	videos := make([]*types.VideoInfo, 0)
	for _, v := range vs {
		if video2 := VideoInfo(v); video2 != nil {
			videos = append(videos, video2)
		}
	}
	return videos
}
