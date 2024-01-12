package xerr

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

/**
常用通用固定错误
*/

type CodeError struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func (e CodeError) Error() string {
	return fmt.Sprintf("code:%d,msg:%s", e.Code, e.Msg)
}

func (e CodeError) GetCode() int32 {
	return e.Code
}

func GetMsg(code int32) string {
	return message[code]

}

func NewErr(code int32, msg string) error {
	return CodeError{
		Code: code,
		Msg:  msg,
	}
}

func NewSprintfErr(code int32, format string, a ...interface{}) error {
	return CodeError{
		Code: code,
		Msg:  fmt.Sprintf(format, a...),
	}
}

func ErrMsg(code int32) error {
	return CodeError{
		Code: code,
		Msg:  message[code],
	}

}
func ParamErr() error {
	return CodeError{
		Code: Param,
		Msg:  message[Param],
	}

}

func SystemErr() error {
	return CodeError{
		Code: SystemErrCode,
		Msg:  message[SystemErrCode],
	}
}

func DbErr(code int32) error {
	return kerrors.NewBizStatusError(code, message[code])
	//return CodeError{
	//	Code: code,
	//	Msg:  message[code],
	//}
}
