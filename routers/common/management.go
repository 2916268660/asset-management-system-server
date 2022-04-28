package common

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
)

type ManagementRouter struct {
}

func (m *ManagementRouter) InitUtilsRouters(router *gin.RouterGroup) {
	group := router.Group("common")
	{
		// 发送验证码
		group.POST("sendValidateCode", commonApi.SendValidateCode)
		// 获取菜单栏信息
		group.GET("menus", middleware.JWTAuthMiddleWare(), commonApi.GetMenus)
		// 获取资产的二维码
		group.GET("qrCode", assetApi.GetQRCode)
	}
}
