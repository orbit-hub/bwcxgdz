package svc

import (
	"bwcxgdz/dousheng/service/video/model"
	"bwcxgdz/dousheng/service/video/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	VideoModel model.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideoModel(conn, c.CacheRedis),
	}
}
