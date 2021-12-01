package controllers

import (
	"fmt"
	"log"
	"server/app/user/logic"
	"server/global"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var userLogic = &logic.UserLogic{}

// RegisterUser 注册用户
func RegisterUser(ctx *gin.Context) {
	var userInfo logic.UserArgsForRegister
	err := ctx.ShouldBind(&userInfo)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			log.Println(fmt.Sprintf("提交的信息有误,请仔细检查||err=%v", err))
			global.Response(ctx, nil, global.ERRARGS)
			return
		}
		global.Response(ctx, errs.Translate(global.Trans), global.ERRARGS)
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
