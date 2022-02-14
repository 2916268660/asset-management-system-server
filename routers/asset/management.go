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
		// 录入资产
		group.POST("addAsset", assetApi.AddAsset)
		// 获取某个资产信息
		group.GET("getAsset/:serialId", assetApi.GetAssets)
		// 获取所有资产信息
		group.GET("getAsset", assetApi.GetAssets)
		// 获取资产的二维码
		group.GET("qrCode/:serialId", assetApi.GetQRCode)
		// 对申请的资产进行审核
		group.POST("audit", chargeApi.Audit)
		// 获取领用待办
		group.GET("getReceiveTodo", assetApi.GetReceiveTodo)
		// 获取归还待办
		group.GET("getRevertTodo", assetApi.GetRevertTodo)
		// 获取维修待办
		group.GET("getRepairsTodo", assetApi.GetRepairsTodo)
		// 获取审批待办
		group.GET("getAuditTodo", assetApi.GetAuditTodo)
		// 撤销申请
		group.POST("rollback", assetApi.Rollback)
	}
}
