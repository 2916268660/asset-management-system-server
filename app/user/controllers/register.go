package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/structs"
)


// RegisterUser 注册用户
func RegisterUser(c *gin.Context) {
	var userInfo structs.UserArgsForRegister
	err := c.ShouldBind(&userInfo)
	if err != nil {
		log.Fatal()
	}


}