package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/global"
	"server/models/common/response"
)

type DetailsLogic struct {
}

// GetUserInfo 获取用户详情
func (d *DetailsLogic) GetUserInfo(ctx *gin.Context, stuId string) (*response.UserInfo, error) {
	user, err := userModel.GetUserByStuId(ctx, stuId)
	if err != nil {
		log.Println(fmt.Sprintf("stuId=%s get userInfo failed, err=%v", stuId, err))
		return nil, global.ERRGETUSERINFO
	}
	userInfo := &response.UserInfo{
		UserName: user.UserName,
		StuId:    user.StuId,
		Email:    user.Email,
		Phone:    user.Phone,
	}
	return userInfo, nil
}
