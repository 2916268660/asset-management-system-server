package routers

import (
	"server/routers/asset"
	"server/routers/common"
	"server/routers/user"
)

type routersGroup struct {
	UserRouterGroup  user.RouterGroup
	UtilsRouterGroup common.RouterGroup
	AssetRouterGroup asset.RouterGroup
}

var RoutersGroupApp = new(routersGroup)
