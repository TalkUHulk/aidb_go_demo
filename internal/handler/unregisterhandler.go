package handler

import (
	"net/http"

	"aidb_go/internal/logic"
	"aidb_go/internal/svc"
	"aidb_go/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func unregisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UnRigsterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUnregisterLogic(r.Context(), svcCtx)
		resp, err := l.Unregister(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
