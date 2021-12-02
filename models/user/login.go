package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/global"
	"server/models/common"
)

type LoginModel struct {
}

// GetUserByUserName 通过username获取user
func (l *LoginModel) GetUserByUserName(ctx *gin.Context, userName string) (user *common.User, err error) {
	if err = global.GLOBAL_DB.First(&user, "user_name = ?", userName).Error; err != nil {
		log.Println(fmt.Sprintf("get user failed, err=%v||user_name=%s", err, userName))
		return nil, err
	}
	return user, nil
}
