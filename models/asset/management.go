package asset

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/global"
	"server/models/common"
)

type ManagementModel struct {
}

func (m *ManagementModel) CreateTask(ctx *gin.Context, task *common.Task) (taskId int64, err error) {
	if err = global.GLOBAL_DB.Create(task).Error; err != nil {
		log.Println(fmt.Sprintf("userId create task failed. userId=%s||err=%v", task.UserId, err))
		return 0, global.ERRDATABASE
	}
	return task.ID, nil
}

func (m *ManagementModel) CreateRepair(ctx *gin.Context, repair *common.Repairs) (repairId int64, err error) {
	if err = global.GLOBAL_DB.Create(repair).Error; err != nil {
		log.Println(fmt.Sprintf("userId create repair failed. userId=%s||err=%v", repair.UserId, err))
		return 0, global.ERRDATABASE
	}
	return repair.ID, nil
}
