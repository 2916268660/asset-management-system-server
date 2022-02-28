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
func (m ManagementLogic) ApplyReceive(ctx *gin.Context, applyInfo *request.ApplyReceiveForm) (taskId int64, err error) {
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
		ExpireTime:   now.AddDate(0, 0, applyInfo.Days),
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
func (m *ManagementLogic) GetReceiveTodo(ctx *gin.Context) (res []*response.Function, err error) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		global.GLOBAL_LOG.Error("获取当前用户信息失败", zap.Error(err))
		return nil, err
	}
	receives, err := assetModel.GetReceiveByUser(ctx, claims.UserId, []int{global.WaitAudit, global.WaitProvide})
	if err != nil {
		global.GLOBAL_LOG.Error("获取领用待办失败", zap.String("user_id", claims.UserId), zap.Error(err))
		return nil, errors.New("获取领用待办失败, 请稍后重试")
	}
	for _, task := range receives {
		tmp := &response.Function{
			ID:        task.ID,
			CreatTime: task.CreateTime,
			Kind:      global.Receive,
		}
		status, ok := global.StatusMap[task.Status]
		if !ok {
			continue
		}
		tmp.Status = status
		res = append(res, tmp)
	}
	return
}

// GetRevertTodo 获取归还待办
func (m *ManagementLogic) GetRevertTodo(ctx *gin.Context) (res []*response.Function, err error) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	reverts, err := assetModel.GetRevertByUser(ctx, claims.UserId, global.Reverting)
	if err != nil {
		global.GLOBAL_LOG.Error("获取归还待办失败", zap.String("user_id", claims.UserId), zap.Error(err))
		return nil, errors.New("获取归还待办失败, 请稍后重试")
	}
	for _, task := range reverts {
		tmp := &response.Function{
			ID:        task.ID,
			CreatTime: task.CreateTime,
			Kind:      global.Revert,
		}
		status, ok := global.StatusMap[task.Status]
		if !ok {
			continue
		}
		tmp.Status = status
		res = append(res, tmp)
	}
	return
}

// GetRepairsTodo 获取维修待办
func (m *ManagementLogic) GetRepairsTodo(ctx *gin.Context) (res []*response.Function, err error) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	repairs, err := assetModel.GetRepairsByUser(ctx, claims.UserId, global.WaitReceive)
	if err != nil {
		global.GLOBAL_LOG.Error("获取维修待办失败", zap.String("user_id", claims.UserId), zap.Error(err))
		return nil, errors.New("获取维修待办失败, 请稍后重试")
	}
	for _, task := range repairs {
		tmp := &response.Function{
			ID:        task.ID,
			CreatTime: task.CreateTime,
			Kind:      global.Repairs,
		}
		status, ok := global.StatusMap[task.Status]
		if !ok {
			continue
		}
		tmp.Status = status
		res = append(res, tmp)
	}
	return
}

// GetAuditTodo 获取审批待办
func (m *ManagementLogic) GetAuditTodo(ctx *gin.Context) (res []*response.Function, err error) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		return nil, err
	}
	reverts, err := assetModel.GetWaitAudit(ctx, claims.UserId, global.WaitAudit)
	if err != nil {
		global.GLOBAL_LOG.Error("获取审批待办失败", zap.String("user_id", claims.UserId), zap.Error(err))
		return nil, errors.New("获取审批待办失败, 请稍后重试")
	}
	for _, task := range reverts {
		tmp := &response.Function{
			ID:        task.ID,
			CreatTime: task.CreateTime,
			Kind:      global.Receive,
		}
		status, ok := global.StatusMap[task.Status]
		if !ok {
			continue
		}
		tmp.Status = status
		res = append(res, tmp)
	}
	return
}

// Rollback 撤回申请
func (m *ManagementLogic) Rollback(ctx *gin.Context, info *request.RollbackInfo) error {
	if info.ID <= 0 {
		return errors.New("该申请不存在")
	}
	var obj interface{}
	switch info.Category {
	case global.Receive:
		obj = new(model.AssetReceive)
	case global.Revert:
		obj = new(model.AssetRevert)
	case global.Repairs:
		obj = new(model.AssetRepairs)
	default:
		return errors.New("非法操作")
	}

	now := time.Now()
	if err := global.GLOBAL_DB.Model(&obj).Select("status", "update_time", "rollback_time").
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
