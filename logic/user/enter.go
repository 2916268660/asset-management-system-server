package user

import model "server/models/user"

type UserGroup struct {
	RegisterLogic
	LoginLogic
}

var registerModel = model.ModelGroupApp.RegisterModel
var loginModel = model.ModelGroupApp.LoginModel
