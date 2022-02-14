package common

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model"
	"server/model/response"
	"server/repository/cache"
	"server/utils"
	"time"
)

type ManagementLogic struct {
}

// 验证码有效时间
const expireTime = 60

// 验证码缓存中key的统一格式
var codeCachePre = "%s_code"

// SendValidateCode 发送验证码
func (m *ManagementLogic) SendValidateCode(ctx *gin.Context, target string, way int) (seconds int64, err error) {
	switch way {
	case global.WayByEmail:
		if err = sendValidateCodeByEmail(ctx, target); err != nil {
			return expireTime, err
		}
	}
	return expireTime, nil
}

// sendValidateCodeByEmail 通过邮箱发送验证码
func sendValidateCodeByEmail(ctx *gin.Context, email string) error {
	// 获取6位数字随机验证码
	code := utils.GetValidateCode()
	// 如果缓存中还存在，则返回，等待过期再进行缓存
	ok := cache.IsExistKey(fmt.Sprintf(codeCachePre, email))
	if ok {
		return errors.New("发送验证码过于频繁")
	}
	// 放缓存，存活60s
	err := cache.SetKey(fmt.Sprintf(codeCachePre, email), code, time.Second*60)
	if err != nil {
		global.GLOBAL_LOG.Error("缓存验证码失败", zap.Error(err))
		return errors.New("缓存验证码失败")
	}
	// 将验证码发送邮箱
	send := &utils.EmailRequest{
		Emails: []string{
			email,
		},
		Title: model.EmailTitle_ValidateCode,
		Body:  fmt.Sprintf(model.EmailBody_validateCode, code),
	}
	err = send.SendEmail()
	if err != nil {
		return errors.New("发送验证码失败")
	}
	return nil
}

// GetMenus 根据不同角色获取不同的菜单栏
func (m *ManagementLogic) GetMenus(ctx *gin.Context, claims *model.MyClaims) (res []*response.Menu, err error) {
	if claims.Role == "" {
		global.GLOBAL_LOG.Error("用户角色为空", zap.String("userId", claims.UserId))
		return nil, errors.New("用户权限错误")
	}
	res, ok := global.RoleMenusMap[claims.Role]
	if !ok {
		global.GLOBAL_LOG.Error("用户角色不存在", zap.String("userId", claims.UserId), zap.String("role", claims.Role))
		return nil, errors.New("用户权限错误")
	}
	return res, nil
}
