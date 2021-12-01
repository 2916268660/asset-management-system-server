package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg": "不存在",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg": "success",
		"data": gin.H{
			"username": username,
		},
	})
	return
}
