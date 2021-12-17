package logic

import (
	"server/logic/user"
	"server/logic/validate_code"
)

type logicGroup struct {
	UserLogic         user.LogicGroup
	ValidateCodeLogic validate_code.ValidateCodeLogic
}

var LogicGroupApp = new(logicGroup)
