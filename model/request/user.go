package request

// RegisterUserInfo 用户注册信息
type RegisterUserInfo struct {
	UserName   string `json:"userName" binding:"required"`                    //姓名
	UserId     string `json:"userId" binding:"required"`                      //学号 唯一
	Password   string `json:"password" binding:"required,min=6,max=16"`       //密码
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"` //第二次输入密码
	Department string `json:"department" binding:"required"`                  //所属部门
	Email      string `json:"email" binding:"required,email"`                 //邮箱
	Phone      string `json:"phone" binding:"required"`                       //电话
	//ValidateCode   string `json:"validateCode" binding:"required"`   //验证码
}

// LoginUserInfo 用户登录
type LoginUserInfo struct {
	UserId   string `json:"userId" binding:"required"`                //用户名 唯一
	Password string `json:"password" binding:"required,min=3,max=16"` //密码
	Email    string `json:"email"`                                    //邮箱
	Phone    string `json:"phone"`                                    //电话
	Way      int    `json:"way" binding:"required"`                   //登录方式  1：账户密码  2：邮箱验证码  3：电话验证码
	//ValidateCode   string `json:"validateCode" binding:"required"`   //验证码
}

type ValidateCodeWay struct {
	Target string `json:"target" binding:"required"` // 发送验证码目标
	Way    int    `json:"way" binding:"required"`    // 发送方式  2、邮箱  3、手机
}

type UserRole struct {
	UserId     string `json:"userId" binding:"required"`
	Department string `json:"department" binding:"required"`
	Role       string `json:"role" binding:"required"`
}

type PasswordInfo struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}
