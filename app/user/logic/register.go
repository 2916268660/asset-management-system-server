package logic

import (
	model "server/models/user_management"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type UserLogic struct {
}

var userModel = &model.UserModel{}

// UserArgsForRegister 用户注册信息
type UserArgsForRegister struct {
	UserName     string `json:"username" binding:"required"`                    //用户名 唯一
	Password     string `json:"password" binding:"required"`                    //密码
	RePassword   string `json:"rePassword" binding:"required,eqfield=Password"` //第二次输入密码
	EmailOrPhone string `json:"emailOrPhone" binding:"required"`                //邮箱或者电话
	//ValidateCode   string `json:"validateCode" binding:"required"`   //验证码
}

func (u *UserLogic) RegisterUser(ctx *gin.Context, info *UserArgsForRegister) error {
	// 验证二维码是否正确
	now := time.Now()
	user := &model.User{
		UserName:     info.UserName,
		Password:     utils.Encrypt(info.Password),
		EmailOrPhone: info.EmailOrPhone,
		CreateTime:   now,
		UpdateTime:   now,
	}
	err := userModel.SaveUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
