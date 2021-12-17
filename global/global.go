package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	GLOBAL_DB    *gorm.DB
	GLOBAL_CACHE *redis.Client
)
