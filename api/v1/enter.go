package v1

import (
	"server/api/v1/asset"
	"server/api/v1/charger"
	"server/api/v1/common"
	"server/api/v1/user"
)

type apiGroup struct {
	UserApi    user.ApiGroup
	CommonApi  common.ApiGroup
	AssetApi   asset.ApiGroup
	ChargerApi charger.ApiGroup
}

var ApiGroupApp = new(apiGroup)
