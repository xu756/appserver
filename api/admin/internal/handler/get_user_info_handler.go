package handler

import (
	"github.com/xu756/appserver/internal/result"
	"net/http"

	"github.com/xu756/appserver/api/admin/internal/logic"
	"github.com/xu756/appserver/api/admin/internal/svc"
	"github.com/xu756/appserver/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmptyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
