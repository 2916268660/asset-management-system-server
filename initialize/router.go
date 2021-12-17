package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "server/docs"
	"server/routers"
)

func InitRouters() *gin.Engine {
	var router = gin.Default()
	privateRouter := router.Group("v1")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 获取用户路由组实例
	userRouters := routers.RoutersGroupApp.UserRoutersGroup
	{
		userRouters.InitRegisterRouters(privateRouter)
		userRouters.InitLoginRouters(privateRouter)
		userRouters.InitDetailsRouters(privateRouter)
	}

	return router
}
