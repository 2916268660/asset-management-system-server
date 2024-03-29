package asset

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model"
)

type ManagementModel struct {
}

func (m *ManagementModel) CreateReceive(ctx *gin.Context, task *model.AssetReceive) (taskId int64, err error) {
	if err = global.GLOBAL_DB.Table("asset_receive").Create(task).Error; err != nil {
		return 0, err
	}
	return task.ID, nil
}

func (m *ManagementModel) CreateRevert(ctx *gin.Context, task *model.AssetRevert) (taskId int64, err error) {
	if err = global.GLOBAL_DB.Create(task).Error; err != nil {
		return 0, err
	}
	return task.ID, nil
}

func (m *ManagementModel) CreateRepair(ctx *gin.Context, repair *model.AssetRepairs) (repairId int64, err error) {
	if err = global.GLOBAL_DB.Create(repair).Error; err != nil {
		return 0, err
	}
	return repair.ID, nil
}

func (m *ManagementModel) CreateAssets(ctx *gin.Context, assets []*model.AssetDetails) error {
	if err := global.GLOBAL_DB.Create(assets).Error; err != nil {
		return err
	}
	return nil
}

func (m *ManagementModel) GetReceiveById(ctx *gin.Context, id int64) (receive *model.AssetReceive, err error) {
	if err = global.GLOBAL_DB.Where("id=?", id).First(&receive).Error; err != nil {
		return nil, err
	}
	return
}

// GetReceiveByUser 通过用户以及单的状态来获取指定申请领用单的ID
func (m *ManagementModel) GetReceiveByUser(ctx *gin.Context, userId string, status []int, page *model.PageType) (res []*model.AssetReceive, total int64, err error) {
	//if err = global.GLOBAL_DB.Raw("select * from asset_receive where user_id=? and status in ? limit ?,?", userId, status, page.PageSize, (page.PageNum-1)*page.PageSize).Scan(&res).Error; err != nil {
	//	return nil, 0, err
	//}
	//total = len(res)
	if err = global.GLOBAL_DB.Table("asset_receive").Where("user_id=? and status in ?", userId, status).Count(&total).Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return
}

func (m *ManagementModel) GetRevertByUser(ctx *gin.Context, userId string, status int, page *model.PageType) (res []*model.AssetRevert, total int64, err error) {
	if err = global.GLOBAL_DB.Table("asset_revert").Where("user_id=? and status=?", userId, status).Count(&total).Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return
}

func (m *ManagementModel) GetRepairsByUser(ctx *gin.Context, userId string, status int, page *model.PageType) (res []*model.AssetRepairs, total int64, err error) {
	if err = global.GLOBAL_DB.Table("asset_repairs").Where("user_id=? and status=?", userId, status).Count(&total).Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return
}

func (m *ManagementModel) GetWaitAudit(ctx *gin.Context, adminId string, status int, page *model.PageType) (res []*model.AssetReceive, total int64, err error) {
	if err = global.GLOBAL_DB.Table("asset_receive").Where("admin_id=? and status=?", adminId, status).Count(&total).Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return
}

func (m *ManagementModel) GetReceiveDetails(ctx *gin.Context, id int64) (res *model.AssetReceive, err error) {
	if err = global.GLOBAL_DB.Table("asset_receive").Where("id = ?", id).First(&res).Error; err != nil {
		return nil, err
	}
	return
}

func (m *ManagementModel) GetRepairsDetails(ctx *gin.Context, id int64) (res *model.AssetRepairs, err error) {
	if err = global.GLOBAL_DB.Table("asset_repairs").Where("id = ?", id).First(&res).Error; err != nil {
		return nil, err
	}
	return
}

func (m *ManagementModel) GetRevertDetails(ctx *gin.Context, id int64) (res *model.AssetRevert, err error) {
	if err = global.GLOBAL_DB.Table("asset_revert").Where("id = ?", id).First(&res).Error; err != nil {
		return nil, err
	}
	return
}
