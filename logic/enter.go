package logic

import (
	"server/logic/asset"
	"server/logic/charger"
	"server/logic/common"
	"server/logic/user"
)

type logicGroup struct {
	UserLogic    user.LogicGroup
	CommonLogic  common.LogicGroup
	AssetLogic   asset.LogicGroup
	ChargerLogic charger.LogicGroup
}

var LogicGroupApp = new(logicGroup)
