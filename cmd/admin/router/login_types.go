package router

type LoginReq struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Mobile    string `json:"mobile"`
	Captcha   string `json:"captcha"`
	SessionId string `json:"sessionId,required"`
}
type SendCaptchaReq struct {
	Mobile    string `json:"mobile,required"`
	SessionId string `json:"sessionId,required"`
}
