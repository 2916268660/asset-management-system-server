package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"server/global"
	"server/models/common/request/user"
)

type RegisterApi struct {
}

// RegisterUser 注册用户
func (r *RegisterApi)RegisterUser(ctx *gin.Context) {
	var userInfo user.RegisterUserInfo
	err := ctx.ShouldBind(&userInfo)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			log.Println(fmt.Sprintf("submited args err||err=%v", err))
			global.Response(ctx, nil, global.ERRARGS.WithMsg("提交的信息有误,请仔细检查"))
			return
		}
		global.Response(ctx, global.ReMoveTopStruct(errs.Translate(global.Trans)), global.ERRARGS)
		return
	}
	// 注册用户
	err = registerLogic.RegisterUser(ctx, &userInfo)
	if err != nil {
		global.Response(ctx, nil, err)
		return
	}
	global.Response(ctx, nil, nil)
}

