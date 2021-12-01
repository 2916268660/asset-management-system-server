package main

import (
	"fmt"
	"log"
	userRouter "server/app/user/routers"
	"server/global"
	"server/init/db"
	"server/routers"
)

func main() {
	if err := global.InitTrans("zh"); err != nil {
		log.Println(fmt.Sprintf("init trans failed, err=%v\n", err))
	}
	db.InitDB()
	routers.Include(userRouter.Routers)
	r := routers.Init()

	r.Run()
}
