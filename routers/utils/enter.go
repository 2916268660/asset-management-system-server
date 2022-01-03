package utils

import v1 "server/api/v1"

type RouterGroup struct {
	ManagementRouter
}

var utilsApi = v1.ApiGroupApp.ValidateCodeApi
