package common

import v1 "server/api/v1"

type RouterGroup struct {
	ManagementRouter
}

var commonApi = v1.ApiGroupApp.CommonApi
var assetApi = v1.ApiGroupApp.AssetApi
