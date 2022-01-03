package asset

import "server/logic"

type ApiGroup struct {
	ManagementApi
}

var assetLogic = logic.LogicGroupApp.AssetLogic.ManagementLogic
