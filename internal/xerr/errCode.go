package xerr

// OK 成功返回
const OK uint32 = 200

// 全局错误码
const (
	StystemEoore      uint32 = 401
	DbError           uint32 = 402 // 数据库错误
	ServerCommonError uint32 = 111 // 服务器通用错误
	MsError           uint32 = 112 // 需要提醒用户的错误
	LogOutError       uint32 = 113 // 重新登录
)
