package repository

import (
	"server/repository/asset"
	"server/repository/user"
)

type modeGroup struct {
	UserModel user.ManagementModel
	AssetMode asset.ManagementModel
}

var ModelGroupApp = new(modeGroup)
