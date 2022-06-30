package pack

import (
	"bwcxgdz/dousheng/service/video/api/internal/types"
	"bwcxgdz/dousheng/service/video/rpc/video"
)

func videoInfo(v *video.VideoInfoResponse) *types.VideoInfo {
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

func VideoInfos(vs []*video.VideoInfoResponse) []*types.VideoInfo {
	videos := make([]*types.VideoInfo, 0)
	for _, v := range vs {
		if video2 := videoInfo(v); video2 != nil {
			videos = append(videos, video2)
		}
	}
	return videos
}
