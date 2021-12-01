package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/structs"
)

type BaseContext struct {
	c *gin.Context
}

func (b *BaseContext) Response (obj interface{}, err error) {
	res := &structs.Response{}
	if err != nil {
		switch e := err.(type) {
		case *Error:
			res.ErrNo = e.ErrNo
			res.ErrMsg = e.ErrMsg
		default:
			res.ErrNo = ERRUNKNOWN.ErrNo
			res.ErrMsg = ERRUNKNOWN.ErrMsg
		}
	} else {
		res.ErrNo = SUCCESS.ErrNo
		res.ErrMsg = SUCCESS.ErrMsg
	}
	res.RespData = obj
	b.c.JSON(http.StatusOK, res)
	return
}