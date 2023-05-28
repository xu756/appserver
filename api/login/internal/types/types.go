// Code generated by goctl. DO NOT EDIT.
package types

type LoginByMobileReq struct {
	Mobile  string `json:"mobile"`
	SmsCode string `json:"smscode"`
}

type LoginResp struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type MiniAuthReq struct {
	Code string `path:"code"`
}

type EmptyReq struct {
}

type CaptchaResp struct {
	Code  int64  `json:"code"`
	Key   string `json:"key"`
	Image string `json:"image"`
	Thumb string `json:"thumb"`
}

type CaptchaReq struct {
	Dots []int64 `json:"dots"`
	Key  string  `json:"key"`
}

type CaptchaCompareResp struct {
	Ok bool `json:"ok"`
}

type LoginByPasswordReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
