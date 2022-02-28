package utils

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"server/global"
	"sync"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func InitCasbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		adapter, _ := gormadapter.NewAdapterByDB(global.GLOBAL_DB)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer("./conf/model.conf", adapter)
		go func() {
			ok, _ := syncedEnforcer.AddPolicies(global.PoliciesList)
			if !ok {
				global.GLOBAL_LOG.Info("权限重复")
			}
		}()
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
