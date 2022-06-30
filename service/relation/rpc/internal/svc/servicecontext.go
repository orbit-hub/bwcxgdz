package svc

import (
	"bwcxgdz/dousheng/service/relation/model"
	"bwcxgdz/dousheng/service/relation/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	RelationModel model.RelationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		RelationModel: model.NewRelationModel(conn, c.CacheRedis),
	}
}
