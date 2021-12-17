package validate_code

import "server/logic"

type ApiGroup struct {
	ValidateCodeApi
}

var validateCodeLogic = logic.LogicGroupApp.ValidateCodeLogic
