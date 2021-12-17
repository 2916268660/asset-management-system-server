package user

import (
	"github.com/gin-gonic/gin"
	"server/global"
)

/**
用户详情
*/
type DetailsApi struct {
}

// GetUserInfo 获取用户相信信息
func (d *DetailsApi) GetUserInfo(ctx *gin.Context) {
	stuId, ok := ctx.Get("stuId")
	if !ok {
		global.Response(ctx, nil, global.ERRGETUSERINFO)
		return
	}
	userInfo, err := detailsLogic.GetUserInfo(ctx, stuId.(string))
	global.Response(ctx, userInfo, err)
}
