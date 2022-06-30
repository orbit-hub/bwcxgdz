package handler

import (
	"net/http"
	"strconv"

	"bwcxgdz/dousheng/service/favorite/api/internal/logic"
	"bwcxgdz/dousheng/service/favorite/api/internal/svc"
	"bwcxgdz/dousheng/service/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		userId, _ := strconv.ParseInt(query.Get("user_id"), 10, 64)
		req := types.FavoriteListRequest{
			UserId: userId,
		}
		l := logic.NewFavoriteListLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
