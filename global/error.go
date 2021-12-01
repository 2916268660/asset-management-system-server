package global

import "strings"

// Error 错误结构
type Error struct {
	ErrNo  int
	ErrMsg string
}

func (e *Error) Error ()string {
	return e.ErrMsg
}

func (e *Error) String() string {
	return e.Error()
}

func (e *Error) WithMsg(msg string) *Error {
	err := *e
	err.ErrMsg = msg
	return &err
}

func NewErr(errNo int, errMsg...string) *Error {
	err := &Error{errNo, ErrMsg[errNo]}
	if len(errMsg) > 0 {
		err.ErrMsg += "(" + strings.Join(errMsg, "; ") + ")"
	}
	return err
}

const (
	OK = 2000  //成功
	UNKNOWN_ERR = 5000 //未知错误
)

var ErrMsg = map[int]string{
	OK: "success",
	UNKNOWN_ERR : "未知错误",
}

var (
	SUCCESS = NewErr(OK)
	ERRUNKNOWN = NewErr(UNKNOWN_ERR)
)


