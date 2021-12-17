package validate_code

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/common/request"
)

type ValidateCodeApi struct {
}

// SendValidateCode 发送验证码
func (v *ValidateCodeApi) SendValidateCode(ctx *gin.Context) {
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
	seconds, err := validateCodeLogic.SendValidateCode(ctx, validateCode.Target, validateCode.Way)
	if err != nil {
		global.Response(ctx, nil, err)
		return
	}
	global.Response(ctx, map[string]int64{"expireTime": seconds}, global.SUCCESS)
}
