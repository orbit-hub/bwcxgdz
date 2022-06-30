package handler

import (
	"bwcxgdz/dousheng/common/errno"
	"bwcxgdz/dousheng/common/response"
	"net/http"

	"bwcxgdz/dousheng/service/user/api/internal/logic"
	"bwcxgdz/dousheng/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
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
