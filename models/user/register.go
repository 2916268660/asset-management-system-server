package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/global"
	"server/models/common"
)

type RegisterModel struct {
}

// SaveUser 保存用户
func (u *RegisterModel) SaveUser(ctx *gin.Context, user *common.User) error {
	err := global.GLOBAL_DB.Create(&user).Error
	if err != nil {
		log.Println(fmt.Sprintf("create user_management failed, err=%v", err))
		return err
	}
	return nil
}
