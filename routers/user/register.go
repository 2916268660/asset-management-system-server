package user

import "github.com/gin-gonic/gin"

type RegisterRouters struct {
}

func (r *RegisterRouters) InitRegisterRouters(router *gin.RouterGroup) {
	group := router.Group("register")
	{
		group.POST("", registerApi.RegisterUser)
		group.POST("validateCode", validateCodeApi.SendValidateCode)
	}
}
