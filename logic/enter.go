package logic

import "server/logic/user"

type logicGroup struct {
	UserLogic user.UserGroup
}

var LogicGroupApp = new(logicGroup)
