package handler

import (
	"net/http"
	"strconv"

	"bwcxgdz/dousheng/service/relation/api/internal/logic"
	"bwcxgdz/dousheng/service/relation/api/internal/svc"
	"bwcxgdz/dousheng/service/relation/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RelationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		userId, _ := strconv.ParseInt(query.Get("to_user_id"), 10, 64)
		token := query.Get("token")
		actionType, _ := strconv.ParseInt(query.Get("action_type"), 10, 32)
		req := types.RelationRequest{
			ToUserId:   userId,
			Token:      token,
			ActionType: int32(actionType),
		}
		l := logic.NewRelationLogic(r.Context(), svcCtx)
		resp, err := l.Relation(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
