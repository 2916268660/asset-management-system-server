package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/global"
	"server/models/common/request"
	"server/models/common/response"
	"server/utils"
)

type ManagementApi struct {
}

// Login 登录
func (m *ManagementApi) Login(ctx *gin.Context) {
	var userInfo request.LoginUserInfo
	if err := ctx.ShouldBind(&userInfo); err != nil {
		global.FailWithMsg(ctx, "提交的信息有误，请仔细检查后再次提交")
		return
	}
	if userInfo.Way <= 0 {
		global.FailWithMsg(ctx, "登录方式错误")
		return
	}
	// 登录
	token, err := userLogic.Login(ctx, &userInfo)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "登录成功", map[string]string{"token": "Bearer " + token})
}

// RegisterUser 注册用户
func (m *ManagementApi) RegisterUser(ctx *gin.Context) {
	var userInfo request.RegisterUserInfo
	err := ctx.ShouldBind(&userInfo)
	if err != nil {
		log.Println(fmt.Sprintf("submited args err||err=%v", err))
		global.FailWithMsg(ctx, "提交的信息有误，请仔细检查后再次提交")
		return
	}
	// 注册用户
	err = userLogic.RegisterUser(ctx, &userInfo)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithMsg(ctx, "注册成功")
}

// GetUserInfo 获取用户相信信息
func (m *ManagementApi) GetUserInfo(ctx *gin.Context) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithData(ctx, response.UserInfo{
		UserName:   claims.UserName,
		UserId:     claims.UserId,
		Email:      claims.Email,
		Phone:      claims.Phone,
		Department: claims.Department,
	})
}
