package utils

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
)

type ManagementRouter struct {
}

func (m *ManagementRouter) InitUtilsRouters(router *gin.RouterGroup) {
	group := router.Group("utils", middleware.JWTAuthMiddleWare())
	{
		// 发送验证码
		group.POST("sendValidateCode", utilsApi.SendValidateCode)
	}
}
