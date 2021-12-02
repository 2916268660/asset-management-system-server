package user

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models/common/request"
)

type LoginApi struct {
}

func (l *LoginApi) Login(ctx *gin.Context) {
	var userInfo request.LoginUserInfo
	if err := ctx.ShouldBind(&userInfo); err != nil {
		global.Response(ctx, nil, global.ERRARGS.WithData("账号或密码错误"))
		return
	}
	if userInfo.Way <= 0 {
		global.Response(ctx, nil, global.ERRARGS.WithData("登录方式错误"))
		return
	}
	err := loginLogic.Login(ctx, &userInfo)
	if err != nil {
		global.Response(ctx, nil, err)
		return
	}
	global.Response(ctx, map[string]string{"msg": "登录成功"}, nil)
}
