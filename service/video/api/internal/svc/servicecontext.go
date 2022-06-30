package svc

import (
	"bwcxgdz/dousheng/service/video/api/internal/config"
	"bwcxgdz/dousheng/service/video/rpc/video"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	VideoRpc video.Video
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		VideoRpc: video.NewVideo(zrpc.MustNewClient(c.VideoRpc)),
	}
}
