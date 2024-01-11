package result

import (
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"net/http"
	"server/internal/xerr"
)

func HttpSuccess(c *app.RequestContext, resp interface{}) {
	r := successRes(resp)
	c.JSON(http.StatusOK, r)
	c.Abort()
}

func HttpError(c *app.RequestContext, err error) {
	log.Print(err)
	var resErr xerr.CodeError
	errors.As(err, &resErr)
	c.JSON(http.StatusOK, errorRes(resErr.Code, resErr.Msg))
	c.Abort()
}

func HttpErrorCode(c *app.RequestContext, code int32) {
	c.JSON(http.StatusOK, errorRes(code, xerr.GetMsg(code)))
	c.Abort()
}
