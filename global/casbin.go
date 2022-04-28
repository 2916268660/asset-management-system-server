package global

// 用户权限
const (
	User     = "user"
	Charger  = "charger"
	Provider = "provider"
)

var RoleMap = map[string]string{
	User:     "普通用户",
	Charger:  "部门负责人",
	Provider: "资产管理员",
}

var Role = map[string]string{
	"普通用户":  User,
	"部门负责人": Charger,
	"资产管理员": Provider,
}
var PoliciesList = [][]string{
	{
		User,
		"/v1/asset/applyReceive",
		"POST",
	},
	{
		User,
		"/v1/asset/applyRevert",
		"POST",
	},
	{
		User,
		"/v1/asset/applyRepair",
		"POST",
	},
	{
		User,
		"/v1/asset/getAsset",
		"GET",
	},
	{
		User,
		"/v1/asset/getReceiveTodo",
		"GET",
	},
	{
		User,
		"/v1/asset/getRevertTodo",
		"GET",
	},
	{
		User,
		"/v1/asset/getRepairsTodo",
		"GET",
	},
	{
		User,
		"/v1/asset/rollback",
		"POST",
	},
	{
		User,
		"/v1/asset/getTodoDetails",
		"GET",
	},
	{
		User,
		"/v1/asset/getAssetsByUser",
		"GET",
	},
	{
		User,
		"/v1/updateUser",
		"PUT",
	},
	{
		User,
		"/v1/updatePass",
		"PUT",
	},
	{
		Charger,
		"/v1/asset/applyReceive",
		"POST",
	},
	{
		Charger,
		"/v1/asset/applyRevert",
		"POST",
	},
	{
		Charger,
		"/v1/asset/applyRepair",
		"POST",
	},
	{
		Charger,
		"/v1/asset/getAsset",
		"GET",
	},
	{
		Charger,
		"/v1/asset/getReceiveTodo",
		"GET",
	},
	{
		Charger,
		"/v1/asset/getRevertTodo",
		"GET",
	},
	{
		Charger,
		"/v1/asset/getRepairsTodo",
		"GET",
	},
	{
		Charger,
		"/v1/asset/rollback",
		"POST",
	},
	{
		Charger,
		"/v1/asset/getAuditTodo",
		"GET",
	},
	{
		Charger,
		"/v1/asset/getTodoDetails",
		"GET",
	},
	{
		Charger,
		"/v1/asset/getAssetsByUser",
		"GET",
	},
	{
		Charger,
		"/v1/asset/audit",
		"POST",
	},
	{
		Charger,
		"/v1/updateUser",
		"PUT",
	},
	{
		Charger,
		"/v1/updatePass",
		"PUT",
	},
}
