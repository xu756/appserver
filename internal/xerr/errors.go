package xerr

import "fmt"

/**
常用通用固定错误
*/

type CodeError struct {
	errCode uint32
	errMsg  string
}

// GetErrCode 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func StystenError(err interface{}) *CodeError {
	return &CodeError{errCode: StystemEoore, errMsg: fmt.Sprintf("%v", err)}
}

// DbErr 数据库错误
func DbErr(err interface{}) *CodeError {
	return &CodeError{errCode: DbError, errMsg: fmt.Sprintf("%v", err)}
}

func MsgError(msg string) *CodeError {
	return &CodeError{errCode: MsError, errMsg: msg}
}

func LogOut(err interface{}) *CodeError {
	return &CodeError{errCode: LogOutError, errMsg: fmt.Sprintf("%v", err)}

}
