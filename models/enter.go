package models

import "server/models/user"

type modeGroup struct {
	UserModel user.UserModel
}

var ModelGroupApp = new(modeGroup)
