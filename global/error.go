package global

import "strings"

// Error 错误结构
type Error struct {
	ErrNo   int
	ErrMsg  string
	ErrData string
}

func (e *Error) Error() string {
	return e.ErrMsg
}

func (e *Error) String() string {
	return e.Error()
}

func (e *Error) WithData(data string) *Error {
	err := *e
	err.ErrData += data
	return &err
}

func NewErr(errNo int, errMsg ...string) *Error {
	err := &Error{ErrNo: errNo, ErrMsg: ErrMsg[errNo]}
	if len(errMsg) > 0 {
		err.ErrMsg += "(" + strings.Join(errMsg, "; ") + ")"
	}
	return err
}

const (
	OK       = 2000 //成功
	ARGS_ERR = 2001 //参数有误

	REGISTER_ERR = 4000 //注册用户失败
	USERNAME_ERR = 4001 //用户名不存在
	PASSWORD_ERR = 4002 //密码错误

	UNKNOWN_ERR  = 5000 //未知错误
	DATABASE_ERR = 5001 //数据库出错,请联系管理员
)

var ErrMsg = map[int]string{
	OK: "success",

	ARGS_ERR:     "参数错误",
	REGISTER_ERR: "注册用户失败",
	USERNAME_ERR: "用户名不存在",
	PASSWORD_ERR: "密码错误",

	UNKNOWN_ERR:  "未知错误",
	DATABASE_ERR: "数据库出错,请联系管理员",
}

var (
	SUCCESS = NewErr(OK)       // 成功
	ERRARGS = NewErr(ARGS_ERR) // 参数错误

	ERRREGISTER = NewErr(REGISTER_ERR) // 注册用户失败
	ERRUSERNAME = NewErr(USERNAME_ERR) // 用户名不存在
	ERRPASSWORD = NewErr(PASSWORD_ERR) // 密码错误

	ERRUNKNOWN  = NewErr(UNKNOWN_ERR)  // 未知错误
	ERRDATABASE = NewErr(DATABASE_ERR) // 数据库出错,请联系管理员
)
