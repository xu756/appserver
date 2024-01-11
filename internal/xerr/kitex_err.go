package xerr

//
//import (
//	"context"
//	"errors"
//	"github.com/cloudwego/kitex/pkg/kerrors"
//)
//
//func ServerErrorHandler(ctx context.Context, err error) error {
//	if errors.Is(err, kerrors.ErrBiz) {
//		err = errors.Unwrap(err)
//	}
//	return err
//}
//
//func ClientErrorHandler(ctx context.Context, err error) error {
//	return err
//}
//
//
//func NewSprintfErr(code int32, format string, a ...interface{}) error {
//	err := CodeError{
//		Code: code,
//		Msg:  fmt.Sprintf(format, a...),
//	}
//	return remote.NewTransError(code, err)
//}
//
//func ErrMsg(code int32) error {
//
//	err := CodeError{
//		Code: code,
//		Msg:  message[code],
//	}
//	return remote.NewTransError(code, err)
//
//}
//func ParamErr() error {
//	err := CodeError{
//		Code: Param,
//		Msg:  message[Param],
//	}
//	return remote.NewTransError(Param, err)
//
//}
//
//func SystemErr() error {
//	err := CodeError{
//		Code: SystemErrCode,
//		Msg:  message[SystemErrCode],
//	}
//	return remote.NewTransError(SystemErrCode, err)
//}
