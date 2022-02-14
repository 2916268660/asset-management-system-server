package common

import "server/logic"

type ApiGroup struct {
	ManagementApi
}

var commonLogic = logic.LogicGroupApp.CommonLogic
