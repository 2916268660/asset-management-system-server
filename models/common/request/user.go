package request

// RegisterUserInfo 用户注册信息
type RegisterUserInfo struct {
	UserName   string `json:"userName"`   //姓名
	UserId     string `json:"userId"`     //学号 唯一
	Password   string `json:"password"`   //密码
	RePassword string `json:"rePassword"` //第二次输入密码
	Department string `json:"department"` //所属部门
	Email      string `json:"email"`      //邮箱
	Phone      string `json:"phone"`      //电话
	//ValidateCode   string `json:"validateCode" binding:"required"`   //验证码
}

// LoginUserInfo 用户登录
type LoginUserInfo struct {
	UserId   string `json:"userId"`   //用户名 唯一
	Password string `json:"password"` //密码
	Email    string `json:"email"`    //邮箱
	Phone    string `json:"phone"`    //电话
	Way      int    `json:"way"`      //登录方式  1：账户密码  2：邮箱验证码  3：电话验证码
	//ValidateCode   string `json:"validateCode" binding:"required"`   //验证码
}

type ValidateCodeWay struct {
	Target string `json:"target"` // 发送验证码目标
	Way    int    `json:"way"`    // 发送方式  2、邮箱  3、手机
}
