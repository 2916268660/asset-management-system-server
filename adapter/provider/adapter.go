package provider

import (
	"github.com/gin-gonic/gin"
	"server/model/response"
)

type Adapter struct {
	UserId string
}

func (a *Adapter) GetTodo(ctx *gin.Context) (res []*response.Function, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) GetDone(ctx *gin.Context) (res []*response.Function, err error) {
	//TODO implement me
	panic("implement me")
}
