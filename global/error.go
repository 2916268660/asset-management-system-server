package global

import "strings"

// Error 错误结构
type Error struct {
	ErrNo  int
	ErrMsg string
}

func (e *Error) Error() string {
	return e.ErrMsg
}

func (e *Error) String() string {
	return e.Error()
}

func (e *Error) WithMsg(msg string) *Error {
	err := *e
	err.ErrMsg += msg
	return &err
}

func NewErr(errNo int, errMsg ...string) *Error {
	err := &Error{errNo, ErrMsg[errNo]}
	if len(errMsg) > 0 {
		err.ErrMsg += "(" + strings.Join(errMsg, "; ") + ")"
	}
	return err
}

const (
	OK       = 2000 //成功
	ARGS_ERR = 2001 //参数有误

	REGISTER_ERR = 40000 //注册用户失败

	UNKNOWN_ERR  = 5000  //未知错误
	DATABASE_ERR = 50001 //数据库出错,请联系管理员
)

var ErrMsg = map[int]string{
	OK:           "success",
	UNKNOWN_ERR:  "未知错误",
	ARGS_ERR:     "参数错误",
	REGISTER_ERR: "注册用户失败",
	DATABASE_ERR: "数据库出错,请联系管理员",
}

var (
	SUCCESS = NewErr(OK)

	ERRARGS = NewErr(ARGS_ERR)

	ERRREGISTER = NewErr(REGISTER_ERR)

	ERRUNKNOWN  = NewErr(UNKNOWN_ERR)
	ERRDATABASE = NewErr(DATABASE_ERR)
)
