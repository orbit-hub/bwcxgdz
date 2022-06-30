package svc

import (
	"bwcxgdz/dousheng/service/favorite/api/internal/config"
	"bwcxgdz/dousheng/service/favorite/rpc/favorite"
	"bwcxgdz/dousheng/service/video/rpc/video"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	FavoriteRpc favorite.Favorite
	VideoRpc    video.Video
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		FavoriteRpc: favorite.NewFavorite(zrpc.MustNewClient(c.FavoriteRpc)),
		VideoRpc:    video.NewVideo(zrpc.MustNewClient(c.VideoRpc)),
	}
}
