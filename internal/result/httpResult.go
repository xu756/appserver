package result

import (
	"github.com/pkg/errors"
	"github.com/xu756/appserver/internal/xerr"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// HttpResult http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		// 错误返回
		errCode := xerr.ServerCommonError
		errMsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			// 自定义CodeError
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		}
		httpx.WriteJson(w, 501, Error(errCode, errMsg))
	}
}
