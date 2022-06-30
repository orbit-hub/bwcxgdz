package handler

import (
	"bwcxgdz/dousheng/common/errno"
	"bwcxgdz/dousheng/common/response"
	"net/http"
	"strconv"

	"bwcxgdz/dousheng/service/video/api/internal/logic"
	"bwcxgdz/dousheng/service/video/api/internal/svc"
	"bwcxgdz/dousheng/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func feedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		latestTime, _ := strconv.ParseInt(query.Get("latest_time"), 10, 64)
		token := query.Get("token")
		req := types.FeedRequest{
			LatestTime: latestTime,
			Token:      token,
		}
		l := logic.NewFeedLogic(r.Context(), svcCtx)
		resp, err := l.Feed(&req)
		if err != nil {
			Err := errno.ConvertErr(err)
			httpx.WriteJson(w, http.StatusInternalServerError, response.ErrResponse{
				StatusCode: Err.ErrCode,
				StatusMsg:  Err.ErrMsg,
			})
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
