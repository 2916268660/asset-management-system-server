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
		global.FailWithMsg(ctx, "参数错误")
		return
	}
	if validateCode.Way <= 0 || validateCode.Target == "" {
		global.FailWithMsg(ctx, "参数错误")
		return
	}
	// 发送验证码
	seconds, err := utilsLogic.SendValidateCode(ctx, validateCode.Target, validateCode.Way)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "发送成功", map[string]int64{"expireTime": seconds})
}
