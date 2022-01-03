package logic

import (
	"server/logic/asset"
	"server/logic/user"
	"server/logic/utils"
)

type logicGroup struct {
	UserLogic  user.LogicGroup
	UtilsLogic utils.LogicGroup
	AssetLogic asset.LogicGroup
}

var LogicGroupApp = new(logicGroup)
