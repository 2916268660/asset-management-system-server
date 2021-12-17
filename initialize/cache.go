package initialize

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/go-redis/redis"
	"log"
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
		log.Println("加载配置文件失败", err)
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
		log.Println(fmt.Sprintf("redis ping failed, err=%v", err))
		return global.ERRCACHE
	}
	return nil
}
