package user

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/common/request"
	"server/utils"
)

type LoginLogic struct {
}

func (l *LoginLogic) Login(ctx *gin.Context, loginInfo *request.LoginUserInfo) (token string, err error) {
	switch loginInfo.Way {
	case global.WayByUserName: // 通过用户名登录
		if err = loginByUserName(ctx, loginInfo); err != nil {
			return "", err
		}
	}
	// 获取token
	token, err = global.GetToken(loginInfo.StuId)
	if err != nil {
		return "", global.ERRTOKENGENERATE
	}
	return token, nil
}

// 通过用户名登录
func loginByUserName(ctx *gin.Context, loginInfo *request.LoginUserInfo) (err error) {
	if loginInfo.StuId == "" || loginInfo.Password == "" {
		return global.ERRARGS.WithMsg("学号或密码不能为空")
	}
	user, err := userModel.GetUserByStuId(ctx, loginInfo.StuId)
	if err != nil || user == nil {
		return global.ERRUSERNAMENOTEXIST
	}
	if user.Password != utils.Encrypt(loginInfo.Password) {
		return global.ERRPASSWORD
	}
	return nil
}
