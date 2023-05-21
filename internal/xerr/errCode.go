package xerr

// OK 成功返回
const OK uint32 = 200

// 全局错误码
const (
	ServerCommonError uint32 = 100001 + iota
	RequestParamError
	TokenExpireError
	TokenGenerateError
	DbError
	DbUpdateAffectedZeroError
	DbErrInLoopFind
	CodeValidateFailure
	TypeConversionErr
)
