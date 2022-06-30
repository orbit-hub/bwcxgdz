package handler

import (
	"net/http"
	"strconv"

	"bwcxgdz/dousheng/service/relation/api/internal/logic"
	"bwcxgdz/dousheng/service/relation/api/internal/svc"
	"bwcxgdz/dousheng/service/relation/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FollowListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		userId, _ := strconv.ParseInt(query.Get("user_id"), 10, 64)
		token := query.Get("token")
		req := types.FollowListRequest{
			UserId: userId,
			Token:  token,
		}
		l := logic.NewFollowListLogic(r.Context(), svcCtx)
		resp, err := l.FollowList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
