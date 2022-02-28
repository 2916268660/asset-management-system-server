package adapter

import (
	"github.com/gin-gonic/gin"
	"server/adapter/charger"
	"server/adapter/provider"
	"server/adapter/user"
	"server/global"
	"server/model/response"
)

type Factory interface {
	// Get提交的信息不合规, 请仔细检查后再次提交 获取待办事件
	GetTodo(ctx *gin.Context) (res []*response.Function, err error)
	// GetDone 获取已完成事件
	GetDone(ctx *gin.Context) (res []*response.Function, err error)
}

// NewFactory 简单工厂模式  不同角色实现不同的方法
func NewFactory(ctx *gin.Context, role string, userId string) Factory {
	switch role {
	case global.User:
		return &user.Adapter{UserId: userId}
	case global.Charger:
		return &charger.Adapter{UserId: userId}
	case global.Provider:
		return &provider.Adapter{UserId: userId}
	}
	return nil
}
