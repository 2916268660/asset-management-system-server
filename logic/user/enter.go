package user

import (
	"server/models"
)

type LogicGroup struct {
	RegisterLogic
	LoginLogic
	DetailsLogic
}

var userModel = models.ModelGroupApp.UserModel
