package asset

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"server/global"
	"server/models/common/request"
)

type ManagementApi struct {
}

// ApplyReceive 申请领用资产
func (m *ManagementApi) ApplyReceive(ctx *gin.Context) {
	var applyInfo *request.ApplyReceiveForm
	if err := ctx.ShouldBind(&applyInfo); err != nil {
		log.Println(fmt.Sprintf("submit msg has err err=%v", err))
		global.Response(ctx, nil, global.ERRARGS.WithMsg("提交的信息有误，请仔细检查后再次提交"))
		return
	}
	taskId, err := assetLogic.ApplyReceive(ctx, applyInfo)
	global.Response(ctx, map[string]int64{"taskId": taskId}, err)
}

// ApplyRevert 申请归还资产
func (m *ManagementApi) ApplyRevert(ctx *gin.Context) {
	var applyInfo *request.ApplyRevertForm
	if err := ctx.ShouldBind(&applyInfo); err != nil {
		log.Println(fmt.Sprintf("submit msg has err err=%v", err))
		global.Response(ctx, nil, global.ERRARGS.WithMsg("提交的信息有误，请仔细检查后再次提交"))
		return
	}
	taskId, err := assetLogic.ApplyRevert(ctx, applyInfo)
	global.Response(ctx, map[string]int64{"taskId": taskId}, err)
}

// ApplyRepair 申请维修资产
func (m *ManagementApi) ApplyRepair(ctx *gin.Context) {
	var applyInfo *request.ApplyRepairForm
	if err := ctx.ShouldBind(&applyInfo); err != nil {
		log.Println(fmt.Sprintf("submit msg has err err=%v", err))
		global.Response(ctx, nil, global.ERRARGS.WithMsg("提交的信息有误，请仔细检查后再次提交"))
		return
	}
	repairId, err := assetLogic.ApplyRepair(ctx, applyInfo)
	global.Response(ctx, map[string]int64{"repairId": repairId}, err)
}