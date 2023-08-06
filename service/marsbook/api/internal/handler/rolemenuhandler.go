package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"marsbook/service/marsbook/api/internal/logic"
	"marsbook/service/marsbook/api/internal/svc"
	"marsbook/service/marsbook/api/internal/mytypes"
)

func rolemenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req mytypes.RoleMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRolemenuLogic(r.Context(), svcCtx)
		resp, err := l.Rolemenu(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
