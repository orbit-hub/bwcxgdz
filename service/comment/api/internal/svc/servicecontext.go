package svc

import (
	"bwcxgdz/dousheng/service/comment/api/internal/config"
	"bwcxgdz/dousheng/service/comment/rpc/comment"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	CommentRpc comment.Comment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CommentRpc: comment.NewComment(zrpc.MustNewClient(c.CommentRpc)),
	}
}
