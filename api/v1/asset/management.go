package asset

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/global"
	"server/model"
	"server/model/request"
)

type ManagementApi struct {
}

// ApplyReceive 申请领用资产
func (m *ManagementApi) ApplyReceive(ctx *gin.Context) {
	var applyInfo *request.ApplyReceiveForm
	if err := ctx.ShouldBind(&applyInfo); err != nil {
		global.GLOBAL_LOG.Error("提交的信息有误", zap.Error(err))
		global.FailWithMsg(ctx, "提交的信息有误，请仔细检查后再次提交")
		return
	}
	taskId, err := assetLogic.ApplyReceive(ctx, applyInfo)
	if err != nil {
		global.GLOBAL_LOG.Error("申请资产失败", zap.Error(err))
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "申请成功", map[string]int64{"taskId": taskId})
}

// ApplyRevert 申请归还资产
func (m *ManagementApi) ApplyRevert(ctx *gin.Context) {
	var applyInfo *request.ApplyRevertForm
	if err := ctx.ShouldBind(&applyInfo); err != nil {
		global.GLOBAL_LOG.Error("提交的信息有误", zap.Error(err))
		global.FailWithMsg(ctx, "提交的信息有误，请仔细检查后再次提交")
		return
	}
	taskId, err := assetLogic.ApplyRevert(ctx, applyInfo)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "申请成功", map[string]int64{"taskId": taskId})
}

// ApplyRepair 申请维修资产
func (m *ManagementApi) ApplyRepair(ctx *gin.Context) {
	var applyInfo *request.ApplyRepairForm
	if err := ctx.ShouldBind(&applyInfo); err != nil {
		global.GLOBAL_LOG.Error("提交的信息有误", zap.Error(err))
		global.FailWithMsg(ctx, "提交的信息有误，请仔细检查后再次提交")
		return
	}
	repairId, err := assetLogic.ApplyRepair(ctx, applyInfo)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "申请成功", map[string]int64{"repairId": repairId})
}

// AddAsset 录入新资产
func (m *ManagementApi) AddAsset(ctx *gin.Context) {
	var assetInfo *request.AssetInfo
	if err := ctx.ShouldBind(&assetInfo); err != nil {
		global.GLOBAL_LOG.Error("提交的信息有误", zap.Error(err))
		global.FailWithMsg(ctx, "提交的信息有误，请仔细检查后再次提交")
		return
	}
	err := assetLogic.AddAsset(ctx, assetInfo)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithMsg(ctx, "新增资产成功")
}

// GetAssets 获取资产信息
func (m *ManagementApi) GetAssets(ctx *gin.Context) {
	serialId := ctx.Param("serialId")
	if serialId == "" {
		var assets []*model.AssetDetails
		if err := global.GLOBAL_DB.Find(&assets).Error; err != nil {
			global.GLOBAL_LOG.Error("查询所有资产信息失败", zap.Error(err))
			global.FailWithMsg(ctx, "查询失败")
			return
		}
		global.OkWithDetails(ctx, "查询成功", map[string][]*model.AssetDetails{"assets": assets})
		return
	}
	var asset *model.AssetDetails
	if err := global.GLOBAL_DB.Where("serial_id=?", serialId).First(&asset).Error; err != nil {
		global.GLOBAL_LOG.Error("查询指定资产信息失败", zap.String("serialId", serialId), zap.Error(err))
		global.FailWithMsg(ctx, "查询失败")
		return
	}
	global.OkWithDetails(ctx, "查询成功", map[string]*model.AssetDetails{"asset": asset})
}

// GetQRCode 获取资产的二维码
func (m *ManagementApi) GetQRCode(ctx *gin.Context) {
	serialId := ctx.Param("serialId")
	var asset *model.AssetDetails
	if errors.Is(global.GLOBAL_DB.Select("serial_img").Where("serial_id=?", serialId).Find(&asset).Error, gorm.ErrRecordNotFound) || asset.SerialImg == "" {
		global.FailWithMsg(ctx, "二维码不存在")
		return
	}
	ctx.File(asset.SerialImg)
}

// GetReceiveTodo 获取领用待办
func (m *ManagementApi) GetReceiveTodo(ctx *gin.Context) {
	todo, err := assetLogic.GetReceiveTodo(ctx)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithData(ctx, todo)
}

// GetRevertTodo 获取归还待办
func (m *ManagementApi) GetRevertTodo(ctx *gin.Context) {
	todo, err := assetLogic.GetRevertTodo(ctx)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithData(ctx, todo)
}

// GetRepairsTodo 获取维修待办
func (m *ManagementApi) GetRepairsTodo(ctx *gin.Context) {
	todo, err := assetLogic.GetRepairsTodo(ctx)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithData(ctx, todo)
}

// GetAuditTodo 获取审批待办
func (m *ManagementApi) GetAuditTodo(ctx *gin.Context) {
	todo, err := assetLogic.GetAuditTodo(ctx)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithData(ctx, todo)
}

// Rollback 撤销申请
func (m *ManagementApi) Rollback(ctx *gin.Context) {
	var info *request.RollbackInfo
	if err := ctx.ShouldBind(&info); err != nil {
		global.GLOBAL_LOG.Error("提交的信息有误", zap.Error(err))
		global.FailWithMsg(ctx, "提交的信息有误")
		return
	}
	err := assetLogic.Rollback(ctx, info)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.Ok(ctx)
}
