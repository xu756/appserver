// Code generated by goctl. DO NOT EDIT.
package types

type EmptyReq struct {
}

type UserInfo struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Role   int64  `json:"role"`
	OpenId string `json:"openId"`
}
