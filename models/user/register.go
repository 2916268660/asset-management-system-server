package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/global"
	"time"
)

type User struct {
	ID           int64     // 主键ID
	UserName     string    // 用户名
	Password     string    // 用户密码
	EmailOrPhone string    //邮箱或者电话
	CreateTime   time.Time //创建时间
	UpdateTime   time.Time //更新时间
}

type RegisterModel struct {
}

// SaveUser 保存用户
func (u *RegisterModel) SaveUser(ctx *gin.Context, user *User) error {
	err := global.GLOBAL_DB.Create(&user).Error
	if err != nil {
		log.Println(fmt.Sprintf("create user_management failed, err=%v", err))
		return err
	}
	return nil
}