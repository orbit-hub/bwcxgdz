package svc

import (
	"bwcxgdz/dousheng/service/favorite/model"
	"bwcxgdz/dousheng/service/favorite/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	FavoriteModel model.FavoriteModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		FavoriteModel: model.NewFavoriteModel(conn, c.CacheRedis),
	}
}
