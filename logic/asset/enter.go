package asset

import "server/repository"

type LogicGroup struct {
	ManagementLogic
}

var userModel = repository.ModelGroupApp.UserModel
var assetModel = repository.ModelGroupApp.AssetMode
