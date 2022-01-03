package models

import (
	"server/models/asset"
	"server/models/user"
)

type modeGroup struct {
	UserModel user.ManagementModel
	AssetMode asset.ManagementModel
}

var ModelGroupApp = new(modeGroup)
