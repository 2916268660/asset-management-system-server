package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"server/app/user/logic"
	"server/global"

	"github.com/gin-gonic/gin"
)

var userLogic = &logic.UserLogic{}

// RegisterUser 注册用户
func RegisterUser(ctx *gin.Context) {
	var userInfo logic.UserArgsForRegister
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
	err = userLogic.RegisterUser(ctx, &userInfo)
	if err != nil {
		global.Response(ctx, nil, err)
		return
	}
	global.Response(ctx, nil, nil)
}
