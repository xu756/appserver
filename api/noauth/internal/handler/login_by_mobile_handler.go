package handler

import (
	"github.com/xu756/appserver/internal/result"
	"net/http"

	"github.com/xu756/appserver/api/noauth/internal/logic"
	"github.com/xu756/appserver/api/noauth/internal/svc"
	"github.com/xu756/appserver/api/noauth/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginByMobileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginByMobileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginByMobileLogic(r.Context(), svcCtx)
		resp, err := l.LoginByMobile(&req)
		result.HttpResult(r, w, resp, err)
	}
}
