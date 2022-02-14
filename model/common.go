package model

import (
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"time"
)

type MyClaims struct {
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Department string `json:"department"`
	Role       string `json:"role"`
	jwt.StandardClaims
}

// SysUser 用户表
type SysUser struct {
	ID         int64     // 主键ID
	UserId     string    // 账号
	UserName   string    // 用户名
	Password   string    // 用户密码
	Email      string    // 邮箱
	Phone      string    // 电话
	Department string    // 部门
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间
	Role       string    // 角色
}

// AssetDetails 资产信息表
type AssetDetails struct {
	ID         int64     `json:"-"`          // 主键ID
	SerialId   uuid.UUID `json:"serialId"`   // 资产序列号
	SerialImg  string    `json:"serialImg"`  // 资产序列号的二维码路径
	Category   string    `json:"category"`   // 资产品类
	Name       string    `json:"name"`       // 资产名称
	Status     int       `json:"status"`     // 资产状态
	Price      float64   `json:"price"`      // 价格
	Provide    string    `json:"provide"`    // 采购地
	CreateTime time.Time `json:"createTime"` // 采购时间
	UpdateTime time.Time `json:"updateTime"` // 更新时间
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

// AssetReceive 领取单表
type AssetReceive struct {
	ID            int64     // 主键ID
	UserId        string    // 申请人账号
	UserName      string    // 申请人姓名
	UserPhone     string    // 申请人联系方式
	Department    string    // 申请人所属部门
	Category      string    // 资产品类
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
	ExpireTime    time.Time // 到期时间
	ProvideTime   time.Time // 发放时间
	CreateTime    time.Time // 申请时间
	AuditTime     time.Time // 审批时间
	UpdateTime    time.Time // 更新时间
	RollbackTime  time.Time // 撤回时间
}

type AssetRevert struct {
	ID             int64     // 主键ID
	UserId         string    // 申请人账号
	UserName       string    // 申请人姓名
	UserPhone      string    // 申请人联系方式
	Department     string    // 申请人所属部门
	Category       string    // 资产品类
	Nums           int       // 申请资产数量
	Assets         string    // 资产的序列号json字符串
	ReclaimerId    string    // 同意领用管理员的账号
	ReclaimerName  string    // 同意领用管理员的姓名
	ReclaimerPhone string    // 同意领用管理员的联系方式
	SignPath       string    // 电子签名生成的图片的存贮地址
	Remake         string    // 备注信息
	Status         int       // 任务单状态
	CreateTime     time.Time // 申请时间
	RevertTime     time.Time // 归还时间
	UpdateTime     time.Time // 更新时间
	RollbackTime   time.Time // 撤回时间
}

// AssetRepairs 维修单表
type AssetRepairs struct {
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

// Charger 部门负责人表
type Charger struct {
	ID         int64  // 主键ID
	Department string // 部门名称
	UserId     string // 负责人的ID
}
