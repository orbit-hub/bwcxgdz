package handler

import (
	"bwcxgdz/dousheng/common/errno"
	"bwcxgdz/dousheng/common/response"
	"net/http"

	"bwcxgdz/dousheng/service/user/api/internal/logic"
	"bwcxgdz/dousheng/service/user/api/internal/svc"
	"bwcxgdz/dousheng/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		username := query.Get("username")
		password := query.Get("password")
		req := types.RegisterRequest{
			UserName: username,
			Password: password,
		}
		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
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
