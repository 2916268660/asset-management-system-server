package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var GlobalDB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/server?charset=utf8mb4&parseTime=True&loc=Local", "root", "wS970107.")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 严格按照模型映射，不自动给表加复数
		},
	})
	if err != nil {
		log.Println(fmt.Sprintf("mysql connect fail||err=%v", err))
		return
	}
	GlobalDB = db
}
