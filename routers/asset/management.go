package asset

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
)

type ManagementRouter struct {
}

func (m *ManagementRouter) InitAssetRouters(router *gin.RouterGroup) {
	group := router.Group("asset", middleware.JWTAuthMiddleWare(), middleware.Casbin())
	{
		// 申请领用资产
		group.POST("applyReceive", assetApi.ApplyReceive)
		// 申请归还资产
		group.POST("applyRevert", assetApi.ApplyRevert)
		// 申请报修资产
		group.POST("applyRepair", assetApi.ApplyRepair)
		// 录入资产
		group.POST("addAsset", assetApi.AddAsset)
		// 删除资产
		group.DELETE("delAsset", assetApi.DelAsset)
		// 更新资产信息
		group.PUT("updateAsset", assetApi.UpdateAsset)
		// 获取某个或所有资产信息
		group.GET("getAsset", assetApi.GetAssets)
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
		// 获取发放人员待办
		group.GET("getProviderTodo", assetApi.GetProviderTodo)
		// 撤销申请
		group.POST("rollback", assetApi.Rollback)
		// 获取某个待办事项的详情信息
		group.GET("getTodoDetails", assetApi.GetTodoDetails)
		// 获取用户已领用的资产
		group.GET("getAssetsByUser", assetApi.GetAssetsByUser)
		// 发放资产
		group.PUT("provideAsset", assetApi.ProvideAsset)
		// 获取某个品类的所有资产
		group.GET("getAssetsByCategory", assetApi.GetAssetsByCategory)
		// 归还资产
		group.POST("revertAssets", assetApi.RevertAssets)
		// 驳回申请领用
		group.POST("rejectReceive", assetApi.RejectReceive)
		// 获取所有资产
		group.GET("getAllAssets", assetApi.GetAllAssets)
		// 获取资产的领用人信息
		group.GET("getAssetUser", assetApi.GetAssetUser)
	}
}
