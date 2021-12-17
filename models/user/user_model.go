package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/global"
	"server/models/common"
)

type UserModel struct {
}

// GetUserByStuId 通过StuId获取user
func (u *UserModel) GetUserByStuId(ctx *gin.Context, studId string) (user *common.User, err error) {
	if err = global.GLOBAL_DB.First(&user, "stu_id = ?", studId).Error; err != nil {
		log.Println(fmt.Sprintf("get user failed, err=%v||user_name=%s", err, studId))
		return nil, err
	}
	if user == nil {
		return nil, global.ERRGETUSERINFO
	}
	return user, nil
}

// SaveUser 保存用户
func (u *UserModel) SaveUser(ctx *gin.Context, user *common.User) error {
	err := global.GLOBAL_DB.Create(&user).Error
	if err != nil {
		log.Println(fmt.Sprintf("create user_management failed, err=%v", err))
		return err
	}
	return nil
}
