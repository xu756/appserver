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

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: GetErrDetail(errCode)}
}

//func NewErrMsg(errMsg string) *CodeError {
//	return &CodeError{errCode: ServerCommonError, errMsg: errMsg}
//}
//
//// TypeConversionError 类型转换出现错误
//func TypeConversionError(err error) *CodeError {
//	fmt.Println("类型转换出现错误: ", err.Error())
//	return &CodeError{errCode: TypeConversionErr, errMsg: "类型转换出现错误: " + err.Error()}
//}

// DatabaseErr 数据库错误
func DatabaseErr(err interface{}) *CodeError {
	return &CodeError{errCode: DbError, errMsg: fmt.Sprintf("%v", err)}
}
