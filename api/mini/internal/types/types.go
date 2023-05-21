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
