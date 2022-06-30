package svc

import (
	"bwcxgdz/dousheng/service/comment/model"
	"bwcxgdz/dousheng/service/comment/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(conn, c.CacheRedis),
	}
}
