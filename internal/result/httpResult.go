package result

import (
	"github.com/pkg/errors"
	"github.com/xu756/appserver/internal/xerr"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

// HttpResult http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	logx.Error("【API-错误】 : %+v ", err)
	if err == nil {
		// 成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		// 错误返回
		errCode := uint32(http.StatusInternalServerError)
		errMsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			// 自定义CodeError
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gStatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gStatus.Code())
				if xerr.HasErrCode(grpcCode) { // 区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errCode = grpcCode
					errMsg = gStatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)

		switch errCode {
		case xerr.DbError:
			httpx.WriteJson(w, http.StatusInternalServerError, Error(http.StatusInternalServerError, errMsg))
			return

		case http.StatusUnauthorized:
			httpx.WriteJson(w, http.StatusUnauthorized, Error(errCode, errMsg))
			return

		case http.StatusNotFound:
			httpx.WriteJson(w, http.StatusNotFound, Error(errCode, errMsg))
			return

		default:
			httpx.WriteJson(w, http.StatusInternalServerError, Error(errCode, errMsg))
			return
		}
	}
}
