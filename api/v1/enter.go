package v1

import (
	"server/api/v1/asset"
	"server/api/v1/user"
	"server/api/v1/utils"
)

type apiGroup struct {
	UserApi         user.ApiGroup
	ValidateCodeApi utils.ApiGroup
	AssetApi        asset.ApiGroup
}

var ApiGroupApp = new(apiGroup)
