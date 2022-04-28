package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/utils"
)

// Casbin 角色权限中间件
func Casbin() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		claims, err := utils.GetClaims(ctx)
		if err != nil {
			global.FailWithMsg(ctx, err.Error())
			ctx.Abort()
			return
		}
		enforcer := utils.InitCasbin()
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		if claims.Role == "" {
			global.GLOBAL_LOG.Error("用户角色为空", zap.String("userId", claims.UserId))
			global.FailWithMsg(ctx, "用户角色错误")
			ctx.Abort()
			return
		}
		if claims.Role == global.Provider {
			ctx.Next()
			return
		}
		ok, _ := enforcer.Enforce(claims.Role, path, method)
		if ok {
			ctx.Next()
		} else {
			global.FailWithMsg(ctx, "权限不足")
			ctx.Abort()
			return
		}
	}
}
