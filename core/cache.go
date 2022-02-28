package core

import (
	"errors"
	"github.com/go-ini/ini"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"server/global"
)

var (
	addr     string
	password string
	db       int
)

func init() {
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		global.GLOBAL_LOG.Error("加载配置文件失败", zap.Error(err))
		return
	}
	redisCfg := cfg.Section("redis")
	addr = redisCfg.Key("addr").MustString("")
	password = redisCfg.Key("password").MustString("")
	db = redisCfg.Key("db").MustInt(0)
}

func InitCache() error {
	global.GLOBAL_CACHE = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := global.GLOBAL_CACHE.Ping().Result()
	if err != nil {
		global.GLOBAL_LOG.Error("redis ping failed", zap.Error(err))
		return errors.New("初始化redis失败")
	}
	return nil
}
