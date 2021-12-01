package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code int
	Msg  string
	Data interface{}
}

func Response(ctx *gin.Context, obj interface{}, err error) {
	res := &ResponseData{}
	if err != nil {
		switch e := err.(type) {
		case *Error:
			res.Code = e.ErrNo
			res.Msg = e.ErrMsg
		default:
			res.Code = ERRUNKNOWN.ErrNo
			res.Msg = ERRUNKNOWN.ErrMsg
		}
	} else {
		res.Code = SUCCESS.ErrNo
		res.Msg = SUCCESS.ErrMsg
	}
	res.Data = obj
	ctx.JSON(http.StatusOK, res)
	return
}
