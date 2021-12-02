package user

import "github.com/gin-gonic/gin"

type RegisterRouters struct {
}

func (r *RegisterRouters) InitRegisterRouters(router *gin.RouterGroup) {
	registerRouter := router.Group("register")
	{
		registerRouter.POST("", registerApi.RegisterUser)
	}
}