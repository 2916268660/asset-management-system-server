package common

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/request"
	"server/utils"
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
	seconds, err := commonLogic.SendValidateCode(ctx, validateCode.Target, validateCode.Way)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "发送成功", map[string]int64{"expireTime": seconds})
}

// GetMenus 获取菜单
func (m *ManagementApi) GetMenus(ctx *gin.Context) {
	// 根据不同用户角色 获取不同的菜单栏
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		global.GLOBAL_LOG.Error("获取用户信息失败", zap.Error(err))
		global.FailWithMsg(ctx, err.Error())
		return
	}
	menus, err := commonLogic.GetMenus(ctx, claims)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "获取菜单列表成功", menus)
}
