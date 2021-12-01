package main

import (
	routers2 "server/app/user/routers"
	routers3 "server/app/work/routers"
	"server/routers"
)

func main() {
	routers.Include(routers2.Routers, routers3.Routers)
	r := routers.Init()
	//r.GET("/*", middleware.JWTAuthMiddleWare())

	r.Run()

}
