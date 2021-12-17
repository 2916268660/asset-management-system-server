package user

import "github.com/gin-gonic/gin"

type LoginRouters struct {
}

func (l *LoginRouters) InitLoginRouters(router *gin.RouterGroup) {
	group := router.Group("login")
	{
		group.POST("", loginApi.Login)
		group.POST("validateCode", validateCodeApi.SendValidateCode)
	}
}
