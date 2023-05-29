package xjwt

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/xu756/appserver/internal/tool"
	"github.com/xu756/appserver/internal/xerr"
	"net/http"
	"strings"
	"time"
)

var (
	ErrTokenExpired    = xerr.MsgError("令牌已过期")
	ErrTokenMalformed  = xerr.MsgError("请求令牌格式有误")
	ErrTokenInvalid    = xerr.MsgError("请求令牌无效")
	ErrHeaderEmpty     = xerr.MsgError("需要认证才能访问！")
	ErrHeaderMalformed = xerr.MsgError("请求头中 Authorization 格式有误")
)

type AuthInfo struct {
	ID     uint64 `json:"Id"`
	Job    string `json:"Job"`
	Issuer string `json:"Issuer"`
}

func AuthToInfo(b any) *AuthInfo {
	a := b.(map[string]interface{})
	i, _ := a["Id"].(json.Number).Int64()
	id := uint64(i)
	job := a["Job"].(string)
	issuer := a["Issuer"].(string)
	return &AuthInfo{
		ID:     id,
		Job:    job,
		Issuer: issuer,
	}
}

func (a *AuthInfo) GetStringID() uint64 {
	return a.ID
}

// GetJob 这里会返回三种角色的字符串注解，普通用户就是 "user" ...
func (a *AuthInfo) GetJob() string {
	return a.Job
}
func (a *AuthInfo) GetIssuer() string {
	return a.Issuer
}

type JWT struct {
	SignKey         string // 秘钥
	MaxRefresh      int64  // 刷新token的最大过期时间
	DevExpireTime   int64  // 开发模式令牌过期时间
	DebugExpireTime int64  // 调试模式令牌过期时间
	ProdExpireTime  int64  // 生产环境令牌过期时间
	Issuer          string // token颁发者
	Mode            string // 当前模块的运行状态
}

// CustomJWTClaims
//
//	 结构体实现了 Claims 接口继承了  Valid() 方法
//		jwt 规定了7个官方字段，提供使用:
//		  - iss (issuer)：发布者
//		  - sub (subject)：主题
//		  - iat (Issued At)：生成签名的时间
//		  - exp (expiration time)：签名过期时间
//		  - aud (audience)：观众，相当于接受者
//		  - nbf (Not Before)：生效时间
//		  - jti (jwt ID)：编号

type CustomJWTClaims struct {
	User AuthInfo `json:"user"`
	jwt.RegisteredClaims
}

// IssueToken 生成jwt，返回 token 字符串
func (j *JWT) IssueToken(u *AuthInfo) (string, error) {
	c := CustomJWTClaims{
		User: *u,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(tool.TimeNowInTimezone()), // 签名生效时间
			IssuedAt:  jwt.NewNumericDate(tool.TimeNowInTimezone()), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: jwt.NewNumericDate(j.expireAtTime()),         // 签名过期时间
			Issuer:    j.Issuer,                                     // 签名颁发者
		},
	}

	// 根据 claims 生成token对象
	token, err := j.createToken(c)
	if err != nil {
		return "", xerr.LogOut(err.Error())
	}

	return token, nil
}

// ParserToken 解析 Token，中间件中调用
func (j *JWT) ParserToken(tokenString string) (*CustomJWTClaims, error) {
	// 解析token
	token, err := j.parseTokenString(tokenString)
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		if ok {
			if validationErr.Errors == jwt.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwt.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	if claims, ok := token.Claims.(*CustomJWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// GetTokenFromHeader 获取请求头中的token
func (j *JWT) GetTokenFromHeader(r *http.Request, headerName string) (string, error) {
	authHeader := r.Header.Get(headerName)
	if authHeader == "" {
		fmt.Println(ErrHeaderEmpty)
		return "", ErrHeaderEmpty
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		fmt.Println(ErrHeaderMalformed)
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}

// createToken 创建 Token，内部使用，外部请调用 IssueToken
func (j *JWT) createToken(claims CustomJWTClaims) (string, error) {
	// 使用HS256算法进行token生成
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString([]byte(j.SignKey))
}

// parseTokenString 解析token
func (j *JWT) parseTokenString(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &CustomJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SignKey), nil
	})
}

// token过期时间
func (j *JWT) expireAtTime() time.Time {
	timezone := tool.TimeNowInTimezone()
	expire := time.Duration(j.GetExpire()) * time.Minute
	return timezone.Add(expire)
}

// GetExpire 获得当前模式下令牌过期时间
func (j *JWT) GetExpire() int64 {
	var expireTime int64
	switch j.Mode {
	case "dev": // 开发环境
		expireTime = j.DevExpireTime
	case "test": // 测试环境
		expireTime = j.DebugExpireTime
	case "pro": // 生产环境
		expireTime = j.ProdExpireTime
	case "rt": // 回归测试
		expireTime = j.ProdExpireTime
	case "pre": // 预发
		expireTime = j.ProdExpireTime
	}
	return expireTime
}

// GetMaxRefresh 获得令牌最大有效期
func (j *JWT) GetMaxRefresh() int64 {
	return j.MaxRefresh
}
