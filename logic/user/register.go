package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/global"
	"server/models/common"
	"server/models/common/request"
	"server/utils"
	"time"
)

type RegisterLogic struct {
}

func (u *RegisterLogic) RegisterUser(ctx *gin.Context, info *request.RegisterUserInfo) error {
	now := time.Now()
	user := &common.User{
		UserName:   info.UserName,
		StuId:      info.StuId,
		Password:   utils.Encrypt(info.Password),
		Email:      info.Email,
		Phone:      info.Phone,
		CreateTime: now,
		UpdateTime: now,
	}
	err := userModel.SaveUser(ctx, user)
	if err != nil {
		log.Println(fmt.Sprintf("username=%s exist, register failed", info.UserName))
		return global.ERRUSERNAMEISEXIST
	}
	return nil
}
