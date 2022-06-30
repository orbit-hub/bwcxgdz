package handler

import (
	"net/http"

	"bwcxgdz/dousheng/service/video/api/internal/logic"
	"bwcxgdz/dousheng/service/video/api/internal/svc"
	"bwcxgdz/dousheng/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishRequest
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewPublishLogic(r.Context(), r, svcCtx)
		resp, err := l.Publish(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
