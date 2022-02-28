package core

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
	"server/routers"
)

func InitRouters() *gin.Engine {
	var router = gin.Default()
	// 配置跨域中间件
	router.Use(middleware.Cors())

	v1Router := router.Group("v1")
	// 获取用户路由组实例
	userRouters := routers.RoutersGroupApp.UserRouterGroup
	// 获取工具路由组实例
	utilsRouters := routers.RoutersGroupApp.UtilsRouterGroup
	// 获取资产路由组实例
	assetRouters := routers.RoutersGroupApp.AssetRouterGroup

	{
		userRouters.InitUserRouters(v1Router)
		utilsRouters.InitUtilsRouters(v1Router)
		assetRouters.InitAssetRouters(v1Router)
	}

	return router
}
