package v1

import (
	"server/api/v1/user"
	"server/api/v1/validate_code"
)

type apiGroup struct {
	UserApi         user.ApiGroup
	ValidateCodeApi validate_code.ApiGroup
}

var ApiGroupApp = new(apiGroup)
