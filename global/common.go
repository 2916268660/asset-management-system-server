package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ResponseData struct {
	Code int
	Msg  string
	Data interface{}
}

// Response 统一响应
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

// 解析validator中的错误，去除结构体前缀
func ReMoveTopStruct(m map[string]string) map[string]string {
	res := make(map[string]string)
	for field, err := range m {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
