package user

import (
	"fmt"
	"log"
	"server/global"
	"server/models/common"

	"github.com/gin-gonic/gin"
)

type ManagementModel struct {
}

// GetUserByUserId 通过UserId获取user
func (m *ManagementModel) GetUserByUserId(ctx *gin.Context, studId string) (user *common.User, err error) {
	if err = global.GLOBAL_DB.First(&user, "user_id = ?", studId).Error; err != nil {
		log.Println(fmt.Sprintf("get user failed, err=%v||user_id=%s", err, studId))
		return nil, err
	}
	if user == nil {
		return nil, global.ERRGETUSERINFO
	}
	return user, nil
}

// SaveUser 保存用户
func (m *ManagementModel) SaveUser(ctx *gin.Context, user *common.User) error {
	err := global.GLOBAL_DB.Create(&user).Error
	if err != nil {
		log.Println(fmt.Sprintf("create user_management failed, err=%v", err))
		return err
	}
	return nil
}
