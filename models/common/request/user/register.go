package user


// RegisterUserInfo 用户注册信息
type RegisterUserInfo struct {
	UserName     string `json:"username" binding:"required"`                    //用户名 唯一
	Password     string `json:"password" binding:"required"`                    //密码
	RePassword   string `json:"rePassword" binding:"required,eqfield=Password"` //第二次输入密码
	EmailOrPhone string `json:"emailOrPhone" binding:"required"`                //邮箱或者电话
	//ValidateCode   string `json:"validateCode" binding:"required"`   //验证码
}
