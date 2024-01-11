package router

type LoginByPasswordReq struct {
	Username  string `json:"username,required"`
	Password  string `json:"password,required"`
	SessionId string `json:"sessionId,required"`
}

type LoginByMobileReq struct {
	Mobile    string `json:"mobile,required"`
	Captcha   string `json:"captcha,required"`
	SessionId string `json:"sessionId,required"`
}

type SendCaptchaReq struct {
	Mobile    string `json:"mobile,required"`
	SessionId string `json:"sessionId,required"`
}
