package user

import (
	"server/repository"
)

type LogicGroup struct {
	ManagementLogic
}

var userModel = repository.ModelGroupApp.UserModel
