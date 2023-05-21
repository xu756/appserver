package handler

import (
	"github.com/xu756/appserver/internal/result"
	"net/http"

	"github.com/xu756/appserver/api/mini/internal/logic"
	"github.com/xu756/appserver/api/mini/internal/svc"
	"github.com/xu756/appserver/api/mini/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MiniAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MiniAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMiniAuthLogic(r.Context(), svcCtx)
		resp, err := l.MiniAuth(&req)
		result.HttpResult(r, w, resp, err)
	}
}
