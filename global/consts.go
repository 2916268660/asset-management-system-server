package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"server/model/response"
	"time"
)

// 初始化相关
var (
	GLOBAL_DB    *gorm.DB      //数据库
	GLOBAL_CACHE *redis.Client //缓存
	GLOBAL_LOG   *zap.Logger   // zap日志库
)

// 登录方式相关
const (
	WayByUserName = iota + 1 //通过用户名登录
	WayByEmail               //通过邮箱登录
	WayByPhone               //通过手机号登录
)

// jwt相关
const (
	TokenExpireDuration = time.Hour * 12
	UserId              = "userId"
)

var MySecret = []byte("a8x0sd.")

// 用户权限
const (
	User     = "user"
	Charger  = "charger"
	Provider = "provider"
)

// 表单种类
const (
	Receive = "领用单"
	Repairs = "维修单"
	Revert  = "归还单"
)

// 申请领用状态
const (
	WaitAudit     = iota + 1 // 审批中
	WaitProvide              // 待发放
	Rejected                 // 已驳回
	ProvideFailed            // 发放失败
	ProvideDone              // 发放成功

	Reverting // 归还中
	Reverted  // 已归还

	Maintaining // 维护中

	WaitReceive // 待接单
	WaitRepair  // 待维修
	RepairDone  // 维修完成

	CanApply // 可领用
	Applied  // 已领用

	Rollback // 已撤销
)

var StatusMap = map[int]string{
	WaitAudit:     "审批中",
	WaitProvide:   "待发放",
	Rejected:      "已驳回",
	ProvideFailed: "发放失败",
	ProvideDone:   "发放成功",

	Reverting: "归还中",
	Reverted:  "已归还",

	Maintaining: "维护中",

	WaitReceive: "待接单",
	WaitRepair:  "待维修",
	RepairDone:  "维修完成",

	CanApply: "可领用",
	Applied:  "已领用",

	Rollback: "已撤销",
}

// RoleMenusMap 菜单列表
var RoleMenusMap = map[string][]*response.Menu{
	User: []*response.Menu{
		&response.Menu{
			ID:       1,
			AuthName: "首页",
			Path:     "/home",
		},
		&response.Menu{
			ID:       2,
			AuthName: "我的待办",
			Path:     "",
			Children: []*response.Menu{
				&response.Menu{
					ID:       3,
					AuthName: "领用待办",
					Path:     "/receiveTodo",
				},
				&response.Menu{
					ID:       4,
					AuthName: "归还待办",
					Path:     "/revertTodo",
				},
				&response.Menu{
					ID:       5,
					AuthName: "维修待办",
					Path:     "/repairsTodo",
				},
			},
		},
	},
	Charger: []*response.Menu{
		&response.Menu{
			ID:       1,
			AuthName: "首页",
			Path:     "/home",
		},
		&response.Menu{
			ID:       2,
			AuthName: "我的待办",
			Path:     "",
			Children: []*response.Menu{
				&response.Menu{
					ID:       6,
					AuthName: "审批待办",
					Path:     "/auditTodo",
				},
				&response.Menu{
					ID:       3,
					AuthName: "领用待办",
					Path:     "/receiveTodo",
				},
				&response.Menu{
					ID:       4,
					AuthName: "归还待办",
					Path:     "/revertTodo",
				},
				&response.Menu{
					ID:       5,
					AuthName: "维修待办",
					Path:     "/repairsTodo",
				},
			},
		},
	},
	Provider: []*response.Menu{
		&response.Menu{
			ID:       1,
			AuthName: "首页",
			Path:     "/home",
		},
		&response.Menu{
			ID:       2,
			AuthName: "我的待办",
			Path:     "/todo",
		},
		&response.Menu{
			ID:       7,
			AuthName: "用户管理",
			Path:     "",
			Children: []*response.Menu{
				&response.Menu{
					ID:       8,
					AuthName: "权限管理",
					Path:     "/auth",
				},
				&response.Menu{
					ID:       9,
					AuthName: "账号管理",
					Path:     "/user",
				},
			},
		},
		&response.Menu{
			ID:       10,
			AuthName: "资产管理",
			Path:     "/asset",
		},
	},
}
