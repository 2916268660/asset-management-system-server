package main

import (
	"go.uber.org/zap"
	"server/core"
	"server/global"
)

func main() {
	// 初始化日志库
	global.GLOBAL_LOG = core.InitLogger()
	// 初始化mysql
	global.GLOBAL_DB = core.InitDB()
	// 初始化redis
	if err := core.InitCache(); err != nil {
		global.GLOBAL_LOG.Error("redis conn failed", zap.Error(err))
		return
	}
	global.GLOBAL_LOG.Debug("redis conn succeed")
	r := core.InitRouters()
	r.Run()
}
