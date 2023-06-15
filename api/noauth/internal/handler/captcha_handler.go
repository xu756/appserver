package handler

import (
	"github.com/xu756/appserver/internal/result"
	"net/http"

	"github.com/xu756/appserver/api/noauth/internal/logic"
	"github.com/xu756/appserver/api/noauth/internal/svc"
)

func CaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.Captcha()
		result.HttpResult(r, w, resp, err)
	}
}
