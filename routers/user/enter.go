package user

import v1 "server/api/v1"

type RouterGroup struct {
	ManagementRouter
}

var userApi = v1.ApiGroupApp.UserApi
