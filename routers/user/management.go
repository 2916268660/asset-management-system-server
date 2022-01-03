package user

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
)

type ManagementRouter struct {
}

func (m ManagementRouter) InitUserRouters(router *gin.RouterGroup) {
	group := router.Group("")
	{
		// 注册
		group.POST("register", userApi.RegisterUser)
		// 登录
		group.POST("login", userApi.Login)
		// 获取用户信息
		group.GET("details", middleware.JWTAuthMiddleWare(), userApi.GetUserInfo)
	}
}
