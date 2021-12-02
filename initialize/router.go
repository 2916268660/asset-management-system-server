package initialize

import (
	"github.com/gin-gonic/gin"
	"server/routers"
)

func InitRouters() *gin.Engine {
	var router = gin.Default()
	privateRouter := router.Group("v1")

	// 获取用户路由组实例
	userRouters := routers.RoutersGroupApp.UserRoutersGroup
	{
		userRouters.InitRegisterRouters(privateRouter)
		userRouters.InitLoginRouters(privateRouter)
	}

	return router
}
