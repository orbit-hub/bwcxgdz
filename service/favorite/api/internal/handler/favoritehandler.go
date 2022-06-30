package handler

import (
	"net/http"
	"strconv"

	"bwcxgdz/dousheng/service/favorite/api/internal/logic"
	"bwcxgdz/dousheng/service/favorite/api/internal/svc"
	"bwcxgdz/dousheng/service/favorite/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		videoId, _ := strconv.ParseInt(query.Get("video_id"), 10, 64)

		actionType, _ := strconv.Atoi(query.Get("action_type"))

		req := types.FavoriteRequest{
			VideoId:    videoId,
			ActionType: int32(actionType),
		}
		l := logic.NewFavoriteLogic(r.Context(), svcCtx)
		resp, err := l.Favorite(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
