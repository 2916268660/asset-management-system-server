package main

import (
	"server/global"
	"server/initialize"
)

func main() {
	// 初始化mysql
	global.GLOBAL_DB = initialize.InitDB()
	// 初始化redis
	if err := initialize.InitCache(); err != nil {
		return
	}
	r := initialize.InitRouters()
	r.Run()
}
