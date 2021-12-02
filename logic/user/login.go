package user

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/common/request"
	"server/utils"
)

type LoginLogic struct {
}

func (l *LoginLogic) Login(ctx *gin.Context, loginInfo *request.LoginUserInfo) (err error) {
	switch loginInfo.Way {
	case global.LoginWayByUserName:
		if err = loginByUserName(ctx, loginInfo); err != nil {
			return err
		}
	}
	return nil
}

// 通过用户名登录
func loginByUserName(ctx *gin.Context, loginInfo *request.LoginUserInfo) (err error) {
	if loginInfo.UserName == "" || loginInfo.Password == "" {
		return global.ERRARGS.WithData("用户名或密码不能为空")
	}
	user, err := loginModel.GetUserByUserName(ctx, loginInfo.UserName)
	if err != nil || user == nil {
		return global.ERRUSERNAME
	}
	if user.Password != utils.Encrypt(loginInfo.Password) {
		return global.ERRPASSWORD
	}
	return nil
}
