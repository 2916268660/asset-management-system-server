package user

import v1 "server/api/v1"

type RoutersGroup struct {
	LoginRouters
	RegisterRouters
	DetailsRouters
}

var registerApi = v1.ApiGroupApp.UserApi.RegisterApi
var loginApi = v1.ApiGroupApp.UserApi.LoginApi
var validateCodeApi = v1.ApiGroupApp.ValidateCodeApi
var detailsApi = v1.ApiGroupApp.UserApi.DetailsApi
