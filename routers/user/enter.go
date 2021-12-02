package user

import v1 "server/api/v1"

type RoutersGroup struct {
	LoginRouters
	RegisterRouters
}

var registerApi = v1.ApiGroupApp.UserApi.RegisterApi