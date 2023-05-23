package ctxdata

import (
	"context"
	"github.com/xu756/appserver/internal/xjwt"
	"github.com/zeromicro/go-zero/core/logx"
)

type key string

var userKey key

// NewContextForJwt 上下文设置用户信息
func NewContextForJwt(ctx context.Context, u *xjwt.AuthInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

// FromContextForJwt 从上下文中获取用户信息
func FromContextForJwt(ctx context.Context) (*xjwt.AuthInfo, bool) {
	if u := ctx.Value(userKey); u != nil {
		if value, ok := u.(*xjwt.AuthInfo); ok {
			return value, true
		}
	}
	logx.WithContext(ctx).Error("not found AuthInfo from ctx")
	return nil, false
}
