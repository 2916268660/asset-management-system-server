package charger

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model/response"
)

type Adapter struct {
	UserId string
}

func (a *Adapter) GetTodo(ctx *gin.Context) (res []*response.Function, err error) {
	// 获取申请领用待办
	todoReceiveStatus := []int{
		global.WaitAudit,
		global.WaitProvide,
	}
	receives, err := assetModel.GetReceiveByUser(ctx, a.UserId, todoReceiveStatus)
	if err != nil {
		return res, err
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
	// 获取申请归还待办
	reverts, err := assetModel.GetRevertByUser(ctx, a.UserId, global.Reverting)
	if err != nil {
		return res, err
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
	// 获取维修待办
	repairs, err := assetModel.GetRepairsByUser(ctx, a.UserId, global.WaitRepair)
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

func (a *Adapter) GetDone(ctx *gin.Context) (res []*response.Function, err error) {
	//TODO implement me
	panic("implement me")
}
