package routers

import (
	"server/routers/asset"
	"server/routers/user"
	"server/routers/utils"
)

type routersGroup struct {
	UserRouterGroup  user.RouterGroup
	UtilsRouterGroup utils.RouterGroup
	AssetRouterGroup asset.RouterGroup
}

var RoutersGroupApp = new(routersGroup)
