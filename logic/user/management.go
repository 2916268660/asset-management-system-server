package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/utils"
	"time"
)

type ManagementLogic struct {
}

// Login 登录
func (m *ManagementLogic) Login(ctx *gin.Context, loginInfo *request.LoginUserInfo) (token string, err error) {
	var user *model.SysUser
	switch loginInfo.Way {
	case global.WayByUserName: // 通过用户名登录
		if user, err = loginByUserName(ctx, loginInfo); err != nil {
			return "", err
		}
	default:
		return "", errors.New("登录方式错误")
	}
	// 获取token
	token, err = global.GetToken(user)
	if err != nil {
		return "", errors.New("生成token失败")
	}
	return token, nil
}

// 通过用户名登录
func loginByUserName(ctx *gin.Context, loginInfo *request.LoginUserInfo) (user *model.SysUser, err error) {
	if loginInfo.UserId == "" || loginInfo.Password == "" {
		return nil, errors.New("学号或密码不能为空")
	}
	user, err = userModel.GetUserByUserId(ctx, loginInfo.UserId)
	if err != nil || user == nil {
		return nil, errors.New("用户不存在")
	}
	if user.Password != utils.Encrypt(loginInfo.Password) {
		return nil, errors.New("密码错误,请重新输入")
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
	err = userModel.FindUserByUserId(ctx, info.UserId)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("用户名已存在")
	}
	now := time.Now()
	user := &model.SysUser{
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
		global.GLOBAL_LOG.Error("创建用户失败", zap.Any("user", user), zap.Error(err))
		return errors.New("注册失败")
	}
	return nil
}

// 前置校验
func preCheck(ctx *gin.Context, userInfo *request.RegisterUserInfo) error {
	if userInfo == nil {
		return errors.New("参数有误")
	}
	switch "" {
	case userInfo.UserId:
		return errors.New("账号不能为空")
	case userInfo.UserName:
		return errors.New("姓名不能为空")
	case userInfo.Password:
		return errors.New("密码不能为空")
	case userInfo.RePassword:
		return errors.New("确认密码不能为空")
	case userInfo.Phone:
		return errors.New("联系方式不能为空")
	case userInfo.Department:
		return errors.New("所属部门不能为空")
	}
	if userInfo.Password != userInfo.RePassword {
		return errors.New("输入的两次密码不一致")
	}
	return nil
}

// GetUserInfo 获取用户详情
func (m *ManagementLogic) GetUserInfo(ctx *gin.Context, userId string) (*response.UserInfo, error) {
	user, err := userModel.GetUserByUserId(ctx, userId)
	if err != nil {
		global.GLOBAL_LOG.Error("获取用户信息失败", zap.String("userId", userId), zap.Error(err))
		return nil, errors.New("获取用户信息失败")
	}
	userInfo := &response.UserInfo{
		UserName: user.UserName,
		UserId:   user.UserId,
		Email:    user.Email,
		Phone:    user.Phone,
	}
	return userInfo, nil
}
