package core

import (
	"github.com/go-ini/ini"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"server/global"
)

var dsn string

func init() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		global.GLOBAL_LOG.Error("加载配置文件失败", zap.Error(err))
		return
	}
	mysqlCfg := cfg.Section("mysql")
	dsn = mysqlCfg.Key("dsn").MustString("")
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 严格按照模型映射，不自动给表加复数
		},
	})
	if err != nil {
		global.GLOBAL_LOG.Error("数据库连接失败", zap.Error(err))
		return nil
	}
	global.GLOBAL_LOG.Debug("数据库连接成功")
	return db
}
