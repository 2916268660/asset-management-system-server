package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/global"
	"server/models/common/request"
)

type ManagementApi struct {
}

// Login 登录
func (m *ManagementApi) Login(ctx *gin.Context) {
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
	token, err := userLogic.Login(ctx, &userInfo)
	if err != nil {
		global.Response(ctx, nil, err)
		return
	}
	global.Response(ctx, map[string]string{"token": "Bearer " + token}, global.SUCCESS.WithMsg("登录成功"))
}

// RegisterUser 注册用户
func (m *ManagementApi) RegisterUser(ctx *gin.Context) {
	var userInfo request.RegisterUserInfo
	err := ctx.ShouldBind(&userInfo)
	if err != nil {
		log.Println(fmt.Sprintf("submited args err||err=%v", err))
		global.Response(ctx, nil, global.ERRARGS.WithMsg("提交的信息有误,请仔细检查"))
		return
	}
	// 注册用户
	err = userLogic.RegisterUser(ctx, &userInfo)
	if err != nil {
		global.Response(ctx, nil, err)
		return
	}
	global.Response(ctx, nil, nil)
}

// GetUserInfo 获取用户相信信息
func (m *ManagementApi) GetUserInfo(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		global.Response(ctx, nil, global.ERRGETUSERINFO)
		return
	}
	userInfo, err := userLogic.GetUserInfo(ctx, userId.(string))
	global.Response(ctx, userInfo, err)
}
