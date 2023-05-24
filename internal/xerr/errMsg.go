package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
}

func GetErrDetail(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	}

	return "服务器开小差啦,稍后再来试一试"
}

func HasErrCode(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	}

	return false
}
