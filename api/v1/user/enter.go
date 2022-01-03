package user

import "server/logic"

type ApiGroup struct {
	ManagementApi
}

var userLogic = logic.LogicGroupApp.UserLogic.ManagementLogic
