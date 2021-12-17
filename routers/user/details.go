package user

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
)

type DetailsRouters struct {
}

func (d *DetailsRouters) InitDetailsRouters(router *gin.RouterGroup) {
	group := router.Group("details", middleware.JWTAuthMiddleWare())
	{
		group.GET("", detailsApi.GetUserInfo)
	}
}
