package adapter

import (
	"github.com/gin-gonic/gin"
	"server/adapter/receive"
	"server/adapter/repairs"
	"server/adapter/revert"
	"server/global"
)

type Adapter interface {
	// 新增记录
	Add(ctx *gin.Context, info interface{}) error
	// 详情
	Details(ctx *gin.Context, id int64) (error, interface{})
	// 更改状态
	UpdateStatus(ctx *gin.Context, id int64, status int) error
}

func NewAdapter(ctx *gin.Context, kind string) Adapter {
	switch kind {
	case global.Receive:
		return &receive.Adapter{}
	case global.Revert:
		return &revert.Adapter{}
	case global.Repairs:
		return &repairs.Adapter{}
	}
	return nil
}
