package svc

import (
	"bwcxgdz/dousheng/service/relation/api/internal/config"
	"bwcxgdz/dousheng/service/relation/rpc/relation"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RelationRpc relation.Relation
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RelationRpc: relation.NewRelation(zrpc.MustNewClient(c.RelationRpc)),
	}
}
