package handler

import (
	"net/http"
	"strconv"

	"bwcxgdz/dousheng/service/video/api/internal/logic"
	"bwcxgdz/dousheng/service/video/api/internal/svc"
	"bwcxgdz/dousheng/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()
		userId, _ := strconv.ParseInt(query.Get("user_id"), 10, 64)
		token := query.Get("token")
		req := types.PublishListRequest{
			UserId: userId,
			Token:  token,
		}
		l := logic.NewPublishListLogic(r.Context(), svcCtx)
		resp, err := l.PublishList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
