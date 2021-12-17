package common

import "time"

// user表映射结构体
type User struct {
	ID         int64     // 主键ID
	UserName   string    // 用户名
	StuId      string    //学号
	Password   string    // 用户密码
	Email      string    //邮箱
	Phone      string    //电话
	CreateTime time.Time //创建时间
	UpdateTime time.Time //更新时间
}
