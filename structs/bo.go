package structs

// UserArgsForLogin 用户登录信息
type UserArgsForLogin struct {
	UserName     string `json:"username"`     // 用户名
	Password     string `json:"password"`     // 密码
	ValidateCode string `json:"validateCode"` //验证码
}

// UserArgsForRegister 用户注册信息
type UserArgsForRegister struct {
	UserName       string `json:"username"`       //用户名 唯一
	Password       string `json:"password"`       //密码
	PasswordRepeat string `json:"passwordRepeat"` //第二次输入密码
	EmailOrPhone   string `json:"emailOrPhone"`   //邮箱或者电话
	ValidateCode   string `json:"validateCode"`   //验证码
}
