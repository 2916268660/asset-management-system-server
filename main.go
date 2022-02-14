package main

import (
	"go.uber.org/zap"
	"server/global"
	"server/initialize"
)

func main() {
	// 初始化日志库
	global.GLOBAL_LOG = initialize.InitLogger()
	// 初始化mysql
	global.GLOBAL_DB = initialize.InitDB()
	// 初始化redis
	if err := initialize.InitCache(); err != nil {
		global.GLOBAL_LOG.Error("redis conn failed", zap.Error(err))
		return
	}
	global.GLOBAL_LOG.Debug("redis conn succeed")
	r := initialize.InitRouters()
	r.Run()
}
