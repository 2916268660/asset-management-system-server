package user

import (
	"github.com/gin-gonic/gin"
	"server/models/common/request/user"
	model "server/models/user"
	"server/utils"
	"time"
)

type RegisterLogic struct {
}




func (u *RegisterLogic) RegisterUser(ctx *gin.Context, info *user.RegisterUserInfo) error {
	// 验证二维码是否正确
	now := time.Now()
	user := &model.User{
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