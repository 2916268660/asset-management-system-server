package utils

import (
	"errors"
	"server/global"
	"server/models/common"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetClaims 通过header的token获取claims
func GetClaims(ctx *gin.Context) (claims *common.MyClaims, err error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("无权限访问,请先登录")
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, errors.New("token格式有误")
	}
	claims, err = global.ParseToken(parts[1])
	return
}
