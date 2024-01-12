package result

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"net/http"
	"server/internal/xerr"
)

func HttpSuccess(c *app.RequestContext, resp interface{}) {
	r := successRes(resp)
	c.JSON(http.StatusOK, r)
	c.Abort()
}

func HttpError(c *app.RequestContext, err error) {
	var resErr xerr.CodeError
	errors.As(err, &resErr)
	c.JSON(http.StatusOK, errorRes(resErr.Code, resErr.Msg))
	c.Abort()
}

func HttpBizErr(c *app.RequestContext, err error) {
	if bizErr, isBizErr := kerrors.FromBizStatusError(err); isBizErr {
		c.JSON(http.StatusOK, errorRes(bizErr.BizStatusCode(), bizErr.BizMessage()))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, errorRes(xerr.SystemErrCode, xerr.GetMsg(xerr.SystemErrCode)))
	c.Abort()

}

func HttpParamErr(c *app.RequestContext) {
	c.JSON(http.StatusOK, errorRes(xerr.Param, xerr.GetMsg(xerr.Param)))
	c.Abort()
}
