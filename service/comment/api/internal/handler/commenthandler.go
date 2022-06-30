package handler

import (
	"net/http"
	"strconv"

	"bwcxgdz/dousheng/service/comment/api/internal/logic"
	"bwcxgdz/dousheng/service/comment/api/internal/svc"
	"bwcxgdz/dousheng/service/comment/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()
		videoId, _ := strconv.ParseInt(query.Get("video_id"), 10, 64)
		token := query.Get("token")
		actionType, _ := strconv.ParseInt(query.Get("action_type"), 10, 32)
		commentText := query.Get("comment_text")
		commentId, _ := strconv.ParseInt(query.Get("comment_id"), 10, 64)
		req := types.CommentRequest{
			VideoId:     videoId,
			Token:       token,
			ActionType:  int32(actionType),
			CommentText: commentText,
			CommentId:   commentId,
		}

		l := logic.NewCommentLogic(r.Context(), svcCtx)
		resp, err := l.Comment(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
