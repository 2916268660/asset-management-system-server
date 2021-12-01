package routers

import "github.com/gin-gonic/gin"

type Router func(r *gin.Engine)

var routers = make([]Router, 0, 8)

func Include(router ...Router) {
	routers = append(routers, router...)
}

func Init() *gin.Engine {
	r := gin.Default()
	for _, router := range routers {
		router(r)
	}
	return r
}
