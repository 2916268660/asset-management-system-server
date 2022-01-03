package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"time"
)

// 初始化相关
var (
	GLOBAL_DB    *gorm.DB      //数据库
	GLOBAL_CACHE *redis.Client //缓存
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

// 任务单相关
// 状态
const (
	WaitApprove   = iota + 1 //审批中
	WaitProvide              //发放中
	Reject                   //未通过
	ProvideFailed            //发放失败
	ProvideDone              //发放成功
)

// 属性
const (
	Receive = iota + 1 // 领取单
	Revert             // 归还单
)

// 维修单状态
const (
	WaitReceive = iota + 1 // 待接单
	WaitRepair             // 待维修
	Repaired               // 维修完成
)

// md5盐
const Solt = "kjh1k2"
