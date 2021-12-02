package user

import (
	"github.com/gin-gonic/gin"
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
		UserName:     info.UserName,
		Password:     utils.Encrypt(info.Password),
		EmailOrPhone: info.EmailOrPhone,
		CreateTime:   now,
		UpdateTime:   now,
	}
	err := registerModel.SaveUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
