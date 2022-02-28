package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

// 表单种类
const (
	Receive = "领用单"
	Repairs = "维修单"
	Revert  = "归还单"
)
