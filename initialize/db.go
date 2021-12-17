package initialize

import (
	"fmt"
	"github.com/go-ini/ini"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var dsn string

func init() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Println("加载配置文件失败", err)
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
		log.Println(fmt.Sprintf("mysql connect fail||err=%v", err))
		return nil
	}
	return db
}
