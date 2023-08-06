package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"marsbook/service/marsbook/api/internal/logic"
	"marsbook/service/marsbook/api/internal/svc"
	"marsbook/service/marsbook/api/internal/mytypes"
)

func userinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req mytypes.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserinfoLogic(r.Context(), svcCtx)
		resp, err := l.Userinfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
