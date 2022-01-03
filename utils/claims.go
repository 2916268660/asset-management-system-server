package utils

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/common"
	"strings"
)

// GetClaims 通过header的token获取claims
func GetClaims(ctx *gin.Context) (claims *common.MyClaims, err error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return nil, global.ERRTOKENNONE
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, global.ERRTOKENFMT
	}
	claims, err = global.ParseToken(parts[1])
	return
}
