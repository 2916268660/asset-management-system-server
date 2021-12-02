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
	global.GLOBAL_DB = initialize.InitDB()
	r := initialize.InitRouters()
	r.Run()
}
