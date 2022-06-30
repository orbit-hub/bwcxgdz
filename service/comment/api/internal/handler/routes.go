// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"bwcxgdz/dousheng/service/comment/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/douyin/comment/list",
				Handler: CommentListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/comment/action",
				Handler: CommentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}