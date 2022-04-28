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
		// 批量注册
		group.POST("register2", userApi.RegisterUsers)
		// 登录
		group.POST("login", userApi.Login)
		// 获取用户信息
		group.GET("details", middleware.JWTAuthMiddleWare(), userApi.GetUserInfo)
		// 给用户设置部门管理员
		group.POST("setRole", middleware.JWTAuthMiddleWare(), middleware.Casbin(), userApi.SetRole)
		// 获取所有用户信息
		group.GET("getAllUser", middleware.JWTAuthMiddleWare(), middleware.Casbin(), userApi.GetAllUser)
		// 修该用户信息
		group.PUT("updateUser", middleware.JWTAuthMiddleWare(), middleware.Casbin(), userApi.UpdateUser)
		// 修改用户密码
		group.PUT("updatePass", middleware.JWTAuthMiddleWare(), middleware.Casbin(), userApi.UpdatePass)
		// 获取指定用户信息
		group.GET("getUser", middleware.JWTAuthMiddleWare(), middleware.Casbin(), userApi.GetUser)
		// 删除用户
		group.DELETE("delUser", middleware.JWTAuthMiddleWare(), middleware.Casbin(), userApi.DelUser)
	}
}
