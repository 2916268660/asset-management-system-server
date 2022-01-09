package cache

import (
	"fmt"
	"log"
	"server/global"
	"time"
)

// SetKey 保存缓存
func SetKey(key, value string, expire time.Duration) error {
	if err := global.GLOBAL_CACHE.Set(key, value, expire).Err(); err != nil {
		log.Println(fmt.Sprintf("set key{%s} value{%s} failed, err=%v", key, value, err))
		return err
	}
	return nil
}

// GetKey 获取缓存
func GetKey(key string) (value string, err error) {
	value, err = global.GLOBAL_CACHE.Get(key).Result()
	if err != nil {
		log.Println(fmt.Sprintf("get key{%s} failed, err=%v", key, err))
		return "", err
	}
	return
}

// IsExistKey 判断key是否存在
func IsExistKey(key string) bool {
	res, err := global.GLOBAL_CACHE.Do("exists", key).Result()
	r := res.(int64)
	if err != nil || r == 0 {
		return false
	}
	return true
}
