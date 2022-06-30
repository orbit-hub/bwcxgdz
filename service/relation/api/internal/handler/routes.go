// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"bwcxgdz/dousheng/service/relation/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/douyin/relation/follow/list",
				Handler: FollowListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/relation/follower/list",
				Handler: FollowerListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/relation/action",
				Handler: RelationHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}