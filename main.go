package main

import (
	"fmt"
	"log"
	"server/global"
	"server/initialize"
)

func main() {
	if err := global.InitTrans("zh"); err != nil {
		log.Println(fmt.Sprintf("init trans failed, err=%v\n", err))
	}
	// 初始化mysql
	global.GLOBAL_DB = initialize.InitDB()
	// 初始化redis
	if err := initialize.InitCache(); err != nil {
		return
	}
	r := initialize.InitRouters()
	r.Run()
}
