package repairs

import (
	"github.com/gin-gonic/gin"
)

type Adapter struct{}

func (a *Adapter) Add(ctx *gin.Context, info interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) Details(ctx *gin.Context, id int64) (error, interface{}) {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) UpdateStatus(ctx *gin.Context, id int64, status int) error {
	//TODO implement me
	panic("implement me")
}
