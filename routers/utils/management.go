package utils

import (
	"github.com/gin-gonic/gin"
)

type ManagementRouter struct {
}

func (m *ManagementRouter) InitUtilsRouters(router *gin.RouterGroup) {
	group := router.Group("utils")
	{
		// 发送验证码
		group.POST("sendValidateCode", utilsApi.SendValidateCode)
	}
}
