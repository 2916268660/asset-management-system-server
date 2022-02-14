package charger

import "server/logic"

type ApiGroup struct {
	ManagementApi
}

var chargerLogic = logic.LogicGroupApp.ChargerLogic
