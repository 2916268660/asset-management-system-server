package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/global"
	"server/structs"
)


func Login(c *gin.Context) {
	var user structs.UserArgs
	err := c.ShouldBind(&user)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
			"data": gin.H{
				"err": err,
			},
		})
		return
	}
	if user.UserName == "ws" && user.Password == "970107" {
		token, _ := global.GetToken(user.UserName)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{
				"token": token,
			},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
}
