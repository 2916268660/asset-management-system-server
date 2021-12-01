package routers

import (
	"github.com/gin-gonic/gin"
	controllers "server/app/user/controllers"
)

func Routers(r *gin.Engine) {
	group := r.Group("/v1")
	{
		// 登录
		group.POST("/login", controllers.Login)
		// 注册
		group.POST("/register", controllers.RegisterUser)
	}

}
