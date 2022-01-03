package asset

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
)

type ManagementRouter struct {
}

func (m *ManagementRouter) InitAssetRouters(router *gin.RouterGroup) {
	group := router.Group("asset", middleware.JWTAuthMiddleWare())
	{
		// 申请领用资产
		group.POST("applyReceive", assetApi.ApplyReceive)
		// 申请归还资产
		group.POST("applyRevert", assetApi.ApplyRevert)
		// 申请报修资产
		group.POST("applyRepair", assetApi.ApplyRepair)
	}
}
