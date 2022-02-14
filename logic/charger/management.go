package charger

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/request"
	"server/utils"
	"time"
)

type ManagementLogic struct {
}

// Audit 审核
func (m *ManagementLogic) Audit(ctx *gin.Context, info *request.AuditStatus) error {
	// 获取当前登录的用户
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		return err
	}
	receive, err := assetModel.GetReceiveById(ctx, info.ID)
	if err != nil {
		return errors.New("申请单不存在")
	}
	if receive.AdminId != claims.UserId {
		return errors.New("非法操作！无权限对该申请进行审核")
	}
	now := time.Now()
	if err = global.GLOBAL_DB.Model(&receive).Select("status", "audit_time").Where("id=?", info.ID).Updates(map[string]interface{}{
		"status":     info.Status,
		"audit_time": now,
	}).Error; err != nil {
		global.GLOBAL_LOG.Error("更新审批状态失败", zap.Error(err))
		return errors.New("审核失败,请联系管理员")
	}
	return nil
}
