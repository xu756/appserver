package xerr

// OK 成功返回
const OK uint32 = 200

// 全局错误码
const (
	StystemEoore      uint32 = 401
	DbError           uint32 = 402 // 数据库错误
	ServerCommonError uint32 = 111 // 服务器通用错误
)
