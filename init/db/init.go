package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
var globalDB *gorm.DB
func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/server?charset=utf8mb4&parseTime=True&loc=Local", "root", "wS970107.")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("mysql connect fail||err=%v", err))
		return
	}
	globalDB = db
}

func GetDb(ctx gin.Context) *gorm.DB {
	return globalDB.WithContext(ctx.Copy())
}
