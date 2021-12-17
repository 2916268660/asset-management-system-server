package middleware

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"strings"
)

func JWTAuthMiddleWare() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			global.Response(ctx, nil, global.ERRTOKENNONE)
			ctx.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			global.Response(ctx, nil, global.ERRTOKENFMT)
			ctx.Abort()
			return
		}
		mc, err := global.ParseToken(parts[1])
		if err != nil {
			if strings.Contains(err.Error(), "expired") {
				global.Response(ctx, nil, global.ERRTOKENTIMEOUT)
				ctx.Abort()
				return
			}
			global.Response(ctx, nil, global.ERRTOKENNONE)
			ctx.Abort()
			return
		}
		ctx.Set("stuId", mc.StuId)
		ctx.Next()
	}
}
