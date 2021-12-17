package user

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/common/request"
)

type LoginApi struct {
}

// @Tags 用户
// @Summary  用户登录API
// @title Swagger Example API
// @version 0.0.1
// @description  用户登录.
// @Host 127.0.0.1:8080
// @Produce  json
// @Param username body string true "账号"
// @Param password body string true "密码"
// @Param email body string false "邮箱"
// @Param phone body string false "电话"
// @Param way body int true "登录方式"
// @Success 200 {string} json "{"code":2000,"data":"token","msg":"成功||登录成功"}"
// @Router /v1/login [post]
func (l *LoginApi) Login(ctx *gin.Context) {
	var userInfo request.LoginUserInfo
	if err := ctx.ShouldBind(&userInfo); err != nil {
		global.Response(ctx, nil, global.ERRARGS.WithMsg("账号或密码错误"))
		return
	}
	if userInfo.Way <= 0 {
		global.Response(ctx, nil, global.ERRARGS.WithMsg("登录方式错误"))
		return
	}
	// 登录
	token, err := loginLogic.Login(ctx, &userInfo)
	if err != nil {
		global.Response(ctx, nil, err)
		return
	}
	global.Response(ctx, map[string]string{"token": "Bearer " + token}, global.SUCCESS.WithMsg("登录成功"))
}
