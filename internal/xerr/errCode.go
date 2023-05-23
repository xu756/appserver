package xerr

// OK 成功返回
const OK uint32 = 200

// 全局错误码
const (
	LoginError        uint32 = 401
	DbError           uint32 = 402 // 数据库错误
	TokenExpireError  uint32 = 403 // token 过期
	ServerCommonError uint32 = 500 // 服务器通用错误
	TypeConversionErr uint32 = 501 // 类型转换错误
)
