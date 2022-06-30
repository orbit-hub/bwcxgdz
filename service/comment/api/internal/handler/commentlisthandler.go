package handler

import (
	"net/http"
	"strconv"

	"bwcxgdz/dousheng/service/comment/api/internal/logic"
	"bwcxgdz/dousheng/service/comment/api/internal/svc"
	"bwcxgdz/dousheng/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()
		videoId, _ := strconv.ParseInt(query.Get("video_id"), 10, 64)
		token := query.Get("token")
		req := types.CommentListRequest{
			VideoId: videoId,
			Token:   token,
		}
		l := logic.NewCommentListLogic(r.Context(), svcCtx)
		resp, err := l.CommentList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
