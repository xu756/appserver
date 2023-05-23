package handler

import (
	"fmt"
	"github.com/xu756/appserver/api/conn/imserver/logic"
	"github.com/xu756/appserver/api/conn/imserver/svc"
	"github.com/xu756/appserver/internal/ctxdata"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"nhooyr.io/websocket"
	"strconv"
)

func addConn(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr, err := svcCtx.Jwt.GetTokenFromHeader(r, "Authorization")
		if err != nil {
			logx.WithContext(r.Context()).Error("【中间件 middleware】 error:", err)
			fmt.Print("【中间件 middleware】 error:", err)
			return
		}
		jwtClaims, err := svcCtx.Jwt.ParserToken(tokenStr)
		if err != nil {
			logx.WithContext(r.Context()).Error("【中间件 middleware】 error:", err)
			return
		}

		strId := strconv.FormatUint(jwtClaims.User.ID, 10)
		if _, err = svcCtx.RedisClient.Get("cache:appserver:user:id:" + strId); err != nil {
			logx.WithContext(r.Context()).Error("中间件 middleware error:", err)
			return
		}

		ctx := ctxdata.NewContextForJwt(r.Context(), &jwtClaims.User)
		r = r.WithContext(ctx)
		fmt.Print("user:", jwtClaims.User.GetJob())
		c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			Subprotocols:         nil,
			InsecureSkipVerify:   false,
			OriginPatterns:       nil,
			CompressionMode:      0,
			CompressionThreshold: 0,
		})
		if err != nil {
			logx.WithContext(r.Context()).Error("【中间件 middleware】升级请求头 error:", err)
			return
		}
		l := logic.NewAddConnLogic(r.Context(), svcCtx)
		err = l.AddConn(w, r, c)
		if err != nil {
			logx.WithContext(r.Context()).Error("【中间件 middleware】添加连接 error:", err)
			return
		}
	}
}
