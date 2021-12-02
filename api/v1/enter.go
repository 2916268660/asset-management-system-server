package v1

import "server/api/v1/user"

type apiGroup struct {
	UserApi user.ApiGroup
}

var ApiGroupApp = new(apiGroup)
