package charger

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model/request"
)

type ManagementApi struct {
}

// Audit 审批
func (m *ManagementApi) Audit(ctx *gin.Context) {
	var info *request.AuditStatus
	if err := ctx.ShouldBind(&info); err != nil {
		global.Fail(ctx)
		return
	}
	if info.ID <= 0 {
		global.FailWithMsg(ctx, "申请单不存在")
		return
	}
	if _, ok := global.StatusMap[info.Status]; !ok {
		global.Fail(ctx)
		return
	}
	err := chargerLogic.Audit(ctx, info)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.Ok(ctx)
}
