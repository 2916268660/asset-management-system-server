package cache

import (
	"go.uber.org/zap"
	"server/global"
	"time"
)

// SetKey 保存缓存
func SetKey(key, value string, expire time.Duration) error {
	if err := global.GLOBAL_CACHE.Set(key, value, expire).Err(); err != nil {
		global.GLOBAL_LOG.Error("设置缓存错误", zap.String("key", key), zap.String("val", value), zap.Error(err))
		return err
	}
	return nil
}

// GetKey 获取缓存
func GetKey(key string) (value string, err error) {
	value, err = global.GLOBAL_CACHE.Get(key).Result()
	if err != nil {
		global.GLOBAL_LOG.Error("获取缓存错误", zap.String("key", key), zap.Error(err))
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
