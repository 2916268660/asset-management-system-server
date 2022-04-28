package asset

import (
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/model"
	"server/model/request"
	"server/model/response"
	"sync"
	"time"

	"server/global"
	"server/utils"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type ManagementLogic struct {
}

// ApplyReceive 申请领用资产
func (m *ManagementLogic) ApplyReceive(ctx *gin.Context, applyInfo *request.ApplyReceiveForm) (taskId int64, err error) {
	// 获取当前用户信息
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		global.GLOBAL_LOG.Error("获取当前用户信息失败", zap.Error(err))
		return 0, err
	}
	// 获取部门负责人信息
	charger, err := userModel.GetChargerByDepart(ctx, claims.Department)
	if err != nil {
		global.GLOBAL_LOG.Error("获取部门负责人信息失败", zap.Error(err))
		return 0, errors.New("获取部门负责人信息失败")
	}
	now := time.Now()
	task := &model.AssetReceive{
		UserId:       claims.UserId,
		UserName:     claims.UserName,
		UserPhone:    claims.Phone,
		AdminId:      charger.UserId,
		AdminName:    charger.UserName,
		AdminPhone:   charger.Phone,
		Department:   claims.Department,
		Category:     applyInfo.Category,
		Nums:         applyInfo.Nums,
		Days:         applyInfo.Days,
		Remake:       applyInfo.Remake,
		Status:       global.WaitAudit,
		ExpireTime:   time.Unix(0, 0),
		CreateTime:   now,
		UpdateTime:   now,
		ProvideTime:  time.Unix(0, 0),
		AuditTime:    time.Unix(0, 0),
		RollbackTime: time.Unix(0, 0),
	}

	// 发送申请
	taskId, err = assetModel.CreateReceive(ctx, task)
	if err != nil {
		global.GLOBAL_LOG.Error("新增申请记录失败", zap.Error(err))
		return 0, errors.New("发起申请失败")
	}
	return
}

// ApplyRevert 申请归还资产
func (m *ManagementLogic) ApplyRevert(ctx *gin.Context, applyInfo *request.ApplyRevertForm) (taskId int64, err error) {
	// 获取当前用户信息
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		global.GLOBAL_LOG.Error("获取当前用户信息失败", zap.Error(err))
		return 0, err
	}
	now := time.Now()
	assets, err := json.Marshal(applyInfo.Assets)
	if err != nil {
		global.GLOBAL_LOG.Error("json 转化失败", zap.Error(err))
		return 0, err
	}
	task := &model.AssetRevert{
		UserId:       claims.UserId,
		UserName:     claims.UserName,
		UserPhone:    claims.Phone,
		Department:   claims.Department,
		Nums:         len(applyInfo.Assets),
		Assets:       string(assets),
		Remake:       applyInfo.Remake,
		Status:       global.Reverting,
		CreateTime:   now,
		RevertTime:   time.Unix(0, 0),
		UpdateTime:   now,
		RollbackTime: time.Unix(0, 0),
	}
	taskId, err = assetModel.CreateRevert(ctx, task)
	if err != nil {
		return 0, errors.New("发起申请失败")
	}
	return
}

// ApplyRepair 申请维修
func (m *ManagementLogic) ApplyRepair(ctx *gin.Context, applyInfo *request.ApplyRepairForm) (repairId int64, err error) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		global.GLOBAL_LOG.Error("获取当前用户信息失败", zap.Error(err))
		return 0, err
	}
	assets, err := json.Marshal(applyInfo.Assets)
	if err != nil {
		return 0, err
	}
	now := time.Now()
	repair := &model.AssetRepairs{
		UserId:       claims.UserId,
		UserName:     claims.UserName,
		UserPhone:    claims.Phone,
		Address:      applyInfo.Address,
		Assets:       string(assets),
		Remake:       applyInfo.Remake,
		Status:       global.WaitReceive,
		CreateTime:   now,
		UpdateTime:   time.Unix(0, 0),
		ReceiveTime:  time.Unix(0, 0),
		RepairedTime: time.Unix(0, 0),
		RollbackTime: time.Unix(0, 0),
	}
	repairId, err = assetModel.CreateRepair(ctx, repair)
	if err != nil {
		global.GLOBAL_LOG.Error("新增维修记录失败", zap.Any("repair", repair), zap.Error(err))
		return 0, errors.New("发起申请失败")
	}
	return
}

// AddAssets 添加资产
func (m *ManagementLogic) AddAssets(ctx *gin.Context, info *request.AssetInfo) error {
	assets := make([]*model.AssetDetails, 0, info.Nums)
	wg, lock := sync.WaitGroup{}, sync.Mutex{}
	wg.Add(info.Nums)
	for i := 0; i < info.Nums; i++ {
		go func() {
			defer wg.Done()
			// 生成uuid来作为每台资产的序列号
			// 通过序列号生成对应的二维码，二维码则是对应的资产信息
			serialId := uuid.NewV4()
			path := utils.GetQRCode(serialId.String())
			now := time.Now()
			asset := &model.AssetDetails{
				SerialId:   serialId.String(),
				SerialImg:  path,
				Category:   info.Category,
				Name:       info.Name,
				Status:     global.CanApply,
				Price:      info.Price,
				Provide:    info.Provide,
				CreateTime: now,
				UpdateTime: now,
			}
			lock.Lock()
			defer lock.Unlock()
			assets = append(assets, asset)
		}()
	}
	wg.Wait()
	err := assetModel.CreateAssets(ctx, assets)
	if err != nil {
		global.GLOBAL_LOG.Error("新增资产失败", zap.Error(err))
		return errors.New("新增资产失败")
	}
	return nil
}

// UpdateAsset 更新资产信息
func (m *ManagementLogic) UpdateAsset(ctx *gin.Context, info *request.UpdateAssetInfo) (asset *model.AssetDetails, err error) {
	if err := global.GLOBAL_DB.Where("serial_id = ?", info.SerialId).First(&asset).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("资产不存在")
		}
	}
	if err := global.GLOBAL_DB.Model(&asset).Select("category", "name", "price", "provide", "update_time").Where("serial_id = ?", info.SerialId).Updates(
		map[string]interface{}{
			"category":    info.Category,
			"name":        info.Name,
			"price":       info.Price,
			"provide":     info.Provide,
			"update_time": time.Now(),
		}).Error; err != nil {
		global.GLOBAL_LOG.Error("更新资产信息失败", zap.String("serial_id", info.SerialId), zap.Error(err))
		return nil, errors.New("更新失败")
	}
	return asset, nil
}

// GetReceiveTodo 获取领用待办
func (m *ManagementLogic) GetReceiveTodo(ctx *gin.Context, page *model.PageType) (res []*response.Function, total int64, err error) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		global.GLOBAL_LOG.Error("获取当前用户信息失败", zap.Error(err))
		return nil, 0, err
	}
	receives, total, err := assetModel.GetReceiveByUser(ctx, claims.UserId, []int{global.WaitAudit, global.WaitProvide}, page)
	if err != nil {
		global.GLOBAL_LOG.Error("获取领用待办失败", zap.String("user_id", claims.UserId), zap.Error(err))
		return nil, 0, errors.New("获取领用待办失败, 请稍后重试")
	}
	for _, task := range receives {
		tmp := &response.Function{
			ID:         task.ID,
			CreateTime: task.CreateTime.Unix(),
			Kind:       global.Receive,
		}
		tmp.Status = task.Status
		res = append(res, tmp)
	}
	return
}

// GetRevertTodo 获取归还待办
func (m *ManagementLogic) GetRevertTodo(ctx *gin.Context, page *model.PageType) (res []*response.Function, total int64, err error) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		return nil, 0, err
	}
	reverts, total, err := assetModel.GetRevertByUser(ctx, claims.UserId, global.Reverting, page)
	if err != nil {
		global.GLOBAL_LOG.Error("获取归还待办失败", zap.String("user_id", claims.UserId), zap.Error(err))
		return nil, 0, errors.New("获取归还待办失败, 请稍后重试")
	}
	for _, task := range reverts {
		tmp := &response.Function{
			ID:         task.ID,
			CreateTime: task.CreateTime.Unix(),
			Kind:       global.Revert,
		}
		tmp.Status = task.Status
		res = append(res, tmp)
	}
	return
}

// GetRepairsTodo 获取维修待办
func (m *ManagementLogic) GetRepairsTodo(ctx *gin.Context, page *model.PageType) (res []*response.Function, total int64, err error) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		return nil, 0, err
	}
	repairs, total, err := assetModel.GetRepairsByUser(ctx, claims.UserId, global.WaitReceive, page)
	if err != nil {
		global.GLOBAL_LOG.Error("获取维修待办失败", zap.String("user_id", claims.UserId), zap.Error(err))
		return nil, 0, errors.New("获取维修待办失败, 请稍后重试")
	}
	for _, task := range repairs {
		tmp := &response.Function{
			ID:         task.ID,
			CreateTime: task.CreateTime.Unix(),
			Kind:       global.Repairs,
		}
		tmp.Status = task.Status
		res = append(res, tmp)
	}
	return
}

// GetAuditTodo 获取审批待办
func (m *ManagementLogic) GetAuditTodo(ctx *gin.Context, page *model.PageType) (res []*response.Function, total int64, err error) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		return nil, 0, err
	}
	reverts, total, err := assetModel.GetWaitAudit(ctx, claims.UserId, global.WaitAudit, page)
	if err != nil {
		global.GLOBAL_LOG.Error("获取审批待办失败", zap.String("user_id", claims.UserId), zap.Error(err))
		return nil, 0, errors.New("获取审批待办失败, 请稍后重试")
	}
	for _, task := range reverts {
		tmp := &response.Function{
			ID:         task.ID,
			CreateTime: task.CreateTime.Unix(),
			Kind:       global.Receive,
		}
		tmp.Status = task.Status
		res = append(res, tmp)
	}
	return
}

// GetProviderTodo 获取发放人员的待办事项 //todo 解决多个发放人员对同一订单不同处理，需要做前置校验是否已经处理过此订单
func (m *ManagementLogic) GetProviderTodo(ctx *gin.Context, page *model.PageType) (res []*response.Function, total int64, err error) {
	// 获取归还订单
	var revertCount int64
	var revertRes []*model.AssetRevert
	if err = global.GLOBAL_DB.Table("asset_revert").Where("status = ?", global.Reverting).Count(&revertCount).Find(&revertRes).Error; err != nil {
		global.GLOBAL_LOG.Error("获取归还中订单列表失败", zap.Error(err))
		return nil, 0, errors.New("获取失败")
	}
	for _, task := range revertRes {
		tmp := &response.Function{
			ID:         task.ID,
			CreateTime: task.CreateTime.Unix(),
			Kind:       global.Revert,
		}
		tmp.Status = task.Status
		res = append(res, tmp)
	}
	// 获取申请领用订单
	var receiveCount int64
	var receiveRes []*model.AssetRevert
	if err = global.GLOBAL_DB.Table("asset_receive").Where("status = ?", global.WaitProvide).Count(&receiveCount).Find(&receiveRes).Error; err != nil {
		global.GLOBAL_LOG.Error("获取待发放订单列表失败", zap.Error(err))
		return nil, 0, errors.New("获取失败")
	}
	for _, task := range receiveRes {
		tmp := &response.Function{
			ID:         task.ID,
			CreateTime: task.CreateTime.Unix(),
			Kind:       global.Receive,
		}
		tmp.Status = task.Status
		res = append(res, tmp)
	}
	total = revertCount + receiveCount
	// 将切片分页
	start := (page.PageNum - 1) * page.PageSize
	end := start + page.PageSize
	if end > len(res) {
		end = len(res)
	}
	return res[start:end], total, nil
}

// Rollback 撤回申请 todo 如果单子处于完成状态 不可撤回申请
func (m *ManagementLogic) Rollback(ctx *gin.Context, info *request.RollbackInfo) error {
	tableName := ""
	if info.ID <= 0 {
		return errors.New("该申请不存在")
	}
	switch info.Category {
	case global.Receive:
		tableName = "asset_receive"
	case global.Revert:
		tableName = "asset_revert"
	case global.Repairs:
		tableName = "asset_repairs"
	default:
		return errors.New("非法操作")
	}
	now := time.Now()
	if err := global.GLOBAL_DB.Table(tableName).Select("status", "update_time", "rollback_time").
		Where("id=?", info.ID).
		Updates(map[string]interface{}{
			"status":        global.Rollback,
			"update_time":   now,
			"rollback_time": now,
		}).Error; err != nil {
		global.GLOBAL_LOG.Error("更新撤回状态失败", zap.Int64("id", info.ID), zap.String("category", info.Category), zap.Error(err))
		return errors.New("撤回申请失败")
	}
	return nil
}

// GetTodoDetails 获取某个待办的详情信息
func (m *ManagementLogic) GetTodoDetails(ctx *gin.Context, id int64, kind string) (res interface{}, err error) {
	switch kind {
	case global.Receive:
		rece, err := assetModel.GetReceiveDetails(ctx, id)
		if err != nil {
			global.GLOBAL_LOG.Error("查询待办事件详情失败", zap.String("kind", kind), zap.Int64("id", id), zap.Error(err))
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("获取失败")
			}
			return nil, errors.New("待办事项不存在")
		}
		var assets []string
		if err = json.Unmarshal([]byte(rece.Assets), &assets); err != nil {
			global.GLOBAL_LOG.Error("资产反序列化失败", zap.Error(err))
			return nil, errors.New("操作失败")
		}
		receive := &response.AssetReceive{
			ID:            rece.ID,
			UserId:        rece.UserId,
			UserName:      rece.UserName,
			UserPhone:     rece.UserPhone,
			Department:    rece.Department,
			Category:      rece.Category,
			Nums:          rece.Nums,
			Days:          rece.Days,
			Assets:        assets,
			AdminId:       rece.AdminId,
			AdminName:     rece.AdminName,
			AdminPhone:    rece.AdminPhone,
			ProviderId:    rece.ProviderId,
			ProviderName:  rece.ProviderName,
			ProviderPhone: rece.ProviderPhone,
			Remake:        rece.Remake,
			Status:        rece.Status,
			ExpireTime:    rece.ExpireTime.Unix(),
			ProvideTime:   rece.ProvideTime.Unix(),
			CreateTime:    rece.CreateTime.Unix(),
			AuditTime:     rece.AuditTime.Unix(),
			UpdateTime:    rece.UpdateTime.Unix(),
			RollbackTime:  rece.RollbackTime.Unix(),
		}
		return receive, nil
	case global.Repairs:
		rep, err := assetModel.GetRepairsDetails(ctx, id)
		if err != nil {
			global.GLOBAL_LOG.Error("查询待办事件详情失败", zap.String("kind", kind), zap.Int64("id", id), zap.Error(err))
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("获取失败")
			}
			return nil, errors.New("待办事项不存在")
		}
		assets := make([]string, 0, 8)
		if err = json.Unmarshal([]byte(rep.Assets), &assets); err != nil {
			global.GLOBAL_LOG.Error("资产列表反序列化失败", zap.Error(err))
			return nil, err
		}
		repairs := &response.AssetRepairs{
			ID:            rep.ID,
			UserId:        rep.UserId,
			UserName:      rep.UserName,
			UserPhone:     rep.UserPhone,
			Address:       rep.Address,
			Assets:        assets,
			Remake:        rep.Remake,
			RepairerName:  rep.RepairerName,
			RepairerPhone: rep.RepairerPhone,
			Status:        rep.Status,
			CreateTime:    rep.CreateTime.Unix(),
			UpdateTime:    rep.UpdateTime.Unix(),
			ReceiveTime:   rep.ReceiveTime.Unix(),
			RepairedTime:  rep.RepairedTime.Unix(),
			RollbackTime:  rep.RollbackTime.Unix(),
		}
		return repairs, nil
	case global.Revert:
		rev, err := assetModel.GetRevertDetails(ctx, id)
		if err != nil {
			global.GLOBAL_LOG.Error("查询待办事件详情失败", zap.String("kind", kind), zap.Int64("id", id), zap.Error(err))
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("获取失败")
			}
			return nil, errors.New("待办事项不存在")
		}
		assets := make([]string, 0, 8)
		if err = json.Unmarshal([]byte(rev.Assets), &assets); err != nil {
			global.GLOBAL_LOG.Error("资产列表反序列化失败", zap.Error(err))
			return nil, err
		}
		revert := &response.AssetRevert{
			ID:             rev.ID,
			UserId:         rev.UserId,
			UserName:       rev.UserName,
			UserPhone:      rev.UserPhone,
			Department:     rev.Department,
			Nums:           rev.Nums,
			Assets:         assets,
			ReclaimerId:    rev.ReclaimerId,
			ReclaimerName:  rev.ReclaimerName,
			ReclaimerPhone: rev.ReclaimerPhone,
			Remake:         rev.Remake,
			Status:         rev.Status,
			CreateTime:     rev.CreateTime.Unix(),
			RevertTime:     rev.RevertTime.Unix(),
			UpdateTime:     rev.UpdateTime.Unix(),
		}
		return revert, nil
	default:
		return nil, errors.New("类型错误")
	}
	return
}

// ProvideAsset 发放资产
// 发放资产，先更改申请表单状态  再 更改资产表状态  再给用户资产表添加记录
func (m *ManagementLogic) ProvideAsset(ctx *gin.Context, receive *request.ReceiveForm) error {
	// 获取当前发放人员信息
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		global.GLOBAL_LOG.Error("获取当前用户信息失败", zap.Error(err))
		return err
	}
	now := time.Now()
	err = global.GLOBAL_DB.Transaction(func(tx *gorm.DB) error {
		// 先更改资产表状态
		if err = tx.Table("asset_details").Where("serial_id in ?", receive.Assets).Updates(
			map[string]interface{}{
				"status":      global.Applied,
				"update_time": now,
			}).Error; err != nil {
			global.GLOBAL_LOG.Error("更新资产表状态失败", zap.Any("serial_id", receive.Assets), zap.String("status", global.StatusMap[global.Applied]), zap.Error(err))
			return errors.New("更新资产表状态失败")
		}
		// 再更新申请表状态
		if err = tx.Table("asset_receive").Where("id = ?", receive.ID).Updates(map[string]interface{}{
			"status":         global.ProvideDone,
			"provider_id":    claims.UserId,
			"provider_name":  claims.UserName,
			"provider_phone": claims.Phone,
			"provide_time":   now,
			"update_time":    now,
			"expire_time":    now.AddDate(0, 0, receive.Days),
		}).Error; err != nil {
			global.GLOBAL_LOG.Error("更新申请表状态失败", zap.Int64("id", receive.ID), zap.Any("task", receive), zap.Error(err))
			return errors.New("更新申请表状态失败")
		}
		// 最后给用户资产表添加记录
		res := make([]*model.UserAssets, 0, 8)
		for _, asset := range receive.Assets {
			record := &model.UserAssets{
				SerialId:   asset,
				UserId:     receive.UserId,
				Status:     global.Applied,
				CreateTime: now,
				ExpireTime: now.AddDate(0, 0, receive.Days),
			}
			res = append(res, record)
		}
		if err = tx.Table("user_assets").Create(res).Error; err != nil {
			global.GLOBAL_LOG.Error("用户资产添加记录失败", zap.String("user_id", receive.UserId), zap.Any("assets", receive.Assets), zap.Error(err))
			return errors.New("用户资产添加记录失败")
		}
		return nil
	})
	if err != nil {
		global.GLOBAL_LOG.Error("发放资产失败", zap.Error(err))
		return errors.New("发放失败")
	}
	return nil
}

// RevertAssets 归还资产
func (m *ManagementLogic) RevertAssets(ctx *gin.Context, id int64) error {
	// 先更改订单状态
	var revert *model.AssetRevert
	if err := global.GLOBAL_DB.Table("asset_revert").Where("id = ?", id).Find(&revert).Update("status", global.Reverted).Error; err != nil {
		global.GLOBAL_LOG.Error("修改状态失败", zap.Int64("id", id), zap.Error(err))
		return errors.New("归还失败")
	}
	// 将资产的状态更改为可领用
	var assets []string
	if err := json.Unmarshal([]byte(revert.Assets), &assets); err != nil {
		global.GLOBAL_LOG.Error("资产列表反序列化失败", zap.Error(err))
		return errors.New("归还失败")
	}
	if err := global.GLOBAL_DB.Table("asset_details").Where("serial_id in ?", assets).Update("status", global.CanApply).Error; err != nil {
		global.GLOBAL_LOG.Error("修改资产状态失败", zap.Any("assets", assets), zap.Error(err))
		return errors.New("归还失败")
	}
	// 删除用户资产表的记录
	var userAssets *model.UserAssets
	if err := global.GLOBAL_DB.Table("user_assets").Where("serial_id in ?", assets).Delete(&userAssets).Error; err != nil {
		global.GLOBAL_LOG.Error("删除用户资产表的记录", zap.Error(err))
		return errors.New("归还失败")
	}
	return nil
}

// GetAllAssets 获取所有资产信息
func (m *ManagementLogic) GetAllAssets(ctx *gin.Context, page *model.PageType) (res []*response.Asset, total int64, err error) {
	var assets []*model.AssetDetails
	if err = global.GLOBAL_DB.Table("asset_details").Count(&total).Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize).Find(&assets).Error; err != nil {
		global.GLOBAL_LOG.Error("获取所有资产信息失败", zap.Error(err))
		return nil, 0, errors.New("获取失败")
	}
	for _, asset := range assets {
		tmp := &response.Asset{
			Name:     asset.Name,
			Category: asset.Category,
			SerialId: asset.SerialId,
			Status:   asset.Status,
		}
		res = append(res, tmp)
	}
	return
}
