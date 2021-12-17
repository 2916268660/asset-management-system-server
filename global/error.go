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
	err.ErrMsg += "||" + msg
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

	REGISTER_ERR         = 4000 //注册用户失败
	USERNAMENOTEXIST_ERR = 4001 //用户名不存在
	USERNAMEISEXIST_ERR  = 4002 //用户已存在
	GETUSERINFO_ERR      = 4003 //获取用户信息失败
	PASSWORD_ERR         = 4004 //密码错误
	SENDCODE_ERR         = 4005 //发送验证码错误

	UNKNOWN_ERR       = 5000 //未知错误
	DATABASE_ERR      = 5001 //数据库出错,请联系管理员
	TOKENGENERATE_ERR = 5002 //生成token错误
	SENDEMAIL_ERR     = 5003 //发送邮箱错误
	CACHE_ERR         = 5004 //缓存出错
	TOKENNONE_ERR     = 5005 //无权限访问，请先进行登录
	TOKENTIMEOUT_ERR  = 5006 //token失效，请重新登录
	TOKENEFMT_ERR     = 5007 //token格式有误
)

var ErrMsg = map[int]string{
	OK: "success",

	ARGS_ERR:             "参数错误",
	REGISTER_ERR:         "注册用户失败",
	USERNAMENOTEXIST_ERR: "用户名不存在",
	USERNAMEISEXIST_ERR:  "用户已存在",
	GETUSERINFO_ERR:      "获取用户信息失败",
	PASSWORD_ERR:         "密码错误",
	SENDCODE_ERR:         "发送验证码错误",

	UNKNOWN_ERR:       "未知错误",
	DATABASE_ERR:      "数据库出错,请联系管理员",
	TOKENGENERATE_ERR: "token生成错误",
	SENDEMAIL_ERR:     "发送邮箱错误",
	CACHE_ERR:         "缓存错误",
	TOKENNONE_ERR:     "无权限访问，请先进行登录",
	TOKENTIMEOUT_ERR:  "token失效，请重新登录",
	TOKENEFMT_ERR:     "token格式有误",
}

var (
	SUCCESS = NewErr(OK)       // 成功
	ERRARGS = NewErr(ARGS_ERR) // 参数错误

	ERRREGISTER         = NewErr(REGISTER_ERR)         // 注册用户失败
	ERRUSERNAMENOTEXIST = NewErr(USERNAMENOTEXIST_ERR) // 用户名不存在
	ERRUSERNAMEISEXIST  = NewErr(USERNAMEISEXIST_ERR)  // 用户已存在
	ERRGETUSERINFO      = NewErr(GETUSERINFO_ERR)      // 获取用户信息失败
	ERRPASSWORD         = NewErr(PASSWORD_ERR)         // 密码错误
	ERRSENDCODE         = NewErr(SENDCODE_ERR)         //发送验证码错误

	ERRUNKNOWN       = NewErr(UNKNOWN_ERR)       // 未知错误
	ERRDATABASE      = NewErr(DATABASE_ERR)      // 数据库出错,请联系管理员
	ERRTOKENGENERATE = NewErr(TOKENGENERATE_ERR) // 生成token错误
	ERRSENDEMAIL     = NewErr(SENDEMAIL_ERR)     // 发送邮箱错误
	ERRCACHE         = NewErr(CACHE_ERR)         //缓存错误
	ERRTOKENNONE     = NewErr(TOKENNONE_ERR)     //无权限访问，请先进行登录
	ERRTOKENTIMEOUT  = NewErr(TOKENTIMEOUT_ERR)  //token失效，请重新登录
	ERRTOKENFMT      = NewErr(TOKENEFMT_ERR)     //token格式有误
)
