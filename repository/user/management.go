package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model"
)

type ManagementModel struct {
}

// GetUserByUserId 通过UserId获取user
func (m *ManagementModel) GetUserByUserId(ctx *gin.Context, userId string) (user *model.SysUser, err error) {
	if err = global.GLOBAL_DB.First(&user, "user_id = ?", userId).Error; err != nil {
		global.GLOBAL_LOG.Error("get userByUserId failed", zap.String("user_id", userId), zap.Error(err))
		return nil, err
	}
	if user == nil {
		return nil, err
	}
	return user, nil
}

// SaveUser 保存用户
func (m *ManagementModel) SaveUser(ctx *gin.Context, user *model.SysUser) error {
	err := global.GLOBAL_DB.Create(&user).Error
	if err != nil {
		global.GLOBAL_LOG.Error("create user failed", zap.Error(err))
		return err
	}
	return nil
}

func (m *ManagementModel) FindUserByUserId(ctx *gin.Context, userId string) (user *model.SysUser, err error) {
	err = global.GLOBAL_DB.Where("user_id=?", userId).First(&user).Error
	return
}

// GetChargerByDepart 获取部门负责人信息
func (m *ManagementModel) GetChargerByDepart(ctx *gin.Context, department string) (user *model.SysUser, err error) {
	var charger *model.Department
	if err = global.GLOBAL_DB.Where("department=?", department).First(&charger).Error; err != nil {
		global.GLOBAL_LOG.Error("获取部门负责人失败", zap.String("department", department), zap.Error(err))
		return nil, err
	}
	if err = global.GLOBAL_DB.Where("user_id", charger.UserId).First(&user).Error; err != nil {
		global.GLOBAL_LOG.Error("获取部门负责人信息失败", zap.String("user_id", charger.UserId), zap.Error(err))
		return nil, err
	}
	return
}
