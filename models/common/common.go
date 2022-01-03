package common

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type MyClaims struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Phone    string `json:"phone"`
	jwt.StandardClaims
}

// User 用户表
type User struct {
	ID         int64     // 主键ID
	UserId     string    // 账号
	UserName   string    // 用户名
	Password   string    // 用户密码
	Email      string    // 邮箱
	Phone      string    // 电话
	Department string    // 部门
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间
	isAdmin    int       // 是否为管理员
}

// AssetDetails 资产信息表
type AssetDetails struct {
	ID         int64     // 主键ID
	SerialId   string    // 资产序列号
	SerialImg  string    // 资产序列号的二维码路径
	Category   int       // 资产品类
	Name       string    // 资产名称
	Status     int       // 资产状态
	Price      int       // 价格（只允许存整数）
	Provide    string    // 采购地
	CreateTime time.Time // 采购时间
	UpdateTime time.Time // 更新时间
}

// AssetUseRecord 资产记录表
type AssetUseRecord struct {
	ID         int64     // 主键ID
	TaskId     int64     // 任务ID
	SerialId   string    // 资产序列号
	Status     int       // 资产状态
	CreateTime time.Time // 申请时间
	ExpireTime time.Time // 到期时间
}

// Task 领取单表
type Task struct {
	ID            int64     // 主键ID
	UserId        string    // 申请人账号
	UserName      string    // 申请人姓名
	UserPhone     string    // 申请人联系方式
	Department    string    // 申请人所属部门
	Category      int       // 资产品类
	Nums          int       // 申请资产数量
	Days          int       // 申请天数
	Assets        string    // 资产的序列号json字符串
	AdminId       string    // 同意领用管理员的账号
	AdminName     string    // 同意领用管理员的姓名
	AdminPhone    string    // 同意领用管理员的联系方式
	ProviderId    string    // 发放资产人的账号
	ProviderName  string    // 发放资产人的姓名
	ProviderPhone string    // 发放资产人的联系方式
	SignPath      string    // 电子签名生成的图片的存贮地址
	Remake        string    // 备注信息
	Status        int       // 任务单状态
	Property      int       // 任务单属性。 1：领用单  2：归还单
	ExpireTime    time.Time // 到期时间
	ProvideTime   time.Time // 发放时间
	CreateTime    time.Time // 申请时间
	AgreeTime     time.Time // 审批时间
	UpdateTime    time.Time // 更新时间
	RollbackTime  time.Time // 撤回时间
}

// Repairs 维修单表
type Repairs struct {
	ID            int64  // 主键ID
	UserId        string // 申请人账号
	UserName      string // 申请人姓名
	UserPhone     string // 申请人联系方式
	Address       string // 地址
	Assets        string // 资产的序列号json字符串
	Remake        string // 备注信息
	RepairerName  string // 维修人员的姓名
	RepairerPhone string // 维修人员的联系方式
	Status        int    // 维修单状态
	CreateTime    time.Time
	UpdateTime    time.Time
	ReceiveTime   time.Time // 接单时间
	RepairedTime  time.Time // 维修完成时间
	RollbackTime  time.Time
}
