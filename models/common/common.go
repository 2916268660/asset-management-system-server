package common

import "time"

type User struct {
	ID           int64     // 主键ID
	UserName     string    // 用户名
	Password     string    // 用户密码
	EmailOrPhone string    //邮箱或者电话
	CreateTime   time.Time //创建时间
	UpdateTime   time.Time //更新时间
}
