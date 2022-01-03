package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/global"
	"server/models/common"
	"server/models/common/request"
	"server/models/common/response"
	"server/utils"
	"time"
)

type ManagementLogic struct {
}

// Login 登录
func (m *ManagementLogic) Login(ctx *gin.Context, loginInfo *request.LoginUserInfo) (token string, err error) {
	var user *common.User
	switch loginInfo.Way {
	case global.WayByUserName: // 通过用户名登录
		if user, err = loginByUserName(ctx, loginInfo); err != nil {
			return "", err
		}
	}
	// 获取token
	token, err = global.GetToken(user)
	if err != nil {
		return "", global.ERRTOKENGENERATE
	}
	return token, nil
}

// 通过用户名登录
func loginByUserName(ctx *gin.Context, loginInfo *request.LoginUserInfo) (user *common.User, err error) {
	if loginInfo.UserId == "" || loginInfo.Password == "" {
		return nil, global.ERRARGS.WithMsg("学号或密码不能为空")
	}
	user, err = userModel.GetUserByUserId(ctx, loginInfo.UserId)
	if err != nil || user == nil {
		return nil, global.ERRUSERNAMENOTEXIST
	}
	if user.Password != utils.Encrypt(loginInfo.Password) {
		return nil, global.ERRPASSWORD
	}
	return
}

// RegisterUser 注册用户
func (m *ManagementLogic) RegisterUser(ctx *gin.Context, info *request.RegisterUserInfo) error {
	// 前置校验
	err := preCheck(ctx, info)
	if err != nil {
		return err
	}
	now := time.Now()
	user := &common.User{
		UserName:   info.UserName,
		UserId:     info.UserId,
		Password:   utils.Encrypt(info.Password),
		Department: info.Department,
		Email:      info.Email,
		Phone:      info.Phone,
		CreateTime: now,
		UpdateTime: now,
	}
	err = userModel.SaveUser(ctx, user)
	if err != nil {
		log.Println(fmt.Sprintf("userId=%s exist, register failed", info.UserId))
		return global.ERRUSERNAMEISEXIST
	}
	return nil
}

// 前置校验
func preCheck(ctx *gin.Context, userInfo *request.RegisterUserInfo) error {
	if userInfo == nil {
		return global.ERRARGS
	}
	switch "" {
	case userInfo.UserId:
		return global.ERRARGS.WithMsg("账号不能为空")
	case userInfo.UserName:
		return global.ERRARGS.WithMsg("姓名不能为空")
	case userInfo.Password:
		return global.ERRARGS.WithMsg("密码不能为空")
	case userInfo.RePassword:
		return global.ERRARGS.WithMsg("确认密码不能为空")
	case userInfo.Phone:
		return global.ERRARGS.WithMsg("联系方式不能为空")
	case userInfo.Department:
		return global.ERRARGS.WithMsg("所属部门不能为空")
	}
	if userInfo.Password != userInfo.RePassword {
		return global.ERRARGS.WithMsg("输入的两次密码不一致")
	}
	return nil
}

// GetUserInfo 获取用户详情
func (m *ManagementLogic) GetUserInfo(ctx *gin.Context, userId string) (*response.UserInfo, error) {
	user, err := userModel.GetUserByUserId(ctx, userId)
	if err != nil {
		log.Println(fmt.Sprintf("userId=%s get userInfo failed, err=%v", userId, err))
		return nil, global.ERRGETUSERINFO
	}
	userInfo := &response.UserInfo{
		UserName: user.UserName,
		UserId:   user.UserId,
		Email:    user.Email,
		Phone:    user.Phone,
	}
	return userInfo, nil
}
