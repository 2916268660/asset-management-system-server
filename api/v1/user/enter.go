package user

import "server/logic"

type ApiGroup struct {
	RegisterApi
	LoginApi
	DetailsApi
}

var registerLogic = logic.LogicGroupApp.UserLogic.RegisterLogic
var loginLogic = logic.LogicGroupApp.UserLogic.LoginLogic
var detailsLogic = logic.LogicGroupApp.UserLogic.DetailsLogic
