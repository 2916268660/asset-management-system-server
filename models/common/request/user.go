package request

// RegisterUserInfo 用户注册信息
type RegisterUserInfo struct {
	UserName     string `json:"username" binding:"required"`                    //用户名 唯一
	Password     string `json:"password" binding:"required"`                    //密码
	RePassword   string `json:"rePassword" binding:"required,eqfield=Password"` //第二次输入密码
	EmailOrPhone string `json:"emailOrPhone" binding:"required"`                //邮箱或者电话
	//ValidateCode   string `json:"validateCode" binding:"required"`   //验证码
}

// LoginUserInfoForUserName 用户登录
type LoginUserInfo struct {
	UserName     string `json:"username"`     //用户名 唯一
	Password     string `json:"password"`     //密码
	EmailOrPhone string `json:"emailOrPhone"` //邮箱或者电话
	Way          int    `json:"way"`          //登录方式  1：账户密码  2：邮箱验证码  3：电话验证码
	//ValidateCode   string `json:"validateCode" binding:"required"`   //验证码
}
