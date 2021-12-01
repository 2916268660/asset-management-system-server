package routers

import (
	"github.com/gin-gonic/gin"
	controllers2 "server/app/work/controllers"
	"server/middleware"
)

func Routers(e *gin.Engine) {
	workGroup := e.Group("/work", middleware.JWTAuthMiddleWare())
	workGroup.GET("/home", controllers2.Home)
}
