package asset

import v1 "server/api/v1"

type RouterGroup struct {
	ManagementRouter
}

var assetApi = v1.ApiGroupApp.AssetApi
