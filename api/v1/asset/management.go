package asset

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/utils"
	"strconv"
	"time"
)

type ManagementApi struct {
}

// ApplyReceive 申请领用资产
func (m *ManagementApi) ApplyReceive(ctx *gin.Context) {
	var applyInfo *request.ApplyReceiveForm
	if err := ctx.ShouldBind(&applyInfo); err != nil {
		global.GLOBAL_LOG.Error("提交的信息有误", zap.Error(err))
		global.FailWithMsg(ctx, "提交的信息不合规, 请仔细检查后再次提交")
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
		global.FailWithMsg(ctx, "提交的信息不合规, 请仔细检查后再次提交")
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
		global.FailWithMsg(ctx, "提交的信息不合规, 请仔细检查后再次提交")
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
		global.FailWithMsg(ctx, "提交的信息不合规, 请仔细检查后再次提交")
		return
	}
	err := assetLogic.AddAssets(ctx, assetInfo)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithMsg(ctx, "新增资产成功")
}

// DelAsset 删除资产
func (m *ManagementApi) DelAsset(ctx *gin.Context) {
	serialId := ctx.Query("serialId")
	if serialId == "" {
		global.GLOBAL_LOG.Error("资产序列号为空")
		global.FailWithMsg(ctx, "资产不存在")
		return
	}
	asset := &model.AssetDetails{}
	if err := global.GLOBAL_DB.Where("serial_id = ?", serialId).Delete(&asset).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.GLOBAL_LOG.Error("资产序列号不存在", zap.String("serial_id", serialId))
			global.FailWithMsg(ctx, "资产不存在")
			return
		}
		global.GLOBAL_LOG.Error("删除资产失败", zap.Error(err))
		global.FailWithMsg(ctx, "删除失败")
		return
	}
	global.Ok(ctx)
}

// UpdateAsset 更新资产信息
func (m *ManagementApi) UpdateAsset(ctx *gin.Context) {
	var assetInfo *request.UpdateAssetInfo
	if err := ctx.ShouldBind(&assetInfo); err != nil {
		global.GLOBAL_LOG.Error("提交的信息有误", zap.Error(err))
		global.FailWithMsg(ctx, "提交的信息不合规, 请仔细检查后再次提交")
		return
	}
	asset, err := assetLogic.UpdateAsset(ctx, assetInfo)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "更新成功", asset)
}

// GetAssets 获取资产信息
func (m *ManagementApi) GetAssets(ctx *gin.Context) {
	serialId := ctx.Query("serialId")
	if serialId == "" {
		var assets []*model.AssetDetails
		if err := global.GLOBAL_DB.Find(&assets).Error; err != nil {
			global.GLOBAL_LOG.Error("查询所有资产信息失败", zap.Error(err))
			global.FailWithMsg(ctx, "查询失败")
			return
		}
		global.OkWithDetails(ctx, "", map[string][]*model.AssetDetails{"assets": assets})
		return
	}
	var asset *model.AssetDetails
	if err := global.GLOBAL_DB.Where("serial_id=?", serialId).First(&asset).Error; err != nil {
		global.GLOBAL_LOG.Error("查询指定资产信息失败", zap.String("serialId", serialId), zap.Error(err))
		global.FailWithMsg(ctx, "查询失败")
		return
	}
	global.OkWithDetails(ctx, "", map[string]*model.AssetDetails{"asset": asset})
}

// GetQRCode 获取资产的二维码
func (m *ManagementApi) GetQRCode(ctx *gin.Context) {
	serialId := ctx.Query("serialId")
	var asset *model.AssetDetails
	if errors.Is(global.GLOBAL_DB.Select("serial_img").Where("serial_id=?", serialId).Find(&asset).Error, gorm.ErrRecordNotFound) || asset.SerialImg == "" {
		global.FailWithMsg(ctx, "二维码不存在")
		return
	}
	ctx.File(asset.SerialImg)
}

// GetReceiveTodo 获取领用待办
func (m *ManagementApi) GetReceiveTodo(ctx *gin.Context) {
	pageNum, _ := ctx.GetQuery("pageNum")
	pageSize, _ := ctx.GetQuery("pageSize")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	page := &model.PageType{
		PageNum:  num,
		PageSize: size,
	}
	todo, total, err := assetLogic.GetReceiveTodo(ctx, page)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "", map[string]interface{}{
		"total": total,
		"todo":  todo,
	})
}

// GetRevertTodo 获取归还待办
func (m *ManagementApi) GetRevertTodo(ctx *gin.Context) {
	pageNum, _ := ctx.GetQuery("pageNum")
	pageSize, _ := ctx.GetQuery("pageSize")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	page := &model.PageType{
		PageNum:  num,
		PageSize: size,
	}
	todo, total, err := assetLogic.GetRevertTodo(ctx, page)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "", map[string]interface{}{
		"total": total,
		"todo":  todo,
	})
}

// GetRepairsTodo 获取维修待办
func (m *ManagementApi) GetRepairsTodo(ctx *gin.Context) {
	pageNum, _ := ctx.GetQuery("pageNum")
	pageSize, _ := ctx.GetQuery("pageSize")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	page := &model.PageType{
		PageNum:  num,
		PageSize: size,
	}
	todo, total, err := assetLogic.GetRepairsTodo(ctx, page)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "", map[string]interface{}{
		"total": total,
		"todo":  todo,
	})
}

// GetAuditTodo 获取审批待办
func (m *ManagementApi) GetAuditTodo(ctx *gin.Context) {
	pageNum, _ := ctx.GetQuery("pageNum")
	pageSize, _ := ctx.GetQuery("pageSize")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	page := &model.PageType{
		PageNum:  num,
		PageSize: size,
	}
	todo, total, err := assetLogic.GetAuditTodo(ctx, page)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "", map[string]interface{}{
		"total": total,
		"todo":  todo,
	})
}

// GetProviderTodo 获取发放待办
func (m *ManagementApi) GetProviderTodo(ctx *gin.Context) {
	pageNum, _ := ctx.GetQuery("pageNum")
	pageSize, _ := ctx.GetQuery("pageSize")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	page := &model.PageType{
		PageNum:  num,
		PageSize: size,
	}
	todo, total, err := assetLogic.GetProviderTodo(ctx, page)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "", map[string]interface{}{
		"total": total,
		"todo":  todo,
	})
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

// GetTodoDetails 获取某个待办的详情信息
func (m *ManagementApi) GetTodoDetails(ctx *gin.Context) {
	query, _ := ctx.GetQuery("id")
	id, _ := strconv.ParseInt(query, 10, 64)
	kind, _ := ctx.GetQuery("kind")
	if id <= 0 || kind == "" {
		global.FailWithMsg(ctx, "获取失败")
		return
	}
	res, err := assetLogic.GetTodoDetails(ctx, id, kind)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "获取成功", res)

}

// GetAssetsByUser 获取用户所拥有的资产
func (m *ManagementApi) GetAssetsByUser(ctx *gin.Context) {
	pageNum, _ := ctx.GetQuery("pageNum")
	pageSize, _ := ctx.GetQuery("pageSize")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	var assets []*model.UserAssets
	var total int64
	if err = global.GLOBAL_DB.Table("user_assets").Where("user_id = ?", claims.UserId).Count(&total).Limit(size).Offset((num - 1) * size).Find(&assets).Error; err != nil {
		global.GLOBAL_LOG.Error("获取用户资产信息失败", zap.String("user_id", claims.UserId), zap.Error(err))
		global.FailWithMsg(ctx, "获取资产失败")
		return
	}
	res := make([]*response.AssetsInfo, 0, 8)
	for _, asset := range assets {
		tmp := &response.AssetsInfo{
			SerialId:   asset.SerialId,
			ExpireTime: asset.ExpireTime.Unix(),
		}
		revert := model.AssetRevert{}
		if err = global.GLOBAL_DB.Table("asset_details").Where("serial_id = ?", asset.SerialId).Select("status").Find(&revert).Error; err != nil {
			continue
		}
		tmp.Status = revert.Status
		res = append(res, tmp)
	}
	global.OkWithDetails(ctx, "", map[string]interface{}{
		"total":  total,
		"assets": res,
	})
}

// ProvideAsset 发放资产
func (m *ManagementApi) ProvideAsset(ctx *gin.Context) {
	var info *request.ReceiveForm
	if err := ctx.ShouldBind(&info); err != nil {
		global.GLOBAL_LOG.Error("提交的信息有误", zap.Error(err))
		global.FailWithMsg(ctx, "提交的信息有误")
		return
	}
	if err := assetLogic.ProvideAsset(ctx, info); err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithMsg(ctx, "发放成功")
}

// GetAssetsByCategory 通过品类获取资产信息
func (m *ManagementApi) GetAssetsByCategory(ctx *gin.Context) {
	category, _ := ctx.GetQuery("category")
	if category == "" {
		global.FailWithMsg(ctx, "无效的种类")
		return
	}
	var info []*model.AssetDetails
	if err := global.GLOBAL_DB.Table("asset_details").Where("category = ? and status = ?", category, global.CanApply).Find(&info).Error; err != nil {
		global.GLOBAL_LOG.Error("获取资产列表失败", zap.String("category", category), zap.String("status", global.StatusMap[global.CanApply]), zap.Error(err))
		global.FailWithMsg(ctx, "获取资产失败")
		return
	}
	var assets []*response.AssetVO
	for _, asset := range info {
		tmp := &response.AssetVO{
			SerialId: asset.SerialId,
			Category: asset.Category,
			Name:     asset.Name,
		}
		assets = append(assets, tmp)
	}
	global.OkWithDetails(ctx, "", assets)
}

// RevertAssets 归还资产
func (m *ManagementApi) RevertAssets(ctx *gin.Context) {
	query, _ := ctx.GetQuery("id")
	id, _ := strconv.ParseInt(query, 10, 64)
	if id <= 0 {
		global.FailWithMsg(ctx, "订单不存在")
		return
	}
	if err := assetLogic.RevertAssets(ctx, id); err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithMsg(ctx, "归还成功")
}

// RejectReceive 驳回申请
func (m *ManagementApi) RejectReceive(ctx *gin.Context) {
	query, _ := ctx.GetQuery("id")
	id, _ := strconv.ParseInt(query, 10, 64)
	if id <= 0 {
		global.FailWithMsg(ctx, "订单不存在")
		return
	}
	if err := global.GLOBAL_DB.Table("asset_receive").Where("id = ?", id).Updates(map[string]interface{}{
		"status":      global.Rejected,
		"update_time": time.Now(),
	}).Error; err != nil {
		global.GLOBAL_LOG.Error("更改申请单状态失败", zap.Int64("id", id), zap.Error(err))
		global.FailWithMsg(ctx, "驳回失败")
		return
	}
	global.OkWithMsg(ctx, "驳回成功")
}

// GetAllAssets 获取所有资产信息
func (m *ManagementApi) GetAllAssets(ctx *gin.Context) {
	pageNum, _ := ctx.GetQuery("pageNum")
	pageSize, _ := ctx.GetQuery("pageSize")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	page := &model.PageType{
		PageNum:  num,
		PageSize: size,
	}
	res, total, err := assetLogic.GetAllAssets(ctx, page)
	if err != nil {
		global.FailWithMsg(ctx, "获取失败")
		return
	}
	global.OkWithDetails(ctx, "", map[string]interface{}{
		"assets": res,
		"total":  total,
	})
}

// GetAssetUser 获取资产领用人信息
func (m *ManagementApi) GetAssetUser(ctx *gin.Context) {
	serialId, _ := ctx.GetQuery("serial_id")
	if serialId == "" {
		global.FailWithMsg(ctx, "资产不存在")
		return
	}
	var userAsset *model.UserAssets
	if err := global.GLOBAL_DB.Table("user_assets").Where("serial_id = ?", serialId).Find(&userAsset).Error; err != nil {
		global.GLOBAL_LOG.Error("获取用户资产记录失败", zap.String("serial_id", serialId), zap.Error(err))
		global.FailWithMsg(ctx, "获取失败")
		return
	}
	var user *model.SysUser
	if err := global.GLOBAL_DB.Table("sys_user").Where("user_id = ?", userAsset.UserId).Find(&user).Error; err != nil {
		global.GLOBAL_LOG.Error("获取领用人信息失败", zap.String("user_id", userAsset.UserId), zap.Error(err))
		global.FailWithMsg(ctx, "获取失败")
		return
	}
	res := &response.ReceiveUser{
		UserId:      user.UserId,
		UserName:    user.UserName,
		UserPhone:   user.Phone,
		ProvideTime: userAsset.CreateTime.Unix(),
		ExpireTime:  userAsset.ExpireTime.Unix(),
	}
	global.OkWithDetails(ctx, "", res)
}
