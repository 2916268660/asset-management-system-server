package global

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS = 0
	ERROR   = 500
)

func result(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, ResponseData{
		code,
		msg,
		data,
	})
}

func Ok(ctx *gin.Context) {
	result(ctx, SUCCESS, "操作成功", nil)
}

func OkWithMsg(ctx *gin.Context, msg string) {
	result(ctx, SUCCESS, msg, nil)
}

func OkWithData(ctx *gin.Context, data interface{}) {
	result(ctx, SUCCESS, "操作成功", data)
}

func OkWithDetails(ctx *gin.Context, msg string, data interface{}) {
	result(ctx, SUCCESS, msg, data)
}

func Fail(ctx *gin.Context) {
	result(ctx, ERROR, "操作失败", nil)
}

func FailWithMsg(ctx *gin.Context, msg string) {
	result(ctx, ERROR, msg, nil)
}

func FailWithDetails(ctx *gin.Context, msg string, data interface{}) {
	result(ctx, ERROR, msg, data)
}
