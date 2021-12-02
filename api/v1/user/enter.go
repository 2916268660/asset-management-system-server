package user

import "server/logic"

type ApiGroup struct {
	RegisterApi
	LoginApi
}

var registerLogic = logic.LogicGroupApp.UserLogic.RegisterLogic
var loginLogic = logic.LogicGroupApp.UserLogic.LoginLogic