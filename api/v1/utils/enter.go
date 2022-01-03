package utils

import "server/logic"

type ApiGroup struct {
	ManagementApi
}

var utilsLogic = logic.LogicGroupApp.UtilsLogic
