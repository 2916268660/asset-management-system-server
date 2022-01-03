package utils

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/common/request"
)

type ManagementApi struct {
}

// SendValidateCode 发送验证码
func (m *ManagementApi) SendValidateCode(ctx *gin.Context) {
	var validateCode request.ValidateCodeWay
	err := ctx.ShouldBind(&validateCode)
	if err != nil {
		global.Response(ctx, nil, global.ERRARGS)
		return
	}
	if validateCode.Way <= 0 || validateCode.Target == "" {
		global.Response(ctx, nil, global.ERRARGS)
		return
	}
	// 发送验证码
	seconds, err := utilsLogic.SendValidateCode(ctx, validateCode.Target, validateCode.Way)
	if err != nil {
		global.Response(ctx, nil, err)
		return
	}
	global.Response(ctx, map[string]int64{"expireTime": seconds}, global.SUCCESS)
}
